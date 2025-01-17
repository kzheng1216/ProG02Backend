package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	driverName = "mysql"
	host       = "106.15.33.153"
	port       = "3306"
	username   = "mysql"
	password   = "Zaq1@wsx"
	dbName     = "mydb01_dev"
)

func GetDB() (*sql.DB, error) {
	// 替换为你的 MySQL 配置信息

	dsnList := []string{
		//"mysql:Zaq1@wsx@tcp(106.15.33.153:3306)/mydb01_dev"
		username, ":", password, "@tcp(", host, ":", port, ")/", dbName,
	}
	dsn := strings.Join(dsnList, "")
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
