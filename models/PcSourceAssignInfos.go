package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"log"
	"strings"
	"acsys/utils"
)

type PcSourceAssignInfos struct {
	Assign_level     int       `orm:"column(assign_level)"`
	AssignLevel      string    `orm:"-"`
	AssignKey        string    `orm:"column(assign_key);pk"`
	SourceType       string    `orm:"-"`
	App_source_type  string    `orm:"column(app_source_type)"`
	App_source_list  string    `orm:"column(app_source_list)"  json:"-"`
	Rec_upd_ts       time.Time `orm:"column(rec_upd_ts);auto_now_add;type(datetime)" json:"-"`
	Rec_crt_ts       time.Time `orm:"column(rec_crt_ts);auto_now;type(datetime)"  json:"-"`
	App_assign_desc  string    `orm:"column(app_assign_desc)"`
	AppSourceVersion string    `orm:"-"`
	//FileImageUrl     string    `orm:"-"`
	//Fileinput        string    `orm:"-"`
}

func (p *PcSourceAssignInfos) TableUnique() [][]string {
	return [][]string{
		[]string{"assign_level", "assign_key", "app_source_type"},
	}
}

func (p *PcSourceAssignInfos) TableName() string {
	return PcSourceAssignInfosTBName()
}

type Pc_source_infos struct {
	App_id                  string    `orm:"column(app_id);pk"`
	App_source_type         string    `orm:"column(app_source_type)"`
	App_source_code         string    `orm:"column(app_source_code)"`
	App_source_url          string    `orm:"column(app_source_url)"`
	App_source_md5          string    `orm:"column(app_source_md5)"`
	App_source_version_code string    `orm:"column(app_source_version_code)"`
	App_source_version_name string    `orm:"column(app_source_version_name)"`
	Rec_upd_ts              time.Time `orm:"column(rec_upd_ts);auto_now_add;type(datetime)"`
	Rec_crt_ts              time.Time `orm:"column(rec_crt_ts);auto_now;type(datetime)"`
	App_ins_id              string    `orm:"column(app_ins_id)"`
	App_prod                string    `orm:"column(app_prod)"`
}

func (p *Pc_source_infos) TableName() string {
	return Pc_source_infosTBName()
}

type PcSourceAssignInfosParam struct {
	BaseQueryParam
	LevelLike    string //模糊查询
	KeyLike      string //模糊查询
	Mobile       string //精确查询
	SearchStatus string //为空不查询，有值精确查询
}

func PcSourceAssignInfosPageList(params *PcSourceAssignInfosParam) ([]*PcSourceAssignInfos, int64) {
	o := orm.NewOrm()
	query := o.QueryTable(PcSourceAssignInfosTBName())
	data := make([]*PcSourceAssignInfos, 0)
	source := &Pc_source_infos{}
	log.Printf("params=[%+v]\n", *params)

	//默认排序
	sortorder := "assign_level"
	if params.Sort == "Assign_key" {
		sortorder = "assign_key"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.LevelLike > "0" {
		query = query.Filter("assign_level__istartswith", params.LevelLike)
	}
	if params.KeyLike != "" {
		query = query.Filter("assign_key__istartswith", params.KeyLike)
	}
	//if len(params.Mobile) > 0 {
	//	query = query.Filter("mobile", params.Mobile) //精确查找
	//}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	for i, _ := range data {
		switch data[i].Assign_level {
		case 1:
			data[i].AssignLevel = "全量"
		case 2:
			data[i].AssignLevel = "商户"
		default:
			data[i].AssignLevel = "异常"
		}
		err := o.QueryTable(Pc_source_infosTBName()).Filter("app_id", data[i].App_source_list).One(source)
		if err != nil {
			continue
		}
		//log.Printf("source=[%+v]\n", *source)
		data[i].AppSourceVersion = source.App_source_version_code
	}
	return data, total
}
func PcSourceAssignInfosOne(id string) (*PcSourceAssignInfos, error) {
	o := orm.NewOrm()
	args := strings.Split(id, "-")
	l := 0
	switch args[0] {
	case "全量":
		l = 1
	default:
		l = 2
	}
	k := args[1]
	st := args[2]
	m := PcSourceAssignInfos{}
	err := o.QueryTable(m.TableName()).Filter("assign_level", l).Filter("assign_key", k).Filter("app_source_type", st).One(&m)
	if err != nil {
		return nil, err
	}

	source := &Pc_source_infos{}
	err = o.QueryTable(Pc_source_infosTBName()).Filter("app_id", m.App_source_list).One(source)
	if err != nil {
		return nil, err
	}
	utils.LogDebugf("source=[%+v]\n", m)
	m.AppSourceVersion = source.App_source_version_code
	return &m, nil
}
