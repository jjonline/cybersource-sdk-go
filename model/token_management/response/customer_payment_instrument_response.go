package response

import "github.com/jjonline/cybersource-sdk-go/model/token_management"

type CustomerPaymentInstrumentResponse struct {
	ID                    string                                                   `json:"id"`
	Object                string                                                   `json:"object"`
	State                 string                                                   `json:"state"`
	Type                  string                                                   `json:"type"`
	Default               bool                                                     `json:"default"`
	BankAccount           *token_management.PaymentInstrumentBankAccount           `json:"bankAccount"`
	Card                  *token_management.PaymentInstrumentCard                  `json:"card"`
	BuyerInformation      *token_management.PaymentInstrumentBuyerInformation      `json:"buyerInformation"`
	BillTo                *token_management.PaymentInstrumentBillTo                `json:"billTo"`
	ProcessingInformation *token_management.PaymentInstrumentProcessingInformation `json:"processingInformation"`
	MerchantInformation   *token_management.PaymentInstrumentMerchantInformation   `json:"merchantInformation"`
	InstrumentIdentifier  *token_management.PaymentInstrumentIdentifier            `json:"instrumentIdentifier"`
	Metadata              *token_management.PaymentInstrumentMetadata              `json:"metadata"`
	// Reason and Message will be populated for 400 and 500 responses
	Reason  string    `json:"reason,omitempty"`
	Message string    `json:"message,omitempty"`
	Errors  []*Errors `json:"errors,omitempty"`
}
