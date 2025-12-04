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
	registry.RegisterModel(krm.TagsTagKeyGVK, NewTagsTagKeyModel)
}

func NewTagsTagKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &TagsTagKeyModel{config: config}, nil
}

var _ directbase.Model = &TagsTagKeyModel{}

type TagsTagKeyModel struct {
	config *config.ControllerConfig
}

func (m *TagsTagKeyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	tagKeysClient, err := newTagKeysClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.TagsTagKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// normalize parent (sadly not a ref, for legacy reasons)
	{
		parent := direct.ValueOf(obj.Spec.Parent)
		tokens := strings.Split(parent, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			// TODO: Use ProjectRef once it uses the ref pattern
			// Ensure uses number
			projectNumber, err := m.config.ProjectMapper.LookupProjectNumber(ctx, tokens[1])
			if err != nil {
				return nil, fmt.Errorf("mapping project ID to number: %w", err)
			}
			obj.Spec.Parent = direct.PtrTo(fmt.Sprintf("projects/%d", projectNumber))
		} else if len(tokens) == 2 && tokens[0] == "organizations" {
			// Always numeric, no normalization needed
		} else {
			return nil, fmt.Errorf("invalid parent %q: expected form is project/{project_id} or organizations/{organization_id}", parent)
		}
	}

	var id *krm.TagsTagKeyIdentity
	if obj.Spec.ResourceID != nil || obj.Status.ExternalRef != nil {
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idFromObject.(*krm.TagsTagKeyIdentity)
	}

	var desired *pb.TagKey
	{
		mapCtx := &direct.MapContext{}
		desired = TagsTagKeySpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &TagsTagKeyAdapter{
		id:            id,
		tagKeysClient: tagKeysClient,
		desired:       desired,
	}, nil
}

func (m *TagsTagKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//cloudresourcemanager.googleapis.com/") {
		return nil, nil
	}

	id := &krm.TagsTagKeyIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	tagKeysClient, err := newTagKeysClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &TagsTagKeyAdapter{
		id:            id,
		tagKeysClient: tagKeysClient,
	}, nil
}

type TagsTagKeyAdapter struct {
	id            *krm.TagsTagKeyIdentity
	tagKeysClient *api.TagKeysClient
	desired       *pb.TagKey
	actual        *pb.TagKey
}

var _ directbase.Adapter = &TagsTagKeyAdapter{}

func (a *TagsTagKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting TagsTagKey", "name", fqn)

	req := &pb.GetTagKeyRequest{Name: fqn}
	actual, err := a.tagKeysClient.GetTagKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TagsTagKey %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *TagsTagKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// There no FQN until after creation (server generated id).

	log := klog.FromContext(ctx)
	log.V(0).Info("creating TagsTagKey")

	req := &pb.CreateTagKeyRequest{
		TagKey: direct.ProtoClone(a.desired),
	}

	op, err := a.tagKeysClient.CreateTagKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagsTagKey: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of TagsTagKey: %w", err)
	}
	log.V(0).Info("created TagsTagKey", "name", created.GetName())

	// For compatibility, we set spec.resourceID after creation because this is a server-generated-id resource that we are migrating from terraform/DCL.
	// More info in docs/ai/server-generated-id.md
	resourceID := strings.TrimPrefix(created.GetName(), "tagKeys/")
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *TagsTagKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	req := &pb.UpdateTagKeyRequest{
		TagKey: direct.ProtoClone(a.desired),
	}
	req.TagKey.Name = fqn

	diff, updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for TagsTagKey %q: %w", fqn, err)
	}
	req.UpdateMask = updateMask

	structuredreporting.ReportDiff(ctx, diff)

	latest := a.actual
	if len(req.UpdateMask.Paths) != 0 {
		log.V(0).Info("updating TagsTagKey", "name", fqn)

		op, err := a.tagKeysClient.UpdateTagKey(ctx, req)
		if err != nil {
			return fmt.Errorf("updating TagsTagKey %q: %w", fqn, err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of TagsTagKey %s: %w", fqn, err)
		}
		log.V(0).Info("updated TagsTagKey", "name", fqn)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *TagsTagKeyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.TagKey) error {
	status := &krm.TagsTagKeyStatus{}

	// NOTYET: observedState
	// {
	// 	mapCtx := &direct.MapContext{}
	// 	status.ObservedState = TagsTagKeyObservedState_v1alpha1_FromProto(mapCtx, latest)
	// 	if mapCtx.Err() != nil {
	// 		return mapCtx.Err()
	// 	}
	// }

	status.ExternalRef = direct.PtrTo(latest.GetName())

	// Legacy status fields
	status.Name = direct.PtrTo(strings.TrimPrefix(latest.GetName(), "tagKeys/"))
	status.NamespacedName = direct.PtrTo(latest.GetNamespacedName())
	status.CreateTime = direct.PtrTo(latest.GetCreateTime().AsTime().Format("2006-01-02T15:04:05Z07:00"))
	status.UpdateTime = direct.PtrTo(latest.GetUpdateTime().AsTime().Format("2006-01-02T15:04:05Z07:00"))

	return op.UpdateStatus(ctx, status, nil)
}

func (a *TagsTagKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("TagsTagKey %q not found", fqn)
	}

	obj := &krm.TagsTagKey{}

	{
		mapCtx := &direct.MapContext{}
		spec := TagsTagKeySpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.TagsTagKeyGVK)
	obj.Name = a.id.TagKey

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting TagsTagKey to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

// Delete implements the Adapter interface.
func (a *TagsTagKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting TagsTagKey", "name", fqn)

	req := &pb.DeleteTagKeyRequest{}
	req.Name = fqn

	op, err := a.tagKeysClient.DeleteTagKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent TagsTagKey, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagsTagKey %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted TagsTagKey", "name", fqn)

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of TagsTagKey %s: %w", fqn, err)
	}

	return true, nil
}

func (a *TagsTagKeyAdapter) changedFields(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	log := klog.FromContext(ctx)

	// Compute the actual with only the spec fields populated.
	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := TagsTagKeySpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := TagsTagKeySpec_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	diff := &structuredreporting.Diff{}

	var paths []string
	fields := actualMasked.Type().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		path := string(fields.Get(i).Name())
		fieldDiff, err := fieldHasChanged(path, a.desired.ProtoReflect(), actualMasked)
		if err != nil {
			log.Error(err, "error determining if field has changed", "field", path)
			// If we can't determine if the field has changed, include it in the update.
		} else if fieldDiff == nil {
			continue
		}
		diff.AddField(fieldDiff.FieldPath, fieldDiff.ActualValue, fieldDiff.DesiredValue)
		paths = append(paths, fieldDiff.FieldPath)
	}
	return diff, &fieldmaskpb.FieldMask{Paths: paths}, nil
}
