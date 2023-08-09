package model

import (
	"gopkg.in/guregu/null.v4"
)

const LinkTable = "Links"

type Link struct {
	ID        uint      `json:"-" gorm:"primary_key;auto_increment"` // 自增ID
	Short     string    `json:"short" gorm:"unique"`                 // 短链，全局唯一
	Origin    string    `json:"origin"`                              // 原始链接
	Comment   string    `json:"comment"`                             // 备注信息
	StartTime null.Time `json:"start_time"`                          // 起始时间，UTC
	EndTime   null.Time `json:"end_time"`                            // 到期时间，UTC
	Active    null.Bool `json:"active"`

	// edge
	Owner   *User `json:"-" gorm:"foreignKey:OwnerID;references:ID"`
	OwnerID uint  `json:"-"`
}
