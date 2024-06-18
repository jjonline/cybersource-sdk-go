package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) GetPayment(requestID string) (*response.GetPaymentResponse, int, error) {
	resource := "/pts/v2/payments/" + requestID
	resp, err := c.doGet(resource)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var getPaymentResp response.GetPaymentResponse
	err = json.Unmarshal(body, &getPaymentResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &getPaymentResp, resp.StatusCode, nil
}
