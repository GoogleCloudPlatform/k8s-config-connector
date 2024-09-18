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

package bigqueryconnection

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/google/uuid"

	gcp "cloud.google.com/go/bigquery/connection/apiv1"

	bigqueryconnectionpb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "bigqueryconnection-controller"
	serviceDomain = "//bigqueryconnection.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.BigQueryConnectionConnectionGVK, NewModel)
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
		return nil, fmt.Errorf("building bigqueryconnection client: %w", err)
	}
	return gcpClient, err
}

func isValidUUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryConnectionConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	connectionRef, err := krm.New(ctx, reader, u)
	if err != nil {
		return nil, err
	}

	// Get bigqueryconnection GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        connectionRef,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.BigQueryConnectionConnectionRef
	gcpClient *gcp.Client
	desired   *krm.BigQueryConnectionConnection
	actual    *bigqueryconnectionpb.Connection
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)

	log.V(2).Info("getting BigQueryConnectionConnection", "name", a.id.External)

	if a.id.External == "" {
		// Cannot retrieve the Connection without ServiceGeneratedID, expecting to create a new Connection.
		return false, nil
	}
	req := &bigqueryconnectionpb.GetConnectionRequest{Name: a.id.External}
	connectionpb, err := a.gcpClient.GetConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryConnectionConnection %q: %w", a.id.External, err)
	}

	a.actual = connectionpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Connection", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &bigqueryconnectionpb.CreateConnectionRequest{
		Parent:     a.id.Parent(),
		Connection: resource,
	}
	created, err := a.gcpClient.CreateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Connection %s: %w", *&a.id.External, err)
	}
	log.V(2).Info("successfully created Connection", "name", created.Name)

	status := &krm.BigQueryConnectionConnectionStatus{}
	status.ObservedState = BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &created.Name
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Connection", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.FriendlyName, a.actual.FriendlyName) {
		updateMask.Paths = append(updateMask.Paths, "friendly_name")
	}
	if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	desired := a.desired.DeepCopy()
	resource := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	fqn := a.id.External
	req := &bigqueryconnectionpb.UpdateConnectionRequest{
		Name:       fqn,
		Connection: resource,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully updated Connection", "name", fqn)

	status := &krm.BigQueryConnectionConnectionStatus{}
	status.ObservedState = BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx, updated)
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

	obj := &krm.BigQueryConnectionConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryConnectionConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.Parent()
	if parent != "" {
		tokens := strings.Split(parent, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			obj.Spec.ProjectRef = &refs.ProjectRef{Name: tokens[1]}
			obj.Spec.Location = tokens[3]
		}
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting Connection", "name", a.id.External)

	fqn := a.id.External
	req := &bigqueryconnectionpb.DeleteConnectionRequest{Name: fqn}
	if err := a.gcpClient.DeleteConnection(ctx, req); err != nil {
		return false, fmt.Errorf("deleting Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted Connection", "name", fqn)
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
		status["serviceGeneratedID"] = old["serviceGeneratedID"]
	}

	u.Object["status"] = status

	return nil
}
