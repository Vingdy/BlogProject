package feedback

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"errors"
)

//返回内容处理

//
//type FeedBack struct {
//	DistWriter http.ResponseWriter `json:"-"`
//	Code     int                 `json:"code"`
//	Msg      string              `json:"msg,omitempty"`
//	Data     interface{}         `json:"data,omitempty"`
//}
//func NewFeedBack(w http.ResponseWriter){
//	return &feedBack{DistWriter:w}
//}
//type FbBuilder interface {
//	Dist(w http.ResponseWriter) FbBuilder
//	Code(code int) FbBuilder
//	Msg(msg string) FbBuilder
//	Data(data interface{}) FbBuilder
//	Response() error
//	Clear()
//}

func NewFeedBack(w http.ResponseWriter) *FeedBack {
	return &FeedBack{DistWriter: w}
}

func (f *FeedBack) Dist(w http.ResponseWriter) *FeedBack {
	f.DistWriter = w
	return f
}

func (f *FeedBack) FbCode(code int) *FeedBack {
	f.Code = code
	return f
}

func (f *FeedBack) FbMsg(msg string) *FeedBack {
	f.Msg = msg
	return f
}

func (f *FeedBack) FbTotal(total int) *FeedBack {
	f.Total = total
	return f
}

func (f *FeedBack) FbData(data interface{}) *FeedBack {
	f.Data = data
	return f
}

//func (f *feedBack) Response() (err error) {
//	if f.DistWriter == nil {
//		return errors.New("DistWriter is empty")
//	}
//	buf, _ := json.Marshal(f)
//	fmt.Fprint(f.DistWriter, string(buf))
//	f.Clear()
//	return nil
//}
//func (f *feedBack) Clear() {
//	f.FbData = nil
//	f.FbMsg = ""
//	f.FbCode = 0
//}
//
//func NewFeedBack(w http.ResponseWriter)FeedBack{
//	return &FeedBack{DistWriter:w}
//}
//
//func (f *FeedBack)Dist(w http.ResponseWriter)FeedBack{
//	f.DistWriter=w
//	return f
//}

func (f *FeedBack)Response()(err error){
	//out:=&FeedBack
	if f.DistWriter == nil {
		return errors.New("DistWriter is empty")
	}
	result,err:=json.Marshal(f)
	if err!=nil{
		log.Print(err)
		return
	}
	//w.WriteHeader(code)
	fmt.Fprintln(f.DistWriter,string(result))
	f.Clear()
	return nil
}

func (f *FeedBack) Clear() {
	f.Data = nil
	f.Msg = ""
	f.Code = 0
}
