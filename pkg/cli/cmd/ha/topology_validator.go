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
	"fmt"
	"strings"
)

// topologyValidator checks cross-cluster HA configuration consistency and global leader agreement.
type topologyValidator struct {
	reports     []clusterStatusReport
	validStates []*clusterHAState
}

func newTopologyValidator(reports []clusterStatusReport, validStates []*clusterHAState) *topologyValidator {
	return &topologyValidator{
		reports:     reports,
		validStates: validStates,
	}
}

func (v *topologyValidator) validateClusterTopology() []string {
	if len(v.validStates) < 2 {
		return nil
	}

	err := validateMultiClusterLeaseConfigs(v.validStates)
	if err == nil {
		return nil
	}

	var lines []string
	for _, line := range strings.Split(err.Error(), "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			lines = append(lines, trimmed)
		}
	}
	return lines
}

const (
	topologyDefaultLeaseNamespace = "cnrm-system"
	topologyDefaultLeaseName      = "kcc-multicluster-lease"
)

func (v *topologyValidator) resolveExpectedLeaseCoordinates() (string, string) {
	for _, r := range v.reports {
		if r.leaseNamespace != "" && r.leaseName != "" {
			return r.leaseNamespace, r.leaseName
		}
	}

	return topologyDefaultLeaseNamespace, topologyDefaultLeaseName
}

func (v *topologyValidator) checkLeaderConsistency(gcsStatus *gcsLockStatus) []string {
	var warnings []string
	leaders := make(map[string][]string)
	activeLeaders := make([]string, 0)

	for _, r := range v.reports {
		if r.observedLeader != "" && !strings.HasPrefix(r.observedLeader, "<") {
			leaders[r.observedLeader] = append(leaders[r.observedLeader], r.contextName)
		}
		if r.role == "LEADER" {
			activeLeaders = append(activeLeaders, fmt.Sprintf("%s (%s)", r.contextName, r.candidateIdentity))
		}
		if gcsStatus != nil && gcsStatus.found && gcsStatus.holderIdentity != "" &&
			r.observedLeader != "" && !strings.HasPrefix(r.observedLeader, "<") &&
			r.observedLeader != gcsStatus.holderIdentity {
			warnings = append(warnings, fmt.Sprintf(
				"leader mismatch on cluster %q: MultiClusterLease status.globalHolderIdentity is %q, but GCS lock holder is %q",
				r.contextName, r.observedLeader, gcsStatus.holderIdentity,
			))
		}
	}

	if len(leaders) > 1 {
		parts := make([]string, 0, len(v.reports))
		for _, r := range v.reports {
			parts = append(parts, fmt.Sprintf("%s=%q", r.contextName, r.observedLeader))
		}
		warnings = append(warnings, fmt.Sprintf("clusters disagree on global leader (%s)", strings.Join(parts, ", ")))
	}
	if len(activeLeaders) > 1 {
		warnings = append(warnings, fmt.Sprintf("multiple clusters report role LEADER simultaneously: [%s]", strings.Join(activeLeaders, ", ")))
	}

	return warnings
}
