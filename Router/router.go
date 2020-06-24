package Router

import (
	"github.com/gin-gonic/gin"
	"io"
	b "mmagic/Backend/Controller/Message"
	m "mmagic/Frontend/Controller"
	"mmagic/Pkg/e"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdin)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	dir := e.GetDir()
	//加载后端js、样式文件
	r.StaticFS("static", http.Dir(dir+"/Resources/Public"))
	//加载后端文件
	r.LoadHTMLGlob("Resources/View/**/*")

	//移动端
	r.GET("/", m.Index)
	r.GET("/sub", m.Subject)
	r.GET("/le", m.Learn)
	r.GET("/om", m.Omo)
	r.GET("/authorize", m.Authorize)
	r.POST("/AddMessage", b.AddMessage)
	return r
}
