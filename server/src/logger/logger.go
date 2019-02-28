package logger

import (
	"time"


	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
)

//const (
//	//项目内日志文件路径
//	fileName string = "server\\src\\logs\\log.log"
//)
//
//var Logger = GetLogger()
//
////type Log struct {
//////	startAt time.Time
//////	conText *gin.Context
//////	writer  responseWriter
//////	error   error
//////
//////	Level     string
//////	Time      string
//////	ClientIp  string
//////	Uri       string
//////	ParamGet  url.Values `json:"pGet"`
//////	ParamPost url.Values `json:"pPost"`
//////	RespBody  string
//////	TimeUse   string
//////}
////获取项目路径
//func GetProjectPath() string {
//	path, _ := os.Getwd()
//	pathArr := strings.SplitN(path, "src", -1)
//	return pathArr[0]
//}
//
////获取日志
//func GetLogger() *log.Logger {
//	var logFilePath string = GetProjectPath() + "\\" + fileName;
//	logFile, err := os.OpenFile(logFilePath, os.O_RDWR | os.O_CREATE, os.ModePerm)
//	if err != nil {
//		log.Panic(err)
//		fmt.Println(logFilePath)
//	}
//	logFile.Seek(0, os.SEEK_END)
//	//3个参数：日志文件，每行信息的前缀，信息携带内容
//	return log.New(logFile,"", log.Ldate | log.Ltime)
//}
//
//func OutPut() string {
//	pc, file_name, line_num, ok := runtime.Caller(2)
//	if !ok { return "" }
//	func_name := runtime.FuncForPC(pc).Name()//package.functionName
//	short_func_name := func_name[strings.LastIndex(func_name, ".")+1:]//functionName
//	short_file_name := file_name[strings.LastIndex(file_name, "/")+1:]//fileName.go
//	str := fmt.Sprintf("%s:%d %s():",short_file_name, line_num, short_func_name)
//	return str
//}
//
//func Error(arg_println ...interface{}) {
//	fmt.Println(OutPut(),"[Error]",arg_println)
//	Logger.Println(OutPut(),"[Error]", arg_println)
//}
//
//func Info(arg_println ...interface{}) {
//	fmt.Println(OutPut(),"[Info]",arg_println)
//	Logger.Println(OutPut(),"[Info]", arg_println)
//}
//
//func Debug(w http.ResponseWriter,r *http.Request,arg_println ...interface{}) {
//	fmt.Println(OutPut(),"[Debug]",arg_println)
//	Logger.Println(OutPut(),"[Debug]", arg_println)
//}
//
//func Warn(arg_println ...interface{}) {
//	fmt.Println(OutPut(),"[Warn]",arg_println)
//	Logger.Println(OutPut(),"[Warn]", arg_println)
//}
//
//var (
//	green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
//	white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
//	yellow       = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
//	red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
//	blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
//	magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
//	cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
//	reset        = string([]byte{27, 91, 48, 109})
//	disableColor = false
//)
//
//// LogFormatterParams is the structure any formatter will be handed when time to log comes
//type LogFormatterParams struct {
//	Request *http.Request
//
//	// TimeStamp shows the time after the server returns a response.
//	TimeStamp time.Time
//	// Latency is how much time the server cost to process a certain request.
//	Latency time.Duration
//	// Method is the HTTP method given to the request.
//	Method string
//	// Path is a path the client requests.
//	Path string
//}
//
//// defaultLogFormatter is the default log format function Logger middleware uses.
//func DefaultLogFormatter(w http.ResponseWriter,r *http.Request,starttime time.Time,endtime time.Duration,param LogFormatterParams) string {
//	var methodColor, resetColor string
//	param.Method=r.Method
//	param.Path = r.URL.Path
//	//raw := r.URL.RawQuery
//	//client := &http.Client{}
//	//response,err:=client.Do(r.RequestURI)
//	//fmt.Println("127.0.0.1:8080"+r.RequestURI)
//	//u, _ := url.Parse("127.0.0.1:4200"+r.RequestURI)
//	//q := u.Query()
//	//u.RawQuery = q.Encode()
//	//res, _ := http.Get(u.String())
//	//resCode := res.StatusCode
//	//res.Body.Close()
//	//fmt.Printf("%d\r\n", resCode)
//	//param.StatusCode = http.StatusBadRequest
//	//w.WriteHeader(http.StatusBadRequest)
//	//param.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String(
//				param.TimeStamp = starttime
//				param.Latency = endtime
//	return fmt.Sprintf("%v | %13v |%s %-7s %s %s\n%s",
//		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
//		//statusColor, param.StatusCode, resetColor,
//		param.Latency,
//		//param.ClientIP,
//		methodColor, param.Method, resetColor,
//		param.Path,
//	)
//}

var Logger *zap.Logger
var LogFile *os.File
var log *zap.SugaredLogger

//const (
//	DebugLevel Level = iota - 1
//
//	InfoLevel
//
//	WarnLevel
//
//	ErrorLevel
//
//	DPanicLevel
//
//	PanicLevel
//
//	FatalLevel
//)

var logLevel = zap.NewAtomicLevel()

func InitLogger(prodir string)error{
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  prodir,
		MaxSize:   1024, //MB
		LocalTime: true,
		Compress:  true,
	})
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		logLevel,
	)
	var err error
	LogFile, err = os.OpenFile(prodir+"\\logs\\log.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0333)
	if err != nil {
		return fmt.Errorf("failed to open tracefile.log: %v",err)
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar()
	logEncoderCfg := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Name",
		CallerKey:      "Caller",
		MessageKey:     "Message",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeTime:     TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	logConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    logEncoderCfg,
		OutputPaths:      []string{"server/bin/logs/log.log"},//记录系统日志
		ErrorOutputPaths: []string{"server/bin/logs/log.log"},//记录日志系统的错误
	}
	//fmt.Println(prodir)
	Logger,err= logConfig.Build()
	if err!=nil{
		return err
	}
	return nil
}

func Close(){
	Logger.Sync()
	LogFile.Close()
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString( t.Format("2006-01-02 15:04:05"))
}