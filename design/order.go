package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("order", func() {
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
	Method("check_order_fulfillability", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", CheckOrderFulfillabilityRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(CheckOrderFulfillabilityResponse)
		HTTP(func() {
			POST("/v3/CheckOrderFulfillability")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("create_order", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", CreateOrderRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(CreateOrderResponse)
		HTTP(func() {
			POST("/v3/CreateOrder")
			Body("body")
			Response(StatusOK)
		})
	})
	Method("list_orders", func() {
		Payload(func() {
			Username("username")
			Password("password")
			Attribute("body", ListOrdersRequest)
			Required(
				"username",
				"password",
				"body",
			)
		})
		Result(ListOrdersResponse)
		HTTP(func() {
			POST("/v3/ListOrders")
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

// CheckOrderFulfillability method

var CheckOrderFulfillabilityRequest = Type("CheckOrderFulfillabilityRequest", func() {
	Attribute("merchant_id", String)
	Attribute("item", ArrayOf(LineItem))
	Attribute("cart_id", String)
	Required(
		"merchant_id",
		"item",
	)
})

var CheckOrderFulfillabilityResponse = Type("CheckOrderFulfillabilityResponse", func() {
	Attribute("fulfillability", OrderFulfillability)
	Attribute("fees_and_taxes", Price)
	Attribute("fees", Fees)
	Attribute("cart_expiration_sec", Int64)
})

var Fees = Type("Fees", func() {
	Attribute("per_ticket_fee", ArrayOf(SpecificPerTicketFee))
	Attribute("per_order_fee", ArrayOf(SpecificPerOrderFee))
	Required(
		"per_ticket_fee",
		"per_order_fee",
	)
})

var SpecificPerTicketFee = Type("SpecificPerTicketFee", func() {
	Attribute("ticket_id", String)
	Attribute("service_id", String)
	Attribute("fee_name", String)
	Attribute("fee_amount", Price)
	Required(
		"ticket_id",
		"service_id",
		"fee_name",
		"fee_amount",
	)
})

var SpecificPerOrderFee = Type("SpecificPerOrderFee", func() {
	Attribute("fee_name", String)
	Attribute("fee_amount", Price)
	Required(
		"fee_name",
		"fee_amount",
	)
})

// CreateOrder method

var CreateOrderRequest = Type("CreateOrderRequest", func() {
	Attribute("order", Order)
	Attribute("payment_processing_parameters", PaymentProcessingParameters)
	Attribute("idempotency_token", String)
	Required(
		"order",
		"idempotency_token",
	)
})

var CreateOrderResponse = Type("CreateOrderResponse", func() {
	Attribute("order", Order)
	Attribute("order_failure", OrderFailure)
})

// ListOrders method

var OrderIds = Type("OrderIds", func() {
	Attribute("order_id", ArrayOf(String))
})

var ListOrdersRequest = Type("ListOrdersRequest", func() {
	Attribute("user_id", String)
	Attribute("order_ids", OrderIds)
})

var ListOrdersResponse = Type("ListOrdersResponse", func() {
	Attribute("order", ArrayOf(Order))
	Required("order")
})

// Order specification

var Order = Type("Order", func() {
	Attribute("order_id", String)
	Attribute("user_information", UserInformation)
	Attribute("payment_information", PaymentInformation)
	Attribute("merchant_id", String)
	Attribute("item", ArrayOf(LineItem))
	Required(
		"user_information",
	)
})

var OrderedTickets = Type("OrderedTickets", func() {
	Attribute("ticket_id", String)
	Attribute("count", Int32)
	Required(
		"ticket_id",
		"count",
	)
})

const (
	WarningReasonUnspecifiedWarningReason = "UNSPECIFIED_WARNING_REASON"
	WarningReasonPriceIncrease            = "PRICE_INCREASE"
	WarningReasonPriceDecrease            = "PRICE_DECREASE"
)

var WarningReason = []interface{}{
	WarningReasonUnspecifiedWarningReason,
	WarningReasonPriceIncrease,
	WarningReasonPriceDecrease,
}

var LineItem = Type("LineItem", func() {
	Attribute("service_id", String)
	Attribute("start_sec", Int64)
	Attribute("duration_sec", Int64)
	Attribute("tickets", ArrayOf(OrderedTickets))
	Attribute("price", Price)
	Attribute("status", String, func() {
		Enum(BookingStatus...)
	})
	Attribute("intake_form_answers", IntakeFormAnswers)
	Attribute("warning_reason", String, func() {
		Enum(WarningReason...)
	})
	Required(
		"service_id",
		"start_sec",
		"duration_sec",
		"price",
	)
})

var IntakeFormAnswers = Type("IntakeFormAnswers", func() {
	Attribute("answer", ArrayOf(IntakeFormFieldAnswer))
})

var IntakeFormFieldAnswer = Type("IntakeFormFieldAnswer", func() {
	Attribute("id", String)
	Attribute("response", ArrayOf(String))
	Required(
		"id",
		"response",
	)
})

var OrderFailure = Type("OrderFailure", func() {
	const (
		CauseUnspecified                  = "CAUSE_UNSPECIFIED"
		CauseOrderUnfulfillable           = "ORDER_UNFULFILLABLE"
		CausePaymentErrorCardTypeRejected = "PAYMENT_ERROR_CARD_TYPE_REJECTED"
		CausePaymentErrorCardDeclined     = "PAYMENT_ERROR_CARD_DECLINED"
		CausePaymentError                 = "PAYMENT_ERROR"
		CauseIncorrectFeeTotal            = "INCORRECT_FEE_TOTAL"
		CausePaymentRequires3ds1          = "PAYMENT_REQUIRES_3DS1"
	)
	var Cause = []interface{}{
		CauseUnspecified,
		CauseOrderUnfulfillable,
		CausePaymentErrorCardTypeRejected,
		CausePaymentErrorCardDeclined,
		CausePaymentError,
		CauseIncorrectFeeTotal,
		CausePaymentRequires3ds1,
	}
	Attribute("cause", String, func() {
		Enum(Cause...)
	})
	Attribute("fulfillability", OrderFulfillability)
	Attribute("rejected_card_type", String, func() {
		Enum(CreditCardType...)
	})
	Attribute("description", String)
	Attribute("payment_failure", PaymentFailureInformation)
	Required("cause")
})

// OrderFulfillability specification

const (
	OrderFulfillabilityResultUnspecified                     = "ORDER_FULFILLABILITY_RESULT_UNSPECIFIED"
	OrderFulfillabilityResultCanFulfill                      = "CAN_FULFILL"
	OrderFulfillabilityResultUnfulfillableLineItem           = "UNFULFILLABLE_LINE_ITEM"
	OrderFulfillabilityResultUnfulfillableServiceCombination = "UNFULFILLABLE_SERVICE_COMBINATION"
	OrderFulfillabilityResultOrderUnfulfillableOtherReason   = "ORDER_UNFULFILLABLE_OTHER_REASON"
)

var OrderFulfillabilityResult = []interface{}{
	OrderFulfillabilityResultUnspecified,
	OrderFulfillabilityResultCanFulfill,
	OrderFulfillabilityResultUnfulfillableLineItem,
	OrderFulfillabilityResultUnfulfillableServiceCombination,
	OrderFulfillabilityResultOrderUnfulfillableOtherReason,
}

var OrderFulfillability = Type("OrderFulfillability", func() {
	Attribute("result", String, func() {
		Enum(OrderFulfillabilityResult...)
	})
	Attribute("item_fulfillability", ArrayOf(LineItemFulfillability))
	Attribute("unfulfillable_reason", String)
	Required(
		"result",
		"item_fulfillability",
	)
})

const (
	ItemFulfillabilityResultUnspecified                    = "ITEM_FULFILLABILITY_RESULT_UNSPECIFIED"
	ItemFulfillabilityResultCanFulfill                     = "CAN_FULFILL"
	ItemFulfillabilityResultSlotUnavailable                = "SLOT_UNAVAILABLE"
	ItemFulfillabilityResultChildTicketsWithoutAdult       = "CHILD_TICKETS_WITHOUT_ADULT"
	ItemFulfillabilityResultUnfulfillableTicketCombination = "UNFULFILLABLE_TICKET_COMBINATION"
	ItemFulfillabilityResultIncorrectPrice                 = "INCORRECT_PRICE"
	ItemFulfillabilityResultTicketConstraintViolated       = "TICKET_CONSTRAINT_VIOLATED"
	ItemFulfillabilityResultItemUnfulfillableOtherReason   = "ITEM_UNFULFILLABLE_OTHER_REASON"
)

var ItemFulfillabilityResult = []interface{}{
	ItemFulfillabilityResultUnspecified,
	ItemFulfillabilityResultCanFulfill,
	ItemFulfillabilityResultSlotUnavailable,
	ItemFulfillabilityResultChildTicketsWithoutAdult,
	ItemFulfillabilityResultUnfulfillableTicketCombination,
	ItemFulfillabilityResultIncorrectPrice,
	ItemFulfillabilityResultTicketConstraintViolated,
	ItemFulfillabilityResultItemUnfulfillableOtherReason,
}

var UpdatedAvailability = Type("UpdatedAvailability", func() {
	Attribute("spots_open", Int32)
	Required("spots_open")
})

var ViolatedTicketConstraint = Type("ViolatedTicketConstraint", func() {
	Attribute("min_ticket_count", Int32)
	Attribute("max_ticket_count", Int32)
	Attribute("ticket_id", String)
	Required(
		"min_ticket_count",
		"max_ticket_count",
		"ticket_id",
	)
})

var LineItemFulfillability = Type("LineItemFulfillability", func() {
	Attribute("item", LineItem)
	Attribute("result", String, func() {
		Enum(ItemFulfillabilityResult...)
	})
	Attribute("unfulfillable_reason", String)
	Attribute("availability", UpdatedAvailability)
	Attribute("ticket_type", ArrayOf(TicketType))
	Attribute("violated_ticket_constraint", ArrayOf(ViolatedTicketConstraint))
	Required(
		"item",
		"result",
	)
})

var TicketType = Type("TicketType", func() {
	Attribute("ticket_type_id", String)
	Attribute("short_description", String)
	Attribute("localized_short_description", Text)
	Attribute("price", Price)
	Attribute("per_ticket_fee", PerTicketFee)
	Attribute("option_description", String)
	Attribute("localized_option_description", Text)
	Required(
		"ticket_type_id",
		"short_description",
		"localized_short_description",
		"price",
		"option_description",
	)
})

var PerTicketFee = Type("PerTicketFee", func() {
	Attribute("service_charge", Price)
	Attribute("facility_fee", Price)
	Attribute("taxes", Price)
	Required(
		"service_charge",
		"facility_fee",
		"taxes",
	)
})

var Text = Type("Text", func() {
	Attribute("value", String)
	Attribute("localized_value", ArrayOf(LocalizedString))
})

var LocalizedString = Type("LocalizedString", func() {
	Attribute("locale", String)
	Attribute("value", String)
	Required(
		"locale",
		"value",
	)
})
