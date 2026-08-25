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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/hypercomputecluster/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/hypercomputecluster/apiv1"
	pb "cloud.google.com/go/hypercomputecluster/apiv1/hypercomputeclusterpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.HypercomputeClusterClusterGVK, NewModel)
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
		return nil, fmt.Errorf("building HypercomputeCluster client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.HypercomputeClusterCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	clusterID := id.(*krm.HypercomputeClusterClusterIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := HypercomputeClusterClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = clusterID.String()

	return &Adapter{
		id:        clusterID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.HypercomputeClusterClusterIdentity
	gcpClient *gcp.Client
	desired   *pb.Cluster
	actual    *pb.Cluster
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HypercomputeClusterCluster", "name", a.id.String())

	req := &pb.GetClusterRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HypercomputeClusterCluster %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HypercomputeClusterCluster", "name", a.id.String())

	req := &pb.CreateClusterRequest{
		Parent:    a.id.ParentString(),
		ClusterId: a.id.Cluster,
		Cluster:   a.desired,
	}

	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating HypercomputeClusterCluster %s: %w", a.id.String(), err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for HypercomputeClusterCluster %s creation: %w", a.id.String(), err)
	}

	// Always GET immediately after creation to fetch fully-populated fields
	getReq := &pb.GetClusterRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetCluster(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching newly created HypercomputeClusterCluster %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created HypercomputeClusterCluster", "name", a.id.String())

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HypercomputeClusterCluster", "name", a.id.String())

	diffs, updateMask, err := compareCluster(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateClusterRequest{
			Cluster:    a.desired,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateCluster(ctx, req)
		if err != nil {
			return fmt.Errorf("updating HypercomputeClusterCluster %s: %w", a.id.String(), err)
		}

		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for HypercomputeClusterCluster %s update: %w", a.id.String(), err)
		}

		// Always GET immediately after update to fetch fully-populated fields
		getReq := &pb.GetClusterRequest{Name: a.id.String()}
		latest, err = a.gcpClient.GetCluster(ctx, getReq)
		if err != nil {
			return fmt.Errorf("fetching updated HypercomputeClusterCluster %s: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Cluster) error {
	mapCtx := &direct.MapContext{}
	status := &krm.HypercomputeClusterClusterStatus{}
	status.ObservedState = HypercomputeClusterClusterObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Cluster)
	u.SetGroupVersionKind(krm.HypercomputeClusterClusterGVK)
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HypercomputeClusterCluster", "name", a.id.String())

	req := &pb.DeleteClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting HypercomputeClusterCluster %s: %w", a.id.String(), err)
	}

	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting HypercomputeClusterCluster %s waiting: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted HypercomputeClusterCluster", "name", a.id.String())
	return true, nil
}

func compareCluster(ctx context.Context, actual, desired *pb.Cluster) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	actualSpec := HypercomputeClusterClusterSpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	maskedActual := HypercomputeClusterClusterSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Cluster)

	// network_resources is immutable and cannot be updated.
	// To prevent false diffs (as GCP returns resolved references) and to avoid
	// disallowed fields in the update mask, we clear it from the compared objects.
	clonedDesired.NetworkResources = nil
	maskedActual.NetworkResources = nil

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
