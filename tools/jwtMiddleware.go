package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/VisualizationPlatform_service/config"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var Secret string

func init() {
	Secret = config.Config()["secret"].String()
}

func JWTAuth(Admin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				// 我们使用固定的secret，直接返回就好
				return []byte(Secret), nil
			})
		if err != nil {
			Response(ctx, UNAUTHORIZED, err.Error())
			ctx.Abort()
			return
		}
		if !token.Valid {
			Response(ctx, UNAUTHORIZED, "Token is invalid!")
			ctx.Abort()
			return
		}
		//如果需要Admin权限
		mapClaims := token.Claims.(jwt.MapClaims)
		if Admin && mapClaims["admin"] == false {
			ResponseError(ctx, errors.New("admin required"))
			ctx.Abort()
			return
		}
		//将名字和sub丢入参数
		ctx.Set("Name", mapClaims["name"])
		ctx.Set("Sub", mapClaims["sub"])
		//在ctx.Next()前的都是before request，之后的是after request
		ctx.Next()
	}
}
