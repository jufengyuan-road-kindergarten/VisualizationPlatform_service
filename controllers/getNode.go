package controllers

import (
	"github.com/jmcvetta/neoism"
	"github.com/gin-gonic/gin"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
	"errors"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
)

func getNode(ctx *gin.Context, name string, labels []string) interface{} {
	//先获得目标结点信息
	nodes := make([]struct {
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Labels     []string    `json:"labels"`
		Id         string      `json:"id"`
		Properties interface{} `json:"properties"`
	}, 0)
	cq := neoism.CypherQuery{
		Statement: `
match (n{name:{name}})
where size({labels})=0 or labels(n)[0] in {labels}
return 	n.name as name,
		labels(n) as labels,
		labels(n)[0] as type,
		toString(id(n)) as id,
		n as properties
`,
		Parameters: neoism.Props{"name": name, "labels": labels},
		Result:     &nodes,
	}
	err := dao.DB().Cypher(&cq)
	if err != nil {
		tools.ResponseError(ctx, err)
		return nil
	}
	//无查询结果
	if len(nodes) < 1 {
		tools.ResponseError(ctx, errors.New("getNode: 无查询结果，请确保Name正确"))
		return nil
	}
	//处理Person的Properties域
	data := nodes[0].Properties.(map[string]interface{})["data"]
	nodes[0].Properties = data
	return nodes[0]
}
