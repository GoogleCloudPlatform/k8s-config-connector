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

package bigtable

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	api "google.golang.org/api/bigtableadmin/v2"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

// AddInstanceController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddInstanceController(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.BigtableInstanceGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &instanceModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m, opts)
}

type instanceModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &instanceModel{}

type instanceAdapter struct {
	projectID  string
	resourceID string

	desired        *krm.BigtableInstance
	actual         *api.Instance
	actualClusters []*api.Cluster

	*gcpClient
	adminClient *api.Service
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &instanceAdapter{}

// AdapterForObject implements the Model interface.
func (m *instanceModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	adminClient, err := m.newAdminClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.BigtableInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	// // resourceID is server-generated, no fallback
	// // TODO: How do we do resource acquisition - maybe by shortname?
	// resourceID = strings.TrimPrefix(resourceID, "instances/")
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource id")
	}

	projectID := obj.Annotations[k8s.ProjectIDAnnotation]
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project id")
	}

	return &instanceAdapter{
		projectID:   projectID,
		resourceID:  resourceID,
		desired:     obj,
		gcpClient:   m.gcpClient,
		adminClient: adminClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	instance, err := a.adminClient.Projects.Instances.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = instance

	clusters, err := a.adminClient.Projects.Instances.Clusters.List(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	// This is documented as not happening
	if clusters.NextPageToken != "" {
		return false, fmt.Errorf("clusters list was unexpected paginated")
	}
	a.actualClusters = clusters.Clusters

	return true, nil
}

// Delete implements the Adapter interface.
func (a *instanceAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	_, err := a.adminClient.Projects.Instances.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting instance %s: %w", a.fullyQualifiedName(), err)
	}

	// if _, err := op.Wait(ctx); err != nil {
	// 	return false, fmt.Errorf("instance deletion failed: %w", err)
	// }
	// TODO: Do we need to check that it was deleted?

	return true, nil
}

