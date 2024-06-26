package cybersource

import (
	"encoding/json"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/request"
	"github.com/jjonline/cybersource-sdk-go/model/token_management/response"
	"io"
	"net/http"
)

func (c *Client) GenerateCaptureContext(req *request.GenerateCaptureContextRequest) (*response.GenerateCaptureContextResponse, int, error) {
	resource := "/microform/v2/sessions"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	var captureContextResp response.GenerateCaptureContextResponse
	if resp.StatusCode != http.StatusCreated {
		err = json.Unmarshal(body, &captureContextResp)
		if err != nil {
			return nil, resp.StatusCode, err
		}
	} else {
		captureContextResp.CaptureContext = string(body)
	}

	return &captureContextResp, resp.StatusCode, nil
}
