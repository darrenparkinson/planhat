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
