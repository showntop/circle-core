package migrations

import (
	"github.com/showntop/circle-core/stores"
	"github.com/showntop/circle-core/utils"
)

func Migrate() {
	////从config里获取数据库名
	utils.LoadConfig("config/config.json")

	store := stores.NewStore(utils.AppConf["SqlSettings"].(map[string]interface{}))

	////创建表，一旦schema之后只允许添加不允许修改
	createUsers(store.GetMaster())

	store.GetMaster().CreateTablesIfNotExists()
	// store.Close()
}
