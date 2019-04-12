package routers

import (
	"acsys/controllers"

	"github.com/astaxie/beego"
	"path"
)

func init() {
	//用户角色路由
	bp := func(p string) string {
		return path.Join("acsys", p)
	}
	beego.Router(bp("/role/index"), &controllers.RoleController{}, "*:Index")
	beego.Router(bp("/role/datagrid"), &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router(bp("/role/edit/?:id"), &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router(bp("/role/delete"), &controllers.RoleController{}, "Post:Delete")
	beego.Router(bp("/role/datalist"), &controllers.RoleController{}, "Post:DataList")
	beego.Router(bp("/role/allocate"), &controllers.RoleController{}, "Post:Allocate")
	beego.Router(bp("/role/updateseq"), &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router(bp("/resource/index"), &controllers.ResourceController{}, "*:Index")
	beego.Router(bp("/resource/treegrid"), &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router(bp("/resource/edit/?:id"), &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router(bp("/resource/parent"), &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router(bp("/resource/delete"), &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router(bp("/resource/updateseq"), &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router(bp("/resource/select"), &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router(bp("/resource/usermenutree"), &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router(bp("/resource/checkurlfor"), &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router(bp("/backenduser/index"), &controllers.BackendUserController{}, "*:Index")
	beego.Router(bp("/backenduser/datagrid"), &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router(bp("/backenduser/edit/?:id"), &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router(bp("/backenduser/delete"), &controllers.BackendUserController{}, "Post:Delete")
	//后台用户中心
	beego.Router(bp("/usercenter/profile"), &controllers.UserCenterController{}, "Get:Profile")
	beego.Router(bp("/usercenter/basicinfosave"), &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router(bp("/usercenter/uploadimage"), &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router(bp("/usercenter/passwordsave"), &controllers.UserCenterController{}, "Post:PasswordSave")

	beego.Router(bp("/home/index"), &controllers.HomeController{}, "*:Index")
	beego.Router(bp("/home/login"), &controllers.HomeController{}, "*:Login")
	beego.Router(bp("/home/dologin"), &controllers.HomeController{}, "Post:DoLogin")
	beego.Router(bp("/home/logout"), &controllers.HomeController{}, "*:Logout")

	beego.Router(bp("/home/404"), &controllers.HomeController{}, "*:Page404")
	beego.Router(bp("/home/error/?:error"), &controllers.HomeController{}, "*:Error")
	//激活码路由
	beego.Router(bp("/activationcode/index"), &controllers.ActivationCodeController{}, "*:Index")
	beego.Router(bp("/activationcode/datagrid"), &controllers.ActivationCodeController{}, "POST:DataGrid")
	beego.Router(bp("/activationcode/edit/?:id"), &controllers.ActivationCodeController{}, "Get,Post:Edit")
	//商户路由
	beego.Router(bp("/pcmchtinfos/index"), &controllers.PcMchtInfosController{}, "*:Index")
	beego.Router(bp("/pcmchtinfos/datagrid"), &controllers.PcMchtInfosController{}, "POST:DataGrid")
	beego.Router(bp("/pcmchtinfos/edit/?:MchtCd"), &controllers.PcMchtInfosController{}, "Get,Post:Edit")
	//beego.Router("/pcmchtinfos/delete", &controllers.PcMchtInfosController{}, "Post:Delete")
	beego.Router(bp("/"), &controllers.HomeController{}, "*:Index")
	//版本控制
	beego.Router(bp("/version/index"), &controllers.VersionController{}, "*:Index")
	beego.Router(bp("/version/datagrid"), &controllers.VersionController{}, "POST:DataGrid")
	beego.Router(bp("/version/edit/?:Id"), &controllers.VersionController{}, "Get,Post:Edit")
	beego.Router(bp("/version/uploadimage"), &controllers.VersionController{}, "Post:UploadImage")

}
