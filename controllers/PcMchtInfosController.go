package controllers

import (
	"sdrms/models"
	"github.com/astaxie/beego/orm"
	"sdrms/enums"
	"strconv"
	"strings"
	"fmt"
	"encoding/json"
)

type PcMchtInfosController struct {
	BaseController
}

func (c *PcMchtInfosController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()

}
func (c *PcMchtInfosController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "pcmchtinfos/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "pcmchtinfos/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("PcMchtInfosController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("PcMchtInfosController", "Delete")
}
func (c *PcMchtInfosController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.PcMchtInfosQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	fmt.Println(params)
	data, total := models.PcMchtInfosPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
func (c *PcMchtInfosController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	MchtCd := c.GetString(":MchtCd", "")
	fmt.Println("----->", c.Ctx.Request.Method, MchtCd)
	m := &models.PcMchtInfos{}
	var err error
	if MchtCd != "" {
		m, err = models.PcMchtInfosOne(MchtCd)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("pcmchtinfos/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "pcmchtinfos/edit_footerjs.html"
}
func (c *PcMchtInfosController) Save() {
	m := models.PcMchtInfos{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.MchtCd)
	}

	if _, err := o.Update(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "编辑失败", m.MchtCd)
	} else {
		c.jsonResult(enums.JRCodeSucc, "保存成功", m.MchtCd)
	}
}
func (c *PcMchtInfosController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(models.BackendUserTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
