package main

import (
	"log"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/router"
	"fmt"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/config"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
)

func main() {
	conf := config.Config()
	address := conf["address"].String()
	port := conf["port"].String()
	log.Println("正在连接数据库...")
	dao.DB()
	log.Println("连接数据库成功")
	log.Fatal(router.Router.Run(fmt.Sprintf("%s:%s", address, port)))
}
