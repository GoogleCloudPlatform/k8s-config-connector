package leaderelection

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	rl "github.com/600lyy/multi-cluster-leader-election/pkg/resourcelock"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

const (
	JitterFactor = 1.2
)

// NewLeaderElector creates a LeaderElector from a LeaderElectionConfig
func NewLeaderElector(lec LeaderElectionConfig) (*LeaderElector, error) {
	if lec.LeaseDuration <= lec.RenewDeadline {
		return nil, fmt.Errorf("leaseDuration must be greater than renewDeadline")
	}
	if lec.RenewDeadline <= time.Duration(JitterFactor*float64(lec.RetryPeriod)) {
		return nil, fmt.Errorf("renewDeadline must be greater than retryPeriod*JitterFactor")
	}
	if lec.LeaseDuration < 1 {
		return nil, fmt.Errorf("leaseDuration must be greater than zero")
	}
	if lec.RenewDeadline < 1 {
		return nil, fmt.Errorf("renewDeadline must be greater than zero")
	}
	if lec.RetryPeriod < 1 {
		return nil, fmt.Errorf("retryPeriod must be greater than zero")
	}

	if lec.Lock == nil {
		return nil, fmt.Errorf("lock must not be nil")
	}

	le := LeaderElector{
		Config: lec,
		Clock:  clock.RealClock{},
	}
	return &le, nil
}

type LeaderElectionConfig struct {
	// Lock is the resource that will be used for locking
	Lock rl.Interface

	// LeaseDuration is the duration that non-leader candidates will
	// wait to force acquire leadership. This is measured against time of
	// last observed ack.
	//
	// A client needs to wait a full LeaseDuration without observing a change to
	// the record before it can attempt to take over. When all clients are
	// shutdown and a new set of clients are started with different names against
	// the same leader record, they must wait the full LeaseDuration before
	// attempting to acquire the lease. Thus LeaseDuration should be as short as
	// possible (within your tolerance for clock skew rate) to avoid a possible
	// long waits in the scenario.
	//
	// Core clients default this value to 15 seconds.
	LeaseDuration time.Duration
	// RenewDeadline is the duration that the acting master will retry
	// refreshing leadership before giving up.
	//
	// Core clients default this value to 10 seconds.
	RenewDeadline time.Duration
	// RetryPeriod is the duration the LeaderElector clients should wait
	// between tries of actions.
	//
	// Core clients default this value to 2 seconds.
	RetryPeriod time.Duration
}

// LeaderElector is a leader election client.
type LeaderElector struct {
	Config LeaderElectionConfig
	// internal bookkeeping
	ObservedRecord    rl.LeaderElectionRecord
	ObservedRawRecord []byte
	ObservedTime      time.Time
	// used to implement OnNewLeader(), may lag slightly from the
	// value observedRecord.HolderIdentity if the transition has
	// not yet been reported.
	ReportedLeader string
	// clock is wrapper around time to allow for less flaky testing
	Clock clock.Clock
	// used to lock the observedRecord
	ObservedRecordLock sync.Mutex
}

// GetLeader returns the identity of the last observed leader or returns the empty string if
// no leader has yet been observed.
// This function is for informational purposes. (e.g. monitoring, logs, etc.)
func (le *LeaderElector) GetLeader() string {
	return le.getObservedRecord().HolderIdentity
}

func (le *LeaderElector) getObservedRecord() rl.LeaderElectionRecord {
	le.ObservedRecordLock.Lock()
	defer le.ObservedRecordLock.Unlock()

	return le.ObservedRecord
}

func (le *LeaderElector) TryAcquireOrRenew(ctx context.Context) error {
	now := metav1.NewTime(le.Clock.Now())
	leaderElectionRecord := rl.LeaderElectionRecord{
		HolderIdentity:       le.Config.Lock.Identity(),
		LeaseDurationSeconds: int(le.Config.LeaseDuration / time.Second),
		RenewTime:            now,
		AcquireTime:          now,
	}
	// 1. obtain or create the ElectionRecord
	oldLeaderElectionRecord, oldLeaderElectionRawRecord, err := le.Config.Lock.Get(ctx)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			klog.Errorf("error retrieving resource lock %v: %v", le.Config.Lock.Describe(), err)
			return err
		}
		if err = le.Config.Lock.Create(ctx, leaderElectionRecord); err != nil {
			klog.Errorf("error initially creating leader election record: %v", err)
			return err
		}

		le.setObservedRecord(&leaderElectionRecord)
		return nil
	}

	// 2. Record obtained, check the Identity & Time
	if !bytes.Equal(le.ObservedRawRecord, oldLeaderElectionRawRecord) {
		le.setObservedRecord(oldLeaderElectionRecord)
		le.ObservedRawRecord = oldLeaderElectionRawRecord
	}
	if len(oldLeaderElectionRecord.HolderIdentity) > 0 &&
		le.ObservedTime.Add(time.Second*time.Duration(oldLeaderElectionRecord.LeaseDurationSeconds)).After(now.Time) &&
		!le.IsLeader() {
		klog.Infof("lock %v is held by %v and has not yet expired", le.Config.Lock.Describe(), oldLeaderElectionRecord.HolderIdentity)
		return nil
	}

	// 3. We're going to try to update. The leaderElectionRecord is set to it's default
	// here. Let's correct it before updating.
	if le.IsLeader() {
		leaderElectionRecord.AcquireTime = oldLeaderElectionRecord.AcquireTime
		leaderElectionRecord.LeaderTransitions = oldLeaderElectionRecord.LeaderTransitions
	} else {
		leaderElectionRecord.LeaderTransitions = oldLeaderElectionRecord.LeaderTransitions + 1
	}

	// update the lock itself
	if err = le.Config.Lock.Update(ctx, leaderElectionRecord); err != nil {
		klog.Errorf("failed to update lock: %v", err)
		return err
	}

	le.setObservedRecord(&leaderElectionRecord)
	return nil
}

func (le *LeaderElector) setObservedRecord(observedRecord *rl.LeaderElectionRecord) {
	le.ObservedRecordLock.Lock()
	defer le.ObservedRecordLock.Unlock()

	le.ObservedRecord = *observedRecord
	le.ObservedTime = le.Clock.Now()
}

func (le *LeaderElector) IsLeader() bool {
	return (le.getObservedRecord().HolderIdentity == le.Config.Lock.Identity())
}
