//go:build !(controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14)

package commonclient

import (
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// SourceKind is a version-indendenent abstraction over calling source.Kind
func SourceKind(cache cache.Cache, obj client.Object) source.Source {
	return source.Kind(cache, obj)
}

// WrapEventHandler is a version-indendenent abstraction over handler.EventHandler
func WrapEventHandler(h handler.EventHandler) handler.EventHandler {
	return h
}

type EventHandler = handler.EventHandler
