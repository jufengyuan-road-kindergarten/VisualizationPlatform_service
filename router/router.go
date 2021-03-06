package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/controllers"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// 跨域
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowMethods("DELETE", "PUT")
	Router.Use(cors.New(corsConfig))
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	Router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	Router.Use(gin.Recovery())
	freeGroup := Router.Group("vp")
	{
		//无需登录也可使用的接口
		freeGroup.GET("/all", controllers.All)
		freeGroup.GET("/person-event", controllers.PersonEvent)
		freeGroup.GET("/event-person", controllers.EventPerson)
		freeGroup.GET("/allName", controllers.AllName)
		freeGroup.GET("/personRelation", controllers.PersonRelation)
		freeGroup.GET("/eventTree", controllers.EventTree)
		freeGroup.GET("/info", controllers.Info)
	}
}
