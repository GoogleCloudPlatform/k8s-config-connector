//go:build controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14

package commonclient

import (
	"context"

	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// SourceKind is a version-indendenent abstraction over calling source.Kind
func SourceKind(cache cache.Cache, obj client.Object) source.SyncingSource {
	return source.NewKindWithCache(obj, cache)
}

// WrapEventHandler is a version-indendenent abstraction over handler.EventHandler
func WrapEventHandler(h EventHandler) handler.EventHandler {
	return &eventHandlerWithoutContext{h: h}
}

type eventHandlerWithoutContext struct {
	h EventHandler
}

func (h *eventHandlerWithoutContext) Create(ev event.CreateEvent, q workqueue.RateLimitingInterface) {
	h.h.Create(context.TODO(), ev, q)
}
func (h *eventHandlerWithoutContext) Update(ev event.UpdateEvent, q workqueue.RateLimitingInterface) {
	h.h.Update(context.TODO(), ev, q)
}
func (h *eventHandlerWithoutContext) Delete(ev event.DeleteEvent, q workqueue.RateLimitingInterface) {
	h.h.Delete(context.TODO(), ev, q)
}
func (h *eventHandlerWithoutContext) Generic(ev event.GenericEvent, q workqueue.RateLimitingInterface) {
	h.h.Generic(context.TODO(), ev, q)
}

// EventHandler is the controller-runtime 0.15 version of EventHandler (with a context argument)
type EventHandler interface {
	// Create is called in response to an create event - e.g. Pod Creation.
	Create(context.Context, event.CreateEvent, workqueue.RateLimitingInterface)

	// Update is called in response to an update event -  e.g. Pod Updated.
	Update(context.Context, event.UpdateEvent, workqueue.RateLimitingInterface)

	// Delete is called in response to a delete event - e.g. Pod Deleted.
	Delete(context.Context, event.DeleteEvent, workqueue.RateLimitingInterface)

	// Generic is called in response to an event of an unknown type or a synthetic event triggered as a cron or
	// external trigger request - e.g. reconcile Autoscaling, or a Webhook.
	Generic(context.Context, event.GenericEvent, workqueue.RateLimitingInterface)
}
