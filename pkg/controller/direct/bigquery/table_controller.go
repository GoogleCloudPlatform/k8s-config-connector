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

package bigquery

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	bigquery "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName                            = "bigquery-controller"
	unmanageDataPoliciesAnnotationValue = "spec.schema.fields.dataPolicies"
	unmanagePolicyTagsAnnotationValue   = "spec.schema.fields.policyTags"
)

func init() {
	rg := &TableReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.BigQueryTableGVK, NewModel, rg)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type TableReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &TableReconcileGate{}

func (r *TableReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	if r.optIn.ShouldReconcile(o) {
		return true
	}
	if _, ok := o.GetAnnotations()[kccpredicate.AnnotationUnmanaged]; ok {
		return true
	}
	obj := &krm.BigQueryTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &obj); err != nil {
		return false
	}
	return obj.Spec.Labels != nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) tableService(ctx context.Context) (*bigquery.TablesService, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpService, err := bigquery.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Table client: %w", err)
	}
	return gcpService.Tables, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryTable{}
	copied := u.DeepCopy()
	if err := label.ComputeLabels(copied); err != nil {
		return nil, err
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(copied.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTableIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	// Create bigquery GCP client for BQ Table.
	gcpService, err := m.tableService(ctx)
	if err != nil {
		return nil, err
	}
	adapter := &Adapter{
		id:         id,
		gcpService: gcpService,
		desired:    obj,
		reader:     reader,
	}
	unmanaged, ok := obj.GetAnnotations()[kccpredicate.AnnotationUnmanaged]
	if ok && unmanaged != "" {
		unmanagedFields := strings.Split(unmanaged, ",")
		for _, field := range unmanagedFields {
			if field != unmanageDataPoliciesAnnotationValue && field != unmanagePolicyTagsAnnotationValue {
				return nil, fmt.Errorf("unmanaging field `%s` is not supported", field)
			}
		}
		adapter.unmanagedFields = unmanagedFields
	}
	return adapter, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id              *krm.TableIdentity
	gcpService      *bigquery.TablesService
	desired         *krm.BigQueryTable
	actual          *bigquery.Table
	reader          client.Reader
	unmanagedFields []string
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting BigQueryTable", "name", a.id.String())
	parent := a.id.Parent()
	getCall := a.gcpService.Get(parent.ProjectID, parent.DatasetID, a.id.ID())
	table, err := getCall.Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryTable %q: %w", a.id.String(), err)
	}
	a.actual = table
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Table", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing reference for creation: %w", err)
	}
	desired := a.desired.DeepCopy()
	table := BigQueryTableSpec_ToProto(mapCtx, &desired.Spec)

	a.customTableLogic(table)
	parent := a.id.Parent()
	if table.View != nil && table.Schema != nil {
		log.V(2).Info("Removing schema from table definition because big query does not support setting schema on view creation")
		schemaBack := table.Schema
		table.Schema = nil
		log.V(2).Info("Creating BigQuery without schema", "name", table.TableReference.TableId)
		res, err := a.gcpService.Insert(parent.ProjectID, parent.DatasetID, table).Do()
		if err != nil {
			return fmt.Errorf("error creating Table %s: %w", a.id.ID(), err)
		}
		log.V(2).Info("successfully created Table", "name", res.Id)
		table.Schema = schemaBack
		log.V(2).Info("Updating BigQuery Table back with schema", "name", res.TableReference.TableId)
		res, err = a.gcpService.Update(parent.ProjectID, parent.DatasetID, res.TableReference.TableId, table).Do()
		if err != nil {
			return fmt.Errorf("error updating Table during CREATE table with both schema and view specified %s: %w", a.id.ID(), err)
		}
		log.V(2).Info("successfully updated Table with schema", "name", res.Id)
		if err := a.UpdateStatusForCreate(ctx, createOp, res); err != nil {
			return err
		}
	} else {
		insertCall := a.gcpService.Insert(parent.ProjectID, parent.DatasetID, table)
		created, err := insertCall.Do()
		if err != nil {
			return fmt.Errorf("error creating Table %s: %w", a.id.ID(), err)
		}
		log.V(2).Info("successfully created Table", "name", a.id.String())
		if err := a.UpdateStatusForCreate(ctx, createOp, created); err != nil {
			return err
		}
	}
	return nil
}

