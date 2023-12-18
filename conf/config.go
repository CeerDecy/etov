package conf

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

type EtovConfig struct {
	OpenAI OpenAI `json:"openAI"`
	Mysql  Mysql  `json:"mysql"`
	Redis  Redis  `json:"redis"`
	Cache  Cache  `json:"cache"`
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
