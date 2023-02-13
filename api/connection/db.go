package connection

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.PlayList{},&model.Track{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
