package controllers

import (
	"study/tools/common"
	"study/tools/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloworldController struct {
}

func (ctr *HelloworldController) Hello(ctx *gin.Context) {
	dao := models.HelloworldDao{}
	ctx.JSON(http.StatusOK, common.Success(*dao.Hello()))
}
