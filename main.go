package main

import (
	"go_learn/funcs"
	httpservice "go_learn/http_service"
)

func main() {
	funcs.InitMinio()
	httpservice.Start()
}
