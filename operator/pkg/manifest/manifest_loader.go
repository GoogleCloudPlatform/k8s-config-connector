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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var (
	mlog = ctrl.Log.WithName("ManifestLoader")
	rlog = ctrl.Log.WithName("LocalRepository")
)

const (
	crdFileName                    = "crds.yaml"
	cnrmSystemFileName             = "0-cnrm-system.yaml"
	perNamespaceComponentsFileName = "per-namespace-components.yaml"
)

type Loader struct {
	repo Repository
}

func NewLoader(repo Repository) *Loader {
	return &Loader{
		repo: repo,
	}
}

// Ensure that ManifestController implements declarative.ManifestController.
var _ declarative.ManifestController = &Loader{}

func (c *Loader) ResolveManifest(ctx context.Context, o runtime.Object) (map[string]string, error) {
	cc, ok := o.(*corev1beta1.ConfigConnector)
	if !ok {
		return nil, fmt.Errorf("expected the resource to be a ConfigConnector, but it was not. Object: %v", o)
	}
	mlog.Info("resolving manifest", "name", cc.Name)

	componentName := cc.ComponentName()
	channelName := StableChannel
	version, err := ResolveVersion(ctx, c.repo, componentName, channelName)
	if err != nil {
		return nil, fmt.Errorf("error resolving the version for %v in %v channel: %w", componentName, channelName, err)
	}
	mlog.Info("resolved version from channel", "channel", channelName, "version", version)
	return c.repo.LoadManifest(ctx, componentName, version, cc)
}
