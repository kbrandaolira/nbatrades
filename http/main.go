package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:dds3@1@/nbatrades")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func main() {
	InitDB()
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Initializing Server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
