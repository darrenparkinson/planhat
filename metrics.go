package planhat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// MetricsListOptions represents query parameters for listing metrics.  They are pointer values
// in order to distinguish between unset fields and those with set to a zero value.  Use the helper
// function planhat.Int() or planhat.String() to set the values.
type MetricsListOptions struct {
	// Id of company.
	CID *string `url:"cid,omitempty"`

	// Id of the dimension data.
	DimID *string `url:"dimid,omitempty"`

	// Days format integer representing the start day period (Days since January 1, 1970, Unix epoch).
	From *int `url:"from,omitempty"`

	// Days format integer representing the end day period (Days since January 1, 1970, Unix epoch).
	To *int `url:"to,omitempty"`

	// Limit the list length.
	Limit *int `url:"limit,omitempty"`

	// Start the list on a specific integer index.
	Offset *int `url:"offset,omitempty"`
}

// Metric represents an item that can be pushed to planhat.
type Metric struct {
	// Any string without spaces or special characters. If you're sending "Share of Active Users" a good dimensionId
	// might be "activeusershare". It's not displayed in Planhat but will be used when building Health Metrics in
	// Planhat. Required.
	DimensionID *string `json:"dimensionId,omitempty"`
	// The raw (number) value you would like to set. Required.
	Value *float64 `json:"value,omitempty"`
	// This is the model (company by default) external id in your systems. For this to work the objects in Planhat
	// will need to have this externalId set. Required.
	ExternalID *string `json:"externalId,omitempty"`
	// Company (default), EndUser, Asset and Project models are supported
	Model *string `json:"model,omitempty"`
	// Pass a valid ISO format date string to specify the date of the event. In none is provided we will use the time the request was received.
	Date *string `json:"date,omitempty"`
}

// DimensionData represents metrics data from planhat for the list operation.
type DimensionData struct {
	ID          string    `json:"_id"`
	DimensionID string    `json:"dimensionId"`
	Time        time.Time `json:"time"`
	Value       float64   `json:"value"`
	Model       string    `json:"model"`
	ParentID    string    `json:"parentId"`
	CompanyID   string    `json:"companyId"`
	CompanyName string    `json:"companyName"`
}

// List returns a list of DimensionData items as per the documentation https://docs.planhat.com/#get_metrics
func (s *MetricsService) List(ctx context.Context, options ...*MetricsListOptions) ([]*DimensionData, error) {
	dd := []*DimensionData{}
	url := fmt.Sprintf("%s/dimensiondata", s.client.BaseURL)
	for _, option := range options {
		var err error
		url, err = addOptions(url, option)
		if err != nil {
			return dd, err
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dd, err
	}
	if err := s.client.makeRequest(ctx, req, &dd); err != nil {
		return dd, err
	}
	return dd, nil
}

// BulkUpsert will update metrics. To push dimension data into Planhat it is required to specify the Tenant
// Token (tenantUUID) in the request URL.  This token is a simple uui identifier for your tenant and it can
// be found in the Developer module under the Tokens section.  Set the TenantUUID on the planhat Client.
// For more information, see the [planhat docs](https://docs.planhat.com/#bulkupsert_metrics)
func (s *MetricsService) BulkUpsert(ctx context.Context, metrics []Metric) (*UpsertMetricsResponse, error) {
	if s.client.TenantUUID == "" {
		return nil, ErrMissingTenantUUID
	}
	url := fmt.Sprintf("%s/%s", s.client.MetricsURL, s.client.TenantUUID)
	payload, err := json.Marshal(metrics)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	umr := &UpsertMetricsResponse{}
	if err := s.client.makeRequest(ctx, req, umr); err != nil {
		return umr, err
	}
	return umr, nil
}
