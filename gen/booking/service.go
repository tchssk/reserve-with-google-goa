// Code generated by goa v3.2.4, DO NOT EDIT.
//
// booking service
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package booking

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the booking service interface.
type Service interface {
	// BatchAvailabilityLookup implements batch_availability_lookup.
	BatchAvailabilityLookup(context.Context, *BatchAvailabilityLookupPayload) (res *BatchAvailabilityLookupResponse, err error)
	// CheckAvailability implements check_availability.
	CheckAvailability(context.Context, *CheckAvailabilityPayload) (res *CheckAvailabilityResponse, err error)
	// CreateBooking implements create_booking.
	CreateBooking(context.Context, *CreateBookingPayload) (res *CreateBookingResponse, err error)
	// UpdateBooking implements update_booking.
	UpdateBooking(context.Context, *UpdateBookingPayload) (res *UpdateBookingResponse, err error)
	// GetBookingStatus implements get_booking_status.
	GetBookingStatus(context.Context, *GetBookingStatusPayload) (res *GetBookingStatusResponse, err error)
	// ListBookings implements list_bookings.
	ListBookings(context.Context, *ListBookingsPayload) (res *ListBookingsResponse, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "booking"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"batch_availability_lookup", "check_availability", "create_booking", "update_booking", "get_booking_status", "list_bookings"}

// BatchAvailabilityLookupPayload is the payload type of the booking service
// batch_availability_lookup method.
type BatchAvailabilityLookupPayload struct {
	Username string
	Password string
	Body     *BatchAvailabilityLookupRequest
}

// BatchAvailabilityLookupResponse is the result type of the booking service
// batch_availability_lookup method.
type BatchAvailabilityLookupResponse struct {
	SlotTimeAvailability []*SlotTimeAvailability
}

// CheckAvailabilityPayload is the payload type of the booking service
// check_availability method.
type CheckAvailabilityPayload struct {
	Username string
	Password string
	Body     *CheckAvailabilityRequest
}

// CheckAvailabilityResponse is the result type of the booking service
// check_availability method.
type CheckAvailabilityResponse struct {
	Slot                     *Slot
	CountAvailable           int32
	LastOnlineCancellableSec *string
	DurationRequirement      *string
	AvailabilityUpdate       *AvailabilityUpdate
}

// CreateBookingPayload is the payload type of the booking service
// create_booking method.
type CreateBookingPayload struct {
	Username string
	Password string
	Body     *CreateBookingRequest
}

// CreateBookingResponse is the result type of the booking service
// create_booking method.
type CreateBookingResponse struct {
	Booking           *Booking
	UserPaymentOption *UserPaymentOption
	BookingFailure    *BookingFailure
}

// UpdateBookingPayload is the payload type of the booking service
// update_booking method.
type UpdateBookingPayload struct {
	Username string
	Password string
	Body     *UpdateBookingRequest
}

// UpdateBookingResponse is the result type of the booking service
// update_booking method.
type UpdateBookingResponse struct {
	Booking           *Booking
	UserPaymentOption *UserPaymentOption
	BookingFailure    *BookingFailure
}

// GetBookingStatusPayload is the payload type of the booking service
// get_booking_status method.
type GetBookingStatusPayload struct {
	Username string
	Password string
	Body     *GetBookingStatusRequest
}

// GetBookingStatusResponse is the result type of the booking service
// get_booking_status method.
type GetBookingStatusResponse struct {
	BookingID        string
	BookingStatus    string
	PrepaymentStatus string
}

// ListBookingsPayload is the payload type of the booking service list_bookings
// method.
type ListBookingsPayload struct {
	Username string
	Password string
	Body     *ListBookingsRequest
}

// ListBookingsResponse is the result type of the booking service list_bookings
// method.
type ListBookingsResponse struct {
	Bookings []*Booking
}

type BatchAvailabilityLookupRequest struct {
	MerchantID string
	SlotTime   []*SlotTime
}

type SlotTime struct {
	ServiceID        string
	StartSec         string
	DurationSec      *string
	AvailabilityTag  *string
	ResourceIds      *ResourceIds
	ConfirmationMode *string
}

type ResourceIds struct {
	StaffID   *string
	RoomID    *string
	PartySize *int32
}

type SlotTimeAvailability struct {
	SlotTime  *SlotTime
	Available bool
}

type CheckAvailabilityRequest struct {
	Slot *Slot
}

type Slot struct {
	MerchantID       *string
	ServiceID        *string
	StartSec         *string
	DurationSec      *string
	AvailabilityTag  *string
	Resources        *ResourceIds
	ConfirmationMode *string
}

type AvailabilityUpdate struct {
	SlotAvailability []*SlotAvailability
}

type SlotAvailability struct {
	Slot           *Slot
	CountAvailable int32
}

type CreateBookingRequest struct {
	Slot                        *Slot
	LeaseRef                    *LeaseReference
	UserInformation             *UserInformation
	PaymentInformation          *PaymentInformation
	PaymentProcessingParameters *PaymentProcessingParameters
	IdempotencyToken            string
	AdditionalRequest           *string
	OfferID                     *string
	DealID                      *string
}

type LeaseReference struct {
	LeaseID string
}

type UserInformation struct {
	UserID       string
	GivenName    string
	FamilyName   string
	Address      *PostalAddress
	Telephone    string
	Email        string
	LanguageCode *string
}

type PostalAddress struct {
	Country       string
	Locality      string
	Region        *string
	PostalCode    string
	StreetAddress string
}

type PaymentInformation struct {
	PrepaymentStatus     string
	PaymentTransactionID *string
	Price                *Price
	TaxAmount            *Price
	Fees                 *Price
	FeesAndTaxes         *Price
	Deposit              *Deposit
	NoShowFee            *NoShowFee
	PaymentProcessedBy   string
	PaymentOptionID      *string
	UserPaymentOptionID  *string
	FraudSignals         *string
	PaResponse           *string
	MdMerchantData       *string
}

type Price struct {
	PriceMicros      int64
	CurrencyCode     string
	PricingOptionTag *string
}

type Deposit struct {
	Deposit                   *Price
	MinAdvanceCancellationSec string
	DepositType               string
}

type NoShowFee struct {
	Fee     *Price
	FeeType string
}

type PaymentProcessingParameters struct {
	Processor                  string
	PaymentMethodToken         *string
	UnparsedPaymentMethodToken *string
	Version                    *string
	PaymentProcessor           string
	TokenizationConfig         *TokenizationConfig
}

type TokenizationConfig struct {
	TokenizationParameter    map[string]string
	BillingInformationFormat string
	MerchantOfRecordName     string
	PaymentCountryCode       string
	CardNetworkParameters    []*CardNetworkParameters
	AllowedAuthMethods       []string
}

type CardNetworkParameters struct {
	CardNetwork        string
	AcquirerBin        string
	AcquirerMerchantID string
}

type Booking struct {
	BookingID          string
	Slot               *Slot
	UserInformation    *UserInformation
	Status             *string
	PaymentInformation *PaymentInformation
	VirtualSessionInfo *VirtualSessionInfo
	OfferInfo          *OfferInfo
}

type VirtualSessionInfo struct {
	SessionURL *string
	MeetingID  *string
	Password   *string
}

type OfferInfo struct {
	OfferID string
}

type UserPaymentOption struct {
	UserPaymentOptionID string
	ValidStartTimeSec   *string
	ValidEndTimeSec     *string
	Type                string
	OriginalCount       int32
	CurrentCount        int32
	PaymentOptionID     string
}

type BookingFailure struct {
	Cause            string
	RejectedCardType *string
	Description      *string
	PaymentFailure   *PaymentFailureInformation
}

type PaymentFailureInformation struct {
	Threeds1Parameters *ThreeDS1Parameters
}

type ThreeDS1Parameters struct {
	AcsURL         string
	PaReq          string
	TransactionID  string
	MdMerchantData string
}

type UpdateBookingRequest struct {
	Booking *Booking
}

type GetBookingStatusRequest struct {
	BookingID string
}

type ListBookingsRequest struct {
	UserID string
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeConflict builds a goa.ServiceError from an error.
func MakeConflict(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "conflict",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeTooManyRequests builds a goa.ServiceError from an error.
func MakeTooManyRequests(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "too_many_requests",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeClientClosedRequest builds a goa.ServiceError from an error.
func MakeClientClosedRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "client_closed_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternalServerError builds a goa.ServiceError from an error.
func MakeInternalServerError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal_server_error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotImplemented builds a goa.ServiceError from an error.
func MakeNotImplemented(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_implemented",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeServiceUnavailable builds a goa.ServiceError from an error.
func MakeServiceUnavailable(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "service_unavailable",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeGatewayTimeout builds a goa.ServiceError from an error.
func MakeGatewayTimeout(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "gateway_timeout",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
