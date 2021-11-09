package store

import (
	"github.com/qqzeng/gorm-practice-tips/batchinsert/store/model"
	"gorm.io/gorm"
)

// Manager is the store interface
type Manager interface {
	// User
	// SaveUser ...
	SaveUser(*model.User) error
	// SaveMultiUser ...
	SaveMultiUser([]*model.User) error
	// DeleteUser ...
	DeleteUser(*model.User) error
	// DeleteMultiUser ...
	DeleteMultiUser([]*model.User) error
	// UpdateUser ...
	UpdateUser(*model.User) error
	// UpdateMultiUser ...
	UpdateMultiUser([]*model.User) error
	// ListUser...
	ListUser(*CommonParams) ([]*model.User, int, error)
}

const (
	desc = "DESC"
	asc  = "ASC"
)

type manager struct {
	mysqlClient *gorm.DB
}

// NewManager ...
func NewManager(mysqlClient *gorm.DB) Manager {
	return &manager{mysqlClient: mysqlClient}
}
