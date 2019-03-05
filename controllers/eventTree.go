package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmcvetta/neoism"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/dao"
	"github.com/jufengyuan-road-kindergarten/VisualizationPlatform_service/tools"
)

type node struct {
	Name     string  `json:"name"`
	Edge     *string `json:"edge"`
	Children []node  `json:"children"`
}
type result struct {
	Source   string `json:"source"`
	Edge     string `json:"edge"`
	Target   string `json:"target"`
	hasBuilt bool
}

func EventTree(ctx *gin.Context) {
	name, _ := ctx.GetQuery("eventName")
	node := getNode(ctx, name, []string{"Event", "Meeting"})
	r := make([]result, 0)
	cq := neoism.CypherQuery{
		Statement: `
match p=(s{name:{name}})-[*]->(t)
where (s:Event or s:Meeting) and (t:Event or t:Meeting)
with relationships(p) as rs
unwind rs as r
unwind [(a)-[r]->(b)|{source:a.name,type:r.type,target:b.name}] as line
with distinct line as l
order by l.source
return l.source as source,l.type as edge,l.target as target
`,
		Parameters: neoism.Props{"name": name},
		Result:     &r,
	}
	err := dao.DB().Cypher(&cq)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	if len(r) < 1 {
		tools.Response(ctx, tools.SUCCESS, gin.H{"tree": gin.H{"name": name, "edge": nil, "children": []interface{}{}}, "node": node})
		return
	}
	tools.Response(ctx, tools.SUCCESS, gin.H{"tree": buildTree(name, nil, r), "node": node})
}

func buildTree(name string, edge *string, r []result) node {
	n := node{
		Name:     name,
		Edge:     edge,
		Children: []node{},
	}
	for j := 0; j < len(r); j++ {
		if r[j].hasBuilt == false && r[j].Source == name {
			r[j].hasBuilt = true
			n.Children = append(n.Children, buildTree(r[j].Target, &r[j].Edge, r))
		}
	}
	return n
}
