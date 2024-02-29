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

package leaser

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leasable"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceLeaser struct {
	kubeClient client.Client
	leaser     *Leaser
}

func NewResourceLeaser(tfProvider *schema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, kubeClient client.Client) *ResourceLeaser {
	return &ResourceLeaser{
		kubeClient: kubeClient,
		leaser:     NewLeaser(tfProvider, smLoader, kubeClient),
	}
}

// Soft obtain obtains a lease for the resource given the live labels. The method should be called on resources that are leasable.
//
// It does not write the results to GCP so the caller must apply the changes to GCP if persistence is desired
func (l *ResourceLeaser) SoftObtain(ctx context.Context, resource *k8s.Resource, liveLabels map[string]string) error {
	uniqueID, err := cluster.GetNamespaceID(ctx, k8s.NamespaceIDConfigMapNN, l.kubeClient, resource.GetNamespace())
	if err != nil {
		return fmt.Errorf("error getting unique id for namespace '%v': %w", resource.GetNamespace(), err)
	}
	if err := l.leaser.SoftObtain(resource, liveLabels, uniqueID, k8s.TimeToLeaseExpiration, k8s.TimeToLeaseRenewal); err != nil {
		return fmt.Errorf("error obtaining lease: %w", err)
	}
	return nil
}

func (l *ResourceLeaser) Release(ctx context.Context, u *unstructured.Unstructured) error {
	uniqueID, err := cluster.GetNamespaceID(ctx, k8s.NamespaceIDConfigMapNN, l.kubeClient, u.GetNamespace())
	if err != nil {
		return fmt.Errorf("error getting unique id for namespace '%v': %w", u.GetNamespace(), err)
	}
	if err := l.leaser.Release(ctx, u, uniqueID); err != nil {
		return fmt.Errorf("error releasing lease on %v with name '%v': %w", u.GroupVersionKind(), u.GetName(), err)
	}
	return nil
}

func (l *ResourceLeaser) IsLeasable(resource *krmtotf.Resource) (ok bool, err error) {
	return leasable.ResourceConfigSupportsLeasing(&resource.ResourceConfig, l.leaser.tfProvider.ResourcesMap)
}
