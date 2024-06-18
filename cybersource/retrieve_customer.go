package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/response"
	"io"
)

func (c *Client) RetrieveCustomer(customerID string) (*response.CustomersResponse, int, error) {
	resource := "/tms/v2/customers/" + customerID
	resp, err := c.doGet(resource)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var customerResp response.CustomersResponse
	err = json.Unmarshal(body, &customerResp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &customerResp, resp.StatusCode, nil
}
