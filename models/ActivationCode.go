package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
)

// Role 用户角色 实体类
type ActivationCode struct {
	ACTIVE_CODE string    `orm:"column(active_code);pk"`
	ACTIVE_FLG  string    `orm:"column(active_flg)"`
	RecUpdTs    time.Time `orm:"column(REC_UPD_TS)"`
	RecCrtTs    time.Time `orm:"column(REC_CRT_TS)"`
	RecUpd      string    `orm:"-"`
	RecCrt      string    `orm:"-"`
}

func (a *ActivationCode) TableName() string {
	return ActivationCodeTBName()
}

// RoleQueryParam 用于搜索的类
type ActivationCodeQueryParam struct {
	BaseQueryParam
	NameLike       string
	ActivationCode string //为空不查询，有值精确查询
	Activation     string //为空不查询，有值精确查询
	Name           string
	Seq            int
}

// RolePageList 获取分页数据
func ActivationCodeList(params *ActivationCodeQueryParam) ([]*ActivationCode, int64) {
	query := orm.NewOrm().QueryTable(ActivationCodeTBName())
	data := make([]*ActivationCode, 0)
	//默认排序
	sortorder := params.Sort
	if params.Sort == "RecCrt" {
		sortorder = "REC_CRT_TS"
	}
	fmt.Println("->", params.Sort)
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("active_code__istartswith", params.ActivationCode)

	if len(params.Activation) > 0 {
		if strings.ContainsAny(params.Activation, "是") {
			query = query.Filter("active_flg", "1")
		} else {
			query = query.Filter("active_flg", "0")
		}
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	for i, _ := range data {
		data[i].RecUpd = data[i].RecUpdTs.Format("2006-01-02 15:04:05")
		data[i].RecCrt = data[i].RecCrtTs.Format("2006-01-02 15:04:05")
	}
	return data, total
}
