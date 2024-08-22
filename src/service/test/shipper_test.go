package test

import (
	"context"
	"testing"
	"time"

	"github.com/dwprz/prasorganic-proto/protogen/order"
	"github.com/dwprz/prasorganic-shipping-service/src/core/grpc/client"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/mock/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/mock/delivery"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	serviceimpl "github.com/dwprz/prasorganic-shipping-service/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestSetvice_ShipperNotif$ -v ./src/service/test/ -count=1

type ShipperNotifTestSuite struct {
	suite.Suite
	notifService service.Notification
	orderGrpc    *delivery.OrderGrpcMock
}

func (s *ShipperNotifTestSuite) SetupSuite() {
	s.orderGrpc = delivery.NewOrderGrpcMock()
	orderConn := new(grpc.ClientConn)

	grpcClient := client.NewGrpc(s.orderGrpc, orderConn)
	shippingCache := cache.NewShippingMock()

	s.notifService = serviceimpl.NewNotification(grpcClient, shippingCache)
}

func (s *ShipperNotifTestSuite) Test_COMPLETED() {
	req := s.CreateShipper()
	req.ExternalStatus.Code = 2000

	s.MockOrderGrpc_UpdateStatus(req.OrderId, nil)

	err := s.notifService.Shipper(context.Background(), req)
	assert.NoError(s.T(), err)
}

func (s *ShipperNotifTestSuite) Test_IN_PROGRESS() {
	req := s.CreateShipper()
	req.ExternalStatus.Code = 1000

	s.MockOrderGrpc_UpdateStatus(req.OrderId, nil)

	err := s.notifService.Shipper(context.Background(), req)
	assert.NoError(s.T(), err)
}

func (s *ShipperNotifTestSuite) Test_RETURN_PROCESSING() {
	req := s.CreateShipper()
	req.ExternalStatus.Code = 1340

	s.MockOrderGrpc_UpdateStatus(req.OrderId, nil)

	err := s.notifService.Shipper(context.Background(), req)
	assert.NoError(s.T(), err)
}

func (s *ShipperNotifTestSuite) Test_FAILED() {
	req := s.CreateShipper()
	req.ExternalStatus.Code = 1370

	s.MockOrderGrpc_UpdateStatus(req.OrderId, nil)

	err := s.notifService.Shipper(context.Background(), req)
	assert.NoError(s.T(), err)
}

func (s *ShipperNotifTestSuite) Test_LOST_OR_DAMAGED() {
	req := s.CreateShipper()
	req.ExternalStatus.Code = 1380

	s.MockOrderGrpc_UpdateStatus(req.OrderId, nil)

	err := s.notifService.Shipper(context.Background(), req)
	assert.NoError(s.T(), err)
}

func (s *ShipperNotifTestSuite) Test_CANCELED() {
	req := s.CreateShipper()
	req.ExternalStatus.Code = 999

	s.MockOrderGrpc_UpdateStatus(req.OrderId, nil)

	err := s.notifService.Shipper(context.Background(), req)
	assert.NoError(s.T(), err)
}

func (s *ShipperNotifTestSuite) MockOrderGrpc_UpdateStatus(orderId string, returnArg1 error) {

	s.orderGrpc.Mock.On("UpdateStatus", mock.Anything, mock.MatchedBy(func(req *order.UpdateStatusReq) bool {
		return orderId == req.OrderId && req.Status != ""
	})).Return(returnArg1)
}

func (s *ShipperNotifTestSuite) CreateShipper() *entity.Shipper {
	return &entity.Shipper{
		Auth:            "auth_token_example",
		ShippingId:      "ship_123456789",
		TrackingId:      "track_987654321",
		OrderTrackingId: "order_track_123456",
		OrderId:         "order_987654",
		StatusDate:      time.Now(),
		Internal: entity.InternalExternal{
			Id:          1,
			Name:        "Internal Name",
			Description: "Internal Description",
		},
		External: entity.InternalExternal{
			Id:          2,
			Name:        "External Name",
			Description: "External Description",
		},
		InternalStatus: entity.Status{
			Code:        100,
			Name:        "Internal Status Name",
			Description: "Internal Status Description",
		},
		ExternalStatus: entity.Status{
			Code:        2000,
			Name:        "External Status Name",
			Description: "External Status Description",
		},
		AWB: "awb_123456789",
	}
}

func TestSetvice_ShipperNotif(t *testing.T) {
	suite.Run(t, new(ShipperNotifTestSuite))
}
