package rabbit

import (
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"rabbitmq_mate_go/pkg/path"
	"rabbitmq_mate_go/pkg/util"
	"time"
)

func Start() error {
	stdout, stderr, err := gutil.CallScript(path.RabbitStartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	if err != nil {
		util.Logger().Error("command wait error ", zap.Error(err))
		return err
	}
	time.Sleep(10 * time.Second)
	stdout, stderr, err = gutil.CallScript(path.RabbitInitUserScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	if err != nil {
		util.Logger().Error("rabbit init user failed")
	}
	return err
}
