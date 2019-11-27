package controllers

import (
	"strings"

	"acsys/enums"
	"acsys/models"
	"time"
	"encoding/json"
	"net/http"
	"bytes"
	"go-common/library/log"
	"io/ioutil"
	"acsys/utils"
	"github.com/astaxie/beego"
)

type WxinController struct {
	BaseController
	TranMsg *TransMessage
}

type TransMessage struct {
	Encoding    string       `json:"encoding"`
	Sign_method string       `json:"sign_method"`
	Signature   string       `json:"signature"`
	Version     string       `json:"version"`
	Msg_body    string       `json:"msg_body"`
	MsgBody     *TransParams `json:"-"`
}

type TransParams struct {
	Tran_cd    string `json:"tran_cd,omitempty"`
	Mcht_cd    string `json:"mcht_cd,omitempty"`
	Resp_cd    string `json:"resp_cd,omitempty"`
	Resp_msg   string `json:"resp_msg,omitempty"`
	Ins_id_cd  string `json:"ins_id_cd,omitempty"`
	Send_time  string `json:"send_time,omitempty"`
	Order_id   string `json:"order_id,omitempty"`
	Phone_no   string `json:"phone_no,omitempty"`
	Msg_conten string `json:"msg_conten,omitempty"`
	User_name  string `json:"user_name,omitempty"`
	Branch_id  string `json:"branch_id,omitempty"`
}

func (t *TransMessage) init() {
	t.Encoding = "UTF-8"
	t.Sign_method = "AA"
	t.Version = "1.0.0"
	t.MsgBody = new(TransParams)
	t.MsgBody.Tran_cd = "10004003"
	t.MsgBody.Order_id = time.Now().Format("200601021504059999")
	t.MsgBody.Branch_id = "9999000001"
	t.MsgBody.Ins_id_cd = "HB000054"

}
func (t *TransMessage) setBody() []byte {
	b, _ := json.Marshal(t.MsgBody)
	t.Msg_body = string(b)
	s, _ := json.Marshal(t)
	return s
}
func (c *WxinController) Index() {
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["headcssjs"] = "wxin/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "wxin/login_footerjs.html"
	//c.setTpl("wxin/index.html", "shared/layout_base.html")
	c.setTpl("wxin/weixin.html", "shared/layout_base.html")
}

func (c *WxinController) Page404() {
	c.setTpl()
}
func (c *WxinController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("wxin/error.html", "shared/layout_pullbox.html")
}

func (c *WxinController) Login() {
	//判断短信验证码是否正确：
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["headcssjs"] = "wxin/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "wxin/login_footerjs.html"
	c.setTpl("wxin/login.html", "shared/layout_base.html")
}

func (c *WxinController) GetVerifyCode() {
	mcht_cd := strings.TrimSpace(c.GetString("UserName"))
	//userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	log.Info("mcht_cd=", mcht_cd)
	//发送短信
	c.TranMsg = new(TransMessage)
	c.TranMsg.init()
	c.TranMsg.MsgBody.Mcht_cd = mcht_cd
	c.TranMsg.MsgBody.Phone_no = "15029973362" //"15890313756" "15029973362"
	sms := strings.ToUpper(string(utils.Krand(6)))
	c.TranMsg.MsgBody.Msg_conten = mcht_cd + "你本次验证码为：" + sms
	b := c.TranMsg.setBody()
	url := beego.AppConfig.String("myself::sendSMS")
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Error("DefaultClient=%+v", err)
		c.jsonResult(enums.JRCodeFailed, "DefaultClient: ", err)
	}
	defer resp.Body.Close()
	rsp, _ := ioutil.ReadAll(resp.Body)
	log.Info("发送应答成功：%v", string(rsp))

	c.jsonResult(enums.JRCodeSucc, "登录成功", "")
}

func (c *WxinController) DoLogin() {
	username := strings.TrimSpace(c.GetString("mcht_cd"))
	userpwd := strings.TrimSpace(c.GetString("code"))
	log.Info("DoLogin  mcht_cd=%s code=%s", username,  userpwd)
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	//userpwd = utils.String2md5(userpwd)
	//user, err := models.BackendUserOneByUserName(username, userpwd)
	//if user != nil && err == nil {
	//	if user.Status == enums.Disabled {
	//		c.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
	//	}
	//	//保存用户信息到session
	//	c.setBackendUser2Session(user.Id)
	//	//获取用户信息
	//	c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	//} else {
	//	c.jsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
	//}
}
func (c *WxinController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.pageLogin()
}
