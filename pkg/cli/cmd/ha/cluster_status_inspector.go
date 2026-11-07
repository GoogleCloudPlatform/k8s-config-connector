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
	"strings"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
)

type syncerSummary struct {
	namespace string
	name      string
	mode      string
	suspended bool
	source    string
	health    string
	errors    []string
}

type clusterStatusReport struct {
	contextName       string
	candidateIdentity string
	leaseNamespace    string
	leaseName         string
	replicationMode   string
	role              string
	observedLeader    string
	renewTime         string
	renewAge          string
	transitions       int32
	syncers           []syncerSummary
	rawState          *clusterHAState
}

// clusterStatusInspector inspects the multi-cluster HA Kubernetes state on a single cluster.
type clusterStatusInspector struct {
	handle *clusterHandle
	now    time.Time
}

func newClusterStatusInspector(handle *clusterHandle, now time.Time) *clusterStatusInspector {
	return &clusterStatusInspector{
		handle: handle,
		now:    now,
	}
}

func (c *clusterStatusInspector) inspect(ctx context.Context) (clusterStatusReport, *clusterHAState, []string) {
	report := newDefaultClusterStatusReport()
	if c.handle == nil || c.handle.client == nil {
		return report, nil, []string{"cluster handle or client is nil"}
	}
	report.contextName = c.handle.contextName

	cc, mclSpec, specValid, warnings := c.inspectConfigConnector(ctx, &report)
	if mclSpec == nil {
		return report, nil, warnings
	}

	lease, leaseWarns := c.inspectMultiClusterLease(ctx, mclSpec, &report)
	warnings = append(warnings, leaseWarns...)

	syncers, syncerWarns := c.inspectKRMSyncers(ctx, mclSpec.Namespace, &report)
	warnings = append(warnings, syncerWarns...)

	var validState *clusterHAState
	if specValid {
		validState = &clusterHAState{
			contextName:       c.handle.contextName,
			configConnector:   cc,
			multiClusterLease: lease,
			krmSyncers:        syncers,
		}
		report.rawState = validState
	}

	return report, validState, warnings
}

func newDefaultClusterStatusReport() clusterStatusReport {
	return clusterStatusReport{
		contextName:       "<unknown>",
		candidateIdentity: "<unconfigured>",
		role:              "UNCONFIGURED",
		observedLeader:    "<none>",
		renewAge:          "<unknown>",
	}
}

const defaultStatusReplicationMode = "Status"

func (c *clusterStatusInspector) inspectConfigConnector(ctx context.Context, report *clusterStatusReport) (*operatorv1beta1.ConfigConnector, *operatorv1beta1.MultiClusterLeaseSpec, bool, []string) {
	cc, err := fetchConfigConnector(ctx, c.handle)
	if err != nil {
		return nil, nil, false, []string{err.Error()}
	}
	if cc == nil {
		return nil, nil, false, []string{fmt.Sprintf("cluster %q is missing the %q ConfigConnector object", c.handle.contextName, configConnectorResourceName)}
	}

	var warnings []string
	mclSpec := multiClusterLeaseSpec(cc)
	specErr := validateClusterLeaseSpec(c.handle.contextName, mclSpec)
	if specErr != nil {
		warnings = append(warnings, splitErrorLines(specErr)...)
		if mclSpec == nil {
			return cc, nil, false, warnings
		}
	}

	report.candidateIdentity = mclSpec.ClusterCandidateIdentity
	report.leaseNamespace = mclSpec.Namespace
	report.leaseName = mclSpec.LeaseName
	report.replicationMode = mclSpec.ResourceReplicationMode
	if report.replicationMode == "" {
		report.replicationMode = defaultStatusReplicationMode
	}
	report.role = "UNDECIDED"

	return cc, mclSpec, specErr == nil, warnings
}

func splitErrorLines(err error) []string {
	var lines []string
	for _, line := range strings.Split(err.Error(), "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			lines = append(lines, trimmed)
		}
	}
	return lines
}

