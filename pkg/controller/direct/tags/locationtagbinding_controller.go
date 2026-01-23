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

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1alpha1"
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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.TagsLocationTagBindingGVK, NewTagsLocationTagBindingModel)
}

func NewTagsLocationTagBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &TagsLocationTagBindingModel{config: config}, nil
}

var _ directbase.Model = &TagsLocationTagBindingModel{}

type TagsLocationTagBindingModel struct {
	config *config.ControllerConfig
}

func (m *TagsLocationTagBindingModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.TagsLocationTagBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	location := direct.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("spec.location must be specified for TagsLocationTagBinding")
	}

	tagBindingsClient, err := newLocationTagBindingsClient(ctx, m.config, location)
	if err != nil {
		return nil, err
	}

	var id *krm.TagsLocationTagBindingIdentity
	if obj.Spec.ResourceID != nil || obj.Status.ExternalRef != nil {
		idFromObject, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idFromObject.(*krm.TagsLocationTagBindingIdentity)
	}

	var desired *pb.TagBinding
	{
		mapCtx := &direct.MapContext{}
		desired = TagsLocationTagBindingSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
	}

	return &TagsLocationTagBindingAdapter{
		id:                id,
		tagBindingsClient: tagBindingsClient,
		desired:           desired,
		projectMapper:     m.config.ProjectMapper,
	}, nil
}

func (m *TagsLocationTagBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Not implemented, because we need a way to encode the location in the URL
	return nil, nil
}

type TagsLocationTagBindingAdapter struct {
	id                *krm.TagsLocationTagBindingIdentity
	tagBindingsClient *api.TagBindingsClient
	desired           *pb.TagBinding
	actual            *pb.TagBinding

	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &TagsLocationTagBindingAdapter{}

func (a *TagsLocationTagBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id == nil {
		return false, nil
	}

	fqn := a.id.String()
	log.V(0).Info("getting TagsLocationTagBinding", "name", fqn)

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

func (a *TagsLocationTagBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// There no FQN until after creation (server generated id).

	log := klog.FromContext(ctx)
	log.V(0).Info("creating TagsLocationTagBinding")

	req := &pb.CreateTagBindingRequest{
		TagBinding: direct.ProtoClone(a.desired),
	}

	op, err := a.tagBindingsClient.CreateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagsLocationTagBinding: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of TagsLocationTagBinding: %w", err)
	}
	log.V(0).Info("created TagsLocationTagBinding", "name", created.GetName())

	// For compatibility, we set spec.resourceID after creation because this is a server-generated-id resource that we are migrating from terraform/DCL.
	// More info in docs/ai/server-generated-id.md
	resourceID := created.GetName()
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *TagsLocationTagBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()

	diff, updateMask, err := a.changedFields(ctx)
	if err != nil {
		return fmt.Errorf("getting changed fields for TagsLocationTagBinding %q: %w", fqn, err)
	}

	structuredreporting.ReportDiff(ctx, diff)

	latest := a.actual
	if len(updateMask.Paths) != 0 {
		return fmt.Errorf("cannot update TagsLocationTagBinding %q: fields changed: %v; TagBindings are immutable after creation", fqn, updateMask.Paths)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *TagsLocationTagBindingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.TagBinding) error {
	status := &krm.TagsLocationTagBindingStatus{}

	// NOTYET: observedState
	// {
	// 	mapCtx := &direct.MapContext{}
	// 	status.ObservedState = TagsLocationTagBindingObservedState_v1alpha1_FromProto(mapCtx, latest)
	// 	if mapCtx.Err() != nil {
	// 		return mapCtx.Err()
	// 	}
	// }

	status.ExternalRef = direct.PtrTo(latest.GetName())

	// Legacy status fields
	status.Name = direct.PtrTo(latest.GetName())

	return op.UpdateStatus(ctx, status, nil)
}

func (a *TagsLocationTagBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	// Not implemented, because we need a way to encode the location in the URL.
	return nil, nil
}

// Delete implements the Adapter interface.
func (a *TagsLocationTagBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()

	log.V(0).Info("deleting TagsLocationTagBinding", "name", fqn)

	req := &pb.DeleteTagBindingRequest{}
	req.Name = fqn

	op, err := a.tagBindingsClient.DeleteTagBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(0).Info("skipping delete for non-existent TagsLocationTagBinding, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagsLocationTagBinding %s: %w", fqn, err)
	}
	log.V(0).Info("successfully deleted TagsLocationTagBinding", "name", fqn)

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of TagsLocationTagBinding %s: %w", fqn, err)
	}

	return true, nil
}

// TODO: Make this function generic and reuse across models.
func (a *TagsLocationTagBindingAdapter) changedFields(ctx context.Context) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {

	actual := direct.ProtoClone(a.actual)
	desired := direct.ProtoClone(a.desired)

	// Normalize both to use project numbers in links.
	if err := a.projectMapper.RemapLinkToProjectNumber(ctx, &actual.Parent); err != nil {
		return nil, nil, fmt.Errorf("remapping project in actual TagBinding: %w", err)
	}
	if err := a.projectMapper.RemapLinkToProjectNumber(ctx, &desired.Parent); err != nil {
		return nil, nil, fmt.Errorf("remapping project in desired TagBinding: %w", err)
	}

	// Compute the actual with only the spec fields populated.
	var actualMasked protoreflect.Message
	{
		mapCtx := &direct.MapContext{}
		actualSpec := TagsLocationTagBindingSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		mapCtx = &direct.MapContext{}
		specProto := TagsLocationTagBindingSpec_ToProto(mapCtx, actualSpec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		actualMasked = specProto.ProtoReflect()
	}

	return buildDiff(ctx, desired.ProtoReflect(), actualMasked)
}
