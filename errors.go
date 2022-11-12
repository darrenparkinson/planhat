package planhat

// Err implements the error interface so we can have constant errors.
type Err string

func (e Err) Error() string {
	return string(e)
}

// Error Constants
// Cisco documents these as the only error responses they will emit.
const (
	ErrBadRequest        = Err("planhat: bad request")
	ErrUnauthorized      = Err("planhat: unauthorized request")
	ErrForbidden         = Err("planhat: forbidden")
	ErrNotFound          = Err("planhat: not found")
	ErrInternalError     = Err("planhat: internal error")
	ErrUnknown           = Err("planhat: unexpected error occurred")
	ErrMissingTenantUUID = Err("planhat: missing required tenant uuid for this request")
)
