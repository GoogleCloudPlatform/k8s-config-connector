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

package monitoring

import (
	"context"
	"fmt"
	"strings"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.MonitoringGroupGVK, newGroupModel)
}

func newGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &groupModel{config: config}, nil
}

type groupModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &groupModel{}

type groupAdapter struct {
	id *krm.MonitoringGroupIdentity

	desired *pb.Group
	actual  *pb.Group

	groupsClient *monitoring.GroupClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &groupAdapter{}

// AdapterForObject implements the Model interface.
func (m *groupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	groupsClient, err := gcpClient.newGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.MonitoringGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := MonitoringGroupSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &groupAdapter{
		id:           id.(*krm.MonitoringGroupIdentity),
		desired:      desiredProto,
		groupsClient: groupsClient,
	}, nil
}

func (m *groupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //monitoring.googleapis.com/projects/PROJECT_NUMBER/groups/GROUP_ID
	if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
		return nil, nil
	}

	id := &krm.MonitoringGroupIdentity{}
	if err := id.FromExternal(strings.TrimPrefix(url, "//monitoring.googleapis.com/")); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	groupsClient, err := gcpClient.newGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &groupAdapter{
		id:           id,
		groupsClient: groupsClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *groupAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Group == "" {
		return false, nil
	}

	req := &pb.GetGroupRequest{
		Name: a.fullyQualifiedName(),
	}
	group, err := a.groupsClient.GetGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = group

	return true, nil
}

// Delete implements the Adapter interface.
func (a *groupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Check if exists / already deleted
	exists, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	req := &pb.DeleteGroupRequest{
		Name: a.fullyQualifiedName(),
	}

	if err := a.groupsClient.DeleteGroup(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting group %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *groupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := a.id.ParentString()

	req := &pb.CreateGroupRequest{
		Name:  parent,
		Group: a.desired,
	}

	log.V(2).Info("creating group", "req", req)
	created, err := a.groupsClient.CreateGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating group: %w", err)
	}
	log.V(2).Info("created group", "group", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *groupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	diffs, _, err := compareGroup(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		report.AddDiff(diffs)
		structuredreporting.ReportDiff(ctx, report)

		req := &pb.UpdateGroupRequest{
			Group: a.desired,
		}
		req.Group.Name = a.fullyQualifiedName()

		log.V(2).Info("updating group", "request", req)
		updated, err := a.groupsClient.UpdateGroup(ctx, req)
		if err != nil {
			return err
		}
		log.V(2).Info("updated group", "group", updated)
		a.actual = updated
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *groupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Group)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetGroupVersionKind(krm.MonitoringGroupGVK)
	u.SetName(a.id.Group)

	return u, nil
}

func (a *groupAdapter) fullyQualifiedName() string {
	return a.id.String()
}

func (a *groupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Group) error {
	status := &krm.MonitoringGroupStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func compareGroup(ctx context.Context, actual, desired *pb.Group) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, MonitoringGroupSpec_FromProto, MonitoringGroupSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Group)

	populateDefaults := func(obj *pb.Group) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
