package routers

import (
	"study/tools/controllers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"regexp"
	"time"
)

var router *gin.Engine

//配置所有路由
func init() {
	router = gin.New()
	router.Use(logger())
	router.HEAD("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
		return
	})
	router.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "%s", "OK")
		return
	})
	helloworldController := controllers.HelloworldController{}
	router.GET("/hello", helloworldController.Hello)
	router.NoRoute(func(c *gin.Context) {
		pwd, _ := os.Getwd()
		staticPath := pwd + "/static/index.html"
		c.File(staticPath)
	})

}

var staticReg = regexp.MustCompile(".(js|css|woff2|html|woff|ttf|svg|png|eot|map)$")

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodHead {
			c.Next()
			return
		}
		t := time.Now()
		c.Next()
		// after request
		latency := time.Since(t)
		// access the status we are sending
		status := c.Writer.Status()
		resource := c.Request.URL.Path
		if !staticReg.MatchString(resource) {
			logrus.Info(latency, status, resource)
		}

	}
}

func Router() *gin.Engine {
	return router
}
