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
	Attribute("valid_start_time_sec", String)
	Attribute("valid_end_time_sec", String)
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
	Attribute("min_advance_cancellation_sec", String)
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
	Attribute("start_sec", String)
	Attribute("duration_sec", String)
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