func (a *Adapter) UpdateStatusForCreate(ctx context.Context, createOp *directbase.CreateOperation, created *bigquery.Table) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryTableStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	external := a.id.String()
	status.ExternalRef = &external
	if err := createOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}
	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Table", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing reference for update: %w", err)
	}

	desired := a.desired.DeepCopy()
	table := BigQueryTableSpec_ToProto(mapCtx, &desired.Spec)

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// For table created from external data configuration,
	// If no schema was provided, the Update request would fail.
	// Getting schema from actual object.
	if table.ExternalDataConfiguration != nil && table.Schema == nil && table.ExternalDataConfiguration.Schema == nil {
		table.Schema = a.actual.Schema
		for _, field := range table.Schema.Fields {
			setEmptyPolicyTagsInSchema(field)
		}
	}
	// Make fields unmanaged before performing desired vs actual comparison.
	if len(a.unmanagedFields) > 0 {
		makeFieldsUnmanaged(table, a.unmanagedFields)
		makeFieldsUnmanaged(a.actual, a.unmanagedFields)
	}
	eq, err := TableEq(a.actual, table)

	if err != nil {
		return err
	}
	// No diff detected.
	if eq {
		log.Info("no diff detected for Table", "name", a.id)
		return a.UpdateStatusForUpdate(ctx, updateOp, a.actual)
	}

	a.customTableLogic(table)
	parent := a.id.Parent()

	if len(a.unmanagedFields) > 0 {
		// Make PATCH call without unmanaged fields to avoid accidental override.
		res, err := a.gcpService.Patch(parent.ProjectID, parent.DatasetID, a.id.ID(), table).Do()
		if err != nil {
			return fmt.Errorf("error patching Table %s: %w", a.id.ID(), err)
		}
		return a.UpdateStatusForUpdate(ctx, updateOp, res)
	}
	// If not unmanaged field is specify, we use PUT to keep backward compatibility with TF controller.
	res, err := a.gcpService.Update(parent.ProjectID, parent.DatasetID, a.id.ID(), table).Do()
	if err != nil {
		return fmt.Errorf("error updating Table %s: %w", a.id.ID(), err)
	}
	return a.UpdateStatusForUpdate(ctx, updateOp, res)
}

func makeFieldsUnmanaged(table *bigquery.Table, unmanagedFields []string) {
	if table == nil {
		return
	}
	for _, fieldName := range unmanagedFields {
		if fieldName == unmanagePolicyTagsAnnotationValue {
			unmanagePolicyTags(table)
		}
		if fieldName == unmanageDataPoliciesAnnotationValue {
			unmanageDataPolicies(table)
		}
	}
}

func unmanagePolicyTags(table *bigquery.Table) {
	if table != nil && table.Schema != nil {
		setEmptyPolicyTags(table.Schema.Fields)
	}
}

// Recursively set PolicyTags to nil for each fields and subfields.
func setEmptyPolicyTags(fields []*bigquery.TableFieldSchema) {
	for _, field := range fields {
		field.PolicyTags = nil
		setEmptyPolicyTags(field.Fields)
	}
}

func unmanageDataPolicies(table *bigquery.Table) {
	if table != nil && table.Schema != nil {
		setEmptyDataPolicies(table.Schema.Fields)
	}
}

// Recursively set DataPolicies to nil for each fields and subfields.
func setEmptyDataPolicies(fields []*bigquery.TableFieldSchema) {
	for _, field := range fields {
		field.DataPolicies = nil
		setEmptyDataPolicies(field.Fields)
	}
}

func (a *Adapter) UpdateStatusForUpdate(ctx context.Context, updateOp *directbase.UpdateOperation, updated *bigquery.Table) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryTableStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	external := a.id.String()
	status.ExternalRef = &external
	if err := updateOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}
	return nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryTable{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryTableSpec_FromProto(mapCtx, a.actual))
	// Populate the reference from the identity
	obj.Spec.DatasetRef = &krm.DatasetRef{External: a.id.Parent().String()}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.BigQueryTableGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting Table", "name", a.id.String())
	parent := a.id.Parent()
	if err := a.gcpService.Delete(parent.ProjectID, parent.DatasetID, a.id.ID()).Do(); err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BigQuery Table, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Table %s: %w", a.id, err)
	}
	return true, nil
}

// customTableLogic added extra logics for the table obj before making the API call.
// These extra logics could not be handled in the mapper functions.
func (a *Adapter) customTableLogic(table *bigquery.Table) {
	if table == nil {
		return
	}
	parent := a.id.Parent()
	// Adding TableReference object as the API requires this object to be populated to be valid.
	table.TableReference = &bigquery.TableReference{
		ProjectId: parent.ProjectID,
		DatasetId: parent.DatasetID,
		TableId:   a.id.ID(),
	}
}

func (a *Adapter) normalizeReferences(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.EncryptionConfiguration != nil && obj.Spec.EncryptionConfiguration.KmsKeyRef != nil {
		key, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, obj, obj.Spec.EncryptionConfiguration.KmsKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.EncryptionConfiguration.KmsKeyRef = key
	}
	return nil
}
