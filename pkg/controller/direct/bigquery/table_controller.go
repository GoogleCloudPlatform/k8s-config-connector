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

	bigquery "google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "bigquery-controller"
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
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTableIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	projectID := id.Parent().ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project ID")
	}

	// Create bigquery GCP client for BQ Table.
	gcpService, err := m.tableService(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:         id,
		gcpService: gcpService,
		desired:    obj,
		reader:     reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id         *krm.TableIdentity
	gcpService *bigquery.TablesService
	desired    *krm.BigQueryTable
	actual     *bigquery.Table
	reader     client.Reader
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
	if desired.Spec.EncryptionConfiguration != nil && desired.Spec.EncryptionConfiguration.KmsKeyRef != nil {
		kmsRef, err := refs.ResolveKMSCryptoKeyRef(ctx, a.reader, a.desired, a.desired.Spec.EncryptionConfiguration.KmsKeyRef)
		if err != nil {
			return err
		}
		table.EncryptionConfiguration.KmsKeyName = kmsRef.External
	}
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
		log.V(2).Info("Updating BigQuery Table back with schema", "name", res.Id)
		res, err = a.gcpService.Update(parent.ProjectID, parent.DatasetID, res.Id, table).Do()
		if err != nil {
			return fmt.Errorf("error updating Table %s: %w", a.id.ID(), err)
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

	eq, err := TableEq(a.actual, table)
	if err != nil {
		return err
	}
	// No diff detected.
	if eq {
		return nil
	}

	a.customTableLogic(table)
	parent := a.id.Parent()

	if mapCtx.Err() != nil {
		return fmt.Errorf("error updating Table %s: %w", a.id.ID(), mapCtx.Err())
	}
	if a.desired.ObjectMeta.Annotations != nil {
		unmanaged, ok := a.desired.ObjectMeta.Annotations["cnrm.cloud.google.com/unmanaged"]
		if ok && unmanaged != "" {
			unmanagedFields := strings.Split(unmanaged, ",")
			for _, field := range unmanagedFields {
				// This ability is only intended for spec.schema field for the moment.
				if field == "spec.schema" {
					table.Schema = nil
				}
			}
			// Make PATCH call with nil schema to avoid schema being updated.
			res, err := a.gcpService.Patch(parent.ProjectID, parent.DatasetID, a.id.ID(), table).Do()
			if err != nil {
				return fmt.Errorf("error updating Table %s: %w", a.id.ID(), err)
			}
			return a.UpdateStatusForUpdate(ctx, updateOp, res)
		}
	}
	res, err := a.gcpService.Update(parent.ProjectID, parent.DatasetID, a.id.ID(), table).Do()
	if err != nil {
		return fmt.Errorf("error updating Table %s: %w", a.id.ID(), err)
	}
	return a.UpdateStatusForUpdate(ctx, updateOp, res)
}

func (a *Adapter) UpdateStatusForUpdate(ctx context.Context, updateOp *directbase.UpdateOperation, created *bigquery.Table) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryTableStatus_FromProto(mapCtx, created)
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
		return false, fmt.Errorf("deleting Table %s: %w", a.id, err)
	}
	return true, nil
}

func (a *Adapter) customTableLogic(table *bigquery.Table) {
	if table == nil {
		return
	}
	parent := a.id.Parent()
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
