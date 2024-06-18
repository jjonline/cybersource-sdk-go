package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/response"
	"io"
)

func (c *Client) RetrieveCustomerPaymentInstrument(customerID string, paymentInstrumentTokenID string) (*response.CustomerPaymentInstrumentResponse, int, error) {
	resource := "/tms/v2/customers/" + customerID + "/payment-instruments/" + paymentInstrumentTokenID
	resp, err := c.doGet(resource)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var instrumentResp response.CustomerPaymentInstrumentResponse
	err = json.Unmarshal(body, &instrumentResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &instrumentResp, resp.StatusCode, nil
}
