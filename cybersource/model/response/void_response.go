package response

type VoidResponse struct {
	ID                         string                      `json:"id,omitempty"`
	SubmitTimeUTC              string                      `json:"submitTimeUtc,omitempty"`
	Status                     string                      `json:"status,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	VoidAmountDetails          *VoidAmountDetails          `json:"voidAmountDetails,omitempty"`
	ProcessorInformation       *ProcessorInformation       `json:"processorInformation,omitempty"`
}

type VoidAmountDetails struct {
	VoidAmount                string `json:"voidAmount,omitempty"`
	OriginalTransactionAmount string `json:"originalTransactionAmount,omitempty"`
	Currency                  string `json:"currency,omitempty"`
}
