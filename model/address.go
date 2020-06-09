package model

import "time"

type Address struct {
	ID        uint  `gorm:`
	age       int
	Address1  string //`gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2  string //`gorm:"type:varchar(100);unique"`
	Post      string `gorm:"type:varchar(100)"`
	CreatedAt *time.Time
}

func (m Address) TableName() string{
return "Address"
}