package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
	"strings"
)

// GetProDir 用于获取项目根目录
func GetProDir() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	end := strings.LastIndex(path, string(os.PathSeparator))
	proPath := path[:end]
	return proPath, nil
}

func round(num float64) int {
	//留整数部分 四舍五入
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	//将要保留的小数通过 x10 放到整数部分 去小数,再/10
	return float64(round(num*output)) / output
}

func HandlePanic(proDir string) {
	if err := recover(); err != nil {
		formatErr := fmt.Sprintf("%s\t%v", time.Now().Local(), err)
		stack := make([]byte, 1024*10)
		if len(stack) == 0 {
			fmt.Println("stack length = 0")
			return
		}
		runtime.Stack(stack, true)
		stack = bytes.Replace(stack, []byte("\u0000"), []byte(""), -1)
		fmt.Printf("%s\n%s", formatErr, stack)
		ioutil.WriteFile(proDir+"/panic", bytes.Join([][]byte{[]byte(formatErr), stack}, []byte("\n")), 0644)
	}
}

// 获取多个url参数 map版本
func GetURLParams(r *http.Request, key ...string) map[string]string {
	urlParams := make(map[string]string)
	for _, param := range key {
		urlParams[param] = r.URL.Query().Get(param)
	}
	return urlParams
}

func Mkdir(dir string) error{
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0770)
	} else {
		return err
	}
	return nil
}