package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) RetrieveTransaction(requestID string) (*response.TransactionDetails, error) {
	resource := "/tss/v2/transactions/" + requestID
	resp, err := c.doGet(resource)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var transactionDetails response.TransactionDetails
	err = json.Unmarshal(body, &transactionDetails)
	if err != nil {
		return nil, err
	}
	return &transactionDetails, nil
}
