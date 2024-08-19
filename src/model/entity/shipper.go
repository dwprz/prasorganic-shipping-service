package entity

import (
	"time"
)

type Shipper struct {
	Auth            string           `json:"auth"`
	ShippingId      string           `json:"order_id"` // order_id ini berisi shipping_id dari aplikasi ini
	TrackingId      string           `json:"tracking_id"`
	OrderTrackingId string           `json:"order_tracking_id"`
	OrderId         string           `json:"external_id"` // external_id ini berisi order_id dari aplikasi ini
	StatusDate      time.Time        `json:"status_date"`
	Internal        InternalExternal `json:"internal"`
	External        InternalExternal `json:"external"`
	InternalStatus  Status           `json:"internal_status"`
	ExternalStatus  Status           `json:"external_status"`
	AWB             string           `json:"awb"`
}

type InternalExternal struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
