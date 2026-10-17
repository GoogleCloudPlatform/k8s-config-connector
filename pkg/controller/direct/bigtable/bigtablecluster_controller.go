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

package bigtable

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	admin "cloud.google.com/go/bigtable/admin/apiv2"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigtableClusterGVK, NewClusterModel)
}

func NewClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCluster{config: *config}, nil
}

var _ directbase.Model = &modelCluster{}

type modelCluster struct {
	config config.ControllerConfig
}

func (m *modelCluster) client(ctx context.Context) (*admin.BigtableInstanceAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableCluster client options: %w", err)
	}
	gcpClient, err := admin.NewBigtableInstanceAdminClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableCluster client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewClusterIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ClusterAdapter struct {
	id        *krm.ClusterIdentity
	gcpClient *admin.BigtableInstanceAdminClient
	desired   *krm.BigtableCluster
	actual    *pb.Cluster
	actualML  *pb.MemoryLayer
}

var _ directbase.Adapter = &ClusterAdapter{}

func (a *ClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableCluster", "name", a.id)

	cluster, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableCluster %q: %w", a.id, err)
	}
	a.actual = cluster

	// Also fetch the memory layer state.
	mlName := a.id.String() + "/memoryLayer"
	ml, err := a.gcpClient.GetMemoryLayer(ctx, &pb.GetMemoryLayerRequest{Name: mlName})
	if err != nil {
		if !direct.IsNotFound(err) {
			return false, fmt.Errorf("getting MemoryLayer for BigtableCluster %q: %w", a.id, err)
		}
	} else {
		a.actualML = ml
	}

	return true, nil
}

func (a *ClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableCluster", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := BigtableClusterSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	op, err := a.gcpClient.CreateCluster(ctx, &pb.CreateClusterRequest{
		Parent:    a.id.ParentString(),
		ClusterId: a.id.ID(),
		Cluster:   desired,
	})
	if err != nil {
		return fmt.Errorf("creating BigtableCluster %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for BigtableCluster %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableCluster", "name", a.id)
	a.actual = created

	if err := a.reconcileMemoryLayer(ctx); err != nil {
		return err
	}

	return a.setStatus(ctx, func(status *krm.BigtableClusterStatus) error {
		return createOp.UpdateStatus(ctx, status, nil)
	})
}

func (a *ClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableCluster", "name", a.id)

	spec := a.desired.Spec
	updateMask := &fieldmaskpb.FieldMask{}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	if direct.ValueOf(spec.ServeNodes) != a.actual.GetServeNodes() &&
		spec.ClusterConfig == nil {
		report.AddField("serve_nodes", a.actual.GetServeNodes(), spec.ServeNodes)
		updateMask.Paths = append(updateMask.Paths, "serve_nodes")
	}

	if !reflect.DeepEqual(spec.ClusterConfig, BigtableClusterSpec_v1alpha1_FromProto(&direct.MapContext{}, a.actual).ClusterConfig) {
		report.AddField("cluster_config.cluster_autoscaling_config", a.actual.GetClusterConfig(), spec.ClusterConfig)
		updateMask.Paths = append(updateMask.Paths, "cluster_config.cluster_autoscaling_config")
	}

	if len(updateMask.Paths) > 0 {
		log.V(2).Info("fields need update", "name", a.id, "paths", updateMask.Paths)
		structuredreporting.ReportDiff(ctx, report)

		mapCtx := &direct.MapContext{}
		desired := BigtableClusterSpec_v1alpha1_ToProto(mapCtx, &spec)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		desired.Name = a.id.String()

		op, err := a.gcpClient.PartialUpdateCluster(ctx, &pb.PartialUpdateClusterRequest{
			Cluster:    desired,
			UpdateMask: updateMask,
		})
		if err != nil {
			return fmt.Errorf("updating BigtableCluster %s: %w", a.id, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for BigtableCluster %s update: %w", a.id, err)
		}
		a.actual = updated
		log.V(2).Info("successfully updated BigtableCluster", "name", a.id)
	} else {
		log.V(2).Info("no cluster fields need update", "name", a.id)
	}

	// Reconcile memory layer separately.
	if err := a.reconcileMemoryLayer(ctx); err != nil {
		return err
	}

	return a.setStatus(ctx, func(status *krm.BigtableClusterStatus) error {
		return updateOp.UpdateStatus(ctx, status, nil)
	})
}

// reconcileMemoryLayer reconciles the singleton memory layer sub-resource.
// spec.memoryLayers is a list for forward-compatibility; at most one entry is
// currently supported by the Bigtable API.
func (a *ClusterAdapter) reconcileMemoryLayer(ctx context.Context) error {
	log := klog.FromContext(ctx)

	if len(a.desired.Spec.MemoryLayers) > 1 {
		return fmt.Errorf("BigtableCluster %s: at most one memoryLayer entry is supported (got %d)", a.id, len(a.desired.Spec.MemoryLayers))
	}

	desiredEnabled := len(a.desired.Spec.MemoryLayers) == 1
	actualEnabled := a.actualML != nil && a.actualML.MemoryConfig != nil

	if desiredEnabled == actualEnabled {
		return nil
	}

	log.V(2).Info("reconciling MemoryLayer", "name", a.id, "desiredEnabled", desiredEnabled)

	mlName := a.id.String() + "/memoryLayer"
	ml := &pb.MemoryLayer{Name: mlName}
	if desiredEnabled {
		ml.MemoryConfig = &pb.MemoryLayer_MemoryConfig{}
	}

	op, err := a.gcpClient.UpdateMemoryLayer(ctx, &pb.UpdateMemoryLayerRequest{
		MemoryLayer: ml,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"memory_config"}},
	})
	if err != nil {
		return fmt.Errorf("updating MemoryLayer for BigtableCluster %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for MemoryLayer update for BigtableCluster %s: %w", a.id, err)
	}
	a.actualML = updated
	log.V(2).Info("successfully reconciled MemoryLayer", "name", a.id)
	return nil
}

func (a *ClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}
	obj := &krm.BigtableCluster{}
	mapCtx := &direct.MapContext{}
	spec := BigtableClusterSpec_v1alpha1_FromProto(mapCtx, a.actual)
	obj.Spec = direct.ValueOf(spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableClusterGVK)
	u.Object = uObj
	return u, nil
}

func (a *ClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableCluster", "name", a.id)

	if err := a.gcpClient.DeleteCluster(ctx, &pb.DeleteClusterRequest{Name: a.id.String()}); err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BigtableCluster, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableCluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableCluster", "name", a.id)
	return true, nil
}

func (a *ClusterAdapter) setStatus(ctx context.Context, updateFn func(status *krm.BigtableClusterStatus) error) error {
	status := &krm.BigtableClusterStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = &krm.BigtableClusterObservedState{}
	if a.actual != nil {
		mapCtx := &direct.MapContext{}
		observed := BigtableClusterObservedState_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ObservedState = observed
	}
	if a.actualML != nil {
		mlState := krm.Cluster_MemoryLayerObservedState{
			State: direct.LazyPtr(a.actualML.State.String()),
		}
		if mc := a.actualML.GetMemoryConfig(); mc != nil {
			gib := mc.GetStorageSizeGib()
			mlState.StorageSizeGib = &gib
		}
		status.ObservedState.MemoryLayers = []krm.Cluster_MemoryLayerObservedState{mlState}
	}
	return updateFn(status)
}
