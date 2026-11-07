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
	"sort"
	"strings"
	"time"
)

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

// gcsLockInspector inspects the global multi-cluster lease lock object and lease key prefix in GCS.
type gcsLockInspector struct {
	runner          commandOutputRunner
	bucket          string
	expectedLockKey string
	lockURI         string
	now             time.Time
}

func newGCSLockInspector(runner commandOutputRunner, bucket, leaseNamespace, leaseName string, now time.Time) *gcsLockInspector {
	expectedKey := fmt.Sprintf("leases/%s/%s", leaseNamespace, leaseName)
	return &gcsLockInspector{
		runner:          runner,
		bucket:          bucket,
		expectedLockKey: expectedKey,
		lockURI:         fmt.Sprintf("gs://%s/%s", bucket, expectedKey),
		now:             now,
	}
}

func (g *gcsLockInspector) inspect(ctx context.Context) (*gcsLockStatus, []string) {
	status := &gcsLockStatus{
		bucket:          g.bucket,
		expectedLockKey: g.expectedLockKey,
		lockURI:         g.lockURI,
	}
	var warnings []string

	extraKeys, lsWarnings := g.findExtraLeaseKeys(ctx)
	status.extraKeys = extraKeys
	warnings = append(warnings, lsWarnings...)

	data, catWarnings := g.readLockObject(ctx)
	warnings = append(warnings, catWarnings...)
	if data == nil {
		return status, warnings
	}

	status.found = true
	status.holderIdentity = data.HolderIdentity
	status.renewTime = data.RenewTime
	status.leaseTransitions = data.LeaseTransitions
	if !data.RenewTime.IsZero() {
		age, staleWarns := g.evaluateLockFreshness(data)
		status.renewAge = age
		warnings = append(warnings, staleWarns...)
	}

	return status, warnings
}

func (g *gcsLockInspector) findExtraLeaseKeys(ctx context.Context) ([]string, []string) {
	lsOut, lsErr := g.runner.output(ctx, "gcloud", "storage", "ls", fmt.Sprintf("gs://%s/leases/**", g.bucket))
	if lsErr != nil {
		if !isGCSObjectNotFoundError(lsErr) {
			return nil, []string{fmt.Sprintf("listing lease objects in gs://%s/leases/**: %v", g.bucket, lsErr)}
		}
		return nil, nil
	}

	var extraKeys []string
	prefix := fmt.Sprintf("gs://%s/", g.bucket)
	for _, line := range strings.Split(lsOut, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasSuffix(trimmed, "/") || strings.HasSuffix(trimmed, ":") {
			continue
		}
		key := strings.TrimPrefix(trimmed, prefix)
		if key != g.expectedLockKey {
			extraKeys = append(extraKeys, key)
		}
	}
	sort.Strings(extraKeys)

	var warnings []string
	if len(extraKeys) > 0 {
		warnings = append(warnings, fmt.Sprintf(
			"extra lease keys found in gs://%s: [%s] (expected only %q; clusters using different lease keys will elect multiple active leaders)",
			g.bucket, strings.Join(extraKeys, ", "), g.expectedLockKey,
		))
	}

	return extraKeys, warnings
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

func (g *gcsLockInspector) readLockObject(ctx context.Context) (*gcsLockData, []string) {
	catOut, catErr := g.runner.output(ctx, "gcloud", "storage", "cat", g.lockURI)
	if catErr != nil {
		if isGCSObjectNotFoundError(catErr) {
			return nil, []string{fmt.Sprintf("GCS lock object %q does not exist yet", g.lockURI)}
		}
		return nil, []string{fmt.Sprintf("reading GCS lock object %q: %v", g.lockURI, catErr)}
	}

	var data gcsLockData
	if err := json.Unmarshal([]byte(strings.TrimSpace(catOut)), &data); err != nil {
		return nil, []string{fmt.Sprintf("parsing GCS lock object %q JSON: %v", g.lockURI, err)}
	}

	return &data, nil
}

const (
	gcsDefaultLeaseDurationSeconds = int32(15)
	gcsStaleLeaseMultiplier        = 2
)

func (g *gcsLockInspector) evaluateLockFreshness(data *gcsLockData) (time.Duration, []string) {
	age := g.now.Sub(data.RenewTime).Truncate(time.Second)
	if age < 0 {
		age = 0
	}

	staleThreshold := time.Duration(gcsStaleLeaseMultiplier*gcsDefaultLeaseDurationSeconds) * time.Second
	if age > staleThreshold {
		return age, []string{fmt.Sprintf(
			"stale GCS lock %q: holder %q last renewed %s ago (at %s, staleness threshold %s)",
			g.lockURI, data.HolderIdentity, age, data.RenewTime.UTC().Format(time.RFC3339), staleThreshold,
		)}
	}

	return age, nil
}
