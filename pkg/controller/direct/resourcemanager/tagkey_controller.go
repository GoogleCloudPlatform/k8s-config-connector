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
	"time"

	api "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	registry.RegisterModel(krm.TagsTagKeyGVK, newTagKeyModel)
}

func newTagKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &tagKeyModel{config: config}, nil
}

type tagKeyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &tagKeyModel{}

type tagKeyAdapter struct {
	resourceID string

	desired *krm.TagsTagKey
	actual  *pb.TagKey

	*gcpClient
	tagKeysClient *api.TagKeysClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &tagKeyAdapter{}

// AdapterForObject implements the Model interface.
func (m *tagKeyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	tagKeysClient, err := gcpClient.newTagKeysClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.TagsTagKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	// resourceID is server-generated, no fallback
	// TODO: How do we do resource acquisition - maybe by shortname?
	resourceID = strings.TrimPrefix(resourceID, "tagKeys/")

	return &tagKeyAdapter{
		resourceID:    resourceID,
		desired:       obj,
		gcpClient:     gcpClient,
		tagKeysClient: tagKeysClient,
	}, nil
}

func (m *tagKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *tagKeyAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.GetTagKeyRequest{
		Name: a.fullyQualifiedName(),
	}
	tagKey, err := a.tagKeysClient.GetTagKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = tagKey

	return true, nil
}

// Delete implements the Adapter interface.
func (a *tagKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteTagKeyRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.tagKeysClient.DeleteTagKey(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting tagKey %s: %w", a.fullyQualifiedName(), err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("tagKey deletion failed: %w", err)
	}
	// TODO: Do we need to check that it was deleted?

	return true, nil
}

// Create implements the Adapter interface.
func (a *tagKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	// TODO: Should be ref
	parent := a.desired.Spec.Parent
	tagKey := &pb.TagKey{
		Parent:      parent,
		ShortName:   a.desired.Spec.ShortName,
		Description: direct.ValueOf(a.desired.Spec.Description),
		PurposeData: a.desired.Spec.PurposeData,
	}

	if s := direct.ValueOf(a.desired.Spec.Purpose); s != "" {
		purpose, ok := pb.Purpose_value[s]
		if !ok {
			return fmt.Errorf("unknown purpose %q", s)
		}
		tagKey.Purpose = pb.Purpose(purpose)
	}
	req := &pb.CreateTagKeyRequest{
		TagKey: tagKey,
	}

	op, err := a.tagKeysClient.CreateTagKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating tagKey: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("tagKey creation failed: %w", err)
	}

	log.V(2).Info("created tagkey", "tagkey", created)

	resourceID := created.Name
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	status := &krm.TagsTagKeyStatus{}
	if err := tagKeyStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func tagKeyStatusToKRM(in *pb.TagKey, out *krm.TagsTagKeyStatus) error {
	out.NamespacedName = &in.NamespacedName
	// TODO: Should be metav1.Time (?)
	out.CreateTime = timeToKRMString(in.GetCreateTime())
	out.UpdateTime = timeToKRMString(in.GetUpdateTime())
	return nil
}

func timeToKRMString(t *timestamppb.Timestamp) *string {
	if t == nil {
		return nil
	}
	s := t.AsTime().Format(time.RFC3339Nano)
	return &s
}

// Update implements the Adapter interface.
func (a *tagKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}
	update := &pb.TagKey{}
	update.Name = a.fullyQualifiedName()

	// description is the only field that can be updated
	if direct.ValueOf(a.desired.Spec.Description) != a.actual.GetDescription() {
		updateMask.Paths = append(updateMask.Paths, "description")
		update.Description = direct.ValueOf(a.desired.Spec.Description)
	}

	// TODO: Where/how do we want to enforce immutability?

	if len(updateMask.Paths) != 0 {
		req := &pb.UpdateTagKeyRequest{
			TagKey:     update,
			UpdateMask: updateMask,
		}

		op, err := a.tagKeysClient.UpdateTagKey(ctx, req)
		if err != nil {
			return err
		}

		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("tagKey update failed: %w", err)
		}
		// TODO: Do we need to check that the operation succeeeded?
	}

	// TODO: Return updated object status
	return nil
}

func (a *tagKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *tagKeyAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("tagKeys/%s", a.resourceID)
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	// TODO: Just fetch this object?
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}
	// TODO: Merge to avoid overwriting conditions?
	u.Object["status"] = status

	return nil
}
