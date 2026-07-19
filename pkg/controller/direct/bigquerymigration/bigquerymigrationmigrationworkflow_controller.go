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

package bigquerymigration

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerymigration/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/bigquery/migration/apiv2alpha"
	bigquerymigrationpb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigQueryMigrationMigrationWorkflowGVK, NewMigrationWorkflowModel)
}

func NewMigrationWorkflowModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMigrationWorkflow{config: *config}, nil
}

var _ directbase.Model = &modelMigrationWorkflow{}

type modelMigrationWorkflow struct {
	config config.ControllerConfig
}

func (m *modelMigrationWorkflow) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MigrationWorkflow client: %w", err)
	}
	return gcpClient, err
}

func (m *modelMigrationWorkflow) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryMigrationMigrationWorkflow{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idInterface, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idInterface.(*krm.BigQueryMigrationMigrationWorkflowIdentity)
	if !ok {
		return nil, fmt.Errorf("identity is not of type *krm.BigQueryMigrationMigrationWorkflowIdentity")
	}

	mapCtx := &direct.MapContext{}
	desiredPb := BigQueryMigrationMigrationWorkflowSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get bigquerymigration GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &MigrationWorkflowAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelMigrationWorkflow) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.BigQueryMigrationMigrationWorkflowIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, fmt.Errorf("parsing url %q: %w", url, err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &MigrationWorkflowAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type MigrationWorkflowAdapter struct {
	id        *krm.BigQueryMigrationMigrationWorkflowIdentity
	gcpClient *gcp.Client
	desired   *bigquerymigrationpb.MigrationWorkflow
	actual    *bigquerymigrationpb.MigrationWorkflow
}

var _ directbase.Adapter = &MigrationWorkflowAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MigrationWorkflowAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MigrationWorkflow", "name", a.id)

	req := &bigquerymigrationpb.GetMigrationWorkflowRequest{Name: a.id.String()}
	migrationworkflowpb, err := a.gcpClient.GetMigrationWorkflow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MigrationWorkflow %q: %w", a.id, err)
	}

	a.actual = migrationworkflowpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MigrationWorkflowAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MigrationWorkflow", "name", a.id)
	mapCtx := &direct.MapContext{}

	req := &bigquerymigrationpb.CreateMigrationWorkflowRequest{
		Parent:            a.id.ParentString(),
		MigrationWorkflow: a.desired,
	}
	created, err := a.gcpClient.CreateMigrationWorkflow(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MigrationWorkflow %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created MigrationWorkflow", "name", a.id)

	status := &krm.BigQueryMigrationMigrationWorkflowStatus{}
	status.ObservedState = BigQueryMigrationMigrationWorkflowObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MigrationWorkflowAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MigrationWorkflow", "name", a.id)

	diffs, _, err := compareResource(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		return fmt.Errorf("BigQueryMigrationMigrationWorkflow is immutable and cannot be updated")
	}

	return nil
}

func compareResource(ctx context.Context, actual, desired *bigquerymigrationpb.MigrationWorkflow) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BigQueryMigrationMigrationWorkflowSpec_FromProto, BigQueryMigrationMigrationWorkflowSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}

	clonedDesired := proto.Clone(desired).(*bigquerymigrationpb.MigrationWorkflow)

	populateDefaults := func(obj *bigquerymigrationpb.MigrationWorkflow) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MigrationWorkflowAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryMigrationMigrationWorkflow{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryMigrationMigrationWorkflowSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Workflow)
	u.SetGroupVersionKind(krm.BigQueryMigrationMigrationWorkflowGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MigrationWorkflowAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MigrationWorkflow", "name", a.id)

	req := &bigquerymigrationpb.DeleteMigrationWorkflowRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteMigrationWorkflow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent MigrationWorkflow, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting MigrationWorkflow %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MigrationWorkflow", "name", a.id)
	return true, nil
}
