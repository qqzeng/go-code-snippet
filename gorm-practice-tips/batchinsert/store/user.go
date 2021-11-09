package store

import (
	"fmt"

	"github.com/qqzeng/gorm-practice-tips/batchinsert/store/model"
)

// SaveUser ...
func (m *manager) SaveUser(u *model.User) error {
	return m.mysqlClient.Create(u).Error
}

// SaveMultiUser ...
func (m *manager) SaveMultiUser(users []*model.User) error {
	return m.mysqlClient.Create(users).Error
}

// DeleteUser ...
func (m *manager) DeleteUser(user *model.User) error {
	return m.mysqlClient.Where("`user_uuid` = ? ", user.UserUUID).Delete(model.User{}).Error
}

// DeleteMultiUser ...
func (m *manager) DeleteMultiUser(users []*model.User) error {
	userUUIDs := make([]string, 0, len(users))
	for _, u := range users {
		userUUIDs = append(userUUIDs, u.UserUUID)
	}

	//return m.mysqlClient.Delete(&model.User{}, ids).Error // delete with ID
	return m.mysqlClient.Where("user_uuid in (?)", userUUIDs).Delete(&model.User{}).Error
}

// UpdateUser ...
func (m *manager) UpdateUser(user *model.User) error {
	//return m.mysqlClient.Model(&user).Select("*").Updates(user).Error
	return m.mysqlClient.Save(user).Error
}

// UpdateMultiUser ...
func (m *manager) UpdateMultiUser(users []*model.User) error {
	userUUIDs := make([]string, 0, len(users))
	for _, u := range users {
		userUUIDs = append(userUUIDs, u.UserUUID)
	}

	var dbUsers []*model.User
	if err := m.mysqlClient.Where("user_uuid in (?)", userUUIDs).Save(&dbUsers).Error; err != nil {
		return err
	}

	userDict := make(map[string]*model.User)
	for _, u := range dbUsers {
		userDict[u.UserUUID] = u
	}

	for _, u := range users {
		u.ID = userDict[u.UserUUID].ID
	}

	tx := m.mysqlClient.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, u := range users {
		if err = tx.Save(u).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}

// ListUser ...
func (m *manager) ListUser(params *CommonParams) ([]*model.User, int, error) {
	db := m.mysqlClient
	var hosts []*model.User
	for key, match := range params.Query.Matches {
		switch m := match.(type) {
		case FuzzyMatch:
			db = db.Where(fmt.Sprintf("`%s` like ?", key), "%"+m.Value+"%")
		case ExactMatch:
			// addresses applies JSON query
			if key == "addresses" {
				db = db.Where(fmt.Sprintf("JSON_CONTAINS(addresses, json_array('%s'))", m.Value))
				continue
			}
			db = db.Where(fmt.Sprintf("`%s` = ?", key), m.Value)
		case InMatch:
			db = db.Where(fmt.Sprintf("`%s` IN(?)", key), m.Values)
		case RangeMatch:
			db = db.Where(fmt.Sprintf("`%s` >= ?", key), m.Start)
			db = db.Where(fmt.Sprintf("`%s` <= ?", key), m.End)
		}
	}

	if len(params.Sort) > 0 {
		var orderStr string
		for i := range params.Sort {
			t := asc
			if params.Sort[i][0:1] == "-" {
				t = desc
			}

			orderStr += fmt.Sprintf("`%s` %s,", params.Sort[i][1:], t)
		}

		db = db.Order(orderStr[:len(orderStr)-1])
	}

	if params.Group != "" {
		db = db.Select(params.Group).Group(params.Group)
	}

	var total int64
	db.Model(model.User{}).Count(&total)
	db = db.Offset(params.Page).Limit(params.PageSize)
	if err := db.Find(&hosts).Error; err != nil {
		return nil, 0, err
	}

	return hosts, int(total), nil
}
