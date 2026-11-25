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
	"os"
	"os/signal"
	"slices"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/wait"
)

// failoverOptions holds flags and configuration for `config-connector ha failover`.
type failoverOptions struct {
	baseOptions

	targetCluster string
	timeout       time.Duration
	dryRun        bool
	// skipSyncCheck lets failover go ahead when the KRMSyncers of a standby don't report that
	// they sync from the leader.
	skipSyncCheck bool

	// pollInterval is how often failover looks at the clusters while it waits.
	pollInterval time.Duration
	// leaderAgreementTimeout is how long the clusters get to agree on the current leader before
	// failover gives up. Each cluster learns the leader from its own MCL controller, which
	// re-reads the lease every 15 seconds, so the clusters can disagree right after a leader
	// change.
	leaderAgreementTimeout time.Duration
	// clusterOpTimeout bounds how long each cluster gets for non-polled reads and for stopping or
	// starting its managers, so one unresponsive cluster cannot block recovery on the others.
	clusterOpTimeout time.Duration
	// signalBounceWindow is how long duplicate interrupt signals are ignored right after the first
	// one before restoring immediate OS termination on a subsequent Ctrl-C.
	signalBounceWindow time.Duration
}

const (
	defaultFailoverTimeout                = 5 * time.Minute
	minFailoverTimeout                    = time.Minute
	defaultFailoverPollInterval           = 2 * time.Second
	defaultFailoverLeaderAgreementTimeout = 30 * time.Second

	// defaultFailoverClusterOpTimeout bounds per-cluster non-polled reads and manager stop/start
	// calls, including starting stopped managers after the command fails or is interrupted.
	defaultFailoverClusterOpTimeout = time.Minute

	// defaultSignalBounceWindow absorbs accidental double Ctrl-C taps while cleanup starts, while
	// still letting a later Ctrl-C terminate at once.
	defaultSignalBounceWindow = time.Second

	// maxLeaseStatusLag is how far the target's view of the lease may be behind the leader's.
	// The target's MCL controller re-reads the lease every 15 seconds.
	maxLeaseStatusLag = time.Minute
)

func newFailoverOptions() failoverOptions {
	return failoverOptions{
		timeout:                defaultFailoverTimeout,
		pollInterval:           defaultFailoverPollInterval,
		leaderAgreementTimeout: defaultFailoverLeaderAgreementTimeout,
		clusterOpTimeout:       defaultFailoverClusterOpTimeout,
		signalBounceWindow:     defaultSignalBounceWindow,
	}
}

const failoverLong = `Make the target cluster the leader of the multi-cluster HA group.

Failover stops Config Connector's manager on every other cluster in --contexts, waits until
the target cluster holds the lease, and then starts those managers again as standbys. From the
moment the leader's manager stops until the target takes over, no cluster reconciles resources.
That takes about a minute, because the leader's lease has to expire first.

Before it stops anything, failover checks that the target is a working standby: its managers
run, and its KRMSyncers sync from the leader. If a KRMSyncer reports that it doesn't, failover
refuses: the target's copy of the leader's status and resource IDs may be out of date, and the
target would lead with it. --skip-sync-check makes failover go ahead anyway.

Name every cluster of the group in --contexts. Failover doesn't stop the managers of a cluster
that isn't named, and such a cluster can take the lease instead of the target. Every cluster
that is named has to be reachable.

A manager is stopped by scaling its StatefulSet to 0. The replicas that it had are kept in the
` + failoverReplicasAnnotation + ` annotation of the StatefulSet. Failover starts
the managers again before it exits, also when it fails or is interrupted with Ctrl-C. If it is
killed before that, run it again: it starts the managers that are still stopped.`

func NewFailoverCmd() *cobra.Command {
	return newFailoverCmd(newFailoverOptions())
}

func newFailoverCmd(opts failoverOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "failover",
		Short: "Trigger a manual failover to a target standby cluster",
		Long:  failoverLong,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.stdout = cmd.OutOrStdout()
			opts.stderr = cmd.ErrOrStderr()
			if err := opts.validate(); err != nil {
				return err
			}
			// From here on, errors are about the clusters, not about how the command was used.
			cmd.SilenceUsage = true

			// When output is piped (for example `| tee failover.log`) and the reader exits on
			// Ctrl-C, let writes return EPIPE instead of killing the process mid-restart.
			signal.Ignore(syscall.SIGPIPE)

			// On Ctrl-C, SIGTERM, or SIGHUP, stop waiting and start the stopped managers again
			// before exiting. A second Ctrl-C after the bounce window exits at once.
			ctx, stop := notifyFailoverSignals(cmd.Context(), opts.stderr, opts.signalBounceWindow)
			defer stop()
			return runFailover(ctx, opts)
		},
	}
	opts.addFlags(cmd.Flags())
	return cmd
}

