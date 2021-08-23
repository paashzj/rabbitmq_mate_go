package rabbit

import (
	"rabbitmq_mate_go/pkg/path"
	"rabbitmq_mate_go/pkg/util"
	"time"
)

func Start() error {
	err := util.CallScript(path.RabbitStartScript)
	if err != nil {
		util.Logger().Error("start rabbit failed")
		return err
	}
	time.Sleep(5 * time.Second)
	err = util.CallScript(path.RabbitInitUserScript)
	if err != nil {
		util.Logger().Error("rabbit init user failed")
	}
	return err
}
