package main

import (
	"flag"
	"imagego-go-api/database"
	"imagego-go-api/httpserver"
	"imagego-go-api/util"
)

func main() {
	defaultConfigGen := flag.Bool("g", false, "기본 설정파일 생성")
	help := flag.Bool("h", false, "도움말")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *defaultConfigGen {
		util.GenerateDefaultServerConfig()
		return
	}

	config := util.NewServerConifg()
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = database.CreateDefaultDBConnection(config.Database).Connect()
	if err != nil {
		panic(err)
	}

	err = database.GeteDefaultDBConnection().GetDB().AutoMigrate(&database.User{}, &database.Image{})
	if err != nil {
		panic(err)
	}
	httpServer := httpserver.NewHttpServer("8080")
	httpServer.HttpStart()
}
