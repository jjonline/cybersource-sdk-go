package main

import (
	"fmt"
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
			Code: "TC50171_3",
		},
		//ProcessingInformation: &request.ProcessingInformation{
		//	CommerceIndicator: "internet",
		//},
		OrderInformation: &request.OrderInformation{
			AmountDetails: &request.OrderInformationAmountDetails{
				TotalAmount: "102.21",
				Currency:    "USD",
			},
			BillTo: &request.OrderInformationBillTo{
				FirstName:          "John",
				LastName:           "Doe",
				Address1:           "1 Market St",
				PostalCode:         "94105",
				Locality:           "san francisco",
				AdministrativeArea: "CA",
				Country:            "US",
				PhoneNumber:        "4158880000",
				Email:              "test@cybs.com",
			},
		},
		PaymentInformation: &request.PaymentInformation{
			Card: &request.PaymentInformationCard{
				ExpirationYear:  "2031",
				Number:          "4111111111111111",
				ExpirationMonth: "12",
				// Type:            "002",
				// SecurityCode:    "123",
			},
		},
	}
	res, statusCode, err := client.ProcessPayment(createPaymentReq)
	fmt.Printf("%#v\n", res)
	fmt.Printf("%#v\n", statusCode)
	fmt.Println(err)
}
