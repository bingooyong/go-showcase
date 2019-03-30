package main

import (
	"study/tools/config"
	"study/tools/models"
	"study/tools/routers"
	"study/tools/log"
	"fmt"
	"flag"
	"runtime/debug"
	"os"
	"strconv"
)

var c = flag.String("c", "./etc/config.yaml", "config file path")
var version = flag.Bool("v", false, "show version")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	flag.Parse()

	if *version {
		fmt.Println(config.VERSION)
		os.Exit(0)
		return
	}
	cfg := config.Instance()
	cfg.Load(*c)

	//config log
	log.Config(cfg)
	//config mysql
	models.Config(cfg)

	router := routers.Router()
	router.Run(":" + strconv.Itoa(cfg.Server.Port))

}
