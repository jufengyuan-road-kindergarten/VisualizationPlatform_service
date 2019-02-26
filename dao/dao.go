package dao

import (
	"fmt"
	"sync"
	"github.com/mzz2017/VisualizationPlatform_service/config"
	"github.com/jmcvetta/neoism"
)

var once sync.Once
var db *neoism.Database

func initDB() {
	conf := config.Config()
	ds := conf["dataSource"]
	ds.Get("host")
	args := fmt.Sprintf("http://%s:%s@%s:%s/db/data", ds.Get("username").String(), ds.Get("password").String(), ds.Get("host").String(), ds.Get("port").String())
	var err error
	db, err = neoism.Connect(args)
	if err != nil {
		panic(err)
	}
}

func DB() *neoism.Database {
	once.Do(initDB)
	return db
}
