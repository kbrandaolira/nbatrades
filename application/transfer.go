package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Transfer struct {
	Id     int    `json:"id"`
	Date   string `json:"date"`
	Player string `json:"player"`
	Status string `json:"status"`
}

func TransfersHandler(w http.ResponseWriter, r *http.Request) {

	sid := ""

	option := ""

	if strings.Contains(r.URL.Path, "asbuyer") {

		sid = strings.TrimPrefix(r.URL.Path, "/transfers/asbuyer/")
		option = "asbuyer"

	} else {

		sid = strings.TrimPrefix(r.URL.Path, "/transfers/asseller/")
		option = "asseller"

	}

	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && option == "asbuyer":
		TransfersAsBuyerByTeam(w, r, id)
	case r.Method == "GET" && option == "asseller":
		TransfersAsSellerByTeam(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry...")
	}

}

func TransfersAsSellerByTeam(w http.ResponseWriter, r *http.Request, id int) {

	db, _ := sql.Open("mysql", "root:dds3@1@/nbatrades")

	defer db.Close()

	rows, _ := db.Query("select id,date_format(date,'%d/%m/%Y'),player,status from transfers t where t.seller = ? order by t.player asc", id)

	defer rows.Close()

	var transfers []Transfer

	for rows.Next() {

		var t Transfer
		rows.Scan(&t.Id, &t.Date, &t.Player, &t.Status)

		transfers = append(transfers, t)

	}

	json, _ := json.Marshal(transfers)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))

}

func TransfersAsBuyerByTeam(w http.ResponseWriter, r *http.Request, id int) {

	db, _ := sql.Open("mysql", "root:dds3@1@/nbatrades")

	defer db.Close()

	rows, _ := db.Query("select id,date,player,status from transfers t where t.buyer = ? order by t.player asc", id)

	defer rows.Close()

	var transfers []Transfer

	for rows.Next() {

		var t Transfer
		rows.Scan(&t.Id, &t.Date, &t.Player, &t.Status)

		transfers = append(transfers, t)

	}

	json, _ := json.Marshal(transfers)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))

}
