package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//ConfigStruct 配置信息结构体
type ConfigStruct struct {
	Port          int
	StartIndex    uint64
	EndIndex      uint64
	NeighborAddrs []string
}

//Config 配置信息
var Config ConfigStruct

func init() {
	data, err := ioutil.ReadFile("conf/app.yml")
	if err != nil {
		Debug("init ReadFile error:", err.Error())
		return
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		Debug("init Unmarshal error: ", err.Error())
		return
	}
}
