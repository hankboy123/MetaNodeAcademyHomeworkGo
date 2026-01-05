package utils

import (
	"sh-ethereum/consts"

	"github.com/gin-gonic/gin"
)

func GetCurrentUserID(c *gin.Context) uint {
	return c.GetUint(consts.UserID)
}
