// Code generated by goa v3.2.4, DO NOT EDIT.
//
// booking HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tchssk/reserve-with-google-goa/design

package client

import (
	"encoding/json"
	"fmt"

	booking "github.com/tchssk/reserve-with-google-goa/gen/booking"
	goa "goa.design/goa/v3/pkg"
)

// BuildBatchAvailabilityLookupPayload builds the payload for the booking
// batch_availability_lookup endpoint from CLI flags.
func BuildBatchAvailabilityLookupPayload(bookingBatchAvailabilityLookupBody string, bookingBatchAvailabilityLookupUsername string, bookingBatchAvailabilityLookupPassword string) (*booking.BatchAvailabilityLookupPayload, error) {
	var err error
	var body BatchAvailabilityLookupRequestBody
	{
		err = json.Unmarshal([]byte(bookingBatchAvailabilityLookupBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"merchant_id\": \"Quia voluptates porro nulla praesentium.\",\n      \"slot_time\": [\n         {\n            \"availability_tag\": \"Doloribus aliquam rerum sed ea et.\",\n            \"confirmation_mode\": \"CONFIRMATION_MODE_ASYNCHRONOUS\",\n            \"duration_sec\": \"Sit aspernatur molestiae.\",\n            \"resource_ids\": {\n               \"party_size\": 32244509,\n               \"room_id\": \"Quidem omnis consectetur voluptatibus.\",\n               \"staff_id\": \"Provident qui omnis reiciendis qui.\"\n            },\n            \"service_id\": \"Omnis modi error veritatis quia.\",\n            \"start_sec\": \"Fugiat ut porro ullam qui.\"\n         },\n         {\n            \"availability_tag\": \"Doloribus aliquam rerum sed ea et.\",\n            \"confirmation_mode\": \"CONFIRMATION_MODE_ASYNCHRONOUS\",\n            \"duration_sec\": \"Sit aspernatur molestiae.\",\n            \"resource_ids\": {\n               \"party_size\": 32244509,\n               \"room_id\": \"Quidem omnis consectetur voluptatibus.\",\n               \"staff_id\": \"Provident qui omnis reiciendis qui.\"\n            },\n            \"service_id\": \"Omnis modi error veritatis quia.\",\n            \"start_sec\": \"Fugiat ut porro ullam qui.\"\n         },\n         {\n            \"availability_tag\": \"Doloribus aliquam rerum sed ea et.\",\n            \"confirmation_mode\": \"CONFIRMATION_MODE_ASYNCHRONOUS\",\n            \"duration_sec\": \"Sit aspernatur molestiae.\",\n            \"resource_ids\": {\n               \"party_size\": 32244509,\n               \"room_id\": \"Quidem omnis consectetur voluptatibus.\",\n               \"staff_id\": \"Provident qui omnis reiciendis qui.\"\n            },\n            \"service_id\": \"Omnis modi error veritatis quia.\",\n            \"start_sec\": \"Fugiat ut porro ullam qui.\"\n         }\n      ]\n   }'")
		}
		if body.SlotTime == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("slot_time", "body"))
		}
		for _, e := range body.SlotTime {
			if e != nil {
				if err2 := ValidateSlotTimeRequestBodyRequestBody(e); err2 != nil {
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
		username = bookingBatchAvailabilityLookupUsername
	}
	var password string
	{
		password = bookingBatchAvailabilityLookupPassword
	}
	v := &booking.BatchAvailabilityLookupRequest{
		MerchantID: body.MerchantID,
	}
	if body.SlotTime != nil {
		v.SlotTime = make([]*booking.SlotTime, len(body.SlotTime))
		for i, val := range body.SlotTime {
			v.SlotTime[i] = marshalSlotTimeRequestBodyRequestBodyToBookingSlotTime(val)
		}
	}
	res := &booking.BatchAvailabilityLookupPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildCheckAvailabilityPayload builds the payload for the booking
// check_availability endpoint from CLI flags.
func BuildCheckAvailabilityPayload(bookingCheckAvailabilityBody string, bookingCheckAvailabilityUsername string, bookingCheckAvailabilityPassword string) (*booking.CheckAvailabilityPayload, error) {
	var err error
	var body CheckAvailabilityRequestBody
	{
		err = json.Unmarshal([]byte(bookingCheckAvailabilityBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"slot\": {\n         \"availability_tag\": \"Aliquid reiciendis expedita consectetur at nulla.\",\n         \"confirmation_mode\": \"CONFIRMATION_MODE_SYNCHRONOUS\",\n         \"duration_sec\": \"Rerum rerum excepturi ea accusamus.\",\n         \"merchant_id\": \"Sequi esse.\",\n         \"resources\": {\n            \"party_size\": 32244509,\n            \"room_id\": \"Quidem omnis consectetur voluptatibus.\",\n            \"staff_id\": \"Provident qui omnis reiciendis qui.\"\n         },\n         \"service_id\": \"Officia nihil at mollitia.\",\n         \"start_sec\": \"Possimus in nostrum quae.\"\n      }\n   }'")
		}
		if body.Slot == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("slot", "body"))
		}
		if body.Slot != nil {
			if err2 := ValidateSlotRequestBodyRequestBody(body.Slot); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var username string
	{
		username = bookingCheckAvailabilityUsername
	}
	var password string
	{
		password = bookingCheckAvailabilityPassword
	}
	v := &booking.CheckAvailabilityRequest{}
	if body.Slot != nil {
		v.Slot = marshalSlotRequestBodyRequestBodyToBookingSlot(body.Slot)
	}
	res := &booking.CheckAvailabilityPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildCreateBookingPayload builds the payload for the booking create_booking
// endpoint from CLI flags.
func BuildCreateBookingPayload(bookingCreateBookingBody string, bookingCreateBookingUsername string, bookingCreateBookingPassword string) (*booking.CreateBookingPayload, error) {
	var err error
	var body CreateBookingRequestBody
	{
		err = json.Unmarshal([]byte(bookingCreateBookingBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"additional_request\": \"Quia asperiores quos rerum.\",\n      \"deal_id\": \"Nemo a quidem ut consequatur.\",\n      \"idempotency_token\": \"Natus recusandae delectus pariatur.\",\n      \"lease_ref\": {\n         \"lease_id\": \"Expedita eius sequi autem unde laudantium aspernatur.\"\n      },\n      \"offer_id\": \"Assumenda qui aut quia necessitatibus sit.\",\n      \"payment_information\": {\n         \"deposit\": {\n            \"deposit\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"deposit_type\": \"PER_PERSON\",\n            \"min_advance_cancellation_sec\": \"Facilis sint velit.\"\n         },\n         \"fees\": {\n            \"currency_code\": \"Dolore omnis est architecto.\",\n            \"price_micros\": 941384116037444211,\n            \"pricing_option_tag\": \"Est quis.\"\n         },\n         \"fees_and_taxes\": {\n            \"currency_code\": \"Dolore omnis est architecto.\",\n            \"price_micros\": 941384116037444211,\n            \"pricing_option_tag\": \"Est quis.\"\n         },\n         \"fraud_signals\": \"Et corrupti sed suscipit qui maiores maiores.\",\n         \"md_merchant_data\": \"Delectus reiciendis voluptatem odio quibusdam.\",\n         \"no_show_fee\": {\n            \"fee\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"fee_type\": \"PER_PERSON\"\n         },\n         \"pa_response\": \"Delectus maiores.\",\n         \"payment_option_id\": \"Nobis facilis vel.\",\n         \"payment_processed_by\": \"PROCESSED_BY_PARTNER\",\n         \"payment_transaction_id\": \"Adipisci quibusdam.\",\n         \"prepayment_status\": \"PREPAYMENT_CREDITED\",\n         \"price\": {\n            \"currency_code\": \"Dolore omnis est architecto.\",\n            \"price_micros\": 941384116037444211,\n            \"pricing_option_tag\": \"Est quis.\"\n         },\n         \"tax_amount\": {\n            \"currency_code\": \"Dolore omnis est architecto.\",\n            \"price_micros\": 941384116037444211,\n            \"pricing_option_tag\": \"Est quis.\"\n         },\n         \"user_payment_option_id\": \"Maxime officia voluptatem.\"\n      },\n      \"payment_processing_parameters\": {\n         \"payment_method_token\": \"Et a quia odio.\",\n         \"payment_processor\": \"Animi unde nobis aliquid rem.\",\n         \"processor\": \"PROCESSOR_BRAINTREE\",\n         \"tokenization_config\": {\n            \"allowed_auth_methods\": [\n               \"CRYPTOGRAM_3DS\",\n               \"AUTH_METHOD_UNSPECIFIED\"\n            ],\n            \"billing_information_format\": \"BILLING_INFORMATION_FORMAT_UNSPECIFIED\",\n            \"card_network_parameters\": [\n               {\n                  \"acquirer_bin\": \"Non eius commodi iure optio.\",\n                  \"acquirer_merchant_id\": \"Laboriosam expedita iusto qui sint minus rerum.\",\n                  \"card_network\": \"MASTERCARD\"\n               },\n               {\n                  \"acquirer_bin\": \"Non eius commodi iure optio.\",\n                  \"acquirer_merchant_id\": \"Laboriosam expedita iusto qui sint minus rerum.\",\n                  \"card_network\": \"MASTERCARD\"\n               },\n               {\n                  \"acquirer_bin\": \"Non eius commodi iure optio.\",\n                  \"acquirer_merchant_id\": \"Laboriosam expedita iusto qui sint minus rerum.\",\n                  \"card_network\": \"MASTERCARD\"\n               }\n            ],\n            \"merchant_of_record_name\": \"Voluptatem rerum ex.\",\n            \"payment_country_code\": \"Est et.\",\n            \"tokenization_parameter\": {\n               \"Aut quibusdam delectus assumenda.\": \"Ut deleniti minima.\",\n               \"Eaque qui a repudiandae.\": \"Facilis voluptatem dolorem.\",\n               \"Neque dolor.\": \"Enim architecto error odio a.\"\n            }\n         },\n         \"unparsed_payment_method_token\": \"Modi iure.\",\n         \"version\": \"Reprehenderit ipsum expedita.\"\n      },\n      \"slot\": {\n         \"availability_tag\": \"Aliquid reiciendis expedita consectetur at nulla.\",\n         \"confirmation_mode\": \"CONFIRMATION_MODE_SYNCHRONOUS\",\n         \"duration_sec\": \"Rerum rerum excepturi ea accusamus.\",\n         \"merchant_id\": \"Sequi esse.\",\n         \"resources\": {\n            \"party_size\": 32244509,\n            \"room_id\": \"Quidem omnis consectetur voluptatibus.\",\n            \"staff_id\": \"Provident qui omnis reiciendis qui.\"\n         },\n         \"service_id\": \"Officia nihil at mollitia.\",\n         \"start_sec\": \"Possimus in nostrum quae.\"\n      },\n      \"user_information\": {\n         \"address\": {\n            \"country\": \"Inventore dolor ab omnis.\",\n            \"locality\": \"Qui blanditiis molestias cum qui eius.\",\n            \"postal_code\": \"Sed quasi asperiores sed eum officia.\",\n            \"region\": \"Quo nihil quaerat et totam sit.\",\n            \"street_address\": \"Quia consectetur veritatis vel suscipit recusandae quasi.\"\n         },\n         \"email\": \"Officiis sunt praesentium molestias excepturi.\",\n         \"family_name\": \"Error perspiciatis.\",\n         \"given_name\": \"Excepturi quisquam necessitatibus dolorem.\",\n         \"language_code\": \"Voluptatem consectetur at repellat quo.\",\n         \"telephone\": \"Maiores laborum porro.\",\n         \"user_id\": \"Eaque aut.\"\n      }\n   }'")
		}
		if body.Slot == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("slot", "body"))
		}
		if body.UserInformation == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("user_information", "body"))
		}
		if body.Slot != nil {
			if err2 := ValidateSlotRequestBodyRequestBody(body.Slot); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.PaymentInformation != nil {
			if err2 := ValidatePaymentInformationRequestBodyRequestBody(body.PaymentInformation); err2 != nil {
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
		username = bookingCreateBookingUsername
	}
	var password string
	{
		password = bookingCreateBookingPassword
	}
	v := &booking.CreateBookingRequest{
		IdempotencyToken:  body.IdempotencyToken,
		AdditionalRequest: body.AdditionalRequest,
		OfferID:           body.OfferID,
		DealID:            body.DealID,
	}
	if body.Slot != nil {
		v.Slot = marshalSlotRequestBodyRequestBodyToBookingSlot(body.Slot)
	}
	if body.LeaseRef != nil {
		v.LeaseRef = marshalLeaseReferenceRequestBodyRequestBodyToBookingLeaseReference(body.LeaseRef)
	}
	if body.UserInformation != nil {
		v.UserInformation = marshalUserInformationRequestBodyRequestBodyToBookingUserInformation(body.UserInformation)
	}
	if body.PaymentInformation != nil {
		v.PaymentInformation = marshalPaymentInformationRequestBodyRequestBodyToBookingPaymentInformation(body.PaymentInformation)
	}
	if body.PaymentProcessingParameters != nil {
		v.PaymentProcessingParameters = marshalPaymentProcessingParametersRequestBodyRequestBodyToBookingPaymentProcessingParameters(body.PaymentProcessingParameters)
	}
	res := &booking.CreateBookingPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildUpdateBookingPayload builds the payload for the booking update_booking
// endpoint from CLI flags.
func BuildUpdateBookingPayload(bookingUpdateBookingBody string, bookingUpdateBookingUsername string, bookingUpdateBookingPassword string) (*booking.UpdateBookingPayload, error) {
	var err error
	var body UpdateBookingRequestBody
	{
		err = json.Unmarshal([]byte(bookingUpdateBookingBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"booking\": {\n         \"booking_id\": \"Debitis nulla tenetur ut ut.\",\n         \"offer_info\": {\n            \"offer_id\": \"Cum voluptatibus omnis expedita.\"\n         },\n         \"payment_information\": {\n            \"deposit\": {\n               \"deposit\": {\n                  \"currency_code\": \"Dolore omnis est architecto.\",\n                  \"price_micros\": 941384116037444211,\n                  \"pricing_option_tag\": \"Est quis.\"\n               },\n               \"deposit_type\": \"PER_PERSON\",\n               \"min_advance_cancellation_sec\": \"Facilis sint velit.\"\n            },\n            \"fees\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"fees_and_taxes\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"fraud_signals\": \"Et corrupti sed suscipit qui maiores maiores.\",\n            \"md_merchant_data\": \"Delectus reiciendis voluptatem odio quibusdam.\",\n            \"no_show_fee\": {\n               \"fee\": {\n                  \"currency_code\": \"Dolore omnis est architecto.\",\n                  \"price_micros\": 941384116037444211,\n                  \"pricing_option_tag\": \"Est quis.\"\n               },\n               \"fee_type\": \"PER_PERSON\"\n            },\n            \"pa_response\": \"Delectus maiores.\",\n            \"payment_option_id\": \"Nobis facilis vel.\",\n            \"payment_processed_by\": \"PROCESSED_BY_PARTNER\",\n            \"payment_transaction_id\": \"Adipisci quibusdam.\",\n            \"prepayment_status\": \"PREPAYMENT_CREDITED\",\n            \"price\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"tax_amount\": {\n               \"currency_code\": \"Dolore omnis est architecto.\",\n               \"price_micros\": 941384116037444211,\n               \"pricing_option_tag\": \"Est quis.\"\n            },\n            \"user_payment_option_id\": \"Maxime officia voluptatem.\"\n         },\n         \"slot\": {\n            \"availability_tag\": \"Aliquid reiciendis expedita consectetur at nulla.\",\n            \"confirmation_mode\": \"CONFIRMATION_MODE_SYNCHRONOUS\",\n            \"duration_sec\": \"Rerum rerum excepturi ea accusamus.\",\n            \"merchant_id\": \"Sequi esse.\",\n            \"resources\": {\n               \"party_size\": 32244509,\n               \"room_id\": \"Quidem omnis consectetur voluptatibus.\",\n               \"staff_id\": \"Provident qui omnis reiciendis qui.\"\n            },\n            \"service_id\": \"Officia nihil at mollitia.\",\n            \"start_sec\": \"Possimus in nostrum quae.\"\n         },\n         \"status\": \"CONFIRMED\",\n         \"user_information\": {\n            \"address\": {\n               \"country\": \"Inventore dolor ab omnis.\",\n               \"locality\": \"Qui blanditiis molestias cum qui eius.\",\n               \"postal_code\": \"Sed quasi asperiores sed eum officia.\",\n               \"region\": \"Quo nihil quaerat et totam sit.\",\n               \"street_address\": \"Quia consectetur veritatis vel suscipit recusandae quasi.\"\n            },\n            \"email\": \"Officiis sunt praesentium molestias excepturi.\",\n            \"family_name\": \"Error perspiciatis.\",\n            \"given_name\": \"Excepturi quisquam necessitatibus dolorem.\",\n            \"language_code\": \"Voluptatem consectetur at repellat quo.\",\n            \"telephone\": \"Maiores laborum porro.\",\n            \"user_id\": \"Eaque aut.\"\n         },\n         \"virtual_session_info\": {\n            \"meeting_id\": \"Voluptatem dolore repudiandae qui et ullam necessitatibus.\",\n            \"password\": \"Eum mollitia eveniet.\",\n            \"session_url\": \"Iste voluptatum.\"\n         }\n      }\n   }'")
		}
		if body.Booking == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("booking", "body"))
		}
		if body.Booking != nil {
			if err2 := ValidateBookingRequestBodyRequestBody(body.Booking); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var username string
	{
		username = bookingUpdateBookingUsername
	}
	var password string
	{
		password = bookingUpdateBookingPassword
	}
	v := &booking.UpdateBookingRequest{}
	if body.Booking != nil {
		v.Booking = marshalBookingRequestBodyRequestBodyToBookingBooking(body.Booking)
	}
	res := &booking.UpdateBookingPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildGetBookingStatusPayload builds the payload for the booking
// get_booking_status endpoint from CLI flags.
func BuildGetBookingStatusPayload(bookingGetBookingStatusBody string, bookingGetBookingStatusUsername string, bookingGetBookingStatusPassword string) (*booking.GetBookingStatusPayload, error) {
	var err error
	var body GetBookingStatusRequestBody
	{
		err = json.Unmarshal([]byte(bookingGetBookingStatusBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"booking_id\": \"Est placeat vel doloribus facilis necessitatibus ut.\"\n   }'")
		}
	}
	var username string
	{
		username = bookingGetBookingStatusUsername
	}
	var password string
	{
		password = bookingGetBookingStatusPassword
	}
	v := &booking.GetBookingStatusRequest{
		BookingID: body.BookingID,
	}
	res := &booking.GetBookingStatusPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}

// BuildListBookingsPayload builds the payload for the booking list_bookings
// endpoint from CLI flags.
func BuildListBookingsPayload(bookingListBookingsBody string, bookingListBookingsUsername string, bookingListBookingsPassword string) (*booking.ListBookingsPayload, error) {
	var err error
	var body ListBookingsRequestBody
	{
		err = json.Unmarshal([]byte(bookingListBookingsBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"user_id\": \"Tenetur at distinctio magnam mollitia sit.\"\n   }'")
		}
	}
	var username string
	{
		username = bookingListBookingsUsername
	}
	var password string
	{
		password = bookingListBookingsPassword
	}
	v := &booking.ListBookingsRequest{
		UserID: body.UserID,
	}
	res := &booking.ListBookingsPayload{
		Body: v,
	}
	res.Username = username
	res.Password = password

	return res, nil
}
