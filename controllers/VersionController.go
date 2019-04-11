package controllers

import (
	"acsys/enums"
	"github.com/astaxie/beego"
	"acsys/utils"
	"encoding/json"
	"acsys/models"
	"github.com/go-xweb/log"
	"github.com/astaxie/beego/orm"
	"encoding/base64"
)

type VersionController struct {
	BaseController
	basePath string
	//fileName string
	urlPath string
}

var fileName string

func (c *VersionController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
	c.basePath = beego.AppConfig.String("myself::basePath")
	c.urlPath = beego.AppConfig.String("myself::urlPath")
	utils.LogDebugf("basePath = [%s]\n", c.basePath)
}

//Index 角色管理首页
func (c *VersionController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "version/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "version/index_footerjs.html"

	c.Data["canEdit"] = c.checkActionAuthor("VersionController", "Edit")
}

func (c *VersionController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.PcSourceAssignInfosParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.PcSourceAssignInfosPageList(&params)
	//定义返回的数据结构

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *VersionController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id := c.GetString(":Id", "")
	log.Infof("Id=[%s]", Id)
	m := &models.PcSourceAssignInfos{}
	var err error
	if Id != "" {
		m, err = models.PcSourceAssignInfosOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("version/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "version/edit_footerjs.html"
}
func (c *VersionController) UploadImage() {
	//这里type没有用，只是为了演示传值
	//stype, _ := c.GetInt32("type", 0)
	f, h, err := c.GetFile("file")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
	defer f.Close()
	filePath := c.basePath + h.Filename
	fileName = h.Filename
	utils.LogDebugf("filePath = [%s]\n", filePath)
	log.Printf("fileName=%s\n", fileName)
	// 保存位置在 static/upload, 没有文件夹要先创建
	err = c.SaveToFile("file", filePath)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
	c.jsonResult(enums.JRCodeSucc, "上传成功", "/"+filePath)
}

func (c *VersionController) Save() {
	m := models.PcSourceAssignInfos{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", 0)
	}
	log.Printf("==>%+v", m)
	v := m.AppSourceVersion
	err = o.QueryTable(models.PcSourceAssignInfosTBName()).Filter("assign_level", m.AssignLevel).Filter("assign_key", m.AssignKey).One(&m)
	log.Printf("c.fileName==>%+v", fileName)
	if fileName != "" {
		bs := c.urlPath + base64.StdEncoding.EncodeToString([]byte( fileName))
		_, err = o.QueryTable(models.Pc_source_infosTBName()).Filter("app_id", m.App_source_list).Update(orm.Params{
			"app_source_version_code": v,
			"app_source_url":          bs,
		})
	} else {
		_, err = o.QueryTable(models.Pc_source_infosTBName()).Filter("app_id", m.App_source_list).Update(orm.Params{
			"app_source_version_code": v,
		})
	}

	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "编辑失败", 0)
	}
	c.jsonResult(enums.JRCodeSucc, "保存成功", m.AssignKey)
}
