package main

import (
	"itocon/db"
	"itocon/entry/db_entry"
)

func main() {
	
	// DBã®åæå
	db.InitDB()

	// Migrattion
	db.GDB.Migrator().AutoMigrate(&db_entry.Sensor{})
}