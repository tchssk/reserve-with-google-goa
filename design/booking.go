package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("booking", func() {
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
	Method("batch_availability_lookup", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", BatchAvailabilityLookupRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(BatchAvailabilityLookupResponse)
		HTTP(func() {
			POST("/v3/BatchAvailabilityLookup")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("check_availability", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", CheckAvailabilityRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(CheckAvailabilityResponse)
		HTTP(func() {
			POST("/v3/CheckAvailability")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("create_booking", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", CreateBookingRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(CreateBookingResponse)
		HTTP(func() {
			POST("/v3/CreateBooking")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("update_booking", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", UpdateBookingRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(UpdateBookingResponse)
		HTTP(func() {
			POST("/v3/UpdateBooking")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("get_booking_status", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", GetBookingStatusRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(GetBookingStatusResponse)
		HTTP(func() {
			POST("/v3/GetBookingStatus")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("list_bookings", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", ListBookingsRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(ListBookingsResponse)
		HTTP(func() {
			POST("/v3/ListBookings")
			Body("body")
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

// BatchAvailabilityLookup method

var BatchAvailabilityLookupRequest = Type("BatchAvailabilityLookupRequest", func() {
	Attribute("merchant_id", String)
	Attribute("slot_time", ArrayOf(SlotTime))
	Required(
		"merchant_id",
		"slot_time",
	)
})

var BatchAvailabilityLookupResponse = Type("BatchAvailabilityLookupResponse", func() {
	Attribute("slot_time_availability", ArrayOf(SlotTimeAvailability))
	Required("slot_time_availability")
})

var SlotTime = Type("SlotTime", func() {
	Attribute("service_id", String)
	Attribute("start_sec", String)
	Attribute("duration_sec", String)
	Attribute("availability_tag", String)
	Attribute("resource_ids", ResourceIds)
	Attribute("confirmation_mode", String, func() {
		Enum(ConfirmationMode...)
	})
	Required(
		"service_id",
		"start_sec",
	)
})

var SlotTimeAvailability = Type("SlotTimeAvailability", func() {
	Attribute("slot_time", SlotTime)
	Attribute("available", Boolean)
	Required(
		"slot_time",
		"available",
	)
})

// CheckAvailability method

var CheckAvailabilityRequest = Type("CheckAvailabilityRequest", func() {
	Attribute("slot", Slot)
	Required("slot")
})

const (
	DurationRequirementUnspecified       = "DURATION_REQUIREMENT_UNSPECIFIED"
	DurationRequirementDoNotShowDuration = "DO_NOT_SHOW_DURATION"
	DurationRequirementMustShowDuration  = "MUST_SHOW_DURATION"
)

var DurationRequirement = []interface{}{
	DurationRequirementUnspecified,
	DurationRequirementDoNotShowDuration,
	DurationRequirementMustShowDuration,
}

var CheckAvailabilityResponse = Type("CheckAvailabilityResponse", func() {
	Attribute("slot", Slot)
	Attribute("count_available", Int32)
	Attribute("last_online_cancellable_sec", String)
	Attribute("duration_requirement", String, func() {
		Enum(DurationRequirement...)
	})
	Attribute("availability_update", AvailabilityUpdate)
	Required(
		"slot",
		"count_available",
	)
})

var AvailabilityUpdate = Type("AvailabilityUpdate", func() {
	Attribute("slot_availability", ArrayOf(SlotAvailability))
	Required("slot_availability")
})

var SlotAvailability = Type("SlotAvailability", func() {
	Attribute("slot", Slot)
	Attribute("count_available", Int32)
	Required(
		"slot",
		"count_available",
	)
})

// GetBookingStatus method

var GetBookingStatusRequest = Type("GetBookingStatusRequest", func() {
	Attribute("booking_id", String)
	Required("booking_id")
})

var GetBookingStatusResponse = Type("GetBookingStatusResponse", func() {
	Attribute("booking_id", String)
	Attribute("booking_status", String, func() {
		Enum(BookingStatus...)
	})
	Attribute("prepayment_status", String, func() {
		Enum(PrepaymentStatus...)
	})
	Required(
		"booking_id",
		"booking_status",
		"prepayment_status",
	)
})

// CreateBooking method

var CreateBookingRequest = Type("CreateBookingRequest", func() {
	Attribute("slot", Slot)
	Attribute("lease_ref", LeaseReference)
	Attribute("user_information", UserInformation)
	Attribute("payment_information", PaymentInformation)
	Attribute("payment_processing_parameters", PaymentProcessingParameters)
	Attribute("idempotency_token", String)
	Attribute("additional_request", String)
	Attribute("offer_id", String)
	Attribute("deal_id", String)
	Required(
		"slot",
		"user_information",
		"idempotency_token",
	)
})

var CreateBookingResponse = Type("CreateBookingResponse", func() {
	Attribute("booking", Booking)
	Attribute("user_payment_option", UserPaymentOption)
	Attribute("booking_failure", BookingFailure)
	Required("booking")
})

// CreateLease method

var CreateLeaseRequest = Type("CreateLeaseRequest", func() {
	Attribute("lease", Lease)
	Required("lease")
})

var CreateLeaseResponse = Type("CreateLeaseResponse", func() {
	Attribute("lease", Lease)
	Attribute("booking_failure", BookingFailure)
	Required("lease")
})

var Lease = Type("Lease", func() {
	Attribute("lease_id", String)
	Attribute("slot", Slot)
	Attribute("user_reference", String)
	Attribute("lease_expiration_time_sec", String)
	Required(
		"lease_id",
		"slot",
		"user_reference",
		"lease_expiration_time_sec",
	)
})

var LeaseReference = Type("LeaseReference", func() {
	Attribute("lease_id", String)
	Required("lease_id")
})

// ListBookings method

var ListBookingsRequest = Type("ListBookingsRequest", func() {
	Attribute("user_id", String)
	Required("user_id")
})

var ListBookingsResponse = Type("ListBookingsResponse", func() {
	Attribute("bookings", ArrayOf(Booking))
	Required("bookings")
})

// UpdateBooking method

var UpdateBookingRequest = Type("UpdateBookingRequest", func() {
	Attribute("booking", Booking)
	Required("booking")
})

var UpdateBookingResponse = Type("UpdateBookingResponse", func() {
	Attribute("booking", Booking)
	Attribute("user_payment_option", UserPaymentOption)
	Attribute("booking_failure", BookingFailure)
	Required("booking")
})

// Booking specification

var Booking = Type("Booking", func() {
	Attribute("booking_id", String)
	Attribute("slot", Slot)
	Attribute("user_information", UserInformation)
	Attribute("status", String, func() {
		Enum(BookingStatus...)
	})
	Attribute("payment_information", PaymentInformation)
	Attribute("virtual_session_info", VirtualSessionInfo)
	Attribute("offer_info", OfferInfo)
	Required("booking_id")
})

// BookingStatus specification

const (
	BookingStatusUnspecified                 = "BOOKING_STATUS_UNSPECIFIED"
	BookingStatusConfirmed                   = "CONFIRMED"
	BookingStatusPendingMerchantConfirmation = "PENDING_MERCHANT_CONFIRMATION"
	BookingStatusCanceled                    = "CANCELED"
	BookingStatusNoShow                      = "NO_SHOW"
	BookingStatusNoShowPenalized             = "NO_SHOW_PENALIZED"
	BookingStatusFailed                      = "FAILED"
	BookingStatusDeclinedByMerchant          = "DECLINED_BY_MERCHANT"
)

var BookingStatus = []interface{}{
	BookingStatusUnspecified,
	BookingStatusConfirmed,
	BookingStatusPendingMerchantConfirmation,
	BookingStatusCanceled,
	BookingStatusNoShow,
	BookingStatusNoShowPenalized,
	BookingStatusFailed,
	BookingStatusDeclinedByMerchant,
}

// VirtualSessionInfo specification

var VirtualSessionInfo = Type("VirtualSessionInfo", func() {
	Attribute("session_url", String)
	Attribute("meeting_id", String)
	Attribute("password", String)
})

// BookingFailure specification

var BookingFailure = Type("BookingFailure", func() {
	const (
		CauseUnspecified                  = "CAUSE_UNSPECIFIED"
		CauseSlotUnavailable              = "SLOT_UNAVAILABLE"
		CauseSlotAlreadyBookedByUser      = "SLOT_ALREADY_BOOKED_BY_USER"
		CauseLeaseExpired                 = "LEASE_EXPIRED"
		CauseOutsideCancellationWindow    = "OUTSIDE_CANCELLATION_WINDOW"
		CausePaymentErrorCardTypeRejected = "PAYMENT_ERROR_CARD_TYPE_REJECTED"
		CausePaymentErrorCardDeclined     = "PAYMENT_ERROR_CARD_DECLINED"
		CausePaymentOptionNotValid        = "PAYMENT_OPTION_NOT_VALID"
		CausePaymentError                 = "PAYMENT_ERROR"
		CauseUserCannotUsePaymentOption   = "USER_CANNOT_USE_PAYMENT_OPTION"
		CauseBookingAlreadyCancelled      = "BOOKING_ALREADY_CANCELLED"
		CauseBookingNotCancellable        = "BOOKING_NOT_CANCELLABLE"
		CauseOverlappingReservation       = "OVERLAPPING_RESERVATION"
		CauseUserOverBookingLimit         = "USER_OVER_BOOKING_LIMIT"
		CauseOfferUnavailable             = "OFFER_UNAVAILABLE"
		CauseDealUnavailable              = "DEAL_UNAVAILABLE"
		CausePaymentRequires3ds1          = "PAYMENT_REQUIRES_3DS1"
	)
	var Cause = []interface{}{
		CauseUnspecified,
		CauseSlotUnavailable,
		CauseSlotAlreadyBookedByUser,
		CauseLeaseExpired,
		CauseOutsideCancellationWindow,
		CausePaymentErrorCardTypeRejected,
		CausePaymentErrorCardDeclined,
		CausePaymentOptionNotValid,
		CausePaymentError,
		CauseUserCannotUsePaymentOption,
		CauseBookingAlreadyCancelled,
		CauseBookingNotCancellable,
		CauseOverlappingReservation,
		CauseUserOverBookingLimit,
		CauseOfferUnavailable,
		CauseDealUnavailable,
		CausePaymentRequires3ds1,
	}
	Attribute("cause", String, func() {
		Enum(Cause...)
	})
	Attribute("rejected_card_type", String, func() {
		Enum(CreditCardType...)
	})
	Attribute("description", String)
	Attribute("payment_failure", PaymentFailureInformation)
	Required("cause")
})
