package commonapi

import (
	"context"
	"fmt"
	"log"

	waitlist "github.com/tchssk/reserve-with-google-goa/gen/waitlist"
	"goa.design/goa/v3/security"
)

// waitlist service example implementation.
// The example methods log the requests and return zero values.
type waitlistsrvc struct {
	logger *log.Logger
}

// NewWaitlist returns the waitlist service implementation.
func NewWaitlist(logger *log.Logger) waitlist.Service {
	return &waitlistsrvc{logger}
}

// BasicAuth implements the authorization logic for service "waitlist" for the
// "basic" security scheme.
func (s *waitlistsrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
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

// BatchGetWaitEstimates implements batch_get_wait_estimates.
func (s *waitlistsrvc) BatchGetWaitEstimates(ctx context.Context, p *waitlist.BatchGetWaitEstimatesPayload) (res *waitlist.BatchGetWaitEstimatesResponse, err error) {
	res = &waitlist.BatchGetWaitEstimatesResponse{}
	s.logger.Print("waitlist.batch_get_wait_estimates")
	return
}

// CreateWaitlistEntry implements create_waitlist_entry.
func (s *waitlistsrvc) CreateWaitlistEntry(ctx context.Context, p *waitlist.CreateWaitlistEntryPayload) (res *waitlist.CreateWaitlistEntryResponse, err error) {
	res = &waitlist.CreateWaitlistEntryResponse{}
	s.logger.Print("waitlist.create_waitlist_entry")
	return
}

// GetWaitlistEntry implements get_waitlist_entry.
func (s *waitlistsrvc) GetWaitlistEntry(ctx context.Context, p *waitlist.GetWaitlistEntryPayload) (res *waitlist.GetWaitlistEntryResponse, err error) {
	res = &waitlist.GetWaitlistEntryResponse{}
	s.logger.Print("waitlist.get_waitlist_entry")
	return
}

// DeleteWaitlistEntry implements delete_waitlist_entry.
func (s *waitlistsrvc) DeleteWaitlistEntry(ctx context.Context, p *waitlist.DeleteWaitlistEntryPayload) (err error) {
	s.logger.Print("waitlist.delete_waitlist_entry")
	return
}
