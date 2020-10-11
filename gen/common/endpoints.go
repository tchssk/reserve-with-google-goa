// Code generated by goa v3.2.4, DO NOT EDIT.
//
// common endpoints
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package common

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "common" service endpoints.
type Endpoints struct {
	HealthCheck goa.Endpoint
}

// NewEndpoints wraps the methods of the "common" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		HealthCheck: NewHealthCheckEndpoint(s, a.BasicAuth),
	}
}

// Use applies the given middleware to all the "common" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.HealthCheck = m(e.HealthCheck)
}

// NewHealthCheckEndpoint returns an endpoint function that calls the method
// "health_check" of service "common".
func NewHealthCheckEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*HealthCheckPayload)
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
		return s.HealthCheck(ctx, p)
	}
}
