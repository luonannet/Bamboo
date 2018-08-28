package utils

import (
	"fmt"
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
		fmt.Println("init ReadFile error:", err.Error())
		return
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		fmt.Println("init Unmarshal error: ", err.Error())
		return
	}
}
