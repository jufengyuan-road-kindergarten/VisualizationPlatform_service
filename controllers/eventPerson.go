package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
	"github.com/jmcvetta/neoism"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
	"github.com/pkg/errors"
)

func EventPerson(ctx *gin.Context) {
	res := make(map[string]interface{})
	db := dao.DB()
	name, ok := ctx.GetQuery("name")
	if !ok {
		tools.ResponseError(ctx, errors.New("name参数不存在"))
		return
	}
	err := epNodes(name, db, &res)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	err = epRelations(name, db, &res)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	tools.Response(ctx, tools.SUCCESS, res)
}

func epNodes(name string, db *neoism.Database, res *map[string]interface{}) (err error) {
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
match (n)-[*1]-(e:Person)
where (n:Event or n:Meeting) and n.name contains {name}
with n, size((n)-[*1]-()) as degree
return 	n.name as name,
		labels(n) as labels,
		labels(n)[0] as type,
		degree,
		toString(id(n)) as id,
		n as properties
union
match (n)-[*1]-(e:Person)
where (n:Event or n:Meeting) and n.name contains {name}
with e,size(()-[*1]-(e)) as degree
return e.name as name,
		labels(e) as labels,
		labels(e)[0] as type,
		degree,
		toString(id(e)) as id,
		e as properties
`,
		Parameters: neoism.Props{"name": name},
		Result:     &nodes,
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

func epRelations(name string, db *neoism.Database, res *map[string]interface{}) (err error) {
	relations := make([]struct {
		Id         string      `json:"id"`
		Type       string      `json:"type"`
		Properties interface{} `json:"properties"`
		Source     string      `json:"source"`
		Target     string      `json:"target"`
	}, 0)
	cq := neoism.CypherQuery{
		Statement: `
match (s)-[r]-(e:Person)
where (s:Event or s:Meeting) and s.name contains {name}
return 	toString(1000000 + id(r)) as id,
		type(r) as type,
		r as properties,
		toString(id(s)) as source,
		toString(id(e)) as target
`,
		Parameters: neoism.Props{"name": name},
		Result:     &relations,
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
