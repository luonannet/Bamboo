package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//ConfigStruct 配置信息结构体
type ConfigStruct struct {
	AppName       string
	Port          int
	NeighborAddrs []string
}

//Config 配置信息
var Config ConfigStruct

func init() {
	data, err := ioutil.ReadFile("conf/app.yml")
	if err != nil {
		Error(err.Error())
		return
	}
	yaml.Unmarshal(data, &Config)
}
