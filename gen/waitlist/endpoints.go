// Code generated by goa v3.2.4, DO NOT EDIT.
//
// waitlist endpoints
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package waitlist

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "waitlist" service endpoints.
type Endpoints struct {
	BatchGetWaitEstimates goa.Endpoint
	CreateWaitlistEntry   goa.Endpoint
	GetWaitlistEntry      goa.Endpoint
	DeleteWaitlistEntry   goa.Endpoint
}

// NewEndpoints wraps the methods of the "waitlist" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		BatchGetWaitEstimates: NewBatchGetWaitEstimatesEndpoint(s, a.BasicAuth),
		CreateWaitlistEntry:   NewCreateWaitlistEntryEndpoint(s, a.BasicAuth),
		GetWaitlistEntry:      NewGetWaitlistEntryEndpoint(s, a.BasicAuth),
		DeleteWaitlistEntry:   NewDeleteWaitlistEntryEndpoint(s, a.BasicAuth),
	}
}

// Use applies the given middleware to all the "waitlist" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.BatchGetWaitEstimates = m(e.BatchGetWaitEstimates)
	e.CreateWaitlistEntry = m(e.CreateWaitlistEntry)
	e.GetWaitlistEntry = m(e.GetWaitlistEntry)
	e.DeleteWaitlistEntry = m(e.DeleteWaitlistEntry)
}

// NewBatchGetWaitEstimatesEndpoint returns an endpoint function that calls the
// method "batch_get_wait_estimates" of service "waitlist".
func NewBatchGetWaitEstimatesEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BatchGetWaitEstimatesPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.BatchGetWaitEstimates(ctx, p)
	}
}

// NewCreateWaitlistEntryEndpoint returns an endpoint function that calls the
// method "create_waitlist_entry" of service "waitlist".
func NewCreateWaitlistEntryEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateWaitlistEntryPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.CreateWaitlistEntry(ctx, p)
	}
}

// NewGetWaitlistEntryEndpoint returns an endpoint function that calls the
// method "get_waitlist_entry" of service "waitlist".
func NewGetWaitlistEntryEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetWaitlistEntryPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetWaitlistEntry(ctx, p)
	}
}

// NewDeleteWaitlistEntryEndpoint returns an endpoint function that calls the
// method "delete_waitlist_entry" of service "waitlist".
func NewDeleteWaitlistEntryEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteWaitlistEntryPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteWaitlistEntry(ctx, p)
	}
}
