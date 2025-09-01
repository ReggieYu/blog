package router

import (
	"blog/app/controllers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.Engine) {
	r.GET("/", controllers.TestCtr.HelloWorld)
	r.GET("/test/:name", controllers.TestCtr.TestParam)
	r.GET("/test1", controllers.TestCtr.TestDefaultFormParam)
	r.POST("/testPost", controllers.TestCtr.TestPost)
	r.POST("/testBodyPost", controllers.TestCtr.TestPostBodyParam)
}
