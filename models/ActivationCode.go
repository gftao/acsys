package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
)

// Role 用户角色 实体类
type ActivationCode struct {
	ACTIVE_CODE string    `orm:"column(active_code);pk"`
	ACTIVE_FLG  string    `orm:"column(active_flg)"`
	RecUpdTs    time.Time `orm:"column(REC_UPD_TS)"`
	RecCrtTs    time.Time `orm:"column(REC_CRT_TS)"`
	//ACTFLG      string      `orm:"-"`
}

func (a *ActivationCode) TableName() string {
	return ActivationCodeTBName()
}

// RolePageList 获取分页数据
func ActivationCodeList(params *RoleQueryParam) ([]*ActivationCode, int64) {
	query := orm.NewOrm().QueryTable(ActivationCodeTBName())
	data := make([]*ActivationCode, 0)
	//默认排序
	sortorder := "active_code"
	//switch params.Sort {
	//case "Id":
	//	sortorder = "Id"
	//case "Seq":
	//	sortorder = "Seq"
	//}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("active_code__istartswith", params.ActivationCode)
	//if len(params.ActivationCode) > 0 {
	//	query = query.Filter("active_code", params.ActivationCode)
	//}
	if len(params.Activation) > 0 {
		if strings.ContainsAny(params.Activation,"未"){
			query = query.Filter("active_flg", "0")
		}else {
			query = query.Filter("active_flg", "1")
		}
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	//for i, _ := range data {
	//	switch data[i].ACTIVE_FLG {
	//	case "0":
	//		data[i].ACTFLG = "未使用"
	//	default:
	//		data[i].ACTFLG = "已使用"
	//	}
	//}
	return data, total
}
