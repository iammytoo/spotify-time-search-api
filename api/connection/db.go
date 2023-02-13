package connection

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// DB初期化
func Init() {
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

// DB接続終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
