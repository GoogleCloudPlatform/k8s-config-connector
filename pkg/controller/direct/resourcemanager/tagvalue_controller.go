// Copyright 2024 Google LLC
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

package resourcemanager

import (
	"context"
	"fmt"
	"strings"

	api "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tags/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.TagsTagValueGVK, newTagValueModel)
}

func newTagValueModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &tagValueModel{config: config}, nil
}

type tagValueModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &tagValueModel{}

type tagValueAdapter struct {
	resourceID string

	desired *krm.TagsTagValue
	actual  *pb.TagValue

	*gcpClient
	tagValuesClient *api.TagValuesClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &tagValueAdapter{}

// AdapterForObject implements the Model interface.
func (m *tagValueModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	tagValuesClient, err := gcpClient.newTagValuesClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.TagsTagValue{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	// resourceID is server-generated, no fallback
	// TODO: How do we do resource acquisition - maybe by shortname?
	resourceID = strings.TrimPrefix(resourceID, "tagValues/")

	return &tagValueAdapter{
		resourceID:    resourceID,
		desired:       obj,
		gcpClient:     gcpClient,
		tagValuesClient: tagValuesClient,
	}, nil
}

func (m *tagValueModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *tagValueAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.GetTagValueRequest{
		Name: a.fullyQualifiedName(),
	}
	tagValue, err := a.tagValuesClient.GetTagValue(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = tagValue

	return true, nil
}

// Delete implements the Adapter interface.
func (a *tagValueAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteTagValueRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.tagValuesClient.DeleteTagValue(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting tagValue %s: %w", a.fullyQualifiedName(), err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("tagValue deletion failed: %w", err)
	}
	// TODO: Do we need to check that it was deleted?

	return true, nil
}

// Create implements the Adapter interface.
func (a *tagValueAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	// TODO: Is this the way to handle conversion of KCC ParentRef to pb Parent?
	parent := a.desired.Spec.ParentRef.External
	tagValue := &pb.TagValue{
		Parent:      parent,
		ShortName:   a.desired.Spec.ShortName,
		Description: direct.ValueOf(a.desired.Spec.Description),
	}

	req := &pb.CreateTagValueRequest{
		TagValue: tagValue,
	}

	op, err := a.tagValuesClient.CreateTagValue(ctx, req)
	if err != nil {
		return fmt.Errorf("creating tagValue: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("tagValue creation failed: %w", err)
	}

	log.V(2).Info("created tagkey", "tagkey", created)

	resourceID := created.Name
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	status := &krm.TagsTagValueStatus{}
	if err := tagValueStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func tagValueStatusToKRM(in *pb.TagValue, out *krm.TagsTagValueStatus) error {
	out.NamespacedName = &in.NamespacedName
	// TODO: Should be metav1.Time (?)
	out.CreateTime = timeToKRMString(in.GetCreateTime())
	out.UpdateTime = timeToKRMString(in.GetUpdateTime())
	return nil
}

// Update implements the Adapter interface.
func (a *tagValueAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}
	update := &pb.TagValue{}
	update.Name = a.fullyQualifiedName()

	// description is the only field that can be updated
	if direct.ValueOf(a.desired.Spec.Description) != a.actual.GetDescription() {
		updateMask.Paths = append(updateMask.Paths, "description")
		update.Description = direct.ValueOf(a.desired.Spec.Description)
	}

	// TODO: Where/how do we want to enforce immutability?

	if len(updateMask.Paths) != 0 {
		req := &pb.UpdateTagValueRequest{
			TagValue:     update,
			UpdateMask: updateMask,
		}

		op, err := a.tagValuesClient.UpdateTagValue(ctx, req)
		if err != nil {
			return err
		}

		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("tagValue update failed: %w", err)
		}
		// TODO: Do we need to check that the operation succeeeded?
	}

	// TODO: Return updated object status
	return nil
}

func (a *tagValueAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *tagValueAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("tagValues/%s", a.resourceID)
}
