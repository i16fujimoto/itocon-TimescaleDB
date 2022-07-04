package main

import (
	"itocon/db"
	"itocon/entry/db_entry"
)

func main() {
	
	// DBの初期化
	db.InitDB()

	// Migrattion
	db.GDB.Migrator().AutoMigrate(&db_entry.Sensor{})
}