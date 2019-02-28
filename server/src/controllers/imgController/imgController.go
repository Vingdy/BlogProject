package imgController

import (
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"errors"
	"utils"
)

var file = make(map[string]bool)
type Pic struct {
	Link string `json:"link"`
}

func UploadPic(w http.ResponseWriter,r *http.Request) {

	//上传文件类型规定
	TypeProvision()

	//获取文件
	File, FileHeader, err := r.FormFile("file")
	if err!=nil{
		msg:="获取前端图片失败"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}
	defer File.Close()

	//获取文件格式
	//以点分隔切片
	Slice := strings.Split(FileHeader.Filename, ".")
	//获取切片的数量
	SliceSum := len(Slice)
	//倒数第一个是文件类型
	fileSuffix := strings.ToLower(Slice[SliceSum-1])

	//判断文件大小
	if FileHeader.Size/1024 <= 10240 {
		if !file[fileSuffix]{
			msg:="图片格式不符合要求"
			log.Println(msg)
			log.Println(err)
			//data:=model.FeedBackErrorHandle(601,msg)
			//fmt.Fprintln(w,string(data))
			return
		}
	}else{
		msg:="图片大小不符合要求"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}

	//读取文件
	FileByte, err := ioutil.ReadAll(File)
	if err!=nil{
		msg:="读取字节错误"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}

	//根据文件字节作为文件标识命名
	FileId, err := GetFileMD5(FileByte)
	if err!=nil{
		msg:="图片名字加密错误"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}

	//规定文件标识长度
	CreateFileid,err:=Substr(FileId,0,250);
	if err!=nil{
		msg:="命名图片错误"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}

	//获取可执行文件的绝对路径
	PATH,err:=exepath()

	//kyoma2187
	//shuaibirunfa
	//创建存放文件的文件夹
	err = utils.Mkdir(PATH)
	if err!=nil{
		msg:="创建存放文件的文件夹失败"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}

	//创建文件
	err = createfile(PATH+CreateFileid+"."+fileSuffix, FileByte)
	//fmt.Println(PATH)
	if err!=nil{
		msg:="创建文件错误"
		log.Println(msg)
		log.Println(err)
		//data:=model.FeedBackErrorHandle(601,msg)
		//fmt.Fprintln(w,string(data))
		return
	}

	//返回前端上传图片的地址
	pic:=Pic{}

	pic.Link="http://localhost:8080/static/"+CreateFileid+"."+fileSuffix
	data,_:=json.Marshal(pic)

	w.Write(data)
	msg:="上传图片成功"
	log.Println(msg)
	//log.Println(err)
	//data:=model.FeedBackSuccessHandle(600,msg,datas)
	//fmt.Fprintln(w,string(data))
	return
}
//上传图片类型规定
func TypeProvision() {
	file["jpg"]  = true
	file["jpeg"] = true
	file["png"]  = true
	file["gif"]  = true
}
//MD5加密
func GetFileMD5(fbyte []byte) (string, error) {
	hash := md5.New()
	_,err:=hash.Write(fbyte)
	if err!=nil{
		return "",err
	}
	result := hash.Sum(nil)
	return hex.EncodeToString(result), nil
}
//创建文件
func createfile(fileurl string, filebyte []byte) error {
	file, err := os.Create(fileurl)
	if err != nil {
		return err
	}
	_, err = file.Write(filebyte)
	if err != nil {
		return err
	}
	return nil
}
//获取可执行文件的绝对路径
func exepath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	dir += "\\files"
	return strings.Replace(dir, "\\", "/", -1) + "/", nil
}
//截取字符串
func Substr(str string, start int, length int) (string,error) {
	RuneStr := []rune(str)
	RuneLen := len(RuneStr)
	if RuneLen==0{
		return "",errors.New("文件字节长度为零")
	}
	if start < 0 {
		start = 0
	}
	if length==0{
		return "",errors.New("文件名长度为零")
	}
	if length<0{
		return "",errors.New("文件名长度为负数")
	}

	end := start + length

	if start > RuneLen {
		return "",errors.New("文件名取值起始位置越界")
	}

	if end > RuneLen {
		end = RuneLen
	}

	return string(RuneStr[start:end]),nil
}
