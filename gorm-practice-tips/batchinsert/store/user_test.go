package store

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/pborman/uuid"
	"github.com/qqzeng/gorm-practice-tips/batchinsert/db"
	"github.com/qqzeng/gorm-practice-tips/batchinsert/proto"
	"github.com/qqzeng/gorm-practice-tips/batchinsert/store/model"
)

var storeMgr Manager

func setup() {
	var err error
	dbClient, err := db.Setup()
	if err != nil {
		panic(err)
	}

	storeMgr = NewManager(dbClient)
}

func shutdown() {
	fmt.Println("all test done")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

// BatchInsert ...
func Test_User_SaveMultiUser(t *testing.T) {
	var users = []*model.User{
		{Name: "turlayon", UserUUID: uuid.New()},
		{Name: "qqzeng", UserUUID: uuid.New()},
		{Name: "qiaoqiaozeng", UserUUID: uuid.New()},
		{Name: "zs", UserUUID: uuid.New()},
		{Name: "ls", UserUUID: uuid.New()},
		{Name: "tom", UserUUID: uuid.New()},
		{Name: "bob", UserUUID: uuid.New()},
	}

	if err := storeMgr.SaveMultiUser(users); err != nil {
		t.Fatal(err)
	}

	ids := make([]uint64, 0, len(users))
	for _, u := range users {
		ids = append(ids, u.BaseModel.ID)
	}

	t.Log("save multi user", ids)
}

func Test_User_DeleteMultiUser(t *testing.T) {
	queryParam := proto.QueryCommonParams{
		Limit:  math.MaxUint32,
		Offset: 0,
		Filters: proto.Filters{
			{
				Name:  "Name",
				Type:  "Fuzzy",
				Value: "zeng",
			},
		},
	}
	param, err := Parse(queryParam, nil)
	if err != nil {
		t.Fatal(err)
	}

	toDeleteUsers, _, err := storeMgr.ListUser(param)
	if err != nil {
		t.Fatal(err)
	}

	if err := storeMgr.DeleteMultiUser(toDeleteUsers); err != nil {
		t.Fatal(err)
	}
}

func Test_User_CreateWithJsonString(t *testing.T) {
	testStr := `{
		"ID": 1,
		"Name": "test",
		"UserUUID": "24aa0fa9-29aa-4db2-9c21-6980c8f68d78",
		"CreatedAt": "2021-08-09 10:00:23",
		"UpdatedAt": "2021-08-09 11:00:23"
	}`

	var user model.User
	if err := json.Unmarshal([]byte(testStr), &user); err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}
