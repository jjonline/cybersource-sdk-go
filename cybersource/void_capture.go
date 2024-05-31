package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) VoidCapture(requestID string, req *request.VoidRequest) (*response.VoidResponse, error) {
	resource := "/pts/v2/payments/" + requestID + "/voids"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var voidResp response.VoidResponse
	err = json.Unmarshal(body, &voidResp)
	if err != nil {
		return nil, err
	}
	return &voidResp, nil
}
