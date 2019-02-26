package config

import (
	"sync"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

type Conf map[string]gjson.Result

var conf Conf

func initConfig() {
	b, err := ioutil.ReadFile("config/app.json")
	if err != nil {
		panic("配置文件读取失败。请将app_reference.json复制为app.json后，修改app.json中的配置")
	}
	conf = gjson.ParseBytes(b).Map()
}

var once sync.Once

func Config() Conf {
	once.Do(initConfig)
	return conf
}
