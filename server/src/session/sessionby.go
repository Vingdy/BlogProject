package session

import (
	"sync"
	"time"
	"net/http"
	"net/url"
	_"crypto/rand"
)

//***
//要搞懂为什么存内存要带锁，存redis不用锁
//***

//var sessionManager *Helper.SessionManager = nil //session管理器
//sessionManager = Helper.NewSessionManager("TestCookieName", 3600)
//Session管理器
type SessionManager struct {
	CookieName string
	Lock sync.Mutex
	MaxLifeTime int64

	Sessions map[string] *Session
}

type Session struct {
	SessionID        string                      //唯一id
	LastTimeAccessed time.Time                   //最后访问时间
	Values           map[interface{}]interface{} //其它对应值(保存用户所对应的一些值，比如用户权限之类)
}

//创建Session管理器
func NewSessionManager(cookieName string,maxLifeTime int64)*SessionManager{
	NewManager:=&SessionManager{
		CookieName:cookieName,
		MaxLifeTime: maxLifeTime,
		Sessions: make(map[string]*Session)}
	go NewManager.GC()//定时回收
	return NewManager
}
//创建Session
func (manager *SessionManager) SetSession(w http.ResponseWriter, r *http.Request) string {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	//无论原来有没有，都重新创建一个新的session
	newSessionID := url.QueryEscape(randomSID())

	//存指针
	var session *Session = &Session{SessionID: newSessionID, LastTimeAccessed: time.Now(), Values: make(map[interface{}]interface{})}
	manager.Sessions[newSessionID] = session
	//让浏览器cookie设置过期时间
	cookie := http.Cookie{Name: manager.CookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(manager.MaxLifeTime)}
	http.SetCookie(w, &cookie)
	//buffer:=new(bytes.Buffer)
	//if err := gob.NewEncoder(buffer).Encode(session); err != nil {
	//	fmt.Sprint("encode session failed: %v",err)
	//	return ""
	//}
	//if err := cache.Cache.Set("SESS_"+newSessionID, buffer.Bytes(), time.Duration(manager.MaxLifeTime)*time.Second).Err(); err != nil {
	//	fmt.Sprint("redis set session failed: %v",err)
	//	return ""
	//}
	return newSessionID
}
//GC回收
func (manager *SessionManager) GC() {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	for sessionID, session := range manager.Sessions {
		//删除超过时限的session
		if session.LastTimeAccessed.Unix()+manager.MaxLifeTime < time.Now().Unix() {
			delete(manager.Sessions, sessionID)
		}
	}

	//定时回收
	time.AfterFunc(time.Duration(manager.MaxLifeTime)*time.Second, func() {manager.GC() })
}

//结束Session
func (manager *SessionManager) EndSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.Lock.Lock()
		defer manager.Lock.Unlock()

		delete(manager.Sessions, cookie.Value)

		//让浏览器cookie立刻过期
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.CookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}
//结束session
func (manager *SessionManager) EndSessionBy(sessionID string) {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	delete(manager.Sessions, sessionID)
}

//设置session里面的值
func (manager *SessionManager) SetSessionVal(sessionID string, key interface{}, value interface{}) {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	if session, ok := manager.Sessions[sessionID]; ok {
		session.Values[key] = value
	}
}

//得到session里面的值
func (manager *SessionManager) GetSessionVal(sessionID string, key interface{}) (interface{}, bool) {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	if session, ok := manager.Sessions[sessionID]; ok {
		if val, ok := session.Values[key]; ok {
			return val, ok
		}
	}

	return nil, false
}

//得到sessionID列表
func (manager *SessionManager) GetSessionIDList() []string {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	sessionIDList := make([]string, 0)

	for k, _ := range manager.Sessions {
		sessionIDList = append(sessionIDList, k)
	}

	return sessionIDList[0:len(sessionIDList)]
}

//判断Cookie的合法性（每进入一个页面都需要判断合法性）
func (manager *SessionManager) CheckCookieValid(w http.ResponseWriter, r *http.Request) string {
	var cookie, err = r.Cookie(manager.CookieName)

	if cookie == nil ||
		err != nil {
		return ""
	}

	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	sessionID := cookie.Value

	if session, ok := manager.Sessions[sessionID]; ok {
		session.LastTimeAccessed = time.Now() //判断合法性的同时，更新最后的访问时间
		return sessionID
	}

	return ""
}

//更新最后访问时间
func (manager *SessionManager) GetLastAccessTime(sessionID string) time.Time {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()

	if session, ok := manager.Sessions[sessionID]; ok {
		return session.LastTimeAccessed
	}

	return time.Now()
}
//
//func randomSID() string {
//	sid := make([]byte, 16)
//	rand.Read(sid)
//	return hex.EncodeToString(sid)
//}

//Controller用
//sessionMgr := session.NewSessionManager("TestCookieName", 3600)
//sessionID := sessionMgr.SetSession(w, r)
//sessionMgr.SetSessionVal(sessionID, "UserInfo", "nonononono")