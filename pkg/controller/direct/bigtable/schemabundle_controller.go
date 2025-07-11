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
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigtable"
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableSchemaBundleGVK, NewBigtableSchemaBundleModel)
}

func NewBigtableSchemaBundleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigtableSchemaBundle{config: *config}, nil
}

var _ directbase.Model = &modelBigtableSchemaBundle{}

type modelBigtableSchemaBundle struct {
	config config.ControllerConfig
}

func (m *modelBigtableSchemaBundle) client(ctx context.Context, parentProject string) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	gcpClient, err := gcp.NewAdminClient(ctx, parentProject, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableSchemaBundle client: %w", err)
	}
	return gcpClient, err
}

// This helper function converts a fully qualified project like "projects/myproject" into
// the unqualified project ID, like "myproject".
func (m *modelBigtableSchemaBundle) getProjectId(fullyQualifiedProject string) (string, error) {
	tokens := strings.Split(fullyQualifiedProject, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return "", fmt.Errorf("Unexpected format for SchemaBundle Parent Project ID=%q was not known (expected projects/{projectID})", fullyQualifiedProject)
	}
	return tokens[1], nil
}

func (m *modelBigtableSchemaBundle) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableSchemaBundle{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSchemaBundleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigtable admin GCP client. Accepts the non-fully qualified project ID.
	// E.G. "myproject" instead of "projects/myproject"
	parentProjectId, err := m.getProjectId(id.Parent().ParentString())
	if err != nil {
		return nil, err
	}
	adminClient, err := m.client(ctx, parentProjectId)
	if err != nil {
		return nil, fmt.Errorf("error creating admin client: %w", err)
	}
	return &BigtableSchemaBundleAdapter{
		id:        id,
		gcpClient: adminClient,
		desired:   obj,
	}, nil
}

func (m *modelBigtableSchemaBundle) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BigtableSchemaBundleAdapter struct {
	id        *krm.SchemaBundleIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableSchemaBundle
	actual    *bigtablepb.SchemaBundle
}

var _ directbase.Adapter = &BigtableSchemaBundleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BigtableSchemaBundleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableSchemaBundle", "name", a.id)

	bigtableschemabundlepb, err := a.gcpClient.GetSchemaBundle(ctx, a.id.ParentTableIdString(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableSchemaBundle %q: %w", a.id, err)
	}

	a.actual = bigtableschemabundlepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BigtableSchemaBundleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableSchemaBundle", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableSchemaBundleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var routingConfig gcp.RoutingPolicyConfig
	if multiClusterRouting := resource.GetMultiClusterRoutingUseAny(); multiClusterRouting != nil {
		routingConfig = &gcp.MultiClusterRoutingUseAnyConfig{
			ClusterIDs: resource.GetMultiClusterRoutingUseAny().ClusterIds,
		}
	} else {
		routingConfig = &gcp.SingleClusterRoutingConfig{
			ClusterID:                resource.GetSingleClusterRouting().GetClusterId(),
			AllowTransactionalWrites: resource.GetSingleClusterRouting().AllowTransactionalWrites,
		}
	}
	var isolation gcp.SchemaBundleIsolation
	if dataBoostIsolationReadOnly := resource.GetDataBoostIsolationReadOnly(); dataBoostIsolationReadOnly != nil {
		isolation = &gcp.DataBoostIsolationReadOnly{
			ComputeBillingOwner: gcp.IsolationComputeBillingOwner(resource.GetDataBoostIsolationReadOnly().GetComputeBillingOwner()),
		}
	} else if standardIsolation := resource.GetStandardIsolation(); standardIsolation != nil {
		isolation = &gcp.StandardIsolation{
			Priority: gcp.SchemaBundlePriority(resource.GetStandardIsolation().GetPriority()),
		}
	}
	if isolation == nil {
		isolation = &gcp.StandardIsolation{
			Priority: gcp.SchemaBundlePriority(bigtablepb.SchemaBundle_PRIORITY_HIGH),
		}
	}
	profileConf := &gcp.SchemaBundleConf{
		Name:          "", /*Name is not used in the RPC*/
		ProfileID:     a.id.ID(),
		TableID:       a.id.ParentTableIdString(),
		Etag:          resource.Etag,
		Description:   resource.Description,
		RoutingConfig: routingConfig,
		Isolation:     isolation,
	}

	// TODO: Make use of returned schema bundle in ObservedState below
	_, err := a.gcpClient.CreateSchemaBundle(ctx, *profileConf)
	if err != nil {
		return fmt.Errorf("creating BigtableSchemaBundle %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableSchemaBundle", "name", a.id)

	status := &krm.BigtableSchemaBundleStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BigtableSchemaBundleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableSchemaBundle", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableSchemaBundleSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO: Add updated profile to ObservedState below after reconciling what the new profile looks like
	// var updated *bigtablepb.SchemaBundle
	// updated = a.actual
	var fieldsToUpdate gcp.ProfileAttrsToUpdate
	// Several field changes (single -> multi cluster; turning on data boost) require us to ignore warnings if we want the update to go through.
	hasChanges := false

	if desired.Spec.DataBoostIsolationReadOnly != nil && !cmp.Equal(resource.GetDataBoostIsolationReadOnly(), a.actual.GetDataBoostIsolationReadOnly(), cmpopts.IgnoreUnexported(bigtablepb.SchemaBundle_DataBoostIsolationReadOnly{})) {
		fieldsToUpdate.Isolation = &gcp.DataBoostIsolationReadOnly{
			ComputeBillingOwner: gcp.IsolationComputeBillingOwner(resource.GetDataBoostIsolationReadOnly().GetComputeBillingOwner()),
		}
		hasChanges = true
	}

	if !hasChanges {
		log.V(2).Info("no changes to update", "name", a.id)
	} else {
		err := a.gcpClient.UpdateSchemaBundle(ctx, a.id.ParentTableIdString(), a.id.ID(), fieldsToUpdate)
		if err != nil {
			return fmt.Errorf("updating BigtableSchemaBundle %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated BigtableSchemaBundle", "name", a.id)
	}

	status := &krm.BigtableSchemaBundleStatus{}
	status.Name = direct.LazyPtr(a.id.String())
	// TODO: Add ObservedState
	// status.ObservedState = SchemaBundleObservedState_FromProto(mapCtx, updated)
	// if mapCtx.Err() != nil {
	// 	return mapCtx.Err()
	// }
	return updateOp.UpdateStatus(ctx, status, nil)

}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BigtableSchemaBundleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableSchemaBundle{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigtableSchemaBundleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceRef = &krm.InstanceRef{External: a.id.ParentInstanceIdString()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableSchemaBundleGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BigtableSchemaBundleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableSchemaBundle", "name", a.id)

	err := a.gcpClient.DeleteSchemaBundle(ctx, a.id.ParentTableIdString(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BigtableSchemaBundle, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableSchemaBundle %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableSchemaBundle", "name", a.id)

	return true, nil
}
