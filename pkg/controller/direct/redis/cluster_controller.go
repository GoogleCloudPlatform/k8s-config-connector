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

	"google.golang.org/api/option"

	api "cloud.google.com/go/redis/cluster/apiv1"
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/redis/cluster/apiv1"
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
	id *krm.RedisClusterIdentity

	desired *pb.Cluster
	actual  *pb.Cluster

	clustersClient *api.CloudRedisClusterClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &redisClusterAdapter{}

func (m *redisClusterModel) client(ctx context.Context) (*gcp.CloudRedisClusterClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudRedisClusterRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building RedisCluster client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *redisClusterModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader

	// Get RedisCluster GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.RedisCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := RedisClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
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
		id:             id.(*krm.RedisClusterIdentity),
		desired:        desired,
		clustersClient: gcpClient,
	}, nil
}

func (m *redisClusterModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *redisClusterAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Cluster == "" {
		return false, nil
	}

	req := &pb.GetClusterRequest{
		Name: a.id.String(),
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
	if a.id.Cluster == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteClusterRequest{
		Name: a.id.String(),
	}

	op, err := a.clustersClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		if !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("deleting redisCluster %s: %w", a.id.String(), err)
		}
	}

	if err := op.Wait(ctx); err != nil {
		if !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("waiting for redisCluster delete %s: %w", a.id.String(), err)
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

	parent := "projects/" + a.id.Project + "/locations/" + a.id.Location

	req := &pb.CreateClusterRequest{
		Parent:    parent,
		ClusterId: a.id.Cluster,
		Cluster:   a.desired,
	}
	log.V(0).Info("making redis CreateCluster call", "request", req)

	op, err := a.clustersClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating cluster: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for redisCluster create %s: %w", a.id.String(), err)
	}

	log.V(0).Info("created redisCluster", "redisCluster", created)

	// Set externalRef
	status := &krm.RedisClusterStatus{}
	status.ExternalRef = direct.PtrTo(a.id.String())

	mapCtx := &direct.MapContext{}
	status.ObservedState = RedisClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update implements the Adapter interface.
func (a *redisClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	diffs, err := compareRedisCluster(ctx, a.actual, a.desired)
	if err != nil {
		return fmt.Errorf("comparing actual and desired RedisCluster: %w", err)
	}

	var latest *pb.Cluster
	if diffs.HasDiff() {
		diffs.Object = u
		structuredreporting.ReportDiff(ctx, diffs)

		// exactly 1 update_mask field must be specified per update request

		for _, field := range diffs.Fields {
			if field.ProtoFieldDescriptor == nil {
				return fmt.Errorf("unexpected diff field without proto descriptor: %s", field.ID)
			}
			path := string(field.ProtoFieldDescriptor.Name())

			req := &pb.UpdateClusterRequest{
				Cluster: a.desired,
			}
			req.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: []string{path},
			}

			req.Cluster.Name = a.id.String()

			log.V(0).Info("making redis UpdateCluster call", "request", req)

			op, err := a.clustersClient.UpdateCluster(ctx, req)
			if err != nil {
				return err
			}

			updated, err := op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting for redisCluster update %s: %w", a.id.String(), err)
			}
			log.V(0).Info("updated redisCluster", "redisCluster", updated)

			latest = updated
		}
	} else {
		latest = a.actual
	}

	status := &krm.RedisClusterStatus{}
	if status.ExternalRef == nil {
		// If it is the first reconciliation after switching to direct controller,
		// or is an acquisition with updates, then fill out the ExternalRef.
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	mapCtx := &direct.MapContext{}
	status.ObservedState = RedisClusterObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func compareRedisCluster(ctx context.Context, actual, desired *pb.Cluster) (*structuredreporting.Diff, error) {
	var maskedActual *pb.Cluster
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := RedisClusterSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		maskedActual = RedisClusterSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	maskedActual = populateDefaults(maskedActual)
	desired = populateDefaults(proto.CloneOf(desired))

	diffs, _, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func populateDefaults(cluster *pb.Cluster) *pb.Cluster {
	if cluster.PersistenceConfig == nil {
		cluster.PersistenceConfig = &pb.ClusterPersistenceConfig{}
	}
	if cluster.PersistenceConfig.Mode == pb.ClusterPersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
		cluster.PersistenceConfig.Mode = pb.ClusterPersistenceConfig_DISABLED
	}
	if cluster.AuthorizationMode == pb.AuthorizationMode_AUTH_MODE_UNSPECIFIED {
		cluster.AuthorizationMode = pb.AuthorizationMode_AUTH_MODE_DISABLED
	}
	if cluster.NodeType == pb.NodeType_NODE_TYPE_UNSPECIFIED {
		cluster.NodeType = pb.NodeType_REDIS_HIGHMEM_MEDIUM
	}
	if cluster.TransitEncryptionMode == pb.TransitEncryptionMode_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED {
		cluster.TransitEncryptionMode = pb.TransitEncryptionMode_TRANSIT_ENCRYPTION_MODE_DISABLED
	}
	if cluster.ZoneDistributionConfig == nil {
		cluster.ZoneDistributionConfig = &pb.ZoneDistributionConfig{}
	}
	if cluster.ZoneDistributionConfig.Mode == pb.ZoneDistributionConfig_ZONE_DISTRIBUTION_MODE_UNSPECIFIED {
		cluster.ZoneDistributionConfig.Mode = pb.ZoneDistributionConfig_MULTI_ZONE
	}

	if cluster.AutomatedBackupConfig != nil {
		if cluster.AutomatedBackupConfig.AutomatedBackupMode == pb.AutomatedBackupConfig_AUTOMATED_BACKUP_MODE_UNSPECIFIED {
			cluster.AutomatedBackupConfig.AutomatedBackupMode = pb.AutomatedBackupConfig_DISABLED
		}
		if cluster.AutomatedBackupConfig.AutomatedBackupMode == pb.AutomatedBackupConfig_DISABLED {
			cluster.AutomatedBackupConfig.Schedule = nil
			cluster.AutomatedBackupConfig.Retention = nil
		}
	}

	// clear pscConfig as it's not included in the response
	if cluster.PscConfigs != nil {
		cluster.PscConfigs = nil
	}

	return cluster
}
