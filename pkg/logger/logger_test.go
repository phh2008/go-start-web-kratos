package logger

import (
	"com.gientech/selection/pkg/config"
	"fmt"
	"go.uber.org/zap"
	"testing"
)

// TestZap zap 日志框架
func TestZap(t *testing.T) {
	var config = config.NewConfig("../../config")
	zapLog := newZapLog(config)
	zapLog.Debug("debug message")
	zapLog.Info("info message")
	zapLog.Warn("warn message")
	zapLog.Error("error message")
}

func TestLogger(t *testing.T) {
	var config = config.NewConfig("../../config")
	NewLogger(config)
	S().Debugf("debug message")
	S().Infof("info message")
	S().Warnf("warn message")
	S().Errorf("error message:%s", "this is message")
	S().Info("msg", "a", "b", "c")
	zap.S().Infof("xxxxxxxxxx:%s", "hello")
}

func TestWrapLogger(t *testing.T) {
	var config = config.NewConfig("../../config")
	NewLogger(config)

	Debugf("wrap debugF message %s %s %s", "aa", "bb", "cc")
	Infof("wrap infoF message")
	Warnf("wrap warnF message")
	Errorf("wrap errorF message:%s", "this is message")
	fmt.Println("------------------------------------------------------")
	Debugs("wrap debugs message %s %s %s ", "aa", "bb", "cc")
	Infos("wrap infos message")
	Warns("wrap warns message")
	Errors("wrap errors message:%s", "this is message")
	fmt.Println("------------------------------------------------------")
	Debug("wrap debug message")
	Info("wrap info message")
	Warn("wrap warn message")
	Error("wrap error message", zap.String("url", "www.baidu.com"), zap.String("ip", "192.168.1.200"))
}
