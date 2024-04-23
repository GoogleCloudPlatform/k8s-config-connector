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

package redis

import (
	"context"
	"fmt"
	"strings"

	api "cloud.google.com/go/redis/cluster/apiv1"
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddRedisClusterController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddRedisClusterController(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.RedisClusterGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &redisClusterModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m, opts)
}

type redisClusterModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &redisClusterModel{}

type redisClusterAdapter struct {
	projectID  string
	location   string
	resourceID string

	// desired *krm.RedisCluster
	desiredProto *pb.Cluster
	actual       *pb.Cluster

	*gcpClient
	clustersClient *api.CloudRedisClusterClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &redisClusterAdapter{}

// AdapterForObject implements the Model interface.
func (m *redisClusterModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	redisClustersClient, err := m.newClusterClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.RedisCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	{
		tokens := strings.Split(projectID, "/")
		if len(tokens) == 1 {
			projectID = tokens[0]
		} else if len(tokens) == 2 && tokens[0] == "projects" {
			projectID = tokens[1]
		} else {
			return nil, fmt.Errorf("cannot resolve project from name %q", projectID)
		}
	}

	mapCtx := &MapContext{
		//	kube: kube,
	}
	desiredProto := ClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &redisClusterAdapter{
		projectID:  projectID,
		location:   location,
		resourceID: resourceID,
		// desired:          obj,
		desiredProto:   desiredProto,
		gcpClient:      m.gcpClient,
		clustersClient: redisClustersClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *redisClusterAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.GetClusterRequest{
		Name: a.fullyQualifiedName(),
	}
	redisCluster, err := a.clustersClient.GetCluster(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = redisCluster

	return true, nil
}

// Delete implements the Adapter interface.
func (a *redisClusterAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteClusterRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.clustersClient.DeleteCluster(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting redisCluster %s: %w", a.fullyQualifiedName(), err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for redisCluster delete %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *redisClusterAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	// TODO: Should be ref
	// parent := a.desired.Spec.Parent

	parent := "projects/" + a.projectID + "/locations/" + a.location

	req := &pb.CreateClusterRequest{
		Parent:    parent,
		ClusterId: a.resourceID,
		Cluster:   a.desiredProto,
	}
	log.V(0).Info("making redis CreateCluster call", "request", req)

	op, err := a.clustersClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating cluster: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for redisCluster create %s: %w", a.fullyQualifiedName(), err)
	}

	log.V(0).Info("created redisCluster", "redisCluster", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &MapContext{
		// kube: kube,
	}
	observedState := ClusterState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

// Update implements the Adapter interface.
func (a *redisClusterAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	// TODO: Where/how do we want to enforce immutability?

	updateMask := &fieldmaskpb.FieldMask{}

	// TODO: What if a different field (immutability again)

	if a.desiredProto.ReplicaCount != nil && ValueOf(a.desiredProto.ReplicaCount) != ValueOf(a.actual.ReplicaCount) {
		updateMask.Paths = append(updateMask.Paths, "replica_count")
	}

	if a.desiredProto.ShardCount != nil && ValueOf(a.desiredProto.ShardCount) != ValueOf(a.actual.ShardCount) {
		updateMask.Paths = append(updateMask.Paths, "shard_count")
	}

	// TODO: exactly 1 update_mask field must be specified per update request

	var latest *pb.Cluster
	if len(updateMask.Paths) != 0 {
		// unsupported path in fieldMask: size_gb.
		// Allowed values are deletion_protection_enabled, redis_configs, display_name, shard_count, replica_count, persistence_config

		req := &pb.UpdateClusterRequest{
			UpdateMask: updateMask,
			Cluster:    a.desiredProto,
		}
		req.Cluster.Name = a.fullyQualifiedName()

		log.V(0).Info("making redis UpdateCluster call", "request", req)

		op, err := a.clustersClient.UpdateCluster(ctx, req)
		if err != nil {
			return err
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for redisCluster update %s: %w", a.fullyQualifiedName(), err)
		}
		log.V(0).Info("updated redisCluster", "redisCluster", updated)

		latest = updated
	} else {
		latest = a.actual
	}

	mapCtx := &MapContext{
		// kube: kube,
	}
	observedState := ClusterState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

func (a *redisClusterAdapter) hasChanges() bool {
	desired := DeepCopyProto(a.desiredProto)
	removeOutputFields(desired)
	actual := DeepCopyProto(a.actual)
	removeOutputFields(actual)

	// Ignore a few "key" fields
	desired.Name = actual.Name

	// TODO: Log diffs?  Always?  Sometimes?
	return !proto.Equal(desired, actual)
}

func removeOutputFields(o *pb.Cluster) {
	o.CreateTime = nil
	o.State = pb.Cluster_STATE_UNSPECIFIED
	o.Uid = ""
	o.SizeGb = nil
	o.DiscoveryEndpoints = nil
	o.PscConnections = nil
	o.StateInfo = nil
}

func (a *redisClusterAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", a.projectID, a.location, a.resourceID)
}
