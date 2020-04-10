package session

import (
	"encoding/hex"
	"net/http"
	"constant"
	"net/url"
	"fmt"
	"bytes"
	"encoding/gob"
	"cache"
	"time"
	"crypto/rand"
	"github.com/go-redis/redis"
	"conf"
)

//session管理器结构体
var (
	cookieName  string//cookie名字
	maxLifeTime int//cookie最大生存时间
	path        string//路径
	httpOnly    bool//是否只能http协议
	maxAge      int//
)

const (
	REDIS_SESSION     = "SESS_"
)

func Init() {
	cookieName = conf.App.SessionKey
	maxLifeTime = conf.App.MaxLifeTime
	path = conf.App.Path
	httpOnly = conf.App.HTTPOnly
	maxAge = conf.App.MaxAge
}

func SetSession(w http.ResponseWriter, r *http.Request, session *constant.Session) error {
	var sid string
	//fmt.Println("test3")
	Init()
	cookie, err := r.Cookie(cookieName)

	if err != nil && err != http.ErrNoCookie {
		return err
	}
	if err == http.ErrNoCookie || cookie.Value == "" {
		sid = randomSID()
		newCookie := http.Cookie{
			Name:     cookieName,
			Value:    url.QueryEscape(sid),
			Path:     path,
			HttpOnly: httpOnly,
			MaxAge:   maxAge,
		}
		http.SetCookie(w, &newCookie)
	} else {
		sid, err = url.QueryUnescape(cookie.Value)
		if err != nil {
			return fmt.Errorf("query cookie value failed: %v", err)
		}
	}
	buffer := new(bytes.Buffer)
	if err := gob.NewEncoder(buffer).Encode(session); err != nil {
		return fmt.Errorf("encode session failed: %v", err)
	}
	if err = cache.Cache.Set(REDIS_SESSION+sid, buffer.Bytes(), time.Duration(maxLifeTime)*time.Second).Err();err != nil {
		return fmt.Errorf("redis set session failed: %v", err)
	}
	return nil
}

func GetSession(r *http.Request) (sess *constant.Session, err error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil
		}
		return nil, err
	}
	if cookie.Value == "" {
		return nil, nil
	}
	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return nil, err
	}

	data, err := cache.Cache.Get(REDIS_SESSION + sid).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("redis get session failed: %v", err)
	}
	sess = &constant.Session{}
	if err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(sess); err != nil {
		return nil, fmt.Errorf("decode session failed: %v", err)
	}
	if err = cache.Cache.Expire(REDIS_SESSION+sid, time.Duration(maxLifeTime)*time.Second).Err(); err != nil {
		return nil, fmt.Errorf("redis expire failed: %v", err)
	}
	sess.UserName = "test"
	sess.Role = 10
	fmt.Println(sess)
	return sess, nil
}

func DestroySession(r *http.Request) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return nil
		}
		return err
	}
	sid, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return err
	}
	return cache.Cache.Del(REDIS_SESSION + sid).Err()
}

func randomSID() string {
	sid := make([]byte, 16)
	rand.Read(sid)
	return hex.EncodeToString(sid)
}