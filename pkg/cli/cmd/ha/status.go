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
	"strings"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// statusOptions holds flags and configuration for `config-connector ha status`.
type statusOptions struct {
	baseOptions

	bucket string
}

func newStatusOptions() statusOptions {
	return statusOptions{}
}

const statusLong = `Display multi-cluster high availability leadership, manager, and replication status.

Status inspects each cluster named in --contexts and checks:
  - the clusters share a consistent MultiClusterLease configuration;
  - each cluster's lease backend is healthy and the lease has a live holder;
  - all clusters agree on the current leader;
  - Config Connector's manager StatefulSets are running and ready;
  - KRMSyncers are suspended on the leader and actively syncing on standbys.

In namespaced mode, all managers on a cluster share the multi-cluster lease under the
cluster's identity, so leadership never moves for one namespace alone.`

func NewStatusCmd() *cobra.Command {
	opts := newStatusOptions()
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Display multi-cluster high availability leadership and replication status",
		Long:  statusLong,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.stdout = cmd.OutOrStdout()
			opts.stderr = cmd.ErrOrStderr()
			if err := opts.validate(); err != nil {
				return err
			}
			// From here on, errors are about the clusters, not about how the command was used.
			cmd.SilenceUsage = true
			return runStatus(cmd.Context(), opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

func (o *statusOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.bucket, bucketFlag, o.bucket, "Optional GCS bucket name to inspect the global lock object directly.")
}

func (o *statusOptions) validate() error {
	if err := o.baseOptions.validate(); err != nil {
		return err
	}
	// Like `ha setup`, require GKE contexts so that every cluster's project and location are known.
	var errs []error
	for _, contextName := range o.contexts {
		if _, err := parseGKEContext(contextName); err != nil {
			errs = append(errs, err)
		}
	}
	o.bucket = normalizeBucketName(o.bucket)
	return errors.Join(errs...)
}

// errHAStatusError is returned when the HA status is ERROR, so that the command exits non-zero.
var errHAStatusError = errors.New("multi-cluster HA status is ERROR; see the problems above")

func runStatus(ctx context.Context, opts statusOptions) error {
	handles, err := buildClusterHandles(ctx, opts.contexts)
	if err != nil {
		return err
	}

	report := executeStatus(ctx, handles, time.Now())

	stdout := opts.stdout
	if stdout == nil {
		stdout = io.Discard
	}
	report.render(stdout)

	if report.state == haStateError {
		return errHAStatusError
	}
	return nil
}

// executeStatus inspects every cluster and builds the HA status report. now is used to tell
// whether the lease has expired; it is a parameter so tests are deterministic.
func executeStatus(ctx context.Context, handles []*clusterHandle, now time.Time) *statusReport {
	clusters := make([]*clusterStatus, 0, len(handles))
	for _, h := range handles {
		clusters = append(clusters, inspectCluster(ctx, h))
	}
	return buildStatusReport(clusters, now)
}

// clusterHA is the value of the HA column for a cluster.
type clusterHA string

const (
	clusterHAConfigured        clusterHA = "Configured"
	clusterHANotConfigured     clusterHA = "Not configured"
	clusterHANoConfigConnector clusterHA = "No ConfigConnector"
	clusterHAError             clusterHA = "Error"
)

// clusterRole is the value of the ROLE column for a cluster.
type clusterRole string

const (
	// roleLeader: the cluster holds the multi-cluster lease, so its Config Connector reconciles.
	roleLeader clusterRole = "Leader"
	// roleStandby: another cluster holds the lease; this cluster waits to take over.
	roleStandby clusterRole = "Standby"
	// roleUnknown: HA is configured, but the lease holder could not be determined.
	roleUnknown clusterRole = "Unknown"
	// roleNone: HA is not configured, or the cluster could not be inspected.
	roleNone clusterRole = "-"
)

// clusterStatus is the observed HA status of one cluster.
type clusterStatus struct {
	gke  gkeClusterInfo
	ha   clusterHA
	role clusterRole
	// state is the cluster's HA objects; set only when HA is configured.
	state *clusterHAState
	// err is the error from fetchClusterHAState; nil when HA is configured.
	err error
	// workloads holds the observed HA workloads on the cluster; set only when HA is configured.
	workloads *clusterWorkloads
	// workloadsErr is non-nil if reading workloads failed with a non-NotFound error.
	workloadsErr error
}

// identity returns the cluster's identity in the election, or "" if HA is not configured.
func (c *clusterStatus) identity() string {
	if c.state == nil {
		return ""
	}
	identity, _ := multiClusterLeaseSpec(c.state.configConnector).CandidateIdentity()
	return identity
}

// statusClusterTimeout bounds how long a single cluster is inspected, so that one unreachable
// cluster cannot hang the command.
const statusClusterTimeout = 30 * time.Second

// inspectCluster fetches the cluster's HA state with fetchClusterHAState and maps the result
// to the HA column.
func inspectCluster(ctx context.Context, h *clusterHandle) *clusterStatus {
	// validate() checked that every context is a GKE context.
	gke, _ := parseGKEContext(h.contextName)
	c := &clusterStatus{gke: gke, role: roleNone}

	ctx, cancel := context.WithTimeout(ctx, statusClusterTimeout)
	defer cancel()

	c.state, c.err = fetchClusterHAState(ctx, h)

	var notConfigured *multiClusterLeaseNotConfiguredError
	var notFound *configConnectorNotFoundError
	switch {
	case c.err == nil:
		c.ha = clusterHAConfigured
		c.workloads, c.workloadsErr = fetchClusterWorkloads(ctx, h, c.state.configConnector)
	case errors.As(c.err, &notConfigured):
		c.ha = clusterHANotConfigured
	case errors.As(c.err, &notFound):
		c.ha = clusterHANoConfigConnector
	default:
		c.ha = clusterHAError
	}
	return c
}

type haState string

const (
	haStateNotDetected haState = "NOT DETECTED"
	// haStateHealthy means HA is configured on every cluster and no problem was found.
	haStateHealthy haState = "HEALTHY"
	haStateError   haState = "ERROR"
)

// statusReport is the overall multi-cluster HA status across the inspected clusters.
type statusReport struct {
	state    haState
	clusters []*clusterStatus
	problems []string
	warnings []string
	// lease describes the multi-cluster lease; nil when its holder is not known.
	lease *leaseSummary
	now   time.Time
}

// leaseSummary describes the multi-cluster lease and the cluster that holds it.
type leaseSummary struct {
	// key is <namespace>/<name> of the MultiClusterLease.
	key string
	// holder is the context of the leader, or its identity if it's not in --contexts.
	holder string
	// renewTime is when the leader last renewed the lease; zero if not reported.
	renewTime time.Time
}

func buildStatusReport(clusters []*clusterStatus, now time.Time) *statusReport {
	r := &statusReport{clusters: clusters, now: now}

	var configured, notConfigured []*clusterStatus
	for _, c := range clusters {
		switch c.ha {
		case clusterHAConfigured:
			configured = append(configured, c)
			c.role = roleUnknown
		case clusterHANotConfigured:
			notConfigured = append(notConfigured, c)
		default:
			r.problems = append(r.problems, c.err.Error())
		}
	}

	// HA must be configured on all clusters or on none of them. A cluster without HA runs its own
	// leader election, so it reconciles alongside whichever cluster holds the multi-cluster lease.
	if len(configured) > 0 {
		for _, c := range notConfigured {
			r.problems = append(r.problems, c.err.Error()+"; it reconciles independently of the HA clusters, so more than one cluster may change the same resources")
		}
	}

	// Check the election, leadership and data sync across the configured clusters, even if some
	// clusters aren't configured or couldn't be read.
	if len(configured) > 0 {
		states := make([]*clusterHAState, 0, len(configured))
		for _, c := range configured {
			if c.workloads != nil && c.state != nil {
				c.state.isHASetupManaged = c.workloads.isHASetupManaged
			}
			states = append(states, c.state)
		}
		leader, problems := haProblems(states, r.contextIdentities(), now)
		r.problems = append(r.problems, problems...)
		if leader != "" {
			r.setLeader(leader, configured)
		}

		for _, c := range configured {
			if c.workloadsErr != nil {
				r.problems = append(r.problems, fmt.Sprintf("cluster %q: could not read Config Connector's workloads: %v", c.gke.contextName, c.workloadsErr))
				continue
			}
			if c.workloads != nil {
				probs, warns := c.workloads.checkWorkloads(c.gke.contextName, c.state, now)
				r.problems = append(r.problems, probs...)
				r.warnings = append(r.warnings, warns...)
			}
		}
	}

	switch {
	case len(r.problems) > 0:
		r.state = haStateError
	case len(configured) > 0:
		r.state = haStateHealthy
	default:
		r.state = haStateNotDetected
	}
	return r
}

// clusterIdentity returns a cluster's identity in the election: from its multiClusterLease spec
// when HA is configured, or else from its GKE context (gke_<project>_<location>_<cluster>), which
// is how `config-connector ha setup` sets it.
func clusterIdentity(c *clusterStatus) string {
	if id := c.identity(); id != "" {
		return id
	}
	gke := c.gke.gkeCluster()
	return gke.ResourceName()
}

// contextIdentities returns the election identity of every cluster in --contexts.
func (r *statusReport) contextIdentities() map[string]bool {
	ids := make(map[string]bool, len(r.clusters))
	for _, c := range r.clusters {
		ids[clusterIdentity(c)] = true
	}
	return ids
}

// setLeader sets each configured cluster's role, and the lease summary, from the identity of the
// cluster that holds the lease. The leader may be a cluster that isn't configured or couldn't be
// read, or one outside --contexts.
func (r *statusReport) setLeader(leader string, configured []*clusterStatus) {
	for _, c := range configured {
		if c.identity() == leader {
			c.role = roleLeader
		} else {
			c.role = roleStandby
		}
	}

	// Name the holder by its context when it's in --contexts, and by its identity otherwise.
	holder := leader
	for _, c := range r.clusters {
		if clusterIdentity(c) == leader {
			holder = c.gke.contextName
			break
		}
	}

	// haProblems checked that the configured clusters use the same lease.
	mcl := multiClusterLeaseSpec(configured[0].state.configConnector)
	r.lease = &leaseSummary{
		key:    mcl.Namespace + "/" + mcl.LeaseName,
		holder: holder,
	}
	// Use the latest renewal reported for the leader by any configured cluster.
	for _, c := range configured {
		lease := c.state.multiClusterLease
		if lease == nil || lease.Status.GlobalHolderIdentity == nil || *lease.Status.GlobalHolderIdentity != leader || lease.Status.GlobalRenewTime == nil {
			continue
		}
		if t, err := time.Parse(time.RFC3339, *lease.Status.GlobalRenewTime); err == nil && t.After(r.lease.renewTime) {
			r.lease.renewTime = t
		}
	}
}

// render writes the report: a table with one row per cluster, then the overall status.
func (r *statusReport) render(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "CONTEXT\tPROJECT\tLOCATION\tHA\tROLE")
	for _, c := range r.clusters {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", c.gke.contextName, c.gke.project, c.gke.location, c.ha, c.role)
	}
	tw.Flush()

	fmt.Fprintln(w)
	if r.lease != nil {
		fmt.Fprintf(w, "Lease: %s, held by %s", r.lease.key, r.lease.holder)
		if !r.lease.renewTime.IsZero() {
			fmt.Fprintf(w, " (renewed %s ago)", r.now.Sub(r.lease.renewTime).Round(time.Second))
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintf(w, "HA Status: %s\n", r.state)
	switch r.state {
	case haStateNotDetected:
		contexts := make([]string, 0, len(r.clusters))
		for _, c := range r.clusters {
			contexts = append(contexts, c.gke.contextName)
		}
		fmt.Fprintf(w, "Next step: config-connector ha setup --%s=%s --%s=<BUCKET>\n", contextsFlag, strings.Join(contexts, ","), bucketFlag)
	case haStateError:
		fmt.Fprintln(w, "Problems:")
		for _, p := range r.problems {
			fmt.Fprintf(w, "  - %s\n", p)
		}
	}
	if len(r.warnings) > 0 {
		fmt.Fprintln(w, "Warnings:")
		for _, warn := range r.warnings {
			fmt.Fprintf(w, "  - %s\n", warn)
		}
	}
}
