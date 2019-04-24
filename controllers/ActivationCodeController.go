package controllers

import (
	"acsys/models"
	"acsys/enums"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
	"acsys/utils"
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
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "activationcode/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "activationcode/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("ActivationCodeController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("RoleController", "Delete")
}
func (c *ActivationCodeController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.ActivationCodeQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	//utils.LogDebugf("%s", string(c.Ctx.Input.RequestBody))
	utils.LogDebugf("%#v", params)
	data, total := models.ActivationCodeList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
func (c *ActivationCodeController) Edit() {
	//如果是Post请求，则由Save处理
	//fmt.Println("--Edit-->", c.Ctx.Request.Method)
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	c.setTpl("activationcode/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "activationcode/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *ActivationCodeController) Save() {

	//获取form里的值
	var params models.ActivationCodeQueryParam
	if err := c.ParseForm(&params); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", nil)
	}
	o := orm.NewOrm()
	if params.Seq == 0 {
		params.Seq = 10
	}
	for i := 0; i < params.Seq; i++ {
		m := models.ActivationCode{ACTIVE_FLG: "0"}
		m.ACTIVE_CODE = params.Name + utils.RandomString(6)
		m.RecCrtTs = time.Now()
		m.RecUpdTs = m.RecCrtTs
		b := models.Tbl_pc_belonged{}
		err := o.QueryTable(b.TableName()).Filter("name_ids", params.Name).One(&b)
		if err != nil {
			utils.LogDebugf("获取地区码失败:%s", err)
			//return nil, err
		}
		m.Active_belong = b.Real_name
		if _, err = o.Insert(&m); err != nil {
			i--
		}
	}
	o.Commit()
	c.jsonResult(enums.JRCodeSucc, "添加成功", nil)
}
