//go:build controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14

package commonclient

import (
	"net/http"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// NewDiscoveryRESTMapper is a version-independent wrapper around apiutil.NewDiscoveryRESTMapper
func NewDiscoveryRESTMapper(c *rest.Config, httpClient *http.Client) (meta.RESTMapper, error) {
	return apiutil.NewDiscoveryRESTMapper(c)
}
