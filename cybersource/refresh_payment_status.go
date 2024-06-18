package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) RefreshPaymentStatus(requestID string, request *request.RefreshPaymentStatusRequest) (*response.RefreshPaymentStatusResponse, int, error) {
	resource := "/pts/v2/refresh-payment-status/" + requestID
	resp, err := c.doPost(resource, request)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var refreshPaymentStatusResp response.RefreshPaymentStatusResponse
	err = json.Unmarshal(body, &refreshPaymentStatusResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &refreshPaymentStatusResp, resp.StatusCode, nil
}
