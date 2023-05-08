package models

type Access struct {
	Id          int
	ModuleName  string // 模块名称
	ActionName  string // 操作名称
	Type        int    // 节点类型 ：1、表示模块 2、 表示菜单 3、操作
	Url         string // 路由器跳转地址
	ModuleId    int    // 此module_id 和当前模型的id相关联
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId; references:Id"`
}

func (Access) TableName() string {
	return "access"
}
