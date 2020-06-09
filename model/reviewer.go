package model

import "time"

type Reviewer struct {
	Id        uint       `gorm:"column:id" form:"id" json:"id" comment:"" sql:"bigint(20) unsigned,PRI"`
	StaffId   uint64     `gorm:"column:staff_id" form:"staff_id" json:"staff_id" comment:"外部主键" sql:"bigint(20) unsigned,MUL"`
	Name      string     `gorm:"column:name" form:"name" json:"name" comment:"展示名称" sql:"varchar(200)"`
	AvatarUrl string     `gorm:"column:avatar_url" form:"avatar_url" json:"avatar_url" comment:"头像地址" sql:"varchar(200)"`
	Cellphone string     `gorm:"column:cellphone" form:"cellphone" json:"cellphone" comment:"手机号" sql:"varchar(60)"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"创建时间" sql:"datetime"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"更新时间" sql:"datetime"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"删除时间" sql:"datetime"`
}

func (m *Reviewer) TableName() string {
	return "reviewer"
}

