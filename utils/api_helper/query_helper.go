package api_helper

import (
	"github.com/gin-gonic/gin"
	"oldman.study/gin-project-template/utils/pagination"
)

var userIdText = "userId"

// 从context获得用户id
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
