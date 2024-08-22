package service

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/stretchr/testify/mock"
)

type NotificationMock struct {
	mock.Mock
}

func NewNotificationMock() *NotificationMock {
	return &NotificationMock{
		Mock: mock.Mock{},
	}
}

func (s *NotificationMock) Shipper(ctx context.Context, data *entity.Shipper) error {
	arguments := s.Mock.Called(ctx, data)

	return arguments.Error(0)
}