func (c *clusterStatusInspector) inspectMultiClusterLease(ctx context.Context, mclSpec *operatorv1beta1.MultiClusterLeaseSpec, report *clusterStatusReport) (*mclv1alpha1.MultiClusterLease, []string) {
	if mclSpec.Namespace == "" || mclSpec.LeaseName == "" {
		return nil, nil
	}

	var warnings []string
	lease, err := fetchMultiClusterLease(ctx, c.handle, mclSpec.Namespace, mclSpec.LeaseName)
	if err != nil {
		warnings = append(warnings, err.Error())
	}
	if lease == nil {
		warnings = append(warnings, fmt.Sprintf("cluster %q is missing MultiClusterLease \"%s/%s\"", c.handle.contextName, mclSpec.Namespace, mclSpec.LeaseName))
		return nil, warnings
	}

	warnings = append(warnings, c.populateLeaseStatus(report, lease)...)
	return lease, warnings
}

func (c *clusterStatusInspector) populateLeaseStatus(report *clusterStatusReport, lease *mclv1alpha1.MultiClusterLease) []string {
	var warnings []string
	holder := ptr.Deref(lease.Status.GlobalHolderIdentity, "")
	if holder != "" {
		report.observedLeader = holder
		if holder == report.candidateIdentity {
			report.role = "LEADER"
		} else {
			report.role = "STANDBY"
		}
	} else {
		warnings = append(warnings, fmt.Sprintf("cluster %q MultiClusterLease \"%s/%s\" has an empty status.globalHolderIdentity", report.contextName, lease.Namespace, lease.Name))
	}

	report.transitions = ptr.Deref(lease.Status.GlobalLeaseTransitions, 0)
	warnings = append(warnings, c.evaluateLeaseRenewal(report, lease)...)
	return warnings
}

const (
	statusDefaultLeaseDurationSeconds = int32(15)
	statusStaleLeaseMultiplier        = 2
)

func (c *clusterStatusInspector) evaluateLeaseRenewal(report *clusterStatusReport, lease *mclv1alpha1.MultiClusterLease) []string {
	leaseDurationSec := ptr.Deref(lease.Status.GlobalLeaseDurationSeconds, statusDefaultLeaseDurationSeconds)
	if leaseDurationSec <= 0 {
		leaseDurationSec = statusDefaultLeaseDurationSeconds
	}
	staleThreshold := time.Duration(statusStaleLeaseMultiplier*leaseDurationSec) * time.Second

	renewStr := strings.TrimSpace(ptr.Deref(lease.Status.GlobalRenewTime, ""))
	if renewStr == "" {
		return []string{fmt.Sprintf("cluster %q MultiClusterLease \"%s/%s\" is missing status.globalRenewTime", report.contextName, lease.Namespace, lease.Name)}
	}

	report.renewTime = renewStr
	parsedTime, err := time.Parse(time.RFC3339Nano, renewStr)
	if err != nil {
		report.renewAge = fmt.Sprintf("invalid (%s)", renewStr)
		return []string{fmt.Sprintf("cluster %q MultiClusterLease \"%s/%s\" has invalid status.globalRenewTime %q: %v", report.contextName, lease.Namespace, lease.Name, renewStr, err)}
	}

	age := c.now.Sub(parsedTime).Truncate(time.Second)
	if age < 0 {
		age = 0
	}
	report.renewAge = age.String()
	if age > staleThreshold {
		return []string{fmt.Sprintf(
			"stale MultiClusterLease status on cluster %q: leader %q last renewed %s ago (at %s, staleness threshold %s)",
			report.contextName, report.observedLeader, age, parsedTime.UTC().Format(time.RFC3339), staleThreshold,
		)}
	}

	return nil
}

