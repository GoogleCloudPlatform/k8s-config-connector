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
	"fmt"
	"io"
	"net/http"
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

// putResource is a request to get a single resource
type putResource struct {
	resourceRequestBase
}

// Run serves the http request
func (req *putResource) Run(ctx context.Context, s *MockKubeAPIServer) error {
	gr := schema.GroupResource{Group: req.Group, Resource: req.Resource}
	resource := s.storage.FindResource(gr)
	if resource == nil {
		return req.writeErrorResponse(http.StatusNotFound)
	}

	id := types.NamespacedName{Namespace: req.Namespace, Name: req.Name}

	existingObj, found, err := s.storage.GetObject(ctx, resource, id)
	if err != nil {
		return err
	}
	if !found {
		return req.writeErrorResponse(http.StatusNotFound)
	}

	bodyBytes, err := io.ReadAll(req.r.Body)
	if err != nil {
		return err
	}

	klog.Infof("put request %#v", string(bodyBytes))

	body := &unstructured.Unstructured{}
	if err := body.UnmarshalJSON(bodyBytes); err != nil {
		return fmt.Errorf("failed to parse payload: %w", err)
	}

	original := existingObj.DeepCopy()

	var updated *unstructured.Unstructured

	if req.SubResource == "" {
		updated = body
	} else if req.SubResource == "status" {
		updated = existingObj.DeepCopyObject().(*unstructured.Unstructured)
		newStatus := body.Object["status"]
		if newStatus == nil {
			// TODO: This might be allowed?
			return fmt.Errorf("status not specified on status subresource update")
		}
		updated.Object["status"] = newStatus
	} else {
		// TODO: We need to implement put properly
		return fmt.Errorf("unknown subresource %q", req.SubResource)
	}

	// We don't want to change the resourceVersion (and trigger watches) when the change is a no-op
	if reflect.DeepEqual(original, updated) {
		klog.Infof("update did not change object")
		return req.writeResponse(original)
	}

	if err := s.storage.UpdateObject(ctx, resource, id, updated); err != nil {
		return err
	}
	return req.writeResponse(updated)
}
