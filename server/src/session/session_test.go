package session

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"view/testTools"
	"constant"
)

func TestMain(m *testing.M) {
	testTool.SetupTestMain(m, "../../../bin")
}

func TestSetSession(t *testing.T) {
	Init()
	req := httptest.NewRequest("GET", "http://127.0.0.1/api", strings.NewReader(""))
	w := httptest.NewRecorder()

	session := &constant.Session{"testid", 1, "xiaoming","",""}
	err := SetSession(w, req, session)
	if err != nil {
		t.Fatal(err)
	}
	h := w.Header()
	t.Log(h["Set-Cookie"])
}

func TestGetSession(t *testing.T) {
	Init()
	req := httptest.NewRequest("GET", "http://127.0.0.1/api", strings.NewReader(""))
	cookie := http.Cookie{Name: cookieName, Value: url.QueryEscape("7086e16b768ffebf51ac12f643045d52")}
	req.AddCookie(&cookie)

	session, err := GetSession(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(session)
}

func TestDestroySession(t *testing.T) {
	Init()
	req := httptest.NewRequest("GET", "http://127.0.0.1/api", strings.NewReader(""))
	cookie := http.Cookie{Name: cookieName, Value: url.QueryEscape("23c4be43599911a00f2a353eaedf1951")}
	req.AddCookie(&cookie)

	err := DestroySession(req)
	if err != nil {
		t.Fatal(err)
	}
}
