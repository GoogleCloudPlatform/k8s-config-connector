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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

// patchResource is a request to patch a single resource
type patchResource struct {
	resourceRequestBase
}

// Run serves the http request
func (req *patchResource) Run(ctx context.Context, s *MockKubeAPIServer) error {
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
		existingObj = nil
	}

	bodyBytes, err := io.ReadAll(req.r.Body)
	if err != nil {
		return err
	}

	body := &unstructured.Unstructured{}
	// Can't use the MarshalJSON overload, it doesn't like missing kind etc
	if err := json.Unmarshal(bodyBytes, &body.Object); err != nil {
		return fmt.Errorf("failed to parse PATCH payload: %w", err)
	}

	// TODO: We need to implement patch properly
	klog.Infof("patch request %#v", string(bodyBytes))

	if !found {
		// TODO: Only if server-side-apply

		if req.SubResource != "" {
			// TODO: Is this correct for server-side-apply?
			return req.writeErrorResponse(http.StatusNotFound)
		}

		// TODO: Should we treat this like an apply to an empty object?

		patched := body
		if err := s.storage.CreateObject(ctx, resource, id, patched); err != nil {
			return err
		}

		return req.writeResponse(patched)
	}

	var updated *unstructured.Unstructured
	changed := true
	if req.SubResource == "" {
		if resource.TypeInfo != nil {
			patchOptions := metav1.PatchOptions{}
			if fieldManager := req.r.URL.Query().Get("fieldManager"); fieldManager != "" {
				patchOptions.FieldManager = fieldManager
			}
			if force := req.r.URL.Query().Get("force"); force != "" {
				forceBool, err := strconv.ParseBool(force)
				if err != nil {
					return fmt.Errorf("invalid value %q for force", force)
				}
				patchOptions.Force = &forceBool
			}
			updated, changed, err = resource.DoServerSideApply(ctx, existingObj, bodyBytes, patchOptions)
			if err != nil {
				klog.Warningf("error from DoServerSideApply: %v", err)
				return err
			}

		} else {
			klog.Warningf("falling back to untyped apply emulation")
			updated = existingObj.DeepCopy()
			if err := applyPatch(updated.Object, body.Object); err != nil {
				klog.Warningf("error from patch: %v", err)
				return err
			}
		}
	} else {
		// TODO: We need to implement put properly
		return fmt.Errorf("unknown subresource %q", req.SubResource)
	}

	if !changed {
		klog.Infof("skipping write, object not changed")
		return req.writeResponse(existingObj)
	} else {
		if err := s.storage.UpdateObject(ctx, resource, id, updated); err != nil {
			return err
		}
		return req.writeResponse(updated)
	}
}

func applyPatch(existing, patch map[string]interface{}) error {
	for k, patchValue := range patch {
		existingValue := existing[k]
		switch patchValue := patchValue.(type) {
		case string, int64, float64:
			existing[k] = patchValue
		case map[string]interface{}:
			if existingValue == nil {
				existing[k] = patchValue
			} else {
				existingMap, ok := existingValue.(map[string]interface{})
				if !ok {
					return fmt.Errorf("unexpected type mismatch, expected map got %T", existingValue)
				}
				if err := applyPatch(existingMap, patchValue); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("type %T not handled in patch", patchValue)
		}
	}
	return nil
}
