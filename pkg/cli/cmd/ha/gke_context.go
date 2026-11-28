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
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	connectGatewayHostRegex = regexp.MustCompile(`^https://(?:[a-z0-9-]+-)?connectgateway\.googleapis\.com/v[^/]+/projects/([^/]+)/locations/([^/]+)/gkeMemberships/([^/]+)/?$`)
	gkeResourceLinkRegex    = regexp.MustCompile(`projects/([^/]+)/(?:locations|zones)/([^/]+)/clusters/([^/]+)`)
)

type fleetMembershipDescribe struct {
	Endpoint *struct {
		GKECluster *struct {
			ResourceLink string `json:"resourceLink"`
		} `json:"gkeCluster"`
	} `json:"endpoint"`
}

// gkeClusterInfo holds the GCP project, location, and cluster name parsed from a GKE context.
type gkeClusterInfo struct {
	contextName string
	project     string
	location    string
	clusterName string
}

// gkeCluster returns the GKE cluster reference used as the cluster's identity in the
// multi-cluster lease election (spec.experiments.multiClusterLease.gkeCluster).
func (i gkeClusterInfo) gkeCluster() operatorv1beta1.GKECluster {
	return operatorv1beta1.GKECluster{
		Project:  i.project,
		Location: i.location,
		Name:     i.clusterName,
	}
}

// resolveGKECluster determines the GKE cluster coordinates (project, location, cluster name)
// for a cluster handle.
//
// Fast path:
//  1. If handle.contextName matches "gke_<project>_<location>_<cluster>", it is parsed directly
//     with zero API calls.
//  2. If handle's underlying kubeconfig context.cluster matches "gke_<project>_<location>_<cluster>",
//     it is parsed with zero API calls, supporting `kubectl config rename-context` aliases.
//
// Connect Gateway path:
//  3. If the cluster endpoint host matches Connect Gateway
//     (https://[<loc>-]connectgateway.googleapis.com/v.../projects/<fleet-proj>/locations/<mem-loc>/gkeMemberships/<membership>),
//     it runs read-only `gcloud container fleet memberships describe` to retrieve
//     endpoint.gkeCluster.resourceLink. If the resource link uses a numeric project number,
//     it resolves the GCP project ID via `gcloud projects describe <proj> --format=value(projectId)`.
func resolveGKECluster(ctx context.Context, handle *clusterHandle, runner commandRunner) (gkeClusterInfo, error) {
	if handle == nil {
		return gkeClusterInfo{}, fmt.Errorf("cluster handle is nil")
	}
	if handle.gkeInfo != nil {
		return *handle.gkeInfo, nil
	}

	trimmedContext := strings.TrimSpace(handle.contextName)
	parts := strings.Split(trimmedContext, "_")
	if isValidGKEContextParts(parts) {
		info := gkeClusterInfo{
			contextName: trimmedContext,
			project:     parts[1],
			location:    parts[2],
			clusterName: parts[3],
		}
		handle.gkeInfo = &info
		return info, nil
	}

	clusterName := handle.clusterName
	var serverURL string
	if clusterName == "" || (handle.client == nil || handle.client.RestConfig == nil || handle.client.RestConfig.Host == "") {
		cfgCluster, cfgServer := getKubeconfigContext(trimmedContext)
		if clusterName == "" {
			clusterName = cfgCluster
		}
		if serverURL == "" {
			serverURL = cfgServer
		}
	}

	clusterParts := strings.Split(strings.TrimSpace(clusterName), "_")
	if isValidGKEContextParts(clusterParts) {
		info := gkeClusterInfo{
			contextName: trimmedContext,
			project:     clusterParts[1],
			location:    clusterParts[2],
			clusterName: clusterParts[3],
		}
		handle.gkeInfo = &info
		return info, nil
	}

	host := ""
	if handle.client != nil && handle.client.RestConfig != nil && handle.client.RestConfig.Host != "" {
		host = strings.TrimSpace(handle.client.RestConfig.Host)
	} else if serverURL != "" {
		host = strings.TrimSpace(serverURL)
	}

	if match := connectGatewayHostRegex.FindStringSubmatch(host); len(match) == 4 {
		fleetProj := match[1]
		memLoc := match[2]
		membership := match[3]

		if runner == nil {
			runner = newExecCommandRunner(nil, nil, false)
		}

		out, err := runner.output(ctx, "gcloud", "container", "fleet", "memberships", "describe", membership, "--location="+memLoc, "--project="+fleetProj, "--format=json")
		if err != nil {
			return gkeClusterInfo{}, fmt.Errorf("describing fleet membership %q (project %q, location %q): %w", membership, fleetProj, memLoc, err)
		}

		var desc fleetMembershipDescribe
		if err := json.Unmarshal(out, &desc); err != nil {
			return gkeClusterInfo{}, fmt.Errorf("parsing fleet membership describe JSON for %q: %w", membership, err)
		}

		if desc.Endpoint == nil || desc.Endpoint.GKECluster == nil || strings.TrimSpace(desc.Endpoint.GKECluster.ResourceLink) == "" {
			return gkeClusterInfo{}, fmt.Errorf("fleet membership %q in project %q is not a GKE cluster; multi-cluster HA requires GKE clusters", membership, fleetProj)
		}

		resourceLink := strings.TrimSpace(desc.Endpoint.GKECluster.ResourceLink)
		linkMatches := gkeResourceLinkRegex.FindStringSubmatch(resourceLink)
		if len(linkMatches) != 4 {
			return gkeClusterInfo{}, fmt.Errorf("parsing GKE cluster resourceLink %q for membership %q: invalid format", resourceLink, membership)
		}

		proj := linkMatches[1]
		loc := linkMatches[2]
		cluster := linkMatches[3]

		if isNumeric(proj) {
			pOut, err := runner.output(ctx, "gcloud", "projects", "describe", proj, "--format=value(projectId)")
			if err != nil {
				return gkeClusterInfo{}, fmt.Errorf("resolving project ID for project number %q: %w", proj, err)
			}
			projID := strings.TrimSpace(string(pOut))
			if projID == "" {
				return gkeClusterInfo{}, fmt.Errorf("resolving project ID for project number %q returned empty project ID", proj)
			}
			proj = projID
		}

		info := gkeClusterInfo{
			contextName: trimmedContext,
			project:     proj,
			location:    loc,
			clusterName: cluster,
		}
		handle.gkeInfo = &info
		return info, nil
	}

	return gkeClusterInfo{}, fmt.Errorf(
		"context %q does not match expected GKE context format \"gke_<project>_<location>_<cluster>\"; run `gcloud container clusters get-credentials <cluster> --location=<location> --project=<project>`",
		handle.contextName,
	)
}

