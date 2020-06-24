package Commands

import (
	"fmt"
	"mmagic/Database"
	"mmagic/Model/Admin"
	"mmagic/Model/Article"
	"mmagic/Model/Banner"
	"mmagic/Model/Campus"
	"mmagic/Model/Message"
	"mmagic/Model/Nav"
	"mmagic/Model/Single"
	"mmagic/Model/Site"
)

func init() {
	fmt.Println("生成数据库文件")
	InitAdminDatabase()
}
func InitAdminDatabase() {
	DropDatabase()
	Database.Db.AutoMigrate(&Admin.SysAdminUser{})
	Database.Db.AutoMigrate(&Admin.SysAdminDepartment{})
	Database.Db.AutoMigrate(&Admin.SysAdminPosition{})
	Database.Db.AutoMigrate(&Admin.SysAdminPower{})
	Database.Db.AutoMigrate(&Banner.Banner{})
	Database.Db.AutoMigrate(&Article.Article{})
	Database.Db.AutoMigrate(&Banner.Banner{})
	Database.Db.AutoMigrate(&Message.Message{})
	Database.Db.AutoMigrate(&Nav.Nav{})
	Database.Db.AutoMigrate(&Site.Site{})
	Database.Db.AutoMigrate(&Single.Single{})
	Database.Db.AutoMigrate(&Campus.Campus{})
}

func DropDatabase() {
	if !Database.Db.HasTable(&Admin.SysAdminUser{}) {
		Database.Db.DropTable(&Admin.SysAdminUser{})
	}
	if !Database.Db.HasTable(&Admin.SysAdminDepartment{}) {
		Database.Db.DropTable(&Admin.SysAdminDepartment{})
	}
	if !Database.Db.HasTable(&Admin.SysAdminPosition{}) {
		Database.Db.DropTable(&Admin.SysAdminPosition{})
	}
	if !Database.Db.HasTable(&Admin.SysAdminPower{}) {
		Database.Db.DropTable(&Admin.SysAdminPower{})
	}
	if !Database.Db.HasTable(&Banner.Banner{}) {
		Database.Db.DropTable(&Banner.Banner{})
	}
	//文章
	if !Database.Db.HasTable(&Article.Article{}) {
		Database.Db.DropTable(&Article.Article{})
	}
	//轮播图
	if !Database.Db.HasTable(&Banner.Banner{}) {
		Database.Db.DropTable(&Banner.Banner{})
	}
	//信息
	if !Database.Db.HasTable(&Message.Message{}) {
		Database.Db.DropTable(&Message.Message{})
	}
	//导航
	if !Database.Db.HasTable(&Nav.Nav{}) {
		Database.Db.DropTable(&Nav.Nav{})
	}
	//站点
	if !Database.Db.HasTable(&Site.Site{}) {
		Database.Db.DropTable(&Site.Site{})
	}
	//单页
	if !Database.Db.HasTable(&Single.Single{}) {
		Database.Db.DropTable(&Single.Single{})
	}
	//校园管理
	if !Database.Db.HasTable(&Campus.Campus{}) {
		Database.Db.DropTable(&Campus.Campus{})
	}
}
