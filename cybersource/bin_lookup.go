package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) LookupBIN(req *request.BINLookupRequest) (*response.BINLookupResponse, error) {
	resource := "/bin/v1/binlookup"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var BINResp response.BINLookupResponse
	err = json.Unmarshal(body, &BINResp)
	if err != nil {
		return nil, err
	}
	return &BINResp, nil
}
