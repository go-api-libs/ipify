package ipify

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-api-libs/api"
	"github.com/go-json-experiment/json"
)

var baseURLv6 = &url.URL{
	Host:   "api64.ipify.org",
	Scheme: "https",
}

// GetIP defines an operation.
//
//	GET /
func (c *Client) GetIPv6(ctx context.Context, params *GetIPParams) (*IPWrapper, error) {
	return GetIPv6[IPWrapper](ctx, c, params)
}

// GetIP defines an operation.
// You can define a custom result to unmarshal the response into.
//
//	GET /
func GetIPv6[R any](ctx context.Context, c *Client, params *GetIPParams) (*R, error) {
	u := baseURLv6.JoinPath("/")

	if params != nil && params.Format != "" {
		u.RawQuery = url.Values{"format": []string{params.Format}}.Encode()
	}

	req := (&http.Request{
		Header:     http.Header{"User-Agent": []string{userAgent}},
		Host:       u.Host,
		Method:     http.MethodGet,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	rsp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// Returns the IP address of the requester.
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			var out R
			if err := json.UnmarshalRead(rsp.Body, &out, jsonOpts); err != nil {
				return nil, api.WrapDecodingError(rsp, err)
			}

			return &out, nil
		default:
			return nil, api.NewErrUnknownContentType(rsp)
		}
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}
