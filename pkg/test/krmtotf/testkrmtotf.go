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

package testkrmtotf

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewKRMResource(t *testing.T, u *unstructured.Unstructured, sm *v1alpha1.ServiceMapping, tfProvider *schema.Provider) *krmtotf.Resource {
	t.Helper()
	krmResource, err := krmtotf.NewResource(u, sm, tfProvider)
	if err != nil {
		t.Fatalf("error creating new krm resource: %v", err)
	}
	return krmResource
}

func FetchLiveState(t *testing.T, resource *krmtotf.Resource, provider *schema.Provider, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) *terraform.InstanceState {
	t.Helper()
	liveState, err := krmtotf.FetchLiveState(context.Background(), resource, provider, kubeClient, smLoader)
	if err != nil {
		t.Fatalf("error fetching live state: %v", err)
	}
	return liveState
}
