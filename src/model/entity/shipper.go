package entity

type Country struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Province struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Lat     float64  `json:"lat"`
	Lng     float64  `json:"lng"`
	Country *Country `json:"country,omitempty"`
}

type City struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
	Province *Province `json:"province,omitempty"`
	Country  *Country  `json:"country,omitempty"`
}

type Pagination struct {
	CurrentPage   int `json:"current_page"`
	TotalPages    int `json:"total_pages"`
	TotalElements int `json:"total_elements"`
}
