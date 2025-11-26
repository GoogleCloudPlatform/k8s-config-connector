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

package clouddms

import (
	"context"
	"fmt"

	api "cloud.google.com/go/clouddms/apiv1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
)

func init() {
	registry.RegisterModel(krm.CloudDMSMigrationJobGVK, newMigrationJobModel)
}

func newMigrationJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &migrationJobModel{config: config}, nil
}

type migrationJobModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &migrationJobModel{}

type migrationJobAdapter struct {
	id      *krm.MigrationJobIdentity
	desired *pb.MigrationJob
	actual  *pb.MigrationJob

	client *api.DataMigrationClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &migrationJobAdapter{}

// AdapterForObject implements the Model interface.
func (m *migrationJobModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newMigrationClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.CloudDMSMigrationJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, err
	}

	identity, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := CloudDMSMigrationJobSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &migrationJobAdapter{
		id:      identity.(*krm.MigrationJobIdentity),
		desired: desired,
		client:  client,
	}, nil
}

func (m *migrationJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}

// Find implements the Adapter interface.
func (a *migrationJobAdapter) Find(ctx context.Context) (bool, error) {
	req := &pb.GetMigrationJobRequest{
		Name: a.id.String(),
	}
	migrationJob, err := a.client.GetMigrationJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = migrationJob

	return true, nil
}

// Delete implements the Adapter interface.
func (a *migrationJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &pb.DeleteMigrationJobRequest{
		Name: a.id.String(),
	}

	op, err := a.client.DeleteMigrationJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting CloudDMSMigrationJob %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for CloudDMSMigrationJob delete %s: %w", a.id, err)
	}

	return true, nil
}

func (a *migrationJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Create implements the Adapter interface.
func (a *migrationJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	req := &pb.CreateMigrationJobRequest{
		Parent:         a.id.Parent().String(),
		MigrationJobId: a.id.ID(),
		MigrationJob:   a.desired,
	}
	log.V(0).Info("making CloudDMS CreateMigrationJob call", "request", req)

	op, err := a.client.CreateMigrationJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CloudDMSMigrationJob: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for CloudDMSMigrationJob create %s: %w", a.id, err)
	}

	log.V(0).Info("created CloudDMSMigrationJob", "migrationJob", created)

	mapCtx := &direct.MapContext{}
	status := &krm.CloudDMSMigrationJobStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = CloudDMSMigrationJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update implements the Adapter interface.
func (a *migrationJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	req := &pb.UpdateMigrationJobRequest{
		MigrationJob: a.desired,
		UpdateMask:   updateMask,
	}
	log.V(0).Info("making CloudDMS UpdateMigrationJob call", "request", req)

	op, err := a.client.UpdateMigrationJob(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CloudDMSMigrationJob: %w", err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for CloudDMSMigrationJob update %s: %w", a.id, err)
	}
	log.V(0).Info("updated CloudDMSMigrationJob", "migrationJob", updated)

	mapCtx := &direct.MapContext{}
	status := &krm.CloudDMSMigrationJobStatus{}
	status.ObservedState = CloudDMSMigrationJobObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}
