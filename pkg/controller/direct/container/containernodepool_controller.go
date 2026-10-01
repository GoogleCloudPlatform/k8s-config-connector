// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	gcp "cloud.google.com/go/container/apiv1"
	pb "cloud.google.com/go/container/apiv1/containerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ContainerNodePoolGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &nodePoolModel{config: config}, nil
}

type nodePoolModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &nodePoolModel{}

type nodePoolAdapter struct {
	id         *krm.ContainerNodePoolIdentity
	desiredKRM *krm.ContainerNodePool
	actual     *pb.NodePool
	client     *gcp.ClusterManagerClient
}

var _ directbase.Adapter = &nodePoolAdapter{}

func (m *nodePoolModel) client(ctx context.Context) (*gcp.ClusterManagerClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	return gcp.NewClusterManagerRESTClient(ctx, opts...)
}

func (m *nodePoolModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.ContainerNodePool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("converting unstructured to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	nodePoolID := id.(*krm.ContainerNodePoolIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	client, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &nodePoolAdapter{
		id:         nodePoolID,
		desiredKRM: obj,
		client:     client,
	}, nil
}

func (m *nodePoolModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *nodePoolAdapter) location() string {
	if a.id.Zone != "" {
		return a.id.Zone
	}
	return a.id.Location
}

func (a *nodePoolAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", a.id.Project, a.location(), a.id.Cluster, a.id.NodePool)
}

func (a *nodePoolAdapter) parentFQN() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", a.id.Project, a.location(), a.id.Cluster)
}

func (a *nodePoolAdapter) waitForOperation(ctx context.Context, op *pb.Operation) error {
	if op == nil {
		return nil
	}
	opName := op.GetName()
	if !strings.HasPrefix(opName, "projects/") {
		opName = fmt.Sprintf("projects/%s/locations/%s/operations/%s", a.id.Project, a.location(), op.GetName())
	}
	pollInterval := 2 * time.Second
	_, err := common.WaitForOperation(ctx, pollInterval, func(op *pb.Operation) (bool, error) {
		if op.GetStatus() == pb.Operation_DONE {
			return true, nil
		}
		if op.GetStatus() == pb.Operation_ABORTING {
			return false, fmt.Errorf("operation %s aborted: %s", op.GetName(), op.GetStatusMessage())
		}
		return false, nil
	}, func() (*pb.Operation, error) {
		return a.client.GetOperation(ctx, &pb.GetOperationRequest{Name: opName})
	})
	return err
}

func (a *nodePoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding ContainerNodePool", "name", a.fullyQualifiedName())

	req := &pb.GetNodePoolRequest{
		Name: a.fullyQualifiedName(),
	}
	resp, err := a.client.GetNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = resp
	return true, nil
}

func (a *nodePoolAdapter) populateStatus(status *krm.ContainerNodePoolStatus, actual *pb.NodePool) {
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.InstanceGroupUrls = actual.GetInstanceGroupUrls()
	var managedIgmUrls []string
	for _, u := range actual.GetInstanceGroupUrls() {
		migUrl := strings.Replace(u, "/instanceGroupManagers/", "/instanceGroups/", 1)
		migUrl = strings.Replace(migUrl, "/compute/v1/", "/compute/v1beta1/", 1)
		managedIgmUrls = append(managedIgmUrls, migUrl)
	}
	status.ManagedInstanceGroupUrls = managedIgmUrls

	observedState := &krm.NodepoolObservedStateStatus{}
	if actual.GetVersion() != "" {
		observedState.Version = direct.LazyPtr(actual.GetVersion())
	}
	if actual.GetConfig() != nil && len(actual.GetConfig().GetTaints()) > 0 {
		mapCtx := &direct.MapContext{}
		taints := direct.Slice_FromProto(mapCtx, actual.GetConfig().GetTaints(), NodeTaint_FromProto)
		observedState.NodeConfig = &krm.NodePoolNodeConfigObservedState{
			Taint: taints,
		}
	}
	status.ObservedState = observedState
}

