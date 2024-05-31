package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) RefreshPaymentStatus(requestID string, request *request.RefreshPaymentStatusRequest) (*response.RefreshPaymentStatusResponse, error) {
	resource := "/pts/v2/refresh-payment-status/" + requestID
	resp, err := c.doPost(resource, request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var refreshPaymentStatusResp response.RefreshPaymentStatusResponse
	err = json.Unmarshal(body, &refreshPaymentStatusResp)
	if err != nil {
		return nil, err
	}
	return &refreshPaymentStatusResp, nil
}
