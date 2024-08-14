package dto

type PricingReq struct {
	COD         bool `json:"cod"`
	Destination struct {
		AreaId   int    `json:"area_id"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
		SuburbId int    `json:"suburb_id"`
	} `json:"destination"`
	ForOrder  bool `json:"for_order"`
	Height    int  `json:"height"`
	ItemValue int  `json:"item_value"`
	Length    int  `json:"length"`
	Limit     int  `json:"limit"`
	Origin    struct {
		AreaId   int    `json:"area_id"`
		Lat      string `json:"lat"`
		Lng      string `json:"lng"`
		SuburbId int    `json:"suburb_id"`
	} `json:"origin"`
	Page   int      `json:"page"`
	SortBy []string `json:"sort_by"`
	Weight float32  `json:"weight"`
	Width  int      `json:"width"`
}
