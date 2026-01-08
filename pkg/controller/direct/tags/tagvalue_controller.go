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

package tags

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	api "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.TagsTagValueGVK, NewTagsTagValueModel)
}

func NewTagsTagValueModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &TagsTagValueModel{config: config}, nil
}

var _ directbase.Model = &TagsTagValueModel{}

type TagsTagValueModel struct {
	config *config.ControllerConfig
}

func (m *TagsTagValueModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	TagValuesClient, err := newTagValuesClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.TagsTagValue{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	var id *krm.TagsTagValueIdentity
	if obj.Spec.ResourceID != nil || obj.Status.ExternalRef != nil {
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idFromObject.(*krm.TagsTagValueIdentity)
	}

	var desired *pb.TagValue
	{
		mapCtx := &direct.MapContext{}
		desired = TagsTagValueSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &TagsTagValueAdapter{
		id:              id,
		tagValuesClient: TagValuesClient,
		desired:         desired,
	}, nil
}

func (m *TagsTagValueModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//cloudresourcemanager.googleapis.com/") {
		return nil, nil
	}

	id := &krm.TagsTagValueIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	TagValuesClient, err := newTagValuesClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &TagsTagValueAdapter{
		id:              id,
		tagValuesClient: TagValuesClient,
	}, nil
}

type TagsTagValueAdapter struct {
	id              *krm.TagsTagValueIdentity
	tagValuesClient *api.TagValuesClient
	desired         *pb.TagValue
	actual          *pb.TagValue
}

var _ directbase.Adapter = &TagsTagValueAdapter{}

func (a *TagsTagValueAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting TagsTagValue", "name", fqn)

	req := &pb.GetTagValueRequest{Name: fqn}
	actual, err := a.tagValuesClient.GetTagValue(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TagsTagValue %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *TagsTagValueAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// There no FQN until after creation (server generated id).

	log := klog.FromContext(ctx)
	log.V(0).Info("creating TagsTagValue")

	req := &pb.CreateTagValueRequest{
		TagValue: direct.ProtoClone(a.desired),
	}

	op, err := a.tagValuesClient.CreateTagValue(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagsTagValue: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of TagsTagValue: %w", err)
	}
	log.V(0).Info("created TagsTagValue", "name", created.GetName())

	// For compatibility, we set spec.resourceID after creation because this is a server-generated-id resource that we are migrating from terraform/DCL.
	// More info in docs/ai/server-generated-id.md
	resourceID := strings.TrimPrefix(created.GetName(), "tagValues/")
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *TagsTagValueAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	req := &pb.UpdateTagValueRequest{
		TagValue: direct.ProtoClone(a.desired),
	}
	req.TagValue.Name = fqn

	diff, updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for TagsTagValue %q: %w", fqn, err)
	}
	req.UpdateMask = updateMask

	structuredreporting.ReportDiff(ctx, diff)

	latest := a.actual
	if len(req.UpdateMask.Paths) != 0 {
		log.V(0).Info("updating TagsTagValue", "name", fqn)

		op, err := a.tagValuesClient.UpdateTagValue(ctx, req)
		if err != nil {
			return fmt.Errorf("updating TagsTagValue %q: %w", fqn, err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of TagsTagValue %s: %w", fqn, err)
		}
		log.V(0).Info("updated TagsTagValue", "name", fqn)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *TagsTagValueAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.TagValue) error {
	status := &krm.TagsTagValueStatus{}

	// NOTYET: observedState
	// {
	// 	mapCtx := &direct.MapContext{}
	// 	status.ObservedState = TagsTagValueObservedState_v1alpha1_FromProto(mapCtx, latest)
	// 	if mapCtx.Err() != nil {
	// 		return mapCtx.Err()
	// 	}
	// }

	status.ExternalRef = direct.PtrTo(latest.GetName())

	// Legacy status fields
	status.Name = direct.PtrTo(strings.TrimPrefix(latest.GetName(), "tagValues/"))
	status.NamespacedName = direct.LazyPtr(latest.GetNamespacedName())
	status.CreateTime = direct.PtrTo(latest.GetCreateTime().AsTime().Format("2006-01-02T15:04:05Z07:00"))
	status.UpdateTime = direct.PtrTo(latest.GetUpdateTime().AsTime().Format("2006-01-02T15:04:05Z07:00"))

	return op.UpdateStatus(ctx, status, nil)
}

func (a *TagsTagValueAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("TagsTagValue %q not found", fqn)
	}

	obj := &krm.TagsTagValue{}

	{
		mapCtx := &direct.MapContext{}
		spec := TagsTagValueSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.TagsTagValueGVK)
	obj.Name = a.id.TagValue

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting TagsTagValue to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

// Delete implements the Adapter interface.
func (a *TagsTagValueAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting TagsTagValue", "name", fqn)

	req := &pb.DeleteTagValueRequest{}
	req.Name = fqn

	op, err := a.tagValuesClient.DeleteTagValue(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent TagsTagValue, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagsTagValue %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted TagsTagValue", "name", fqn)

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of TagsTagValue %s: %w", fqn, err)
	}

	return true, nil
}

func (a *TagsTagValueAdapter) changedFields(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	// Compute the actual with only the spec fields populated.
	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := TagsTagValueSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := TagsTagValueSpec_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	return buildDiff(ctx, a.desired.ProtoReflect(), actualMasked)
}
