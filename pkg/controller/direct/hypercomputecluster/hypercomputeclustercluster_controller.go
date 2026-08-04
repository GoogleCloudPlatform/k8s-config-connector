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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/hypercomputecluster/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

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
		return nil, fmt.Errorf("building hypercomputecluster client: %w", err)
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

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	clusterIdentity := identity.(*krm.HypercomputeClusterClusterIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := HypercomputeClusterClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Propagate labels
	if desired.Labels == nil {
		desired.Labels = make(map[string]string)
	}
	for k, v := range obj.GetLabels() {
		desired.Labels[k] = v
	}

	return &Adapter{
		id:        clusterIdentity,
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

	if a.id.Cluster == "" { // resource is not yet created
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("getting HypercomputeClusterCluster", "name", fqn)

	req := &pb.GetClusterRequest{Name: fqn}
	clusterpb, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HypercomputeClusterCluster %q: %w", fqn, err)
	}

	a.actual = clusterpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HypercomputeClusterCluster", "name", fqn)

	parent := a.id.ParentString()
	req := &pb.CreateClusterRequest{
		Parent:    parent,
		ClusterId: a.id.Cluster,
		Cluster:   a.desired,
	}
	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Cluster %s: %w", fqn, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Cluster %s creation: %w", fqn, err)
	}

	log.V(2).Info("successfully created Cluster", "name", created.Name)

	latest, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: fqn})
	if err != nil {
		return fmt.Errorf("getting Cluster %s after creation: %w", fqn, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HypercomputeClusterCluster", "name", fqn)

	maskedActual, err := mappers.OnlySpecFields(a.actual, HypercomputeClusterClusterSpec_FromProto, HypercomputeClusterClusterSpec_ToProto)
	if err != nil {
		return err
	}

	clonedDesired := proto.Clone(a.desired).(*pb.Cluster)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", fqn)
		return nil
	}

	diffs.Object = u
	structuredreporting.ReportDiff(ctx, diffs)

	clonedDesired.Name = a.actual.Name
	req := &pb.UpdateClusterRequest{
		Cluster:    clonedDesired,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Cluster %s: %w", fqn, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Cluster %s update: %w", fqn, err)
	}
	log.V(2).Info("successfully updated Cluster", "name", fqn)

	latest, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: fqn})
	if err != nil {
		return fmt.Errorf("getting Cluster %s after update: %w", fqn, err)
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
	status.ExternalRef = &latest.Name
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
	obj.Spec.ResourceID = &a.id.Cluster
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.Cluster)
	u.SetGroupVersionKind(krm.HypercomputeClusterClusterGVK)
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting Cluster", "name", fqn)

	req := &pb.DeleteClusterRequest{Name: fqn}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting Cluster %s: %w", fqn, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting Cluster %s deletion: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted Cluster", "name", fqn)
	return true, nil
}
