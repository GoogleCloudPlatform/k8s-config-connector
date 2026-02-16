// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkservices

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesEdgeCacheServiceGVK, NewEdgeCacheServiceModel)
}

func NewEdgeCacheServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEdgeCacheService{config: *config}, nil
}

var _ directbase.Model = &modelEdgeCacheService{}

type modelEdgeCacheService struct {
	config config.ControllerConfig
}

func (m *modelEdgeCacheService) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesEdgeCacheService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEdgeCacheServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting HTTP client: %w", err)
	}

	return &EdgeCacheServiceAdapter{
		id:         id,
		httpClient: httpClient,
		desired:    obj,
	}, nil
}

func (m *modelEdgeCacheService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type EdgeCacheServiceAdapter struct {
	id         *krm.EdgeCacheServiceIdentity
	httpClient *http.Client
	desired    *krm.NetworkServicesEdgeCacheService
	actual     *EdgeCacheService
}

var _ directbase.Adapter = &EdgeCacheServiceAdapter{}

func (a *EdgeCacheServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EdgeCacheService", "name", a.id)

	url := fmt.Sprintf("https://networkservices.googleapis.com/v1/%s", a.id.String())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("building GET request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("executing GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("GET request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var actual EdgeCacheService
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		return false, fmt.Errorf("decoding GET response: %w", err)
	}

	a.actual = &actual
	return true, nil
}

func (a *EdgeCacheServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EdgeCacheService", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := EdgeCacheServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if resource.Labels == nil {
		resource.Labels = make(map[string]string)
	}
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	url := fmt.Sprintf("https://networkservices.googleapis.com/v1/%s/edgeCacheServices?edgeCacheServiceId=%s", a.id.Parent().String(), a.id.ID())
	body, err := json.Marshal(resource)
	if err != nil {
		return fmt.Errorf("marshalling resource: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("building POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("POST request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var op Operation
	if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
		return fmt.Errorf("decoding POST response: %w", err)
	}

	created, err := a.waitForOperation(ctx, &op)
	if err != nil {
		return fmt.Errorf("waiting for creation operation: %w", err)
	}

	status := EdgeCacheServiceStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EdgeCacheServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EdgeCacheService", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := EdgeCacheServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if resource.Labels == nil {
		resource.Labels = make(map[string]string)
	}
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	url := fmt.Sprintf("https://networkservices.googleapis.com/v1/%s", a.id.String())
	body, err := json.Marshal(resource)
	if err != nil {
		return fmt.Errorf("marshalling resource: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("building PATCH request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing PATCH request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("PATCH request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var op Operation
	if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
		return fmt.Errorf("decoding PATCH response: %w", err)
	}

	updated, err := a.waitForOperation(ctx, &op)
	if err != nil {
		return fmt.Errorf("waiting for update operation: %w", err)
	}

	status := EdgeCacheServiceStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *EdgeCacheServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *EdgeCacheServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EdgeCacheService", "name", a.id)

	url := fmt.Sprintf("https://networkservices.googleapis.com/v1/%s", a.id.String())
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return false, fmt.Errorf("building DELETE request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("executing DELETE request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("DELETE request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var op Operation
	if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
		return false, fmt.Errorf("decoding DELETE response: %w", err)
	}

	_, err = a.waitForOperation(ctx, &op)
	if err != nil {
		return false, fmt.Errorf("waiting for delete operation: %w", err)
	}

	return true, nil
}

func (a *EdgeCacheServiceAdapter) waitForOperation(ctx context.Context, op *Operation) (*EdgeCacheService, error) {
	for {
		if op.Done {
			if op.Error != nil {
				return nil, fmt.Errorf("operation failed: %s", op.Error.Message)
			}
			var created EdgeCacheService
			if err := json.Unmarshal(op.Response, &created); err != nil {
				// If response is empty (e.g. for delete), this is expected.
				return &created, nil
			}
			return &created, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(5 * time.Second):
			url := fmt.Sprintf("https://networkservices.googleapis.com/v1/%s", op.Name)
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return nil, fmt.Errorf("building operation GET request: %w", err)
			}
			resp, err := a.httpClient.Do(req)
			if err != nil {
				return nil, fmt.Errorf("executing operation GET request: %w", err)
			}
			defer resp.Body.Close()
			if err := json.NewDecoder(resp.Body).Decode(op); err != nil {
				return nil, fmt.Errorf("decoding operation GET response: %w", err)
			}
		}
	}
}

type Operation struct {
	Name     string          `json:"name"`
	Done     bool            `json:"done"`
	Error    *OperationError `json:"error,omitempty"`
	Response json.RawMessage `json:"response,omitempty"`
}

type OperationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
