package Commands

import "mmagic/Model/Admin"

func init() {
	AddUser()
}
func AddUser() {
	var data = make(map[string]interface{})
	if !Admin.ExistsByLoginName("admin") {
		data["nick_name"] = "admin"
		data["login_name"] = "admin"
		data["email"] = "admin@126.com"
		data["pwd"] = Admin.Md5Pwd("admin")
		data["statues"] = int64(1)
		data["tel"] = "1"
		Admin.AddUser(data)
	}
}
