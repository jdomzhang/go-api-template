package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jdomzhang/resty"
)

func GetRawResponse(resp *resty.Response) string {
	if resp.RawResponse == nil {
		return ""
	}
	rawHeader := fmt.Sprintf("%s %s\n", resp.RawResponse.Proto, resp.Status())
	for k, v := range resp.RawResponse.Header {
		rawHeader = fmt.Sprintf("%s%s: %v\n", rawHeader, k, strings.Join(v[:], "; "))
	}

	rawBody := ""
	if resp.RawBody() != nil {
		rawBody = "\n" + fmt.Sprintf("%v", resp)
	}

	return rawHeader + rawBody
}

func GetRawRequest(req *resty.Request) string {
	rawHeader := fmt.Sprintf("%s %s %s\n", req.Method, req.URL, req.RawRequest.Proto)
	for k, v := range req.RawRequest.Header {
		rawHeader = fmt.Sprintf("%s%s: %v\n", rawHeader, k, strings.Join(v[:], "; "))
	}

	rawBody := ""
	if req.Body != nil {
		if jsonString, err := json.Marshal(req.Body); err == nil {
			// rawBody = "\n" + fmt.Sprintf("%v", req.Body)
			rawBody = "\n" + string(jsonString)
		}
	}

	return rawHeader + rawBody
}
