package entity

type Pagination struct {
	CurrentPage   int `json:"current_page"`
	TotalPages    int `json:"total_pages"`
	TotalElements int `json:"total_elements"`
}

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

type Suburb struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
	City     *City     `json:"city,omitempty"`
	Province *Province `json:"province,omitempty"`
	Country  *Country  `json:"country,omitempty"`
}

type Area struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
	Postcode string    `json:"postcode"`
	Suburb   *Suburb   `json:"suburb,omitempty"`
	City     *City     `json:"city,omitempty"`
	Province *Province `json:"province,omitempty"`
	Country  *Country  `json:"country,omitempty"`
}
