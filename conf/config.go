package conf

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"

	"etov/conf/db"
	"etov/conf/openai"
	"etov/conf/redis"
)

type EtovConfig struct {
	OpenAI openai.OpenAI `json:"openAI"`
	Mysql  db.Mysql      `json:"mysql"`
	Redis  redis.Redis   `json:"redis"`
}

var EtovCfg EtovConfig

//go:embed config.yml
var config string

func init() {
	err := yaml.Unmarshal([]byte(config), &EtovCfg)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	//fmt.Printf("config → %+v\n", EtovCfg)
}
