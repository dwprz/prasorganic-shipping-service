package delivery

import (
	"context"

	pb "github.com/dwprz/prasorganic-proto/protogen/order"
	"github.com/stretchr/testify/mock"
)

type OrderGrpcMock struct {
	mock.Mock
}

func NewOrderGrpcMock() *OrderGrpcMock {
	return &OrderGrpcMock{
		Mock: mock.Mock{},
	}
}

func (o *OrderGrpcMock) AddShippingId(ctx context.Context, data *pb.AddShippingIdReq) error {
	arguments := o.Mock.Called(ctx, data)

	return arguments.Error(0)
}

func (o *OrderGrpcMock) UpdateStatus(ctx context.Context, data *pb.UpdateStatusReq) error {
	arguments := o.Mock.Called(ctx, data)

	return arguments.Error(0)
}
