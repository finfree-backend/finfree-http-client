package finfree_http_client

type pair interface {
	getKey() string
	getVal() string
}

// Header parameters
type header struct {
	key string
	val string
}

func (h *header) getKey() string {
	return h.key
}

func (h *header) getVal() string {
	return h.val
}

func NewHeader(key, val string) *header {
	return &header{key: key, val: val}
}

// Query parameters
type query struct {
	key string
	val string
}

func (q *query) getKey() string {
	return q.key
}

func (q *query) getVal() string {
	return q.val
}

func NewQuery(key, val string) *query {
	return &query{key: key, val: val}
}
