package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(ActivationCode),new(Tbl_pc_belonged), new(Pc_source_infos), new(PcSourceAssignInfos), new(PcMchtInfos), new(BackendUser), new(Resource), new(Role), new(RoleResourceRel), new(RoleBackendUserRel))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func ActivationCodeTBName() string {
	return TableName("pc_active_infos")
}

func PcBelongedTBName() string {
	return TableName("pc_belonged")
}

// PcMchtInfosTBName 获取 pc_mcht_infos 对应的表名称
func PcMchtInfosTBName() string {
	return TableName("pc_mcht_infos")
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func BackendUserTBName() string {
	return TableName("backend_user")
}

// ResourceTBName 获取 Resource 对应的表名称
func ResourceTBName() string {
	return TableName("resource")
}

// RoleTBName 获取 Role 对应的表名称
func RoleTBName() string {
	return TableName("role")
}

// RoleResourceRelTBName 角色与资源多对多关系表
func RoleResourceRelTBName() string {
	return TableName("role_resource_rel")
}

// RoleBackendUserRelTBName 角色与用户多对多关系表
func RoleBackendUserRelTBName() string {
	return TableName("role_backenduser_rel")
}

// 终端资源表
func PcSourceAssignInfosTBName() string {
	return TableName("pc_source_assign_infos")
}
func Pc_source_infosTBName() string {
	return TableName("pc_source_infos")
}
