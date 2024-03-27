package mockkubeapiserver

import (
	"context"
	"encoding/json"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

type WatchCallback func(ev *watchEvent) error

type WatchOptions struct {
	Namespace string
}

func (s *MemoryStorage) Watch(ctx context.Context, resource *ResourceInfo, opt WatchOptions, callback WatchCallback) error {
	return resource.storage.watch(ctx, opt, callback)
}

type resourceStorage struct {
	GroupResource schema.GroupResource

	mutex   sync.Mutex
	watches []*watch

	objects map[types.NamespacedName]*unstructured.Unstructured
}

type watch struct {
	callback WatchCallback
	opt      WatchOptions
	errChan  chan error
}

func (r *resourceStorage) watch(ctx context.Context, opt WatchOptions, callback WatchCallback) error {
	w := &watch{
		callback: callback,
		opt:      opt,
		errChan:  make(chan error),
	}

	r.mutex.Lock()
	pos := -1
	for i := range r.watches {
		if r.watches[i] == nil {
			r.watches[i] = w
			pos = i
			break
		}
	}
	if pos == -1 {
		r.watches = append(r.watches, w)
		pos = len(r.watches) - 1
	}
	r.mutex.Unlock()

	// TODO: Delay / buffer watch notifications until after the list

	// TODO: Only send list if no rv specified?

	// TODO: Locking on r.objects

	for _, obj := range r.objects {
		if opt.Namespace != "" {
			if obj.GetNamespace() != opt.Namespace {
				continue
			}
		}

		ev := buildWatchEvent("ADDED", obj)
		if err := w.callback(ev); err != nil {
			klog.Warningf("error sending backfill watch notification; stopping watch: %v", err)

			// remove watch from list
			r.mutex.Lock()
			r.watches[pos] = nil
			r.mutex.Unlock()

			return err
		}
	}

	return <-w.errChan
}

func buildWatchEvent(evType string, u *unstructured.Unstructured) *watchEvent {
	ev := &watchEvent{
		internalObject: u,
		eventType:      evType,
		Namespace:      u.GetNamespace(),
	}
	return ev
}

func (ev *watchEvent) JSON() []byte {
	ev.mutex.Lock()
	defer ev.mutex.Unlock()

	if ev.json != nil {
		return ev.json
	}
	u := ev.internalObject

	msg := messageV1{
		Type:   ev.eventType,
		Object: u,
	}

	j, err := json.Marshal(&msg)
	if err != nil {
		klog.Fatalf("error from json.Marshal(%T): %v", &msg, err)
	}

	j = append(j, byte('\n'))
	ev.json = j

	return j
}

// Constructs the message for a PartialObjectMetadata response
func (ev *watchEvent) PartialObjectMetadataJSON() []byte {
	ev.mutex.Lock()
	defer ev.mutex.Unlock()

	if ev.partialObjectMetadataJSON != nil {
		return ev.partialObjectMetadataJSON
	}
	u := ev.internalObject

	partialObjectMetadata := &metav1.PartialObjectMetadata{}
	partialObjectMetadata.APIVersion = u.GetAPIVersion()
	partialObjectMetadata.Kind = u.GetKind()

	partialObjectMetadata.APIVersion = "meta.k8s.io/v1beta1"
	partialObjectMetadata.Kind = "PartialObjectMetadata"
	// {"kind":"PartialObjectMetadata","apiVersion":"meta.k8s.io/v1beta1","metadata"":

	partialObjectMetadata.Annotations = u.GetAnnotations()
	partialObjectMetadata.Labels = u.GetLabels()
	partialObjectMetadata.Name = u.GetName()
	partialObjectMetadata.Namespace = u.GetNamespace()
	partialObjectMetadata.ResourceVersion = u.GetResourceVersion()
	partialObjectMetadata.Generation = u.GetGeneration()
	partialObjectMetadata.CreationTimestamp = u.GetCreationTimestamp()
	partialObjectMetadata.DeletionTimestamp = u.GetDeletionTimestamp()
	partialObjectMetadata.DeletionGracePeriodSeconds = u.GetDeletionGracePeriodSeconds()
	partialObjectMetadata.GenerateName = u.GetGenerateName()

	msg := messageV1{
		Type:   ev.eventType,
		Object: partialObjectMetadata,
	}

	j, err := json.Marshal(&msg)
	if err != nil {
		klog.Fatalf("error from json.Marshal(%T): %v", &msg, err)
	}

	j = append(j, byte('\n'))
	ev.partialObjectMetadataJSON = j
	return j
}

func (r *resourceStorage) broadcastEventHoldingLock(ctx context.Context, evType string, u *unstructured.Unstructured) {
	ev := buildWatchEvent(evType, u)

	// r.mutex should be locked
	for i := range r.watches {
		w := r.watches[i]
		if w == nil {
			continue
		}
		if w.opt.Namespace != "" && ev.Namespace != w.opt.Namespace {
			continue
		}
		if err := w.callback(ev); err != nil {
			klog.Warningf("error sending watch notification; stopping watch: %v", err)
			w.errChan <- err
			r.watches[i] = nil
		}
	}
}
