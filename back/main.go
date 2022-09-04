package main

import (
	"github.com/cubdesign/dailyfj/config"
	"github.com/cubdesign/dailyfj/controllers"
	"github.com/cubdesign/dailyfj/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StartWebserver()
}
