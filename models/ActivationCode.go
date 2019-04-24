package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"acsys/utils"
)

// Role 用户角色 实体类
type ActivationCode struct {
	ACTIVE_CODE   string    `orm:"column(active_code);pk"`
	ACTIVE_FLG    string    `orm:"column(active_flg)"`
	Active_belong string    `orm:"column(active_belong)"`
	RecUpdTs      time.Time `orm:"column(REC_UPD_TS)"`
	RecCrtTs      time.Time `orm:"column(REC_CRT_TS)"`
	RecUpd        string    `orm:"-"`
	RecCrt        string    `orm:"-"`
}

func (a *ActivationCode) TableName() string {
	return ActivationCodeTBName()
}

type Tbl_pc_belonged struct {
	Belong_id  int    `orm:"column(belong_id);pk"`
	Name_ids   string `orm:"column(name_ids)"`
	App_name   string `orm:"column(app_name)"`
	Real_name  string `orm:"column(real_name)"`
	Rec_crt_ts time.Time
	Rec_upd_ts time.Time
}

func (t Tbl_pc_belonged) TableName() string {
	return PcBelongedTBName()
}

// RoleQueryParam 用于搜索的类
type ActivationCodeQueryParam struct {
	BaseQueryParam
	NameLike       string
	ActivationCode string //为空不查询，有值精确查询
	Activation     string //为空不查询，有值精确查询
	Name           string
	BelongName     string
	Seq            int
}

// RolePageList 获取分页数据
func ActivationCodeList(params *ActivationCodeQueryParam) ([]*ActivationCode, int64) {
	o := orm.NewOrm()
	query := o.QueryTable(ActivationCodeTBName())
	data := make([]*ActivationCode, 0)
	//默认排序
	sortorder := params.Sort
	if params.Sort == "RecCrt" {
		sortorder = "REC_CRT_TS"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("active_code__istartswith", params.ActivationCode)
	if params.BelongName != "" {
		query = query.Filter("active_belong__istartswith", params.BelongName)
	}

	if len(params.Activation) > 0 {
		if strings.ContainsAny(params.Activation, "是") {
			query = query.Filter("active_flg", "1")
		} else {
			query = query.Filter("active_flg", "0")
		}
	}
	total, _ := query.Count()
	_, err := query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	if err != nil {
		o.Rollback()
		utils.LogDebugf("查询失败:%s", err)
		return nil, 0
	}

	for i, _ := range data {
		data[i].RecUpd = data[i].RecUpdTs.Format("2006-01-02 15:04:05")
		data[i].RecCrt = data[i].RecCrtTs.Format("2006-01-02 15:04:05")

		if data[i].Active_belong == "" {
			b := Tbl_pc_belonged{}
			err := o.QueryTable(b.TableName()).Filter("name_ids", data[i].ACTIVE_CODE[:2]).One(&b)
			if err != nil {
				utils.LogDebugf("获取地区码失败:%s", err)
			}
			//if num, err := o.QueryTable(ActivationCodeTBName()).Filter("active_code", data[i].ACTIVE_CODE).Update(orm.Params{"active_belong": b.Real_name}); err == nil {
			//	utils.LogDebugf("更新失败:%dm,%s", num, err)
			//	break
			//}
			data[i].Active_belong = b.Real_name
		}
	}
	o.Commit()
	return data, total
}
