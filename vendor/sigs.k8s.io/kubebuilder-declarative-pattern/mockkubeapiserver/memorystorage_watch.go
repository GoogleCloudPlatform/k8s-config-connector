package mockkubeapiserver

import (
	"context"
	"encoding/json"
	"sync"

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
			klog.Warningf("error sending watch notification; stopping watch: %v", err)

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
		Message: messageV1{
			Type:   evType,
			Object: u,
		},
		Namespace: u.GetNamespace(),
	}
	evJSON, err := json.Marshal(&ev.Message)
	if err != nil {
		klog.Fatalf("error from json.Marshal(%T): %v", &ev.Message, err)
	}

	evJSON = append(evJSON, byte('\n'))
	ev.JSON = evJSON

	return ev
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
