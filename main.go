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
	dao.DB()
	log.Fatal(router.Router.Run(fmt.Sprintf("%s:%s", address, port)))
}
