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

package bigtable

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/bigtable"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableAppProfileGVK, NewBigtableAppProfileModel)
}

func NewBigtableAppProfileModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigtableAppProfile{config: *config}, nil
}

var _ directbase.Model = &modelBigtableAppProfile{}

type modelBigtableAppProfile struct {
	config config.ControllerConfig
}

func (m *modelBigtableAppProfile) client(ctx context.Context, parent *krm.AppProfileParent) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	tokens := strings.Split(parent.BigtableInstance, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "instances" {
		return nil, fmt.Errorf("invalid format for BigtableInstance=%q. Expected: projects/{{projectID}}/instances/{{instance}}", parent.BigtableInstance)
	}

	projectID := tokens[1]
	instanceID := tokens[3]
	gcpClient, err := gcp.NewClient(ctx, projectID, instanceID, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableAppProfile client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBigtableAppProfile) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableAppProfile{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAppProfileIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigtable GCP client
	gcpClient, err := m.client(ctx, id.Parent())
	if err != nil {
		return nil, err
	}
	return &BigtableAppProfileAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBigtableAppProfile) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BigtableAppProfileAdapter struct {
	id        *krm.AppProfileIdentity
	gcpClient *gcp.Client
	desired   *krm.BigtableAppProfile
	actual    *bigtablepb.AppProfile
}

var _ directbase.Adapter = &BigtableAppProfileAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BigtableAppProfileAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableAppProfile", "name", a.id)

	bigtableappprofilepb, err := a.gcpClient.GetAppProfile(ctx, a.id.Parent().BigtableInstance, a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableAppProfile %q: %w", a.id, err)
	}

	a.actual = bigtableappprofilepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BigtableAppProfileAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableAppProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableAppProfileSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	appProfileConf := &gcp.ProfileConf{
		Name:                     a.id.ID(),
		ProfileID:                a.id.String(),
		Etag:                     resource.Etag,
		Description:              resource.Description,
		RoutingPolicy:            resource.RoutingPolicy,
		AllowTransactionalWrites: resource.GetSingleClusterRouting().AllowTransactionalWrites,
		// RoutingPolicy            string
		// ClusterID                string
	}

	op, err := a.gcpClient.CreateAppProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigtableAppProfile %s: %w", a.id, err)
	}
	// TODO: make use of the created resource for Observed State
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("BigtableAppProfile %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableAppProfile", "name", a.id)

	status := &krm.BigtableAppProfileStatus{}
	// TODO: Add Observed State
	// status.ObservedState = BigtableAppProfileObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BigtableAppProfileAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableAppProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableAppProfileSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])

	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = paths.Insert("description")
	}
	if desired.Spec.MultiClusterRoutingUseAny != nil && !reflect.DeepEqual(resource.GetRoutingPolicy(), a.actual.GetRoutingPolicy()) {
		paths = paths.Insert("multi_cluster_routing_use_any")
	}
	if desired.Spec.MultiClusterRoutingClusterIds != nil && !reflect.DeepEqual(resource.GetMultiClusterRoutingUseAny().ClusterIds, a.actual.GetMultiClusterRoutingUseAny().ClusterIds) {
		paths = paths.Insert("multi_cluster_routing_use_any.cluster_ids")
	}
	if desired.Spec.SingleClusterRouting != nil && !reflect.DeepEqual(resource.GetSingleClusterRouting(), a.actual.GetSingleClusterRouting()) {
		paths = paths.Insert("single_cluster_routing")
	}
	if desired.Spec.StandardIsolation != nil && !reflect.DeepEqual(resource.GetStandardIsolation(), a.actual.GetStandardIsolation()) {
		paths = paths.Insert("isolation.standard_isolation")
	}

	// TODO: Add updated profile to ObservedState below
	// var updated *bigtablepb.AppProfile
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// updated = a.actual
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		req := &bigtablepb.UpdateAppProfileRequest{
			AppProfile: resource,
			UpdateMask: updateMask,
		}
		op, err := a.gcpClient.UpdateBigtableAppProfile(ctx, req)
		if err != nil {
			return fmt.Errorf("updating BigtableAppProfile %s: %w", a.id, err)
		}
		// TODO: Make use of the updated resource in ObservedState below
		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("BigtableAppProfile %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated BigtableAppProfile", "name", a.id)
	}

	status := &krm.BigtableAppProfileStatus{}
	// TODO: Add ObservedState
	// status.ObservedState = AppProfileObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BigtableAppProfileAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableAppProfile{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigtableAppProfileSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceRef = &krm.InstanceRef{External: a.id.Parent().BigtableInstance}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableAppProfileGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BigtableAppProfileAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableAppProfile", "name", a.id)

	req := &bigtablepb.DeleteAppProfileRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAppProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BigtableAppProfile, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableAppProfile %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableAppProfile", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BigtableAppProfile %s: %w", a.id, err)
	}
	return true, nil
}
