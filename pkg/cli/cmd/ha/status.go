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

// clusterStatus is the observed HA status of one cluster.
type clusterStatus struct {
	gke gkeClusterInfo
	ha  clusterHA
	// state is the cluster's HA objects; set only when HA is configured.
	state *clusterHAState
	// err is the error from fetchClusterHAState; nil when HA is configured.
	err error
	// workloads holds the observed HA workloads on the cluster; set only when HA is configured.
	workloads *clusterWorkloads
	// workloadsErr is non-nil if reading workloads failed with a non-NotFound error.
	workloadsErr error
}

// statusClusterTimeout bounds how long a single cluster is inspected, so that one unreachable
// cluster cannot hang the command.
const statusClusterTimeout = 30 * time.Second

// inspectCluster fetches the cluster's HA state with fetchClusterHAState and maps the result
// to the HA column.
func inspectCluster(ctx context.Context, h *clusterHandle) *clusterStatus {
	// validate() checked that every context is a GKE context.
	gke, _ := parseGKEContext(h.contextName)
	c := &clusterStatus{gke: gke}

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
}

func buildStatusReport(clusters []*clusterStatus, now time.Time) *statusReport {
	r := &statusReport{clusters: clusters}

	var configured, notConfigured []*clusterStatus
	for _, c := range clusters {
		switch c.ha {
		case clusterHAConfigured:
			configured = append(configured, c)
		case clusterHANotConfigured:
			notConfigured = append(notConfigured, c)
		default:
			r.problems = append(r.problems, c.err.Error())
		}
	}

	// HA must be configured on all clusters or on none of them.
	if len(configured) > 0 {
		for _, c := range notConfigured {
			r.problems = append(r.problems, c.err.Error())
		}
	}

	// When every cluster is configured, check the election and data sync across them.
	if len(configured) == len(clusters) {
		states := make([]*clusterHAState, 0, len(configured))
		for _, c := range configured {
			states = append(states, c.state)
		}
		r.problems = append(r.problems, haProblems(states, now)...)

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

// render writes the report: a table with one row per cluster, then the overall status.
func (r *statusReport) render(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "CONTEXT\tPROJECT\tLOCATION\tHA")
	for _, c := range r.clusters {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", c.gke.contextName, c.gke.project, c.gke.location, c.ha)
	}
	tw.Flush()

	fmt.Fprintln(w)
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
