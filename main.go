package main

import (
	"itocon/server"
	"itocon/db"
)

func main() {
	
	// DBの初期化
	db.InitDB()

	defer db.GDB.Close()

	//サーバの初期化
	server.Init()

}