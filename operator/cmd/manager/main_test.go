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

package main

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/leaderelection"
)

type fakeRecorderProvider struct{}

func (fakeRecorderProvider) GetEventRecorderFor(string) record.EventRecorder {
	return record.NewFakeRecorder(10)
}

func TestNewManagerOptionsLeaderElection(t *testing.T) {
	grid := []struct {
		name                 string
		enableLeaderElection bool
		leaderElectionID     string
	}{
		{name: "disabled", enableLeaderElection: false, leaderElectionID: defaultLeaderElectionID},
		{name: "enabled with default id", enableLeaderElection: true, leaderElectionID: defaultLeaderElectionID},
		{name: "enabled with custom id", enableLeaderElection: true, leaderElectionID: "custom-lease"},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			opts := newManagerOptions(runtime.NewScheme(), ":0", g.enableLeaderElection, g.leaderElectionID)

			if opts.LeaderElection != g.enableLeaderElection {
				t.Errorf("LeaderElection = %v, want %v", opts.LeaderElection, g.enableLeaderElection)
			}
			if opts.LeaderElectionID != g.leaderElectionID {
				t.Errorf("LeaderElectionID = %q, want %q", opts.LeaderElectionID, g.leaderElectionID)
			}
			if !opts.LeaderElectionReleaseOnCancel {
				t.Errorf("LeaderElectionReleaseOnCancel = false, want true")
			}

			// The manager builds its lock with these same options; this is where
			// an empty LeaderElectionID used to fail with "LeaderElectionID must be configured".
			lock, err := leaderelection.NewResourceLock(&rest.Config{Host: "https://127.0.0.1"}, fakeRecorderProvider{}, leaderelection.Options{
				LeaderElection:             opts.LeaderElection,
				LeaderElectionResourceLock: opts.LeaderElectionResourceLock,
				LeaderElectionNamespace:    "configconnector-operator-system",
				LeaderElectionID:           opts.LeaderElectionID,
			})
			if err != nil {
				t.Fatalf("NewResourceLock() error = %v", err)
			}
			if !g.enableLeaderElection {
				if lock != nil {
					t.Errorf("NewResourceLock() = %v, want nil when leader election is disabled", lock)
				}
				return
			}
			if want := "configconnector-operator-system/" + g.leaderElectionID; lock.Describe() != want {
				t.Errorf("lock.Describe() = %q, want %q", lock.Describe(), want)
			}
		})
	}
}
