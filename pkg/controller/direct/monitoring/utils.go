package monitoring

import (
	"errors"

	"github.com/googleapis/gax-go/v2/apierror"
	"k8s.io/klog/v2"
)

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {
	if err == nil {
		return false
	}
	apiError := &apierror.APIError{}
	if errors.As(err, &apiError) {
		return apiError.HTTPCode() == code
	}
	klog.Warningf("unexpected error type %T", err)
	return false
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}
