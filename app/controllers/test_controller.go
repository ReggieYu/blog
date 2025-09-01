package controllers

import (
	"blog/app/pojo/req"
	"blog/app/service"
	"blog/config/log"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
}

func (t *TestController) HelloWorld(context *gin.Context) {
	log.Logger.Info("test hello world method")
	context.String(http.StatusOK, "hello world")
}

func (t *TestController) TestParam(context *gin.Context) {
	log.Logger.Info("test param method")
	name := context.Param("name")
	context.String(http.StatusOK, "check param %s", name)
}

func (t *TestController) TestDefaultFormParam(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "zhang san")
	gender := ctx.Query("gender")
	log.Logger.Info("test default form param")
	ctx.String(http.StatusOK, "his name is %s and gender is %s", name, gender)
}

func (t *TestController) TestPost(cxt *gin.Context) {
	name := cxt.PostForm("name")
	age := cxt.DefaultPostForm("nick", "zs")
	log.Logger.Info("test post")
	cxt.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusOK,
			"success": true,
		},
		"name": name,
		"age":  age,
	})
}

func (t *TestController) TestPostBodyParam(ctx *gin.Context) {
	var request req.TestPostRequest
	log.Logger.Info("test post body param")
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Logger.Panic("param invalid")
	}

	if _, err := json.Marshal(request); err != nil {
		log.Logger.Panic("param parse failed")
	}

	service.TestSrv.PrintInfo(&request)
}
