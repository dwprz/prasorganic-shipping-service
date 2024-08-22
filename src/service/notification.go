package service

import (
	"context"

	"github.com/dwprz/prasorganic-proto/protogen/order"
	"github.com/dwprz/prasorganic-shipping-service/src/common/helper"
	"github.com/dwprz/prasorganic-shipping-service/src/core/grpc/client"
	v "github.com/dwprz/prasorganic-shipping-service/src/infrastructure/validator"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type NotificationImpl struct {
	grpcClient    *client.Grpc
	shippingCache cache.Shipping
}

func NewNotification(gc *client.Grpc, cs cache.Shipping) service.Notification {
	return &NotificationImpl{
		grpcClient:    gc,
		shippingCache: cs,
	}
}

func (n *NotificationImpl) Shipper(ctx context.Context, data *entity.Shipper) (err error) {
	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	shippingId := data.ShippingId
	trackingData := helper.FormatTrackingData(data)

	n.shippingCache.UpdateTracking(ctx, shippingId, trackingData)

	if data.ExternalStatus.Code == 1000 {
		err = n.grpcClient.Order.UpdateStatus(ctx, &order.UpdateStatusReq{
			OrderId: data.OrderId,
			Status:  string(entity.IN_PROGRESS),
		})
	}

	if data.ExternalStatus.Code == 2000 ||
		data.ExternalStatus.Code == 2010 ||
		data.ExternalStatus.Code == 3000 {

		err = n.grpcClient.Order.UpdateStatus(ctx, &order.UpdateStatusReq{
			OrderId: data.OrderId,
			Status:  string(entity.COMPLETED),
		})
	}

	if data.ExternalStatus.Code == 1340 {
		err = n.grpcClient.Order.UpdateStatus(ctx, &order.UpdateStatusReq{
			OrderId: data.OrderId,
			Status:  string(entity.RETURN_PROCESSING),
		})
	}

	if data.ExternalStatus.Code == 1370 {
		err = n.grpcClient.Order.UpdateStatus(ctx, &order.UpdateStatusReq{
			OrderId: data.OrderId,
			Status:  string(entity.FAILED),
		})
	}

	if data.ExternalStatus.Code == 1380 {
		err = n.grpcClient.Order.UpdateStatus(ctx, &order.UpdateStatusReq{
			OrderId: data.OrderId,
			Status:  string(entity.LOST_OR_DAMAGED),
		})
	}

	if data.ExternalStatus.Code == 999 {
		err = n.grpcClient.Order.UpdateStatus(ctx, &order.UpdateStatusReq{
			OrderId: data.OrderId,
			Status:  string(entity.CANCELLED),
		})
	}

	return err
}
