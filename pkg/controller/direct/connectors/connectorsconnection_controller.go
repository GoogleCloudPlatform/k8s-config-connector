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

package connectors

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/connectors/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ConnectorsConnectionGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConnectorsConnection{config: config}, nil
}

var _ directbase.Model = &modelConnectorsConnection{}

type modelConnectorsConnection struct {
	config *config.ControllerConfig
}

type gcpClient struct {
	config  config.ControllerConfig
	service *api.Service
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: *config,
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient.service, err = api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building gcp service client: %w", err)
	}

	return gcpClient, nil
}

func (m *gcpClient) connectionsClient() *api.ProjectsLocationsConnectionsService {
	return api.NewProjectsLocationsConnectionsService(m.service)
}

func (m *gcpClient) operationsClient() *api.ProjectsLocationsOperationsService {
	return api.NewProjectsLocationsOperationsService(m.service)
}

func (m *modelConnectorsConnection) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConnectorsConnection{}

	copied := u.DeepCopy()
	if err := label.ComputeLabels(copied); err != nil {
		return nil, err
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(copied.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ConnectorsConnectionIdentity)

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	return &ConnectorsConnectionAdapter{
		id:                id,
		k8sClient:         reader,
		connectionsClient: gcpClient.connectionsClient(),
		operationsClient:  gcpClient.operationsClient(),
		desired:           obj,
	}, nil
}

func (m *modelConnectorsConnection) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ConnectorsConnectionAdapter struct {
	id                *krm.ConnectorsConnectionIdentity
	k8sClient         client.Reader
	connectionsClient *api.ProjectsLocationsConnectionsService
	operationsClient  *api.ProjectsLocationsOperationsService
	desired           *krm.ConnectorsConnection
	actual            *api.Connection
}

var _ directbase.Adapter = &ConnectorsConnectionAdapter{}

func WaitForConnectorsOp(ctx context.Context, client *api.ProjectsLocationsOperationsService, op *api.Operation) error {
	return common.WaitForDoneOrTimeout(ctx, 2*time.Second, func() (bool, error) {
		current, err := client.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return false, fmt.Errorf("getting operation status of %q: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return true, fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error.Message)
			} else {
				return true, nil
			}
		}
		return false, nil
	})
}

func (a *ConnectorsConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConnectorsConnection", "name", a.id)

	found, err := a.connectionsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConnectorsConnection %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

func (a *ConnectorsConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConnectorsConnection", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := ConnectorsConnectionSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Propagate metadata labels to GCP Connection labels
	if resource.Labels == nil {
		resource.Labels = make(map[string]string)
	}
	for k, v := range a.desired.GetLabels() {
		resource.Labels[k] = v
	}

	op, err := a.connectionsClient.Create(a.id.ParentString(), resource).ConnectionId(a.id.Connection).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ConnectorsConnection %s: %w", a.id, err)
	}
	if err := WaitForConnectorsOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ConnectorsConnection %s waiting creation: %w", a.id, err)
	}

	created, err := a.connectionsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ConnectorsConnection: %w", err)
	}

	log.V(2).Info("successfully created ConnectorsConnection", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ConnectorsConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ConnectorsConnection", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := ConnectorsConnectionSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Propagate metadata labels to GCP Connection labels
	if resource.Labels == nil {
		resource.Labels = make(map[string]string)
	}
	for k, v := range a.desired.GetLabels() {
		resource.Labels[k] = v
	}

	var updateMask []string
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, resource.Labels)
		updateMask = append(updateMask, "labels")
	}
	if resource.Description != a.actual.Description {
		report.AddField("description", a.actual.Description, resource.Description)
		updateMask = append(updateMask, "description")
	}
	if resource.ConnectorVersion != a.actual.ConnectorVersion {
		report.AddField("connectorVersion", a.actual.ConnectorVersion, resource.ConnectorVersion)
		updateMask = append(updateMask, "connector_version")
	}
	if resource.ServiceAccount != a.actual.ServiceAccount {
		report.AddField("serviceAccount", a.actual.ServiceAccount, resource.ServiceAccount)
		updateMask = append(updateMask, "service_account")
	}
	if !reflect.DeepEqual(resource.ConfigVariables, a.actual.ConfigVariables) {
		report.AddField("configVariables", a.actual.ConfigVariables, resource.ConfigVariables)
		updateMask = append(updateMask, "config_variables")
	}
	if !reflect.DeepEqual(resource.AuthConfig, a.actual.AuthConfig) {
		report.AddField("authConfig", a.actual.AuthConfig, resource.AuthConfig)
		updateMask = append(updateMask, "auth_config")
	}
	if !reflect.DeepEqual(resource.LockConfig, a.actual.LockConfig) {
		report.AddField("lockConfig", a.actual.LockConfig, resource.LockConfig)
		updateMask = append(updateMask, "lock_config")
	}
	if !reflect.DeepEqual(resource.DestinationConfigs, a.actual.DestinationConfigs) {
		report.AddField("destinationConfigs", a.actual.DestinationConfigs, resource.DestinationConfigs)
		updateMask = append(updateMask, "destination_configs")
	}
	if resource.Suspended != a.actual.Suspended {
		report.AddField("suspended", a.actual.Suspended, resource.Suspended)
		updateMask = append(updateMask, "suspended")
	}
	if !reflect.DeepEqual(resource.NodeConfig, a.actual.NodeConfig) {
		report.AddField("nodeConfig", a.actual.NodeConfig, resource.NodeConfig)
		updateMask = append(updateMask, "node_config")
	}
	if !reflect.DeepEqual(resource.SslConfig, a.actual.SslConfig) {
		report.AddField("sslConfig", a.actual.SslConfig, resource.SslConfig)
		updateMask = append(updateMask, "ssl_config")
	}

	if len(updateMask) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, report)

	op, err := a.connectionsClient.Patch(a.id.String(), resource).UpdateMask(strings.Join(updateMask, ",")).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ConnectorsConnection %s: %w", a.id, err)
	}
	if err := WaitForConnectorsOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ConnectorsConnection %s waiting update: %w", a.id, err)
	}

	updated, err := a.connectionsClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated ConnectorsConnection: %w", err)
	}

	log.V(2).Info("successfully updated ConnectorsConnection", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ConnectorsConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConnectorsConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConnectorsConnectionSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Connection

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Connection)
	u.SetGroupVersionKind(krm.ConnectorsConnectionGVK)

	u.Object = uObj
	return u, nil
}

func (a *ConnectorsConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConnectorsConnection", "name", a.id)

	op, err := a.connectionsClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ConnectorsConnection, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ConnectorsConnection %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ConnectorsConnection", "name", a.id)

	if err := WaitForConnectorsOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting delete ConnectorsConnection %s: %w", a.id, err)
	}
	return true, nil
}

func (a *ConnectorsConnectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.Connection) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ConnectorsConnectionStatus{}
	status.ObservedState = ConnectorsConnectionObservedState_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
