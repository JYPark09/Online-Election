package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabaseManager(dbfile string) {
	var err error

	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatalln("[dbms] cannot open database ", err)
	}

	log.Println("[dbms] initialized")
}

func shutdownDatabaseManager() {
	db.Close()

	log.Println("[dbms] shutdowned")
}

func getUserInfo(id int) *User {
	usr := new(User)

	usr.id = id
	err := db.QueryRow("SELECT pw FROM Users WHERE id=" + strconv.Itoa(id)).Scan(&usr.pw)
	if err != nil {
		return nil
	}

	return usr
}
