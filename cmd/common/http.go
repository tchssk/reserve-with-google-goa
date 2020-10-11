package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	booking "github.com/tchssk/reserve-with-google-goa/gen/booking"
	common "github.com/tchssk/reserve-with-google-goa/gen/common"
	bookingsvr "github.com/tchssk/reserve-with-google-goa/gen/http/booking/server"
	commonsvr "github.com/tchssk/reserve-with-google-goa/gen/http/common/server"
	ordersvr "github.com/tchssk/reserve-with-google-goa/gen/http/order/server"
	waitlistsvr "github.com/tchssk/reserve-with-google-goa/gen/http/waitlist/server"
	order "github.com/tchssk/reserve-with-google-goa/gen/order"
	waitlist "github.com/tchssk/reserve-with-google-goa/gen/waitlist"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, commonEndpoints *common.Endpoints, bookingEndpoints *booking.Endpoints, orderEndpoints *order.Endpoints, waitlistEndpoints *waitlist.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		commonServer   *commonsvr.Server
		bookingServer  *bookingsvr.Server
		orderServer    *ordersvr.Server
		waitlistServer *waitlistsvr.Server
	)
	{
		eh := errorHandler(logger)
		commonServer = commonsvr.New(commonEndpoints, mux, dec, enc, eh, nil)
		bookingServer = bookingsvr.New(bookingEndpoints, mux, dec, enc, eh, nil)
		orderServer = ordersvr.New(orderEndpoints, mux, dec, enc, eh, nil)
		waitlistServer = waitlistsvr.New(waitlistEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				commonServer,
				bookingServer,
				orderServer,
				waitlistServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	commonsvr.Mount(mux, commonServer)
	bookingsvr.Mount(mux, bookingServer)
	ordersvr.Mount(mux, orderServer)
	waitlistsvr.Mount(mux, waitlistServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range commonServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range bookingServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range orderServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range waitlistServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
