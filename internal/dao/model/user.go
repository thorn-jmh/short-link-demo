package model

const UserTable = "Users"

type User struct {
	ID       uint   `json:"id"`       // 用户ID
	Email    string `json:"email"`    // 用户邮箱
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码

	// edge
	Links []Link `json:"-" gorm:"foreignKey:OwnerID;references:ID"`
}
