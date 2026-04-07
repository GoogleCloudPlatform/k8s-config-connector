//go:build !(controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14)

package commonclient

import (
	"net/http"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// NewDiscoveryRESTMapper is a version-independent wrapper around creating a meta.RESTMapper
// It calls NewDynamicRESTMapper as of kubebuilder-declarative-pattern 0.17.
// Deprecated: prefer NewDynamicRESTMapper
func NewDiscoveryRESTMapper(c *rest.Config, httpClient *http.Client) (meta.RESTMapper, error) {
	return NewDynamicRESTMapper(c, httpClient)
}

// NewDynamicRESTMapper is a version-independent wrapper around apiutil.NewDynamicRESTMapper
func NewDynamicRESTMapper(c *rest.Config, httpClient *http.Client) (meta.RESTMapper, error) {
	return apiutil.NewDynamicRESTMapper(c, httpClient)
}
