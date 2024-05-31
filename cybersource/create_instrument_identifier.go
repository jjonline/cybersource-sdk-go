package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/request"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/response"
	"io"
)

func (c *Client) CreateInstrumentIdentifier(req *request.InstrumentIdentifierRequest) (*response.InstrumentIdentifierResponse, error) {
	resource := "/tms/v1/instrumentidentifiers"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var instrumentIdentifierResp response.InstrumentIdentifierResponse
	err = json.Unmarshal(body, &instrumentIdentifierResp)
	if err != nil {
		return nil, err
	}
	return &instrumentIdentifierResp, nil
}
