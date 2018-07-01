package structs

type transfers struct{
	Id int `json:"id"`,
	Date time `json:"date"`,
	Seller teams `json:"seller"`,
	Buyer teams `json:"buyer"`,
	Player string `json:"player"`,
	Status string `json:"status"`
}