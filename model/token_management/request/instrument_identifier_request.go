package request

import "github.com/jjonline/cybersource-sdk-go/model/token_management"

type InstrumentIdentifierRequest struct {
	Card          *token_management.InstrumentIdentifierCard          `json:"card,omitempty"`
	BankAccount   *token_management.InstrumentIdentifierBankAccount   `json:"bankAccount,omitempty"`
	TokenizedCard *token_management.InstrumentIdentifierTokenizedCard `json:"tokenizedCard,omitempty"`
}
