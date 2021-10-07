package planhat

// UpsertResponse is the result of a bulk upsert operation as documented in the [planhat docs](https://docs.planhat.com/#bulk_upsert).
// Note the use of slices of empty interfaces due to the lack of documentation on what an error object etc. are
type UpsertResponse struct {
	Created          int           `json:"created"`
	CreatedErrors    []interface{} `json:"createdErrors"`
	InsertsKeys      []interface{} `json:"insertsKeys"`
	Updated          int           `json:"updated"`
	UpdatedErrors    []interface{} `json:"updatedErrors"`
	UpdatesKeys      []interface{} `json:"updatesKeys"`
	NonUpdates       int           `json:"nonupdates"`
	Modified         []string      `json:"modified"`
	UpsertedIDs      []string      `json:"upsertedIds"`
	PermissionErrors []interface{} `json:"permissionErrors"`
}

// UpsertMetricsResponse is the result of a bulk upsert operation as documented in the [planhat docs](https://docs.planhat.com/#bulkupsert_metrics).
// Note the use of slices of empty interfaces due to the lack of documentation on what an error object is.
type UpsertMetricsResponse struct {
	Processed int           `json:"processed"`
	Errors    []interface{} `json:"errors"`
}

// DeleteResponse is returned by planhat when deleting an object
type DeleteResponse struct {
	N            int `json:"n"`
	OK           int `json:"ok"`
	DeletedCount int `json:"deletedCount"`
}
