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
	"sort"
	"strings"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	configConnectorResourceName = "configconnector.core.cnrm.cloud.google.com"
)

// clusterHAState holds the observed HA-related resources on a single cluster.
type clusterHAState struct {
	contextName       string
	configConnector   *operatorv1beta1.ConfigConnector
	multiClusterLease *mclv1alpha1.MultiClusterLease
	krmSyncers        []syncerv1alpha1.KRMSyncer
}

// fetchAllClusterHAStates fetches the HA state across all configured clusters.
func fetchAllClusterHAStates(ctx context.Context, handles []*clusterHandle) ([]*clusterHAState, error) {
	var errs []error
	states := make([]*clusterHAState, 0, len(handles))
	for _, h := range handles {
		state, err := fetchClusterHAState(ctx, h)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		states = append(states, state)
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return states, nil
}

// fetchClusterHAState reads and validates the ConfigConnector, MultiClusterLease, and KRMSyncer objects from a cluster.
func fetchClusterHAState(ctx context.Context, handle *clusterHandle) (*clusterHAState, error) {
	if handle == nil || handle.client == nil {
		return nil, fmt.Errorf("cluster handle or client is nil")
	}

	cc, err := fetchConfigConnector(ctx, handle)
	if err != nil {
		return nil, err
	}

	mclSpec := multiClusterLeaseSpec(cc)
	if err := validateClusterLeaseSpec(handle.contextName, mclSpec); err != nil {
		return nil, err
	}

	lease, err := fetchMultiClusterLease(ctx, handle, mclSpec.Namespace, mclSpec.LeaseName)
	if err != nil {
		return nil, err
	}

	syncers, err := listKRMSyncers(ctx, handle, mclSpec.Namespace)
	if err != nil {
		return nil, err
	}

	return &clusterHAState{
		contextName:       handle.contextName,
		configConnector:   cc,
		multiClusterLease: lease,
		krmSyncers:        syncers,
	}, nil
}

func fetchConfigConnector(ctx context.Context, handle *clusterHandle) (*operatorv1beta1.ConfigConnector, error) {
	cc := &operatorv1beta1.ConfigConnector{}
	key := types.NamespacedName{Name: configConnectorResourceName}
	if err := handle.client.Get(ctx, key, cc); err != nil {
		if isResourceOrKindAbsent(err) {
			return nil, fmt.Errorf(
				"cluster %q is missing the %q ConfigConnector object; install Config Connector and create the %q resource before running HA commands",
				handle.contextName, configConnectorResourceName, configConnectorResourceName,
			)
		}

		return nil, fmt.Errorf("fetching ConfigConnector on cluster %q: %w", handle.contextName, err)
	}

	return cc, nil
}

func multiClusterLeaseSpec(cc *operatorv1beta1.ConfigConnector) *operatorv1beta1.MultiClusterLeaseSpec {
	if cc == nil || cc.Spec.Experiments == nil {
		return nil
	}
	return cc.Spec.Experiments.MultiClusterLease
}

func validateClusterLeaseSpec(contextName string, mcl *operatorv1beta1.MultiClusterLeaseSpec) error {
	if mcl == nil {
		return fmt.Errorf("cluster %q does not have spec.experiments.multiClusterLease configured on ConfigConnector %q; run `config-connector ha setup` or configure spec.experiments.multiClusterLease", contextName, configConnectorResourceName)
	}
	var errs []error
	if mcl.LeaseName == "" {
		errs = append(errs, fmt.Errorf("cluster %q has an empty spec.experiments.multiClusterLease.leaseName; set it to the shared lease name across all HA clusters", contextName))
	}
	if mcl.Namespace == "" {
		errs = append(errs, fmt.Errorf("cluster %q has an empty spec.experiments.multiClusterLease.namespace; set it to the shared lease namespace (e.g. %q) across all HA clusters", contextName, defaultLeaseNamespace))
	}
	if mcl.ClusterCandidateIdentity == "" {
		errs = append(errs, fmt.Errorf("cluster %q has an empty spec.experiments.multiClusterLease.clusterCandidateIdentity; set a unique DNS-1123 subdomain identity per cluster", contextName))
	} else if dnsErrs := validation.IsDNS1123Subdomain(mcl.ClusterCandidateIdentity); len(dnsErrs) > 0 {
		errs = append(errs, fmt.Errorf("cluster %q has an invalid spec.experiments.multiClusterLease.clusterCandidateIdentity %q (used as the peer kubeconfig Secret name, must be a valid DNS-1123 subdomain: %s)", contextName, mcl.ClusterCandidateIdentity, strings.Join(dnsErrs, "; ")))
	}
	return errors.Join(errs...)
}

func fetchMultiClusterLease(ctx context.Context, handle *clusterHandle, namespace, name string) (*mclv1alpha1.MultiClusterLease, error) {
	lease := &mclv1alpha1.MultiClusterLease{}
	key := types.NamespacedName{Namespace: namespace, Name: name}
	if err := handle.client.Get(ctx, key, lease); err != nil {
		if isResourceOrKindAbsent(err) {
			return nil, nil
		}

		return nil, fmt.Errorf("fetching MultiClusterLease %s on cluster %q: %w", key, handle.contextName, err)
	}

	return lease, nil
}

func listKRMSyncers(ctx context.Context, handle *clusterHandle, namespace string) ([]syncerv1alpha1.KRMSyncer, error) {
	syncerList := &syncerv1alpha1.KRMSyncerList{}
	if err := handle.client.List(ctx, syncerList, client.InNamespace(namespace)); err != nil {
		if isResourceOrKindAbsent(err) {
			return nil, nil
		}

		return nil, fmt.Errorf("listing KRMSyncers in namespace %q on cluster %q: %w", namespace, handle.contextName, err)
	}

	return syncerList.Items, nil
}

func isResourceOrKindAbsent(err error) bool {
	return apierrors.IsNotFound(err) || meta.IsNoMatchError(err) || runtime.IsNotRegisteredError(err)
}

// validateMultiClusterLeaseConfigs checks that all clusters have consistent MultiClusterLease
// settings (identical leaseName and namespace, and unique clusterCandidateIdentity per cluster).
func validateMultiClusterLeaseConfigs(states []*clusterHAState) error {
	if len(states) < 2 {
		return fmt.Errorf("at least 2 clusters are required for multi-cluster HA validation, got %d", len(states))
	}

	var errs []error
	leaseNames := make(map[string][]string, len(states))
	namespaces := make(map[string][]string, len(states))
	identities := make(map[string][]string, len(states))

	for _, s := range states {
		if s == nil {
			errs = append(errs, fmt.Errorf("cluster HA state is nil"))
			continue
		}
		mcl := multiClusterLeaseSpec(s.configConnector)
		if err := validateClusterLeaseSpec(s.contextName, mcl); err != nil {
			errs = append(errs, err)
			continue
		}
		leaseNames[mcl.LeaseName] = append(leaseNames[mcl.LeaseName], s.contextName)
		namespaces[mcl.Namespace] = append(namespaces[mcl.Namespace], s.contextName)
		identities[mcl.ClusterCandidateIdentity] = append(identities[mcl.ClusterCandidateIdentity], s.contextName)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	if len(leaseNames) != 1 {
		errs = append(errs, fmt.Errorf(
			"mismatched leaseName across clusters (%s): set spec.experiments.multiClusterLease.leaseName to the same value on every cluster because the GCS lock object is leases/<namespace>/<leaseName>",
			formatClusterLeaseValues(states, func(m *operatorv1beta1.MultiClusterLeaseSpec) string { return m.LeaseName }),
		))
	}
	if len(namespaces) != 1 {
		errs = append(errs, fmt.Errorf(
			"mismatched lease namespace across clusters (%s): set spec.experiments.multiClusterLease.namespace to the same value on every cluster because the GCS lock object is leases/<namespace>/<leaseName>",
			formatClusterLeaseValues(states, func(m *operatorv1beta1.MultiClusterLeaseSpec) string { return m.Namespace }),
		))
	}
	if len(identities) != len(states) {
		errs = append(errs, fmt.Errorf(
			"duplicate clusterCandidateIdentity across clusters (%s): set a unique spec.experiments.multiClusterLease.clusterCandidateIdentity on each cluster",
			formatDuplicateIdentities(identities),
		))
	}

	return errors.Join(errs...)
}

func formatClusterLeaseValues(states []*clusterHAState, field func(*operatorv1beta1.MultiClusterLeaseSpec) string) string {
	parts := make([]string, 0, len(states))
	for _, s := range states {
		mcl := multiClusterLeaseSpec(s.configConnector)
		parts = append(parts, fmt.Sprintf("%s=%q", s.contextName, field(mcl)))
	}
	return strings.Join(parts, ", ")
}

func formatDuplicateIdentities(identities map[string][]string) string {
	keys := make([]string, 0, len(identities))
	for id, clusters := range identities {
		if len(clusters) > 1 {
			keys = append(keys, id)
		}
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, id := range keys {
		parts = append(parts, fmt.Sprintf("identity %q shared by [%s]", id, strings.Join(identities[id], ", ")))
	}
	return strings.Join(parts, "; ")
}