// Create implements the Adapter interface.
func (a *instanceAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := "projects/" + a.projectID

	mapCtx := &MapContext{ProjectID: a.projectID}
	instance := Instance_ToProto(mapCtx, &a.desired.Spec)
	if err := mapCtx.Err(); err != nil {
		return fmt.Errorf("mapping to API: %w", err)
	}

	// TODO instance.Labels

	// TODO: Is this needed?
	// instance.Name = a.resourceID

	req := &api.CreateInstanceRequest{}
	req.Instance = &api.Instance{}
	req.Clusters = make(map[string]api.Cluster)
	req.InstanceId = a.resourceID
	req.Parent = parent

	if err := convertViaJSON(instance, req.Instance); err != nil {
		return fmt.Errorf("converting instance via JSON: %w", err)
	}

	for _, cluster := range a.desired.Spec.Cluster {
		out := Cluster_ToProto(mapCtx, &cluster)
		if err := mapCtx.Err(); err != nil {
			return fmt.Errorf("mapping to API: %w", err)
		}
		outAPI := api.Cluster{}
		if err := convertViaJSON(out, &outAPI); err != nil {
			return fmt.Errorf("converting cluster via JSON: %w", err)
		}
		req.Clusters[cluster.ClusterId] = outAPI
	}

	log.Info("creating instance", "request", req)
	op, err := a.adminClient.Projects.Instances.Create(parent, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating instance: %w", err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("instance create operation failed: %w", err)
	}

	created, err := a.adminClient.Projects.Instances.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created instance: %w", err)
	}

	log.V(2).Info("created instance", "instance", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	status := &krm.BigtableInstanceStatus{}
	if err := instanceStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func convertViaJSON(in proto.Message, dest any) error {
	klog.Infof("convertViaJSON: %v", in)
	b, err := protojson.Marshal(in)
	if err != nil {
		return fmt.Errorf("converting proto to json: %w", err)
	}
	klog.Infof("convertViaJSON: json=%v", string(b))
	if err := json.Unmarshal(b, dest); err != nil {
		return fmt.Errorf("converting protojson to API: %w", err)
	}
	klog.Infof("convertViaJSON: result=%v", dest)
	return nil
}

func (a *instanceAdapter) waitForOp(ctx context.Context, op *api.Operation) error {
	// TODO: Only wait a short time, then update status?
	for {
		current, err := a.adminClient.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func instanceStatusToKRM(in *api.Instance, out *krm.BigtableInstanceStatus) error {
	return nil
}

// func timeToKRMString(t *timestamppb.Timestamp) *string {
// 	if t == nil {
// 		return nil
// 	}
// 	s := t.AsTime().Format(time.RFC3339Nano)
// 	return &s
// }

// Update implements the Adapter interface.
func (a *instanceAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}
	update := &pb.Instance{}
	update.Name = a.fullyQualifiedName()

	mapCtx := &MapContext{ProjectID: a.projectID}
	desiredProto := Instance_ToProto(mapCtx, &a.desired.Spec)
	if err := mapCtx.Err(); err != nil {
		return fmt.Errorf("mapping to API: %w", err)
	}

	if a.desired.Spec.DisplayName != nil && desiredProto.GetDisplayName() != a.actual.DisplayName {
		updateMask.Paths = append(updateMask.Paths, "displayName")
	}

	if a.desired.Spec.InstanceType != nil && desiredProto.GetType().String() != a.actual.Type {
		klog.Infof("type changed from %s to %s", a.actual.Type, desiredProto.GetType().String())
		updateMask.Paths = append(updateMask.Paths, "type")
	}

	// TODO: Where/how do we want to enforce immutability?

	if len(updateMask.Paths) != 0 {

		instanceAPI := &api.Instance{}
		if err := convertViaJSON(desiredProto, instanceAPI); err != nil {
			return fmt.Errorf("converting instance via JSON: %w", err)
		}

		log.Info("update instance", "request", instanceAPI, "updateMask", updateMask.Paths)
		op, err := a.adminClient.Projects.Instances.PartialUpdateInstance(a.fullyQualifiedName(), instanceAPI).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
		if err != nil {
			return err
		}

		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("instance update operation failed: %w", err)
		}

		updated, err := a.adminClient.Projects.Instances.Get(a.fullyQualifiedName()).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated instance: %w", err)
		}
		// TODO: Do we need to check that the operation succeeeded?
		log.V(2).Info("updated instance", "instance", updated)
	}

	// TODO!

	actualClusterMap := make(map[string]*api.Cluster)
	for _, cluster := range a.actualClusters {
		actualClusterMap[lastComponent(cluster.Name)] = cluster
	}

	inSpec := make(map[string]bool)
	for _, desiredCluster := range a.desired.Spec.Cluster {
		actual := actualClusterMap[desiredCluster.ClusterId]
		if actual == nil {
			if err := a.createCluster(ctx, &desiredCluster); err != nil {
				return err
			}
		} else {
			if err := a.updateCluster(ctx, &desiredCluster, actual); err != nil {
				return err
			}
		}
		inSpec[desiredCluster.ClusterId] = true
	}

	for key, cluster := range actualClusterMap {
		if !inSpec[key] {
			if err := a.deleteCluster(ctx, cluster); err != nil {
				return err
			}
		}
	}

	// TODO: Return updated object status
	return nil
}

func (a *instanceAdapter) updateCluster(ctx context.Context, desired *krm.InstanceCluster, actual *api.Cluster) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating cluster", "desired", a.desired)

	updateMask := &fieldmaskpb.FieldMask{}

	fqn := actual.Name

	mapCtx := &MapContext{ProjectID: a.projectID}
	desiredProto := Cluster_ToProto(mapCtx, desired)
	if err := mapCtx.Err(); err != nil {
		return fmt.Errorf("mapping to API: %w", err)
	}

	isAutoscaling := desired.AutoscalingConfig != nil

	if isAutoscaling {
		desiredProto.ServeNodes = 0 // Treat as output-only

		if actual.ClusterConfig != nil && actual.ClusterConfig.ClusterAutoscalingConfig != nil {
			updateMask.Paths = append(updateMask.Paths, "clusterConfig.clusterAutoscalingConfig")
		}
	} else {
		if desired.NumNodes != nil && ValueOf(desired.NumNodes) != int(actual.ServeNodes) {
			updateMask.Paths = append(updateMask.Paths, "serveNodes")
		}
	}

	// TODO: Where/how do we want to enforce immutability?

	if len(updateMask.Paths) != 0 {

		clusterAPI := &api.Cluster{}
		if err := convertViaJSON(desiredProto, clusterAPI); err != nil {
			return fmt.Errorf("converting instance via JSON: %w", err)
		}

		log.Info("update cluster", "request", clusterAPI, "updateMask", updateMask.Paths)
		op, err := a.adminClient.Projects.Instances.Clusters.PartialUpdateCluster(fqn, clusterAPI).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
		if err != nil {
			return err
		}

		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("cluster update operation failed: %w", err)
		}

		updated, err := a.adminClient.Projects.Instances.Clusters.Get(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated cluster: %w", err)
		}
		// TODO: Do we need to check that the operation succeeeded?
		log.V(2).Info("updated cluster", "cluster", updated)

	}

	// TODO: Return updated object status
	return nil
}

