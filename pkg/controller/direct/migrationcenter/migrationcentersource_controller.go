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

package migrationcenter

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/migrationcenter/apiv1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.MigrationCenterSourceGVK, NewSourceModel)
}

func NewSourceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sourceModel{config: *config}, nil
}

var _ directbase.Model = &sourceModel{}

type sourceModel struct {
	config config.ControllerConfig
}

func (m *sourceModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building migrationcenter client: %w", err)
	}
	return gcpClient, nil
}

func (m *sourceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MigrationCenterSource{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	sourceIdentity := identity.(*krm.MigrationCenterSourceIdentity)

	mapCtx := &direct.MapContext{}
	desired := MigrationCenterSourceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &sourceAdapter{
		id:        sourceIdentity,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *sourceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type sourceAdapter struct {
	id        *krm.MigrationCenterSourceIdentity
	gcpClient *gcp.Client
	desired   *pb.Source
	actual    *pb.Source
}

var _ directbase.Adapter = &sourceAdapter{}

func (a *sourceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id.Source == "" { // resource is not yet created
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("getting MigrationCenterSource", "name", fqn)

	req := &pb.GetSourceRequest{Name: fqn}
	sourcePb, err := a.gcpClient.GetSource(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MigrationCenterSource %q: %w", fqn, err)
	}

	a.actual = sourcePb
	return true, nil
}

func (a *sourceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MigrationCenterSource", "name", fqn)

	parent := a.id.ParentString()
	clonedDesired := proto.Clone(a.desired).(*pb.Source)
	req := &pb.CreateSourceRequest{
		Parent:   parent,
		SourceId: a.id.Source,
		Source:   clonedDesired,
	}
	op, err := a.gcpClient.CreateSource(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MigrationCenterSource %s: %w", fqn, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting MigrationCenterSource %s creation: %w", fqn, err)
	}
	log.V(2).Info("successfully created MigrationCenterSource", "name", created.Name)

	latest, err := a.gcpClient.GetSource(ctx, &pb.GetSourceRequest{Name: fqn})
	if err != nil {
		latest = created
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *sourceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MigrationCenterSource", "name", fqn)

	maskedActual, err := mappers.OnlySpecFields(a.actual, MigrationCenterSourceSpec_FromProto, MigrationCenterSourceSpec_ToProto)
	if err != nil {
		return err
	}

	clonedDesired := proto.Clone(a.desired).(*pb.Source)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", fqn)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = u
	structuredreporting.ReportDiff(ctx, diffs)

	clonedDesired.Name = a.actual.Name
	req := &pb.UpdateSourceRequest{
		Source:     clonedDesired,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateSource(ctx, req)
	if err != nil {
		return fmt.Errorf("updating MigrationCenterSource %s: %w", fqn, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting MigrationCenterSource %s update: %w", fqn, err)
	}
	log.V(2).Info("successfully updated MigrationCenterSource", "name", fqn)

	latest, err := a.gcpClient.GetSource(ctx, &pb.GetSourceRequest{Name: fqn})
	if err != nil {
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *sourceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Source) error {
	mapCtx := &direct.MapContext{}
	status := &krm.MigrationCenterSourceStatus{}
	status.ObservedState = MigrationCenterSourceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &latest.Name
	return op.UpdateStatus(ctx, status, nil)
}

func (a *sourceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MigrationCenterSource{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MigrationCenterSourceSpec_FromProto(mapCtx, a.actual))
	obj.Spec.ResourceID = &a.id.Source
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.Source)
	u.SetGroupVersionKind(krm.MigrationCenterSourceGVK)
	return u, nil
}

// Delete implements the Adapter interface.
func (a *sourceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting MigrationCenterSource", "name", fqn)

	req := &pb.DeleteSourceRequest{Name: fqn}
	op, err := a.gcpClient.DeleteSource(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting MigrationCenterSource %s: %w", fqn, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting MigrationCenterSource %s deletion: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted MigrationCenterSource", "name", fqn)
	return true, nil
}
