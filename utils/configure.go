package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var AppConf map[string]interface{}
var DatabaseCfg map[string]interface{}

// LoadConfig will try to search around for the corresponding config file.
// It will search /tmp/fileName then attempt ./config/fileName,
// then ../config/fileName and last it will look at fileName
func LoadConfig(fileName string) {

	fileName = findConfigFile(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic("utils.config.load_config.opening.panic")
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&AppConf)
	if err != nil {
		panic("utils.config.load_config.decoding.panic")
	}

}

func findConfigFile(fileName string) string {
	if _, err := os.Stat("/tmp/" + fileName); err == nil {
		fileName, _ = filepath.Abs("/tmp/" + fileName)
	} else if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	}

	return fileName
}
