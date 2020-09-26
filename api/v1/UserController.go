package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ohmyray/my-blog/model"
	"github.com/ohmyray/my-blog/utils"
	"net/http"
)

func InsertUser(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)
	if code%2 == 0 {
		code = model.AddUser(&user)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": utils.GetMessage(code),
		"code":    code,
	})
}

func FindUserById(c *gin.Context) {
	id := c.Param("id")
	data, code := model.UserInfo(id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": utils.GetMessage(code),
		"code":    code,
		"data":    data,
	})
}
