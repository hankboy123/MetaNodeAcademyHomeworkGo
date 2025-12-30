package middleware

import (
	"net/http"
	"sh-manage/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 这里实现JWT认证逻辑
// 示例：检查Authorization头部，验证JWT令牌等
// 如果认证失败，返回401错误
// 如果认证成功，调用c.Next()继续处理请求
func Auth(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			//c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is missing"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header format must be Bearer {token}")
			c.Abort()
			//c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		tokenString := parts[1]

		claims, error := utils.ParseToken(tokenString, jwtSecret)
		if error != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid token: "+error.Error())
			c.Abort()
			//c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token: " + error.Error()})
			return
		}

		c.Set("userID", claims.UserId)
		c.Set("username", claims.Username)

		c.Next()
	}
}
