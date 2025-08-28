package main

import (
	"CurrencyExchangeApp/config"
	"CurrencyExchangeApp/router"
)

func main() {

	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080" //默认配置
	}
	r.Run(port)
}
