package main

import (
	"fmt"
	"log"
)

func main() {
	db, _ := NewSQLiteDB()
	db.Init()

	u := NewUser("jack", "will", "jack@will.com", "mysecret")
	if err := db.AddUser(u); err != nil {
		log.Fatal("error adding user")
	}

	user, err := db.GetUserByEmail(u.Email)
	if err != nil {
		log.Fatal("can't get user")
	}

	fmt.Println(*user)
}
