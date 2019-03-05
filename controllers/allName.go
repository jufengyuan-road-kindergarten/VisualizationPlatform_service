package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
)

func AllName(ctx *gin.Context) {
	types,_:=ctx.GetQueryArray("type[]")
	r := make([]struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}, 0)
	db:=dao.DB()
	cq:=neoism.CypherQuery{
		Statement:`
match(n) where labels(n)[0] in {types}
return n.name as name,labels(n)[0] as type
`,
Parameters:neoism.Props{"types":types},
Result:&r,
	}
	err:=db.Cypher(&cq)
	if err != nil {
		tools.ResponseError(ctx,err)
		return
	}
	tools.Response(ctx,tools.SUCCESS,r)
}