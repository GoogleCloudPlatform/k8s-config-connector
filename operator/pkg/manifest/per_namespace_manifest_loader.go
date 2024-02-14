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

package manifest

import (
	"context"
	"fmt"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

type PerNamespaceManifestLoader struct {
	repo Repository
}

// Ensure that PerNamespaceManifestLoader implements declarative.ManifestController.
var _ declarative.ManifestController = &PerNamespaceManifestLoader{}

func NewPerNamespaceManifestLoader(repo Repository) *PerNamespaceManifestLoader {
	return &PerNamespaceManifestLoader{
		repo: repo,
	}
}

func (p *PerNamespaceManifestLoader) ResolveManifest(ctx context.Context, o runtime.Object) (map[string]string, error) {
	_, ok := o.(*corev1beta1.ConfigConnectorContext)
	if !ok {
		return nil, fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
	}

	componentName := k8s.ConfigConnectorComponentName
	channelName := k8s.StableChannel
	v, err := ResolveVersion(ctx, p.repo, componentName, channelName)
	if err != nil {
		return nil, fmt.Errorf("error resolving the version for %v in %v channel: %w", componentName, channelName, err)
	}

	return p.repo.LoadNamespacedComponents(ctx, componentName, v)
}