func getKubeconfigContext(contextName string) (clusterName, server string) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	cfg, err := loadingRules.Load()
	if err != nil || cfg == nil || cfg.Contexts == nil {
		return "", ""
	}
	ctxObj := cfg.Contexts[contextName]
	if ctxObj == nil {
		return "", ""
	}
	clusterName = ctxObj.Cluster
	if cfg.Clusters != nil {
		if clObj := cfg.Clusters[clusterName]; clObj != nil {
			server = clObj.Server
		}
	}
	return clusterName, server
}

func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// parseGKEContext parses a standard gcloud kubeconfig context name of the form
// "gke_<project>_<location>_<cluster-name>".
func parseGKEContext(contextName string) (gkeClusterInfo, error) {
	trimmed := strings.TrimSpace(contextName)
	parts := strings.Split(trimmed, "_")
	if !isValidGKEContextParts(parts) {
		return gkeClusterInfo{}, fmt.Errorf(
			"context %q does not match expected GKE context format \"gke_<project>_<location>_<cluster>\"; run `gcloud container clusters get-credentials <cluster> --location=<location> --project=<project>`",
			contextName,
		)
	}

	return gkeClusterInfo{
		contextName: trimmed,
		project:     parts[1],
		location:    parts[2],
		clusterName: parts[3],
	}, nil
}

func isValidGKEContextParts(parts []string) bool {
	if len(parts) != 4 {
		return false
	}

	hasGKEPrefix := parts[0] == "gke"
	hasNonEmptySegments := parts[1] != "" && parts[2] != "" && parts[3] != ""
	return hasGKEPrefix && hasNonEmptySegments
}
