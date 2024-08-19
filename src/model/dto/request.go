package dto

type PricingReq struct {
	COD         bool `json:"cod"`
	Destination struct {
		AreaId   int    `json:"area_id" validate:"required"`
		Lat      string `json:"lat" validate:"required"`
		Lng      string `json:"lng" validate:"required"`
		SuburbId int    `json:"suburb_id" validate:"required"`
	} `json:"destination" validate:"required"`
	ForOrder  bool `json:"for_order"`
	Height    int  `json:"height" validate:"required"`
	ItemValue int  `json:"item_value" validate:"required"`
	Length    int  `json:"length" validate:"required"`
	Limit     int  `json:"limit" validate:"required"`
	Origin    struct {
		AreaId   int    `json:"area_id" validate:"required"`
		Lat      string `json:"lat" validate:"required"`
		Lng      string `json:"lng" validate:"required"`
		SuburbId int    `json:"suburb_id" validate:"required"`
	} `json:"origin" validate:"required"`
	Page   int      `json:"page" validate:"required"`
	SortBy []string `json:"sort_by" validate:"dive,required"`
	Weight float32  `json:"weight" validate:"required"`
	Width  int      `json:"width" validate:"required"`
}

type CreateLabelReq struct {
	Id   []string `json:"id" validate:"dive,required"`
	Type string   `json:"type" validate:"required"`
}
