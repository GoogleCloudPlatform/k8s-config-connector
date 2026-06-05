//go:build !(controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14)

package commonclient

import (
	"net/http"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// SourceKind is a version-indendenent abstraction over calling source.Kind
func SourceKind(cache cache.Cache, obj client.Object) source.Source {
	return source.Kind(cache, obj, &handler.TypedEnqueueRequestForObject[client.Object]{})
}

// SourceKind is a version-indendenent abstraction over calling source.Kind
func SourceKindWithHandler(cache cache.Cache, obj client.Object, handler handler.TypedEventHandler[client.Object, reconcile.Request]) source.Source {
	return source.Kind(cache, obj, handler)
}

// GetHTTPClient returns the http.Client associated with the Cluster
func GetHTTPClient(c cluster.Cluster) (*http.Client, error) {
	return c.GetHTTPClient(), nil
}

type EventHandler = handler.EventHandler
