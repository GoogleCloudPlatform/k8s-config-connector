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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tags/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddTagValueController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddTagValueController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.TagsTagValueGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &tagValueModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m)
}

type tagValueModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &tagValueModel{}

type tagValueAdapter struct {
	resourceID string

	desired *krm.TagsTagValue
	actual  *pb.TagValue

	tagValuesClient *api.TagValuesClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &tagValueAdapter{}

// AdapterForObject implements the Model interface.
func (m *tagValueModel) AdapterForObject(ctx context.Context, client client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	tagValuesClient, err := m.newTagValuesClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.TagsTagValue{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve the parent ref
	if obj.Spec.ParentRef.Name != "" {
		parentObj := &unstructured.Unstructured{}
		parentObj.SetGroupVersionKind(krm.TagsTagKeyGVK)
		key := types.NamespacedName{
			Name:      obj.Spec.ParentRef.Name,
			Namespace: obj.Spec.ParentRef.Namespace,
		}
		if key.Namespace == "" {
			key.Namespace = obj.GetNamespace()
		}
		if err := client.Get(ctx, key, parentObj); err != nil {
			return nil, fmt.Errorf("getting parent %v: %w", key, err)
		}
		name, _, err := unstructured.NestedString(parentObj.Object, "status", "name")
		if err != nil {
			return nil, fmt.Errorf("getting status.name: %w", err)
		}
		if name == "" {
			// TODO: Return correct dependency-not-ready value
			return nil, fmt.Errorf("not ready")
		}
		external := "tagKeys/" + name
		obj.Spec.ParentRef = v1alpha1.ResourceRef{
			External: external,
		}
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	// resourceID is server-generated, no fallback
	// TODO: How do we do resource acquisition - maybe by shortname?
	resourceID = strings.TrimPrefix(resourceID, "tagValues/")

	return &tagValueAdapter{
		resourceID:      resourceID,
		desired:         obj,
		tagValuesClient: tagValuesClient,
	}, nil
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
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting tagValue %q: %w", req.Name, err)
	}

	a.actual = tagValue

	return true, nil
}

// Delete implements the Adapter interface.
func (a *tagValueAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deletd
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteTagValueRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.tagValuesClient.DeleteTagValue(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting tagValue: %w", err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("tagValue deletion failed: %w", err)
	}
	// TODO: Do we need to check that it was deleted?

	return true, nil
}

// Create implements the Adapter interface.
func (a *tagValueAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := a.desired.Spec.ParentRef.External
	if parent == "" {
		return fmt.Errorf("parent is empty")
	}
	tagValue := &pb.TagValue{
		Parent:      parent,
		ShortName:   a.desired.Spec.ShortName,
		Description: ValueOf(a.desired.Spec.Description),
	}

	req := &pb.CreateTagValueRequest{
		TagValue: tagValue,
	}

	log.Info("creating tagValue", "request", req)

	op, err := a.tagValuesClient.CreateTagValue(ctx, req)
	if err != nil {
		return fmt.Errorf("creating tagValue: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("tagValue creation failed: %w", err)
	}

	log.V(2).Info("created tagvalue", "tagvalue", created)

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
	name := in.Name
	name = strings.TrimPrefix(name, "tagValues/")
	out.Name = &name
	// TODO: Should be metav1.Time (?)
	out.CreateTime = timeToKRMString(in.GetCreateTime())
	out.UpdateTime = timeToKRMString(in.GetUpdateTime())
	return nil
}

// Update implements the Adapter interface.
func (a *tagValueAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// TODO: Skip updates at the higher level if no changes?
	updateMask := &fieldmaskpb.FieldMask{}
	update := &pb.TagValue{}
	update.Name = a.fullyQualifiedName()

	// description is the only field that can be updated
	if ValueOf(a.desired.Spec.Description) != a.actual.GetDescription() {
		updateMask.Paths = append(updateMask.Paths, "description")
		update.Description = ValueOf(a.desired.Spec.Description)
	}

	// TODO: Where/how do we want to enforce immutability?

	if len(updateMask.Paths) != 0 {
		req := &pb.UpdateTagValueRequest{
			TagValue:   update,
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

func (a *tagValueAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("tagValues/%s", a.resourceID)
}
