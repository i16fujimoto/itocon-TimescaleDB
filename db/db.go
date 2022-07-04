package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
  	"github.com/jinzhu/gorm"
)

var GDB *gorm.DB

// Databaseの初期化
func InitDB() {
	var err error
	GDB, err = gorm.Open("postgres", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}
	fmt.Println("db connected: ", GDB)

	// 詳細なログを表示（エラー以外も）
	GDB.LogMode(true)
	
	// DBコネクションプールの設定
	sqlDB := GDB.DB()
	sqlDB.SetMaxOpenConns(25) // 同時接続数（最大DBにいくつコネクションできるか）（デフォ無制限）
	sqlDB.SetMaxIdleConns(25) // コネクションのアイドル数（コネクションの再利用できる数）（デフォは全て新規接続つまり０）
	sqlDB.SetConnMaxLifetime(25 * time.Second)	// コネクションをどれだけの時間維持しておくか（コネクションを再利用できる最大の期間を設定）（デフォは無制限→接続に失敗するまで）
}

// DB取得
// func GetDB() *gorm.DB {
// 	return GDB
// }

// DB接続終了
// func Close() {
// 	if err := GDB.Close(); err != nil {
// 		panic(err)
// 	}
// }