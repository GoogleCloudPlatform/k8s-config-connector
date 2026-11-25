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

package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// CanonicalResourceFilter is a normalized representation of ResourceFilter for deterministic hashing.
type CanonicalResourceFilter struct {
	Group string `json:"group"`
	Kind  string `json:"kind,omitempty"`
}

// CanonicalResourceSettings is a normalized representation of ResourceSettings for deterministic hashing.
type CanonicalResourceSettings struct {
	Mode      string                    `json:"mode"`
	Resources []CanonicalResourceFilter `json:"resources,omitempty"`
}

// CanonicalGKECluster is a normalized representation of GKECluster for deterministic hashing.
type CanonicalGKECluster struct {
	Project  string `json:"project"`
	Location string `json:"location"`
	Name     string `json:"name"`
}

// CanonicalMultiClusterLease is a normalized representation of MultiClusterLeaseSpec for deterministic hashing.
type CanonicalMultiClusterLease struct {
	LeaseName               string               `json:"leaseName"`
	Namespace               string               `json:"namespace"`
	GKECluster              *CanonicalGKECluster `json:"gkeCluster,omitempty"`
	LeaderEndpointType      string               `json:"leaderEndpointType"`
	ResourceReplicationMode string               `json:"resourceReplicationMode"`
}

// CanonicalCCConfig is hashed to produce "cnrm.cloud.google.com/cc-config-hash".
type CanonicalCCConfig struct {
	ResourceSettings                 *CanonicalResourceSettings  `json:"resourceSettings,omitempty"`
	MaxConcurrentReconcilesOverrides map[string]int              `json:"maxConcurrentReconcilesOverrides,omitempty"`
	MultiClusterLease                *CanonicalMultiClusterLease `json:"multiClusterLease,omitempty"`
}

// CanonicalCCCConfig is hashed to produce "cnrm.cloud.google.com/ccc-config-hash".
type CanonicalCCCConfig struct {
	ResourceSettings                 *CanonicalResourceSettings `json:"resourceSettings,omitempty"`
	MaxConcurrentReconcilesOverrides map[string]int             `json:"maxConcurrentReconcilesOverrides,omitempty"`
}

