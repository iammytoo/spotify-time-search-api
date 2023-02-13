package connection

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/mirito333/spotify-time-search-api/api/model"
)

var (
	db  *gorm.DB
	err error
)

// DB初期化
func Init() {

	// 環境変数取得
	godotenv.Load(".env")

	// DB接続
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&model.PlayList{},&model.Track{})
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
