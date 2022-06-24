package initialize

import (
	"fmt"
	"gin-practice/src/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// LoadConfig load config
func LoadConfig() error {

	buf, err := ioutil.ReadFile(global.ConfigPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, global.GlobalConfig)
	if err != nil {
		return err
	}
	fmt.Println(global.GlobalConfig)
	return nil
}
