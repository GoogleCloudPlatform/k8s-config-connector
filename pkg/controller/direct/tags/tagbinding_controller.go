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
	"errors"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	api "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.TagsTagBindingGVK, NewTagsTagBindingModel)
}

func NewTagsTagBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &TagsTagBindingModel{config: config}, nil
}

var _ directbase.Model = &TagsTagBindingModel{}

type TagsTagBindingModel struct {
	config *config.ControllerConfig
}

func (m *TagsTagBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	tagBindingsClient, err := newTagBindingsClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.TagsTagBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	var id *krm.TagsTagBindingIdentity
	if obj.Spec.ResourceID != nil || obj.Status.ExternalRef != nil {
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idFromObject.(*krm.TagsTagBindingIdentity)
	}

	var desired *pb.TagBinding
	{
		mapCtx := &direct.MapContext{}
		desired = TagsTagBindingSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &TagsTagBindingAdapter{
		id:                id,
		tagBindingsClient: tagBindingsClient,
		desired:           desired,
	}, nil
}

func (m *TagsTagBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//cloudresourcemanager.googleapis.com/") {
		return nil, nil
	}

	id := &krm.TagsTagBindingIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	tagBindingsClient, err := newTagBindingsClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &TagsTagBindingAdapter{
		id:                id,
		tagBindingsClient: tagBindingsClient,
	}, nil
}

type TagsTagBindingAdapter struct {
	id                *krm.TagsTagBindingIdentity
	tagBindingsClient *api.TagBindingsClient
	desired           *pb.TagBinding
	actual            *pb.TagBinding
}

var _ directbase.Adapter = &TagsTagBindingAdapter{}

func (a *TagsTagBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting TagsTagBinding", "name", fqn)

	// There is not GetTagBindingRequest, we can only List
	var found *pb.TagBinding
	parent := a.desired.GetParent()
	req := &pb.ListTagBindingsRequest{Parent: parent}
	it := a.tagBindingsClient.ListTagBindings(ctx, req)
	for {
		tagBinding, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return false, fmt.Errorf("listing tag bindings: %w", err)
		}
		if tagBinding.TagValue == a.desired.GetTagValue() {
			if found != nil {
				return false, fmt.Errorf("found multiple tag bindings matching %q and %q", parent, a.desired.GetTagValue())
			}
			found = tagBinding
		}
	}
	if found == nil {
		return false, nil
	}

	a.actual = found
	return true, nil
}

func (a *TagsTagBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// There no FQN until after creation (server generated id).

	log := klog.FromContext(ctx)
	log.V(0).Info("creating TagsTagBinding")

	req := &pb.CreateTagBindingRequest{
		TagBinding: direct.ProtoClone(a.desired),
	}

	op, err := a.tagBindingsClient.CreateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagsTagBinding: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of TagsTagBinding: %w", err)
	}
	log.V(0).Info("created TagsTagBinding", "name", created.GetName())

	// For compatibility, we set spec.resourceID after creation because this is a server-generated-id resource that we are migrating from terraform/DCL.
	// More info in docs/ai/server-generated-id.md
	resourceID := strings.TrimPrefix(created.GetName(), "tagBindings/")
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *TagsTagBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()

	diff, updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for TagsTagBinding %q: %w", fqn, err)
	}

	structuredreporting.ReportDiff(ctx, diff)

	latest := a.actual
	if len(updateMask.Paths) != 0 {
		return fmt.Errorf("cannot update TagsTagBinding %q: fields changed: %v; TagBindings are immutable after creation", fqn, updateMask.Paths)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *TagsTagBindingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.TagBinding) error {
	status := &krm.TagsTagBindingStatus{}

	// NOTYET: observedState
	// {
	// 	mapCtx := &direct.MapContext{}
	// 	status.ObservedState = TagsTagBindingObservedState_v1alpha1_FromProto(mapCtx, latest)
	// 	if mapCtx.Err() != nil {
	// 		return mapCtx.Err()
	// 	}
	// }

	status.ExternalRef = direct.PtrTo(latest.GetName())

	// Legacy status fields
	status.Name = direct.PtrTo(strings.TrimPrefix(latest.GetName(), "tagBindings/"))

	return op.UpdateStatus(ctx, status, nil)
}

func (a *TagsTagBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	fqn := a.id.String()
	if a.actual == nil {
		return nil, fmt.Errorf("TagsTagBinding %q not found", fqn)
	}

	obj := &krm.TagsTagBinding{}

	{
		mapCtx := &direct.MapContext{}
		spec := TagsTagBindingSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		obj.Spec = *spec
	}

	obj.SetGroupVersionKind(krm.TagsTagBindingGVK)
	name := a.actual.TagValueNamespacedName
	name = strings.ReplaceAll(name, "/", "_")
	obj.Name = name
	obj.Spec.ResourceID = direct.PtrTo(a.actual.GetName())

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting TagsTagBinding to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: u}, nil
}

// Delete implements the Adapter interface.
func (a *TagsTagBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting TagsTagBinding", "name", fqn)

	req := &pb.DeleteTagBindingRequest{}
	req.Name = fqn

	op, err := a.tagBindingsClient.DeleteTagBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent TagsTagBinding, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagsTagBinding %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted TagsTagBinding", "name", fqn)

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of TagsTagBinding %s: %w", fqn, err)
	}

	return true, nil
}

// TODO: Make this function generic and reuse across models.
func (a *TagsTagBindingAdapter) changedFields(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	// Compute the actual with only the spec fields populated.
	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := TagsTagBindingSpec_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := TagsTagBindingSpec_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	return buildDiff(ctx, a.desired.ProtoReflect(), actualMasked)
}
