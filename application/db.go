package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {

	db, err := sql.Open("mysql", "root:dds3@1@/nbatrades")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db

}
