// Copyright 2025 Google LLC
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

package memorystore

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/memorystore/apiv1"

	memorystorepb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.MemorystoreInstanceGVK, NewInstanceModel)
}

func NewInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstance{config: *config}, nil
}

var _ directbase.Model = &modelInstance{}

type modelInstance struct {
	config config.ControllerConfig
}

func (m *modelInstance) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Instance client: %w", err)
	}
	return gcpClient, err
}

func (m *modelInstance) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MemorystoreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get memorystore GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := MemorystoreInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Mask out some fields in the persistenceConfig, to accommodate KRM "discriminated union" semantics
	// https://github.com/kubernetes/enhancements/blob/master/keps/sig-api-machinery/1027-api-unions/README.md
	switch desired.GetPersistenceConfig().GetMode() {
	case memorystorepb.PersistenceConfig_DISABLED:
		desired.PersistenceConfig.AofConfig = nil
		desired.PersistenceConfig.RdbConfig = nil
	case memorystorepb.PersistenceConfig_RDB:
		desired.PersistenceConfig.AofConfig = nil
	case memorystorepb.PersistenceConfig_AOF:
		desired.PersistenceConfig.RdbConfig = nil
	}

	return &InstanceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type InstanceAdapter struct {
	id        *krm.InstanceIdentity
	gcpClient *gcp.Client
	desired   *memorystorepb.Instance
	actual    *memorystorepb.Instance
}

var _ directbase.Adapter = &InstanceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *InstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Instance", "name", a.id)

	req := &memorystorepb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Instance %q: %w", a.id, err)
	}

	a.actual = instancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Instance", "name", a.id)

	desired := direct.ProtoClone(a.desired)

	req := &memorystorepb.CreateInstanceRequest{
		Parent:     a.id.Parent().String(),
		Instance:   desired,
		InstanceId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Instance %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Instance %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Instance", "name", a.id)

	if err := createOp.SetLastModifiedCookie(ctx, a.desired, created); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *InstanceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *memorystorepb.Instance) error {
	mapCtx := &direct.MapContext{}
	status := &krm.MemorystoreInstanceStatus{}
	status.ObservedState = MemorystoreInstanceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Instance", "name", a.id)

	if upToDate, err := updateOp.CompareLastModifiedCookie(a.desired, a.actual); err == nil && upToDate {
		log.Info("resource is up to date (cookie match)", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs, err := compareInstance(ctx, a.actual, a.desired)
	if err != nil {
		return fmt.Errorf("comparing actual and desired Instance: %w", err)
	}

	latest := a.actual

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)

		// The memorystore UpdateInstance API allows exactly one field to be updated per request.
		for _, field := range diffs.Fields {
			if field.ProtoFieldDescriptor == nil {
				return fmt.Errorf("unexpected diff field without proto descriptor: %s", field.ID)
			}

			path := string(field.ProtoFieldDescriptor.Name())

			desired := direct.ProtoClone(a.desired)

			// Workaround: engine_version cannot be updated to empty string (API gives 400 error).
			// We don't want to upgrade/downgrade engine versions if the user didn't specify one (how would we choose a version?)
			if path == "engine_version" && desired.EngineVersion == "" {
				log.V(2).Info("skipping update for engine_version since desired value is empty", "name", a.id)
				continue
			}

			updateMask := &fieldmaskpb.FieldMask{
				Paths: []string{path},
			}
			req := &memorystorepb.UpdateInstanceRequest{
				UpdateMask: updateMask,
				Instance:   desired,
			}
			req.Instance.Name = a.id.String()

			log.Info("updating memorystore instance", "path", path, "name", a.id)

			op, err := a.gcpClient.UpdateInstance(ctx, req)
			if err != nil {
				return fmt.Errorf("updating instance %s: %w", a.id, err)
			}
			updated, err := op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("instance %s waiting update: %w", a.id, err)
			}
			latest = updated
		}

		log.V(2).Info("successfully updated Instance", "name", a.id)
	} else {
		log.Info("resource is up to date (no diff)", "name", a.id)
	}

	if err := updateOp.SetLastModifiedCookie(ctx, a.desired, latest); err != nil {
		return err
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareInstance(ctx context.Context, actual, desired *memorystorepb.Instance) (*structuredreporting.Diff, error) {
	populateDefaults := func(instance *memorystorepb.Instance) *memorystorepb.Instance {
		if instance.AuthorizationMode == memorystorepb.Instance_AUTHORIZATION_MODE_UNSPECIFIED {
			instance.AuthorizationMode = memorystorepb.Instance_AUTH_DISABLED
		}

		if instance.TransitEncryptionMode == memorystorepb.Instance_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED {
			instance.TransitEncryptionMode = memorystorepb.Instance_TRANSIT_ENCRYPTION_DISABLED
		}

		if instance.PersistenceConfig == nil {
			instance.PersistenceConfig = &memorystorepb.PersistenceConfig{}
		}
		if instance.PersistenceConfig.Mode == memorystorepb.PersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
			instance.PersistenceConfig.Mode = memorystorepb.PersistenceConfig_DISABLED
		}

		// mode is immutable, so we must default this one!
		if instance.Mode == memorystorepb.Instance_MODE_UNSPECIFIED {
			instance.Mode = memorystorepb.Instance_CLUSTER
		}

		// cannot specify nodeType = 0 in an update
		if instance.NodeType == memorystorepb.Instance_NODE_TYPE_UNSPECIFIED {
			instance.NodeType = memorystorepb.Instance_HIGHMEM_MEDIUM
		}

		// zone_distribution_config is immutable, so we must default this one!
		if instance.ZoneDistributionConfig == nil {
			instance.ZoneDistributionConfig = &memorystorepb.ZoneDistributionConfig{}
		}
		if instance.ZoneDistributionConfig.Mode == memorystorepb.ZoneDistributionConfig_ZONE_DISTRIBUTION_MODE_UNSPECIFIED {
			instance.ZoneDistributionConfig.Mode = memorystorepb.ZoneDistributionConfig_MULTI_ZONE
		}

		return instance
	}

	var maskedActual *memorystorepb.Instance
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := MemorystoreInstanceSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		maskedActual = MemorystoreInstanceSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	maskedActual = populateDefaults(maskedActual)
	desired = populateDefaults(direct.ProtoClone(desired))

	diffs, _, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MemorystoreInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MemorystoreInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MemorystoreInstanceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Instance", "name", a.id)

	req := &memorystorepb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Instance, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Instance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Instance", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Instance %s: %w", a.id, err)
	}
	return true, nil
}
