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

package bigqueryanalyticshub

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/analyticshub/apiv1"
	bigqueryanalyticshubpb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "bigqueryanalyticshub-dataexchange-controller"
)

func init() {
	registry.RegisterModel(krm.BigQueryAnalyticsHubDataExchangeGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataExchange client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryAnalyticsHubDataExchange{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get bigqueryanalyticshub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id.(*krm.BigQueryAnalyticsHubDataExchangeIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.BigQueryAnalyticsHubDataExchangeIdentity
	gcpClient *gcp.Client
	desired   *krm.BigQueryAnalyticsHubDataExchange
	actual    *bigqueryanalyticshubpb.DataExchange
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataExchange", "name", a.id.String())

	req := &bigqueryanalyticshubpb.GetDataExchangeRequest{Name: a.id.String()}
	dataexchangepb, err := a.gcpClient.GetDataExchange(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataExchange %q: %w", a.id.String(), err)
	}

	a.actual = dataexchangepb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataExchange", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigQueryAnalyticsHubDataExchangeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &bigqueryanalyticshubpb.CreateDataExchangeRequest{
		Parent:         parent,
		DataExchangeId: a.id.DataExchange,
		DataExchange:   resource,
	}
	created, err := a.gcpClient.CreateDataExchange(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataExchange %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DataExchange", "name", a.id.String())

	status := &krm.BigQueryAnalyticsHubDataExchangeStatus{}
	status.ObservedState = BigQueryAnalyticsHubDataExchangeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataExchange", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// TODO(kcc): Autogen "func immutable()" for each field
	// TODO(kcc): autogen updateMastk.path for mutable gcp fields.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(a.desired.Spec.PrimaryContact, a.actual.PrimaryContact) {
		updateMask.Paths = append(updateMask.Paths, "primary_contact")
	}
	if !reflect.DeepEqual(a.desired.Spec.Documentation, a.actual.Documentation) {
		updateMask.Paths = append(updateMask.Paths, "documentation")
	}
	// not yet
	// if !reflect.DeepEqual(a.desired.Spec.Icon, string(a.actual.Icon)) {
	// 	updateMask.Paths = append(updateMask.Paths, "icon")
	// }
	if a.desired.Spec.DiscoveryType != nil && !reflect.DeepEqual(a.desired.Spec.DiscoveryType, a.actual.DiscoveryType.String()) {
		updateMask.Paths = append(updateMask.Paths, "discovery_type")
	}

	desired := a.desired.DeepCopy()
	resource := BigQueryAnalyticsHubDataExchangeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.actual.Name

	req := &bigqueryanalyticshubpb.UpdateDataExchangeRequest{
		UpdateMask:   updateMask,
		DataExchange: resource,
	}

	updated, err := a.gcpClient.UpdateDataExchange(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DataExchange %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated DataExchange", "name", a.id.String())

	status := &krm.BigQueryAnalyticsHubDataExchangeStatus{}
	status.ObservedState = BigQueryAnalyticsHubDataExchangeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryAnalyticsHubDataExchange{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryAnalyticsHubDataExchangeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataExchange", "name", a.id.String())

	req := &bigqueryanalyticshubpb.DeleteDataExchangeRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteDataExchange(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting DataExchange %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DataExchange", "name", a.id.String())

	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
