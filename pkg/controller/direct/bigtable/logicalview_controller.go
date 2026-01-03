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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	// TODO: This gcp should not be used, it takes a different proto definition than the one defined in kcc apis.
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
	registry.RegisterModel(krm.BigtableLogicalViewGVK, NewLogicalViewModel)
}

func NewLogicalViewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLogicalView{config: *config}, nil
}

var _ directbase.Model = &modelLogicalView{}

type modelLogicalView struct {
	config config.ControllerConfig
}

func (m *modelLogicalView) client(ctx context.Context, parentProject string) (*gcp.InstanceAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableLogicalView client options: %w", err)
	}
	gcpClient, err := gcp.NewInstanceAdminClient(ctx, parentProject, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableLogicalView client: %w", err)
	}
	return gcpClient, err
}

// This helper function converts a fully qualified project like "projects/myproject" into
// the unqualified project ID, like "myproject".
func (m *modelLogicalView) getProjectId(fullyQualifiedProject string) (string, error) {
	tokens := strings.Split(fullyQualifiedProject, "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return "", fmt.Errorf("unexpected format for LogicalView Parent Project ID=%q was not known (expected projects/{projectID})", fullyQualifiedProject)
	}
	return tokens[1], nil
}

func (m *modelLogicalView) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableLogicalView{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewLogicalViewIdentity(ctx, reader, obj)
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
	return &LogicalViewAdapter{
		id:        id,
		gcpClient: instanceAdminClient,
		desired:   obj,
	}, nil
}

func (m *modelLogicalView) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LogicalViewAdapter struct {
	id        *krm.LogicalViewIdentity
	gcpClient *gcp.InstanceAdminClient
	desired   *krm.BigtableLogicalView
	actual    *bigtablepb.LogicalView
}

var _ directbase.Adapter = &LogicalViewAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *LogicalViewAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableLogicalView", "name", a.id)

	logicalViewInfo, err := a.gcpClient.LogicalViewInfo(ctx, a.id.ParentInstanceIdString(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableLogicalView %q: %w", a.id, err)
	}

	deletionProtection := false
	if logicalViewInfo.DeletionProtection != 0 {
		deletionProtection = true
	}

	a.actual = &bigtablepb.LogicalView{
		Name:               a.id.String(),
		Query:              logicalViewInfo.Query,
		DeletionProtection: deletionProtection,
	}
	return true, nil
}

func convertLogicalViewToLogicalViewInfo(in *bigtablepb.LogicalView) *gcp.LogicalViewInfo {
	if in == nil {
		return nil
	}
	out := &gcp.LogicalViewInfo{}
	if in.DeletionProtection {
		out.DeletionProtection = gcp.Protected
	} else {
		out.DeletionProtection = gcp.Unprotected
	}
	out.Query = in.Query
	out.LogicalViewID = in.Name
	return out
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogicalViewAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LogicalView", "name", a.id)
	mapCtx := &direct.MapContext{}

	lv := BigtableLogicalViewSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	lv.Name = a.id.ID()
	lvi := convertLogicalViewToLogicalViewInfo(lv)
	err := a.gcpClient.CreateLogicalView(ctx, a.id.ParentInstanceIdString(), lvi)
	if err != nil {
		return fmt.Errorf("creating LogicalView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created LogicalView", "name", a.id)

	status := &krm.BigtableLogicalViewStatus{}
	// TODO: Add Observed State
	// status.ObservedState = BigtableLogicalViewObservedState_FromProto(mapCtx, created)
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
func (a *LogicalViewAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LogicalView", "name", a.id)

	spec := a.desired.Spec

	updateMask := &fieldmaskpb.FieldMask{}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	// Only set query in update mask if it's set.
	if (spec.Query != nil) && (*spec.Query != a.actual.Query) {
		report.AddField("query", a.actual.Query, spec.Query)
		updateMask.Paths = append(updateMask.Paths, "query")
	}
	// Deletion protection can either be unset (which is in itself a possible resource state), false, or true.
	if !reflect.DeepEqual(spec.DeletionProtection, a.actual.DeletionProtection) {
		report.AddField("deletion_protection", a.actual.DeletionProtection, spec.DeletionProtection)
		updateMask.Paths = append(updateMask.Paths, "deletion_protection")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", updateMask.Paths)
		structuredreporting.ReportDiff(ctx, report)

		mapCtx := &direct.MapContext{}

		lv := BigtableLogicalViewSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		lv.Name = a.id.ID()
		lvi := convertLogicalViewToLogicalViewInfo(lv)

		log.V(2).Info("Updating logical view with desired logical view", lvi)

		err := a.gcpClient.UpdateLogicalView(ctx, a.id.ParentInstanceIdString(), *lvi)
		if err != nil {
			return fmt.Errorf("updating LogicalView %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated LogicalView", "name", a.id)
		status := &krm.BigtableLogicalViewStatus{}
		status.Name = direct.LazyPtr(a.id.String())
		// TODO: Add ObservedState
		// status.ObservedState = LogicalViewObservedState_FromProto(mapCtx, updated)
		// if mapCtx.Err() != nil {
		// 	return mapCtx.Err()
		// }
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	status := &krm.BigtableLogicalViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.String())
	// TODO: Add ObservedState
	// status.ObservedState = LogicalViewObservedState_FromProto(mapCtx, updated)
	// if mapCtx.Err() != nil {
	// 	return mapCtx.Err()
	// }
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *LogicalViewAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableLogicalView{}
	mapCtx := &direct.MapContext{}
	spec := BigtableLogicalViewSpec_v1alpha1_FromProto(mapCtx, a.actual)
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
	u.SetGroupVersionKind(krm.BigtableLogicalViewGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *LogicalViewAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogicalView", "name", a.id)

	err := a.gcpClient.DeleteLogicalView(ctx, a.id.ParentInstanceIdString(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent LogicalView, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting LogicalView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LogicalView", "name", a.id)
	return true, nil
}