func (a *nodePoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ContainerNodePool", "name", a.fullyQualifiedName())

	mapCtx := &direct.MapContext{}
	desired := ContainerNodePoolSpec_ToProto(mapCtx, &a.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = a.id.NodePool

	req := &pb.CreateNodePoolRequest{
		Parent:   a.parentFQN(),
		NodePool: desired,
	}

	op, err := a.client.CreateNodePool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for creation of ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}

	created, err := a.client.GetNodePool(ctx, &pb.GetNodePoolRequest{Name: a.fullyQualifiedName()})
	if err != nil {
		return fmt.Errorf("getting created ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = created

	status := &krm.ContainerNodePoolStatus{}
	a.populateStatus(status, created)
	return createOp.UpdateStatus(ctx, status, nil)
}

func hasFieldChange(paths []string, prefix string) bool {
	for _, p := range paths {
		if p == prefix || strings.HasPrefix(p, prefix+".") {
			return true
		}
	}
	return false
}

func (a *nodePoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ContainerNodePool", "name", a.fullyQualifiedName())

	if a.actual == nil {
		return fmt.Errorf("actual is nil in Update")
	}

	_, fieldMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desiredKRM.Spec,
		a.actual,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		nil,
	)
	if err != nil {
		return fmt.Errorf("comparing specs for ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}

	if len(fieldMask.Paths) == 0 {
		log.V(2).Info("no changes detected for ContainerNodePool", "name", a.fullyQualifiedName())
		status := &krm.ContainerNodePoolStatus{}
		a.populateStatus(status, a.actual)
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	mapCtx := &direct.MapContext{}
	desired := ContainerNodePoolSpec_ToProto(mapCtx, &a.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// 1. Update Autoscaling via UpdateCluster if changed
	if hasFieldChange(fieldMask.Paths, "autoscaling") {
		op, err := a.client.UpdateCluster(ctx, &pb.UpdateClusterRequest{
			Name: a.parentFQN(),
			Update: &pb.ClusterUpdate{
				DesiredNodePoolId:          a.id.NodePool,
				DesiredNodePoolAutoscaling: desired.GetAutoscaling(),
			},
		})
		if err != nil {
			return fmt.Errorf("updating autoscaling for ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for autoscaling update for ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
		}
	}

	// 2. Update Management if changed
	if hasFieldChange(fieldMask.Paths, "management") {
		op, err := a.client.SetNodePoolManagement(ctx, &pb.SetNodePoolManagementRequest{
			Name:       a.fullyQualifiedName(),
			Management: desired.GetManagement(),
		})
		if err != nil {
			return fmt.Errorf("updating management for ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for management update for ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
		}
	}

	// 3. Update NodePool fields via UpdateNodePool
	updateReq := &pb.UpdateNodePoolRequest{
		Name: a.fullyQualifiedName(),
	}
	hasUpdates := false

	if hasFieldChange(fieldMask.Paths, "config") && desired.GetConfig() != nil {
		if !reflect.DeepEqual(desired.GetConfig().GetTaints(), a.actual.GetConfig().GetTaints()) {
			taints := desired.GetConfig().GetTaints()
			if taints == nil {
				taints = []*pb.NodeTaint{}
			}
			updateReq.Taints = &pb.NodeTaints{
				Taints: taints,
			}
			hasUpdates = true
		}
		if !reflect.DeepEqual(desired.GetConfig().GetTags(), a.actual.GetConfig().GetTags()) {
			updateReq.Tags = &pb.NetworkTags{
				Tags: desired.GetConfig().GetTags(),
			}
			hasUpdates = true
		}
		if !reflect.DeepEqual(desired.GetConfig().GetLabels(), a.actual.GetConfig().GetLabels()) {
			updateReq.Labels = &pb.NodeLabels{
				Labels: desired.GetConfig().GetLabels(),
			}
			hasUpdates = true
		}
		if desired.GetConfig().GetLinuxNodeConfig() != nil && !proto.Equal(desired.GetConfig().GetLinuxNodeConfig(), a.actual.GetConfig().GetLinuxNodeConfig()) {
			updateReq.LinuxNodeConfig = desired.GetConfig().GetLinuxNodeConfig()
			hasUpdates = true
		}
		if desired.GetConfig().GetKubeletConfig() != nil && !proto.Equal(desired.GetConfig().GetKubeletConfig(), a.actual.GetConfig().GetKubeletConfig()) {
			updateReq.KubeletConfig = desired.GetConfig().GetKubeletConfig()
			hasUpdates = true
		}
		if desired.GetConfig().GetResourceManagerTags() != nil && !proto.Equal(desired.GetConfig().GetResourceManagerTags(), a.actual.GetConfig().GetResourceManagerTags()) {
			updateReq.ResourceManagerTags = desired.GetConfig().GetResourceManagerTags()
			hasUpdates = true
		}
		if desired.GetConfig().GetContainerdConfig() != nil && !proto.Equal(desired.GetConfig().GetContainerdConfig(), a.actual.GetConfig().GetContainerdConfig()) {
			updateReq.ContainerdConfig = desired.GetConfig().GetContainerdConfig()
			hasUpdates = true
		}
		if desired.GetConfig().GetWorkloadMetadataConfig() != nil && !proto.Equal(desired.GetConfig().GetWorkloadMetadataConfig(), a.actual.GetConfig().GetWorkloadMetadataConfig()) {
			updateReq.WorkloadMetadataConfig = desired.GetConfig().GetWorkloadMetadataConfig()
			hasUpdates = true
		}
		if desired.GetConfig().GetLoggingConfig() != nil && !proto.Equal(desired.GetConfig().GetLoggingConfig(), a.actual.GetConfig().GetLoggingConfig()) {
			updateReq.LoggingConfig = desired.GetConfig().GetLoggingConfig()
			hasUpdates = true
		}
		if desired.GetConfig().GetResourceLabels() != nil && !reflect.DeepEqual(desired.GetConfig().GetResourceLabels(), a.actual.GetConfig().GetResourceLabels()) {
			updateReq.ResourceLabels = &pb.ResourceLabels{
				Labels: desired.GetConfig().GetResourceLabels(),
			}
			hasUpdates = true
		}
		if desired.GetConfig().GetWindowsNodeConfig() != nil && !proto.Equal(desired.GetConfig().GetWindowsNodeConfig(), a.actual.GetConfig().GetWindowsNodeConfig()) {
			updateReq.WindowsNodeConfig = desired.GetConfig().GetWindowsNodeConfig()
			hasUpdates = true
		}
		if desired.GetConfig().GetImageType() != "" && desired.GetConfig().GetImageType() != a.actual.GetConfig().GetImageType() {
			updateReq.ImageType = desired.GetConfig().GetImageType()
			hasUpdates = true
		}
	}
	if hasFieldChange(fieldMask.Paths, "upgrade_settings") {
		updateReq.UpgradeSettings = desired.GetUpgradeSettings()
		hasUpdates = true
	}
	if hasFieldChange(fieldMask.Paths, "locations") {
		updateReq.Locations = desired.GetLocations()
		hasUpdates = true
	}
	if hasFieldChange(fieldMask.Paths, "version") {
		updateReq.NodeVersion = desired.GetVersion()
		hasUpdates = true
	}
	if hasFieldChange(fieldMask.Paths, "queued_provisioning") {
		updateReq.QueuedProvisioning = desired.GetQueuedProvisioning()
		hasUpdates = true
	}

	if hasUpdates {
		op, err := a.client.UpdateNodePool(ctx, updateReq)
		if err != nil {
			return fmt.Errorf("updating ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
		}
	}

	updated, err := a.client.GetNodePool(ctx, &pb.GetNodePoolRequest{Name: a.fullyQualifiedName()})
	if err != nil {
		return fmt.Errorf("getting updated ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = updated

	status := &krm.ContainerNodePoolStatus{}
	a.populateStatus(status, updated)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *nodePoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ContainerNodePool", "name", a.fullyQualifiedName())

	req := &pb.DeleteNodePoolRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.client.DeleteNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for deletion of ContainerNodePool %q: %w", a.fullyQualifiedName(), err)
	}

	return false, nil
}

func (a *nodePoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("actual is nil for Export")
	}

	mapCtx := &direct.MapContext{}
	spec := ContainerNodePoolSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.ContainerNodePool{
		Spec: *spec,
	}
	obj.Spec.Location = a.location()
	obj.Spec.ClusterRef = krm.ContainerClusterRef{
		External: a.id.ParentString(),
	}

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	uObj := &unstructured.Unstructured{Object: u}
	uObj.SetGroupVersionKind(krm.ContainerNodePoolGVK)
	uObj.SetName(a.id.NodePool)
	return uObj, nil
}
