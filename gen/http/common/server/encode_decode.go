// Code generated by goa v3.2.4, DO NOT EDIT.
//
// common HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package server

import (
	"context"
	"net/http"

	common "github.com/tchssk/reserve-with-google-goa/gen/common"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeHealthCheckResponse returns an encoder for responses returned by the
// common health_check endpoint.
func EncodeHealthCheckResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*common.HealthCheckResponse)
		enc := encoder(ctx, w)
		body := NewHealthCheckResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeHealthCheckRequest returns a decoder for requests sent to the common
// health_check endpoint.
func DecodeHealthCheckRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		payload := NewHealthCheckPayload()
		user, pass, ok := r.BasicAuth()
		if !ok {
			return nil, goa.MissingFieldError("Authorization", "header")
		}
		payload.Username = user
		payload.Password = pass

		return payload, nil
	}
}

// EncodeHealthCheckError returns an encoder for errors returned by the
// health_check common endpoint.
func EncodeHealthCheckError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "conflict":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckConflictResponseBody(res)
			}
			w.Header().Set("goa-error", "conflict")
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "too_many_requests":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckTooManyRequestsResponseBody(res)
			}
			w.Header().Set("goa-error", "too_many_requests")
			w.WriteHeader(http.StatusTooManyRequests)
			return enc.Encode(body)
		case "client_closed_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckClientClosedRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "client_closed_request")
			w.WriteHeader(499)
			return enc.Encode(body)
		case "internal_server_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal_server_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_implemented":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckNotImplementedResponseBody(res)
			}
			w.Header().Set("goa-error", "not_implemented")
			w.WriteHeader(http.StatusNotImplemented)
			return enc.Encode(body)
		case "service_unavailable":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckServiceUnavailableResponseBody(res)
			}
			w.Header().Set("goa-error", "service_unavailable")
			w.WriteHeader(http.StatusServiceUnavailable)
			return enc.Encode(body)
		case "gateway_timeout":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewHealthCheckGatewayTimeoutResponseBody(res)
			}
			w.Header().Set("goa-error", "gateway_timeout")
			w.WriteHeader(http.StatusGatewayTimeout)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
