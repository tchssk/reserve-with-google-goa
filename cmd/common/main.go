package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	commonapi "github.com/tchssk/reserve-with-google-goa"
	booking "github.com/tchssk/reserve-with-google-goa/gen/booking"
	common "github.com/tchssk/reserve-with-google-goa/gen/common"
	order "github.com/tchssk/reserve-with-google-goa/gen/order"
	waitlist "github.com/tchssk/reserve-with-google-goa/gen/waitlist"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[commonapi] ", log.Ltime)
	}

	// Initialize the services.
	var (
		commonSvc   common.Service
		bookingSvc  booking.Service
		orderSvc    order.Service
		waitlistSvc waitlist.Service
	)
	{
		commonSvc = commonapi.NewCommon(logger)
		bookingSvc = commonapi.NewBooking(logger)
		orderSvc = commonapi.NewOrder(logger)
		waitlistSvc = commonapi.NewWaitlist(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		commonEndpoints   *common.Endpoints
		bookingEndpoints  *booking.Endpoints
		orderEndpoints    *order.Endpoints
		waitlistEndpoints *waitlist.Endpoints
	)
	{
		commonEndpoints = common.NewEndpoints(commonSvc)
		bookingEndpoints = booking.NewEndpoints(bookingSvc)
		orderEndpoints = order.NewEndpoints(orderSvc)
		waitlistEndpoints = waitlist.NewEndpoints(waitlistSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:80"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, commonEndpoints, bookingEndpoints, orderEndpoints, waitlistEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}