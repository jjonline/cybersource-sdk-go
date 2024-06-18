package cybersource

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func (c *Client) dumpDebugRequest(req *http.Request) {
	if c.isDebugPrint {
		bag, _ := httputil.DumpRequest(req, true)
		fmt.Println(string(bag))
	}
}

func (c *Client) dumpDebugResponse(res *http.Response) {
	if c.isDebugPrint {
		bag, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(bag))
	}
}
