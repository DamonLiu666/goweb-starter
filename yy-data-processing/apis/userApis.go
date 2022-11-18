package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yy-data-processing/model"
	"yy-data-processing/service"
)

type UserApis struct {
}

func (u *UserApis) SaveUser(c *gin.Context) {

	user := model.User{}
	user.Name = "user1"
	user.Age = 20
	s := service.UserService{}
	err := s.SaveUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "新增数据失败",
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "新增数据成功",
	})
}

func (u *UserApis) FindUser(c *gin.Context) {
	s := service.UserService{}
	user := s.FindUser(6)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询数据成功",
		"data": user,
	})
}
