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

package hypercomputecluster

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/hypercomputecluster/apiv1"
	pb "cloud.google.com/go/hypercomputecluster/apiv1/hypercomputeclusterpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/hypercomputecluster/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.HypercomputeClusterClusterGVK, NewClusterModel)
}

func NewClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCluster{config: *config}, nil
}

var _ directbase.Model = &modelCluster{}

type modelCluster struct {
	config config.ControllerConfig
}

func (m *modelCluster) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.HypercomputeClusterCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resolvedIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := resolvedIdentity.(*krm.HypercomputeClusterClusterIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", resolvedIdentity)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := HypercomputeClusterClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &ClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.HypercomputeClusterClusterIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ClusterAdapter struct {
	id        *krm.HypercomputeClusterClusterIdentity
	gcpClient *gcp.Client
	desired   *pb.Cluster
	actual    *pb.Cluster
}

var _ directbase.Adapter = &ClusterAdapter{}

func (a *ClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Cluster", "name", a.id)

	req := &pb.GetClusterRequest{Name: a.id.String()}
	clusterpb, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Cluster %q: %w", a.id, err)
	}

	a.actual = clusterpb
	return true, nil
}

func (a *ClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Cluster", "name", a.id)

	req := &pb.CreateClusterRequest{
		Parent:    a.id.ParentString(),
		ClusterId: a.id.Cluster,
		Cluster:   a.desired,
	}
	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Cluster %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for Cluster %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Cluster", "name", a.id)

	latest, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching Cluster after create: %w", err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *ClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Cluster", "name", a.id)

	diffs, updateMask, err := compareCluster(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for Cluster", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateClusterRequest{
		Cluster:    a.desired,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Cluster %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for Cluster %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Cluster", "name", a.id)

	latest, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching Cluster after update: %w", err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.HypercomputeClusterCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(HypercomputeClusterClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Cluster)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Cluster)
	u.SetGroupVersionKind(krm.HypercomputeClusterClusterGVK)

	return u, nil
}

func (a *ClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Cluster", "name", a.id)

	req := &pb.DeleteClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Cluster, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Cluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Cluster", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Cluster %s: %w", a.id, err)
	}
	return true, nil
}

func compareCluster(ctx context.Context, actual, desired *pb.Cluster) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, HypercomputeClusterClusterSpec_FromProto, HypercomputeClusterClusterSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Cluster)

	populateDefaults := func(obj *pb.Cluster) {
		// Populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *ClusterAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Cluster) error {
	mapCtx := &direct.MapContext{}
	status := &krm.HypercomputeClusterClusterStatus{}
	status.ObservedState = HypercomputeClusterClusterObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
