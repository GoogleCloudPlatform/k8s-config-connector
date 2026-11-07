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
	"slices"
	"strings"
	"testing"
	"time"
)

type fakeOutputRunner struct {
	responses map[string]string
	errors    map[string]error
}

func (f *fakeOutputRunner) output(_ context.Context, name string, args ...string) (string, error) {
	cmdLine := formatCommandLine(name, args)
	if err, ok := f.errors[cmdLine]; ok {
		return "", err
	}
	if out, ok := f.responses[cmdLine]; ok {
		return out, nil
	}
	return "", fmt.Errorf("CommandException: One or more URLs matched no objects: %s", cmdLine)
}

func TestGCSLockInspectorHealthyLock(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 0, 10, 0, time.UTC)
	runner := &fakeOutputRunner{
		responses: map[string]string{
			"gcloud storage ls gs://my-ha-bucket/leases/**":                                  "gs://my-ha-bucket/leases/cnrm-system/kcc-multicluster-lease\n",
			"gcloud storage cat gs://my-ha-bucket/leases/cnrm-system/kcc-multicluster-lease": `{"holderIdentity":"cluster-a","renewTime":"2026-10-28T03:00:05Z","leaseTransitions":2}`,
		},
	}
	inspector := newGCSLockInspector(runner, "my-ha-bucket", "cnrm-system", "kcc-multicluster-lease", now)

	status, warnings := inspector.inspect(context.Background())

	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %v", warnings)
	}
	if !status.found || status.holderIdentity != "cluster-a" || status.renewAge != 5*time.Second || status.leaseTransitions != 2 {
		t.Fatalf("unexpected lock status: %+v", status)
	}
}

func TestGCSLockInspectorDetectsExtraKeysAndStaleLock(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 5, 0, 0, time.UTC)
	runner := &fakeOutputRunner{
		responses: map[string]string{
			"gcloud storage ls gs://my-ha-bucket/leases/**": strings.Join([]string{
				"gs://my-ha-bucket/leases/cnrm-system/lease-a",
				"gs://my-ha-bucket/leases/cnrm-system/lease-b",
			}, "\n"),
			"gcloud storage cat gs://my-ha-bucket/leases/cnrm-system/lease-a": `{"holderIdentity":"cluster-a","renewTime":"2026-10-28T03:00:00Z","leaseTransitions":1}`,
		},
	}
	inspector := newGCSLockInspector(runner, "my-ha-bucket", "cnrm-system", "lease-a", now)

	status, warnings := inspector.inspect(context.Background())

	if !slices.Equal(status.extraKeys, []string{"leases/cnrm-system/lease-b"}) {
		t.Fatalf("got extraKeys %v, want [leases/cnrm-system/lease-b]", status.extraKeys)
	}
	if len(warnings) != 2 {
		t.Fatalf("expected 2 warnings (extra keys and stale lock), got %v", warnings)
	}
}
