package cybersource

import (
	"errors"
	"io"
)

func (c *Client) DeleteCustomer(customerID string) error {
	resource := "/tms/v2/customers/" + customerID
	resp, err := c.doDelete(resource)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}
