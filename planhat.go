// Copyright 2021 Darren Parkinson. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run gen-accessors.go

package planhat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

// Client is the main planhat client for interacting with the library.  It can be created using NewClient
type Client struct {
	// BaseURL for Planhat API.  Set to https://api-eu3.planhat.com using `planhat.New()`, or set directly.
	BaseURL string

	// MetricsURL for Planhat API.  Set to https://analytics.planhat.com/dimensiondata as per the planhat docs.
	MetricsURL string

	//HTTP Client to use for making requests, allowing the user to supply their own if required.
	HTTPClient *http.Client

	//API Key for Planhat.
	APIKey string

	//TenantUUID for posting to the metrics endpoint.  Only required if you're sending in metrics.
	TenantUUID string

	MetricsService *MetricsService
	AssetService   *AssetService
	CompanyService *CompanyService
	EndUserService *EndUserService
	UserService    *UserService

	lim *rate.Limiter
}

// MetricsService represents the Metrics methods
type MetricsService struct {
	client *Client
}

// AssetService represents the Assets group
type AssetService struct {
	client *Client
}

// CompanyService represents the Company group
type CompanyService struct {
	client *Client
}

// EndUserService represents the End Users group
type EndUserService struct {
	client *Client
}

// UserService represents the Users group
type UserService struct {
	client *Client
}

// NewClient is a helper function that returns an new planhat client given a region and an API Key.
// Optionally you can provide your own http client or use nil to use the default.  This is done to
// ensure you're aware of the decision you're making to not provide your own http client.
func NewClient(apikey string, cluster string, client *http.Client) (*Client, error) {
	if apikey == "" {
		return nil, errors.New("apikey required")
	}
	apicluster := ""
	if cluster == "" {
		apicluster = "api"
	} else {
		apicluster = fmt.Sprintf("api-%s", cluster)
	}
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	rl := rate.NewLimiter(150, 1)
	c := &Client{
		BaseURL:    fmt.Sprintf("https://%s.planhat.com", apicluster),
		MetricsURL: "https://analytics.planhat.com/dimensiondata",
		HTTPClient: client,
		APIKey:     apikey,
		lim:        rl,
	}
	c.MetricsService = &MetricsService{client: c}
	c.AssetService = &AssetService{client: c}
	c.CompanyService = &CompanyService{client: c}
	c.EndUserService = &EndUserService{client: c}
	c.UserService = &UserService{client: c}

	return c, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float64 is a helper routine that allocates a new Float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// makeRequest provides a single function to add common items to the request.
func (c *Client) makeRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if !c.lim.Allow() {
		c.lim.Wait(ctx)
	}

	rc := req.WithContext(ctx)
	res, err := c.HTTPClient.Do(rc)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {

		var planhatErr error

		switch res.StatusCode {
		case 400:
			planhatErr = ErrBadRequest
		case 401:
			planhatErr = ErrUnauthorized
		case 403:
			planhatErr = ErrForbidden
		case 500:
			planhatErr = ErrInternalError
		default:
			planhatErr = ErrUnknown
		}

		return planhatErr

	}

	if res.StatusCode == http.StatusCreated {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
