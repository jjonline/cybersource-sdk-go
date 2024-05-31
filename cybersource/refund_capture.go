package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) RefundCapture(requestID string, req *request.RefundRequest) (*response.RefundResponse, error) {
	resource := "/pts/v2/captures/" + requestID + "/refunds"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var refundResp response.RefundResponse
	err = json.Unmarshal(body, &refundResp)
	if err != nil {
		return nil, err
	}
	return &refundResp, nil
}
