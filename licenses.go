package planhat

import "time"

// License represents a planhat license
type License struct {
	ID         string  `json:"_id"`
	ExternalID string  `json:"externalId"`
	Value      float64 `json:"value"`
	Currency   struct {
		ID        string `json:"_id"`
		Symbol    string `json:"symbol"`
		Rate      int    `json:"rate"`
		IsBase    bool   `json:"isBase"`
		Overrides struct {
		} `json:"overrides"`
	} `json:"_currency"`
	FromDate           time.Time `json:"fromDate"`
	ToDate             time.Time `json:"toDate"`
	Product            string    `json:"product"`
	CompanyID          string    `json:"companyId"`
	Custom             Custom    `json:"custom"`
	CompanyName        string    `json:"companyName"`
	Status             string    `json:"status"`
	RenewalStatus      string    `json:"renewalStatus"`
	FixedPeriod        bool      `json:"fixedPeriod"`
	ToDateIncluded     bool      `json:"toDateIncluded"`
	Length             float64   `json:"length"`
	Mrr                float64   `json:"mrr"`
	RenewalPeriod      float64   `json:"renewalPeriod"`
	RenewalUnit        string    `json:"renewalUnit"`
	RenewalDate        time.Time `json:"renewalDate"`
	RenewalDaysFromNow int       `json:"renewalDaysFromNow"`
	NoticePeriod       float64   `json:"noticePeriod"`
	NoticeUnit         string    `json:"noticeUnit"`
	IsOverdue          bool      `json:"isOverdue"`
}
