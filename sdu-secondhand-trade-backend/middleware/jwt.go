package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/util"
	"strings"
)

func JWT(minRoleID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		aw := app.NewWrapper(c)
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = authHeader[7:]
		}
		claims, err := util.ParseJWT(authHeader)
		if err != nil {
			aw.Ctx.JSON(401, gin.H{
				"code": -1,
				"msg":  "该接口需要登录",
				"data": nil,
			})
			aw.Ctx.Abort()
			return
		}
		if claims.RoleID < minRoleID {
			aw.Error(fmt.Sprintf("该接口需要角色大于或等于：%v（请尝试退出后重登录）", minRoleID))
			aw.Ctx.Abort()
			return
		}
		aw.Ctx.Set("userClaims", claims)
	}
}
