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

package gcpclient

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrNotFound = fmt.Errorf("resource not found")
)

type Client interface {
	Get(ctx context.Context, u *unstructured.Unstructured) (*unstructured.Unstructured, error)
	Apply(u *unstructured.Unstructured) (*unstructured.Unstructured, error)
	Delete(u *unstructured.Unstructured) error
	IsSupported(kind string) bool
}

type gcpClient struct {
	erroringK8sClient client.Client
	smLoader          *servicemappingloader.ServiceMappingLoader
	tfProvider        *schema.Provider
	supportedKinds    map[string]bool
}

func New(provider *schema.Provider, smLoader *servicemappingloader.ServiceMappingLoader) Client {
	client := gcpClient{
		erroringK8sClient: k8s.NewErroringClient(),
		smLoader:          smLoader,
		tfProvider:        provider,
		supportedKinds:    buildSupportedKindSet(smLoader),
	}
	return &client
}

func (c *gcpClient) IsSupported(kind string) bool {
	_, ok := c.supportedKinds[kind]
	return ok
}

func (c *gcpClient) Get(ctx context.Context, u *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	sm, err := c.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return nil, err
	}
	resource, err := krmtotf.NewResource(u, sm, c.tfProvider)
	if err != nil {
		return nil, fmt.Errorf("could not parse resource %s: %w", u.GetName(), err)
	}
	state, err := krmtotf.FetchLiveState(ctx, resource, c.tfProvider, c.erroringK8sClient, c.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state: %w", err)
	}
	if state == nil {
		return nil, ErrNotFound
	}
	return updateResourceAndNewUnstructuredFromState(resource, state)
}

func (c *gcpClient) Apply(u *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	ctx := context.Background()
	sm, err := c.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return nil, err
	}
	krmResource, err := krmtotf.NewResource(u, sm, c.tfProvider)
	if err != nil {
		return nil, fmt.Errorf("could not parse resource %s: %w", u.GetName(), err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, krmResource, c.tfProvider, c.erroringK8sClient, c.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state: %w", err)
	}
	config, _, err := krmtotf.KRMResourceToTFResourceConfig(krmResource, c.erroringK8sClient, c.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error expanding resource configuration: %w", err)
	}
	diff, err := krmResource.TFResource.Diff(ctx, liveState, config, c.tfProvider.Meta())
	if err != nil {
		return nil, fmt.Errorf("error calculating diff: %w", err)
	}
	if !liveState.Empty() && diff.RequiresNew() {
		return nil, k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff))
	}
	if !diff.Empty() {
		appliedState, diagnostics := krmResource.TFResource.Apply(ctx, liveState, diff, c.tfProvider.Meta())
		err := krmtotf.NewErrorFromDiagnostics(diagnostics)
		if err != nil {
			return nil, fmt.Errorf("error applying desired state: %w", err)
		}
		return updateResourceAndNewUnstructuredFromState(krmResource, appliedState)
	}
	// Even when there are no changes to the resource on GCP the status field may not be present in the 'u' parameter.
	// The implication of the return value of this method is that it returns the state of the resource post-apply.
	// Therefore, always create a new unstructured from the live state of the resource and return it.
	return updateResourceAndNewUnstructuredFromState(krmResource, liveState)
}

func (c *gcpClient) Delete(u *unstructured.Unstructured) error {
	ctx := context.Background()
	sm, err := c.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return err
	}
	krmResource, err := krmtotf.NewResource(u, sm, c.tfProvider)
	if err != nil {
		return fmt.Errorf("could not parse resource %s: %w", u.GetName(), err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, krmResource, c.tfProvider, c.erroringK8sClient, c.smLoader)
	if err != nil {
		return fmt.Errorf("error fetching live state: %w", err)
	}
	if liveState.Empty() {
		return nil
	}
	_, diagnostics := krmResource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, c.tfProvider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error deleting resource: %w", err)
	}
	return nil
}

func updateResourceAndNewUnstructuredFromState(resource *krmtotf.Resource, state *terraform.InstanceState) (*unstructured.Unstructured, error) {
	resource.Name = krmtotf.GetNameFromState(resource, state)
	resource.Labels = krmtotf.GetLabelsFromState(resource, state)
	resource.Annotations = krmtotf.GetAnnotationsFromState(resource, state)
	resource.Spec, resource.Status = krmtotf.ResolveSpecAndStatusWithResourceID(resource, state)
	return resource.MarshalAsUnstructured()
}

// this map is temporary as we fix the bugs in some of the resource kinds -- most bugs are in the import functionality
var unsupportedKinds = map[string]bool{
	// these need tfiamclient -- need to bring that in
	"IAMPolicy":       true,
	"IAMPolicyMember": true,
	// doesn't work with end-user credentials
	"AccessContextManagerAccessLevel": true,
}

func buildSupportedKindSet(smLoader *servicemappingloader.ServiceMappingLoader) map[string]bool {
	kinds := make(map[string]bool)
	sms := smLoader.GetServiceMappings()
	for _, sm := range sms {
		for _, rc := range sm.Spec.Resources {
			if !isRCSupported(rc) {
				continue
			}
			kinds[rc.Kind] = true
		}
	}
	return kinds
}

func isRCSupported(rc v1alpha1.ResourceConfig) bool {
	if _, ok := unsupportedKinds[rc.Kind]; ok {
		return false
	}
	return !hasServerGeneratedID(rc)
}

func hasServerGeneratedID(rc v1alpha1.ResourceConfig) bool {
	return rc.ServerGeneratedIDField != ""
}
