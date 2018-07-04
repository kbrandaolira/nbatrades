package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Team struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Shortname  string `json:"shortname"`
	Country    string `json:"country"`
	Conference string `json:"conference"`
	Division   string `json:"division"`
}

func TeamsHandler(w http.ResponseWriter, r *http.Request) {

	sid := ""
	option := ""

	if strings.Contains(r.URL.Path, "logo") {

		sid = strings.TrimPrefix(r.URL.Path, "/teams/logo/")
		option = "logo"

	} else {

		sid = strings.TrimPrefix(r.URL.Path, "/teams/")
		option = "list"

	}

	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && option == "list" && id > 0:
		Get(w, r, id)
	case r.Method == "GET" && option == "list":
		All(w, r)
	case r.Method == "GET" && option == "logo":
		Logo(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry...")
	}

}

func Logo(w http.ResponseWriter, r *http.Request, id int) {

	filename := fmt.Sprintf("logo/%d.png", id)

	//Check if file exists and open
	Openfile, _ := os.Open(filename)
	defer Openfile.Close() //Close after function return

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client

	return

}

func Get(w http.ResponseWriter, r *http.Request, id int) {

	db, _ := sql.Open("mysql", "root:dds3@1@/nbatrades")

	defer db.Close()

	var t Team

	db.QueryRow("select id,name,shortname,country,conference,division from teams where id = ? order by name", id).Scan(&t.Id, &t.Name, &t.Shortname, &t.Country, &t.Conference, &t.Division)

	json, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))

}

func All(w http.ResponseWriter, r *http.Request) {

	db, _ := sql.Open("mysql", "root:dds3@1@/nbatrades")

	defer db.Close()

	rows, _ := db.Query("select id,name,shortname,country,conference,division from teams order by name")

	defer rows.Close()

	var teams []Team

	for rows.Next() {

		var t Team
		rows.Scan(&t.Id, &t.Name, &t.Shortname, &t.Country, &t.Conference, &t.Division)

		teams = append(teams, t)

	}

	json, _ := json.Marshal(teams)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))

}
