package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io"
)

func (c *Client) ProcessCredit(req *request.CreateCreditRequest) (*response.CreditResponse, int, error) {
	resource := "/pts/v2/credits"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var creditResp response.CreditResponse
	err = json.Unmarshal(body, &creditResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &creditResp, resp.StatusCode, nil
}
