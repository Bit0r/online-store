package model

import (
	"database/sql"
	"fmt"

	"github.com/Bit0r/online-store/conf"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var config struct {
		User     string
		Password string
		Host     string
		Port     uint16
		Name     string
	}

	conf.Unmarshal("database", "online_store_database", &config)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.User, config.Password,
		config.Host, config.Port,
		config.Name)
	db, _ = sql.Open("mysql", dsn)
}
