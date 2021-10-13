package planhat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AssetListOptions represents query parameters for listing assets.  They are pointer values
// in order to distinguish between unset fields and those with set to a zero value.  Use the helper
// function planhat.Int() or planhat.String() to set the values.
type AssetListOptions struct {
	// Limit the list length.
	Limit *int `url:"limit,omitempty"`

	// Start the list on a specific integer index.
	Offset *int `url:"offset,omitempty"`

	// Sort based on a specific property. Prefix the property "-" to change the sort order.
	Sort *string `url:"sort,omitempty"`

	// Select specific properties. This is case sensitive and currently needs to be the planhat names as
	// a comma separated string, e.g. "companyid,name".
	Select *string `url:"select,omitempty"`
}

// Asset represents a planhat asset.
type Asset struct {
	ID         *string                `json:"_id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	CompanyID  *string                `json:"companyId,omitempty"`
	ExternalID *string                `json:"externalId,omitempty"`
	SourceID   *string                `json:"sourceId,omitempty"`
	Custom     map[string]interface{} `json:"custom,omitempty"`
}

// Create creates a new asset record
func (s *AssetService) Create(ctx context.Context, asset Asset) (*Asset, error) {
	as := &Asset{}
	url := fmt.Sprintf("%s/assets", s.client.BaseURL)
	payload, err := json.Marshal(asset)
	if err != nil {
		return as, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return as, err
	}
	if err := s.client.makeRequest(ctx, req, as); err != nil {
		return as, err
	}
	return as, nil
}

// Update will update a planhat asset.
// To update an asset it is required to pass the asset _id in the request.
// Alternately it is possible to update using the asset externalId and/or sourceId adding a prefix and passing one of
// these keyables as identifiers. e.g. extid-{{externalId}} or srcid-{{sourceId}}
func (s *AssetService) Update(ctx context.Context, id string, asset Asset) (*Asset, error) {
	as := &Asset{}
	url := fmt.Sprintf("%s/assets/%s", s.client.BaseURL, id)
	payload, err := json.Marshal(asset)
	if err != nil {
		return as, err
	}
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(payload)))
	if err != nil {
		return as, err
	}
	if err := s.client.makeRequest(ctx, req, as); err != nil {
		return as, err
	}
	return as, nil
}

// Get returns a single asset given it's planhat ID
// Alternately it's possible to get an asset using its externalId and/or sourceId adding a prefix and passing one of
// these keyables as identifiers. e.g. extid-{{externalId}} or srcid-{{sourceId}}.  Helper functions have also
// been provided for this.
func (s *AssetService) Get(ctx context.Context, id string) (*Asset, error) {
	as := &Asset{}
	url := fmt.Sprintf("%s/assets/%s", s.client.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return as, err
	}
	if err := s.client.makeRequest(ctx, req, &as); err != nil {
		return as, err
	}
	return as, nil
}

// GetByExternalID retrieves an asset using it's external ID
func (s *AssetService) GetByExternalID(ctx context.Context, externalID string) (*Asset, error) {
	return s.Get(ctx, fmt.Sprintf("extid-%s", externalID))
}

// GetBySourceID retrieves an asset using it's source ID
func (s *AssetService) GetBySourceID(ctx context.Context, sourceID string) (*Asset, error) {
	return s.Get(ctx, fmt.Sprintf("srcid-%s", sourceID))
}

// List will list assets based on the AssetListOptions provided
func (s *AssetService) List(ctx context.Context, options ...*AssetListOptions) ([]*Asset, error) {
	ar := []*Asset{}

	url := fmt.Sprintf("%s/assets", s.client.BaseURL)
	for _, option := range options {
		var err error
		url, err = addOptions(url, option)
		if err != nil {
			return ar, err
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ar, err
	}
	if err := s.client.makeRequest(ctx, req, &ar); err != nil {
		return ar, err
	}
	return ar, nil
}

// Delete is used delete an asset. It is required to pass the _id (ID).
func (s *AssetService) Delete(ctx context.Context, id string) (*DeleteResponse, error) {
	url := fmt.Sprintf("%s/assets/%s", s.client.BaseURL, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	dr := &DeleteResponse{}
	if err := s.client.makeRequest(ctx, req, dr); err != nil {
		return dr, err
	}
	return dr, nil
}

// BulkUpsert will update or insert assets.
// To create an asset it's required define a name and a valid companyId.
// To update an asset it is required to specify in the payload one of the following keyables:
//   _id, sourceId and/or externalId.
// Since this is a bulk upsert operation it's possible create and/or update multiple assets with the same payload.
// Note there is an upper limit of 50,000 items per request.
// For more information, see the [planhat docs](https://docs.planhat.com/#bulk_upsert)
func (s *AssetService) BulkUpsert(ctx context.Context, assets []Asset) (*UpsertResponse, error) {
	url := fmt.Sprintf("%s/assets", s.client.BaseURL)
	payload, err := json.Marshal(assets)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	ur := &UpsertResponse{}
	if err := s.client.makeRequest(ctx, req, ur); err != nil {
		return ur, err
	}
	return ur, nil
}
