package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) CapturePayment(requestID string, req *request.CapturePaymentRequest) (*response.CaptureResponse, int, error) {
	resource := "/pts/v2/payments/" + requestID + "/captures"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var capturePaymentResp response.CaptureResponse
	err = json.Unmarshal(body, &capturePaymentResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &capturePaymentResp, resp.StatusCode, nil
}
