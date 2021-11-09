package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConf ...
type DBConf struct {
	Type            string `toml:"type"`
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	User            string `toml:"user"`
	Password        string `json:"-" toml:"passwd"`
	Name            string `toml:"name"`
	PoolSize        int    `toml:"pool_size"`
	CommandTimeOut  int    `toml:"command_time_out"`
	ConnMaxIdleTime int    `toml:"conn_max_idle_time"`
}

// Setup setups a database connection
func Setup() (*gorm.DB, error) {
	dbConf := &DBConf{
		Type:            "mysql",
		Host:            "127.0.0.1",
		Port:            3306,
		User:            "root",
		Password:        "1234567",
		Name:            "test",
		PoolSize:        100,
		ConnMaxIdleTime: 60,
	}

	dbClient, err := gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConf.User,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.Name,
		)), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbInst, err := dbClient.DB()
	if err != nil {
		return nil, err
	}

	dbInst.SetMaxIdleConns(dbConf.PoolSize)
	dbInst.SetMaxOpenConns(dbConf.PoolSize)
	dbInst.SetConnMaxIdleTime(time.Duration(dbConf.ConnMaxIdleTime) * time.Minute)

	return dbClient, nil
}
