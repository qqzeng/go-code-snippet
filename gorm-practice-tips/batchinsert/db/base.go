package db

import (
	"time"

	"github.com/qqzeng/gorm-practice-tips/batchinsert/proto"

	"gorm.io/gorm"
)

// BaseModel ...
type BaseModel struct {
	ID        uint64         `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt string         `gorm:"column:created_at" json:"CreatedAt"`
	UpdatedAt string         `gorm:"column:updated_at" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"DeletedAt"` // soft delete
}

// BeforeCreate ...
func (v *BaseModel) BeforeCreate(tx *gorm.DB) error {
	v.CreatedAt = time.Now().Format(proto.TimeDefaultLayout)
	v.UpdatedAt = time.Now().Format(proto.TimeDefaultLayout)
	return nil
}

// BeforeUpdate ...
func (v *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	v.UpdatedAt = time.Now().Format(proto.TimeDefaultLayout)
	return nil
}
