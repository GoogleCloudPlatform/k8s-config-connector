// Copyright 2022 Google LLC
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

package singleresourceiamclient

import (
	"context"
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type iamClient struct {
	tfProvider *schema.Provider
	smLoader   *servicemappingloader.ServiceMappingLoader
}

func New(tfProvider *schema.Provider, smloader *servicemappingloader.ServiceMappingLoader) *iamClient { //nolint:revive
	return &iamClient{
		smLoader:   smloader,
		tfProvider: tfProvider,
	}
}

func (i *iamClient) SupportsIAM(unstructured *unstructured.Unstructured) (bool, error) {
	groundKind := unstructured.GroupVersionKind().GroupKind()
	if registry.IsDirectByGK(groundKind) {
		return registry.SupportsIAM(groundKind)
	}

	rc, err := i.smLoader.GetResourceConfig(unstructured)
	if err != nil {
		return false, fmt.Errorf("error getting resource config for %q with name %q: %w",
			unstructured.GetKind(), unstructured.GetName(), err)
	}
	return krmtotf.SupportsIAM(rc), nil
}

func (i *iamClient) GetPolicy(ctx context.Context, resource *unstructured.Unstructured) (*v1beta1.IAMPolicy, error) {
	// This iamclient is being passed a "single resource client" which only supports Get(...) requests on 'resource'.
	// The reason to do this is to use the iam client without an instance of the API server. The IAM client needs the
	// client because it is required by krmtotf.
	kubeClient := newSingleResourceClient(resource)
	iamClient := kcciamclient.New(i.tfProvider, i.smLoader, kubeClient, nil, nil).TFIAMClient
	iamPolicySkeleton, err := i.newIAMPolicySkeleton(resource, kubeClient)
	if err != nil {
		return nil, fmt.Errorf("error creating new iam policy skeleton: %w", err)
	}
	return iamClient.GetPolicy(ctx, iamPolicySkeleton)
}

func (i *iamClient) newIAMPolicySkeleton(u *unstructured.Unstructured, kubeClient client.Client) (*v1beta1.IAMPolicy, error) {
	sm, err := i.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return nil, fmt.Errorf("error getting service mapping for '%v': %w", u.GroupVersionKind().Group, err)
	}
	resource, err := krmtotf.NewResource(u, sm, i.tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error converting '%v' with name '%v' to krmtotf resource: %w", u.GroupVersionKind(), u.GetName(), err)
	}
	importID, err := resource.GetImportID(kubeClient, i.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting import id for '%v' with name '%v': %w",
			resource.GroupVersionKind(),
			resource.GetName(),
			err)
	}
	policy := v1beta1.IAMPolicy{
		TypeMeta: v1.TypeMeta{
			APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
			Kind:       v1beta1.IAMPolicyGVK.Kind,
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      u.GetName(),
			Namespace: u.GetNamespace(),
		},
		Spec: v1beta1.IAMPolicySpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       u.GetKind(),
				APIVersion: u.GetAPIVersion(),
				External:   importID,
			},
		},
	}
	return &policy, nil
}

func newSingleResourceClient(resource *unstructured.Unstructured) client.Client {
	return &singleResourceClient{
		Resource: resource,
	}
}

type singleResourceClient struct {
	Resource *unstructured.Unstructured
}

func (c *singleResourceClient) Get(_ context.Context, key client.ObjectKey, obj client.Object, _ ...client.GetOption) error {
	unstructObj, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return fmt.Errorf("unexpected argument to single resource client, type is '%v' instead of '%v'",
			reflect.TypeOf(obj).Name(), reflect.TypeOf(&unstructured.Unstructured{}).Name())
	}
	if unstructObj.GroupVersionKind() != c.Resource.GroupVersionKind() {
		return fmt.Errorf("unexpected gvk for argument to single resource client: got '%v', expected '%v'",
			unstructObj.GroupVersionKind(), c.Resource.GroupVersionKind())
	}
	if key.Namespace != c.Resource.GetNamespace() {
		return fmt.Errorf("unexpected namespace for argument to single resource client: got '%v', expected '%v'",
			key.Namespace, c.Resource.GetNamespace())
	}
	if key.Name != c.Resource.GetName() {
		return fmt.Errorf("unexpected name for argument to single resource client: got '%v', expected '%v'",
			key.Name, c.Resource.GetName())
	}
	unstructObj.Object = deepcopy.MapStringInterface(c.Resource.Object)
	return nil
}

func (c *singleResourceClient) List(_ context.Context, list client.ObjectList, _ ...client.ListOption) error {
	return fmt.Errorf("unexpected call to client.List(...) for object with kind %v", list.GetObjectKind())
}

func (c *singleResourceClient) Create(_ context.Context, obj client.Object, _ ...client.CreateOption) error {
	return fmt.Errorf("unexpected call to client.Create(...) for object with kind %v", obj.GetObjectKind())
}

func (c *singleResourceClient) Delete(_ context.Context, obj client.Object, _ ...client.DeleteOption) error {
	return fmt.Errorf("unexpected call to client.Delete(...) for object with kind %v", obj.GetObjectKind())
}

func (c *singleResourceClient) DeleteAllOf(_ context.Context, obj client.Object, _ ...client.DeleteAllOfOption) error {
	return fmt.Errorf("unexpected call to client.DeleteAllOf(...) for object with kind %v", obj.GetObjectKind())
}

func (c *singleResourceClient) Update(_ context.Context, obj client.Object, _ ...client.UpdateOption) error {
	return fmt.Errorf("unexpected call to client.Update(...) for object with kind %v", obj.GetObjectKind())
}

func (c *singleResourceClient) Patch(_ context.Context, obj client.Object, _ client.Patch, _ ...client.PatchOption) error {
	return fmt.Errorf("unexpected call to client.Patch(...) for object with kind %v", obj.GetObjectKind())
}

func (c *singleResourceClient) Status() client.SubResourceWriter {
	panic("unexpected call to client.Status(...)")
}

func (c *singleResourceClient) Scheme() *runtime.Scheme {
	panic("unexpected call to client.Scheme(...)")
}

func (c *singleResourceClient) RESTMapper() meta.RESTMapper {
	panic("unexpected call to client.RESTMapper(...)")
}

func (c *singleResourceClient) GroupVersionKindFor(_ runtime.Object) (k8sschema.GroupVersionKind, error) {
	panic("unexpected call to client.GroupVersionKindFor(...)")
}

func (c *singleResourceClient) IsObjectNamespaced(_ runtime.Object) (bool, error) {
	panic("unexpected call to client.IsObjectNamespaced(...)")
}

func (c *singleResourceClient) SubResource(_ string) client.SubResourceClient {
	panic("unexpected call to client.SubResource(...)")
}

func (c *singleResourceClient) SubResourceWriter(_ string) client.SubResourceClient {
	panic("unexpected call to client.SubResource(...)")
}
