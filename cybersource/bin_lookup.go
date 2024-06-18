package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) LookupBIN(req *request.BINLookupRequest) (*response.BINLookupResponse, int, error) {
	resource := "/bin/v1/binlookup"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var BINResp response.BINLookupResponse
	err = json.Unmarshal(body, &BINResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &BINResp, resp.StatusCode, nil
}
