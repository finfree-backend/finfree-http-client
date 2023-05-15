package finfree_http_client

type method string

const (
	MethodGet     = method("GET")
	MethodHead    = method("HEAD")
	MethodPost    = method("POST")
	MethodPut     = method("PUT")
	MethodPatch   = method("PATCH") // RFC 5789
	MethodDelete  = method("DELETE")
	MethodConnect = method("CONNECT")
	MethodOptions = method("OPTIONS")
	MethodTrace   = method("TRACE")
)
