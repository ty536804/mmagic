package Banner

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"mmagic/Model/Banner"
	"mmagic/Pkg/e"
	"mmagic/Services"
)

// @Summer列表
func List(c *gin.Context) {
	c.HTML(e.SUCCESS, "banner/list.html", gin.H{
		"title": "banner列表",
	})
}

// @Summer 获取所有图片
func GetBanners(c *gin.Context) {
	page := com.StrTo(c.Query("page")).MustInt()
	var data = make(map[string]interface{})
	data["count"] = e.GetPageNum(Banner.GetBannerTotal())
	data["list"] = Banner.GetBanners(page)

	e.Success(c, "获取banner列表", data)
}

// @Summer 详情
func Detail(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	c.HTML(e.SUCCESS, "banner/detail.html", gin.H{
		"title": "banner详情",
		"id":    id,
	})
}

func GetBanner(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	var data = make(map[string]interface{})
	data["list"] = Services.GetNavs(data)
	data["detail"] = Banner.GetBanner(id)
	e.Success(c, "获取banner详情", data)
}

// @Summer banner保存
func AddBanner(c *gin.Context) {
	code, msg := Services.AddBanner(c)
	e.SendRes(c, code, msg, "")
}

// @Summer 删除banner
func DelBanner(c *gin.Context) {
	code, msg := Services.DelBanner(c)
	e.SendRes(c, code, msg, "")
}
