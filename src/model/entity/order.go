package entity

type OrderStatus string

const (
	PENDING_PAYMENT   OrderStatus = "PENDING_PAYMENT"
	PAID              OrderStatus = "PAID"
	IN_PROGRESS       OrderStatus = "IN_PROGRESS"
	COMPLETED         OrderStatus = "COMPLETED"
	CANCELLED         OrderStatus = "CANCELLED"
	FAILED            OrderStatus = "FAILED"
	REFUND_PROCESSING OrderStatus = "REFUND_PROCESSING"
	REFUND_COMPLETED  OrderStatus = "REFUND_COMPLETED"
	RETURN_PROCESSING OrderStatus = "RETURN_PROCESSING"
	RETURN_COMPLETED  OrderStatus = "RETURN_COMPLETED"
	LOST_OR_DAMAGED   OrderStatus = "LOST_OR_DAMAGED"
)
