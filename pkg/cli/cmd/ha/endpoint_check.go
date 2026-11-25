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
	"errors"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
)

type gkeClusterDescribe struct {
	Endpoint             string `json:"endpoint"`
	PrivateClusterConfig struct {
		PrivateEndpoint          string `json:"privateEndpoint"`
		MasterGlobalAccessConfig *struct {
			Enabled bool `json:"enabled"`
		} `json:"masterGlobalAccessConfig"`
	} `json:"privateClusterConfig"`
	ControlPlaneEndpointsConfig struct {
		DNSEndpointConfig struct {
			Endpoint             string `json:"endpoint"`
			AllowExternalTraffic *bool  `json:"allowExternalTraffic"`
		} `json:"dnsEndpointConfig"`
		IPEndpointsConfig struct {
			PrivateEndpoint string `json:"privateEndpoint"`
			GlobalAccess    *bool  `json:"globalAccess"`
		} `json:"ipEndpointsConfig"`
	} `json:"controlPlaneEndpointsConfig"`
	NetworkConfig struct {
		Network string `json:"network"`
	} `json:"networkConfig"`
	MasterAuthorizedNetworksConfig struct {
		Enabled                           bool `json:"enabled"`
		PrivateEndpointEnforcementEnabled bool `json:"privateEndpointEnforcementEnabled"`
	} `json:"masterAuthorizedNetworksConfig"`
}

func fetchClusterDescribe(ctx context.Context, runner commandRunner, info gkeClusterInfo) (*gkeClusterDescribe, error) {
	out, err := runner.output(ctx, "gcloud", "container", "clusters", "describe", info.clusterName,
		"--location="+info.location,
		"--project="+info.project,
		"--format=json",
	)
	if err != nil {
		return nil, fmt.Errorf("cluster %q: describing cluster: %w", info.contextName, err)
	}

	var desc gkeClusterDescribe
	if err := json.Unmarshal(out, &desc); err != nil {
		return nil, fmt.Errorf("cluster %q: parsing cluster describe JSON: %w", info.contextName, err)
	}
	return &desc, nil
}

func clusterRegion(location string) string {
	loc := strings.ToLower(strings.TrimSpace(location))
	if isZonalGKELocation(loc) {
		lastHyphen := strings.LastIndex(loc, "-")
		return loc[:lastHyphen]
	}
	return loc
}

