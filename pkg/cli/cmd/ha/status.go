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
	"io"
	"strings"
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
			return runStatus(cmd.Context(), opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

const statusBucketFlag = "bucket"

func (o *statusOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.bucket, statusBucketFlag, o.bucket, "Optional GCS bucket name to inspect the global lock object directly.")
}

func (o *statusOptions) validate() error {
	if err := o.baseOptions.validate(); err != nil {
		return err
	}

	o.bucket = normalizeBucketName(o.bucket)
	if o.bucket != "" {
		if err := validateBucketName(o.bucket); err != nil {
			return err
		}
	}

	return nil
}

func runStatus(ctx context.Context, opts statusOptions) error {
	handles, err := buildClusterHandles(ctx, opts.contexts)
	if err != nil {
		return err
	}

	runner := newExecCommandRunner(opts.stdout, opts.stderr, false)
	return executeStatus(ctx, opts, handles, runner, time.Now().UTC())
}

func executeStatus(ctx context.Context, opts statusOptions, handles []*clusterHandle, outRunner commandOutputRunner, now time.Time) error {
	reports, validStates, discoveredBucket, warnings := inspectAllClusters(ctx, handles, now)

	validator := newTopologyValidator(reports, validStates)
	warnings = append(warnings, validator.validateClusterTopology()...)

	bucket := opts.bucket
	if bucket == "" {
		bucket = discoveredBucket
	}

	var gcsStatus *gcsLockStatus
	if bucket != "" && outRunner != nil {
		leaseNS, leaseName := validator.resolveExpectedLeaseCoordinates()
		lockInspector := newGCSLockInspector(outRunner, bucket, leaseNS, leaseName, now)
		var gcsWarnings []string
		gcsStatus, gcsWarnings = lockInspector.inspect(ctx)
		warnings = append(warnings, gcsWarnings...)
	}

	warnings = append(warnings, validator.checkLeaderConsistency(gcsStatus)...)

	if opts.stdout != nil {
		renderStatusReport(opts.stdout, gcsStatus, reports, warnings)
	}

	return nil
}

func inspectAllClusters(ctx context.Context, handles []*clusterHandle, now time.Time) ([]clusterStatusReport, []*clusterHAState, string, []string) {
	reports := make([]clusterStatusReport, 0, len(handles))
	validStates := make([]*clusterHAState, 0, len(handles))
	var discoveredBucket string
	var warnings []string

	for _, h := range handles {
		inspector := newClusterStatusInspector(h, now)
		report, state, clusterWarns := inspector.inspect(ctx)
		reports = append(reports, report)
		if state != nil {
			validStates = append(validStates, state)
		}
		warnings = append(warnings, clusterWarns...)
		if discoveredBucket == "" {
			discoveredBucket = inspector.discoverGCSBucket(ctx)
		}
	}

	return reports, validStates, discoveredBucket, warnings
}

func renderStatusReport(w io.Writer, gcsStatus *gcsLockStatus, reports []clusterStatusReport, warnings []string) {
	renderGCSLockSection(w, gcsStatus)
	renderClusterSection(w, reports)
	renderWarningsSection(w, warnings)
}

func renderGCSLockSection(w io.Writer, gcsStatus *gcsLockStatus) {
	if gcsStatus == nil {
		return
	}

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

func renderClusterSection(w io.Writer, reports []clusterStatusReport) {
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
			continue
		}
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

func renderWarningsSection(w io.Writer, warnings []string) {
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
