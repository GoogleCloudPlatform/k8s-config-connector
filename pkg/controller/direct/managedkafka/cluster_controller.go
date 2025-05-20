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

package managedkafka

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/managedkafka/apiv1"
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ManagedKafkaClusterGVK, NewClusterModel)
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
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCluster) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ManagedKafkaCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewClusterIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get managedkafka GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ClusterAdapter struct {
	id        *krm.ClusterIdentity
	gcpClient *gcp.Client
	desired   *krm.ManagedKafkaCluster
	actual    *pb.Cluster
	reader    client.Reader
}

var _ directbase.Adapter = &ClusterAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
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

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Cluster", "name", a.id)

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := ManagedKafkaClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateClusterRequest{
		Parent:    a.id.Parent().String(),
		ClusterId: a.id.ID(), // Note: this is not the fully qualified name for this resource, it is just the resource ID
		Cluster:   resource,
	}
	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Cluster %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of Cluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Cluster", "name", a.id)

	status := &krm.ManagedKafkaClusterStatus{}
	status.ObservedState = ManagedKafkaClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Cluster", "name", a.id)

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := ManagedKafkaClusterSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set the name field to ensure the GCP API can identity the resource during UpdateCluster().
	// This also prevents incorrect diffs, as the name field is not populated by ManagedKafkaClusterSpec_ToProto.
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.ManagedKafkaClusterStatus{}
		status.ObservedState = ManagedKafkaClusterObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateClusterRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths)},
		Cluster: desiredPb,
	}
	op, err := a.gcpClient.UpdateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Cluster %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting update for Cluster %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Cluster", "name", a.id.String())

	status := &krm.ManagedKafkaClusterStatus{}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	status.ObservedState = ManagedKafkaClusterObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ManagedKafkaCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ManagedKafkaClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ManagedKafkaClusterGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Cluster", "name", a.id)

	req := &pb.DeleteClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Cluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Cluster", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Cluster %s: %w", a.id, err)
	}
	return true, nil
}

func (a *ClusterAdapter) normalizeReference(ctx context.Context) error {
	obj := a.desired

	// Normalize the subnetworkRef in the accessConfig.networkConfigs
	if obj.Spec.GcpConfig != nil && obj.Spec.GcpConfig.AccessConfig != nil && obj.Spec.GcpConfig.AccessConfig.NetworkConfigs != nil {
		for i := range obj.Spec.GcpConfig.AccessConfig.NetworkConfigs {
			networkConfig := &obj.Spec.GcpConfig.AccessConfig.NetworkConfigs[i]
			if networkConfig.SubnetworkRef != nil {
				subnet, err := refs.ResolveComputeSubnetwork(ctx, a.reader, obj, networkConfig.SubnetworkRef)
				if err != nil {
					return err
				}
				networkConfig.SubnetworkRef = subnet
			}
		}
	}

	// Normalize the kmsKeyRef in the gcpConfig
	if obj.Spec.GcpConfig != nil && obj.Spec.GcpConfig.KmsKeyRef != nil {
		kmsKey, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj.Namespace, obj.Spec.GcpConfig.KmsKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.GcpConfig.KmsKeyRef = kmsKey
	}

	return nil
}
