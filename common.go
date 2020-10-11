package commonapi

import (
	"context"
	"fmt"
	"log"

	common "github.com/tchssk/reserve-with-google-goa/gen/common"
	"goa.design/goa/v3/security"
)

// common service example implementation.
// The example methods log the requests and return zero values.
type commonsrvc struct {
	logger *log.Logger
}

// NewCommon returns the common service implementation.
func NewCommon(logger *log.Logger) common.Service {
	return &commonsrvc{logger}
}

// BasicAuth implements the authorization logic for service "common" for the
// "basic" security scheme.
func (s *commonsrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
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

// HealthCheck implements health_check.
func (s *commonsrvc) HealthCheck(ctx context.Context, p *common.HealthCheckPayload) (res *common.HealthCheckResponse, err error) {
	res = &common.HealthCheckResponse{}
	s.logger.Print("common.health_check")
	return
}
