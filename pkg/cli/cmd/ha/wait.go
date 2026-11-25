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
	"sort"
	"strings"
	"time"

	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
)

func waitForHAHealthy(ctx context.Context, opts setupOptions, configurers []*clusterConfigurer, handles []*clusterHandle) error {
	stdout := opts.stdout
	if stdout == nil {
		stdout = io.Discard
	}

	if !opts.wait {
		if !opts.dryRun {
			fmt.Fprintf(stdout, "config-connector ha status --%s=%s\n", contextsFlag, strings.Join(opts.contexts, ","))
		}
		return nil
	}

	if opts.dryRun {
		fmt.Fprintln(stdout, "[dry-run] would wait for multi-cluster HA to become healthy")
		return nil
	}

	fmt.Fprintf(stdout, "Waiting up to %s for multi-cluster HA to become healthy (use --wait=false to skip)...\n", opts.waitTimeout)

	statusFn := opts.statusFunc
	if statusFn == nil {
		statusFn = executeStatus
	}
	nowFn := opts.nowFunc
	if nowFn == nil {
		nowFn = time.Now
	}

	pollInterval := opts.waitPollInterval
	if pollInterval <= 0 {
		pollInterval = defaultWaitPollInterval
	}

	sleepFn := opts.sleepFunc
	if sleepFn == nil {
		sleepFn = func(ctx context.Context, d time.Duration) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(d):
				return nil
			}
		}
	}

	startTime := nowFn()
	deadline := startTime.Add(opts.waitTimeout)

	var lastReport *statusReport
	for {
		now := nowFn()
		remaining := deadline.Sub(now)
		if remaining <= 0 {
			break
		}

		pollCtx, cancel := context.WithTimeout(ctx, remaining)
		rep := statusFn(pollCtx, handles, now)
		cancel()

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if rep != nil && rep.state == haStateHealthy {
			printHASummary(stdout, rep, nowFn())
			return nil
		}

		if pollCtx.Err() == nil || lastReport == nil {
			lastReport = rep
		}

		remaining = deadline.Sub(nowFn())
		if remaining <= 0 {
			break
		}

		sleepDuration := pollInterval
		if sleepDuration > remaining {
			sleepDuration = remaining
		}

		if err := sleepFn(ctx, sleepDuration); err != nil {
			return err
		}
	}

	var problems []string
	if lastReport != nil {
		problems = lastReport.problems
	}
	hasNamespaced := isAnyClusterNamespaced(configurers)
	return waitTimeoutError(problems, opts.contexts, hasNamespaced)
}

func printHASummary(w io.Writer, report *statusReport, now time.Time) {
	var states []*clusterHAState
	for _, c := range report.clusters {
		if c.state != nil {
			states = append(states, c.state)
		}
	}
	leaderID, _ := leaseHolder(states, now)

	var leaderContext string
	var standbyContexts []string
	var repMode string

	for _, c := range report.clusters {
		if c.state == nil {
			continue
		}
		mcl := multiClusterLeaseSpec(c.state.configConnector)
		if mcl == nil {
			continue
		}
		if repMode == "" {
			repMode = replicationMode(mcl)
		}
		id, _ := mcl.CandidateIdentity()
		if id == leaderID {
			leaderContext = c.gke.contextName
		} else {
			standbyContexts = append(standbyContexts, c.gke.contextName)
		}
	}
	sort.Strings(standbyContexts)

	fmt.Fprintln(w, "Multi-cluster HA is healthy.")
	fmt.Fprintf(w, "  Leader:      %s\n", leaderContext)
	fmt.Fprintf(w, "  Standbys:    %s\n", strings.Join(standbyContexts, ", "))
	fmt.Fprintf(w, "  Replication: %s\n", repMode)
}

func waitTimeoutError(problems []string, contexts []string, hasNamespaced bool) error {
	var lines []string
	lines = append(lines, "timed out waiting for multi-cluster HA to become healthy")
	if len(problems) > 0 {
		lines = append(lines, "Problems:")
		for _, p := range problems {
			lines = append(lines, "  - "+p)
		}
	}
	lines = append(lines, "The HA configuration is applied and was not rolled back.")
	lines = append(lines, fmt.Sprintf("Check status again with: config-connector ha status --%s=%s", contextsFlag, strings.Join(contexts, ",")))
	if hasNamespaced {
		lines = append(lines, "In namespaced mode, the manager of each namespace restarts on its own, so a cluster with many namespaces can need a longer --wait-timeout.")
	}
	return errors.New(strings.Join(lines, "\n"))
}

func isAnyClusterNamespaced(configurers []*clusterConfigurer) bool {
	for _, cfg := range configurers {
		if cfg.configConnector != nil && cfg.configConnector.GetMode() == operatork8s.NamespacedMode {
			return true
		}
	}
	return false
}
