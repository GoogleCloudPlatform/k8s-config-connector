// Copyright 2024 Google LLC
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

package leaderelection

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/record"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/resourcelock"
)

func New(c Config) *LeaderElector {
	// TODO: validate config
	return &LeaderElector{
		config: c,
	}
}

type Config struct {
	ElectionID string
	Identity   string
	// Lock is the resorurce lock that will be used to hold leader election lease record.
	Lock resourcelock.Interface
	// LeaseDuration is the duration that non-leader candidates will
	// wait to force acquire leadership. This is measured against time of
	// last observed ack.
	LeaseDuration time.Duration
	// RenewDeadline is the duration that the acting master will retry
	// refreshing leadership before giving up.
	RenewDeadline time.Duration
	// RetryPeriod is the duration the LeaderElector clients should wait
	// between tries of actions.
	RetryPeriod   time.Duration
	Log           logr.Logger
	EventRecorder record.EventRecorder
}

type LeaderElector struct {
	config Config

	// observedRecordMutex protects the observedRecord field
	observedRecordMutex sync.Mutex
	observedRecord      resourcelock.LeaderElectionRecord
	observedRawRecord   []byte

	observedTime time.Time
}

func (le *LeaderElector) Identity() string {
	return le.config.Identity
}

func (le *LeaderElector) ElectionID() string {
	return le.config.ElectionID
}

func (le *LeaderElector) LeaderIdentity() string {
	return le.getObservedRecord().HolderIdentity
}

func (le *LeaderElector) IsLeader() bool {
	return le.getObservedRecord().HolderIdentity == le.config.Identity
}

func (le *LeaderElector) Renew(ctx context.Context) error {
	log := le.config.Log.WithValues("lease", le.config.Lock.Describe())
	log.Info("attempting to renew lease")

	internalCtx, cancel := context.WithTimeout(ctx, le.config.RenewDeadline)
	defer cancel()
	if err := wait.PollUntilContextTimeout(ctx, le.config.RetryPeriod, le.config.RenewDeadline, true, func(context.Context) (bool, error) {
		return le.tryAcquireOrRenew(internalCtx), nil
	}); err != nil {
		log.Info("failed to renew lease")
		return err
	}

	log.Info("successfully renewed lease")
	return nil
}

func (le *LeaderElector) Acquire(ctx context.Context) error {
	log := le.config.Log.WithValues("lease", le.config.Lock.Describe())
	log.Info("attempting to acquire leader lease")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	succeeded := le.tryAcquireOrRenew(ctx)
	// TODO: return (bool, error) so that we can differentiate between transient error (e.g. network error) and someone else has the lease
	if !succeeded {
		log.Info("failed to acquire lease")
		return fmt.Errorf("failed to acquire lease")
	}

	log.Info("successfully acquired lease")
	return nil
}

func (le *LeaderElector) tryAcquireOrRenew(ctx context.Context) bool {
	log := le.config.Log.WithValues("lease", le.config.Lock.Describe())

	now := metav1.NewTime(time.Now())
	leaderElectionRecord := resourcelock.LeaderElectionRecord{
		HolderIdentity:       le.config.Lock.Identity(),
		LeaseDurationSeconds: int(le.config.LeaseDuration / time.Second),
		RenewTime:            now,
		AcquireTime:          now,
	}

	// 1. obtain or create the ElectionRecord
	oldLeaderElectionRecord, oldLeaderElectionRawRecord, err := le.config.Lock.Get(ctx)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			log.Error(err, "error retrieving resource lock")
			return false
		}
		if err = le.config.Lock.Create(ctx, leaderElectionRecord); err != nil {
			log.Error(err, "error initially creating leader election record")
			return false
		}

		le.setObservedRecord(&leaderElectionRecord)

		return true
	}

	// 2. Record obtained, check the Identity & Time
	if !bytes.Equal(le.observedRawRecord, oldLeaderElectionRawRecord) {
		le.setObservedRecord(oldLeaderElectionRecord)
		le.observedRawRecord = oldLeaderElectionRawRecord
	}
	if len(oldLeaderElectionRecord.HolderIdentity) > 0 &&
		le.observedTime.Add(time.Second*time.Duration(oldLeaderElectionRecord.LeaseDurationSeconds)).After(now.Time) &&
		!le.IsLeader() {
		log.Info("lock is held by someone else and has not yet expired", "holder identity", oldLeaderElectionRecord.HolderIdentity)
		return false
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
	if err = le.config.Lock.Update(ctx, leaderElectionRecord); err != nil {
		log.Error(err, "error updating lock")
		return false
	}

	le.setObservedRecord(&leaderElectionRecord)
	return true
}

func (le *LeaderElector) getObservedRecord() resourcelock.LeaderElectionRecord {
	le.observedRecordMutex.Lock()
	defer le.observedRecordMutex.Unlock()

	return le.observedRecord
}

func (le *LeaderElector) setObservedRecord(record *resourcelock.LeaderElectionRecord) {
	le.observedRecordMutex.Lock()
	defer le.observedRecordMutex.Unlock()

	le.observedRecord = *record
	le.observedTime = time.Now()
}
