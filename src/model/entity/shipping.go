package entity

type ShippingOrder struct {
	Consignee   Consignee   `json:"consignee" validate:"required"`
	Consigner   Consigner   `json:"consigner" validate:"required"`
	Courier     Courier     `json:"courier" validate:"required"`
	Coverage    string      `json:"coverage" validate:"required"`
	Destination Destination `json:"destination" validate:"required"`
	ExternalId  string      `json:"external_id" validate:"required"`
	Origin      Origin      `json:"origin" validate:"required"`
	Package     Package     `json:"package" validate:"required"`
	PaymentType string      `json:"payment_type" validate:"required"`
}

type Consignee struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Consigner struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Courier struct {
	COD          bool `json:"cod"`
	RateId       int  `json:"rate_id" validate:"required"`
	UseInsurance bool `json:"use_insurance"`
}

type Destination struct {
	Address string `json:"address" validate:"required"`
	AreaId  int    `json:"area_id" validate:"required"`
	Lat     string `json:"lat" validate:"required"`
	Lng     string `json:"lng" validate:"required"`
}

type Origin struct {
	Address string `json:"address" validate:"required"`
	AreaId  int    `json:"area_id" validate:"required"`
	Lat     string `json:"lat" validate:"required"`
	Lng     string `json:"lng" validate:"required"`
}

type Package struct {
	Height      int     `json:"height" validate:"required"`
	Items       []Item  `json:"items" validate:"required,dive"`
	Length      int     `json:"length" validate:"required"`
	PackageType int     `json:"package_type" validate:"required"`
	Price       int     `json:"price" validate:"required"`
	Weight      float32 `json:"weight" validate:"required"`
	Width       int     `json:"width" validate:"required"`
}

type Item struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Qty   int    `json:"qty" validate:"required"`
}
