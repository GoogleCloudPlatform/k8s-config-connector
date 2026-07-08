// Copyright 2026 Google LLC
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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	gcp "cloud.google.com/go/redis/apiv1"
	redispb "cloud.google.com/go/redis/apiv1/redispb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.RedisInstanceGVK, newRedisInstanceModel)
}

func newRedisInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &redisInstanceModel{config: config}, nil
}

type redisInstanceModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &redisInstanceModel{}

type redisInstanceAdapter struct {
	id *krm.RedisInstanceIdentity

	desired *redispb.Instance
	actual  *redispb.Instance

	client *gcp.CloudRedisClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &redisInstanceAdapter{}

func (m *redisInstanceModel) client(ctx context.Context) (*gcp.CloudRedisClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudRedisRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building RedisInstance client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *redisInstanceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader

	// Get RedisInstance GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.RedisInstance{}
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
	desired := RedisInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &redisInstanceAdapter{
		id:      id.(*krm.RedisInstanceIdentity),
		desired: desired,
		client:  gcpClient,
	}, nil
}

func (m *redisInstanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *redisInstanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Instance == "" {
		return false, nil
	}

	req := &redispb.GetInstanceRequest{
		Name: a.id.String(),
	}
	redisInstance, err := a.client.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = redisInstance

	return true, nil
}

// Delete implements the Adapter interface.
func (a *redisInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Already deleted
	if a.id.Instance == "" {
		return false, nil
	}

	req := &redispb.DeleteInstanceRequest{
		Name: a.id.String(),
	}

	op, err := a.client.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		if !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("deleting redisInstance %s: %w", a.id.String(), err)
		}
	}

	if err := op.Wait(ctx); err != nil {
		if !strings.Contains(err.Error(), "missing \"value\" field") {
			return false, fmt.Errorf("waiting for redisInstance delete %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *redisInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Create implements the Adapter interface.
func (a *redisInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	parent := "projects/" + a.id.Project + "/locations/" + a.id.Location

	req := &redispb.CreateInstanceRequest{
		Parent:     parent,
		InstanceId: a.id.Instance,
		Instance:   a.desired,
	}

	log.V(0).Info("making redis CreateInstance call", "request", req)

	op, err := a.client.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for redisInstance create %s: %w", a.id.String(), err)
	}

	log.V(0).Info("created redisInstance", "redisInstance", created)

	status := &krm.RedisInstanceStatus{}

	mapCtx := &direct.MapContext{}
	if err := a.populateStatus(ctx, status, created, mapCtx); err != nil {
		return err
	}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update implements the Adapter interface.
func (a *redisInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	diffs, err := compareRedisInstance(ctx, a.actual, a.desired)
	if err != nil {
		return fmt.Errorf("comparing actual and desired RedisInstance: %w", err)
	}

	var latest *redispb.Instance
	if diffs.HasDiff() {
		diffs.Object = u
		structuredreporting.ReportDiff(ctx, diffs)

		// exactly 1 update_mask field must be specified per update request
		for _, field := range diffs.Fields {
			if field.ProtoFieldDescriptor == nil {
				return fmt.Errorf("unexpected diff field without proto descriptor: %s", field.ID)
			}
			path := string(field.ProtoFieldDescriptor.Name())

			req := &redispb.UpdateInstanceRequest{
				Instance: a.desired,
			}

			req.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: []string{path},
			}

			req.Instance.Name = a.id.String()

			log.V(0).Info("making redis UpdateInstance call", "request", req)

			op, err := a.client.UpdateInstance(ctx, req)
			if err != nil {
				return err
			}

			updated, err := op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting for redisInstance update %s: %w", a.id.String(), err)
			}
			log.V(0).Info("updated redisInstance", "redisInstance", updated)

			latest = updated
		}
	} else {
		latest = a.actual
	}

	status := &krm.RedisInstanceStatus{}
	mapCtx := &direct.MapContext{}
	if err := a.populateStatus(ctx, status, latest, mapCtx); err != nil {
		return err
	}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *redisInstanceAdapter) populateStatus(ctx context.Context, status *krm.RedisInstanceStatus, actual *redispb.Instance, mapCtx *direct.MapContext) error {
	status.CreateTime = Timestamp_FromProto(mapCtx, actual.GetCreateTime())
	if actual.GetCurrentLocationId() != "" {
		status.CurrentLocationId = direct.LazyPtr(actual.GetCurrentLocationId())
	}
	if actual.GetHost() != "" {
		status.Host = direct.LazyPtr(actual.GetHost())
	}
	if actual.GetPersistenceIamIdentity() != "" {
		status.PersistenceIamIdentity = direct.LazyPtr(actual.GetPersistenceIamIdentity())
	}
	if actual.GetPort() != 0 {
		status.Port = direct.LazyPtr(int64(actual.GetPort()))
	}
	if actual.GetReadEndpoint() != "" {
		status.ReadEndpoint = direct.LazyPtr(actual.GetReadEndpoint())
	}
	if actual.GetReadEndpointPort() != 0 {
		status.ReadEndpointPort = direct.LazyPtr(int64(actual.GetReadEndpointPort()))
	}

	if actual.GetMaintenanceSchedule() != nil {
		sched := actual.GetMaintenanceSchedule()
		status.MaintenanceSchedule = []krm.InstanceMaintenanceScheduleStatus{
			{
				EndTime:              Timestamp_FromProto(mapCtx, sched.GetEndTime()),
				ScheduleDeadlineTime: Timestamp_FromProto(mapCtx, sched.GetScheduleDeadlineTime()),
				StartTime:            Timestamp_FromProto(mapCtx, sched.GetStartTime()),
			},
		}
	}

	if actual.GetNodes() != nil {
		status.Nodes = nil
		for _, node := range actual.GetNodes() {
			status.Nodes = append(status.Nodes, krm.InstanceNodesStatus{
				Id:   direct.LazyPtr(node.GetId()),
				Zone: direct.LazyPtr(node.GetZone()),
			})
		}
	}

	if actual.GetServerCaCerts() != nil {
		status.ServerCaCerts = nil
		for _, cert := range actual.GetServerCaCerts() {
			status.ServerCaCerts = append(status.ServerCaCerts, krm.InstanceServerCaCertsStatus{
				Cert:            direct.LazyPtr(cert.GetCert()),
				CreateTime:      Timestamp_FromProto(mapCtx, cert.GetCreateTime()),
				ExpireTime:      Timestamp_FromProto(mapCtx, cert.GetExpireTime()),
				SerialNumber:    direct.LazyPtr(cert.GetSerialNumber()),
				Sha1Fingerprint: direct.LazyPtr(cert.GetSha1Fingerprint()),
			})
		}
	}

	// For ObservedState, fetch AuthString if auth is enabled
	authString := ""
	if actual.GetAuthEnabled() {
		authReq := &redispb.GetInstanceAuthStringRequest{
			Name: a.id.String(),
		}
		authResp, err := a.client.GetInstanceAuthString(ctx, authReq)
		if err != nil {
			return fmt.Errorf("getting instance auth string: %w", err)
		}
		authString = authResp.GetAuthString()
	}

	status.ObservedState = &krm.InstanceObservedStateStatus{}
	if authString != "" {
		status.ObservedState.AuthString = direct.LazyPtr(authString)
	}

	return nil
}

func compareRedisInstance(ctx context.Context, actual, desired *redispb.Instance) (*structuredreporting.Diff, error) {
	var maskedActual *redispb.Instance
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := RedisInstanceSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		maskedActual = RedisInstanceSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	maskedActual = populateInstanceDefaults(maskedActual, nil)
	desired = populateInstanceDefaults(proto.CloneOf(desired), maskedActual)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func populateInstanceDefaults(instance *redispb.Instance, actual *redispb.Instance) *redispb.Instance {
	if instance.PersistenceConfig == nil {
		instance.PersistenceConfig = &redispb.PersistenceConfig{}
	}
	if instance.PersistenceConfig.PersistenceMode == redispb.PersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
		instance.PersistenceConfig.PersistenceMode = redispb.PersistenceConfig_DISABLED
	}
	if instance.ReadReplicasMode == redispb.Instance_READ_REPLICAS_MODE_UNSPECIFIED {
		instance.ReadReplicasMode = redispb.Instance_READ_REPLICAS_DISABLED
	}

	if actual != nil {
		if instance.AuthorizedNetwork == "" && actual.AuthorizedNetwork != "" {
			instance.AuthorizedNetwork = actual.AuthorizedNetwork
		}
		if instance.LocationId == "" && actual.LocationId != "" {
			instance.LocationId = actual.LocationId
		}
		if instance.RedisVersion == "" && actual.RedisVersion != "" {
			instance.RedisVersion = actual.RedisVersion
		}
		if instance.ReservedIpRange == "" && actual.ReservedIpRange != "" {
			instance.ReservedIpRange = actual.ReservedIpRange
		}
		if (instance.SecondaryIpRange == "" || instance.SecondaryIpRange == "auto") && actual.SecondaryIpRange != "" {
			instance.SecondaryIpRange = actual.SecondaryIpRange
		}
	}

	return instance
}
