package entity

type Province struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Country  Country `json:"country"`
}