//go:build !(controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14)

package restmapper

import (
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

// NewControllerRESTMapper is the constructor for a ControllerRESTMapper
func NewControllerRESTMapper(cfg *rest.Config, httpClient *http.Client) (meta.RESTMapper, error) {
	discoveryClient, err := discovery.NewDiscoveryClientForConfigAndClient(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	return &ControllerRESTMapper{
		uncached: discoveryClient,
		cache:    newCache(),
	}, nil
}

// NewForTest creates a ControllerRESTMapper, but is intended to be a common interface for use by tests.
func NewForTest(restConfig *rest.Config) (meta.RESTMapper, error) {
	client, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error from rest.HTTPClientFor: %w", err)
	}
	return NewControllerRESTMapper(restConfig, client)
}
