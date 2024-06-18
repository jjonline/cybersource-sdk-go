package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) RetrieveTransaction(requestID string) (*response.TransactionDetails, int, error) {
	resource := "/tss/v2/transactions/" + requestID
	resp, err := c.doGet(resource)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var transactionDetails response.TransactionDetails
	err = json.Unmarshal(body, &transactionDetails)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &transactionDetails, resp.StatusCode, nil
}
