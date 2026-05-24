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

package biglakedatabase

import (
	"context"
	"fmt"

	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigquerybiglake"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/bigquery/biglake/apiv1"
	bigquerybiglakepb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krmv1alpha1.BigLakeDatabaseGVK, NewDatabaseModel)
}

func NewDatabaseModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDatabase{config: *config}, nil
}

var _ directbase.Model = &modelDatabase{}

type modelDatabase struct {
	config config.ControllerConfig
}

func (m *modelDatabase) client(ctx context.Context) (*gcp.MetastoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewMetastoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Database client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDatabase) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krmv1alpha1.BigLakeDatabase{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := bigquerybiglake.BigLakeDatabaseSpec_v1alpha1_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get bigquerybiglake GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DatabaseAdapter{
		id:        id.(*krmv1alpha1.DatabaseIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelDatabase) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DatabaseAdapter struct {
	id        *krmv1alpha1.DatabaseIdentity
	gcpClient *gcp.MetastoreClient
	desired   *bigquerybiglakepb.Database
	actual    *bigquerybiglakepb.Database
}

var _ directbase.Adapter = &DatabaseAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DatabaseAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Database", "name", a.id)

	req := &bigquerybiglakepb.GetDatabaseRequest{Name: a.id.String()}
	databasepb, err := a.gcpClient.GetDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Database %q: %w", a.id, err)
	}

	a.actual = databasepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DatabaseAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Database", "name", a.id)
	req := &bigquerybiglakepb.CreateDatabaseRequest{
		Parent:     a.id.Parent().String(),
		Database:   a.desired,
		DatabaseId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Database %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created Database", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krmv1alpha1.BigLakeDatabaseStatus{}
	status.ObservedState = bigquerybiglake.BigLakeDatabaseObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DatabaseAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Database", "name", a.id)

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}
	a.desired.Name = a.id.String()

	req := &bigquerybiglakepb.UpdateDatabaseRequest{
		UpdateMask: updateMask,
		Database:   a.desired,
	}
	updated, err := a.gcpClient.UpdateDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Database %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated Database", "name", a.id)

	mapCtx := &direct.MapContext{}

	status := &krmv1alpha1.BigLakeDatabaseStatus{}
	status.ObservedState = bigquerybiglake.BigLakeDatabaseObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DatabaseAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krmv1alpha1.BigLakeDatabase{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(bigquerybiglake.BigLakeDatabaseSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	externalRef := a.actual.GetName()
	var id krmv1alpha1.DatabaseIdentity
	if err := id.FromExternal(externalRef); err != nil {
		return nil, fmt.Errorf("parsing external ref %q: %w", externalRef, err)
	}
	obj.Spec.ParentRef = &krmv1alpha1.BigQueryBigLakeCatalogRef{
		External: id.Parent().String(),
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krmv1alpha1.BigLakeDatabaseGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DatabaseAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Database", "name", a.id)

	req := &bigquerybiglakepb.DeleteDatabaseRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Database, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Database %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Database", "name", a.id)

	return true, nil
}
