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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"

	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.TagsTagBindingGVK, NewTagsTagBindingModel)
}

func NewTagsTagBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTagsTagBinding{config: *config}, nil
}

var _ directbase.Model = &modelTagsTagBinding{}

type modelTagsTagBinding struct {
	config config.ControllerConfig
}

func (m *modelTagsTagBinding) client(ctx context.Context) (*resourcemanager.TagBindingsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := resourcemanager.NewTagBindingsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TagsTagBinding client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTagsTagBinding) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.TagsTagBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating gcp client: %w", err)
	}
	return &TagsTagBindingAdapter{
		id:        id.(*krm.TagBindingIdentity),
		gcpClient: gcpClient,
	}, nil
}

func (m *modelTagsTagBinding) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type TagsTagBindingAdapter struct {
	id        *krm.TagBindingIdentity
	gcpClient *resourcemanager.TagBindingsClient
	actual    *resourcemanagerpb.TagBinding
}

var _ directbase.Adapter = &TagsTagBindingAdapter{}

func (a *TagsTagBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TagsTagBinding", "name", a.id)

	req := &resourcemanagerpb.ListTagBindingsRequest{
		Parent: a.id.Parent().String(),
		// TODO: PageSize and PageToken
	}
	it := a.gcpClient.ListTagBindings(ctx, req)

	for {
		tagBinding, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return false, err
		}
		if tagBinding.TagValue == a.id.TagValue() {
			a.actual = tagBinding
			return true, nil
		}
	}

	return false, nil
}

func (a *TagsTagBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TagsTagBinding", "name", a.id)

	req := &resourcemanagerpb.CreateTagBindingRequest{
		TagBinding: &resourcemanagerpb.TagBinding{
			Parent:   a.id.Parent().String(),
			TagValue: a.id.TagValue(),
		},
	}
	op, err := a.gcpClient.CreateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagsTagBinding %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of TagsTagBinding %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created TagsTagBinding", "name", a.id)

	status := &krm.TagsTagBindingStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	// TF-based trims the prefix "tagBindings/", which is weird because then it is is invalid to retrieve the object. Shall we keep this behavior?
	// This modified value is also the only accepted value in `spec.resourceID`.
	status.Name = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *TagsTagBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("TagsTagBinding does not have updateable fields; skipping update", "name", a.id)
	return nil
}

func (a *TagsTagBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TagsTagBinding", "name", a.id)

	req := &resourcemanagerpb.DeleteTagBindingRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteTagBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent TagsTagBinding, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagsTagBinding %s: %w", a.id, err)
	}
	if err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for deletion of TagsTagBinding %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted TagsTagBinding", "name", a.id)

	return true, nil
}

func (a *TagsTagBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
