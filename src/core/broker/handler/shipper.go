package handler

import (
	"context"
	"encoding/json"

	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type ShipperKafka struct {
	NotifService service.Notification
}

func NewShipperKafka(ns service.Notification) *ShipperKafka {
	return &ShipperKafka{
		NotifService: ns,
	}
}

func (s *ShipperKafka) ProcessMessage(ctx context.Context, msg kafka.Message) {
	const maxRetries = 3
	for i := 0; i < maxRetries; i++ {

		shipperMsg := new(entity.Shipper)
		if err := json.Unmarshal(msg.Value, &shipperMsg); err != nil {
			log.Logger.WithFields(logrus.Fields{"location": "handler.ShipperKafka", "section": "json.Unmarshal"})
			continue
		}

		if err := s.NotifService.Shipper(ctx, shipperMsg); err != nil {
			log.Logger.WithFields(logrus.Fields{"location": "handler.ShipperKafka", "section": "NotifService.Shipper"})
			continue
		}

		break
	}
}
