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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
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

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	nodePoolID := id.(*krm.ContainerNodePoolIdentity)

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

type operationType int

const (
	operationCreate operationType = iota
	operationUpdate
)

func (a *nodePoolAdapter) prepare(ctx context.Context, opType operationType) error {
	if err := resolveNodeCount(a.desiredKRM, a.id, opType); err != nil {
		return err
	}
	// Normalize references in Create/Update rather than AdapterForObject so that
	// already-deleted non-parent references do not block resource deletion.
	if err := common.NormalizeReferences(ctx, a.reader, a.desiredKRM, nil); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}
	return nil
}

func resolveNodeCount(obj *krm.ContainerNodePool, id *krm.ContainerNodePoolIdentity, opType operationType) error {
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
	// In the legacy TF controller (ResolveLegacyGCPManagedFields ->
	// resolveContainerNodePoolNodeCount in pkg/krmtotf/legacygcpmanagedfields.go),
	// on update (when liveState is non-empty) and when spec.autoscaling is enabled,
	// spec.nodeCount is treated as GCP-managed and ignored. On create (when liveState
	// is empty), ResolveLegacyGCPManagedFields returns early and spec.nodeCount is still used.
	//
	// TODO: Fix this behavior properly so that nodeCount and autoscaling handling is consistent
	// across create and update (rather than relying on legacy ResolveLegacyGCPManagedFields quirks).
	if obj.Spec.NodeCount != nil && (opType == operationCreate || obj.Spec.Autoscaling == nil) {
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
	// The GKE NodePool API only returns instanceGroupManagers URLs in instance_group_urls
	// (e.g. https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/instanceGroupManagers/{name}).
	// The legacy TF controller queried GCE instanceGroupManagers.get (primarily to read targetSize for node_count)
	// and used igm.instanceGroup, which GCE returns with /compute/beta/ because GKE provisions MIGs via the
	// Compute beta API. Since a managed instance group's instanceGroup always shares the same project, zone,
	// and name as its instanceGroupManager, we derive managedInstanceGroupUrls directly using /compute/beta/
	// to match the legacy controller's real-GCP output without extra GCE API calls.
	//
	// TODO: Node pools don't expose the current node count in their API, so read the instance groups instead.
	// If we need to expose nodeCount in status.observedState (e.g. when autoscaling is enabled),
	// we will need to query GCE instanceGroupManagers.get to read targetSize.
	var managedIgmUrls []string
	for _, u := range actual.GetInstanceGroupUrls() {
		migUrl := strings.Replace(u, "/instanceGroupManagers/", "/instanceGroups/", 1)
		migUrl = strings.Replace(migUrl, "/compute/v1/", "/compute/beta/", 1)
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

	if err := a.prepare(ctx, operationCreate); err != nil {
		return err
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

	// GKE rejects node pool creation if autoscaling is disabled with min_node_count specified.
	// When autoscaling is disabled, omit autoscaling on create.
	if as := desired.GetAutoscaling(); as != nil && !as.GetEnabled() {
		desired.Autoscaling = nil
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

	// Autoscaling defaults to enabled=true when configured unless explicitly disabled via spec.autoscaling.enabled=false.
	if as := pbNodePool.GetAutoscaling(); as != nil {
		if a.desiredKRM != nil && a.desiredKRM.Spec.Autoscaling != nil && a.desiredKRM.Spec.Autoscaling.Enabled != nil && !*a.desiredKRM.Spec.Autoscaling.Enabled {
			// When autoscaling is disabled, GKE does not maintain an autoscaling config on the node pool (it is nil).
			// Normalize disabled autoscaling to nil so that when actual is nil, no diff is detected.
			if !as.GetEnabled() {
				pbNodePool.Autoscaling = nil
			}
		} else {
			as.Enabled = true
		}
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

	if err := a.prepare(ctx, operationUpdate); err != nil {
		return err
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

	mergedDesired := diffRes.MergedDesired
	updateAutoscaling := diffRes.HasPathWithPrefix("autoscaling")
	updateManagement := diffRes.HasPathWithPrefix("management")
	updateNodeCount := diffRes.HasPathWithPrefix("initial_node_count")
	updateNodePoolRequest, updatePaths := a.buildUpdateNodePoolRequest(diffRes)
	updateNodePool := len(updatePaths) > 0

	diffRes.RemovePathWithPrefix("autoscaling", "management", "initial_node_count")
	diffRes.RemovePathWithPrefix(updatePaths...)
	if !diffRes.Empty() {
		return k8s.NewImmutableFieldsMutationError(diffRes.FieldMask.GetPaths())
	}

	// 1. Update Autoscaling via UpdateCluster if changed
	if updateAutoscaling {
		desiredAutoscaling := mergedDesired.GetAutoscaling()
		if desiredAutoscaling == nil || !desiredAutoscaling.GetEnabled() {
			// GKE requires desiredNodePoolAutoscaling to only have enabled: false without min/max counts.
			desiredAutoscaling = &pb.NodePoolAutoscaling{
				Enabled: false,
			}
		}
		op, err := a.client.UpdateCluster(ctx, &pb.UpdateClusterRequest{
			Name: a.id.ParentString(),
			Update: &pb.ClusterUpdate{
				DesiredNodePoolId:          a.id.NodePool,
				DesiredNodePoolAutoscaling: desiredAutoscaling,
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
	if updateManagement {
		op, err := a.client.SetNodePoolManagement(ctx, &pb.SetNodePoolManagementRequest{
			Name:       a.id.String(),
			Management: mergedDesired.GetManagement(),
		})
		if err != nil {
			return fmt.Errorf("updating management for ContainerNodePool %q: %w", a.id.String(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for management update for ContainerNodePool %q: %w", a.id.String(), err)
		}
	}

	// 3. Update Node Pool Size if node count changed
	if updateNodeCount {
		op, err := a.client.SetNodePoolSize(ctx, &pb.SetNodePoolSizeRequest{
			Name:      a.id.String(),
			NodeCount: mergedDesired.GetInitialNodeCount(),
		})
		if err != nil {
			return fmt.Errorf("setting size for ContainerNodePool %q: %w", a.id.String(), err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for size update for ContainerNodePool %q: %w", a.id.String(), err)
		}
	}

	// 4. Update NodePool fields via UpdateNodePool
	if updateNodePool {
		op, err := a.client.UpdateNodePool(ctx, updateNodePoolRequest)
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

func (a *nodePoolAdapter) buildUpdateNodePoolRequest(diffRes *common.SpecDiffResult[*pb.NodePool]) (*pb.UpdateNodePoolRequest, []string) {
	mergedDesired := diffRes.MergedDesired
	updateReq := &pb.UpdateNodePoolRequest{
		Name: a.id.String(),
	}
	var updatePaths []string

	if diffRes.HasPathWithPrefix("config.taints") {
		updateReq.Taints = &pb.NodeTaints{
			Taints: mergedDesired.GetConfig().GetTaints(),
		}
		updatePaths = append(updatePaths, "config.taints")
	}
	if diffRes.HasPathWithPrefix("config.tags") {
		updateReq.Tags = &pb.NetworkTags{
			Tags: mergedDesired.GetConfig().GetTags(),
		}
		updatePaths = append(updatePaths, "config.tags")
	}
	if diffRes.HasPathWithPrefix("config.labels") {
		updateReq.Labels = &pb.NodeLabels{
			Labels: mergedDesired.GetConfig().GetLabels(),
		}
		updatePaths = append(updatePaths, "config.labels")
	}
	if diffRes.HasPathWithPrefix("config.linux_node_config") {
		linuxNodeConfig := proto.Clone(mergedDesired.GetConfig().GetLinuxNodeConfig()).(*pb.LinuxNodeConfig)
		if linuxNodeConfig != nil {
			var rawSwap *krm.LinuxNodeConfig_SwapConfig
			if a.desiredKRM.Spec.NodeConfig != nil && a.desiredKRM.Spec.NodeConfig.LinuxNodeConfig != nil {
				rawSwap = a.desiredKRM.Spec.NodeConfig.LinuxNodeConfig.SwapConfig
			}
			if rawSwap != nil && rawSwap.Enabled != nil && !*rawSwap.Enabled {
				// GKE API requires the request not to preserve the live state of other fields when the swap feature is disabled.
				// Reset them to the raw desired state to preserve user intention.
				linuxNodeConfig.SwapConfig = LinuxNodeConfig_SwapConfig_ToProto(&direct.MapContext{}, rawSwap)
			}
		}
		updateReq.LinuxNodeConfig = linuxNodeConfig
		updatePaths = append(updatePaths, "config.linux_node_config")
	}
	if diffRes.HasPathWithPrefix("config.kubelet_config") {
		updateReq.KubeletConfig = mergedDesired.GetConfig().GetKubeletConfig()
		updatePaths = append(updatePaths, "config.kubelet_config")
	}
	if diffRes.HasPathWithPrefix("config.resource_manager_tags") {
		updateReq.ResourceManagerTags = mergedDesired.GetConfig().GetResourceManagerTags()
		updatePaths = append(updatePaths, "config.resource_manager_tags")
	}
	if diffRes.HasPathWithPrefix("config.containerd_config") {
		containerdConfig := proto.Clone(mergedDesired.GetConfig().GetContainerdConfig()).(*pb.ContainerdConfig)
		if containerdConfig != nil {
			var rawRegAccess *krm.PrivateRegistryAccessConfig
			if a.desiredKRM.Spec.NodeConfig != nil && a.desiredKRM.Spec.NodeConfig.ContainerdConfig != nil {
				rawRegAccess = a.desiredKRM.Spec.NodeConfig.ContainerdConfig.PrivateRegistryAccessConfig
			}
			if rawRegAccess != nil && rawRegAccess.Enabled != nil && !*rawRegAccess.Enabled {
				// GKE API requires the request not to preserve the live state of other fields when private registry access is disabled.
				// Reset them to the raw desired state to preserve user intention.
				containerdConfig.PrivateRegistryAccessConfig = PrivateRegistryAccessConfig_ToProto(&direct.MapContext{}, rawRegAccess)
			}
		}
		updateReq.ContainerdConfig = containerdConfig
		updatePaths = append(updatePaths, "config.containerd_config")
	}
	if diffRes.HasPathWithPrefix("config.workload_metadata_config") {
		updateReq.WorkloadMetadataConfig = mergedDesired.GetConfig().GetWorkloadMetadataConfig()
		updatePaths = append(updatePaths, "config.workload_metadata_config")
	}
	if diffRes.HasPathWithPrefix("config.logging_config") {
		updateReq.LoggingConfig = mergedDesired.GetConfig().GetLoggingConfig()
		updatePaths = append(updatePaths, "config.logging_config")
	}
	if diffRes.HasPathWithPrefix("config.resource_labels") {
		updateReq.ResourceLabels = &pb.ResourceLabels{
			Labels: mergedDesired.GetConfig().GetResourceLabels(),
		}
		updatePaths = append(updatePaths, "config.resource_labels")
	}
	if diffRes.HasPathWithPrefix("config.windows_node_config") {
		updateReq.WindowsNodeConfig = mergedDesired.GetConfig().GetWindowsNodeConfig()
		updatePaths = append(updatePaths, "config.windows_node_config")
	}
	if diffRes.HasPathWithPrefix("config.image_type") {
		updateReq.ImageType = mergedDesired.GetConfig().GetImageType()
		updatePaths = append(updatePaths, "config.image_type")
	}
	if diffRes.HasPathWithPrefix("config.fast_socket") {
		updateReq.FastSocket = mergedDesired.GetConfig().GetFastSocket()
		updatePaths = append(updatePaths, "config.fast_socket")
	}
	if diffRes.HasPathWithPrefix("config.gvnic") {
		updateReq.Gvnic = mergedDesired.GetConfig().GetGvnic()
		updatePaths = append(updatePaths, "config.gvnic")
	}
	if diffRes.HasPathWithPrefix("config.confidential_nodes") {
		updateReq.ConfidentialNodes = mergedDesired.GetConfig().GetConfidentialNodes()
		updatePaths = append(updatePaths, "config.confidential_nodes")
	}
	if diffRes.HasPathWithPrefix("network_config") {
		updateReq.NodeNetworkConfig = mergedDesired.GetNetworkConfig()
		updatePaths = append(updatePaths, "network_config")
	}
	if diffRes.HasPathWithPrefix("upgrade_settings") {
		updateReq.UpgradeSettings = mergedDesired.GetUpgradeSettings()
		updatePaths = append(updatePaths, "upgrade_settings")
	}
	if diffRes.HasPathWithPrefix("locations") {
		updateReq.Locations = mergedDesired.GetLocations()
		updatePaths = append(updatePaths, "locations")
	}
	if diffRes.HasPathWithPrefix("version") {
		updateReq.NodeVersion = mergedDesired.GetVersion()
		updatePaths = append(updatePaths, "version")
	}
	if diffRes.HasPathWithPrefix("queued_provisioning") {
		updateReq.QueuedProvisioning = mergedDesired.GetQueuedProvisioning()
		updatePaths = append(updatePaths, "queued_provisioning")
	}

	return updateReq, updatePaths
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