func (a *instanceAdapter) createCluster(ctx context.Context, desired *krm.InstanceCluster) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating cluster", "desired", desired)

	fqn := a.fullyQualifiedName() + "/clusters/" + desired.ClusterId

	mapCtx := &MapContext{ProjectID: a.projectID}
	desiredProto := Cluster_ToProto(mapCtx, desired)
	if err := mapCtx.Err(); err != nil {
		return fmt.Errorf("mapping to API: %w", err)
	}

	isAutoscaling := desired.AutoscalingConfig != nil
	if isAutoscaling {
		desiredProto.ServeNodes = 0 // Treat as output-only
	}

	// TODO: Where/how do we want to enforce immutability?

	clusterAPI := &api.Cluster{}
	if err := convertViaJSON(desiredProto, clusterAPI); err != nil {
		return fmt.Errorf("converting instance via JSON: %w", err)
	}

	log.Info("create cluster", "name", fqn, "request", clusterAPI)
	op, err := a.adminClient.Projects.Instances.Clusters.Create(fqn, clusterAPI).Context(ctx).Do()
	if err != nil {
		return err
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("cluster create operation failed: %w", err)
	}

	created, err := a.adminClient.Projects.Instances.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created cluster: %w", err)
	}
	// TODO: Do we need to check that the operation succeeeded?
	log.V(2).Info("created cluster", "cluster", created)

	// TODO: Return updated object status
	return nil
}

func (a *instanceAdapter) deleteCluster(ctx context.Context, cluster *api.Cluster) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting cluster", "remove", cluster)

	fqn := cluster.Name
	log.Info("delete cluster", "name", fqn)
	_, err := a.adminClient.Projects.Instances.Clusters.Delete(fqn).Context(ctx).Do()
	if err != nil {
		return err
	}

	log.V(2).Info("deleted cluster", "cluster", cluster)

	// TODO: Return updated object status
	return nil
}

func (a *instanceAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/instances/%s", a.projectID, a.resourceID)
}

func Cluster_ToProto(ctx *MapContext, in *krm.InstanceCluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Name = in.ClusterId
	out.Location = fmt.Sprintf("projects/%s/locations/%s", ctx.ProjectID, in.Zone)
	// out.State = Enum_ToProto(ctx, &in.State)
	out.ServeNodes = int32(ValueOf(in.NumNodes))
	out.DefaultStorageType = Enum_ToProto[pb.StorageType](ctx, in.StorageType)
	if in.KmsKeyRef != nil {
		if in.KmsKeyRef.External == "" {
			ctx.Errorf("kmsKeyRef was not fully resolved")
		}
		out.EncryptionConfig = &pb.Cluster_EncryptionConfig{
			KmsKeyName: in.KmsKeyRef.External,
		}
	}

	if in.AutoscalingConfig != nil {
		clusterConfig := &pb.Cluster_ClusterConfig{
			ClusterAutoscalingConfig: &pb.Cluster_ClusterAutoscalingConfig{
				AutoscalingLimits: &pb.AutoscalingLimits{
					MinServeNodes: int32(in.AutoscalingConfig.MinNodes),
					MaxServeNodes: int32(in.AutoscalingConfig.MaxNodes),
				},
			},
		}

		clusterConfig.ClusterAutoscalingConfig.AutoscalingTargets = &pb.AutoscalingTargets{
			CpuUtilizationPercent: int32(in.AutoscalingConfig.CpuTarget),
		}
		if in.AutoscalingConfig.StorageTarget != nil {
			clusterConfig.ClusterAutoscalingConfig.AutoscalingTargets.StorageUtilizationGibPerNode = int32(ValueOf(in.AutoscalingConfig.StorageTarget))
		}

		out.Config = &pb.Cluster_ClusterConfig_{
			ClusterConfig: clusterConfig,
		}
	}
	return out
}
