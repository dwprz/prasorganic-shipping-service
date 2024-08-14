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

type Location struct {
	AreaId       int     `json:"area_id"`
	AreaName     string  `json:"area_name"`
	SuburbId     int     `json:"suburb_id"`
	SuburbName   string  `json:"suburb_name"`
	CityId       int     `json:"city_id"`
	CityName     string  `json:"city_name"`
	ProvinceId   int     `json:"province_id"`
	ProvinceName string  `json:"province_name"`
	CountryId    int     `json:"country_id"`
	CountryName  string  `json:"country_name"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
}
