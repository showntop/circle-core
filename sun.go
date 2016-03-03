package main

import (
	"github.com/showntop/circle-core/logger"
	"github.com/showntop/circle-core/server"
	"github.com/showntop/circle-core/stores"
	"github.com/showntop/circle-core/utils"
)

////
////development、test、staging、production

func main() {

	//加载配置信息
	utils.LoadConfig("config/config.json")
	stores.NewStore(utils.AppConf["SqlSettings"].(map[string]interface{}))
	// utils.Server
	server.Fire(utils.AppConf["ServerSettings"].(map[string]interface{}))
	logger.Info("fjflksajdfklsfj")
}
