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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouderrorreporting/v1beta1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	httransport "google.golang.org/api/transport/http"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ErrorReportingNotificationSettingsGVK, NewErrorReportingNotificationSettingsModel)
}

func NewErrorReportingNotificationSettingsModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelErrorReportingNotificationSettings{config: *config}, nil
}

var _ directbase.Model = &modelErrorReportingNotificationSettings{}

type modelErrorReportingNotificationSettings struct {
	config config.ControllerConfig
}

type NotificationSettings struct {
	Name                    string   `json:"name,omitempty"`
	NotificationChannels    []string `json:"notificationChannels,omitempty"`
	VersionSkewReportDelays []string `json:"versionSkewReportDelays,omitempty"`
}

type gcpClient struct {
	httpClient *http.Client
}

func (m *modelErrorReportingNotificationSettings) client(ctx context.Context) (*gcpClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	httpClient, _, err := httransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building clouderrorreporting http client: %w", err)
	}
	return &gcpClient{httpClient: httpClient}, nil
}

func (m *modelErrorReportingNotificationSettings) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ErrorReportingNotificationSettings{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ErrorReportingNotificationSettingsAdapter{
		id:        id.(*krm.ErrorReportingNotificationSettingsIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelErrorReportingNotificationSettings) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ErrorReportingNotificationSettingsAdapter struct {
	id        *krm.ErrorReportingNotificationSettingsIdentity
	gcpClient *gcpClient
	desired   *krm.ErrorReportingNotificationSettings
	actual    *NotificationSettings
}

var _ directbase.Adapter = &ErrorReportingNotificationSettingsAdapter{}

func (a *ErrorReportingNotificationSettingsAdapter) Find(ctx context.Context) (bool, error) {
	url := fmt.Sprintf("https://clouderrorreporting.googleapis.com/v1beta1/%s", a.id.String())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, err
	}
	resp, err := a.gcpClient.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("getting ErrorReportingNotificationSettings %s: %s", a.id.String(), string(body))
	}

	var actual NotificationSettings
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		return false, fmt.Errorf("decoding ErrorReportingNotificationSettings %s: %w", a.id.String(), err)
	}
	a.actual = &actual
	return true, nil
}

func (a *ErrorReportingNotificationSettingsAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// For singleton settings, Create is same as Update
	return a.Update(ctx, &directbase.UpdateOperation{})
}

func (a *ErrorReportingNotificationSettingsAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ErrorReportingNotificationSettings", "name", a.id.String())

	desiredPb := ErrorReportingNotificationSettingsSpec_ToProto(&a.desired.Spec)
	desiredPb.Name = a.id.String()

	// Calculate update mask
	var paths []string
	if a.actual == nil || !reflect.DeepEqual(desiredPb.NotificationChannels, a.actual.NotificationChannels) {
		paths = append(paths, "notification_channels")
	}
	if a.actual == nil || !reflect.DeepEqual(desiredPb.VersionSkewReportDelays, a.actual.VersionSkewReportDelays) {
		paths = append(paths, "version_skew_report_delays")
	}

	if len(paths) == 0 {
		return nil
	}

	report := &structuredreporting.Diff{}
	for _, path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	url := fmt.Sprintf("https://clouderrorreporting.googleapis.com/v1beta1/%s?updateMask=%s", a.id.String(), strings.Join(paths, ","))
	body, err := json.Marshal(desiredPb)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.gcpClient.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("updating ErrorReportingNotificationSettings %s: %s", a.id.String(), string(respBody))
	}

	var latest NotificationSettings
	if err := json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		return fmt.Errorf("decoding ErrorReportingNotificationSettings %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, &latest)
}

func (a *ErrorReportingNotificationSettingsAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting (resetting) ErrorReportingNotificationSettings", "name", a.id.String())

	emptySettings := &NotificationSettings{
		Name: a.id.String(),
	}
	url := fmt.Sprintf("https://clouderrorreporting.googleapis.com/v1beta1/%s", a.id.String())
	body, err := json.Marshal(emptySettings)
	if err != nil {
		return false, err
	}
	// Omit updateMask to reset all fields
	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.gcpClient.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("deleting (resetting) ErrorReportingNotificationSettings %s: %s", a.id.String(), string(respBody))
	}

	return true, nil
}

func (a *ErrorReportingNotificationSettingsAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ErrorReportingNotificationSettings{}
	obj.Spec = *ErrorReportingNotificationSettingsSpec_FromProto(a.actual)

	// We need to set the projectRef in spec for Export to work correctly in some tests
	obj.Spec.ProjectRef = refsv1beta1.ProjectRef{
		External: "projects/" + a.id.Project,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ErrorReportingNotificationSettingsGVK)

	u.Object = uObj
	return u, nil
}

func (a *ErrorReportingNotificationSettingsAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *NotificationSettings) error {
	status := &krm.ErrorReportingNotificationSettingsStatus{}
	// No status fields to map currently, other than observedGeneration
	return op.UpdateStatus(ctx, status, nil)
}

func ErrorReportingNotificationSettingsSpec_ToProto(in *krm.ErrorReportingNotificationSettingsSpec) *NotificationSettings {
	if in == nil {
		return nil
	}
	out := &NotificationSettings{}
	for _, ref := range in.NotificationChannels {
		out.NotificationChannels = append(out.NotificationChannels, ref.External)
	}
	out.VersionSkewReportDelays = in.VersionSkewReportDelays
	return out
}

func ErrorReportingNotificationSettingsSpec_FromProto(in *NotificationSettings) *krm.ErrorReportingNotificationSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.ErrorReportingNotificationSettingsSpec{}
	for _, channel := range in.NotificationChannels {
		out.NotificationChannels = append(out.NotificationChannels, monitoringv1beta1.MonitoringNotificationChannelRef{External: channel})
	}
	out.VersionSkewReportDelays = in.VersionSkewReportDelays
	return out
}