// ComputeCCConfigHash returns a deterministic SHA-256 hex digest for the restart-dependent
// configuration of a ConfigConnector object. Returns "" if no hashed configuration is present.
func ComputeCCConfigHash(cc *corev1beta1.ConfigConnector) string {
	if cc == nil || cc.Spec.Experiments == nil {
		return ""
	}
	cfg := CanonicalCCConfig{
		ResourceSettings:                 canonicalizeResourceSettings(cc.Spec.Experiments.ResourceSettings),
		MaxConcurrentReconcilesOverrides: canonicalizeMaxConcurrentReconcilesOverrides(cc.Spec.Experiments.MaxConcurrentReconcilesOverrides),
		MultiClusterLease:                canonicalizeMultiClusterLease(cc.Spec.Experiments.MultiClusterLease),
	}
	if cfg.ResourceSettings == nil && len(cfg.MaxConcurrentReconcilesOverrides) == 0 && cfg.MultiClusterLease == nil {
		return ""
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// ComputeCCCConfigHash returns a deterministic SHA-256 hex digest for the restart-dependent
// configuration of a ConfigConnectorContext object. Returns "" if no hashed configuration is present.
func ComputeCCCConfigHash(ccc *corev1beta1.ConfigConnectorContext) string {
	if ccc == nil || ccc.Spec.Experiments == nil {
		return ""
	}
	cfg := CanonicalCCCConfig{
		ResourceSettings:                 canonicalizeResourceSettings(ccc.Spec.Experiments.ResourceSettings),
		MaxConcurrentReconcilesOverrides: canonicalizeMaxConcurrentReconcilesOverrides(ccc.Spec.Experiments.MaxConcurrentReconcilesOverrides),
	}
	if cfg.ResourceSettings == nil && len(cfg.MaxConcurrentReconcilesOverrides) == 0 {
		return ""
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func canonicalizeMaxConcurrentReconcilesOverrides(overrides map[string]int) map[string]int {
	if len(overrides) == 0 {
		return nil
	}
	out := make(map[string]int, len(overrides))
	for k, v := range overrides {
		if k != "" && v > 0 {
			out[k] = v
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func canonicalizeResourceSettings(rs *corev1beta1.ResourceSettings) *CanonicalResourceSettings {
	if rs == nil {
		return nil
	}
	mode := strings.ToLower(strings.TrimSpace(string(rs.Mode)))
	if mode == "" {
		mode = string(corev1beta1.ResourceSettingsModeExclude)
	}

	filters := make([]CanonicalResourceFilter, 0, len(rs.Resources))
	for _, r := range rs.Resources {
		var group, kind string
		if r.Group != nil {
			group = strings.TrimSpace(*r.Group)
		}
		if r.Kind != nil {
			kind = strings.TrimSpace(*r.Kind)
		}
		if group == "" && kind == "" {
			continue
		}
		filters = append(filters, CanonicalResourceFilter{
			Group: group,
			Kind:  kind,
		})
	}

	sort.Slice(filters, func(i, j int) bool {
		if filters[i].Group != filters[j].Group {
			return filters[i].Group < filters[j].Group
		}
		return filters[i].Kind < filters[j].Kind
	})

	deduped := make([]CanonicalResourceFilter, 0, len(filters))
	for i, f := range filters {
		if i > 0 && f.Group == filters[i-1].Group && f.Kind == filters[i-1].Kind {
			continue
		}
		deduped = append(deduped, f)
	}

	return &CanonicalResourceSettings{
		Mode:      mode,
		Resources: deduped,
	}
}

func canonicalizeMultiClusterLease(lease *corev1beta1.MultiClusterLeaseSpec) *CanonicalMultiClusterLease {
	if lease == nil {
		return nil
	}
	leaderEndpointType := lease.LeaderEndpointType
	if leaderEndpointType == "" {
		leaderEndpointType = corev1beta1.LeaderEndpointTypeDefault
	}
	resourceReplicationMode := lease.ResourceReplicationMode
	if resourceReplicationMode == "" {
		resourceReplicationMode = "Status"
	}
	var gkeCluster *CanonicalGKECluster
	if lease.GKECluster != nil {
		gkeCluster = &CanonicalGKECluster{
			Project:  lease.GKECluster.Project,
			Location: lease.GKECluster.Location,
			Name:     lease.GKECluster.Name,
		}
	}
	return &CanonicalMultiClusterLease{
		LeaseName:               lease.LeaseName,
		Namespace:               lease.Namespace,
		GKECluster:              gkeCluster,
		LeaderEndpointType:      leaderEndpointType,
		ResourceReplicationMode: resourceReplicationMode,
	}
}

// SetPodTemplateAnnotation sets or removes an annotation under spec.template.metadata.annotations
// on an unstructured workload object (such as a StatefulSet). If value is empty, the annotation key
// is removed if present.
func SetPodTemplateAnnotation(u *unstructured.Unstructured, key, value string) error {
	annotations, found, err := unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if err != nil {
		return fmt.Errorf("failed to get pod template annotations: %w", err)
	}
	if value == "" {
		if !found || annotations == nil {
			return nil
		}
		if _, exists := annotations[key]; !exists {
			return nil
		}
		delete(annotations, key)
		if len(annotations) == 0 {
			unstructured.RemoveNestedField(u.Object, "spec", "template", "metadata", "annotations")
			return nil
		}
		if err := unstructured.SetNestedStringMap(u.Object, annotations, "spec", "template", "metadata", "annotations"); err != nil {
			return fmt.Errorf("failed to set pod template annotations: %w", err)
		}
		return nil
	}

	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[key] = value
	if err := unstructured.SetNestedStringMap(u.Object, annotations, "spec", "template", "metadata", "annotations"); err != nil {
		return fmt.Errorf("failed to set pod template annotations: %w", err)
	}
	return nil
}
