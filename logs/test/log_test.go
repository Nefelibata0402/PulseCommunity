package test

import (
	"fmt"
	"go.uber.org/zap"
	"pulseCommunity/logs/test/config"
	"testing"
	"time"
)

func TestWriteLog(t *testing.T) {
	config.InitConfig()
	go func() {
		// 为了演示 ELK，我直接输出日志
		ticker := time.NewTicker(3 * time.Second)
		for range ticker.C {
			zap.L().Info("模拟输出日志")
			fmt.Println("yes")
		}
	}()
	for {

	}
}
