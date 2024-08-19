package delivery

import (
	"context"

	pb "github.com/dwprz/prasorganic-proto/protogen/order"
	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/dwprz/prasorganic-shipping-service/src/core/grpc/interceptor"
	"github.com/dwprz/prasorganic-shipping-service/src/infrastructure/cbreaker"
	"github.com/dwprz/prasorganic-shipping-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/delivery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderGrpcImpl struct {
	client   pb.OrderServiceClient
}

func NewOrderGrpc(unaryRequest *interceptor.UnaryRequest) (delivery.OrderGrpc, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryRequest.AddBasicAuth),
	)

	conn, err := grpc.NewClient(config.Conf.ApiGateway.BaseUrl, opts...)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "delivery.NewOrderGrpc", "section": "grpc.NewClient"}).Fatal(err)
	}

	client := pb.NewOrderServiceClient(conn)

	return &OrderGrpcImpl{
		client:   client,
	}, conn
}

func (o *OrderGrpcImpl) AddShippingId(ctx context.Context, data *pb.AddShippingIdReq) error {
	_, err := cbreaker.OrderGrpc.Execute(func() (any, error) {
		_, err := o.client.AddShippingId(ctx, data)
		return nil, err
	})

	return err
}