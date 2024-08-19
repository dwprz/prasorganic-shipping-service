package service

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Notification interface {
	Shipper(ctx context.Context, data *entity.Shipper) error
}
