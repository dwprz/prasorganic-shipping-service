package client

import (
	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/delivery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// this main grpc client
type Grpc struct {
	Order     delivery.OrderGrpc
	orderConn *grpc.ClientConn
}

func NewGrpc(ugd delivery.OrderGrpc, orderConn *grpc.ClientConn) *Grpc {

	return &Grpc{
		Order:     ugd,
		orderConn: orderConn,
	}
}

func (g *Grpc) Close() {
	if err := g.orderConn.Close(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "client.Grpc/Close", "section": "orderConn.Close"}).Error(err)
	}
}
