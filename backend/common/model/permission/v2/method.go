package v2

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

var methods = []string{MethodGet,MethodPost,MethodDelete,MethodPut,
	MethodConnect,MethodHead,MethodPatch,MethodOptions,MethodTrace}

