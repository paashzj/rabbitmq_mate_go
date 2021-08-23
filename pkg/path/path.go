package path

import (
	"os"
	"path/filepath"
)

// rabbitmq
var (
	RabbitHome = os.Getenv("RABBITMQ_HOME")
)

// mate
var (
	RabbitMatePath       = filepath.FromSlash(RabbitHome + "/mate")
	RabbitScripts        = filepath.FromSlash(RabbitMatePath + "/scripts")
	RabbitStartScript    = filepath.FromSlash(RabbitScripts + "/start-rabbitmq.sh")
	RabbitInitUserScript = filepath.FromSlash(RabbitScripts + "/init-rabbit-user.sh")
)
