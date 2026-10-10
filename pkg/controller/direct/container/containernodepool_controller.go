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
	"strings"
	"time"

	gcp "cloud.google.com/go/container/apiv1"
	pb "cloud.google.com/go/container/apiv1/containerpb"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	reader     client.Reader
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

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	nodePoolID := id.(*krm.ContainerNodePoolIdentity)

	if err := resolveNodeCount(obj, nodePoolID); err != nil {
		return nil, err
	}

	client, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &nodePoolAdapter{
		id:         nodePoolID,
		desiredKRM: obj,
		client:     client,
		reader:     reader,
	}, nil
}

func resolveNodeCount(obj *krm.ContainerNodePool, id *krm.ContainerNodePoolIdentity) error {
	// Terraform legacy behavior compatibility:
	// In the GKE API, the node pool size is specified via 'initial_node_count' at creation time,
	// and GKE updates this field upon resizing via SetNodePoolSize.
	// In Terraform's schema, both 'initial_node_count' (ForceNew) and 'node_count' (mutable in-place) exist.
	// In TF's expandNodePool, setting both initial_node_count and node_count is disallowed:
	// "Cannot set both initial_node_count and node_count on node pool".
	// However, if initialNodeCount == 0 (e.g. scale-from-zero or queued provisioning),
	// Terraform's d.GetOk("initial_node_count") evaluates to false (0 is zero-value),
	// allowing node_count to be set alongside initialNodeCount: 0.
	// We preserve this exact logic here by setting initialNodeCount from nodeCount so that
	// mappers and diffing work natively without customization.
	if obj.Spec.InitialNodeCount != nil && *obj.Spec.InitialNodeCount != 0 && obj.Spec.NodeCount != nil {
		idStr := obj.GetName()
		if id != nil {
			idStr = id.String()
		}
		return fmt.Errorf("cannot set both initialNodeCount and nodeCount on ContainerNodePool %s", idStr)
	}
	if obj.Spec.NodeCount != nil {
		obj.Spec.InitialNodeCount = obj.Spec.NodeCount
	}
	return nil
}

func (m *nodePoolModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *nodePoolAdapter) waitForOperation(ctx context.Context, op *pb.Operation) error {
	if op == nil {
		return nil
	}
	opName := op.GetName()
	if !strings.HasPrefix(opName, "projects/") {
		opName = fmt.Sprintf("projects/%s/locations/%s/operations/%s", a.id.Project, a.id.LocationValue(), op.GetName())
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
	log.V(2).Info("finding ContainerNodePool", "name", a.id.String())

	req := &pb.GetNodePoolRequest{
		Name: a.id.String(),
	}
	resp, err := a.client.GetNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContainerNodePool %q: %w", a.id.String(), err)
	}

	a.actual = resp
	return true, nil
}

func (a *nodePoolAdapter) populateStatus(status *krm.ContainerNodePoolStatus, actual *pb.NodePool) error {
	// Note on migration diff: We use the canonical location-based externalRef (projects/.../locations/.../clusters/.../nodePools/...)
	// across all operations to keep create, read, update, and delete consistent. This intentionally differs
	// from the legacy controller which used "zones/{zone}" for create and "locations/{location}" for update.
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.InstanceGroupUrls = actual.GetInstanceGroupUrls()
	var managedIgmUrls []string
	for _, u := range actual.GetInstanceGroupUrls() {
		migUrl := strings.Replace(u, "/instanceGroupManagers/", "/instanceGroups/", 1)
		migUrl = strings.Replace(migUrl, "/compute/v1/", "/compute/v1beta1/", 1)
		managedIgmUrls = append(managedIgmUrls, migUrl)
	}
	status.ManagedInstanceGroupUrls = managedIgmUrls

	mapCtx := &direct.MapContext{}
	status.ObservedState = NodepoolObservedStateStatus_FromProto(mapCtx, actual)
	return mapCtx.Err()
}

