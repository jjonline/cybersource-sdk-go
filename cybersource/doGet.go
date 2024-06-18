package cybersource

import (
	"net/http"
	"time"
)

func (c *Client) doGet(resource string) (*http.Response, error) {
	gmtDateTime := time.Now().UTC().String()
	url := "https://" + c.requestHost + resource

	digest := "" // Digest is not required for GET calls
	signature := c.createSignatureHeader(http.MethodGet, resource, digest, gmtDateTime)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("v-c-merchant-id", c.merchantId)
	req.Header.Set("Date", gmtDateTime)
	req.Header.Set("Host", c.requestHost)
	req.Header.Set("Signature", signature)

	httpClient := http.Client{}
	res, err := httpClient.Do(req)

	c.dumpDebugRequest(req)
	c.dumpDebugResponse(res)

	// confirm response res is not nil
	if res == nil {
		return &http.Response{}, err
	}
	return res, err
}
