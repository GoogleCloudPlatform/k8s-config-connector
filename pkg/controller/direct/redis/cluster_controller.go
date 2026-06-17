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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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

	if obj.Spec.KMSKeyRef != nil {
		resolvedKMSKey, err := refs.ResolveKMSCryptoKeyRef(ctx, kube, obj, obj.Spec.KMSKeyRef)
		if err != nil {
			return nil, fmt.Errorf("resolving KMSKeyRef: %w", err)
		}
		obj.Spec.KMSKeyRef = resolvedKMSKey
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

func (a *redisClusterAdapter) normalizeReplicationConfig(config *pb.CrossClusterReplicationConfig) {
	if config == nil {
		return
	}
	// The Redis API is strict:
	// - PRIMARY clusters must NOT have primary_cluster set.
	// - SECONDARY clusters must NOT have secondary_clusters set.
	// - NONE clusters must NOT have either set.
	role := config.GetClusterRole()
	switch role {
	case pb.CrossClusterReplicationConfig_PRIMARY:
		config.PrimaryCluster = nil
	case pb.CrossClusterReplicationConfig_SECONDARY:
		config.SecondaryClusters = nil
	case pb.CrossClusterReplicationConfig_NONE, pb.CrossClusterReplicationConfig_CLUSTER_ROLE_UNSPECIFIED:
		config.PrimaryCluster = nil
		config.SecondaryClusters = nil
	default:
		// do nothing, keep existing configs
	}
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
	a.normalizeReplicationConfig(req.Cluster.CrossClusterReplicationConfig)

	log.V(0).Info("making redis CreateCluster call", "request", req)

	op, err := a.clustersClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating cluster: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for redisCluster create %s: %w", a.id.String(), err)
	}

	if err := a.syncCrossClusterReplication(ctx); err != nil {
		return fmt.Errorf("syncing cross cluster replication: %w", err)
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
			a.normalizeReplicationConfig(req.Cluster.CrossClusterReplicationConfig)

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
	} else {
		cluster.AutomatedBackupConfig = &pb.AutomatedBackupConfig{AutomatedBackupMode: pb.AutomatedBackupConfig_DISABLED}
	}

	// clear pscConfig as it's not included in the response
	if cluster.PscConfigs != nil {
		cluster.PscConfigs = nil
	}

	return cluster
}

func (a *redisClusterAdapter) syncCrossClusterReplication(ctx context.Context) error {
	if a.desired.CrossClusterReplicationConfig == nil {
		return nil
	}

	role := a.desired.GetCrossClusterReplicationConfig().GetClusterRole()
	switch role {
	case pb.CrossClusterReplicationConfig_PRIMARY:
		// Ensure all secondary clusters have this cluster as primary
		for _, sc := range a.desired.GetCrossClusterReplicationConfig().GetSecondaryClusters() {
			if sc.GetCluster() == "" {
				continue
			}
			// Update secondary cluster to be SECONDARY with current cluster as primary
			remoteConfig := &pb.CrossClusterReplicationConfig{
				ClusterRole:    pb.CrossClusterReplicationConfig_SECONDARY,
				PrimaryCluster: &pb.CrossClusterReplicationConfig_RemoteCluster{Cluster: a.id.String()},
			}
			if err := a.updateRemoteCluster(ctx, sc.GetCluster(), remoteConfig); err != nil {
				return err
			}
		}
	case pb.CrossClusterReplicationConfig_SECONDARY:
		// Ensure primary cluster has this cluster as secondary
		pc := a.desired.GetCrossClusterReplicationConfig().GetPrimaryCluster()
		if pc.GetCluster() == "" {
			return nil
		}
		// Get primary cluster
		primary, err := a.clustersClient.GetCluster(ctx, &pb.GetClusterRequest{Name: pc.GetCluster()})
		if err != nil {
			if direct.IsNotFound(err) {
				log := klog.FromContext(ctx)
				log.V(0).Info("primary cluster not found, skipping sync", "primary", pc.GetCluster())
				return nil
			}
			return fmt.Errorf("getting primary cluster %s: %w", pc.GetCluster(), err)
		}
		// Add current cluster to secondary_clusters if not already there
		found := false
		for _, sc := range primary.GetCrossClusterReplicationConfig().GetSecondaryClusters() {
			if sc.GetCluster() == a.id.String() {
				found = true
				break
			}
		}
		if !found {
			newConfig := primary.GetCrossClusterReplicationConfig()
			if newConfig == nil {
				newConfig = &pb.CrossClusterReplicationConfig{
					ClusterRole: pb.CrossClusterReplicationConfig_PRIMARY,
				}
			} else {
				newConfig = proto.Clone(newConfig).(*pb.CrossClusterReplicationConfig)
				newConfig.ClusterRole = pb.CrossClusterReplicationConfig_PRIMARY
			}
			newConfig.PrimaryCluster = nil // PRIMARY should not have PrimaryCluster set
			newConfig.SecondaryClusters = append(newConfig.SecondaryClusters, &pb.CrossClusterReplicationConfig_RemoteCluster{
				Cluster: a.id.String(),
			})
			if err := a.updateRemoteCluster(ctx, pc.GetCluster(), newConfig); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *redisClusterAdapter) updateRemoteCluster(ctx context.Context, clusterName string, desiredConfig *pb.CrossClusterReplicationConfig) error {
	log := klog.FromContext(ctx)

	remote, err := a.clustersClient.GetCluster(ctx, &pb.GetClusterRequest{Name: clusterName})
	if err != nil {
		return fmt.Errorf("getting remote cluster %s: %w", clusterName, err)
	}

	// Only update if there is a difference in role or primary/secondary config
	diff := false
	if remote.GetCrossClusterReplicationConfig().GetClusterRole() != desiredConfig.GetClusterRole() {
		diff = true
	} else {
		switch desiredConfig.GetClusterRole() {
		case pb.CrossClusterReplicationConfig_PRIMARY:
			if len(remote.GetCrossClusterReplicationConfig().GetSecondaryClusters()) != len(desiredConfig.GetSecondaryClusters()) {
				diff = true
			} else {
				for i, sc := range desiredConfig.GetSecondaryClusters() {
					if sc.GetCluster() != remote.GetCrossClusterReplicationConfig().GetSecondaryClusters()[i].GetCluster() {
						diff = true
						break
					}
				}
			}
		case pb.CrossClusterReplicationConfig_SECONDARY:
			if remote.GetCrossClusterReplicationConfig().GetPrimaryCluster().GetCluster() != desiredConfig.GetPrimaryCluster().GetCluster() {
				diff = true
			}
		}
	}

	if !diff {
		return nil
	}

	normalizedConfig := proto.Clone(desiredConfig).(*pb.CrossClusterReplicationConfig)
	a.normalizeReplicationConfig(normalizedConfig)

	req := &pb.UpdateClusterRequest{
		Cluster: &pb.Cluster{
			Name:                          clusterName,
			CrossClusterReplicationConfig: normalizedConfig,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"cross_cluster_replication_config"},
		},
	}
	log.V(0).Info("making remote redis UpdateCluster call", "request", req)
	op, err := a.clustersClient.UpdateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating remote cluster %s: %w", clusterName, err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for remote cluster update %s: %w", clusterName, err)
	}
	return nil
}
