package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jjonline/cybersource-sdk-go/cybersource"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
)

func main() {
	testProcessPayment()
}

func testProcessPayment() {
	client := cybersource.NewClient(cybersource.Params{
		MerchantKeyId:     "your_merchant_key_id",
		MerchantSecretKey: "your_merchant_secret_key",
		MerchantId:        "your_merchant_id",
		RequestHost:       "apitest.cybersource.com",
	})

	createPaymentReq := &request.CreatePaymentRequest{
		ClientReferenceInformation: &request.ClientReferenceInformation{
			Code: uuid.New().String(),
		},
		ProcessingInformation: &request.ProcessingInformation{
			CommerceIndicator: "internet",
		},
		OrderInformation: &request.OrderInformation{
			AmountDetails: &request.OrderInformationAmountDetails{
				TotalAmount: "60.00",
				Currency:    "USD",
			},
			BillTo: &request.OrderInformationBillTo{
				FirstName:          "John",
				LastName:           "Doe",
				Address1:           "201 S. Division St.",
				PostalCode:         "48104-2201",
				Locality:           "Ann Arbor",
				AdministrativeArea: "MI",
				Country:            "US",
				PhoneNumber:        "999999999",
				Email:              "test@cybs.com",
			},
		},
		PaymentInformation: &request.PaymentInformation{
			Card: &request.PaymentInformationCard{
				ExpirationYear:  "2031",
				Number:          "5555555555554444",
				SecurityCode:    "123",
				ExpirationMonth: "12",
				Type:            "002",
			},
		},
	}
	res, err := client.ProcessPayment(createPaymentReq)
	fmt.Println(res)
	fmt.Println(err)
}
