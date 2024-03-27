package mockkubeapiserver

import (
	"net/http"
)

// Hook is the base interface implemented by a hook
type Hook interface {
}

// HTTPOperation contains the details of an HTTP operation
type HTTPOperation struct {
	Request *http.Request
}

// BeforeHTTPOperation is implemented by hooks that want to be called before every HTTP operation
type BeforeHTTPOperation interface {
	BeforeHTTPOperation(op *HTTPOperation)
}
