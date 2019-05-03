package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
)

func Info(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	types, _ := ctx.GetQueryArray("types")
	node := getNode(ctx, name, types)
	tools.Response(ctx, tools.SUCCESS, gin.H{"node": node})
}
