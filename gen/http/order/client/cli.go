// Code generated by goa v3.2.4, DO NOT EDIT.
//
// order HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package client

import (
	"encoding/json"
	"fmt"

	order "github.com/tchssk/reserve-with-google-goa/gen/order"
	goa "goa.design/goa/v3/pkg"
)

// BuildCheckOrderFulfillabilityPayload builds the payload for the order
// check_order_fulfillability endpoint from CLI flags.
func BuildCheckOrderFulfillabilityPayload(orderCheckOrderFulfillabilityBody string, orderCheckOrderFulfillabilityUsername string, orderCheckOrderFulfillabilityPassword string) (*order.CheckOrderFulfillabilityPayload, error) {
	var err error
	var body CheckOrderFulfillabilityRequestBody
	{
		err = json.Unmarshal([]byte(orderCheckOrderFulfillabilityBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cart_id\": \"Nisi odit alias et at magni.\",\n      \"item\": [\n         {\n            \"duration_sec\": 5218063573410113424,\n            \"intake_form_answers\": {\n               \"answer\": [\n                  {\n                     \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                     \"response\": [\n                        \"Quas harum ad qui excepturi sint.\",\n                        \"Dignissimos et.\",\n                        \"Quasi est placeat vel doloribus facilis.\",\n                        \"Ut est vel laborum.\"\n                     ]\n                  },\n                  {\n                     \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                     \"response\": [\n                        \"Quas harum ad qui excepturi sint.\",\n                        \"Dignissimos et.\",\n                        \"Quasi est placeat vel doloribus facilis.\",\n                        \"Ut est vel laborum.\"\n                     ]\n                  }\n               ]\n            },\n            \"price\": {\n               \"currency_code\": \"Sit quae.\",\n               \"price_micros\": 6250006249689697416,\n               \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n            },\n            \"service_id\": \"Quo at harum cupiditate officia.\",\n            \"start_sec\": 2726691727586715899,\n            \"status\": \"BOOKING_STATUS_UNSPECIFIED\",\n            \"tickets\": [\n               {\n                  \"count\": 1295690387,\n                  \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n               },\n               {\n                  \"count\": 1295690387,\n                  \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n               },\n               {\n                  \"count\": 1295690387,\n                  \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n               }\n            ],\n            \"warning_reason\": \"PRICE_INCREASE\"\n         },\n         {\n            \"duration_sec\": 5218063573410113424,\n            \"intake_form_answers\": {\n               \"answer\": [\n                  {\n                     \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                     \"response\": [\n                        \"Quas harum ad qui excepturi sint.\",\n                        \"Dignissimos et.\",\n                        \"Quasi est placeat vel doloribus facilis.\",\n                        \"Ut est vel laborum.\"\n                     ]\n                  },\n                  {\n                     \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                     \"response\": [\n                        \"Quas harum ad qui excepturi sint.\",\n                        \"Dignissimos et.\",\n                        \"Quasi est placeat vel doloribus facilis.\",\n                        \"Ut est vel laborum.\"\n                     ]\n                  }\n               ]\n            },\n            \"price\": {\n               \"currency_code\": \"Sit quae.\",\n               \"price_micros\": 6250006249689697416,\n               \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n            },\n            \"service_id\": \"Quo at harum cupiditate officia.\",\n            \"start_sec\": 2726691727586715899,\n            \"status\": \"BOOKING_STATUS_UNSPECIFIED\",\n            \"tickets\": [\n               {\n                  \"count\": 1295690387,\n                  \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n               },\n               {\n                  \"count\": 1295690387,\n                  \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n               },\n               {\n                  \"count\": 1295690387,\n                  \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n               }\n            ],\n            \"warning_reason\": \"PRICE_INCREASE\"\n         }\n      ],\n      \"merchant_id\": \"Voluptatibus omnis expedita autem ea.\"\n   }'")
		}
		if body.Item == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("item", "body"))
		}
		for _, e := range body.Item {
			if e != nil {
				if err2 := ValidateLineItemRequestBodyRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var username string
	{
		username = orderCheckOrderFulfillabilityUsername
	}
	var password string
	{
		password = orderCheckOrderFulfillabilityPassword
	}
	v := &order.CheckOrderFulfillabilityRequest{
		MerchantID: body.MerchantID,
		CartID:     body.CartID,
	}
	if body.Item != nil {
		v.Item = make([]*order.LineItem, len(body.Item))
		for i, val := range body.Item {
			v.Item[i] = marshalLineItemRequestBodyRequestBodyToOrderLineItem(val)
		}
	}
	res := &order.CheckOrderFulfillabilityPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildCreateOrderPayload builds the payload for the order create_order
// endpoint from CLI flags.
func BuildCreateOrderPayload(orderCreateOrderBody string, orderCreateOrderUsername string, orderCreateOrderPassword string) (*order.CreateOrderPayload, error) {
	var err error
	var body CreateOrderRequestBody
	{
		err = json.Unmarshal([]byte(orderCreateOrderBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"idempotency_token\": \"Exercitationem et voluptas velit voluptas.\",\n      \"order\": {\n         \"item\": [\n            {\n               \"duration_sec\": 5218063573410113424,\n               \"intake_form_answers\": {\n                  \"answer\": [\n                     {\n                        \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                        \"response\": [\n                           \"Quas harum ad qui excepturi sint.\",\n                           \"Dignissimos et.\",\n                           \"Quasi est placeat vel doloribus facilis.\",\n                           \"Ut est vel laborum.\"\n                        ]\n                     },\n                     {\n                        \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                        \"response\": [\n                           \"Quas harum ad qui excepturi sint.\",\n                           \"Dignissimos et.\",\n                           \"Quasi est placeat vel doloribus facilis.\",\n                           \"Ut est vel laborum.\"\n                        ]\n                     }\n                  ]\n               },\n               \"price\": {\n                  \"currency_code\": \"Sit quae.\",\n                  \"price_micros\": 6250006249689697416,\n                  \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n               },\n               \"service_id\": \"Quo at harum cupiditate officia.\",\n               \"start_sec\": 2726691727586715899,\n               \"status\": \"BOOKING_STATUS_UNSPECIFIED\",\n               \"tickets\": [\n                  {\n                     \"count\": 1295690387,\n                     \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n                  },\n                  {\n                     \"count\": 1295690387,\n                     \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n                  },\n                  {\n                     \"count\": 1295690387,\n                     \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n                  }\n               ],\n               \"warning_reason\": \"PRICE_INCREASE\"\n            },\n            {\n               \"duration_sec\": 5218063573410113424,\n               \"intake_form_answers\": {\n                  \"answer\": [\n                     {\n                        \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                        \"response\": [\n                           \"Quas harum ad qui excepturi sint.\",\n                           \"Dignissimos et.\",\n                           \"Quasi est placeat vel doloribus facilis.\",\n                           \"Ut est vel laborum.\"\n                        ]\n                     },\n                     {\n                        \"id\": \"Nesciunt eum sit voluptate vel unde ipsam.\",\n                        \"response\": [\n                           \"Quas harum ad qui excepturi sint.\",\n                           \"Dignissimos et.\",\n                           \"Quasi est placeat vel doloribus facilis.\",\n                           \"Ut est vel laborum.\"\n                        ]\n                     }\n                  ]\n               },\n               \"price\": {\n                  \"currency_code\": \"Sit quae.\",\n                  \"price_micros\": 6250006249689697416,\n                  \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n               },\n               \"service_id\": \"Quo at harum cupiditate officia.\",\n               \"start_sec\": 2726691727586715899,\n               \"status\": \"BOOKING_STATUS_UNSPECIFIED\",\n               \"tickets\": [\n                  {\n                     \"count\": 1295690387,\n                     \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n                  },\n                  {\n                     \"count\": 1295690387,\n                     \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n                  },\n                  {\n                     \"count\": 1295690387,\n                     \"ticket_id\": \"Corporis ex quis reprehenderit molestiae fugit.\"\n                  }\n               ],\n               \"warning_reason\": \"PRICE_INCREASE\"\n            }\n         ],\n         \"merchant_id\": \"Iste unde et deleniti rem illum officia.\",\n         \"order_id\": \"Necessitatibus velit consequuntur et.\",\n         \"payment_information\": {\n            \"deposit\": {\n               \"deposit\": {\n                  \"currency_code\": \"Sit quae.\",\n                  \"price_micros\": 6250006249689697416,\n                  \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n               },\n               \"deposit_type\": \"FIXED_RATE_DEFAULT\",\n               \"min_advance_cancellation_sec\": 2682111425714885120\n            },\n            \"fees\": {\n               \"currency_code\": \"Sit quae.\",\n               \"price_micros\": 6250006249689697416,\n               \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n            },\n            \"fees_and_taxes\": {\n               \"currency_code\": \"Sit quae.\",\n               \"price_micros\": 6250006249689697416,\n               \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n            },\n            \"fraud_signals\": \"Iusto amet pariatur.\",\n            \"md_merchant_data\": \"Saepe quia.\",\n            \"no_show_fee\": {\n               \"fee\": {\n                  \"currency_code\": \"Sit quae.\",\n                  \"price_micros\": 6250006249689697416,\n                  \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n               },\n               \"fee_type\": \"PER_PERSON\"\n            },\n            \"pa_response\": \"Pariatur quia et.\",\n            \"payment_option_id\": \"Facere tempore provident impedit veniam incidunt suscipit.\",\n            \"payment_processed_by\": \"PROCESSED_BY_GOOGLE\",\n            \"payment_transaction_id\": \"Tenetur exercitationem non.\",\n            \"prepayment_status\": \"PREPAYMENT_PROVIDED\",\n            \"price\": {\n               \"currency_code\": \"Sit quae.\",\n               \"price_micros\": 6250006249689697416,\n               \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n            },\n            \"tax_amount\": {\n               \"currency_code\": \"Sit quae.\",\n               \"price_micros\": 6250006249689697416,\n               \"pricing_option_tag\": \"Perferendis iusto ea quae est quam.\"\n            },\n            \"user_payment_option_id\": \"Sapiente saepe libero quaerat velit quam.\"\n         },\n         \"user_information\": {\n            \"address\": {\n               \"country\": \"Vel et vel.\",\n               \"locality\": \"Illum optio.\",\n               \"postal_code\": \"Porro molestiae quam dignissimos illum quia praesentium.\",\n               \"region\": \"Praesentium nemo quia molestias.\",\n               \"street_address\": \"Esse necessitatibus officia nihil at mollitia nobis.\"\n            },\n            \"email\": \"Rerum excepturi ea accusamus illo.\",\n            \"family_name\": \"Repudiandae enim voluptas quia sed tenetur possimus.\",\n            \"given_name\": \"Debitis hic id qui.\",\n            \"language_code\": \"Reiciendis expedita consectetur at nulla.\",\n            \"telephone\": \"In nostrum quae sint.\",\n            \"user_id\": \"Et voluptatibus eius soluta.\"\n         }\n      },\n      \"payment_processing_parameters\": {\n         \"payment_method_token\": \"Sunt quod quos architecto iure quia quibusdam.\",\n         \"payment_processor\": \"Laboriosam est delectus quo quo.\",\n         \"processor\": \"PAYMENT_PROCESSOR_UNSPECIFIED\",\n         \"tokenization_config\": {\n            \"allowed_auth_methods\": [\n               \"PAN_ONLY\",\n               \"AUTH_METHOD_UNSPECIFIED\",\n               \"AUTH_METHOD_UNSPECIFIED\"\n            ],\n            \"billing_information_format\": \"BILLING_INFORMATION_FORMAT_UNSPECIFIED\",\n            \"card_network_parameters\": [\n               {\n                  \"acquirer_bin\": \"Sit molestiae sed quasi.\",\n                  \"acquirer_merchant_id\": \"Sed eum officia.\",\n                  \"card_network\": \"DISCOVER\"\n               },\n               {\n                  \"acquirer_bin\": \"Sit molestiae sed quasi.\",\n                  \"acquirer_merchant_id\": \"Sed eum officia.\",\n                  \"card_network\": \"DISCOVER\"\n               },\n               {\n                  \"acquirer_bin\": \"Sit molestiae sed quasi.\",\n                  \"acquirer_merchant_id\": \"Sed eum officia.\",\n                  \"card_network\": \"DISCOVER\"\n               },\n               {\n                  \"acquirer_bin\": \"Sit molestiae sed quasi.\",\n                  \"acquirer_merchant_id\": \"Sed eum officia.\",\n                  \"card_network\": \"DISCOVER\"\n               }\n            ],\n            \"merchant_of_record_name\": \"Blanditiis molestias.\",\n            \"payment_country_code\": \"Qui eius expedita quo nihil.\",\n            \"tokenization_parameter\": {\n               \"Eaque aut.\": \"Excepturi quisquam necessitatibus dolorem.\",\n               \"Error perspiciatis.\": \"Inventore dolor ab omnis.\",\n               \"Sed exercitationem rem similique aut necessitatibus ullam.\": \"Eius sequi autem unde laudantium aspernatur.\"\n            }\n         },\n         \"unparsed_payment_method_token\": \"Possimus consequatur sequi asperiores quasi.\",\n         \"version\": \"Animi quidem veniam molestiae.\"\n      }\n   }'")
		}
		if body.Order == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("order", "body"))
		}
		if body.Order != nil {
			if err2 := ValidateOrderRequestBodyRequestBody(body.Order); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.PaymentProcessingParameters != nil {
			if err2 := ValidatePaymentProcessingParametersRequestBodyRequestBody(body.PaymentProcessingParameters); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var username string
	{
		username = orderCreateOrderUsername
	}
	var password string
	{
		password = orderCreateOrderPassword
	}
	v := &order.CreateOrderRequest{
		IdempotencyToken: body.IdempotencyToken,
	}
	if body.Order != nil {
		v.Order = marshalOrderRequestBodyRequestBodyToOrderOrder(body.Order)
	}
	if body.PaymentProcessingParameters != nil {
		v.PaymentProcessingParameters = marshalPaymentProcessingParametersRequestBodyRequestBodyToOrderPaymentProcessingParameters(body.PaymentProcessingParameters)
	}
	res := &order.CreateOrderPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildListOrdersPayload builds the payload for the order list_orders endpoint
// from CLI flags.
func BuildListOrdersPayload(orderListOrdersBody string, orderListOrdersUsername string, orderListOrdersPassword string) (*order.ListOrdersPayload, error) {
	var err error
	var body ListOrdersRequestBody
	{
		err = json.Unmarshal([]byte(orderListOrdersBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"order_ids\": {\n         \"order_id\": [\n            \"Voluptatem temporibus itaque error dolores odio.\",\n            \"Quis modi autem quidem quaerat.\",\n            \"Repellat qui quod nesciunt sint officiis.\",\n            \"Et voluptatem neque et id nisi.\"\n         ]\n      },\n      \"user_id\": \"Omnis qui similique quos tempore maiores.\"\n   }'")
		}
	}
	var username string
	{
		username = orderListOrdersUsername
	}
	var password string
	{
		password = orderListOrdersPassword
	}
	v := &order.ListOrdersRequest{
		UserID: body.UserID,
	}
	if body.OrderIds != nil {
		v.OrderIds = marshalOrderIdsRequestBodyRequestBodyToOrderOrderIds(body.OrderIds)
	}
	res := &order.ListOrdersPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}
