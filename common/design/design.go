package design

import (
	. "goa.design/goa/v3/dsl"
)

const (
	StatusClientClosedRequest = 499
)

var BasicAuth = BasicAuthSecurity("basic")

var _ = Service("common", func() {
	Security(BasicAuth)
	Error("bad_request")
	Error("unauthorized")
	Error("forbidden")
	Error("not_found")
	Error("conflict")
	Error("too_many_requests")
	Error("client_closed_request")
	Error("internal_server_error")
	Error("not_implemented")
	Error("service_unavailable")
	Error("gateway_timeout")
	Method("health_check", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Required(
				"username",
				"password",
			)
		})
		Result(HealthCheckResponse)
		HTTP(func() {
			GET("/v3/HealthCheck")
			Response(StatusOK)
		})
	})
	HTTP(func() {
		Response(StatusBadRequest, "bad_request")
		Response(StatusUnauthorized, "unauthorized")
		Response(StatusForbidden, "forbidden")
		Response(StatusNotFound, "not_found")
		Response(StatusConflict, "conflict")
		Response(StatusTooManyRequests, "too_many_requests")
		Response(StatusClientClosedRequest, "client_closed_request")
		Response(StatusInternalServerError, "internal_server_error")
		Response(StatusNotImplemented, "not_implemented")
		Response(StatusServiceUnavailable, "service_unavailable")
		Response(StatusGatewayTimeout, "gateway_timeout")
	})
})

const (
	ServingStatusUnknown        = "UNKNOWN"
	ServingStatusServing        = "SERVING"
	ServingStatusNotServing     = "NOT_SERVING"
	ServingStatusServiceUnknown = "SERVICE_UNKNOWN"
)

var ServingStatus = []interface{}{
	ServingStatusUnknown,
	ServingStatusServing,
	ServingStatusNotServing,
	ServingStatusServiceUnknown,
}

var HealthCheckResponse = Type("HealthCheckResponse", func() {
	Attribute("status", String, func() {
		Enum(ServingStatus...)
	})
	Required("status")
})
