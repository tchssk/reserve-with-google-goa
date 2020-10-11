// Code generated by goa v3.2.4, DO NOT EDIT.
//
// booking HTTP server
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package server

import (
	"context"
	"net/http"

	booking "github.com/tchssk/reserve-with-google-goa/gen/booking"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the booking service endpoint HTTP handlers.
type Server struct {
	Mounts                  []*MountPoint
	BatchAvailabilityLookup http.Handler
	CheckAvailability       http.Handler
	CreateBooking           http.Handler
	UpdateBooking           http.Handler
	GetBookingStatus        http.Handler
	ListBookings            http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the booking service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *booking.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"BatchAvailabilityLookup", "POST", "/v3/BatchAvailabilityLookup"},
			{"CheckAvailability", "POST", "/v3/CheckAvailability"},
			{"CreateBooking", "POST", "/v3/CreateBooking"},
			{"UpdateBooking", "POST", "/v3/UpdateBooking"},
			{"GetBookingStatus", "POST", "/v3/GetBookingStatus"},
			{"ListBookings", "POST", "/v3/ListBookings"},
		},
		BatchAvailabilityLookup: NewBatchAvailabilityLookupHandler(e.BatchAvailabilityLookup, mux, decoder, encoder, errhandler, formatter),
		CheckAvailability:       NewCheckAvailabilityHandler(e.CheckAvailability, mux, decoder, encoder, errhandler, formatter),
		CreateBooking:           NewCreateBookingHandler(e.CreateBooking, mux, decoder, encoder, errhandler, formatter),
		UpdateBooking:           NewUpdateBookingHandler(e.UpdateBooking, mux, decoder, encoder, errhandler, formatter),
		GetBookingStatus:        NewGetBookingStatusHandler(e.GetBookingStatus, mux, decoder, encoder, errhandler, formatter),
		ListBookings:            NewListBookingsHandler(e.ListBookings, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "booking" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.BatchAvailabilityLookup = m(s.BatchAvailabilityLookup)
	s.CheckAvailability = m(s.CheckAvailability)
	s.CreateBooking = m(s.CreateBooking)
	s.UpdateBooking = m(s.UpdateBooking)
	s.GetBookingStatus = m(s.GetBookingStatus)
	s.ListBookings = m(s.ListBookings)
}

// Mount configures the mux to serve the booking endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountBatchAvailabilityLookupHandler(mux, h.BatchAvailabilityLookup)
	MountCheckAvailabilityHandler(mux, h.CheckAvailability)
	MountCreateBookingHandler(mux, h.CreateBooking)
	MountUpdateBookingHandler(mux, h.UpdateBooking)
	MountGetBookingStatusHandler(mux, h.GetBookingStatus)
	MountListBookingsHandler(mux, h.ListBookings)
}

// MountBatchAvailabilityLookupHandler configures the mux to serve the
// "booking" service "batch_availability_lookup" endpoint.
func MountBatchAvailabilityLookupHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v3/BatchAvailabilityLookup", f)
}

// NewBatchAvailabilityLookupHandler creates a HTTP handler which loads the
// HTTP request and calls the "booking" service "batch_availability_lookup"
// endpoint.
func NewBatchAvailabilityLookupHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeBatchAvailabilityLookupRequest(mux, decoder)
		encodeResponse = EncodeBatchAvailabilityLookupResponse(encoder)
		encodeError    = EncodeBatchAvailabilityLookupError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "batch_availability_lookup")
		ctx = context.WithValue(ctx, goa.ServiceKey, "booking")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCheckAvailabilityHandler configures the mux to serve the "booking"
// service "check_availability" endpoint.
func MountCheckAvailabilityHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v3/CheckAvailability", f)
}

// NewCheckAvailabilityHandler creates a HTTP handler which loads the HTTP
// request and calls the "booking" service "check_availability" endpoint.
func NewCheckAvailabilityHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCheckAvailabilityRequest(mux, decoder)
		encodeResponse = EncodeCheckAvailabilityResponse(encoder)
		encodeError    = EncodeCheckAvailabilityError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "check_availability")
		ctx = context.WithValue(ctx, goa.ServiceKey, "booking")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateBookingHandler configures the mux to serve the "booking" service
// "create_booking" endpoint.
func MountCreateBookingHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v3/CreateBooking", f)
}

// NewCreateBookingHandler creates a HTTP handler which loads the HTTP request
// and calls the "booking" service "create_booking" endpoint.
func NewCreateBookingHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateBookingRequest(mux, decoder)
		encodeResponse = EncodeCreateBookingResponse(encoder)
		encodeError    = EncodeCreateBookingError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_booking")
		ctx = context.WithValue(ctx, goa.ServiceKey, "booking")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateBookingHandler configures the mux to serve the "booking" service
// "update_booking" endpoint.
func MountUpdateBookingHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v3/UpdateBooking", f)
}

// NewUpdateBookingHandler creates a HTTP handler which loads the HTTP request
// and calls the "booking" service "update_booking" endpoint.
func NewUpdateBookingHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateBookingRequest(mux, decoder)
		encodeResponse = EncodeUpdateBookingResponse(encoder)
		encodeError    = EncodeUpdateBookingError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_booking")
		ctx = context.WithValue(ctx, goa.ServiceKey, "booking")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetBookingStatusHandler configures the mux to serve the "booking"
// service "get_booking_status" endpoint.
func MountGetBookingStatusHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v3/GetBookingStatus", f)
}

// NewGetBookingStatusHandler creates a HTTP handler which loads the HTTP
// request and calls the "booking" service "get_booking_status" endpoint.
func NewGetBookingStatusHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetBookingStatusRequest(mux, decoder)
		encodeResponse = EncodeGetBookingStatusResponse(encoder)
		encodeError    = EncodeGetBookingStatusError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_booking_status")
		ctx = context.WithValue(ctx, goa.ServiceKey, "booking")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListBookingsHandler configures the mux to serve the "booking" service
// "list_bookings" endpoint.
func MountListBookingsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v3/ListBookings", f)
}

// NewListBookingsHandler creates a HTTP handler which loads the HTTP request
// and calls the "booking" service "list_bookings" endpoint.
func NewListBookingsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListBookingsRequest(mux, decoder)
		encodeResponse = EncodeListBookingsResponse(encoder)
		encodeError    = EncodeListBookingsError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_bookings")
		ctx = context.WithValue(ctx, goa.ServiceKey, "booking")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
