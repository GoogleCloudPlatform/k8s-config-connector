package testcontext

import (
	"context"
	"net/http"
)

// httpRoundTripperKey is the key value for http.RoundTripper in a context.Context
var httpRoundTripperKey httpRoundTripperKeyType

type httpRoundTripperKeyType int

// HTTPRoundTripperFromContext returns the http.RoundTripper from the context, or nil if none is set.
func HTTPRoundTripperFromContext(ctx context.Context) http.RoundTripper {
	v := ctx.Value(httpRoundTripperKey)
	if v != nil {
		return v.(http.RoundTripper)
	}
	return nil
}

// WithHTTPRoundTripper stores the http.RoundTripper in the context.
func WithHTTPRoundTripper(ctx context.Context, httpRoundTripper http.RoundTripper) context.Context {
	return context.WithValue(ctx, httpRoundTripperKey, httpRoundTripper)
}
