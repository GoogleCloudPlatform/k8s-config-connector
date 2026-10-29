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
	"errors"
	"fmt"
	"io"
	"regexp"
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
	"k8s.io/apimachinery/pkg/util/validation"
)

const contextsFlag = "contexts"

// baseOptions holds shared configuration across `config-connector ha` subcommands.
type baseOptions struct {
	// contexts is the list of kubeconfig context names participating in the HA setup.
	contexts []string

	// stdout and stderr are populated from the Cobra command streams before execution.
	stdout io.Writer
	stderr io.Writer
}

func (o *baseOptions) addFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&o.contexts, contextsFlag, o.contexts, "Comma-separated list of kubeconfig contexts participating in the HA group (at least 2 required).")
}

// validate checks that the shared base options are valid for multi-cluster HA operations
// and normalizes context names in place.
func (o *baseOptions) validate() error {
	for i, c := range o.contexts {
		o.contexts[i] = strings.TrimSpace(c)
	}
	if len(o.contexts) < 2 {
		return fmt.Errorf("at least 2 cluster contexts must be specified via --%s", contextsFlag)
	}
	return validateUniqueContextNames(o.contexts)
}

func validateUniqueContextNames(names []string) error {
	var errs []error
	seenContexts := make(map[string]bool, len(names))
	seenIdentities := make(map[string]string, len(names))
	for _, name := range names {
		if name == "" {
			errs = append(errs, fmt.Errorf("context name in --%s cannot be empty", contextsFlag))
			continue
		}
		if seenContexts[name] {
			errs = append(errs, fmt.Errorf("duplicate context %q specified in --%s", name, contextsFlag))
			continue
		}
		seenContexts[name] = true

		identity, err := sanitizeClusterIdentity(name)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if prevCtx, exists := seenIdentities[identity]; exists {
			errs = append(errs, fmt.Errorf("contexts %q and %q sanitize to the same DNS-1123 cluster identity %q; rename one of the kubeconfig contexts", prevCtx, name, identity))
			continue
		}
		seenIdentities[identity] = name
	}
	return errors.Join(errs...)
}

// sanitizeClusterIdentity derives a DNS-1123 subdomain identity from a kubeconfig context name
// (e.g. "gke_my-proj_us-central1_cluster-a" -> "gke-my-proj-us-central1-cluster-a") so it can be
// used as clusterCandidateIdentity.
func sanitizeClusterIdentity(contextName string) (string, error) {
	identity := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(contextName), "_", "-"))
	if dnsErrs := validation.IsDNS1123Subdomain(identity); len(dnsErrs) > 0 {
		return "", fmt.Errorf("context %q (sanitized to %q) is not a valid DNS-1123 subdomain for cluster identity (%s); rename the kubeconfig context using lowercase alphanumeric characters, '-' or '.'", contextName, identity, strings.Join(dnsErrs, "; "))
	}
	return identity, nil
}

func normalizeBucketName(bucket string) string {
	trimmed := strings.TrimSpace(bucket)
	trimmed = strings.TrimPrefix(trimmed, "gs://")
	trimmed = strings.TrimSuffix(trimmed, "/")
	return strings.TrimSpace(trimmed)
}

var gcsBucketNameRegex = regexp.MustCompile(`^[a-z0-9][a-z0-9._-]{1,61}[a-z0-9]$`)

func validateBucketName(bucket string) error {
	if !gcsBucketNameRegex.MatchString(bucket) {
		return fmt.Errorf("invalid GCS bucket name %q; bucket names must be 3-63 characters of lowercase letters, numbers, '-', '_', or '.', starting and ending with a letter or number", bucket)
	}
	return nil
}

// clusterHandle wraps a kubeconfig context name and its initialized Kubernetes client.
type clusterHandle struct {
	contextName string
	client      *kubecli.Client
}

// buildClusterHandles initializes a Kubernetes client for each kubeconfig context name.
func buildClusterHandles(ctx context.Context, contexts []string) ([]*clusterHandle, error) {
	scheme, err := newHAScheme()
	if err != nil {
		return nil, fmt.Errorf("creating HA runtime scheme: %w", err)
	}
	var errs []error
	handles := make([]*clusterHandle, 0, len(contexts))
	for _, ctxName := range contexts {
		c, err := kubecli.NewClient(ctx, kubecli.ClusterOptions{
			Context: ctxName,
			Scheme:  scheme,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("building kubernetes client for context %q: %w", ctxName, err))
			continue
		}
		handles = append(handles, &clusterHandle{
			contextName: ctxName,
			client:      c,
		})
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return handles, nil
}

// newHAScheme constructs a runtime.Scheme registered with all API types needed by HA commands.
func newHAScheme() (*runtime.Scheme, error) {
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
