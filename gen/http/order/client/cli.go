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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cart_id\": \"Molestias est minima eum porro sunt earum.\",\n      \"item\": [\n         {\n            \"duration_sec\": \"Eos atque illo voluptas et.\",\n            \"intake_form_answers\": {\n               \"answer\": [\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  },\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  }\n               ]\n            },\n            \"price\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"service_id\": \"Aut rem ut.\",\n            \"start_sec\": \"Sed neque omnis quia praesentium dolor.\",\n            \"status\": \"CONFIRMED\",\n            \"tickets\": [\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               },\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               }\n            ],\n            \"warning_reason\": \"UNSPECIFIED_WARNING_REASON\"\n         },\n         {\n            \"duration_sec\": \"Eos atque illo voluptas et.\",\n            \"intake_form_answers\": {\n               \"answer\": [\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  },\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  }\n               ]\n            },\n            \"price\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"service_id\": \"Aut rem ut.\",\n            \"start_sec\": \"Sed neque omnis quia praesentium dolor.\",\n            \"status\": \"CONFIRMED\",\n            \"tickets\": [\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               },\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               }\n            ],\n            \"warning_reason\": \"UNSPECIFIED_WARNING_REASON\"\n         },\n         {\n            \"duration_sec\": \"Eos atque illo voluptas et.\",\n            \"intake_form_answers\": {\n               \"answer\": [\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  },\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  }\n               ]\n            },\n            \"price\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"service_id\": \"Aut rem ut.\",\n            \"start_sec\": \"Sed neque omnis quia praesentium dolor.\",\n            \"status\": \"CONFIRMED\",\n            \"tickets\": [\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               },\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               }\n            ],\n            \"warning_reason\": \"UNSPECIFIED_WARNING_REASON\"\n         },\n         {\n            \"duration_sec\": \"Eos atque illo voluptas et.\",\n            \"intake_form_answers\": {\n               \"answer\": [\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  },\n                  {\n                     \"id\": \"Consequuntur incidunt maiores.\",\n                     \"response\": [\n                        \"Dicta ipsam.\",\n                        \"Et dolorem sit maiores.\",\n                        \"Exercitationem in voluptatem aspernatur iure.\"\n                     ]\n                  }\n               ]\n            },\n            \"price\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"service_id\": \"Aut rem ut.\",\n            \"start_sec\": \"Sed neque omnis quia praesentium dolor.\",\n            \"status\": \"CONFIRMED\",\n            \"tickets\": [\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               },\n               {\n                  \"count\": 2034419787,\n                  \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n               }\n            ],\n            \"warning_reason\": \"UNSPECIFIED_WARNING_REASON\"\n         }\n      ],\n      \"merchant_id\": \"Commodi vel placeat est debitis odit dicta.\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"idempotency_token\": \"Dolorem facere et.\",\n      \"order\": {\n         \"item\": [\n            {\n               \"duration_sec\": \"Eos atque illo voluptas et.\",\n               \"intake_form_answers\": {\n                  \"answer\": [\n                     {\n                        \"id\": \"Consequuntur incidunt maiores.\",\n                        \"response\": [\n                           \"Dicta ipsam.\",\n                           \"Et dolorem sit maiores.\",\n                           \"Exercitationem in voluptatem aspernatur iure.\"\n                        ]\n                     },\n                     {\n                        \"id\": \"Consequuntur incidunt maiores.\",\n                        \"response\": [\n                           \"Dicta ipsam.\",\n                           \"Et dolorem sit maiores.\",\n                           \"Exercitationem in voluptatem aspernatur iure.\"\n                        ]\n                     }\n                  ]\n               },\n               \"price\": {\n                  \"currency_code\": \"Dolore omnis est architecto.\",\n                  \"price_micros\": 941384116037444211,\n                  \"pricing_option_tag\": \"Est quis.\"\n               },\n               \"service_id\": \"Aut rem ut.\",\n               \"start_sec\": \"Sed neque omnis quia praesentium dolor.\",\n               \"status\": \"CONFIRMED\",\n               \"tickets\": [\n                  {\n                     \"count\": 2034419787,\n                     \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n                  },\n                  {\n                     \"count\": 2034419787,\n                     \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n                  }\n               ],\n               \"warning_reason\": \"UNSPECIFIED_WARNING_REASON\"\n            },\n            {\n               \"duration_sec\": \"Eos atque illo voluptas et.\",\n               \"intake_form_answers\": {\n                  \"answer\": [\n                     {\n                        \"id\": \"Consequuntur incidunt maiores.\",\n                        \"response\": [\n                           \"Dicta ipsam.\",\n                           \"Et dolorem sit maiores.\",\n                           \"Exercitationem in voluptatem aspernatur iure.\"\n                        ]\n                     },\n                     {\n                        \"id\": \"Consequuntur incidunt maiores.\",\n                        \"response\": [\n                           \"Dicta ipsam.\",\n                           \"Et dolorem sit maiores.\",\n                           \"Exercitationem in voluptatem aspernatur iure.\"\n                        ]\n                     }\n                  ]\n               },\n               \"price\": {\n                  \"currency_code\": \"Dolore omnis est architecto.\",\n                  \"price_micros\": 941384116037444211,\n                  \"pricing_option_tag\": \"Est quis.\"\n               },\n               \"service_id\": \"Aut rem ut.\",\n               \"start_sec\": \"Sed neque omnis quia praesentium dolor.\",\n               \"status\": \"CONFIRMED\",\n               \"tickets\": [\n                  {\n                     \"count\": 2034419787,\n                     \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n                  },\n                  {\n                     \"count\": 2034419787,\n                     \"ticket_id\": \"Voluptas ex voluptatem molestias veritatis laboriosam qui.\"\n                  }\n               ],\n               \"warning_reason\": \"UNSPECIFIED_WARNING_REASON\"\n            }\n         ],\n         \"merchant_id\": \"Impedit voluptatem suscipit incidunt dolore.\",\n         \"order_id\": \"Quo soluta adipisci error omnis.\",\n         \"payment_information\": {\n            \"deposit\": {\n               \"deposit\": {\n                  \"currency_code\": \"Dolore omnis est architecto.\",\n                  \"price_micros\": 941384116037444211,\n                  \"pricing_option_tag\": \"Est quis.\"\n               },\n               \"deposit_type\": \"PER_PERSON\",\n               \"min_advance_cancellation_sec\": \"Facilis sint velit.\"\n            },\n            \"fees\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"fees_and_taxes\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"fraud_signals\": \"Et corrupti sed suscipit qui maiores maiores.\",\n            \"md_merchant_data\": \"Delectus reiciendis voluptatem odio quibusdam.\",\n            \"no_show_fee\": {\n               \"fee\": {\n                  \"currency_code\": \"Dolore omnis est architecto.\",\n                  \"price_micros\": 941384116037444211,\n                  \"pricing_option_tag\": \"Est quis.\"\n               },\n               \"fee_type\": \"PER_PERSON\"\n            },\n            \"pa_response\": \"Delectus maiores.\",\n            \"payment_option_id\": \"Nobis facilis vel.\",\n            \"payment_processed_by\": \"PROCESSED_BY_PARTNER\",\n            \"payment_transaction_id\": \"Adipisci quibusdam.\",\n            \"prepayment_status\": \"PREPAYMENT_CREDITED\",\n            \"price\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"tax_amount\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"user_payment_option_id\": \"Maxime officia voluptatem.\"\n         },\n         \"user_information\": {\n            \"address\": {\n               \"country\": \"Inventore dolor ab omnis.\",\n               \"locality\": \"Qui blanditiis molestias cum qui eius.\",\n               \"postal_code\": \"Sed quasi asperiores sed eum officia.\",\n               \"region\": \"Quo nihil quaerat et totam sit.\",\n               \"street_address\": \"Quia consectetur veritatis vel suscipit recusandae quasi.\"\n            },\n            \"email\": \"Officiis sunt praesentium molestias excepturi.\",\n            \"family_name\": \"Error perspiciatis.\",\n            \"given_name\": \"Excepturi quisquam necessitatibus dolorem.\",\n            \"language_code\": \"Voluptatem consectetur at repellat quo.\",\n            \"telephone\": \"Maiores laborum porro.\",\n            \"user_id\": \"Eaque aut.\"\n         }\n      },\n      \"payment_processing_parameters\": {\n         \"payment_method_token\": \"Et a quia odio.\",\n         \"payment_processor\": \"Animi unde nobis aliquid rem.\",\n         \"processor\": \"PROCESSOR_BRAINTREE\",\n         \"tokenization_config\": {\n            \"allowed_auth_methods\": [\n               \"CRYPTOGRAM_3DS\",\n               \"AUTH_METHOD_UNSPECIFIED\"\n            ],\n            \"billing_information_format\": \"BILLING_INFORMATION_FORMAT_UNSPECIFIED\",\n            \"card_network_parameters\": [\n               {\n                  \"acquirer_bin\": \"Non eius commodi iure optio.\",\n                  \"acquirer_merchant_id\": \"Laboriosam expedita iusto qui sint minus rerum.\",\n                  \"card_network\": \"MASTERCARD\"\n               },\n               {\n                  \"acquirer_bin\": \"Non eius commodi iure optio.\",\n                  \"acquirer_merchant_id\": \"Laboriosam expedita iusto qui sint minus rerum.\",\n                  \"card_network\": \"MASTERCARD\"\n               },\n               {\n                  \"acquirer_bin\": \"Non eius commodi iure optio.\",\n                  \"acquirer_merchant_id\": \"Laboriosam expedita iusto qui sint minus rerum.\",\n                  \"card_network\": \"MASTERCARD\"\n               }\n            ],\n            \"merchant_of_record_name\": \"Voluptatem rerum ex.\",\n            \"payment_country_code\": \"Est et.\",\n            \"tokenization_parameter\": {\n               \"Aut quibusdam delectus assumenda.\": \"Ut deleniti minima.\",\n               \"Eaque qui a repudiandae.\": \"Facilis voluptatem dolorem.\",\n               \"Neque dolor.\": \"Enim architecto error odio a.\"\n            }\n         },\n         \"unparsed_payment_method_token\": \"Modi iure.\",\n         \"version\": \"Reprehenderit ipsum expedita.\"\n      }\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"order_ids\": {\n         \"order_id\": [\n            \"Earum ab est nam.\",\n            \"Qui odio molestias aut corrupti illum.\",\n            \"Unde molestias.\",\n            \"Qui voluptatem et sed natus unde.\"\n         ]\n      },\n      \"user_id\": \"Dolores quae.\"\n   }'")
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
