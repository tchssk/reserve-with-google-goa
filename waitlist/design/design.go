package design

import (
	. "goa.design/goa/v3/dsl"
)

const (
	StatusClientClosedRequest = 499
)

var BasicAuth = BasicAuthSecurity("basic")

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
	Method("batch_get_wait_estimates", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", BatchGetWaitEstimatesRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(BatchGetWaitEstimatesResponse)
		HTTP(func() {
			POST("/v3/BatchGetWaitEstimates")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("create_waitlist_entry", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", CreateWaitlistEntryRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(CreateWaitlistEntryResponse)
		HTTP(func() {
			POST("/v3/CreateWaitlistEntry")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("get_waitlist_entry", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", GetWaitlistEntryRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(GetWaitlistEntryResponse)
		HTTP(func() {
			POST("/v3/GetWaitlistEntry")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("delete_waitlist_entry", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", DeleteWaitlistEntryRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		HTTP(func() {
			POST("/v3/DeleteWaitlistEntry")
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

// BatchGetWaitEstimates method

var BatchGetWaitEstimatesRequest = Type("BatchGetWaitEstimatesRequest", func() {
	Attribute("merchant_id", String)
	Attribute("service_id", String)
	Attribute("party_size", ArrayOf(Int32))
	Required(
		"merchant_id",
		"service_id",
		"party_size",
	)
})

var BatchGetWaitEstimatesResponse = Type("BatchGetWaitEstimatesResponse", func() {
	Attribute("waitlist_status", String, func() {
		Enum(WaitlistStatus...)
	})
	Attribute("wait_estimate", ArrayOf(WaitEstimate))
	Required(
		"waitlist_status",
	)
})

const (
	WaitlistStatusUnspecified  = "WAITLIST_STATUS_UNSPECIFIED"
	WaitlistStatusOpen         = "OPEN"
	WaitlistStatusClosedNoWait = "CLOSED_NO_WAIT"
	WaitlistStatusClosedOther  = "CLOSED_OTHER"
)

var WaitlistStatus = []interface{}{
	WaitlistStatusUnspecified,
	WaitlistStatusOpen,
	WaitlistStatusClosedNoWait,
	WaitlistStatusClosedOther,
}

var EstimatedSeatTimeRange = Type("EstimatedSeatTimeRange", func() {
	Attribute("start_seconds", Int64)
	Attribute("end_seconds", Int64)
	Required(
		"start_seconds",
		"end_seconds",
	)
})

var WaitLength = Type("WaitLength", func() {
	Attribute("parties_ahead_count", Int32)
	Attribute("estimated_seat_time_range", EstimatedSeatTimeRange)
	Required(
		"parties_ahead_count",
		"estimated_seat_time_range",
	)
})

const (
	WaitlistConfirmationModeUnspecified  = "WAITLIST_CONFIRMATION_MODE_UNSPECIFIED"
	WaitlistConfirmationModeSynchronous  = "WAITLIST_CONFIRMATION_MODE_SYNCHRONOUS"
	WaitlistConfirmationModeAsynchronous = "WAITLIST_CONFIRMATION_MODE_ASYNCHRONOUS"
)

var WaitlistConfirmationMode = []interface{}{
	WaitlistConfirmationModeUnspecified,
	WaitlistConfirmationModeSynchronous,
	WaitlistConfirmationModeAsynchronous,
}

var WaitEstimate = Type("WaitEstimate", func() {
	Attribute("party_size", Int32)
	Attribute("wait_length", WaitLength)
	Attribute("waitlist_confirmation_mode", String, func() {
		Enum(WaitlistConfirmationMode...)
	})
	Required(
		"party_size",
		"waitlist_confirmation_mode",
	)
})

// CreateWaitlistEntry method

var CreateWaitlistEntryRequest = Type("CreateWaitlistEntryRequest", func() {
	Attribute("merchant_id", String)
	Attribute("service_id", String)
	Attribute("party_size", Int32)
	Attribute("user_information", UserInformation)
	Attribute("additional_request", String)
	Attribute("idempotency_token", String)
	Required(
		"merchant_id",
		"service_id",
		"party_size",
		"user_information",
		"idempotency_token",
	)
})

var CreateWaitlistEntryResponse = Type("CreateWaitlistEntryResponse", func() {
	Attribute("waitlist_entry_id", String)
	Attribute("waitlist_business_logic_failure", WaitlistBusinessLogicFailure)
})

// GetWaitlistEntry method

var GetWaitlistEntryRequest = Type("GetWaitlistEntryRequest", func() {
	Attribute("waitlist_entry_id", String)
	Required("waitlist_entry_id")
})

var GetWaitlistEntryResponse = Type("GetWaitlistEntryResponse", func() {
	Attribute("waitlist_entry", WaitlistEntry)
	Required("waitlist_entry")
})

// DeleteWaitlistEntry method

var DeleteWaitlistEntryRequest = Type("DeleteWaitlistEntryRequest", func() {
	Attribute("waitlist_entry_id", String)
	Required("waitlist_entry_id")
})

// WaitlistEntry specification

const (
	WaitlistEntryStateUnspecified                 = "WAITLIST_ENTRY_STATE_UNSPECIFIED"
	WaitlistEntryStateWaiting                     = "WAITING"
	WaitlistEntryStatePendingMerchantConfirmation = "PENDING_MERCHANT_CONFIRMATION"
	WaitlistEntryStateCanceled                    = "CANCELED"
	WaitlistEntryStateDeclinedByMerchant          = "DECLINED_BY_MERCHANT"
	WaitlistEntryStateServiceReady                = "SERVICE_READY"
	WaitlistEntryStateCheckedIn                   = "CHECKED_IN"
	WaitlistEntryStateSeated                      = "SEATED"
	WaitlistEntryStateNoShow                      = "NO_SHOW"
)

var WaitlistEntryState = []interface{}{
	WaitlistEntryStateUnspecified,
	WaitlistEntryStateWaiting,
	WaitlistEntryStatePendingMerchantConfirmation,
	WaitlistEntryStateCanceled,
	WaitlistEntryStateDeclinedByMerchant,
	WaitlistEntryStateServiceReady,
	WaitlistEntryStateCheckedIn,
	WaitlistEntryStateSeated,
	WaitlistEntryStateNoShow,
}

var WaitlistEntryStateTimes = Type("WaitlistEntryStateTimes", func() {
	Attribute("created_time_seconds", Int64)
	Attribute("canceled_time_seconds", Int64)
	Attribute("service_readied_time_seconds", Int64)
	Attribute("checked_in_time_seconds", Int64)
	Attribute("seated_time_seconds", Int64)
	Attribute("marked_no_show_time_seconds", Int64)
	Required("created_time_seconds")
})

var WaitlistEntry = Type("WaitlistEntry", func() {
	Attribute("waitlist_entry_state", String, func() {
		Enum(WaitlistEntryState...)
	})
	Attribute("waitlist_entry_state_times", WaitlistEntryStateTimes)
	Attribute("wait_estimate", WaitEstimate)
	Required(
		"waitlist_entry_state",
		"waitlist_entry_state_times",
		"wait_estimate",
	)
})

// WaitlistBusinessLogicFailure specification

const (
	CauseUnspecified                   = "CAUSE_UNSPECIFIED"
	CauseExistingWaitlistEntry         = "EXISTING_WAITLIST_ENTRY"
	CauseBelowMinPartySize             = "BELOW_MIN_PARTY_SIZE"
	CauseAboveMaxPartySize             = "ABOVE_MAX_PARTY_SIZE"
	CauseMerchantClosed                = "MERCHANT_CLOSED"
	CauseNoWait                        = "NO_WAIT"
	CauseWaitlistFull                  = "WAITLIST_FULL"
	CausePhoneNumberCountryUnsupported = "PHONE_NUMBER_COUNTRY_UNSUPPORTED"
	CauseWaitlistClosed                = "WAITLIST_CLOSED"
)

var Cause = []interface{}{
	CauseUnspecified,
	CauseExistingWaitlistEntry,
	CauseBelowMinPartySize,
	CauseAboveMaxPartySize,
	CauseMerchantClosed,
	CauseNoWait,
	CauseWaitlistFull,
	CausePhoneNumberCountryUnsupported,
	CauseWaitlistClosed,
}

var WaitlistBusinessLogicFailure = Type("WaitlistBusinessLogicFailure", func() {
	Attribute("cause", String, func() {
		Enum(Cause...)
	})
	Attribute("description", String)
	Required("cause")
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
