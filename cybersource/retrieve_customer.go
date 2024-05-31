package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/response"
	"io/ioutil"
)

func (c *Client) RetrieveCustomer(customerID string) (*response.CustomersResponse, error) {
	resource := "/tms/v2/customers/" + customerID
	resp, err := c.doGet(resource)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customerResp response.CustomersResponse
	err = json.Unmarshal(body, &customerResp)
	if err != nil {
		return nil, err
	}
	return &customerResp, nil
}
