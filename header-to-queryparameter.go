// Package plugindemo a demo plugin.
package header_to_queryparameter

import (
	"context"
	"fmt"
	"net/http"
)

// the plugin configuration.
type Config struct {
	QueryParameter string `json:"query_parameter"`
	Header         string `json:"header"`
}

// CreateConfig creates the default plugin configuration
func CreateConfig() *Config {
	return &Config{
		QueryParameter: "v",
		Header:         "X-Version",
	}
}

type HeaderToQueryParameterMiddleware struct {
	next           http.Handler
	queryParameter string
	header         string
	name           string
}

// Creates a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Header) < 1 {
		return nil, fmt.Errorf("header cannot be empty string")
	}
	if len(config.QueryParameter) < 1 {
		return nil, fmt.Errorf("query parameter cannot be empty string")
	}

	return &HeaderToQueryParameterMiddleware{
		header:         config.Header,
		queryParameter: config.QueryParameter,
		next:           next,
		name:           name,
	}, nil
}

func (m *HeaderToQueryParameterMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	headers := req.Header
	header := headers[m.header]
	if len(header) > 0 {
		req.URL.Query().Set(m.queryParameter, header[0])
	}
	m.next.ServeHTTP(rw, req)
}
