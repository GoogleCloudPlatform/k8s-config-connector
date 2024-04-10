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
	"google.golang.org/api/iterator"
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

// AddTagBindingController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddTagBindingController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.TagsTagBindingGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &tagBindingModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m)
}

type tagBindingModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &tagBindingModel{}

type tagBindingAdapter struct {
	resourceID string

	desired *krm.TagsTagBinding
	actual  *pb.TagBinding

	tagBindingsClient *api.TagBindingsClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &tagBindingAdapter{}

// AdapterForObject implements the Model interface.
func (m *tagBindingModel) AdapterForObject(ctx context.Context, client client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	tagBindingsClient, err := m.newTagBindingsClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.TagsTagBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve the parent ref
	if obj.Spec.ParentRef.Name != "" {
		return nil, fmt.Errorf("refs in parent are not yet implemented")
	}

	// Resolve the tagValue ref
	if obj.Spec.TagValueRef.Name != "" {
		parentObj := &unstructured.Unstructured{}
		parentObj.SetGroupVersionKind(krm.TagsTagValueGVK)
		key := types.NamespacedName{
			Name:      obj.Spec.TagValueRef.Name,
			Namespace: obj.Spec.TagValueRef.Namespace,
		}
		if key.Namespace == "" {
			key.Namespace = obj.GetNamespace()
		}
		if err := client.Get(ctx, key, parentObj); err != nil {
			return nil, fmt.Errorf("getting tagValueRef %v: %w", key, err)
		}
		name, _, err := unstructured.NestedString(parentObj.Object, "status", "name")
		if err != nil {
			return nil, fmt.Errorf("getting status.name: %w", err)
		}
		if name == "" {
			// TODO: Return correct dependency-not-ready value
			return nil, fmt.Errorf("not ready")
		}
		external := "tagValues/" + name
		obj.Spec.TagValueRef = v1alpha1.ResourceRef{
			External: external,
		}
	}

	// TODO: validate that tagValueRef.External is well-formed if the user provided it

	resourceID := ValueOf(obj.Spec.ResourceID)
	// // resourceID is server-generated, no fallback
	// // TODO: How do we do resource acquisition - maybe by shortname?
	// resourceID = strings.TrimPrefix(resourceID, "tagBindings/")

	return &tagBindingAdapter{
		resourceID:        resourceID,
		desired:           obj,
		tagBindingsClient: tagBindingsClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *tagBindingAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	parent := a.desired.Spec.ParentRef.External
	if parent == "" {
		return false, fmt.Errorf("parent is empty")
	}
	tagValue := a.desired.Spec.TagValueRef.External
	if tagValue == "" {
		return false, fmt.Errorf("tagValue is empty")
	}

	req := &pb.ListTagBindingsRequest{
		Parent: parent,
	}

	tagBindingIterator := a.tagBindingsClient.ListTagBindings(ctx, req)

	var found []*pb.TagBinding
	for {
		// Its second return value is iterator.Done if there are no more
		// // results. Once Next returns Done, all subsequent calls will return Done.
		next, err := tagBindingIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, fmt.Errorf("listing tagBindings: %w", err)
		}
		if next.TagValue == tagValue {
			found = append(found, next)
		}
	}

	if len(found) == 0 {
		return false, nil
	}
	if len(found) > 1 {
		return false, fmt.Errorf("found multiple values matching tagValue %q", tagValue)
	}
	a.actual = found[0]

	return true, nil
}

// Delete implements the Adapter interface.
func (a *tagBindingAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deletd
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteTagBindingRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.tagBindingsClient.DeleteTagBinding(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting tagBinding: %w", err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("tagBinding deletion failed: %w", err)
	}
	// TODO: Do we need to check that it was deleted?

	return true, nil
}

// Create implements the Adapter interface.
func (a *tagBindingAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := a.desired.Spec.ParentRef.External
	if parent == "" {
		return fmt.Errorf("parent is empty")
	}
	tagValue := a.desired.Spec.TagValueRef.External
	if tagValue == "" {
		return fmt.Errorf("tagValue is empty")
	}

	tagBinding := &pb.TagBinding{
		Parent:   parent,
		TagValue: tagValue,
	}

	req := &pb.CreateTagBindingRequest{
		TagBinding: tagBinding,
	}

	log.Info("creating tagBinding", "request", req)

	op, err := a.tagBindingsClient.CreateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating tagBinding: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("tagBinding creation failed: %w", err)
	}

	log.V(2).Info("created tagbinding", "tagbinding", created)

	resourceID := strings.TrimPrefix(created.Name, "tagBindings/")
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	status := &krm.TagsTagBindingStatus{}
	if err := tagBindingStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func tagBindingStatusToKRM(in *pb.TagBinding, out *krm.TagsTagBindingStatus) error {
	name := in.Name
	out.Name = &name
	return nil
}

// Update implements the Adapter interface.
func (a *tagBindingAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// TODO: Where/how do we want to enforce immutability?

	// TODO: Make this an error if anything has changed

	// TODO: Return updated object status
	return nil
}

func (a *tagBindingAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("tagBindings/%s", a.resourceID)
}
