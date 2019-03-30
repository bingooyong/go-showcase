package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"study/tools/config"
)

var db *sql.DB
var dbx *sqlx.DB

func Config(config *config.Config) {
	db, _ = sql.Open("mysql", config.Mysql.Host)
	db.SetMaxOpenConns(config.Mysql.Max)
	db.SetMaxIdleConns(config.Mysql.MaxIdle)
	ping := db.Ping()
	if ping != nil {
		logrus.Errorf("can not connect to mysql %s", config.Mysql.Host)
	}
	dbx = sqlx.NewDb(db, "mysql")

}

func Connection() *sql.DB {
	return db
}

func ConnectionX() *sqlx.DB {
	return dbx
}
