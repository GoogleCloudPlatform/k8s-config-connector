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

package managedkafka

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/managedkafka/apiv1"
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ManagedKafkaConnectClusterGVK, NewConnectClusterModel)
}

func NewConnectClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConnectCluster{config: *config}, nil
}

var _ directbase.Model = &modelConnectCluster{}

type modelConnectCluster struct {
	config config.ControllerConfig
}

func (m *modelConnectCluster) client(ctx context.Context) (*gcp.ManagedKafkaConnectClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewManagedKafkaConnectRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ConnectCluster client: %w", err)
	}
	return gcpClient, err
}

func (m *modelConnectCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ManagedKafkaConnectCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewConnectClusterIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConnectClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelConnectCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ConnectClusterAdapter struct {
	id        *krm.ManagedKafkaConnectClusterIdentity
	gcpClient *gcp.ManagedKafkaConnectClient
	desired   *krm.ManagedKafkaConnectCluster
	actual    *pb.ConnectCluster
	reader    client.Reader
}

var _ directbase.Adapter = &ConnectClusterAdapter{}

// Find retrieves the GCP resource.
func (a *ConnectClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConnectCluster", "name", a.id)

	req := &pb.GetConnectClusterRequest{Name: a.id.String()}
	connectClusterPb, err := a.gcpClient.GetConnectCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConnectCluster %q: %w", a.id, err)
	}

	a.actual = connectClusterPb
	return true, nil
}

// Create creates the resource in GCP.
func (a *ConnectClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConnectCluster", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := ManagedKafkaConnectClusterSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateConnectClusterRequest{
		Parent:           a.id.ParentString(),
		ConnectClusterId: a.id.ID(),
		ConnectCluster:   resource,
	}
	op, err := a.gcpClient.CreateConnectCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConnectCluster %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of ConnectCluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ConnectCluster", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

// Update updates the resource in GCP.
func (a *ConnectClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ConnectCluster", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := ManagedKafkaConnectClusterSpec_v1alpha1_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Name = a.id.String()

	diffs, updateMask, err := compareConnectCluster(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateConnectClusterRequest{
		UpdateMask:     updateMask,
		ConnectCluster: desiredPb,
	}
	op, err := a.gcpClient.UpdateConnectCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ConnectCluster %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting update for ConnectCluster %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated ConnectCluster", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, updated)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConnectClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ManagedKafkaConnectCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ManagedKafkaConnectClusterSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ManagedKafkaConnectClusterGVK)

	u.Object = uObj
	return u, nil
}

// Delete deletes the resource from GCP.
func (a *ConnectClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConnectCluster", "name", a.id)

	req := &pb.DeleteConnectClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConnectCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ConnectCluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ConnectCluster", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ConnectCluster %s: %w", a.id, err)
	}
	return true, nil
}

func (a *ConnectClusterAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ConnectCluster) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ManagedKafkaConnectClusterStatus{}
	status.ObservedState = ManagedKafkaConnectClusterObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

func compareConnectCluster(ctx context.Context, actual, desired *pb.ConnectCluster) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := ManagedKafkaConnectClusterSpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := ManagedKafkaConnectClusterSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.ConnectCluster)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
