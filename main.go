package main

import (
	"github.com/jesusslim/slimgo"
	"github.com/slimgotest/controller"
	"github.com/slimgotest/controller/Admin"
	_ "github.com/slimgotest/hook"
	_ "github.com/slimgotest/model"
	_ "github.com/slimgotest/task"
)

func main() {
	slimgo.SlimApp.Handerlers.Register(&controller.IndexController{}, &controller.TestController{})
	slimgo.SlimApp.Handerlers.Register(&Admin.CommonController{})
	slimgo.Run(":9022")
}
