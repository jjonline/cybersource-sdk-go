package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/request"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/response"
	"io"
)

func (c *Client) CreateCustomerPaymentInstrument(customerID string, req *request.CustomerPaymentInstrumentRequest) (*response.CustomerPaymentInstrumentResponse, int, error) {
	resource := "/tms/v2/customers/" + customerID + "/payment-instruments"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	var paymentInstrumentResp response.CustomerPaymentInstrumentResponse
	err = json.Unmarshal(body, &paymentInstrumentResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &paymentInstrumentResp, resp.StatusCode, nil
}
