package main

import (
	"github.com/kiplimoboor/favorit/api"
	"github.com/kiplimoboor/favorit/api/database"
	"os"
)

func main() {
	os.Remove("favorit.db")
	db, _ := database.NewSQLiteDB()
	db.Init()
	server := api.NewServer(":8080", db)
	server.Start()
}
