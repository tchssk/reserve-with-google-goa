package design

import (
	. "goa.design/goa/v3/dsl"
)

const (
	StatusClientClosedRequest = 499
)

var _ = Service("booking", func() {
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
		Payload(BatchAvailabilityLookupRequest)
		Result(BatchAvailabilityLookupResponse)
		HTTP(func() {
			POST("/v3/BatchAvailabilityLookup")
			Response(StatusOK)
		})
	})
	Method("check_availability", func() {
		Payload(CheckAvailabilityRequest)
		Result(CheckAvailabilityResponse)
		HTTP(func() {
			POST("/v3/CheckAvailability")
			Response(StatusOK)
		})
	})
	Method("create_booking", func() {
		Payload(CreateBookingRequest)
		Result(CreateBookingResponse)
		HTTP(func() {
			POST("/v3/CreateBooking")
			Response(StatusOK)
		})
	})
	Method("update_booking", func() {
		Payload(UpdateBookingRequest)
		Result(UpdateBookingResponse)
		HTTP(func() {
			POST("/v3/UpdateBooking")
			Response(StatusOK)
		})
	})
	Method("get_booking_status", func() {
		Payload(GetBookingStatusRequest)
		Result(GetBookingStatusResponse)
		HTTP(func() {
			POST("/v3/GetBookingStatus")
			Response(StatusOK)
		})
	})
	Method("list_bookings", func() {
		Payload(ListBookingsRequest)
		Result(ListBookingsResponse)
		HTTP(func() {
			POST("/v3/ListBookings")
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
	Attribute("start_sec", Int64)
	Attribute("duration_sec", Int64)
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
	Attribute("last_online_cancellable_sec", Int64)
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
	Attribute("slot", Slot, func() {
		Required(
			"merchant_id",
			"service_id",
			"start_sec",
			"duration_sec",
		)
	})
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
	Attribute("booking", Booking, func() {
		Required(
			"slot",
			"user_information",
			"status",
		)
	})
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
	Attribute("lease_expiration_time_sec", Int64)
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

var ThreeDS1Parameters = Type("ThreeDS1Parameters", func() {
	Attribute("acs_url", String)
	Attribute("pa_req", String)
	Attribute("transaction_id", String)
	Attribute("md_merchant_data", String)
	Required(
		"acs_url",
		"pa_req",
		"transaction_id",
		"md_merchant_data",
	)
})

var PaymentFailureInformation = Type("PaymentFailureInformation", func() {
	Attribute("threeds1_parameters", ThreeDS1Parameters)
	Required("threeds1_parameters")
})

var BookingFailure = Type("BookingFailure", func() {
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

const (
	CreditCardTypeUnspecified     = "CREDIT_CARD_TYPE_UNSPECIFIED"
	CreditCardTypeVisa            = "VISA"
	CreditCardTypeMastercard      = "MASTERCARD"
	CreditCardTypeAmericanExpress = "AMERICAN_EXPRESS"
	CreditCardTypeDiscover        = "DISCOVER"
	CreditCardTypeJCB             = "JCB"
)

var CreditCardType = []interface{}{
	CreditCardTypeUnspecified,
	CreditCardTypeVisa,
	CreditCardTypeMastercard,
	CreditCardTypeAmericanExpress,
	CreditCardTypeDiscover,
	CreditCardTypeJCB,
}

// Other specifications

var ResourceIds = Type("ResourceIds", func() {
	Attribute("staff_id", String)
	Attribute("room_id", String)
	Attribute("party_size", Int32)
})

var OfferInfo = Type("OfferInfo", func() {
	Attribute("offer_id", String)
	Required("offer_id")
})

// Payment specification

const (
	PaymentProcessorUnspecified = "PAYMENT_PROCESSOR_UNSPECIFIED"
	PaymentProcessorStripe      = "PROCESSOR_STRIPE"
	PaymentProcessorBraintree   = "PROCESSOR_BRAINTREE"
)

var PaymentProcessor = []interface{}{
	PaymentProcessorUnspecified,
	PaymentProcessorStripe,
	PaymentProcessorBraintree,
}

var PaymentProcessingParameters = Type("PaymentProcessingParameters", func() {
	Attribute("processor", String, func() {
		Enum(PaymentProcessor...)
	})
	Attribute("payment_method_token", String)
	Attribute("unparsed_payment_method_token", String)
	Attribute("version", String)
	Attribute("payment_processor", String)
	Attribute("tokenization_config", TokenizationConfig)
	Required(
		"processor",
		"payment_processor",
	)
})

const (
	PaymentOptionTypeUnspecified  = "PAYMENT_OPTION_TYPE_UNSPECIFIED"
	PaymentOptionTypeSingleUse    = "PAYMENT_OPTION_SINGLE_USE"
	PaymentOptionTypeMultiUse     = "PAYMENT_OPTION_MULTI_USE"
	PaymentOptionTypeUnlimitedUse = "PAYMENT_OPTION_UNLIMITED_USE"
)

var PaymentOptionType = []interface{}{
	PaymentOptionTypeUnspecified,
	PaymentOptionTypeSingleUse,
	PaymentOptionTypeMultiUse,
	PaymentOptionTypeUnlimitedUse,
}

var UserPaymentOption = Type("UserPaymentOption", func() {
	Attribute("user_payment_option_id", String)
	Attribute("valid_start_time_sec", Int64)
	Attribute("valid_end_time_sec", Int64)
	Attribute("type", String, func() {
		Enum(PaymentOptionType...)
	})
	Attribute("original_count", Int32)
	Attribute("current_count", Int32)
	Attribute("payment_option_id", String)
	Required(
		"user_payment_option_id",
		"type",
		"original_count",
		"current_count",
		"payment_option_id",
	)
})

const (
	PaymentProcessedByUnspecified = "PAYMENT_PROCESSED_BY_UNSPECIFIED"
	PaymentProcessedByGoogle      = "PROCESSED_BY_GOOGLE"
	PaymentProcessedByPartner     = "PROCESSED_BY_PARTNER"
)

var PaymentProcessedBy = []interface{}{
	PaymentProcessedByUnspecified,
	PaymentProcessedByGoogle,
	PaymentProcessedByPartner,
}

var PaymentInformation = Type("PaymentInformation", func() {
	Attribute("prepayment_status", String, func() {
		Enum(PrepaymentStatus...)
	})
	Attribute("payment_transaction_id", String)
	Attribute("price", Price)
	Attribute("tax_amount", Price)
	Attribute("fees", Price)
	Attribute("fees_and_taxes", Price)
	Attribute("deposit", Deposit)
	Attribute("no_show_fee", NoShowFee)
	Attribute("payment_processed_by", String, func() {
		Enum(PaymentProcessedBy...)
	})
	Attribute("payment_option_id", String)
	Attribute("user_payment_option_id", String)
	Attribute("fraud_signals", String)
	Attribute("pa_response", String)
	Attribute("md_merchant_data", String)
	Required(
		"prepayment_status",
		"price",
		"payment_processed_by",
	)
})

const (
	PrepaymentStatusUnspecified = "PREPAYMENT_STATUS_UNSPECIFIED"
	PrepaymentStatusProvided    = "PREPAYMENT_PROVIDED"
	PrepaymentStatusNotProvided = "PREPAYMENT_NOT_PROVIDED"
	PrepaymentStatusRefunded    = "PREPAYMENT_REFUNDED"
	PrepaymentStatusCredited    = "PREPAYMENT_CREDITED"
)

var PrepaymentStatus = []interface{}{
	PrepaymentStatusUnspecified,
	PrepaymentStatusProvided,
	PrepaymentStatusNotProvided,
	PrepaymentStatusRefunded,
	PrepaymentStatusCredited,
}

var Price = Type("Price", func() {
	Attribute("price_micros", Int64)
	Attribute("currency_code", String)
	Attribute("pricing_option_tag", String)
	Required(
		"price_micros",
		"currency_code",
	)
})

const (
	PriceTypeFixedRateDefault = "FIXED_RATE_DEFAULT"
	PriceTypePerPerson        = "PER_PERSON"
)

var PriceType = []interface{}{
	PriceTypeFixedRateDefault,
	PriceTypePerPerson,
}

var NoShowFee = Type("NoShowFee", func() {
	Attribute("fee", Price)
	Attribute("fee_type", String, func() {
		Enum(PriceType...)
	})
	Required(
		"fee",
		"fee_type",
	)
})

var Deposit = Type("Deposit", func() {
	Attribute("deposit", Price)
	Attribute("min_advance_cancellation_sec", Int64)
	Attribute("deposit_type", String, func() {
		Enum(PriceType...)
	})
	Required(
		"deposit",
		"min_advance_cancellation_sec",
		"deposit_type",
	)
})

// TokenizationConfig specification

const (
	BillingInformationFormatUnspecified = "BILLING_INFORMATION_FORMAT_UNSPECIFIED"
	BillingInformationFormatMin         = "MIN"
	BillingInformationFormatFull        = "FULL"
)

var BillingInformationFormat = []interface{}{
	BillingInformationFormatUnspecified,
	BillingInformationFormatMin,
	BillingInformationFormatFull,
}

var CardNetworkParameters = Type("CardNetworkParameters", func() {
	Attribute("card_network", String, func() {
		Enum(CreditCardType...)
	})
	Attribute("acquirer_bin", String)
	Attribute("acquirer_merchant_id", String)
	Required(
		"card_network",
		"acquirer_bin",
		"acquirer_merchant_id",
	)
})

const (
	AuthMethodUnspecified   = "AUTH_METHOD_UNSPECIFIED"
	AuthMethodPanOnly       = "PAN_ONLY"
	AuthMethodCryptogram3DS = "CRYPTOGRAM_3DS"
)

var AuthMethod = []interface{}{
	AuthMethodUnspecified,
	AuthMethodPanOnly,
	AuthMethodCryptogram3DS,
}

var TokenizationConfig = Type("TokenizationConfig", func() {
	Attribute("tokenization_parameter", MapOf(String, String))
	Attribute("billing_information_format", String, func() {
		Enum(BillingInformationFormat...)
	})
	Attribute("merchant_of_record_name", String)
	Attribute("payment_country_code", String)
	Attribute("card_network_parameters", ArrayOf(CardNetworkParameters))
	Attribute("allowed_auth_methods", ArrayOf(String, func() {
		Enum(AuthMethod...)
	}))
	Required(
		"tokenization_parameter",
		"billing_information_format",
		"merchant_of_record_name",
		"payment_country_code",
		"card_network_parameters",
		"allowed_auth_methods",
	)
})

var CreditCardRestrictions = Type("CreditCardRestrictions", func() {
	Attribute("credit_card_type", ArrayOf(String, func() {
		Enum(CreditCardType...)
	}))
})

// ConfirmationMode specification

const (
	ConfirmationModeUnspecified  = "CONFIRMATION_MODE_UNSPECIFIED"
	ConfirmationModeSynchronous  = "CONFIRMATION_MODE_SYNCHRONOUS"
	ConfirmationModeAsynchronous = "CONFIRMATION_MODE_ASYNCHRONOUS"
)

var ConfirmationMode = []interface{}{
	ConfirmationModeUnspecified,
	ConfirmationModeSynchronous,
	ConfirmationModeAsynchronous,
}

// Slot specification

var Slot = Type("Slot", func() {
	Attribute("merchant_id", String)
	Attribute("service_id", String)
	Attribute("start_sec", Int64)
	Attribute("duration_sec", Int64)
	Attribute("availability_tag", String)
	Attribute("resources", ResourceIds)
	Attribute("confirmation_mode", String, func() {
		Enum(ConfirmationMode...)
	})
})

// User specification

var UserInformation = Type("UserInformation", func() {
	Attribute("user_id", String)
	Attribute("given_name", String)
	Attribute("family_name", String)
	Attribute("address", PostalAddress)
	Attribute("telephone", String)
	Attribute("email", String)
	Attribute("language_code", String)
	Required(
		"user_id",
		"given_name",
		"family_name",
		"telephone",
		"email",
	)
})

var PostalAddress = Type("PostalAddress", func() {
	Attribute("country", String)
	Attribute("locality", String)
	Attribute("region", String)
	Attribute("postal_code", String)
	Attribute("street_address", String)
	Required(
		"country",
		"locality",
		"postal_code",
		"street_address",
	)
})
