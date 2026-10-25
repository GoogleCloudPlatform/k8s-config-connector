// Copyright 2025 Google LLC
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

// +generated:controller
// krm.group: managedkafka.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ManagedKafkaConnector

package managedkafka

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/managedkafka/apiv1"
	managedkafkapb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ManagedKafkaConnectorGVK, NewConnectorModel)
}

func NewConnectorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConnector{config: *config}, nil
}

var _ directbase.Model = &modelConnector{}

type modelConnector struct {
	config config.ControllerConfig
}

func (m *modelConnector) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ManagedKafkaConnector{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ManagedKafkaConnectorIdentity)

	// Get managedkafka GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConnectorAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelConnector) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//managedkafka.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ManagedKafkaConnectorIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConnectorAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelConnector) client(ctx context.Context) (*gcp.ManagedKafkaConnectClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewManagedKafkaConnectRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Connector client: %w", err)
	}
	return gcpClient, err
}

type ConnectorAdapter struct {
	id        *krm.ManagedKafkaConnectorIdentity
	gcpClient *gcp.ManagedKafkaConnectClient
	desired   *krm.ManagedKafkaConnector
	actual    *managedkafkapb.Connector
}

var _ directbase.Adapter = &ConnectorAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ConnectorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Connector", "name", a.id)

	req := &managedkafkapb.GetConnectorRequest{
		Name: a.id.String(),
	}
	connectorpb, err := a.gcpClient.GetConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Connector %q: %w", a.id, err)
	}

	a.actual = connectorpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Connector", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ManagedKafkaConnectorSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &managedkafkapb.CreateConnectorRequest{
		Parent:      a.id.ParentString(),
		ConnectorId: a.id.Connector,
		Connector:   resource,
	}
	created, err := a.gcpClient.CreateConnector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Connector %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Connector", "name", a.id)

	status := &krm.ManagedKafkaConnectorStatus{}
	status.ObservedState = ManagedKafkaConnectorObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Connector", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ManagedKafkaConnectorSpec_v1alpha1_FromProto,
		ManagedKafkaConnectorSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &managedkafkapb.UpdateConnectorRequest{
			UpdateMask: updateMask,
			Connector:  resource,
		}
		updated, err = a.gcpClient.UpdateConnector(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Connector %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Connector", "name", a.id)
	}

	status := &krm.ManagedKafkaConnectorStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ManagedKafkaConnectorObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConnectorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ManagedKafkaConnector{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ManagedKafkaConnectorSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ManagedKafkaConnectorGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ConnectorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Connector", "name", a.id)

	req := &managedkafkapb.DeleteConnectorRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Connector, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Connector %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Connector", "name", a.id)
	return true, nil
}
