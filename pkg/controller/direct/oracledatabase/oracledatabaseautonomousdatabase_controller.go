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

package oracledatabase

import (
	"context"
	"fmt"
	"slices"

	oracledatabase "cloud.google.com/go/oracledatabase/apiv1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.OracleDatabaseAutonomousDatabaseGVK, NewAutonomousDatabaseModel)
}

func NewAutonomousDatabaseModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAutonomousDatabase{config: config}, nil
}

type modelAutonomousDatabase struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &modelAutonomousDatabase{}

func (m *modelAutonomousDatabase) client(ctx context.Context) (*oracledatabase.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := oracledatabase.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building OracleDatabase REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelAutonomousDatabase) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.OracleDatabaseAutonomousDatabase{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %v: %w", krm.OracleDatabaseAutonomousDatabaseGVK, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.OracleDatabaseAutonomousDatabaseIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := OracleDatabaseAutonomousDatabaseSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &autonomousDatabaseAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *modelAutonomousDatabase) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.OracleDatabaseAutonomousDatabaseIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &autonomousDatabaseAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type autonomousDatabaseAdapter struct {
	id        *krm.OracleDatabaseAutonomousDatabaseIdentity
	gcpClient *oracledatabase.Client
	desired   *pb.AutonomousDatabase
	actual    *pb.AutonomousDatabase
	reader    client.Reader
}

var _ directbase.Adapter = &autonomousDatabaseAdapter{}

func (a *autonomousDatabaseAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AutonomousDatabase", "name", a.id.String())

	req := &pb.GetAutonomousDatabaseRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetAutonomousDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AutonomousDatabase %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *autonomousDatabaseAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AutonomousDatabase", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.AutonomousDatabase)

	req := &pb.CreateAutonomousDatabaseRequest{
		Parent:               a.id.ParentString(),
		AutonomousDatabaseId: a.id.AutonomousDatabase,
		AutonomousDatabase:   desired,
	}

	createOp, err := a.gcpClient.CreateAutonomousDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AutonomousDatabase %s: %w", a.id.String(), err)
	}

	if _, err := createOp.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for AutonomousDatabase %s creation: %w", a.id.String(), err)
	}

	latest, err := a.gcpClient.GetAutonomousDatabase(ctx, &pb.GetAutonomousDatabaseRequest{
		Name: a.id.String(),
	})
	if err != nil {
		return fmt.Errorf("getting AutonomousDatabase %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, op, latest)
}

func (a *autonomousDatabaseAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AutonomousDatabase", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.AutonomousDatabase)

	diffs, updateMask, err := a.compareAutonomousDatabase(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateAutonomousDatabaseRequest{
		UpdateMask:         updateMask,
		AutonomousDatabase: desired,
	}
	req.AutonomousDatabase.Name = a.id.String()

	op, err := a.gcpClient.UpdateAutonomousDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AutonomousDatabase %s: %w", a.id.String(), err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for AutonomousDatabase %s update: %w", a.id.String(), err)
	}

	latest, err := a.gcpClient.GetAutonomousDatabase(ctx, &pb.GetAutonomousDatabaseRequest{
		Name: a.id.String(),
	})
	if err != nil {
		return fmt.Errorf("getting AutonomousDatabase %s after update: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *autonomousDatabaseAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AutonomousDatabase", "name", a.id.String())

	req := &pb.DeleteAutonomousDatabaseRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteAutonomousDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AutonomousDatabase %s: %w", a.id.String(), err)
	}

	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for AutonomousDatabase %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *autonomousDatabaseAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.OracleDatabaseAutonomousDatabase{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OracleDatabaseAutonomousDatabaseSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.AutonomousDatabase

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.AutonomousDatabase)
	u.SetGroupVersionKind(krm.OracleDatabaseAutonomousDatabaseGVK)

	return u, nil
}

func (a *autonomousDatabaseAdapter) compareAutonomousDatabase(ctx context.Context, actual, desired *pb.AutonomousDatabase) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	desired.Name = a.id.String()

	maskedActual, err := mappers.OnlySpecFields(actual, OracleDatabaseAutonomousDatabaseSpec_FromProto, OracleDatabaseAutonomousDatabaseSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	if desired.AdminPasswordSecretVersion != "" && maskedActual.AdminPasswordSecretVersion == "" {
		maskedActual.AdminPasswordSecretVersion = desired.AdminPasswordSecretVersion
	}
	if desired.AdminPassword != "" && maskedActual.AdminPassword == "" {
		maskedActual.AdminPassword = desired.AdminPassword
	}

	clonedDesired := proto.Clone(desired).(*pb.AutonomousDatabase)

	diffPaths, diff, err := common.CompareProtoMessageStructuredDiff(clonedDesired, maskedActual, common.BasicDiff)
	if err != nil {
		return nil, nil, err
	}

	diff.Controller = k8s.ReconcilerTypeDirect

	// Filter out non-updatable identifier fields if present in diff
	diffPaths.Delete("name")
	var filteredFields []structuredreporting.DiffField
	for _, f := range diff.Fields {
		if f.ID != "name" {
			filteredFields = append(filteredFields, f)
		}
	}
	diff.Fields = filteredFields

	paths := diffPaths.UnsortedList()
	slices.Sort(paths)
	updateMask := &fieldmaskpb.FieldMask{Paths: paths}

	return diff, updateMask, nil
}

func (a *autonomousDatabaseAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AutonomousDatabase) error {
	mapCtx := &direct.MapContext{}
	status := &krm.OracleDatabaseAutonomousDatabaseStatus{}
	status.ObservedState = OracleDatabaseAutonomousDatabaseObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
