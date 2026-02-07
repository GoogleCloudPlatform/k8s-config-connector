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

package cluster

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	api "cloud.google.com/go/redis/cluster/apiv1"
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.RedisClusterGVK, newRedisClusterModel)
}

func newRedisClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &redisClusterModel{config: config}, nil
}

type redisClusterModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &redisClusterModel{}

type redisClusterAdapter struct {
	projectID  string
	location   string
	resourceID string

	desired *pb.Cluster
	actual  *pb.Cluster

	clustersClient *api.CloudRedisClusterClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &redisClusterAdapter{}

// AdapterForObject implements the Model interface.
func (m *redisClusterModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	redisClustersClient, err := gcpClient.newClusterClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.RedisCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := direct.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectRef, err := refs.ResolveProject(ctx, kube, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if err := common.VisitFields(obj, &refNormalizer{ctx: ctx, src: obj, project: *projectRef, kube: kube}); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := RedisClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Default to disabled persistence config
	if desired.PersistenceConfig == nil {
		desired.PersistenceConfig = &pb.ClusterPersistenceConfig{}
	}
	if desired.PersistenceConfig.Mode == pb.ClusterPersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
		desired.PersistenceConfig.Mode = pb.ClusterPersistenceConfig_DISABLED
	}

	// Mask out some fields in the persistenceConfig, to accommodate KRM "discriminated union" semantics
	switch desired.GetPersistenceConfig().GetMode() {
	case pb.ClusterPersistenceConfig_DISABLED:
		desired.PersistenceConfig.AofConfig = nil
		desired.PersistenceConfig.RdbConfig = nil
	case pb.ClusterPersistenceConfig_RDB:
		desired.PersistenceConfig.AofConfig = nil
	case pb.ClusterPersistenceConfig_AOF:
		desired.PersistenceConfig.RdbConfig = nil
	}

	return &redisClusterAdapter{
		projectID:      projectID,
		location:       location,
		resourceID:     resourceID,
		desired:        desired,
		clustersClient: redisClustersClient,
	}, nil
}

func (m *redisClusterModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
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
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = redisCluster

	return true, nil
}

// Delete implements the Adapter interface.
func (a *redisClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
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
		if direct.IsNotFound(err) {
			return false, nil
		}
		if !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("deleting redisCluster %s: %w", a.fullyQualifiedName(), err)
		}
	}

	if err := op.Wait(ctx); err != nil {
		if !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("waiting for redisCluster delete %s: %w", a.fullyQualifiedName(), err)
		}
	}

	return true, nil
}

func (a *redisClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Create implements the Adapter interface.
func (a *redisClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	parent := "projects/" + a.projectID + "/locations/" + a.location

	req := &pb.CreateClusterRequest{
		Parent:    parent,
		ClusterId: a.resourceID,
		Cluster:   a.desired,
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

	/* TODO: Any reason to write resourceID back?  It's required anyway
	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}
	*/

	mapCtx := &direct.MapContext{}
	observedState := RedisClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

// Update implements the Adapter interface.
func (a *redisClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	updateMask := &fieldmaskpb.FieldMask{}

	// TODO: What if a different field (immutability again)

	report := &structuredreporting.Diff{Object: u}
	if direct.ValueOf(a.desired.ReplicaCount) != direct.ValueOf(a.actual.ReplicaCount) {
		updateMask.Paths = append(updateMask.Paths, "replica_count")
		report.AddField("replica_count", a.actual.ReplicaCount, a.desired.ReplicaCount)
	}

	if direct.ValueOf(a.desired.ShardCount) != direct.ValueOf(a.actual.ShardCount) {
		updateMask.Paths = append(updateMask.Paths, "shard_count")
		report.AddField("shard_count", a.actual.ShardCount, a.desired.ShardCount)
	}

	if direct.ValueOf(a.desired.DeletionProtectionEnabled) != direct.ValueOf(a.actual.DeletionProtectionEnabled) {
		updateMask.Paths = append(updateMask.Paths, "deletion_protection_enabled")
		report.AddField("deletion_protection_enabled", a.actual.DeletionProtectionEnabled, a.desired.DeletionProtectionEnabled)
	}

	if !proto.Equal(a.desired.PersistenceConfig, a.actual.PersistenceConfig) {
		report.AddField("persistence_config", a.actual.PersistenceConfig, a.desired.PersistenceConfig)
		updateMask.Paths = append(updateMask.Paths, "persistence_config")
	}

	if !reflect.DeepEqual(a.desired.GetRedisConfigs(), a.actual.GetRedisConfigs()) {
		report.AddField("redis_configs", a.actual.GetRedisConfigs(), a.desired.GetRedisConfigs())
		updateMask.Paths = append(updateMask.Paths, "redis_configs")
	}

	var latest *pb.Cluster
	if len(updateMask.Paths) != 0 {
		structuredreporting.ReportDiff(ctx, report)

		// exactly 1 update_mask field must be specified per update request
		for _, path := range updateMask.Paths {
			req := &pb.UpdateClusterRequest{
				Cluster: a.desired,
			}
			req.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: []string{path},
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
		}
	} else {
		latest = a.actual
	}

	mapCtx := &direct.MapContext{}
	observedState := RedisClusterObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

func (a *redisClusterAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", a.projectID, a.location, a.resourceID)
}