// notifyFailoverSignals cancels the returned context on the first interrupt/termination signal,
// prints immediate feedback, absorbs bounced signals during bounceWindow, and then restores
// default OS termination so a deliberate second Ctrl-C exits immediately.
func notifyFailoverSignals(parent context.Context, errOut io.Writer, bounceWindow time.Duration) (context.Context, context.CancelFunc) {
	sigs := []os.Signal{os.Interrupt, syscall.SIGTERM}
	if !signal.Ignored(syscall.SIGHUP) {
		sigs = append(sigs, syscall.SIGHUP)
	}
	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, sigs...)

	ctx, cancel := context.WithCancel(parent)
	done := make(chan struct{})
	stopped := make(chan struct{})
	var once sync.Once
	stop := func() {
		once.Do(func() {
			signal.Stop(sigCh)
			close(done)
			cancel()
		})
		<-stopped
	}

	go func() {
		defer close(stopped)
		select {
		case <-done:
			return
		case <-ctx.Done():
			signal.Stop(sigCh)
			return
		case <-sigCh:
			if errOut != nil {
				fmt.Fprintln(errOut, "\nInterrupted; starting any stopped managers again before exiting (press Ctrl-C again to exit at once).")
			}
			cancel()
		}

		if bounceWindow <= 0 {
			signal.Stop(sigCh)
			return
		}
		timer := time.NewTimer(bounceWindow)
		defer timer.Stop()
		for {
			select {
			case <-done:
				return
			case <-sigCh:
				// Ignore an immediate bounce so cleanup can finish.
			case <-timer.C:
				signal.Stop(sigCh)
				return
			}
		}
	}()

	return ctx, stop
}

const (
	targetClusterFlag   = "target-cluster"
	failoverTimeoutFlag = "timeout"
	failoverDryRunFlag  = "dry-run"
	skipSyncCheckFlag   = "skip-sync-check"
)

func (o *failoverOptions) addFlags(flags *pflag.FlagSet) {
	o.baseOptions.addFlags(flags)
	flags.StringVar(&o.targetCluster, targetClusterFlag, o.targetCluster, "Kubeconfig context name of the standby cluster to promote to leader.")
	flags.DurationVar(&o.timeout, failoverTimeoutFlag, o.timeout, "How long to wait for each step (at least 1m): for the other clusters' managers to stop, for the target to become the leader, and for the other clusters to follow it.")
	flags.BoolVar(&o.dryRun, failoverDryRunFlag, o.dryRun, "If true, print planned actions without making changes.")
	flags.BoolVar(&o.skipSyncCheck, skipSyncCheckFlag, o.skipSyncCheck, "If true, fail over even if the target's KRMSyncers don't report that they sync from the leader, and don't wait for the other clusters' KRMSyncers to report that they sync from the new leader. The target's copy of the leader's status and resource IDs may then be out of date.")
}

func (o *failoverOptions) validate() error {
	var errs []error
	if err := o.baseOptions.validate(); err != nil {
		errs = append(errs, err)
	}
	o.targetCluster = strings.TrimSpace(o.targetCluster)
	if o.targetCluster == "" {
		errs = append(errs, fmt.Errorf("--%s is required", targetClusterFlag))
	} else if !slices.Contains(o.contexts, o.targetCluster) {
		errs = append(errs, fmt.Errorf("--%s %q must be one of the contexts specified in --%s (%v)", targetClusterFlag, o.targetCluster, contextsFlag, o.contexts))
	}
	if o.timeout < minFailoverTimeout {
		errs = append(errs, fmt.Errorf("--%s must be at least %s (lease expiry takes about a minute), got %s", failoverTimeoutFlag, minFailoverTimeout, o.timeout))
	}
	return errors.Join(errs...)
}

func runFailover(ctx context.Context, opts failoverOptions) error {
	handles, err := buildClusterHandles(ctx, opts.contexts)
	if err != nil {
		return err
	}
	return executeFailover(ctx, opts, handles)
}

