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
// See the License for the() exact code and formatting.
// See the LICENSE file for details.

package databasemigrationmigrationjob

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/clouddms/apiv1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datamigration/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/datamigration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DatabaseMigrationMigrationJobGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.DataMigrationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataMigrationClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building clouddms client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DatabaseMigrationMigrationJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	migrationJobID := id.(*krm.DatabaseMigrationMigrationJobIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := datamigration.DatabaseMigrationMigrationJobSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = migrationJobID.String()

	return &Adapter{
		id:        migrationJobID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.DatabaseMigrationMigrationJobIdentity
	gcpClient *gcp.DataMigrationClient
	desired   *pb.MigrationJob
	actual    *pb.MigrationJob
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DatabaseMigrationMigrationJob", "name", a.id.String())

	req := &pb.GetMigrationJobRequest{Name: a.id.String()}
	migrationJob, err := a.gcpClient.GetMigrationJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DatabaseMigrationMigrationJob %q: %w", a.id.String(), err)
	}

	a.actual = migrationJob
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DatabaseMigrationMigrationJob", "name", a.id.String())

	req := &pb.CreateMigrationJobRequest{
		Parent:         a.id.ParentString(),
		MigrationJobId: a.id.MigrationJob,
		MigrationJob:   a.desired,
	}
	op, err := a.gcpClient.CreateMigrationJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DatabaseMigrationMigrationJob %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DatabaseMigrationMigrationJob %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DatabaseMigrationMigrationJob", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DatabaseMigrationMigrationJob", "name", a.id.String())

	diffs, updateMask, err := compareDatabaseMigrationMigrationJob(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateMigrationJobRequest{
			UpdateMask:   updateMask,
			MigrationJob: a.desired,
		}
		op, err := a.gcpClient.UpdateMigrationJob(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DatabaseMigrationMigrationJob %s: %w", a.id.String(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DatabaseMigrationMigrationJob %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated DatabaseMigrationMigrationJob", "name", a.id.String())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MigrationJob) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DatabaseMigrationMigrationJobStatus{}
	status.ObservedState = datamigration.DatabaseMigrationMigrationJobObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatabaseMigrationMigrationJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(datamigration.DatabaseMigrationMigrationJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.MigrationJob)
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.MigrationJob)
	u.SetGroupVersionKind(krm.DatabaseMigrationMigrationJobGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DatabaseMigrationMigrationJob", "name", a.id.String())

	req := &pb.DeleteMigrationJobRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMigrationJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DatabaseMigrationMigrationJob %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DatabaseMigrationMigrationJob", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DatabaseMigrationMigrationJob %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareDatabaseMigrationMigrationJob(ctx context.Context, actual, desired *pb.MigrationJob) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, datamigration.DatabaseMigrationMigrationJobSpec_FromProto, datamigration.DatabaseMigrationMigrationJobSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
