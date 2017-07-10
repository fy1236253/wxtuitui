package g

import (
	"fmt"
	"log"
	"os"
)

const (
	// VERSION 版本号
	VERSION = "wechatv1 0.1.0"
)

// Root 获取当前路径
var Root string

// InitRootDir 初始化路径
func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	if err != nil {
		log.Fatalln("getwd fail:", err)
	}
}
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
		panic(fmt.Sprintf("%s: %s", err, msg))
	}
}
