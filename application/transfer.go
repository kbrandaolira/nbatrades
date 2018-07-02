package main

import (
	"time"
)

type Transfer struct {
	Id     int       `json:"id"`
	Date   time.Time `json:"date"`
	Seller Team      `json:"seller"`
	Buyer  Team      `json:"buyer"`
	Player string    `json:"player"`
	Status string    `json:"status"`
}
