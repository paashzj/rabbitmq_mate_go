package main

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"rabbitmq_mate_go/pkg/rabbit"
	"rabbitmq_mate_go/pkg/util"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	err := rabbit.Start()
	if err != nil {
		util.Logger().Error("rabbit mq server init failed ", zap.Error(err))
		os.Exit(1)
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}