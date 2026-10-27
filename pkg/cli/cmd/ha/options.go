// Copyright 2026 Google LLC
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

package ha

import (
	"context"
	"fmt"
	"io"
	"strings"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	"github.com/spf13/pflag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ContextsFlag = "contexts"
	DryRunFlag   = "dry-run"
)

// BaseOptions holds shared configuration across `config-connector ha` subcommands.
type BaseOptions struct {
	// Contexts is the list of kubeconfig context names participating in the HA setup.
	Contexts []string

	// DryRun indicates that mutating subcommands should print planned actions without applying them.
	DryRun bool

	// Stdout and Stderr are populated from the Cobra command streams before execution.
	Stdout io.Writer
	Stderr io.Writer

	// Clusters holds the initialized per-context Kubernetes clients.
	// May be pre-populated in unit tests.
	Clusters []*ClusterHandle
}

func NewBaseOptions() *BaseOptions {
	return &BaseOptions{}
}

func (o *BaseOptions) AddFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&o.Contexts, ContextsFlag, o.Contexts, "Comma-separated list of kubeconfig contexts participating in the HA group (at least 2 required).")
	flags.BoolVar(&o.DryRun, DryRunFlag, o.DryRun, "If true, print planned actions without making changes.")
}

// Validate checks that the shared base options are valid for multi-cluster HA operations
// and normalizes context names in place.
func (o *BaseOptions) Validate() error {
	for i, c := range o.Contexts {
		o.Contexts[i] = strings.TrimSpace(c)
	}
	for _, h := range o.Clusters {
		if h == nil {
			return fmt.Errorf("cluster handle cannot be nil")
		}
		h.ContextName = strings.TrimSpace(h.ContextName)
	}

	clusterNames := o.ClusterNames()
	if len(clusterNames) < 2 {
		return fmt.Errorf("at least 2 cluster contexts must be specified via --%s", ContextsFlag)
	}
	return validateUniqueContextNames(clusterNames)
}

func validateUniqueContextNames(names []string) error {
	seen := make(map[string]bool, len(names))
	for _, name := range names {
		if name == "" {
			return fmt.Errorf("context name in --%s cannot be empty", ContextsFlag)
		}
		if seen[name] {
			return fmt.Errorf("duplicate context %q specified in --%s", name, ContextsFlag)
		}
		seen[name] = true
	}
	return nil
}

// ClusterNames returns the configured kubeconfig context names from either Contexts or pre-populated Clusters.
func (o *BaseOptions) ClusterNames() []string {
	if len(o.Contexts) > 0 {
		return o.Contexts
	}
	names := make([]string, 0, len(o.Clusters))
	for _, h := range o.Clusters {
		if h != nil {
			names = append(names, h.ContextName)
		}
	}
	return names
}

// InitClusterHandles initializes a Kubernetes client for each context in o.Contexts
// if o.Clusters has not already been pre-populated (e.g. in tests).
func (o *BaseOptions) InitClusterHandles(ctx context.Context) error {
	if len(o.Clusters) > 0 {
		return nil
	}
	scheme, err := NewHAScheme()
	if err != nil {
		return fmt.Errorf("creating HA runtime scheme: %w", err)
	}
	handles := make([]*ClusterHandle, 0, len(o.Contexts))
	for _, ctxName := range o.Contexts {
		c, err := kubecli.NewClient(ctx, kubecli.ClusterOptions{
			Context: ctxName,
			Scheme:  scheme,
		})
		if err != nil {
			return fmt.Errorf("building kubernetes client for context %q: %w", ctxName, err)
		}
		handles = append(handles, &ClusterHandle{
			ContextName: ctxName,
			Client:      c,
		})
	}
	o.Clusters = handles
	return nil
}

// ClusterHandle wraps a kubeconfig context name and its initialized Kubernetes client.
type ClusterHandle struct {
	ContextName string
	Client      *kubecli.Client
}

// NewHAScheme constructs a runtime.Scheme registered with all API types needed by HA commands.
func NewHAScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	builders := []func(*runtime.Scheme) error{
		corev1.AddToScheme,
		appsv1.AddToScheme,
		rbacv1.AddToScheme,
		apiextensionsv1.AddToScheme,
		operatorv1beta1.AddToScheme,
		mclv1alpha1.AddToScheme,
		syncerv1alpha1.AddToScheme,
	}
	for _, add := range builders {
		if err := add(scheme); err != nil {
			return nil, err
		}
	}
	return scheme, nil
}
