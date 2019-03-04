package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
	"strconv"
	"github.com/pkg/errors"
)

func PersonRelation(ctx *gin.Context) {
	r:=make([]struct{
		Name string `json:"name"`
		Weight int `json:"weight"`
		Events []interface{} `json:"events"`
	},0)
	name,_:=ctx.GetQuery("person")
	p,_:=ctx.GetQuery("page")
	ps,_:=ctx.GetQuery("pageSize")
	page,_:=strconv.ParseInt(p,10,strconv.IntSize)
	pageSize,_:=strconv.ParseInt(ps,10,strconv.IntSize)
	db:=dao.DB()

	//先获得目标Person信息
	nodes := make([]struct {
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Labels     []string    `json:"labels"`
		Id         string      `json:"id"`
		Properties interface{} `json:"properties"`
	}, 0)
	cq:=neoism.CypherQuery{
		Statement:`
match (n:Person{name:{name}})
return 	n.name as name,
		labels(n) as labels,
		labels(n)[0] as type,
		toString(id(n)) as id,
		n as properties
`,
Parameters:neoism.Props{"name":name},
Result:&nodes,
	}
	err:=db.Cypher(&cq)
	if err != nil {
		tools.ResponseError(ctx,err)
		return
	}
	//无查询结果
	if len(nodes)<1{
		tools.ResponseError(ctx,errors.New("无查询结果，请确保Name正确"))
		return
	}
	//处理Person的Properties域
	data := nodes[0].Properties.(map[string]interface{})["data"]
	nodes[0].Properties = data

	//现在计算关系
	cq=neoism.CypherQuery{
		Statement:`
match (n:Person{name:{name}})--()--(p:Person)
with distinct p,n
with size((p)--()--(n)) as weight,[(p)--(s)--(n)|{name:s.name, type:labels(s)[0]}] as t,p
order by weight desc
skip {page}*{pageSize}
limit {pageSize}
return p.name as name,weight,t as events
`,
Parameters:neoism.Props{"name":name,"page":page,"pageSize":pageSize},
Result:&r,
	}
	err=db.Cypher(&cq)
	if err != nil {
		tools.ResponseError(ctx,err)
		return
	}

	//现在计算总条数
	cnt := make([]struct {
		Cnt int `json:"cnt"`
	}, 0)
	cq=neoism.CypherQuery{
		Statement:`match (n:Person{name:{name}})--()--(p:Person) return count(p) as cnt`,
		Parameters:neoism.Props{"name":name},
		Result:&cnt,
	}
	err=db.Cypher(&cq)
	if err != nil {
		tools.ResponseError(ctx,err)
		return
	}
	tools.Response(ctx,tools.SUCCESS,gin.H{"relations":r,"node":nodes[0],"count":cnt[0].Cnt})

}