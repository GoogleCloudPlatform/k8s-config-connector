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

package databasemigrationconnectionprofile

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/clouddms/apiv1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datamigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/datamigration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DatabaseMigrationConnectionProfileGVK, NewModel)
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
		return nil, fmt.Errorf("building clouddms client options: %w", err)
	}
	gcpClient, err := gcp.NewDataMigrationClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building clouddms client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DatabaseMigrationConnectionProfile{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identityObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identityObj.(*krm.DatabaseMigrationConnectionProfileIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := datamigration.DatabaseMigrationConnectionProfileSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.DatabaseMigrationConnectionProfileIdentity
	gcpClient *gcp.DataMigrationClient
	desired   *pb.ConnectionProfile
	actual    *pb.ConnectionProfile
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DatabaseMigrationConnectionProfile", "name", a.id)

	req := &pb.GetConnectionProfileRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DatabaseMigrationConnectionProfile %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DatabaseMigrationConnectionProfile", "name", a.id)

	req := &pb.CreateConnectionProfileRequest{
		Parent:              a.id.ParentString(),
		ConnectionProfileId: a.id.ConnectionProfile,
		ConnectionProfile:   a.desired,
	}
	op, err := a.gcpClient.CreateConnectionProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DatabaseMigrationConnectionProfile %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting DatabaseMigrationConnectionProfile %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DatabaseMigrationConnectionProfile", "name", a.id)

	// Fetch fully-populated resource before calling updateStatus (Rule: Reconciling Empty or Incomplete LRO Responses)
	latest, err := a.gcpClient.GetConnectionProfile(ctx, &pb.GetConnectionProfileRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting DatabaseMigrationConnectionProfile %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DatabaseMigrationConnectionProfile", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareConnectionProfile(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for DatabaseMigrationConnectionProfile", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateConnectionProfileRequest{
		ConnectionProfile: a.desired,
		UpdateMask:        updateMask,
	}
	op, err := a.gcpClient.UpdateConnectionProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DatabaseMigrationConnectionProfile %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting DatabaseMigrationConnectionProfile %s update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated DatabaseMigrationConnectionProfile", "name", a.id)

	// Fetch fully-populated resource before calling updateStatus (Rule: Reconciling Empty or Incomplete LRO Responses)
	latest, err := a.gcpClient.GetConnectionProfile(ctx, &pb.GetConnectionProfileRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting DatabaseMigrationConnectionProfile %s after update: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatabaseMigrationConnectionProfile{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(datamigration.DatabaseMigrationConnectionProfileSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DatabaseMigrationConnectionProfileGVK)
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DatabaseMigrationConnectionProfile", "name", a.id)

	req := &pb.DeleteConnectionProfileRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DatabaseMigrationConnectionProfile, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DatabaseMigrationConnectionProfile %s: %w", a.id, err)
	}
	log.V(2).Info("successfully initiated deletion of DatabaseMigrationConnectionProfile", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete DatabaseMigrationConnectionProfile %s: %w", a.id, err)
	}
	log.V(2).Info("successfully waited for DatabaseMigrationConnectionProfile deletion LRO", "name", a.id)

	return true, nil
}

func compareConnectionProfile(ctx context.Context, actual, desired *pb.ConnectionProfile) (*structuredreporting.Diff, *field_mask.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, datamigration.DatabaseMigrationConnectionProfileSpec_FromProto, datamigration.DatabaseMigrationConnectionProfileSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.ConnectionProfile)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ConnectionProfile) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DatabaseMigrationConnectionProfileStatus{}
	status.ObservedState = datamigration.DatabaseMigrationConnectionProfileObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
