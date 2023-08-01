/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockkubeapiserver

import (
	"context"
	"net/http"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

// listResource is a request to list resources
type listResource struct {
	baseRequest

	Group     string
	Version   string
	Resource  string
	Namespace string
}

// Run serves the http request
func (req *listResource) Run(ctx context.Context, s *MockKubeAPIServer) error {
	gr := schema.GroupResource{Group: req.Group, Resource: req.Resource}

	resource := s.storage.FindResource(gr)
	if resource == nil {
		return req.writeErrorResponse(http.StatusNotFound)
	}

	query := req.r.URL.Query()

	watchParam := query.Get("watch")
	if watchParam == "true" {
		return req.doWatch(ctx, s, resource)
	}

	var filter ListFilter
	filter.Namespace = req.Namespace
	objects, err := s.storage.ListObjects(ctx, resource, filter)
	if err != nil {
		return err
	}

	objects.SetGroupVersionKind(resource.ListGVK)

	return req.writeResponse(objects)
}

type watchEvent struct {
	Message messageV1
	JSON    []byte

	Namespace string
}

type messageV1 struct {
	Type   string                     `json:"type"`
	Object *unstructured.Unstructured `json:"object"`
}

func (req *listResource) doWatch(ctx context.Context, s *MockKubeAPIServer, resource *ResourceInfo) error {
	w := req.w
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "no-cache, private")

	var opt WatchOptions
	opt.Namespace = req.Namespace

	return s.storage.Watch(ctx, resource, opt, func(ev *watchEvent) error {
		klog.Infof("sending watch event %s", string(ev.JSON))
		if _, err := w.Write(ev.JSON); err != nil {
			return err
		}
		// TODO: Should we try to debounce flushes?  We could have a queue and send everything in the queue before flushing
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		return nil
	})
}
