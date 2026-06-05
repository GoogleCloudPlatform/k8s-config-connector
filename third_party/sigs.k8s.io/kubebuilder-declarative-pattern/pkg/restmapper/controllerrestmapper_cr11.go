//go:build controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14

package restmapper

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

// NewControllerRESTMapper is the constructor for a ControllerRESTMapper.
func NewControllerRESTMapper(restConfig *rest.Config) (meta.RESTMapper, error) {
	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error from HTTPClientFor: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return nil, err
	}

	return &ControllerRESTMapper{
		uncached: discoveryClient,
		cache:    newCache(),
	}, nil
}

// NewForTest creates a ControllerRESTMapper, but is intended to be a common interface for use by tests.
func NewForTest(cfg *rest.Config) (meta.RESTMapper, error) {
	return NewControllerRESTMapper(cfg)
}