func executeFailover(ctx context.Context, opts failoverOptions, handles []*clusterHandle) error {
	f, err := newFailover(ctx, opts, handles)
	if err != nil {
		return err
	}
	return f.run(ctx)
}

// failover makes the target cluster the leader of the HA group.
//
// A cluster leads while it holds the lease, and it keeps the lease for as long as its Config
// Connector managers renew it. So failover stops the managers of every cluster but the target.
// The lease then expires, and the target, the only cluster that still asks for it, takes it.
// After that the other clusters' managers are started again, and they find the target leading.
type failover struct {
	clusters []*failoverCluster
	target   *failoverCluster

	out    io.Writer
	dryRun bool
	// skipSyncCheck is --skip-sync-check.
	skipSyncCheck bool

	timeout                time.Duration
	pollInterval           time.Duration
	leaderAgreementTimeout time.Duration
	clusterOpTimeout       time.Duration
}

func newFailover(ctx context.Context, opts failoverOptions, handles []*clusterHandle) (*failover, error) {
	states, err := fetchClusterStatesForFailover(ctx, handles, opts.clusterOpTimeout)
	if err != nil {
		return nil, err
	}
	if err := validateMultiClusterLeaseConfigs(states); err != nil {
		return nil, err
	}

	opTimeout := opts.clusterOpTimeout
	if opTimeout <= 0 {
		opTimeout = defaultFailoverClusterOpTimeout
	}
	f := &failover{
		out:                    opts.stdout,
		dryRun:                 opts.dryRun,
		skipSyncCheck:          opts.skipSyncCheck,
		timeout:                opts.timeout,
		pollInterval:           opts.pollInterval,
		leaderAgreementTimeout: opts.leaderAgreementTimeout,
		clusterOpTimeout:       opTimeout,
	}
	if f.out == nil {
		f.out = io.Discard
	}
	for i, state := range states {
		// fetchClusterStatesForFailover returns one state per handle, in order.
		c := newFailoverCluster(handles[i], state)
		f.clusters = append(f.clusters, c)
		if c.name() == opts.targetCluster {
			f.target = c
		}
	}
	if f.target == nil {
		return nil, fmt.Errorf("--%s %q is not one of the clusters in --%s", targetClusterFlag, opts.targetCluster, contextsFlag)
	}
	return f, nil
}

