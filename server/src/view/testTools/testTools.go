package testTool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"db"
	"cache"
	"logger"
	"constant"
)

type FB struct {
	FbCode int         `json:"code"`
	FbMsg  string      `json:"msg,omitempty"`
	FbData interface{} `json:"data,omitempty"`
}

// Case 表示一个测试样例
type Case struct {
	Method   string                                       //http请求方法
	RawQuery string                                       //Get方法携带参数
	Data     string                                       //POST方法携带数据
	HTTPFunc func(w http.ResponseWriter, r *http.Request) //测试函数
	WantCode int                                          //期望返回值
	UserInfo *constant.Session                         //用户信息
	FeedBack *FB
}

// CheckErrCode 检查API接口返回值
//func (c *Case) MockRequest() error {
////
////	req := httptest.NewRequest(c.Method, "http://127.0.0.1/api", strings.NewReader(c.Data))
////	w := httptest.NewRecorder()
////
////	req.URL.RawQuery = c.RawQuery
////	c.HTTPFunc(w, req.WithContext(context.WithValue(req.Context(), "ctx", constant.ContextContent{UserInfo: c.UserInfo})))
////
////	resp := w.Result()
////	body, err := ioutil.ReadAll(resp.Body)
////	if err != nil {
////		return fmt.Errorf("read result body failed: %v", err)
////	}
////	defer resp.Body.Close()
////
////	//skip check code
////	if c.WantCode == 0 {
////		return nil
////	}
////
////	if c.FeedBack == nil {
////		var fb FB
////		if err = json.Unmarshal(body, &fb); err != nil {
////			return fmt.Errorf("unmarshal feedback data failed: %v", err)
////		}
////
////		if fb.FbCode != c.WantCode {
////			return fmt.Errorf("want %d but got %d", c.WantCode, fb.FbCode)
////		}
////	} else {
////		if err = json.Unmarshal(body, &c.FeedBack); err != nil {
////			return fmt.Errorf("unmarshal feedback data failed: %v", err)
////		}
////		if c.FeedBack.FbCode != c.WantCode {
////			return fmt.Errorf("want %d but got %d", c.WantCode, c.FeedBack.FbCode)
////		}
////	}
////	return nil
////}

func ClearAndInitData(proDir string) {
	if proDir == "" {
		proDir = "../../../../bin"
	}
	fmt.Println("Clear TestData Before Test")
	ClearTestData()

	fmt.Println("Init TestData Before Test")

	err := InitTestDataWithFile(proDir, "testinitdata")
	if err != nil {
		panic(err)
	}
	err = InitTestDataWithFile(proDir, "testInitData_chy")
	if err != nil {
		panic(err)
	}
}

func SetupTestMain(m *testing.M, proDir string) {
	if proDir == "" {
		proDir = "../../../../bin"
	}

	fmt.Println("Setup Env")
	SetupTestEnv(proDir)

	ClearAndInitData(proDir)

	iRsl := m.Run()

	fmt.Println("Clear TestData After Test")
	ClearTestData()

	fmt.Println("Test Complete")
	os.Exit(iRsl)
}

//准备测试环境
func SetupTestEnv(proDir string) {
	//var err error
	//conf.Init(proDir, "DEV")
	if db.Db == nil {
		db.InitDB("localhost","5432", "postgres", "VingB2by","test","postgres")
	}
	logger.InitLogger(proDir)

	cache.Init()
}

func InitTestDataWithFile(proDir, fileName string) error {
	if proDir == "" || fileName == "" {
		return errors.New("proDir or fileName empty")
	}
	if db.Db == nil {
		return errors.New("Db nil")
	}
	data, err := ioutil.ReadFile(proDir + "/testData/" + fileName)
	if err != nil {
		return err
	}
	_, err = db.Db.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}

func ClearTestData() error {
	_, err := db.Db.Exec(`
		DELETE FROM event WHERE title LIKE '__T_%';
		DELETE FROM mgr WHERE account LIKE '__T_%';
		DELETE FROM mgr WHERE name LIKE '__T_%';
		DELETE FROM record WHERE openid LIKE '__T_%';
		DELETE FROM best WHERE openid LIKE '__T_%';
		DELETE FROM player WHERE openid LIKE '__T_%';
		DELETE FROM actionlog WHERE operator LIKE '__T_%';
		DELETE FROM stats WHERE openid LIKE '__T_%' or eventid between 10000 AND 10020;
	`)
	if err != nil {
		return err
	}
	return nil
}
