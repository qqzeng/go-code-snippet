package model

import "github.com/qqzeng/gorm-practice-tips/batchinsert/db"

// User ...
type User struct {
	db.BaseModel
	UserUUID string `json:"UserUUID" gorm:"column:user_uuid;not null;unique_index"`
	Name     string
}

// TableName ...
func (User) TableName() string {
	return "user"
}
