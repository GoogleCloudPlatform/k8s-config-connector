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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/loaders"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/yaml"
)

type Repository interface {
	LoadChannel(ctx context.Context, name string) (*loaders.Channel, error)
	LoadManifest(ctx context.Context, component string, version string, o declarative.DeclarativeObject) (map[string]string, error)
	LoadNamespacedComponents(ctx context.Context, componentName string, version string) (map[string]string, error)
}

type LocalRepository struct {
	basedir string
}

// Ensure that LocalRepository implements Repository.
var _ Repository = &LocalRepository{}

func NewLocalRepository(basedir string) *LocalRepository {
	return &LocalRepository{
		basedir: basedir,
	}
}

func (r *LocalRepository) LoadChannel(_ context.Context, name string) (*loaders.Channel, error) {
	rlog.Info("loading channel", "base", r.basedir, "name", name)

	p := filepath.Join(r.basedir, name)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		rlog.Error(err, "error reading channel", "path", p)
		return nil, fmt.Errorf("error reading channel %s: %w", p, err)
	}

	channel := &loaders.Channel{}
	if err := yaml.Unmarshal(b, channel); err != nil {
		return nil, fmt.Errorf("error parsing channel %s: %w", p, err)
	}

	return channel, nil
}

func (r *LocalRepository) LoadManifest(_ context.Context, componentName string, version string, o declarative.DeclarativeObject) (map[string]string, error) {
	cc, ok := o.(*corev1beta1.ConfigConnector)
	if !ok {
		return nil, fmt.Errorf("expected the resource to be a ConfigConnector, but it was not. Object: %v", o)
	}
	mode := cc.GetMode()

	p := filepath.Join(r.basedir, "packages", componentName, version, crdFileName)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", p, err)
	}
	var sb strings.Builder
	sb.Write(b)
	sb.WriteString("---\n")
	if mode == corev1beta1.ClusterMode {
		var authIdentity string
		if cc.Spec.GoogleServiceAccount != "" {
			authIdentity = "workload-identity"
		} else {
			authIdentity = "gcp-identity"
		}
		rlog.Info("loading manifest", "component", componentName, "version", version, "mode", mode, "identity", authIdentity)
		p := filepath.Join(r.basedir, "packages", componentName, version, mode, authIdentity, cnrmSystemFileName)
		b, err := ioutil.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", p, err)
		}
		sb.Write(b)
		path := strings.Join([]string{r.basedir, "packages", componentName, version, mode, authIdentity}, "/")
		return map[string]string{path: sb.String()}, nil
	}

	// otherwise we are in namesapce mode
	rlog.Info("loading manifest", "component", componentName, "version", version, "mode", mode)
	p = filepath.Join(r.basedir, "packages", componentName, version, "namespaced", cnrmSystemFileName)
	b, err = os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", p, err)
	}
	sb.Write(b)
	path := strings.Join([]string{r.basedir, "packages", componentName, version, mode}, "/")
	return map[string]string{path: sb.String()}, nil
}

func (r *LocalRepository) LoadNamespacedComponents(_ context.Context, componentName string, version string) (map[string]string, error) {
	p := filepath.Join(r.basedir, "packages", componentName, version, "namespaced", perNamespaceComponentsFileName)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", p, err)
	}
	return map[string]string{p: string(b)}, nil
}
