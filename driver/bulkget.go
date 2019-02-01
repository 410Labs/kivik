package driver

import "context"

// BulkDocReference is a reference to a document given in a BulkGet query.
type BulkDocReference struct {
	ID        string `json:"id"`
	Rev       string `json:"rev,omitempty"`
	AttsSince string `json:"atts_since,omitempty"`
}

// BulkGetter is an optional interface which may be implemented by a driver to
// support bulk get operations.
type BulkGetter interface {
	// BulkGet uses the _bulk_get interface to fetch multiple documents in a single query.
	BulkGet(ctx context.Context, docs []BulkDocReference, options map[string]interface{}) (Rows, error)
}
