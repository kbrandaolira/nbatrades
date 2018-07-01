package structs

type teams struct{
	Id int `json:"id"`,
	Name string `json:"name"`,
	Shortname string `json:"shortname"`,
	Country string `json:"country"`,
	Conference string `json:"conference"`,
	Division string `json:"division"`,
	transfers []transfers `json:"transfers"`
}