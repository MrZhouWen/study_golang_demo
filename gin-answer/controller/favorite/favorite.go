package favorite

import (
	"gin-answer/util"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	name, _ := c.GetQuery("name")
	//list := []int{1, 2, 3, 4, 5}
	//util.ResponseSuccess(c, list)
	util.ResponseSuccess(c, name)
}

func PostSubmit(c *gin.Context) {
	title := c.PostForm("title") //获取post参数
	util.ResponseSuccess(c, title)
}
