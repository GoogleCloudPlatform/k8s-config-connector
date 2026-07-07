// Copyright 2026 Google LLC
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

package clouderrorreporting

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouderrorreporting/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ErrorReportingNotificationSettingsGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

type model struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &model{}

type GCPNotificationSettings struct {
	Name                 string   `json:"name,omitempty"`
	NotificationChannels []string `json:"notificationChannels,omitempty"`
}

type adapter struct {
	id         *krm.ErrorReportingNotificationSettingsIdentity
	httpClient *http.Client
	desired    *GCPNotificationSettings
	actual     *GCPNotificationSettings
}

var _ directbase.Adapter = &adapter{}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building http client: %w", err)
	}

	obj := &krm.ErrorReportingNotificationSettings{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	var resolvedChannels []string
	for _, ref := range obj.Spec.NotificationChannels {
		resolved, err := refs.ResolveMonitoringNotificationChannelRef(ctx, reader, obj, &ref)
		if err != nil {
			return nil, fmt.Errorf("resolving notificationChannelRef: %w", err)
		}
		resolvedChannels = append(resolvedChannels, resolved.String())
	}

	desired := &GCPNotificationSettings{
		Name:                 fmt.Sprintf("projects/%s/locations/global/notificationSettings", id.(*krm.ErrorReportingNotificationSettingsIdentity).Project),
		NotificationChannels: resolvedChannels,
	}

	return &adapter{
		id:         id.(*krm.ErrorReportingNotificationSettingsIdentity),
		httpClient: httpClient,
		desired:    desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//clouderrorreporting.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//clouderrorreporting.googleapis.com/")

	id := &krm.ErrorReportingNotificationSettingsIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building http client: %w", err)
	}

	return &adapter{
		id:         id,
		httpClient: httpClient,
	}, nil
}

func (a *adapter) makeRequest(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return a.httpClient.Do(req)
}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	url := fmt.Sprintf("https://clouderrorreporting.googleapis.com/v1beta1/projects/%s/locations/global/notificationSettings", a.id.Project)
	resp, err := a.makeRequest(ctx, "GET", url, nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("getting notification settings: status %s, body: %s", resp.Status, string(body))
	}

	var actual GCPNotificationSettings
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		return false, err
	}
	a.actual = &actual
	return true, nil
}

func (a *adapter) patch(ctx context.Context, channels []string) error {
	url := fmt.Sprintf("https://clouderrorreporting.googleapis.com/v1beta1/projects/%s/locations/global/notificationSettings?updateMask=notificationChannels", a.id.Project)
	payload := &GCPNotificationSettings{
		Name:                 fmt.Sprintf("projects/%s/locations/global/notificationSettings", a.id.Project),
		NotificationChannels: channels,
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := a.makeRequest(ctx, "PATCH", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("patching notification settings: status %s, body: %s", resp.Status, string(body))
	}
	return nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	if err := a.patch(ctx, a.desired.NotificationChannels); err != nil {
		return err
	}
	return nil
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	if reflect.DeepEqual(a.desired.NotificationChannels, a.actual.NotificationChannels) {
		return nil
	}
	if err := a.patch(ctx, a.desired.NotificationChannels); err != nil {
		return err
	}
	return nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting object")
	if err := a.patch(ctx, nil); err != nil {
		return false, err
	}
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("no actual state found to export")
	}

	var refsList []interface{}
	for _, ch := range a.actual.NotificationChannels {
		ref := map[string]interface{}{
			"external": ch,
		}
		refsList = append(refsList, ref)
	}

	spec := map[string]interface{}{
		"notificationChannels": refsList,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.ErrorReportingNotificationSettingsGVK)
	u.SetName(a.id.Project + "-notification-settings") // Set a deterministic name
	u.SetNamespace(a.id.Project)
	if err := unstructured.SetNestedMap(u.Object, spec, "spec"); err != nil {
		return nil, err
	}
	return u, nil
}
