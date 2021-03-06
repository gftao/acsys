package controllers

import (
	"acsys/enums"
	"github.com/astaxie/beego"
	"acsys/utils"
	"encoding/json"
	"acsys/models"
	"github.com/astaxie/beego/orm"
	"encoding/base64"
	"strconv"
	"fmt"
	"math/rand"
	"time"
	"strings"
)

type VersionController struct {
	BaseController
	basePath string
	//fileName string
	urlPath string
}

//var fileName string

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
	var err error
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id := c.GetString(":Id", "")
	m := &models.PcSourceAssignInfos{}

	utils.LogDebugf("Id=[%s]", Id)
	if len(Id) > 1 {
		m, err = models.PcSourceAssignInfosOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	} else {
		m.Assign_level = 2
		m.App_source_type = "NXY"
	}
	c.Data["m"] = m
	c.setTpl("version/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "version/edit_footerjs.html"
}
func (c *VersionController) UploadImage() {
	f, h, err := c.GetFile("file")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
	defer f.Close()
	filePath := c.basePath + h.Filename
	c.SetSession(utils.ClientIP(c.Ctx.Request), h.Filename)

	utils.LogDebugf("filePath = [%s]\n", filePath)
	//log.Printf("Id = %s,fileName=%s\n", utils.ClientIP(c.Ctx.Request), (c.GetSession(utils.ClientIP(c.Ctx.Request))).(string))

	err = c.SaveToFile("file", filePath)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
	c.jsonResult(enums.JRCodeSucc, "上传成功", "/"+filePath)
}

func (c *VersionController) Save() {
	m := models.PcSourceAssignInfos{}
	mdb := models.PcSourceAssignInfos{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", 0)
	}
	utils.LogDebugf("PcSourceAssignInfos->[%+v]", m)
	m.Assign_level, _ = strconv.Atoi(m.AssignLevel)
	m.App_source_type = m.SourceType
	fileName, ok := (c.GetSession(utils.ClientIP(c.Ctx.Request))).(string)
	if !ok {
		fileName = ""
	}

	err = o.QueryTable(models.PcSourceAssignInfosTBName()).Filter("assign_key", m.AssignKey).Filter("app_source_type", m.SourceType).One(&mdb)
	utils.LogDebugf("AssignInfos=[%+v]", err)
	if err == orm.ErrNoRows {
		//utils.LogDebugf("------[%+v]", m)
		if strings.TrimSpace(m.AssignKey) == "G" {
			rand.Seed(time.Now().UnixNano())
			m.App_source_list = fmt.Sprintf("G%d", rand.Int()/10000000)
			//utils.LogDebugf("--- [%+v]", m)
		}else {
			m.App_source_list = m.AssignKey
		}
		_, err = o.Insert(&m)
		utils.LogDebugf("[%+v]", err)
		if err == nil {
			bs := c.urlPath + base64.StdEncoding.EncodeToString([]byte(fileName))
			pc := models.Pc_source_infos{
				App_id:                  m.App_source_list,
				App_source_type:         m.SourceType,
				App_source_version_code: m.AppSourceVersion,
				App_source_url:          bs,
			}
			_, err = o.InsertOrUpdate(&pc)
		}

 	} else if err == nil {
		if m.Assign_level != mdb.Assign_level || m.SourceType != mdb.App_source_type {
			_, err = o.QueryTable(models.PcSourceAssignInfosTBName()).Filter("assign_key", m.AssignKey).Filter("app_source_type", m.SourceType).Update(orm.Params{
				"assign_level":    m.Assign_level,
				"app_source_type": m.SourceType,
			})
			if err != nil {
				c.jsonResult(enums.JRCodeFailed, "编辑失败", 0)
				return
			}
		}
		if fileName != "" {
			bs := c.urlPath + base64.StdEncoding.EncodeToString([]byte( fileName))
			_, err = o.QueryTable(models.Pc_source_infosTBName()).Filter("app_id", mdb.App_source_list).Update(orm.Params{
				"app_source_version_code": m.AppSourceVersion,
				"app_source_url":          bs,
			})
		} else {
			utils.LogDebugf("-- else if --"  )
			_, err = o.QueryTable(models.Pc_source_infosTBName()).Filter("app_id", mdb.App_source_list).Update(orm.Params{
				"app_source_version_code": m.AppSourceVersion,
			})
		}
	}
	//utils.LogDebugf("c.fileName[%+v]", fileName)

	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "编辑失败", 0)
	}
	c.jsonResult(enums.JRCodeSucc, "保存成功", m.AssignKey)
}
