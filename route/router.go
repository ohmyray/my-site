package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ohmyray/my-blog/api/v1"
)

func InitRouter() {
	gin.SetMode("debug")
	router := gin.Default()
	applicationV1(router)
	router.Run()
}

func applicationV1(r *gin.Engine) {
	r.GET("user/:id", v1.FindUserById)
	r.POST("user", v1.InsertUser)

}
