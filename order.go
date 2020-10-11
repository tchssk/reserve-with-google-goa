package commonapi

import (
	"context"
	"fmt"
	"log"

	order "github.com/tchssk/reserve-with-google-goa/gen/order"
	"goa.design/goa/v3/security"
)

// order service example implementation.
// The example methods log the requests and return zero values.
type ordersrvc struct {
	logger *log.Logger
}

// NewOrder returns the order service implementation.
func NewOrder(logger *log.Logger) order.Service {
	return &ordersrvc{logger}
}

// BasicAuth implements the authorization logic for service "order" for the
// "basic" security scheme.
func (s *ordersrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// CheckOrderFulfillability implements check_order_fulfillability.
func (s *ordersrvc) CheckOrderFulfillability(ctx context.Context, p *order.CheckOrderFulfillabilityPayload) (res *order.CheckOrderFulfillabilityResponse, err error) {
	res = &order.CheckOrderFulfillabilityResponse{}
	s.logger.Print("order.check_order_fulfillability")
	return
}

// CreateOrder implements create_order.
func (s *ordersrvc) CreateOrder(ctx context.Context, p *order.CreateOrderPayload) (res *order.CreateOrderResponse, err error) {
	res = &order.CreateOrderResponse{}
	s.logger.Print("order.create_order")
	return
}

// ListOrders implements list_orders.
func (s *ordersrvc) ListOrders(ctx context.Context, p *order.ListOrdersPayload) (res *order.ListOrdersResponse, err error) {
	res = &order.ListOrdersResponse{}
	s.logger.Print("order.list_orders")
	return
}