func (a *nodePoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ContainerNodePool", "name", a.id.String())

	if err := common.NormalizeReferences(ctx, a.reader, a.desiredKRM, nil); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ContainerNodePoolSpec_ToProto(mapCtx, &a.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desired.Name = a.id.NodePool

	if err := a.normalizeNodePool(ctx, desired); err != nil {
		return err
	}

	req := &pb.CreateNodePoolRequest{
		Parent:   a.id.ParentString(),
		NodePool: desired,
	}

	op, err := a.client.CreateNodePool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ContainerNodePool %q: %w", a.id.String(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for creation of ContainerNodePool %q: %w", a.id.String(), err)
	}

	created, err := a.client.GetNodePool(ctx, &pb.GetNodePoolRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created ContainerNodePool %q: %w", a.id.String(), err)
	}
	a.actual = created

	status := &krm.ContainerNodePoolStatus{}
	if err := a.populateStatus(status, created); err != nil {
		return err
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *nodePoolAdapter) normalizeNodePool(ctx context.Context, pbNodePool *pb.NodePool) error {
	if pbNodePool == nil {
		return nil
	}

	// In the KRM CRD, autoscaling does not expose an explicit `enabled` field; the presence
	// of the autoscaling configuration implies autoscaling is enabled. GCP GKE requires
	// enabled=true in the request whenever autoscaling is configured.
	if as := pbNodePool.GetAutoscaling(); as != nil {
		as.Enabled = true
	}

	if netConfig := pbNodePool.GetNetworkConfig(); netConfig != nil {
		if netConfig.GetSubnetwork() != "" {
			canonicalized, err := computev1beta1.CanonicalizeSubnetworkValue(ctx, netConfig.GetSubnetwork(), a.id.Project, a.id.LocationValue(), nil)
			if err != nil {
				return err
			}
			netConfig.Subnetwork = canonicalized
		}
		for _, addNodeNet := range netConfig.GetAdditionalNodeNetworkConfigs() {
			if addNodeNet.GetNetwork() != "" {
				addNodeNet.Network = computerefs.CanonicalizeNetworkValue(ctx, addNodeNet.GetNetwork(), a.id.Project, nil)
			}
			if addNodeNet.GetSubnetwork() != "" {
				canonicalized, err := computev1beta1.CanonicalizeSubnetworkValue(ctx, addNodeNet.GetSubnetwork(), a.id.Project, a.id.LocationValue(), nil)
				if err != nil {
					return err
				}
				addNodeNet.Subnetwork = canonicalized
			}
		}
		for _, addPodNet := range netConfig.GetAdditionalPodNetworkConfigs() {
			if addPodNet.GetSubnetwork() != "" {
				canonicalized, err := computev1beta1.CanonicalizeSubnetworkValue(ctx, addPodNet.GetSubnetwork(), a.id.Project, a.id.LocationValue(), nil)
				if err != nil {
					return err
				}
				addPodNet.Subnetwork = canonicalized
			}
		}
	}
	return nil
}

