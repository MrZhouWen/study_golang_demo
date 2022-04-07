package favorite

import (
	"gin-answer/util"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	list := []int{1, 2, 3, 4, 5}
	util.ResponseSuccess(c, list)
}
