package controllers

import "github.com/kiplimoboor/favorit/api/database"

type Controller struct {
	db *database.SQLiteDB
}

func NewController(db *database.SQLiteDB) *Controller {
	return &Controller{db: db}
}