func (a *nodePoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ContainerNodePool", "name", a.id.String())

	if a.actual == nil {
		return fmt.Errorf("actual is nil in Update")
	}

	if err := common.NormalizeReferences(ctx, a.reader, a.desiredKRM, nil); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&a.desiredKRM.Spec,
		a.actual,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		a.normalizeNodePool,
		true, // fineGrained
	)
	if err != nil {
		return fmt.Errorf("comparing specs for ContainerNodePool %q: %w", a.id.String(), err)
	}

	if diffRes.Empty() {
		log.V(2).Info("no changes detected for ContainerNodePool", "name", a.id.String())
		status := &krm.ContainerNodePoolStatus{}
		if err := a.populateStatus(status, a.actual); err != nil {
			return err
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	structuredreporting.ReportDiff(ctx, diffRes.Diff)

	desired := diffRes.MergedDesired

	// 1. Update Autoscaling via UpdateCluster if changed
	if diffRes.Has("autoscaling") {
		op, err := a.client.UpdateCluster(ctx, &pb.UpdateClusterRequest{
			Name: a.id.ParentString(),
			Update: &pb.ClusterUpdate{
				DesiredNodePoolId:          a.id.NodePool,
				DesiredNodePoolAutoscaling: desired.GetAutoscaling(),
			},
		})
		if err != nil {
			return fmt.Errorf("updating autoscaling for ContainerNodePool %q: %w", a.id.String(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for autoscaling update for ContainerNodePool %q: %w", a.id.String(), err)
		}
	}

	// 2. Update Management if changed
	if diffRes.Has("management") {
		op, err := a.client.SetNodePoolManagement(ctx, &pb.SetNodePoolManagementRequest{
			Name:       a.id.String(),
			Management: desired.GetManagement(),
		})
		if err != nil {
			return fmt.Errorf("updating management for ContainerNodePool %q: %w", a.id.String(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for management update for ContainerNodePool %q: %w", a.id.String(), err)
		}
	}

	// 3. Update Node Pool Size if node count changed
	if diffRes.Has("initial_node_count") {
		op, err := a.client.SetNodePoolSize(ctx, &pb.SetNodePoolSizeRequest{
			Name:      a.id.String(),
			NodeCount: desired.GetInitialNodeCount(),
		})
		if err != nil {
			return fmt.Errorf("setting size for ContainerNodePool %q: %w", a.id.String(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for size update for ContainerNodePool %q: %w", a.id.String(), err)
		}
	}

	// 4. Update NodePool fields via UpdateNodePool
	updateReq := &pb.UpdateNodePoolRequest{
		Name: a.id.String(),
	}
	hasUpdates := false

	if diffRes.Has("config.taints") {
		updateReq.Taints = &pb.NodeTaints{
			Taints: desired.GetConfig().GetTaints(),
		}
		hasUpdates = true
	}
	if diffRes.Has("config.tags") {
		updateReq.Tags = &pb.NetworkTags{
			Tags: desired.GetConfig().GetTags(),
		}
		hasUpdates = true
	}
	if diffRes.Has("config.labels") {
		updateReq.Labels = &pb.NodeLabels{
			Labels: desired.GetConfig().GetLabels(),
		}
		hasUpdates = true
	}
	if diffRes.Has("config.linux_node_config") {
		updateReq.LinuxNodeConfig = desired.GetConfig().GetLinuxNodeConfig()
		hasUpdates = true
	}
	if diffRes.Has("config.kubelet_config") {
		updateReq.KubeletConfig = desired.GetConfig().GetKubeletConfig()
		hasUpdates = true
	}
	if diffRes.Has("config.resource_manager_tags") {
		updateReq.ResourceManagerTags = desired.GetConfig().GetResourceManagerTags()
		hasUpdates = true
	}
	if diffRes.Has("config.containerd_config") {
		updateReq.ContainerdConfig = desired.GetConfig().GetContainerdConfig()
		hasUpdates = true
	}
	if diffRes.Has("config.workload_metadata_config") {
		updateReq.WorkloadMetadataConfig = desired.GetConfig().GetWorkloadMetadataConfig()
		hasUpdates = true
	}
	if diffRes.Has("config.logging_config") {
		updateReq.LoggingConfig = desired.GetConfig().GetLoggingConfig()
		hasUpdates = true
	}
	if diffRes.Has("config.resource_labels") {
		updateReq.ResourceLabels = &pb.ResourceLabels{
			Labels: desired.GetConfig().GetResourceLabels(),
		}
		hasUpdates = true
	}
	if diffRes.Has("config.windows_node_config") {
		updateReq.WindowsNodeConfig = desired.GetConfig().GetWindowsNodeConfig()
		hasUpdates = true
	}
	if diffRes.Has("config.image_type") {
		updateReq.ImageType = desired.GetConfig().GetImageType()
		hasUpdates = true
	}
	if diffRes.Has("config.fast_socket") {
		updateReq.FastSocket = desired.GetConfig().GetFastSocket()
		hasUpdates = true
	}
	if diffRes.Has("config.gvnic") || diffRes.Has("config.virtual_nic") {
		updateReq.Gvnic = desired.GetConfig().GetGvnic()
		hasUpdates = true
	}
	if diffRes.Has("config.confidential_nodes") {
		updateReq.ConfidentialNodes = desired.GetConfig().GetConfidentialNodes()
		hasUpdates = true
	}
	if diffRes.Has("network_config") {
		updateReq.NodeNetworkConfig = desired.GetNetworkConfig()
		hasUpdates = true
	}
	if diffRes.Has("upgrade_settings") {
		updateReq.UpgradeSettings = desired.GetUpgradeSettings()
		hasUpdates = true
	}
	if diffRes.Has("locations") {
		updateReq.Locations = desired.GetLocations()
		hasUpdates = true
	}
	if diffRes.Has("version") {
		updateReq.NodeVersion = desired.GetVersion()
		hasUpdates = true
	}
	if diffRes.Has("queued_provisioning") {
		updateReq.QueuedProvisioning = desired.GetQueuedProvisioning()
		hasUpdates = true
	}

	if hasUpdates {
		op, err := a.client.UpdateNodePool(ctx, updateReq)
		if err != nil {
			return fmt.Errorf("updating ContainerNodePool %q: %w", a.id.String(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of ContainerNodePool %q: %w", a.id.String(), err)
		}
	}

	updated, err := a.client.GetNodePool(ctx, &pb.GetNodePoolRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting updated ContainerNodePool %q: %w", a.id.String(), err)
	}
	a.actual = updated

	status := &krm.ContainerNodePoolStatus{}
	if err := a.populateStatus(status, updated); err != nil {
		return err
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *nodePoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ContainerNodePool", "name", a.id.String())

	req := &pb.DeleteNodePoolRequest{
		Name: a.id.String(),
	}

	op, err := a.client.DeleteNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ContainerNodePool %q: %w", a.id.String(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for deletion of ContainerNodePool %q: %w", a.id.String(), err)
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
	obj.Spec.Location = a.id.LocationValue()
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
