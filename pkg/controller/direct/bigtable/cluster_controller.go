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

// +tool:controller
// proto.service: google.bigtable.admin.v2.Bigtable
// proto.message: google.bigtable.admin.v2.Cluster
// crd.type: BigtableCluster
// crd.version: v1alpha1

package bigtable

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/bigtable"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.BigtableClusterGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewClusterIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigtable GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newBigtableInstanceAdminClient(ctx)
	if err != nil {
		return nil, err
	}
	return &adapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type adapter struct {
	gcpClient *gcp.InstanceAdminClient
	id        *krm.ClusterIdentity
	desired   *krm.BigtableCluster
	actual    *pb.Cluster
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting bigtable cluster", "name", a.id)

	req := &pb.GetClusterRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCluster(ctx, req.)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting bigtable cluster %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating bigtable cluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateClusterRequest{
		Parent:    a.id.Parent().String(),
		ClusterId: a.id.ID(),
		Cluster:   resource,
	}
	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating bigtable cluster %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("bigtable cluster %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created bigtable cluster in gcp", "name", a.id)

	status := &krm.BigtableClusterStatus{}
	status.ObservedState = BigtableClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating bigtable cluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.actual.Name

	paths := []string{}
	if desired.Spec.ClusterConfig != nil && a.actual.Config != nil {
		if desired.Spec.ClusterConfig.ClusterAutoscalingConfig != nil && a.actual.Config == nil {
			// from manual scaling to autoscaling
			// note, this switch will also need serveNodes to be zero, we check this below
			paths = append(paths, "cluster_config.cluster_autoscaling_config")
		} else if desired.Spec.ClusterConfig.ClusterAutoscalingConfig == nil && a.actual.Config != nil {
			// from autoscaling to manual scaling
			// note, this switch will also need a non-zero serveNodes, we check this below
			paths = append(paths, "cluster_config.cluster_autoscaling_config")
		} else if desired.Spec.ClusterConfig.ClusterAutoscalingConfig != nil && a.actual.Config != nil && !reflect.DeepEqual(resource.ClusterConfig.ClusterAutoscalingConfig, a.actual.ClusterConfig.ClusterAutoscalingConfig) {
			// update autoscaling config
			paths = append(paths, "cluster_config.cluster_autoscaling_config")
		}
	}
	if desired.Spec.ServeNodes != nil && !reflect.DeepEqual(resource.ServeNodes, a.actual.ServeNodes) {
		if *desired.Spec.ServeNodes == 0 {
			// We do not support serveNodes to be zero.
			return fmt.Errorf("zero value is only allowed for serveNodes if clusterConfig.clusterAutoscalingConfig is set, and the intent is to switch to autoscaling")
		}
		if desired.Spec.ClusterConfig != nil && desired.Spec.ClusterConfig.ClusterAutoscalingConfig != nil {
			// We do not support setting serveNodes if clusterConfig.clusterAutoscalingConfig is set.
			return fmt.Errorf("non-zero value is only allowed for serveNodes if clusterConfig.clusterAutoscalingConfig is not set, and the intent is to switch to manual scaling")
		}
		paths = append(paths, "serve_nodes")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	log.V(2).Info("updating bigtable cluster: fields", "name", a.id, "fields", paths)
	op, err := a.gcpClient.UpdateCluster(ctx, resource)
	if err != nil {
		return fmt.Errorf("updating bigtable cluster %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("bigtable cluster %s waiting for update: %w", a.id, err)
	}

	status := &krm.BigtableClusterStatus{}

	log.V(2).Info("successfully updated bigtable cluster", "name", a.id)

	// update status as `ready` after operations succeeds
	status.ObservedState = BigtableClusterObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigtableClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceRef = &refs.InstanceRef{External: a.id.Parent().Instance}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.BigtableClusterGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting bigtable cluster", "name", a.id)

	req := &pb.DeleteClusterRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting bigtable cluster %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted bigtable cluster", "name", a.id)
	return true, nil
}
