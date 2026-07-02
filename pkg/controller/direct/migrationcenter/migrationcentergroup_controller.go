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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.MigrationCenterGroupGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
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

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MigrationCenterGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	groupIdentity := identity.(*krm.MigrationCenterGroupIdentity)
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        groupIdentity,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.MigrationCenterGroupIdentity
	gcpClient *gcp.Client
	desired   *krm.MigrationCenterGroup
	actual    *pb.Group
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id.Group == "" { // resource is not yet created
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("getting MigrationCenterGroup", "name", fqn)

	req := &pb.GetGroupRequest{Name: fqn}
	grouppb, err := a.gcpClient.GetGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MigrationCenterGroup %q: %w", fqn, err)
	}

	a.actual = grouppb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MigrationCenterGroup", "name", fqn)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := MigrationCenterGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.ParentString()
	req := &pb.CreateGroupRequest{
		Parent:  parent,
		GroupId: a.id.Group,
		Group:   resource,
	}
	op, err := a.gcpClient.CreateGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Group %s: %w", fqn, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Group %s creation: %w", fqn, err)
	}
	log.V(2).Info("successfully created Group", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MigrationCenterGroup", "name", fqn)

	mapCtx := &direct.MapContext{}
	desiredPb := MigrationCenterGroupSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	maskedActual, err := mappers.OnlySpecFields(a.actual, MigrationCenterGroupSpec_FromProto, MigrationCenterGroupSpec_ToProto)
	if err != nil {
		return err
	}

	clonedDesired := proto.Clone(desiredPb).(*pb.Group)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", fqn)
		return nil
	}

	diffs.Object = u
	structuredreporting.ReportDiff(ctx, diffs)

	desiredPb.Name = a.actual.Name
	req := &pb.UpdateGroupRequest{
		Group:      desiredPb,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Group %s: %w", fqn, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Group %s update: %w", fqn, err)
	}
	log.V(2).Info("successfully updated Group", "name", fqn)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Group) error {
	mapCtx := &direct.MapContext{}
	status := &krm.MigrationCenterGroupStatus{}
	status.ObservedState = MigrationCenterGroupObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &latest.Name
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MigrationCenterGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MigrationCenterGroupSpec_FromProto(mapCtx, a.actual))
	obj.Spec.ResourceID = &a.id.Group
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.Group)
	u.SetGroupVersionKind(krm.MigrationCenterGroupGVK)
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting Group", "name", fqn)

	req := &pb.DeleteGroupRequest{Name: fqn}
	op, err := a.gcpClient.DeleteGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting Group %s: %w", fqn, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting Group %s deletion: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted Group", "name", fqn)
	return true, nil
}