func (c *clusterStatusInspector) inspectKRMSyncers(ctx context.Context, namespace string, report *clusterStatusReport) ([]syncerv1alpha1.KRMSyncer, []string) {
	if namespace == "" {
		return nil, nil
	}

	var warnings []string
	syncers, err := listKRMSyncers(ctx, c.handle, namespace)
	if err != nil {
		warnings = append(warnings, err.Error())
	}

	for _, s := range syncers {
		summary, syncerWarns := summarizeKRMSyncer(c.handle.contextName, s)
		report.syncers = append(report.syncers, summary)
		warnings = append(warnings, syncerWarns...)
	}

	return syncers, warnings
}

func summarizeKRMSyncer(contextName string, s syncerv1alpha1.KRMSyncer) (syncerSummary, []string) {
	mode := string(s.Spec.Mode)
	if mode == "" {
		mode = string(syncerv1alpha1.ModePull)
	}

	source := resolveSyncerSource(s)
	health, syncerErrors, warnings := evaluateSyncerConditions(contextName, s)

	return syncerSummary{
		namespace: s.Namespace,
		name:      s.Name,
		mode:      mode,
		suspended: s.Spec.Suspend,
		source:    source,
		health:    health,
		errors:    syncerErrors,
	}, warnings
}

func resolveSyncerSource(s syncerv1alpha1.KRMSyncer) string {
	if s.Spec.Remote == nil || s.Spec.Remote.ClusterConfig == nil || s.Spec.Remote.ClusterConfig.KubeConfigSecretRef == nil {
		return "<unconfigured>"
	}

	ref := s.Spec.Remote.ClusterConfig.KubeConfigSecretRef
	if ref.Namespace != "" && ref.Namespace != s.Namespace {
		return fmt.Sprintf("%s/%s", ref.Namespace, ref.Name)
	}
	if ref.Name != "" {
		return ref.Name
	}

	return "<unconfigured>"
}

func evaluateSyncerConditions(contextName string, s syncerv1alpha1.KRMSyncer) (string, []string, []string) {
	var condParts []string
	var syncerErrors []string
	var warnings []string

	for _, cond := range s.Status.Conditions {
		condParts = append(condParts, fmt.Sprintf("%s=%s", cond.Type, cond.Status))
		if isSyncerErrorCondition(cond, s.Spec.Suspend) {
			errMsg := fmt.Sprintf("%s=%s (%s): %s", cond.Type, cond.Status, cond.Reason, cond.Message)
			syncerErrors = append(syncerErrors, errMsg)
			warnings = append(warnings, fmt.Sprintf("cluster %q KRMSyncer \"%s/%s\" reported error condition %s", contextName, s.Namespace, s.Name, errMsg))
		}
	}

	health := "Unknown"
	if len(condParts) > 0 {
		health = strings.Join(condParts, ",")
	} else if s.Spec.Suspend {
		health = "Suspended"
	}

	return health, syncerErrors, warnings
}

func isSyncerErrorCondition(cond metav1.Condition, suspended bool) bool {
	reasonLower := strings.ToLower(cond.Reason)
	return (cond.Type == "InvalidRule" && cond.Status == metav1.ConditionTrue) ||
		(cond.Type == "Active" && cond.Status == metav1.ConditionFalse && !suspended) ||
		strings.Contains(reasonLower, "error") ||
		strings.Contains(reasonLower, "fail")
}

const (
	statusMCLDeploymentNamespace = "cnrm-system"
	statusMCLDeploymentName      = "multiclusterlease-controller-manager"
)

func (c *clusterStatusInspector) discoverGCSBucket(ctx context.Context) string {
	if c.handle == nil || c.handle.client == nil {
		return ""
	}

	deploy := &appsv1.Deployment{}
	key := types.NamespacedName{Namespace: statusMCLDeploymentNamespace, Name: statusMCLDeploymentName}
	if err := c.handle.client.Get(ctx, key, deploy); err != nil {
		return ""
	}

	for _, container := range deploy.Spec.Template.Spec.Containers {
		for _, token := range append(container.Command, container.Args...) {
			if strings.HasPrefix(token, "--gcs-bucket=") {
				bucket := normalizeBucketName(strings.TrimPrefix(token, "--gcs-bucket="))
				if bucket != "" {
					return bucket
				}
			}
		}
	}

	return ""
}
