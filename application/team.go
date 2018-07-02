package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	sid := strings.TrimPrefix(r.URL.Path, "/teams/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		Get(w, r, id)
	case r.Method == "GET":
		All(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry...")
	}

}

func Get(w http.ResponseWriter, r *http.Request, id int) {

	db := Init()

	var t Team

	db.QueryRow("select id,name from teams where id = ? order by name", id).Scan(&t.Id, &t.Name)

	fmt.Println(t)

	json, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))

}

func All(w http.ResponseWriter, r *http.Request) {

	db := Init()

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