// fetchClusterStatesForFailover fetches each cluster's HA state under its own opTimeout so one
// stalled cluster cannot prevent reporting errors from the others.
func fetchClusterStatesForFailover(ctx context.Context, handles []*clusterHandle, opTimeout time.Duration) ([]*clusterHAState, error) {
	if opTimeout <= 0 {
		opTimeout = defaultFailoverClusterOpTimeout
	}
	var errs []error
	states := make([]*clusterHAState, 0, len(handles))
	for _, h := range handles {
		callCtx, cancel := context.WithTimeout(ctx, opTimeout)
		state, err := fetchClusterHAState(callCtx, h)
		cancel()
		if err != nil {
			errs = append(errs, err)
			if ctx.Err() != nil {
				break
			}
			continue
		}
		states = append(states, state)
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return states, nil
}

func (f *failover) logf(format string, args ...any) {
	fmt.Fprintf(f.out, format+"\n", args...)
}

func (f *failover) run(ctx context.Context) error {
	leader, views, err := f.currentLeader(ctx)
	if err != nil {
		return err
	}
	f.logf("Leader: %s", leader)

	if leader == f.target {
		return f.finishEarlierFailover(ctx, views)
	}
	f.logf("Target: %s", f.target)

	// Once the leader stops, nothing is reconciled until the target takes over. So first make
	// sure that the target can.
	if err := f.checkTargetSeesLease(leader, views); err != nil {
		return err
	}
	if err := f.prepareTarget(ctx, leader, views[f.target]); err != nil {
		return err
	}

	if f.dryRun {
		f.printPlan(leader, views)
		return nil
	}

	f.logf("If this command doesn't finish, run it again: it starts the managers that are still stopped.")
	handOverErr := f.handOver(ctx, leader, views[f.target].renewTime)
	// Start the other clusters' managers again, whether or not the target took over.
	startErr := f.startOtherManagers(ctx)
	if handOverErr != nil {
		if startErr != nil {
			return fmt.Errorf("failover to cluster %q did not complete: %w", f.target.name(), errors.Join(handOverErr, startErr))
		}
		// With every manager running again, any cluster can end up with the lease.
		return fmt.Errorf("failover to cluster %q did not complete, and the other clusters' managers were started again; run the command with --%s to see which cluster leads now: %w", f.target.name(), failoverDryRunFlag, handOverErr)
	}
	if startErr != nil {
		return fmt.Errorf("cluster %q is the leader, but not every manager on the other clusters was started again: %w", f.target.name(), startErr)
	}

	if err := f.waitForStandbys(ctx); err != nil {
		return fmt.Errorf("cluster %q is the leader, but: %w", f.target.name(), err)
	}
	f.logf("Failover complete: %s is the leader.", f.target)
	return nil
}

// currentLeader returns the cluster that holds the lease, and what every cluster looks like,
// once the clusters agree on the holder.
func (f *failover) currentLeader(ctx context.Context) (*failoverCluster, map[*failoverCluster]*failoverView, error) {
	views := make(map[*failoverCluster]*failoverView, len(f.clusters))
	var holder string
	var readErr error
	err := poll(ctx, f.pollInterval, f.leaderAgreementTimeout, func(ctx context.Context) (string, error) {
		for _, c := range f.clusters {
			v, err := c.view(ctx)
			if err != nil {
				if ctx.Err() != nil {
					// The time ran out during the request. That isn't the cluster's fault.
					return err.Error(), nil
				}
				readErr = err
				return "", err
			}
			views[c] = v
		}

		holder = views[f.target].holder
		if holder == "" {
			return fmt.Sprintf("cluster %q doesn't report a leader: the status of its MultiClusterLease %s names no holder; check that the MultiClusterLease controller runs on it",
				f.target.name(), f.target.lease), nil
		}
		for _, c := range f.clusters {
			if h := views[c].holder; h != "" && h != holder {
				return "the clusters disagree: " + f.describeHolders(views), nil
			}
		}
		return "", nil
	})
	if readErr != nil {
		return nil, nil, readErr
	}
	if err != nil {
		return nil, nil, fmt.Errorf("cannot tell which cluster is the leader: %w", err)
	}

	leader := f.clusterOf(holder)
	if leader == nil {
		return nil, nil, fmt.Errorf("the leader, %s, is not one of the clusters in --%s; add its context, because failover has to stop its managers", holder, contextsFlag)
	}
	return leader, views, nil
}

func (f *failover) describeHolders(views map[*failoverCluster]*failoverView) string {
	parts := make([]string, 0, len(f.clusters))
	for _, c := range f.clusters {
		holder := views[c].holder
		if holder == "" {
			holder = "no leader"
		}
		parts = append(parts, fmt.Sprintf("cluster %q reports %s", c.name(), holder))
	}
	return strings.Join(parts, "; ")
}

// checkTargetSeesLease checks that the target's view of the lease is current. The target learns
// the state of the lease from its own MCL controller, and takes the lease through it. If that
// controller has stopped reading the lease, the target can't take over.
//
// Both times come from the lease itself, so the check doesn't depend on this machine's clock.
func (f *failover) checkTargetSeesLease(leader *failoverCluster, views map[*failoverCluster]*failoverView) error {
	for _, c := range []*failoverCluster{f.target, leader} {
		if views[c].renewTime.IsZero() {
			return fmt.Errorf("cluster %q doesn't report when the lease was last renewed: the status of its MultiClusterLease %s has no valid globalRenewTime; check the MultiClusterLease controller on it",
				c.name(), c.lease)
		}
	}
	targetRenewal, leaderRenewal := views[f.target].renewTime, views[leader].renewTime
	if lag := leaderRenewal.Sub(targetRenewal); lag > maxLeaseStatusLag {
		return fmt.Errorf("cluster %q is not ready to take over: its view of the lease is %s behind the leader's (it last saw a renewal at %s, the leader at %s); check that the MultiClusterLease controller runs on it and can read the lease bucket",
			f.target.name(), lag.Round(time.Second), targetRenewal.Format(time.RFC3339), leaderRenewal.Format(time.RFC3339))
	}
	return nil
}

// prepareTarget checks that the target is a working standby: its managers run, and its
// KRMSyncers pull from the leader. If an earlier failover was interrupted and left the target's
// managers stopped, it starts them first.
func (f *failover) prepareTarget(ctx context.Context, leader *failoverCluster, view *failoverView) error {
	target := f.target
	if len(view.managers) == 0 {
		return fmt.Errorf("cluster %q is not ready to take over: it has no StatefulSet of Config Connector's manager", target.name())
	}

	if stopped := stoppedByFailover(view.managers); len(stopped) > 0 {
		if f.dryRun {
			f.logf("[dry-run] cluster %q: would start %s, which an earlier failover stopped, and wait for the cluster to follow the leader", target.name(), describeManagers(stopped))
			return nil
		}
		f.logf("cluster %q: an earlier failover left its managers stopped", target.name())
		callCtx, cancel := context.WithTimeout(ctx, f.clusterOpTimeout)
		startErr := target.startManagers(callCtx, f.out)
		cancel()
		if startErr != nil {
			return startErr
		}
		f.logf("cluster %q: waiting up to %s for it to follow the leader", target.name(), f.timeout)
		var moved error
		err := poll(ctx, f.pollInterval, f.timeout, func(ctx context.Context) (string, error) {
			v, err := target.view(ctx)
			if err != nil {
				return err.Error(), nil
			}
			// The earlier failover may have stopped the leader too. Its lease then expires
			// while the target's managers start.
			if v.holder != "" && v.holder != leader.identity {
				moved = fmt.Errorf("the lease went from %s to %s while failover started the managers of cluster %q; run the command again", leader.identity, v.holder, target.name())
				return "", moved
			}
			if pending := target.pendingAsTarget(v, leader); pending != "" {
				return pending, nil
			}
			return f.pendingSync(target, v, leader), nil
		})
		if moved != nil {
			return moved
		}
		if err != nil {
			return fmt.Errorf("cluster %q is not ready to take over: %w", target.name(), err)
		}
		return nil
	}

	if pending := target.pendingAsTarget(view, leader); pending != "" {
		return fmt.Errorf("cluster %q is not ready to take over: %s", target.name(), pending)
	}
	if problem := target.notSyncing(view, leader); problem != "" {
		if !f.skipSyncCheck {
			return fmt.Errorf("cluster %q is not ready to take over: it does not sync from the leader (--%s fails over anyway): %s", target.name(), skipSyncCheckFlag, problem)
		}
		f.logf("cluster %q: --%s is set, so failing over although it does not sync from the leader: %s", target.name(), skipSyncCheckFlag, problem)
	}
	for _, m := range view.managers {
		if m.replicas() == 0 {
			f.logf("cluster %q: note: StatefulSet %s is scaled to 0, so that manager will not run after the failover either", target.name(), m.key())
		}
	}
	return nil
}

// pendingSync returns why the KRMSyncers of a standby don't report that they sync from the
// leader, or "" if they do. With --skip-sync-check it returns "": failover then only checks
// what the managers ask the syncer for (pendingSyncers).
func (f *failover) pendingSync(c *failoverCluster, v *failoverView, leader *failoverCluster) string {
	if f.skipSyncCheck {
		return ""
	}
	return c.notSyncing(v, leader)
}

func (f *failover) printPlan(leader *failoverCluster, views map[*failoverCluster]*failoverView) {
	for _, c := range f.othersLeaderLast(leader) {
		var running []*managerStatefulSet
		for _, m := range views[c].managers {
			if m.replicas() > 0 {
				running = append(running, m)
			}
		}
		if len(running) == 0 {
			f.logf("[dry-run] cluster %q: has no running manager to stop", c.name())
			continue
		}
		f.logf("[dry-run] cluster %q: would stop %s, and wait up to %s for its pods to stop", c.name(), describeManagers(running), f.timeout)
	}
	f.logf("[dry-run] cluster %q: would wait up to %s for it to become the leader", f.target.name(), f.timeout)
	f.logf("[dry-run] would then start the stopped managers again, and wait up to %s for their clusters to follow the new leader", f.timeout)
}

// handOver stops the managers of every cluster but the target, and waits for the target to take
// the lease. lastRenewal is the last renewal of the lease that the target saw before.
func (f *failover) handOver(ctx context.Context, leader *failoverCluster, lastRenewal time.Time) error {
	// The leader goes last. Once it stops, the lease is up for grabs, and by then the target must
	// be the only cluster left that asks for it.
	stopped := f.othersLeaderLast(leader)
	for _, c := range stopped {
		callCtx, cancel := context.WithTimeout(ctx, f.clusterOpTimeout)
		waiting, err := c.stopManagers(callCtx, f.out)
		cancel()
		if err != nil {
			return err
		}
		if !waiting {
			f.logf("cluster %q: has no running manager to stop", c.name())
			continue
		}
		f.logf("cluster %q: waiting up to %s for its manager pods to stop", c.name(), f.timeout)
		err = poll(ctx, f.pollInterval, f.timeout, c.pendingStop)
		if err != nil {
			return fmt.Errorf("cluster %q: Config Connector's manager did not stop: %w", c.name(), err)
		}
	}
	f.logf("No cluster reconciles resources until cluster %q takes the lease. The old lease has to expire first, which takes about a minute.", f.target.name())
	return f.waitForTargetToLead(ctx, lastRenewal, stopped)
}

// waitForTargetToLead waits until the target holds the lease, has renewed it after lastRenewal,
// and has suspended its KRMSyncers. stopped are the clusters whose managers failover has stopped.
//
// The renewal shows that the target's managers are running: the status of a lease keeps naming
// its last holder after that holder has stopped.
//
// A stopped cluster can hold the lease for a while yet. Its MCL controller keeps asking for the
// lease for up to 30 seconds after its managers are gone, so it renews the lease that the
// cluster holds, and it can even take the lease when the leader's expires. Nothing renews the
// lease after that, so the target gets it in the end. Any other holder has running managers, and
// keeps the lease.
func (f *failover) waitForTargetToLead(ctx context.Context, lastRenewal time.Time, stopped []*failoverCluster) error {
	target := f.target
	f.logf("cluster %q: waiting up to %s for it to become the leader", target.name(), f.timeout)
	err := poll(ctx, f.pollInterval, f.timeout, func(ctx context.Context) (string, error) {
		v, err := target.view(ctx)
		if err != nil {
			return err.Error(), nil
		}
		if v.holder == "" {
			return fmt.Sprintf("its MultiClusterLease %s names no holder", target.lease), nil
		}
		holder := f.clusterOf(v.holder)
		switch {
		case holder == target:
			if !v.renewTime.After(lastRenewal) {
				return "its MultiClusterLease names it as the holder, but it has not renewed the lease yet", nil
			}
			return target.pendingAsTarget(v, target), nil
		case slices.Contains(stopped, holder):
			return fmt.Sprintf("its MultiClusterLease %s still names %s as the holder", target.lease, v.holder), nil
		case holder != nil:
			return "", fmt.Errorf("cluster %q took the lease instead", holder.name())
		default:
			return "", fmt.Errorf("%s took the lease instead; it is not one of the clusters in --%s, so failover did not stop its managers", v.holder, contextsFlag)
		}
	})
	if err != nil {
		return fmt.Errorf("cluster %q did not become the leader: %w", target.name(), err)
	}
	f.logf("cluster %q: is the leader", target.name())
	return nil
}

// clusterOf returns the cluster with the given identity in the election, or nil if it is not one
// of the clusters in --contexts.
func (f *failover) clusterOf(identity string) *failoverCluster {
	for _, c := range f.clusters {
		if c.identity == identity {
			return c
		}
	}
	return nil
}

// startManagersOn starts any managers that failover stopped on c, under a fresh per-cluster
// timeout detached from ctx's cancellation so one unresponsive cluster cannot starve the others.
func (f *failover) startManagersOn(ctx context.Context, c *failoverCluster) error {
	callCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), f.clusterOpTimeout)
	defer cancel()
	return c.startManagers(callCtx, f.out)
}

