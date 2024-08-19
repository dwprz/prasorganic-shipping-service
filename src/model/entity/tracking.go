package entity

import (
	"time"
)

type Tracking struct {
	Trackings []TrackingData `json:"trackings"`
}

type TrackingData struct {
	ShipperStatus  Status    `json:"shipper_status"`
	LogisticStatus Status    `json:"logistic_status"`
	CreatedDate    time.Time `json:"created_date"`
}

type Status struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
