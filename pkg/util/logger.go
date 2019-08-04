package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goCart/pkg/setting"
	"os"
	"path"
	"sync"
	"time"
)

type logger struct {
}

var lock *sync.RWMutex

func Logger() *logger {
	return &logger{}
}
func getLoggerInstance() *logrus.Logger {
	//日志文件
	fileName := path.Join(fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath), fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	))
	_, err := os.Stat(fileName)
	if err != nil {
		lock.Lock()
		f, _ := os.Create(fileName)
		_ = f.Close()
		lock.Unlock()
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//写入文件
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	//实例化
	//设置输出
	logger.Out = src

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger
}
func (log *logger) Info(args ...interface{}) {
	getLoggerInstance().SetLevel(logrus.InfoLevel)
	fmt.Println(args)
	getLoggerInstance().Infof("| %v |", args)
}
func (log *logger) Error(args ...interface{}) {
	getLoggerInstance().SetLevel(logrus.ErrorLevel)
	fmt.Println(args)
	getLoggerInstance().Errorf("| %v |", args)
}
func (log *logger) Warn(args ...interface{}) {
	getLoggerInstance().SetLevel(logrus.WarnLevel)
	fmt.Println(args)
	getLoggerInstance().Warnf("| %v |", args)
}
func (log *logger) Debug(args ...interface{}) {
	getLoggerInstance().SetLevel(logrus.DebugLevel)
	fmt.Println(args)
	getLoggerInstance().Debugf("| %v |", args)
}
