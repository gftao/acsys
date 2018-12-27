package controllers

import (
	"sdrms/models"
	"encoding/json"
	"fmt"
)

type ActivationCodeController struct {
	BaseController
}

func (c *ActivationCodeController) Prepare() {
	//先执行
	c.BaseController.Prepare()
}
func (c *ActivationCodeController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "activationcode/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "activationcode/index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("RoleController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("RoleController", "Delete")
	//c.Data["canAllocate"] = c.checkActionAuthor("RoleController", "Allocate")
}
func (c *ActivationCodeController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.RoleQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	fmt.Printf("->>>>%#v", params)
	data, total := models.ActivationCodeList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
