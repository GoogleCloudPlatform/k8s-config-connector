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

package bigtable

import (
	"context"
	"fmt"
	"reflect"

	bigtable "cloud.google.com/go/bigtable/admin/apiv2"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableTableGVK, NewBigtableTableModel)
}

func NewBigtableTableModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigtableTable{config: *config}, nil
}

var _ directbase.Model = &modelBigtableTable{}

type modelBigtableTable struct {
	config config.ControllerConfig
}

func (m *modelBigtableTable) client(ctx context.Context) (*bigtable.BigtableTableAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableTable client options: %w", err)
	}
	gcpClient, err := bigtable.NewBigtableTableAdminClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableTable admin client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelBigtableTable) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identity.(*krm.BigtableTableIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identity)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &BigtableTableAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelBigtableTable) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type BigtableTableAdapter struct {
	id        *krm.BigtableTableIdentity
	gcpClient *bigtable.BigtableTableAdminClient
	desired   *krm.BigtableTable
	reader    client.Reader
	actual    *pb.Table
}

var _ directbase.Adapter = &BigtableTableAdapter{}

func (a *BigtableTableAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *BigtableTableAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableTable", "name", a.id)

	req := &pb.GetTableRequest{
		Name: a.id.String(),
		View: pb.Table_SCHEMA_VIEW,
	}

	table, err := a.gcpClient.GetTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableTable %q: %w", a.id, err)
	}

	a.actual = table
	return true, nil
}

func (a *BigtableTableAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableTable", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredTable := BigtableTableSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Prepare initial splits
	var initialSplits []*pb.CreateTableRequest_Split
	for _, key := range a.desired.Spec.SplitKeys {
		initialSplits = append(initialSplits, &pb.CreateTableRequest_Split{
			Key: []byte(key),
		})
	}

	parent := fmt.Sprintf("projects/%s/instances/%s", a.id.Project, a.id.Instance)
	req := &pb.CreateTableRequest{
		Parent:        parent,
		TableId:       a.id.Table,
		Table:         desiredTable,
		InitialSplits: initialSplits,
	}

	created, err := a.gcpClient.CreateTable(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigtableTable %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created BigtableTable", "name", a.id)
	a.actual = created

	return a.updateStatus(ctx, createOp)
}

func (a *BigtableTableAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableTable", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredTable := BigtableTableSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Check column family modifications (create/drop/update)
	var mods []*pb.ModifyColumnFamiliesRequest_Modification
	actualCFs := make(map[string]*pb.ColumnFamily)
	if a.actual != nil {
		actualCFs = a.actual.GetColumnFamilies()
	}
	desiredCFs := desiredTable.GetColumnFamilies()

	// Drop removed column families
	for cfName := range actualCFs {
		if _, exists := desiredCFs[cfName]; !exists {
			mods = append(mods, &pb.ModifyColumnFamiliesRequest_Modification{
				Id:  cfName,
				Mod: &pb.ModifyColumnFamiliesRequest_Modification_Drop{Drop: true},
			})
		}
	}

	// Add or update column families
	for cfName, desiredCF := range desiredCFs {
		actualCF, exists := actualCFs[cfName]
		if !exists {
			mods = append(mods, &pb.ModifyColumnFamiliesRequest_Modification{
				Id:  cfName,
				Mod: &pb.ModifyColumnFamiliesRequest_Modification_Create{Create: desiredCF},
			})
		} else if !reflect.DeepEqual(actualCF, desiredCF) {
			mods = append(mods, &pb.ModifyColumnFamiliesRequest_Modification{
				Id:  cfName,
				Mod: &pb.ModifyColumnFamiliesRequest_Modification_Update{Update: desiredCF},
			})
		}
	}

	if len(mods) > 0 {
		req := &pb.ModifyColumnFamiliesRequest{
			Name:          a.id.String(),
			Modifications: mods,
		}
		updatedTable, err := a.gcpClient.ModifyColumnFamilies(ctx, req)
		if err != nil {
			return fmt.Errorf("modifying column families for BigtableTable %s: %w", a.id, err)
		}
		a.actual = updatedTable
	}

	// Check deletion protection update if changed
	if a.actual.GetDeletionProtection() != desiredTable.GetDeletionProtection() {
		req := &pb.UpdateTableRequest{
			Table: &pb.Table{
				Name:               a.id.String(),
				DeletionProtection: desiredTable.GetDeletionProtection(),
			},
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"deletion_protection"}},
		}
		op, err := a.gcpClient.UpdateTable(ctx, req)
		if err != nil {
			return fmt.Errorf("updating deletion protection for BigtableTable %s: %w", a.id, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for deletion protection update on BigtableTable %s: %w", a.id, err)
		}
		a.actual = updated
	}

	return a.updateStatus(ctx, updateOp)
}

func (a *BigtableTableAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableTable", "name", a.id)

	req := &pb.DeleteTableRequest{
		Name: a.id.String(),
	}

	err := a.gcpClient.DeleteTable(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableTable %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted BigtableTable", "name", a.id)
	return true, nil
}

func (a *BigtableTableAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	mapCtx := &direct.MapContext{}
	_ = BigtableTableObservedState_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.BigtableTableStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}
