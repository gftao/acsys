package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type PcMchtInfosQueryParam struct {
	BaseQueryParam
	MchtCdLike     string //模糊查询
	ActiveCodeLike string //模糊查询
	TermID         string //模糊查询
	SearchStatus   string //为空不查询，有值精确查询
}

// Role 用户角色 实体类
type PcMchtInfos struct {
	MchtCd          string    `orm:"column(MCHT_CD);pk"`
	TermId          string    `orm:"column(TERM_ID)"`
	MchtNm          string    `orm:"column(MCHT_NM)"`
	ActiveFlg       string    `orm:"column(ACTIVE_FLG)"`
	SignFlg         string    `orm:"column(SIGN_FLG)"`
	Term_batch      string    `orm:"column(TERM_BATCH)"`
	Term_seq        string    `orm:"column(TERM_SEQ)"`
	TermProd        string    `orm:"column(TERM_PROD)"`
	TermModel       string    `orm:"column(TERM_MODEL)"`
	BrandKsn        string    `orm:"column(BRAND_KSN)"`
	ActiveCode      string    `orm:"column(active_code)"`
	KeyDownFlg      string    `orm:"column(KEY_DOWN_FLG)"`
	SessionKey      string    `orm:"column(SESSION_KEY)"`
	TermPubKey      string    `orm:"column(TERM_PUB_KEY)"`
	ServerPubKey    string    `orm:"column(SERVER_PUB_KEY)"`
	ServerPriKey    string    `orm:"column(SERVER_PRI_KEY)"`
	PbocIcKey       string    `orm:"column(PBOC_IC_KEY)"`
	PbocIcParam     string    `orm:"column(PBOC_IC_PARAM)"`
	RecUpdTs        time.Time `orm:"column(REC_UPD_TS);type(date)"`
	RecCrtTs        time.Time `orm:"column(REC_CRT_TS);type(date)"`
	Pc_expired_date int       `orm:"column(pc_expired_date)"`
}

func (a *PcMchtInfos) TableName() string {
	return PcMchtInfosTBName()
}

func (a *PcMchtInfos) TableUnique() [][]string {
	return [][]string{
		[]string{"MCHT_CD", "active_code"},
	}
}
func PcMchtInfosOne(MchtCd string) (*PcMchtInfos, error) {
	o := orm.NewOrm()
	m := PcMchtInfos{MchtCd: MchtCd}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// BackendUserPageList 获取分页数据
func PcMchtInfosPageList(params *PcMchtInfosQueryParam) ([]*PcMchtInfos, int64) {
	query := orm.NewOrm().QueryTable(PcMchtInfosTBName())
	data := make([]*PcMchtInfos, 0)
	//默认排序
	sortorder := "MCHT_CD"
	switch params.Sort {
	case "MCHT_CD":
		sortorder = "MCHT_CD"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("MCHT_CD__istartswith", params.MchtCdLike)
	query = query.Filter("active_code__istartswith", params.ActiveCodeLike)
	//if len(params.MchtCdLike) > 0 {
	//	query = query.Filter("mobile", params.ActiveCodeLike)
	//}
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder, "TERM_ID").Limit(params.Limit, params.Offset).All(&data)
	for i, _ := range data {
		switch data[i].ActiveFlg {
		case "1":
			data[i].ActiveFlg = "激活"
		case "2":
			data[i].ActiveFlg = "过期"
		case "3":
			data[i].ActiveFlg = "注销"
		default:
			data[i].ActiveFlg = "异常"
		}
	}
	return data, total
}
