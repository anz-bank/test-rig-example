// Code generated by sysl DO NOT EDIT.
package dep2

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for dep2
type Service interface {
	GetBarbar(ctx context.Context, req *GetBarbarRequest) (*Str, error)
}

// Client for dep2 API
type Client struct {
	client *http.Client
	url    string
}

// NewClient for dep2
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL}
}

// GetBarbar ...
func (s *Client) GetBarbar(ctx context.Context, req *GetBarbarRequest) (*Str, error) {
	required := []string{}
	var okResponse Str

	u, err := url.Parse(fmt.Sprintf("%s/barbar/%v", s.url, req.ID))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, &okResponse, nil)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: dep2 <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}

	OkStrResponse, ok := result.Response.(*Str)
	if ok {
		valErr := validator.Validate(OkStrResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkStrResponse, nil
	}

	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}