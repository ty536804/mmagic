package Admin

import (
	"github.com/gin-gonic/gin"
	"mmagic/Pkg/e"
)

// @summer权限列表
func PowerShow(c *gin.Context) {
	c.HTML(e.SUCCESS, "admin/power.html", gin.H{
		"title": "权限管理",
		"small": "权限列表",
	})
}

// @summer 添加权限
func PowerAdd(c *gin.Context) {
	c.Request.Body = e.GetBody(c)
}
