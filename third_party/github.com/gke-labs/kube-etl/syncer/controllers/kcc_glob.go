// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import "strings"

const (
	// kccGroupSuffix is the suffix of every Config Connector API group.
	kccGroupSuffix = "cnrm.cloud.google.com"
	// kccSystemGroup holds the objects that configure a cluster's Config
	// Connector (ConfigConnector, ConfigConnectorContext). Its subgroups hold
	// related per-cluster objects, such as MultiClusterLease in
	// multicluster.core.cnrm.cloud.google.com.
	kccSystemGroup = "core.cnrm.cloud.google.com"
)

// kccGlobMatchesGroup reports whether the "*.cnrm.cloud.google.com" rule
// matches an API group.
//
// The rule matches Config Connector's resource groups. It skips the system
// group and its subgroups: those objects describe one cluster's Config
// Connector installation and its leader election state, so they must not be
// copied from another cluster. Skipping them also means that the remote
// cluster only has to grant read access to Config Connector resources.
//
// A rule that names one of these groups exactly still matches it.
func kccGlobMatchesGroup(group string) bool {
	if !strings.HasSuffix(group, kccGroupSuffix) {
		return false
	}
	return group != kccSystemGroup && !strings.HasSuffix(group, "."+kccSystemGroup)
}
