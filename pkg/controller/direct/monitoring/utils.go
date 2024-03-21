package monitoring

import (
	"github.com/googleapis/gax-go/v2/apierror"
	"k8s.io/klog/v2"
)

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {
	if err == nil {
		return false
	}
	switch err := err.(type) {
	case *apierror.APIError:
		if err.HTTPCode() == code {
			return true
		}
	default:
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}
