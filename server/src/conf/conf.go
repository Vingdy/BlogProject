package conf

import (
	"os"

	"time"

	"github.com/BurntSushi/toml"
	"constant"
)

type srvFile struct {
	AppEnv string

	DBDriver   string `toml:"DBDriver"`
	DBHost     string `toml:"DBHost"`
	DBPort     string `toml:"DBPort"`
	DBUser     string `toml:"DBUser"`
	DBPassword string `toml:"DBPassword"`
	DBName     string `toml:"DBName"`
	DBTestName string `toml:"DBTestName"`

	RedisHost string `toml:"RedisHost"`
	RedisPort string `toml:"RedisPort"`
	RedisPwd  string `toml:"RedisPwd"`
	RedisDB   int    `toml:"RedisDB"`

	ServerHost     string `toml:"serverHost"`
	ServerPort     string `toml:"serverPort"`
	ServerTestPort string `toml:"serverTestPort"`

	WebPageHost     string `toml:"WebPageHost"`
	WebPageTestHost string `toml:"WebPageTestHost"`

	//log
	OutPutPath    string `toml:"outPutPath"`
	ErrOutPutPath string `toml:"errOutPutPath"`

	//文件有效后缀
	ValidSuffix []string `toml:"suffixes"`

	//文件最大尺寸 M为单位
	MaxFileSize int64 `toml:"MaxFileSize"`

	FileDownloadHost string `toml:"FileDownloadHost"`

	FileDownloadPrefix string `toml:"FileDownloadPrefix"`

	LoggerBuffer int `toml:"LoggerBuffer"`

	LoggerBufDuration int64 `toml:"LoggerBufDuration"`

	SessionKey  string `toml:"SessionKey"`
	MaxLifeTime int    `toml:"MaxLifeTime"`
	Path        string `toml:"Path"`
	HTTPOnly    bool   `toml:"HTTPOnly"`
	MaxAge      int    `toml:"MaxAge"`

	SignKey string `toml:"SignKey"`

	SignSwitch bool `toml:"SignSwitch"`

	LocationCacheStay time.Duration `toml:"LocationCacheStay"`

	TemplateDir string
	ProDir      string

	LogDir  string
	FileDir string

	ExePath string
}

var App *srvFile

func init() {
	App = new(srvFile)
}

func Init(proDir string, buildType string) {
	var confFileName string
	if buildType == constant.BUILD_TYPE_PROD {
		confFileName = "app-prod.toml"
	} else {
		confFileName = "app-dev.toml"
	}
	_, err := toml.DecodeFile(proDir+"/conf-file/"+confFileName, App)
	if err != nil {
		panic(err)
	}

	App.ProDir = proDir
	App.TemplateDir = proDir + "/dist/"

	App.LogDir = proDir + "/logs/"
	App.FileDir = proDir + "/files/"

	App.AppEnv = buildType

	err = mkdirProDir()
	if err != nil {
		panic("mkdirProDir failed " + err.Error())
	}
}

func mkdirProDir() error {
	//创建log目录在项目目录下
	if _, err := os.Stat(App.LogDir); os.IsNotExist(err) {
		os.Mkdir(App.LogDir, 0770)
	} else {
		return err
	}

	//创建file目录
	if _, err := os.Stat(App.FileDir); os.IsNotExist(err) {
		os.Mkdir(App.FileDir, 0770)
	} else {
		return err
	}

	return nil
}

