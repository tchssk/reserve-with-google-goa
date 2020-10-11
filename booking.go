package commonapi

import (
	"context"
	"fmt"
	"log"

	booking "github.com/tchssk/reserve-with-google-goa/gen/booking"
	"goa.design/goa/v3/security"
)

// booking service example implementation.
// The example methods log the requests and return zero values.
type bookingsrvc struct {
	logger *log.Logger
}

// NewBooking returns the booking service implementation.
func NewBooking(logger *log.Logger) booking.Service {
	return &bookingsrvc{logger}
}

// BasicAuth implements the authorization logic for service "booking" for the
// "basic" security scheme.
func (s *bookingsrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
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

// BatchAvailabilityLookup implements batch_availability_lookup.
func (s *bookingsrvc) BatchAvailabilityLookup(ctx context.Context, p *booking.BatchAvailabilityLookupPayload) (res *booking.BatchAvailabilityLookupResponse, err error) {
	res = &booking.BatchAvailabilityLookupResponse{}
	s.logger.Print("booking.batch_availability_lookup")
	return
}

// CheckAvailability implements check_availability.
func (s *bookingsrvc) CheckAvailability(ctx context.Context, p *booking.CheckAvailabilityPayload) (res *booking.CheckAvailabilityResponse, err error) {
	res = &booking.CheckAvailabilityResponse{}
	s.logger.Print("booking.check_availability")
	return
}

// CreateBooking implements create_booking.
func (s *bookingsrvc) CreateBooking(ctx context.Context, p *booking.CreateBookingPayload) (res *booking.CreateBookingResponse, err error) {
	res = &booking.CreateBookingResponse{}
	s.logger.Print("booking.create_booking")
	return
}

// UpdateBooking implements update_booking.
func (s *bookingsrvc) UpdateBooking(ctx context.Context, p *booking.UpdateBookingPayload) (res *booking.UpdateBookingResponse, err error) {
	res = &booking.UpdateBookingResponse{}
	s.logger.Print("booking.update_booking")
	return
}

// GetBookingStatus implements get_booking_status.
func (s *bookingsrvc) GetBookingStatus(ctx context.Context, p *booking.GetBookingStatusPayload) (res *booking.GetBookingStatusResponse, err error) {
	res = &booking.GetBookingStatusResponse{}
	s.logger.Print("booking.get_booking_status")
	return
}

// ListBookings implements list_bookings.
func (s *bookingsrvc) ListBookings(ctx context.Context, p *booking.ListBookingsPayload) (res *booking.ListBookingsResponse, err error) {
	res = &booking.ListBookingsResponse{}
	s.logger.Print("booking.list_bookings")
	return
}
