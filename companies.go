package planhat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// CompanyListOptions represents query parameters for listing companies.  They are pointer values
// in order to distinguish between unset fields and those with set to a zero value.  Use the helper
// function planhat.Int() or planhat.String() to set the values.
type CompanyListOptions struct {
	// Limit the list length.
	Limit *int `url:"limit,omitempty"`

	// Start the list on a specific integer index.
	Offset *int `url:"offset,omitempty"`

	// Sort based on a specific property. Prefix the property "-" to change the sort order.
	Sort *string `url:"sort,omitempty"`
}

// LeanCompanyListOptions represents query parameters for listing companies.  They are pointer values
// in order to distinguish between unset fields and those with set to a zero value.  Use the helper
// function planhat.String() to set the value.
type LeanCompanyListOptions struct {
	ExternalID *string `url:"externalId,omitempty"`
	SourceID   *string `url:"sourceId,omitempty"`
	Status     *string `url:"status,omitempty"`
}

// LeanCompany is the result of the Get lean list endpoint
type LeanCompany struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	ExternalID string `json:"externalId"`
	SourceID   string `json:"sourceId"`
	Slug       string `json:"slug"`
}

// Company represents a planhat company.
type Company struct {
	// CoOwner of the company.  Empty interface due to GetCompany returning an ID string and GetCompanies returning an ID and Nickname
	CoOwner       *interface{}           `json:"coOwner,omitempty"`
	CSMScore      *int                   `json:"csmScore,omitempty"`
	Custom        map[string]interface{} `json:"custom,omitempty"`
	CustomerFrom  *time.Time             `json:"customerFrom,omitempty"`
	CustomerTo    *time.Time             `json:"customerTo,omitempty"`
	ExternalID    *string                `json:"externalId,omitempty"`
	H             *int                   `json:"h,omitempty"`
	ID            *string                `json:"_id,omitempty"`
	LastRenewal   *time.Time             `json:"lastRenewal,omitempty"`
	LastTouch     *interface{}           `json:"lastTouch,omitempty"`
	LastTouchType *interface{}           `json:"lastTouchType,omitempty"`
	Licenses      *[]License             `json:"licenses,omitempty"`
	MR            *float64               `json:"mr,omitempty"`
	MRR           *float64               `json:"mrr,omitempty"`
	MRRTotal      *float64               `json:"mrrTotal,omitempty"`
	MRTotal       *float64               `json:"mrTotal,omitempty"`
	Name          *string                `json:"name,omitempty"`
	NRR30         *int                   `json:"nrr30,omitempty"`
	NRRTotal      *int                   `json:"nrrTotal,omitempty"`
	// Owner of the company.  Empty interface due to GetCompany returning an ID string and GetCompanies returning an ID and Nickname
	Owner              *interface{} `json:"owner,omitempty"`
	Phase              *string      `json:"phase,omitempty"`
	PhaseSince         *time.Time   `json:"phaseSince,omitempty"`
	Products           *[]string    `json:"products,omitempty"`
	RenewalDate        *time.Time   `json:"renewalDate,omitempty"`
	RenewalDaysFromNow *int         `json:"renewalDaysFromNow,omitempty"`
	Status             *string      `json:"status,omitempty"`
}

// List will list companies based on the CompanyListOptions provided
func (s *CompanyService) List(ctx context.Context, options ...*CompanyListOptions) ([]*Company, error) {
	cr := []*Company{}

	url := fmt.Sprintf("%s/companies", s.client.BaseURL)
	for _, option := range options {
		var err error
		url, err = addOptions(url, option)
		if err != nil {
			return cr, err
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return cr, err
	}
	if err := s.client.makeRequest(ctx, req, &cr); err != nil {
		return cr, err
	}
	return cr, nil
}

// LeanList returns a lightweight list of all companies in Planhat to match against your own ids etc.
func (s *CompanyService) LeanList(ctx context.Context, options ...*LeanCompanyListOptions) ([]*LeanCompany, error) {
	cr := []*LeanCompany{}

	url := fmt.Sprintf("%s/leancompanies", s.client.BaseURL)
	for _, option := range options {
		var err error
		url, err = addOptions(url, option)
		if err != nil {
			return cr, err
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return cr, err
	}
	if err := s.client.makeRequest(ctx, req, &cr); err != nil {
		return cr, err
	}
	return cr, nil
}

func (s *CompanyService) Get(ctx context.Context, id string) (*Company, error) {
	co := &Company{}
	url := fmt.Sprintf("%s/companies/%s", s.client.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return co, err
	}
	if err := s.client.makeRequest(ctx, req, &co); err != nil {
		return co, err
	}
	return co, nil
}

// GetByExternalID retrieves a company using it's external ID
func (s *CompanyService) GetByExternalID(ctx context.Context, externalID string) (*Company, error) {
	co := &Company{}
	url := fmt.Sprintf("%s/companies/extid-%s", s.client.BaseURL, externalID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return co, err
	}
	if err := s.client.makeRequest(ctx, req, co); err != nil {
		return co, err
	}
	return co, nil
}

// GetBySourceID retrieves a company using it's source ID
func (s *CompanyService) GetBySourceID(ctx context.Context, sourceID string) (*Company, error) {
	co := &Company{}
	url := fmt.Sprintf("%s/companies/srcid-%s", s.client.BaseURL, sourceID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return co, err
	}
	if err := s.client.makeRequest(ctx, req, co); err != nil {
		return co, err
	}
	return co, nil
}

// Create creates a new company record
func (s *CompanyService) Create(ctx context.Context, company Company) (*Company, error) {
	co := &Company{}
	url := fmt.Sprintf("%s/companies", s.client.BaseURL)
	payload, err := json.Marshal(company)
	if err != nil {
		return co, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return co, err
	}
	if err := s.client.makeRequest(ctx, req, co); err != nil {
		return co, err
	}
	return co, nil
}

// BulkUpsert will update or insert companies.
// Note there is an upper limit of 50,000 items per request.
// For more information, see the [planhat docs](https://docs.planhat.com/#bulk_upsert)
func (s *CompanyService) BulkUpsert(ctx context.Context, companies []Company) (*UpsertResponse, error) {
	url := fmt.Sprintf("%s/companies", s.client.BaseURL)
	payload, err := json.Marshal(companies)
	log.Println(string(payload))
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
