package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/payment/request"
	"github.com/jjonline/cybersource-sdk-go/model/payment/response"
	"io/ioutil"
)

func (c *Client) CapturePayment(requestID string, req *request.CapturePaymentRequest) (*response.CaptureResponse, error) {
	resource := "/pts/v2/payments/" + requestID + "/captures"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var capturePaymentResp response.CaptureResponse
	err = json.Unmarshal(body, &capturePaymentResp)
	if err != nil {
		return nil, err
	}
	return &capturePaymentResp, nil
}
