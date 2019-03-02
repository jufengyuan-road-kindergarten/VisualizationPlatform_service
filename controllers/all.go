package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
	"github.com/jmcvetta/neoism"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
)

func All(ctx *gin.Context) {
	res := make(map[string]interface{})
	db := dao.DB()

	err := allNodes(db, &res)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	err = allRelation(db, &res)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	tools.Response(ctx, tools.SUCCESS, res)
}

func allNodes(db *neoism.Database, res *map[string]interface{}) (err error) {
	nodes := make([]struct {
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Labels     []string    `json:"labels"`
		Degree     int         `json:"degree"`
		Id         string      `json:"id"`
		Properties interface{} `json:"properties"`
	}, 0)
	cq := neoism.CypherQuery{
		Statement: `
match (n)
with n,size((n)-[*1]-()) as degree
return 	n.name as name,
		labels(n) as labels,
		labels(n)[0] as type,
		degree,
		toString(id(n)) as id,
		n as properties
`,
		Result: &nodes,
	}
	err = db.Cypher(&cq)
	if err != nil {
		return
	}
	for i := range nodes {
		data := nodes[i].Properties.(map[string]interface{})["data"]
		nodes[i].Properties = data
	}
	(*res)["nodes"] = nodes
	return
}

func allRelation(db *neoism.Database, res *map[string]interface{}) (err error) {
	relations := make([]struct {
		Id         string      `json:"id"`
		Type       string      `json:"type"`
		Properties interface{} `json:"properties"`
		Source     string      `json:"source"`
		Target     string      `json:"target"`
	}, 0)
	cq := neoism.CypherQuery{
		Statement: `
match (s)-[r]->(t) 
return 	toString(1000000 + id(r)) as id,
		type(r) as type,
		r as properties,
		toString(id(s)) as source,
		toString(id(t)) as target
`,
		Result: &relations,
	}
	err = db.Cypher(&cq)
	if err != nil {
		return
	}
	for i := range relations {
		data := relations[i].Properties.(map[string]interface{})["data"]
		relations[i].Properties = data
	}
	(*res)["edges"] = relations
	return
}
