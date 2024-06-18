package cybersource

import (
	"errors"
	"io"
)

func (c *Client) DeleteCustomer(customerID string) (error, int) {
	resource := "/tms/v2/customers/" + customerID
	resp, err := c.doDelete(resource)
	defer resp.Body.Close()
	if err != nil {
		return err, resp.StatusCode
	}

	if resp.StatusCode != 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err, resp.StatusCode
		}
		return errors.New(string(body)), resp.StatusCode
	}

	return nil, resp.StatusCode
}
