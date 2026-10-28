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
	"io"
	"sort"
	"strings"
	"time"

	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	defaultLeaseDurationSeconds = int32(15)
	staleLeaseMultiplier        = 2
	kubeconfigSecretKey         = "kubeconfig"
)

// statusOptions holds flags and configuration for `config-connector ha status`.
type statusOptions struct {
	baseOptions

	bucket string
}

func newStatusOptions() statusOptions {
	return statusOptions{}
}

func NewStatusCmd() *cobra.Command {
	opts := newStatusOptions()
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Display multi-cluster high availability leadership and replication status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.stdout = cmd.OutOrStdout()
			opts.stderr = cmd.ErrOrStderr()
			if err := opts.validate(); err != nil {
				return err
			}
			return runStatus(cmd.Context(), opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

func runStatus(ctx context.Context, opts statusOptions) error {
	handles, err := buildClusterHandles(ctx, opts.contexts)
	if err != nil {
		return err
	}
	runner := newExecCommandRunner(opts.stdout, opts.stderr, false)
	return executeStatus(ctx, opts, handles, runner, time.Now().UTC())
}

// gcsLockData matches the JSON payload written to GCS by multicluster-leader-election (`storage.LeaseData`).
type gcsLockData struct {
	HolderIdentity   string    `json:"holderIdentity"`
	RenewTime        time.Time `json:"renewTime"`
	LeaseTransitions int32     `json:"leaseTransitions"`
}

type gcsLockStatus struct {
	bucket           string
	expectedLockKey  string
	lockURI          string
	found            bool
	holderIdentity   string
	renewTime        time.Time
	renewAge         time.Duration
	leaseTransitions int32
	extraKeys        []string
}

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

func executeStatus(ctx context.Context, opts statusOptions, handles []*clusterHandle, outRunner commandOutputRunner, now time.Time) error {
	var warnings []string

	reports, validStates, clusterWarnings := collectAllClusterStatusReports(ctx, handles, now)
	warnings = append(warnings, clusterWarnings...)

	if len(validStates) >= 2 {
		if err := validateMultiClusterLeaseConfigs(validStates); err != nil {
			for _, line := range strings.Split(err.Error(), "\n") {
				if trimmed := strings.TrimSpace(line); trimmed != "" {
					warnings = append(warnings, trimmed)
				}
			}
		}
	}

	warnings = append(warnings, checkMissingPeerSecrets(ctx, handles, reports)...)

	bucket := opts.bucket
	if bucket == "" {
		bucket = discoverGCSBucketFromClusters(ctx, handles)
	}

	leaseNS, leaseName := resolveExpectedLeaseCoordinates(reports)
	var gcsStatus *gcsLockStatus
	if bucket != "" && outRunner != nil {
		var gcsWarnings []string
		gcsStatus, gcsWarnings = inspectGCSLock(ctx, bucket, leaseNS, leaseName, outRunner, now)
		warnings = append(warnings, gcsWarnings...)
	}

	warnings = append(warnings, checkLeaderConsistency(reports, gcsStatus)...)

	if opts.stdout != nil {
		renderStatusReport(opts.stdout, gcsStatus, reports, warnings)
	}
	return nil
}

func collectAllClusterStatusReports(ctx context.Context, handles []*clusterHandle, now time.Time) ([]clusterStatusReport, []*clusterHAState, []string) {
	reports := make([]clusterStatusReport, 0, len(handles))
	validStates := make([]*clusterHAState, 0, len(handles))
	var warnings []string

	for _, h := range handles {
		report, state, clusterWarns := collectSingleClusterStatus(ctx, h, now)
		reports = append(reports, report)
		if state != nil {
			validStates = append(validStates, state)
		}
		warnings = append(warnings, clusterWarns...)
	}
	return reports, validStates, warnings
}

func collectSingleClusterStatus(ctx context.Context, handle *clusterHandle, now time.Time) (clusterStatusReport, *clusterHAState, []string) {
	report := clusterStatusReport{
		contextName:       "<unknown>",
		candidateIdentity: "<unconfigured>",
		role:              "UNCONFIGURED",
		observedLeader:    "<none>",
		renewAge:          "<unknown>",
	}
	if handle == nil || handle.client == nil {
		return report, nil, []string{"cluster handle or client is nil"}
	}
	report.contextName = handle.contextName

	var warnings []string
	cc, err := fetchConfigConnector(ctx, handle)
	if err != nil {
		warnings = append(warnings, err.Error())
		return report, nil, warnings
	}
	if cc == nil {
		warnings = append(warnings, fmt.Sprintf("cluster %q is missing the %q ConfigConnector object", handle.contextName, configConnectorResourceName))
		return report, nil, warnings
	}

	mclSpec := multiClusterLeaseSpec(cc)
	if err := validateClusterLeaseSpec(handle.contextName, mclSpec); err != nil {
		for _, line := range strings.Split(err.Error(), "\n") {
			if trimmed := strings.TrimSpace(line); trimmed != "" {
				warnings = append(warnings, trimmed)
			}
		}
		if mclSpec == nil {
			return report, nil, warnings
		}
	}

	report.candidateIdentity = mclSpec.ClusterCandidateIdentity
	report.leaseNamespace = mclSpec.Namespace
	report.leaseName = mclSpec.LeaseName
	report.replicationMode = mclSpec.ResourceReplicationMode
	if report.replicationMode == "" {
		report.replicationMode = "Status"
	}
	report.role = "UNDECIDED"

	var lease *mclv1alpha1.MultiClusterLease
	if mclSpec.Namespace != "" && mclSpec.LeaseName != "" {
		lease, err = fetchMultiClusterLease(ctx, handle, mclSpec.Namespace, mclSpec.LeaseName)
		if err != nil {
			warnings = append(warnings, err.Error())
		}
	}

	if lease == nil {
		if mclSpec.Namespace != "" && mclSpec.LeaseName != "" {
			warnings = append(warnings, fmt.Sprintf("cluster %q is missing MultiClusterLease \"%s/%s\"", handle.contextName, mclSpec.Namespace, mclSpec.LeaseName))
		}
	} else {
		leaseWarnings := populateLeaseStatus(&report, lease, now)
		warnings = append(warnings, leaseWarnings...)
	}

	var syncers []syncerv1alpha1.KRMSyncer
	if mclSpec.Namespace != "" {
		syncers, err = listKRMSyncers(ctx, handle, mclSpec.Namespace)
		if err != nil {
			warnings = append(warnings, err.Error())
		}
	}
	for _, s := range syncers {
		summary, syncerWarns := summarizeKRMSyncer(handle.contextName, s)
		report.syncers = append(report.syncers, summary)
		warnings = append(warnings, syncerWarns...)
	}

	var validState *clusterHAState
	if validateClusterLeaseSpec(handle.contextName, mclSpec) == nil {
		validState = &clusterHAState{
			contextName:       handle.contextName,
			configConnector:   cc,
			multiClusterLease: lease,
			krmSyncers:        syncers,
		}
		report.rawState = validState
	}

	return report, validState, warnings
}

func populateLeaseStatus(report *clusterStatusReport, lease *mclv1alpha1.MultiClusterLease, now time.Time) []string {
	var warnings []string
	holder := derefString(lease.Status.GlobalHolderIdentity)
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

	report.transitions = derefInt32(lease.Status.GlobalLeaseTransitions)
	leaseDurationSec := derefInt32(lease.Status.GlobalLeaseDurationSeconds)
	if leaseDurationSec <= 0 {
		leaseDurationSec = defaultLeaseDurationSeconds
	}
	staleThreshold := time.Duration(staleLeaseMultiplier*leaseDurationSec) * time.Second

	renewStr := strings.TrimSpace(derefString(lease.Status.GlobalRenewTime))
	if renewStr == "" {
		warnings = append(warnings, fmt.Sprintf("cluster %q MultiClusterLease \"%s/%s\" is missing status.globalRenewTime", report.contextName, lease.Namespace, lease.Name))
		return warnings
	}

	report.renewTime = renewStr
	parsedTime, err := parseTimestamp(renewStr)
	if err != nil {
		report.renewAge = fmt.Sprintf("invalid (%s)", renewStr)
		warnings = append(warnings, fmt.Sprintf("cluster %q MultiClusterLease \"%s/%s\" has invalid status.globalRenewTime %q: %v", report.contextName, lease.Namespace, lease.Name, renewStr, err))
		return warnings
	}

	age := now.Sub(parsedTime).Truncate(time.Second)
	if age < 0 {
		age = 0
	}
	report.renewAge = age.String()
	if age > staleThreshold {
		warnings = append(warnings, fmt.Sprintf(
			"stale MultiClusterLease status on cluster %q: leader %q last renewed %s ago (at %s, staleness threshold %s)",
			report.contextName, report.observedLeader, age, parsedTime.UTC().Format(time.RFC3339), staleThreshold,
		))
	}
	return warnings
}

func summarizeKRMSyncer(contextName string, s syncerv1alpha1.KRMSyncer) (syncerSummary, []string) {
	mode := string(s.Spec.Mode)
	if mode == "" {
		mode = string(syncerv1alpha1.ModePull)
	}
	source := "<unconfigured>"
	if s.Spec.Remote != nil && s.Spec.Remote.ClusterConfig != nil && s.Spec.Remote.ClusterConfig.KubeConfigSecretRef != nil {
		ref := s.Spec.Remote.ClusterConfig.KubeConfigSecretRef
		if ref.Namespace != "" && ref.Namespace != s.Namespace {
			source = fmt.Sprintf("%s/%s", ref.Namespace, ref.Name)
		} else if ref.Name != "" {
			source = ref.Name
		}
	}

	var condParts []string
	var syncerErrors []string
	var warnings []string

	for _, c := range s.Status.Conditions {
		condParts = append(condParts, fmt.Sprintf("%s=%s", c.Type, c.Status))
		isErrorCond := (c.Type == "InvalidRule" && c.Status == metav1.ConditionTrue) ||
			(c.Type == "Active" && c.Status == metav1.ConditionFalse && !s.Spec.Suspend) ||
			strings.Contains(strings.ToLower(c.Reason), "error") ||
			strings.Contains(strings.ToLower(c.Reason), "fail")
		if isErrorCond {
			errMsg := fmt.Sprintf("%s=%s (%s): %s", c.Type, c.Status, c.Reason, c.Message)
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

func checkMissingPeerSecrets(ctx context.Context, handles []*clusterHandle, reports []clusterStatusReport) []string {
	var warnings []string
	handleByContext := make(map[string]*clusterHandle, len(handles))
	for _, h := range handles {
		if h != nil {
			handleByContext[h.contextName] = h
		}
	}

	for i, r := range reports {
		h := handleByContext[r.contextName]
		if h == nil || h.client == nil || r.leaseNamespace == "" {
			continue
		}

		expectedSecrets := make(map[types.NamespacedName]string)
		for j, peer := range reports {
			if i == j || peer.candidateIdentity == "" || strings.HasPrefix(peer.candidateIdentity, "<") {
				continue
			}
			expectedSecrets[types.NamespacedName{Namespace: r.leaseNamespace, Name: peer.candidateIdentity}] = fmt.Sprintf("peer cluster %q", peer.contextName)
		}

		if r.rawState != nil {
			for _, syncer := range r.rawState.krmSyncers {
				if syncer.Spec.Remote != nil && syncer.Spec.Remote.ClusterConfig != nil && syncer.Spec.Remote.ClusterConfig.KubeConfigSecretRef != nil {
					ref := syncer.Spec.Remote.ClusterConfig.KubeConfigSecretRef
					if ref.Name != "" {
						ns := ref.Namespace
						if ns == "" {
							ns = r.leaseNamespace
						}
						expectedSecrets[types.NamespacedName{Namespace: ns, Name: ref.Name}] = fmt.Sprintf("KRMSyncer %q", syncer.Name)
					}
				}
			}
		}

		keys := make([]types.NamespacedName, 0, len(expectedSecrets))
		for k := range expectedSecrets {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(a, b int) bool {
			return keys[a].String() < keys[b].String()
		})

		for _, key := range keys {
			secret := &corev1.Secret{}
			if err := h.client.Get(ctx, key, secret); err != nil {
				if isResourceOrKindAbsent(err) {
					warnings = append(warnings, fmt.Sprintf(
						"cluster %q is missing peer kubeconfig Secret %q (required for %s)",
						r.contextName, key.String(), expectedSecrets[key],
					))
				} else {
					warnings = append(warnings, fmt.Sprintf("fetching Secret %q on cluster %q: %v", key.String(), r.contextName, err))
				}
				continue
			}
			if len(secret.Data[kubeconfigSecretKey]) == 0 && len(secret.StringData[kubeconfigSecretKey]) == 0 {
				warnings = append(warnings, fmt.Sprintf(
					"cluster %q peer Secret %q is missing required %q key",
					r.contextName, key.String(), kubeconfigSecretKey,
				))
			}
		}
	}

	return warnings
}

func discoverGCSBucketFromClusters(ctx context.Context, handles []*clusterHandle) string {
	for _, h := range handles {
		if h == nil || h.client == nil {
			continue
		}
		deploy := &appsv1.Deployment{}
		key := types.NamespacedName{Namespace: cnrmSystemNamespace, Name: mclDeploymentName}
		if err := h.client.Get(ctx, key, deploy); err != nil {
			continue
		}
		for _, c := range deploy.Spec.Template.Spec.Containers {
			for _, token := range append(c.Command, c.Args...) {
				if strings.HasPrefix(token, "--gcs-bucket=") {
					bucket := normalizeBucketName(strings.TrimPrefix(token, "--gcs-bucket="))
					if bucket != "" {
						return bucket
					}
				}
			}
		}
	}
	return ""
}

func resolveExpectedLeaseCoordinates(reports []clusterStatusReport) (string, string) {
	for _, r := range reports {
		if r.leaseNamespace != "" && r.leaseName != "" {
			return r.leaseNamespace, r.leaseName
		}
	}
	return defaultLeaseNamespace, defaultLeaseName
}

func inspectGCSLock(ctx context.Context, bucket, leaseNamespace, leaseName string, outRunner commandOutputRunner, now time.Time) (*gcsLockStatus, []string) {
	expectedKey := fmt.Sprintf("leases/%s/%s", leaseNamespace, leaseName)
	lockURI := fmt.Sprintf("gs://%s/%s", bucket, expectedKey)
	status := &gcsLockStatus{
		bucket:          bucket,
		expectedLockKey: expectedKey,
		lockURI:         lockURI,
	}
	var warnings []string

	lsOut, lsErr := outRunner.output(ctx, "gcloud", "storage", "ls", fmt.Sprintf("gs://%s/leases/**", bucket))
	if lsErr != nil {
		if !isGCSObjectNotFoundError(lsErr) {
			warnings = append(warnings, fmt.Sprintf("listing lease objects in gs://%s/leases/**: %v", bucket, lsErr))
		}
	} else {
		prefix := fmt.Sprintf("gs://%s/", bucket)
		for _, line := range strings.Split(lsOut, "\n") {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasSuffix(trimmed, "/") || strings.HasSuffix(trimmed, ":") {
				continue
			}
			key := strings.TrimPrefix(trimmed, prefix)
			if key != expectedKey {
				status.extraKeys = append(status.extraKeys, key)
			}
		}
		sort.Strings(status.extraKeys)
		if len(status.extraKeys) > 0 {
			warnings = append(warnings, fmt.Sprintf(
				"extra lease keys found in gs://%s: [%s] (expected only %q; clusters using different lease keys will elect multiple active leaders)",
				bucket, strings.Join(status.extraKeys, ", "), expectedKey,
			))
		}
	}

	catOut, catErr := outRunner.output(ctx, "gcloud", "storage", "cat", lockURI)
	if catErr != nil {
		if isGCSObjectNotFoundError(catErr) {
			warnings = append(warnings, fmt.Sprintf("GCS lock object %q does not exist yet", lockURI))
		} else {
			warnings = append(warnings, fmt.Sprintf("reading GCS lock object %q: %v", lockURI, catErr))
		}
		return status, warnings
	}

	var data gcsLockData
	if err := json.Unmarshal([]byte(strings.TrimSpace(catOut)), &data); err != nil {
		warnings = append(warnings, fmt.Sprintf("parsing GCS lock object %q JSON: %v", lockURI, err))
		return status, warnings
	}

	status.found = true
	status.holderIdentity = data.HolderIdentity
	status.renewTime = data.RenewTime
	status.leaseTransitions = data.LeaseTransitions
	if !data.RenewTime.IsZero() {
		age := now.Sub(data.RenewTime).Truncate(time.Second)
		if age < 0 {
			age = 0
		}
		status.renewAge = age
		staleThreshold := time.Duration(staleLeaseMultiplier*defaultLeaseDurationSeconds) * time.Second
		if age > staleThreshold {
			warnings = append(warnings, fmt.Sprintf(
				"stale GCS lock %q: holder %q last renewed %s ago (at %s, staleness threshold %s)",
				lockURI, data.HolderIdentity, age, data.RenewTime.UTC().Format(time.RFC3339), staleThreshold,
			))
		}
	}

	return status, warnings
}

func isGCSObjectNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "404") ||
		strings.Contains(msg, "not found") ||
		strings.Contains(msg, "matched no objects") ||
		strings.Contains(msg, "no urls matched")
}

func checkLeaderConsistency(reports []clusterStatusReport, gcsStatus *gcsLockStatus) []string {
	var warnings []string
	leaders := make(map[string][]string)
	activeLeaders := make([]string, 0)

	for _, r := range reports {
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
		parts := make([]string, 0, len(reports))
		for _, r := range reports {
			parts = append(parts, fmt.Sprintf("%s=%q", r.contextName, r.observedLeader))
		}
		warnings = append(warnings, fmt.Sprintf("clusters disagree on global leader (%s)", strings.Join(parts, ", ")))
	}
	if len(activeLeaders) > 1 {
		warnings = append(warnings, fmt.Sprintf("multiple clusters report role LEADER simultaneously: [%s]", strings.Join(activeLeaders, ", ")))
	}
	return warnings
}

func renderStatusReport(w io.Writer, gcsStatus *gcsLockStatus, reports []clusterStatusReport, warnings []string) {
	if gcsStatus != nil {
		fmt.Fprintln(w, "=== GCS Lock Status ===")
		fmt.Fprintf(w, "Bucket:      gs://%s\n", gcsStatus.bucket)
		fmt.Fprintf(w, "Lock Object: %s\n", gcsStatus.lockURI)
		if gcsStatus.found {
			fmt.Fprintf(w, "Leader:      %s\n", gcsStatus.holderIdentity)
			fmt.Fprintf(w, "Renew Age:   %s (renewed at %s, transitions: %d)\n",
				gcsStatus.renewAge, gcsStatus.renewTime.UTC().Format(time.RFC3339), gcsStatus.leaseTransitions)
		} else {
			fmt.Fprintln(w, "Leader:      <none> (lock object not found)")
		}
		if len(gcsStatus.extraKeys) > 0 {
			fmt.Fprintf(w, "Extra Keys:  %s\n", strings.Join(gcsStatus.extraKeys, ", "))
		}
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "=== Cluster HA Status ===")
	for _, r := range reports {
		fmt.Fprintf(w, "Context:            %s\n", r.contextName)
		fmt.Fprintf(w, "  Candidate ID:     %s\n", r.candidateIdentity)
		if r.leaseNamespace != "" || r.leaseName != "" {
			fmt.Fprintf(w, "  Lease:            %s/%s (replication: %s)\n", r.leaseNamespace, r.leaseName, r.replicationMode)
		} else {
			fmt.Fprintln(w, "  Lease:            <unconfigured>")
		}
		fmt.Fprintf(w, "  Role:             %s\n", r.role)
		fmt.Fprintf(w, "  Observed Leader:  %s\n", r.observedLeader)
		if r.renewTime != "" {
			fmt.Fprintf(w, "  Renew Age:        %s (renewed at %s, transitions: %d)\n", r.renewAge, r.renewTime, r.transitions)
		} else {
			fmt.Fprintf(w, "  Renew Age:        %s\n", r.renewAge)
		}
		if len(r.syncers) == 0 {
			fmt.Fprintln(w, "  Syncer:           <none>")
		} else {
			for _, s := range r.syncers {
				errSummary := "none"
				if len(s.errors) > 0 {
					errSummary = strings.Join(s.errors, "; ")
				}
				fmt.Fprintf(w, "  Syncer:           %s/%s (mode: %s, suspended: %t, source: %s, health: %s, errors: %s)\n",
					s.namespace, s.name, s.mode, s.suspended, s.source, s.health, errSummary)
			}
		}
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "=== Warnings ===")
	if len(warnings) == 0 {
		fmt.Fprintln(w, "None")
		return
	}
	for _, warn := range warnings {
		fmt.Fprintf(w, "- %s\n", warn)
	}
}

func parseTimestamp(raw string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339Nano, raw); err == nil {
		return t, nil
	}
	return time.Parse(time.RFC3339, raw)
}

func derefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func derefInt32(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}

func (o *statusOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.bucket, bucketFlag, o.bucket, "Optional GCS bucket name to inspect the global lock object directly.")
}

func (o *statusOptions) validate() error {
	if err := o.baseOptions.validate(); err != nil {
		return err
	}
	o.bucket = normalizeBucketName(o.bucket)
	return nil
}
