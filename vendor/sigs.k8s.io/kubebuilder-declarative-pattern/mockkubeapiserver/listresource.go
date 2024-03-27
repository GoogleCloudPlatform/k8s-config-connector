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
	"strings"
	"sync"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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

	// TODO: we should parse the accept header correctly
	partialObjectMetadata := false
	accept := req.r.Header.Get("accept")
	if strings.Contains(accept, "PartialObjectMetadata") {
		partialObjectMetadata = true
	}

	watchParam := query.Get("watch")
	if watchParam == "true" {
		return req.doWatch(ctx, s, resource, partialObjectMetadata)
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
	internalObject *unstructured.Unstructured
	eventType      string

	Namespace string

	mutex                     sync.Mutex
	partialObjectMetadataJSON []byte
	json                      []byte
}

type messageV1 struct {
	Type   string         `json:"type"`
	Object runtime.Object `json:"object"`
}

func (req *listResource) doWatch(ctx context.Context, s *MockKubeAPIServer, resource *ResourceInfo, partialObjectMetadata bool) error {
	w := req.w

	contentType := "application/json"
	if partialObjectMetadata {
		contentType = "application/json;as=PartialObjectMetadataList;g=meta.k8s.io;v=v1"
	}

	w.Header().Add("Content-Type", contentType)
	w.Header().Add("Cache-Control", "no-cache, private")

	var opt WatchOptions
	opt.Namespace = req.Namespace

	w.WriteHeader(200)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	return s.storage.Watch(ctx, resource, opt, func(ev *watchEvent) error {
		klog.V(2).Infof("sending watch event %s", string(ev.JSON()))
		if partialObjectMetadata {
			if _, err := w.Write(ev.PartialObjectMetadataJSON()); err != nil {
				return err
			}
		} else {
			if _, err := w.Write(ev.JSON()); err != nil {
				return err
			}
		}
		// TODO: Should we try to debounce flushes?  We could have a queue and send everything in the queue before flushing
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		return nil
	})
}