// checkLeaderEndpoints validates that each cluster exposes the endpoint required by the other
// clusters before any mutations are made.
func checkLeaderEndpoints(ctx context.Context, configurers []*clusterConfigurer, runner commandRunner, stdout, stderr io.Writer) error {
	if len(configurers) == 0 {
		return nil
	}

	descs := make([]*gkeClusterDescribe, len(configurers))
	for i, cfg := range configurers {
		desc, err := fetchClusterDescribe(ctx, runner, cfg.clusterInfo())
		if err != nil {
			return err
		}
		descs[i] = desc
	}

	var errs []error
	failedEndpoints := make(map[string]bool)
	checkedEndpoints := make(map[string][]string)

	var warnings []string
	seenWarnings := make(map[string]bool)
	warn := func(msg string) {
		if !seenWarnings[msg] {
			seenWarnings[msg] = true
			warnings = append(warnings, msg)
		}
	}

	for sIdx, cfgS := range configurers {
		sInfo := cfgS.clusterInfo()
		mcl := cfgS.targetMultiClusterLeaseSpec()
		if mcl == nil {
			continue
		}

		mode := replicationMode(mcl)
		if mode == replicationModeDisabled {
			continue
		}

		sEndpointType := mcl.LeaderEndpointType
		if sEndpointType == "" {
			sEndpointType = defaultLeaderEndpointType
		}

		for lIdx, cfgL := range configurers {
			if sIdx == lIdx {
				continue
			}
			lInfo := cfgL.clusterInfo()
			lDesc := descs[lIdx]

			checkedKey := fmt.Sprintf("%s on %s", sEndpointType, lInfo.contextName)
			if !slices.Contains(checkedEndpoints[sEndpointType], lInfo.contextName) {
				checkedEndpoints[sEndpointType] = append(checkedEndpoints[sEndpointType], lInfo.contextName)
			}

			// 1. Check failure conditions
			switch sEndpointType {
			case operatorv1beta1.LeaderEndpointTypeDNS:
				dns := lDesc.ControlPlaneEndpointsConfig.DNSEndpointConfig
				if dns.Endpoint == "" || dns.AllowExternalTraffic == nil || !*dns.AllowExternalTraffic {
					if !failedEndpoints[checkedKey] {
						failedEndpoints[checkedKey] = true
						errs = append(errs, fmt.Errorf(
							"cluster %q: DNS endpoint is not enabled for external traffic; run `gcloud container clusters update %s --location=%s --project=%s --enable-dns-access`",
							lInfo.contextName, lInfo.clusterName, lInfo.location, lInfo.project,
						))
					}
				}
			case operatorv1beta1.LeaderEndpointTypePrivateIP:
				privateEndpoint := lDesc.ControlPlaneEndpointsConfig.IPEndpointsConfig.PrivateEndpoint
				if privateEndpoint == "" {
					privateEndpoint = lDesc.PrivateClusterConfig.PrivateEndpoint
				}
				if privateEndpoint == "" {
					if !failedEndpoints[checkedKey] {
						failedEndpoints[checkedKey] = true
						errs = append(errs, fmt.Errorf(
							"cluster %q: cluster has no private endpoint; choose --leader-endpoint-type=DNS or Default",
							lInfo.contextName,
						))
					}
				}
			case operatorv1beta1.LeaderEndpointTypeDefault:
				if lDesc.Endpoint == "" {
					if !failedEndpoints[checkedKey] {
						failedEndpoints[checkedKey] = true
						errs = append(errs, fmt.Errorf(
							"cluster %q: cluster has no default endpoint; choose --leader-endpoint-type=DNS",
							lInfo.contextName,
						))
					}
				}
			}

			if failedEndpoints[checkedKey] {
				continue
			}

			// 2. Check warning conditions
			sDesc := descs[sIdx]
			switch sEndpointType {
			case operatorv1beta1.LeaderEndpointTypePrivateIP:
				sRegion := clusterRegion(sInfo.location)
				lRegion := clusterRegion(lInfo.location)
				hasGlobalAccess := (lDesc.PrivateClusterConfig.MasterGlobalAccessConfig != nil && lDesc.PrivateClusterConfig.MasterGlobalAccessConfig.Enabled) ||
					(lDesc.ControlPlaneEndpointsConfig.IPEndpointsConfig.GlobalAccess != nil && *lDesc.ControlPlaneEndpointsConfig.IPEndpointsConfig.GlobalAccess)
				if sRegion != lRegion && !hasGlobalAccess {
					warn(fmt.Sprintf(
						"connections from another region to the private endpoint time out unless global access is on: gcloud container clusters update %s --location=%s --project=%s --enable-master-global-access",
						lInfo.clusterName, lInfo.location, lInfo.project,
					))
				}

				if sDesc.NetworkConfig.Network != "" && lDesc.NetworkConfig.Network != "" && sDesc.NetworkConfig.Network != lDesc.NetworkConfig.Network {
					warn("the clusters need a route between their VPC networks (VPC peering, Cloud VPN or Interconnect)")
				}

				if lDesc.MasterAuthorizedNetworksConfig.PrivateEndpointEnforcementEnabled {
					warn(fmt.Sprintf("add %s's node and Pod ranges to %s's authorized networks", sInfo.contextName, lInfo.contextName))
				}

			case operatorv1beta1.LeaderEndpointTypeDefault:
				if lDesc.MasterAuthorizedNetworksConfig.Enabled {
					warn(fmt.Sprintf("add %s's egress addresses (for example its Cloud NAT addresses) to %s's authorized networks", sInfo.contextName, lInfo.contextName))
				}
			}
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	for _, msg := range warnings {
		if stderr != nil {
			fmt.Fprintf(stderr, "Warning: %s\n", msg)
		}
	}

	if stdout != nil {
		reportCheckedEndpoints(stdout, checkedEndpoints)
	}

	return nil
}

func reportCheckedEndpoints(w io.Writer, checked map[string][]string) {
	if len(checked) == 0 {
		fmt.Fprintln(w, "Checked leader endpoints: none (replication disabled).")
		return
	}

	var epTypes []string
	for epType := range checked {
		epTypes = append(epTypes, epType)
	}
	sort.Strings(epTypes)

	if len(epTypes) == 1 {
		epType := epTypes[0]
		clusters := checked[epType]
		sort.Strings(clusters)
		fmt.Fprintf(w, "Checked %s leader endpoint on %s.\n", epType, strings.Join(clusters, ", "))
		return
	}

	var parts []string
	for _, epType := range epTypes {
		clusters := checked[epType]
		sort.Strings(clusters)
		parts = append(parts, fmt.Sprintf("%s on [%s]", epType, strings.Join(clusters, ", ")))
	}
	fmt.Fprintf(w, "Checked leader endpoints: %s.\n", strings.Join(parts, "; "))
}
