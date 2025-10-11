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
	"slices"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigtable"
	bigtablepb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableMaterializedViewGVK, NewMaterializedViewModel)
}

func NewMaterializedViewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMaterializedView{config: *config}, nil
}

var _ directbase.Model = &modelMaterializedView{}

type modelMaterializedView struct {
	config config.ControllerConfig
}

func (m *modelMaterializedView) client(ctx context.Context, parentProject string) (*gcp.InstanceAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	gcpClient, err := gcp.NewInstanceAdminClient(ctx, parentProject, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableMaterializedView client: %w", err)
	}
	return gcpClient, err
}

// This helper function converts a fully qualified project like "projects/myproject" into
// the unqualified project ID, like "myproject".
func (m *modelMaterializedView) getProjectId(fullyQualifiedProject string) (string, error) {
	tokens := strings.Split(fullyQualifiedProject, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return "", fmt.Errorf("Unexpected format for MaterializedView Parent Project ID=%q was not known (expected projects/{projectID})", fullyQualifiedProject)
	}
	return tokens[1], nil
}

func (m *modelMaterializedView) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableMaterializedView{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewMaterializedViewIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigtable instance admin GCP client. Accepts the non-fully qualified project ID.
	// E.G. "myproject" instead of "projects/myproject"
	parentProjectId, err := m.getProjectId(id.Parent().ParentString())
	if err != nil {
		return nil, err
	}
	instanceAdminClient, err := m.client(ctx, parentProjectId)
	if err != nil {
		return nil, fmt.Errorf("error creating instance admin client: %w", err)
	}
	return &MaterializedViewAdapter{
		id:        id,
		gcpClient: instanceAdminClient,
		desired:   obj,
	}, nil
}

func (m *modelMaterializedView) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MaterializedViewAdapter struct {
	id        *krm.MaterializedViewIdentity
	gcpClient *gcp.InstanceAdminClient
	desired   *krm.BigtableMaterializedView
	actual    *bigtablepb.MaterializedView
}

var _ directbase.Adapter = &MaterializedViewAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MaterializedViewAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableMaterializedView", "name", a.id)

	MaterializedViewInfo, err := a.gcpClient.MaterializedViewInfo(ctx, a.id.ParentInstanceIdString(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableMaterializedView %q: %w", a.id, err)
	}

	deletionProtection := false
	if MaterializedViewInfo.DeletionProtection == gcp.Protected {
		deletionProtection = true
	}

	a.actual = &bigtablepb.MaterializedView{
		Name:               a.id.String(),
		Query:              MaterializedViewInfo.Query,
		DeletionProtection: deletionProtection,
	}
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MaterializedViewAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MaterializedView", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigtableMaterializedViewSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	spec := a.desired.Spec

	gcpDeletionProtection := gcp.None
	if spec.DeletionProtection != nil {
		if *spec.DeletionProtection {
			gcpDeletionProtection = gcp.Protected
		} else {
			gcpDeletionProtection = gcp.Unprotected
		}
	}

	MaterializedViewInfo := &gcp.MaterializedViewInfo{
		MaterializedViewID: a.id.ID(),
		Query:              resource.Query,
		DeletionProtection: gcpDeletionProtection,
	}
	err := a.gcpClient.CreateMaterializedView(ctx, a.id.ParentInstanceIdString(), MaterializedViewInfo)
	if err != nil {
		return fmt.Errorf("creating MaterializedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created MaterializedView", "name", a.id)

	status := &krm.BigtableMaterializedViewStatus{}
	// TODO: Add Observed State
	// status.ObservedState = BigtableMaterializedViewObservedState_FromProto(mapCtx, created)
	// if mapCtx.Err() != nil {
	// 	return mapCtx.Err()
	// }
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.String())
	if err := createOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}

	// Write resourceID into spec.
	if err := unstructured.SetNestedField(createOp.GetUnstructured().Object, a.id.ID(), "spec", "resourceID"); err != nil {
		return fmt.Errorf("error setting spec.resourceID: %w", err)
	}
	return nil
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MaterializedViewAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MaterializedView", "name", a.id)

	spec := a.desired.Spec
	updateMask := &fieldmaskpb.FieldMask{}
	if (spec.DeletionProtection != nil) && (*spec.DeletionProtection != a.actual.DeletionProtection) {
		updateMask.Paths = append(updateMask.Paths, "deletion_protection")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", updateMask.Paths)

		desiredmaterializedview := gcp.MaterializedViewInfo{
			MaterializedViewID: a.id.ID(),
			Query:              a.actual.Query,
		}

		if slices.Contains(updateMask.Paths, "deletion_protection") {
			gcpDeletionProtection := gcp.None
			if a.desired.Spec.DeletionProtection != nil {
				if *a.desired.Spec.DeletionProtection {
					gcpDeletionProtection = gcp.Protected
				} else {
					gcpDeletionProtection = gcp.Unprotected
				}
			}
			desiredmaterializedview.DeletionProtection = gcpDeletionProtection
		}

		err := a.gcpClient.UpdateMaterializedView(ctx, a.id.ParentInstanceIdString(), desiredmaterializedview)
		if err != nil {
			return fmt.Errorf("updating MaterializedView %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated MaterializedView", "name", a.id)
		status := &krm.BigtableMaterializedViewStatus{}
		status.Name = direct.LazyPtr(a.id.String())
		// TODO: Add ObservedState
		// status.ObservedState = MaterializedViewObservedState_FromProto(mapCtx, updated)
		// if mapCtx.Err() != nil {
		// 	return mapCtx.Err()
		// }
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	status := &krm.BigtableMaterializedViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.String())
	// TODO: Add ObservedState
	// status.ObservedState = MaterializedViewObservedState_FromProto(mapCtx, updated)
	// if mapCtx.Err() != nil {
	// 	return mapCtx.Err()
	// }
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MaterializedViewAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableMaterializedView{}
	mapCtx := &direct.MapContext{}
	spec := BigtableMaterializedViewSpec_FromProto(mapCtx, a.actual)
	obj.Spec = direct.ValueOf(spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceRef = &krmv1beta1.InstanceRef{External: a.id.ParentInstanceIdString()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableMaterializedViewGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MaterializedViewAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MaterializedView", "name", a.id)

	err := a.gcpClient.DeleteMaterializedView(ctx, a.id.ParentInstanceIdString(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MaterializedView, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting MaterializedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MaterializedView", "name", a.id)
	return true, nil
}
