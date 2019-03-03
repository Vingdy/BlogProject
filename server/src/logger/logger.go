package logger

import (
	"time"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"runtime"
	"strconv"
	"path"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"utils"
)

var loggerfile = logrus.New()

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) string {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+"%Y%m%d%H%M"+".log",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),        // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		loggerfile.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}


	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%file%:%line% [%lvl%] %time% - %msg%\n",
			Skip:11,
		})
	loggerfile.AddHook(lfHook)
	return writer.CurrentFileName()
}

func InitLogger(prodir string)(err error){
	err=utils.Mkdir(prodir+"/logs")
	if err!=nil{
		return err
	}
	fileName:=ConfigLocalFilesystemLogger(prodir+"/logs", "log", time.Hour*24*365, time.Hour*24)

	loggerfile.Formatter=&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%file%:%line% [%lvl%] %time% - %msg%\n",
		Skip:8,
	}

	loggerfile.Level=logrus.DebugLevel
	//loggerfile.Out=os.Stderr
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0333)
	//file, err := os.OpenFile(prodir+"/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0333)
	if err == nil {
		loggerfile.Out = file
	}

	return nil
}

func Info(msg string){
	loggerfile.Infoln(msg)
}

func Warn(msg string){
	loggerfile.Warnln(msg)
}

func Debug(msg string){
	loggerfile.Debugln(msg)
}

func Fatal(msg string){
	loggerfile.Fatalln(msg)
}

func Error(msg string){
	loggerfile.Errorln(msg)
}

const (
	defaultLogFormat       = "[%lvl%]: %time% - %msg% %%"
	defaultTimestampFormat = time.RFC3339
)

type Formatter struct {
	// Timestamp format
	TimestampFormat string
	LogFormat string
	Skip int
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)
	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)
	//entry.Line
	_, file, line, ok := runtime.Caller(f.Skip)
	if ok {
		output = strings.Replace(output, "%file%", file, 1)
		output = strings.Replace(output, "%line%", strconv.Itoa(line), 1)
	}
	for k, v := range entry.Data {
		if s, ok := v.(string); ok {
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}
	return []byte(output), nil
}


