package delivery

import (
	"context"

	pb "github.com/dwprz/prasorganic-proto/protogen/order"
)


type OrderGrpc interface {
	AddShippingId(ctx context.Context, data *pb.AddShippingIdReq) error
	UpdateStatus(ctx context.Context, data *pb.UpdateStatusReq) error 
}