// startOtherManagers starts the managers that failover stopped on the clusters other than the
// target. It also runs after the command was interrupted, so it doesn't use ctx's deadline or
// cancellation.
func (f *failover) startOtherManagers(ctx context.Context) error {
	var errs []error
	for _, c := range f.clusters {
		if c == f.target {
			continue
		}
		if err := f.startManagersOn(ctx, c); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// waitForStandbys waits until every cluster other than the target is a standby of the target:
// its managers run, and its KRMSyncers pull from the target.
func (f *failover) waitForStandbys(ctx context.Context) error {
	var errs []error
	for _, c := range f.clusters {
		if c == f.target {
			continue
		}
		f.logf("cluster %q: waiting up to %s for it to follow the new leader", c.name(), f.timeout)
		var hasRunning bool
		err := poll(ctx, f.pollInterval, f.timeout, func(ctx context.Context) (string, error) {
			v, err := c.view(ctx)
			if err != nil {
				return err.Error(), nil
			}
			hasRunning = hasRunningManager(v.managers)
			if pending := c.pendingAsStandby(v, f.target); pending != "" {
				return pending, nil
			}
			return f.pendingSync(c, v, f.target), nil
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("cluster %q does not follow it yet: %w", c.name(), err))
			if ctx.Err() != nil {
				break
			}
			continue
		}
		if hasRunning {
			f.logf("cluster %q: follows the new leader", c.name())
		} else {
			f.logf("cluster %q: has no running manager", c.name())
		}
	}
	return errors.Join(errs...)
}

// finishEarlierFailover handles a target that already holds the lease. Usually there is nothing
// to do. But if an earlier failover was interrupted, some managers are still stopped: it starts
// them.
func (f *failover) finishEarlierFailover(ctx context.Context, views map[*failoverCluster]*failoverView) error {
	target := f.target
	var stopped []*failoverCluster
	for _, c := range f.clusters {
		if len(stoppedByFailover(views[c].managers)) > 0 {
			stopped = append(stopped, c)
		}
	}
	if len(stopped) == 0 {
		// MultiClusterLease.status.globalHolderIdentity keeps naming the last holder after its
		// managers stop, so verify the target's lease status, managers, and syncers before
		// reporting that nothing needs to be done.
		if err := f.checkTargetSeesLease(target, views); err != nil {
			return err
		}
		if pending := target.pendingAsTarget(views[target], target); pending != "" {
			return fmt.Errorf("cluster %q holds the lease, but: %s", target.name(), pending)
		}
		f.logf("Cluster %q is already the leader. Nothing to do.", target.name())
		return nil
	}

	if f.dryRun {
		for _, c := range stopped {
			f.logf("[dry-run] cluster %q: would start %s, which an earlier failover stopped", c.name(), describeManagers(stoppedByFailover(views[c].managers)))
		}
		// A real run checks the target's view of the lease once it has started the managers.
		return f.checkTargetSeesLease(target, views)
	}

	f.logf("Cluster %q already holds the lease. Starting the managers that an earlier failover stopped.", target.name())
	// Start the managers of every cluster, also if those of one can't be started.
	targetErr := f.startManagersOn(ctx, target)
	if err := errors.Join(targetErr, f.startOtherManagers(ctx)); err != nil {
		return err
	}
	// The target's view of the lease is checked only now, so that a lease status without a
	// renewal time doesn't keep the managers stopped. Without that time, waitForTargetToLead
	// can't tell a new renewal from an old one.
	if err := f.checkTargetSeesLease(target, views); err != nil {
		return fmt.Errorf("the stopped managers were started again, but: %w", err)
	}
	// No cluster's managers are stopped now, so another cluster that holds the lease keeps it.
	if err := f.waitForTargetToLead(ctx, views[target].renewTime, nil); err != nil {
		return err
	}
	if err := f.waitForStandbys(ctx); err != nil {
		return fmt.Errorf("cluster %q is the leader, but: %w", target.name(), err)
	}
	f.logf("Failover complete: %s is the leader.", target)
	return nil
}

// othersLeaderLast returns every cluster but the target, with the leader last.
func (f *failover) othersLeaderLast(leader *failoverCluster) []*failoverCluster {
	var others []*failoverCluster
	for _, c := range f.clusters {
		if c != f.target && c != leader {
			others = append(others, c)
		}
	}
	if leader != f.target {
		others = append(others, leader)
	}
	return others
}

// poll calls check every interval until it reports that nothing is pending, for at most timeout.
// check returns what it is still waiting for, or an error to stop waiting at once. If the time
// runs out, the error says what check was last waiting for.
func poll(ctx context.Context, interval, timeout time.Duration, check func(ctx context.Context) (pending string, err error)) error {
	var pending string
	err := wait.PollUntilContextTimeout(ctx, interval, timeout, true, func(pollCtx context.Context) (bool, error) {
		p, err := check(pollCtx)
		if err != nil {
			return false, err
		}
		// When the time runs out during a request, the request fails. What check was waiting
		// for before that says more than the failed request does.
		if p == "" || pending == "" || pollCtx.Err() == nil {
			pending = p
		}
		return p == "", nil
	})
	switch {
	case err == nil:
		return nil
	case ctx.Err() != nil:
		return fmt.Errorf("interrupted: %w", ctx.Err())
	case wait.Interrupted(err):
		return fmt.Errorf("timed out after %s: %s", timeout, pending)
	default:
		return err
	}
}
