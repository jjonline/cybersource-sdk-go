package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) RefundPayment(requestID string, req *request.RefundRequest) (*response.RefundResponse, int, error) {
	resource := "/pts/v2/payments/" + requestID + "/refunds"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var refundResp response.RefundResponse
	err = json.Unmarshal(body, &refundResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &refundResp, resp.StatusCode, nil
}
