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
)

// gcpProvisioner provisions the GCS lease bucket and IAM bindings required for multi-cluster HA.
type gcpProvisioner struct {
	runner         commandRunner
	bucketURI      string
	bucketLocation string
	outWriter      io.Writer
	warnWriter     io.Writer

	retryInterval time.Duration
	retryTimeout  time.Duration
}

const (
	defaultIAMRetryInterval = 5 * time.Second
	defaultIAMRetryTimeout  = 2 * time.Minute
)

func newGCPProvisioner(runner commandRunner, bucket, bucketLocation string, outWriter, warnWriter io.Writer) *gcpProvisioner {
	return &gcpProvisioner{
		runner:         runner,
		bucketURI:      "gs://" + bucket,
		bucketLocation: bucketLocation,
		outWriter:      outWriter,
		warnWriter:     warnWriter,
		retryInterval:  defaultIAMRetryInterval,
		retryTimeout:   defaultIAMRetryTimeout,
	}
}

func (p *gcpProvisioner) provision(ctx context.Context, clusters []gkeClusterInfo) error {
	if len(clusters) == 0 {
		return fmt.Errorf("no clusters provided to provision GCP resources")
	}

	if err := p.createLeaseBucket(ctx, clusters); err != nil {
		return err
	}

	projects := dedupeProjects(clusters)
	accountCreated, err := p.ensureServiceAccounts(ctx, projects)
	if err != nil {
		return err
	}

	for _, project := range projects {
		if err := p.configureProjectIAM(ctx, project.project, projects, accountCreated); err != nil {
			return err
		}
	}

	return nil
}

func (p *gcpProvisioner) ensureServiceAccounts(ctx context.Context, projects []gkeClusterInfo) (bool, error) {
	accountNames := []string{gcpMCLServiceAccountName, gcpSyncerServiceAccountName}
	anyCreated := false

	for _, project := range projects {
		for _, name := range accountNames {
			email := gsaEmail(name, project.project)
			_, describeErr := p.runner.output(ctx, "gcloud", "iam", "service-accounts", "describe", email, "--project="+project.project)
			if describeErr != nil {
				createErr := p.runner.run(ctx, "gcloud", "iam", "service-accounts", "create", name, "--project="+project.project)
				if createErr != nil {
					return false, fmt.Errorf("ensuring service account %q in project %q: describe failed: %w; create failed: %w", email, project.project, describeErr, createErr)
				}
				anyCreated = true
			}
		}
	}

	return anyCreated, nil
}

func (p *gcpProvisioner) createLeaseBucket(ctx context.Context, clusters []gkeClusterInfo) error {
	_, describeErr := p.runner.output(ctx, "gcloud", "storage", "buckets", "describe", p.bucketURI)
	if describeErr == nil {
		if p.outWriter != nil {
			fmt.Fprintf(p.outWriter, "Lease bucket %s already exists; reusing it.\n", p.bucketURI)
			if p.bucketLocation != "" {
				fmt.Fprintf(p.outWriter, "--%s has no effect on an existing bucket.\n", bucketLocationFlag)
			}
		}
		return nil
	}

	location, err := resolveBucketLocation(clusters, p.bucketLocation, p.warnWriter)
	if err != nil {
		return err
	}

	primary := clusters[0]
	if err := p.runner.run(ctx, "gcloud", "storage", "buckets", "create", p.bucketURI, "--project="+primary.project, "--location="+location); err != nil {
		return fmt.Errorf("creating GCS lease bucket %q in location %q (project %q): describe failed: %w; create failed: %w", p.bucketURI, location, primary.project, describeErr, err)
	}

	return nil
}

// resolveBucketLocation determines the GCS bucket location for the lease bucket.
// If explicitLocation is non-empty, it is used directly. Otherwise, it maps all
// cluster locations to candidate bucket locations. If all clusters agree, that
// location is used (with a warning printed if falling back to a regional bucket).
// If clusters differ, an error is returned instructing the operator to provide
// --bucket-location explicitly.
func resolveBucketLocation(clusters []gkeClusterInfo, explicitLocation string, warnWriter io.Writer) (string, error) {
	if explicitLocation != "" {
		return explicitLocation, nil
	}

	if len(clusters) == 0 {
		return "", fmt.Errorf("no clusters provided to determine lease bucket location")
	}

	type clusterResult struct {
		cluster        gkeClusterInfo
		bucketLocation string
		isMultiRegion  bool
	}

	results := make([]clusterResult, len(clusters))
	for i, c := range clusters {
		loc, isMulti := mapClusterToBucketLocation(c.location)
		results[i] = clusterResult{
			cluster:        c,
			bucketLocation: loc,
			isMultiRegion:  isMulti,
		}
	}

	first := results[0]
	if first.bucketLocation == "" {
		return "", fmt.Errorf("unable to determine bucket location from cluster location %q; specify an explicit GCS bucket location using --%s", first.cluster.location, bucketLocationFlag)
	}

	for _, r := range results[1:] {
		if r.bucketLocation != first.bucketLocation {
			var details []string
			for _, res := range results {
				details = append(details, fmt.Sprintf("cluster %q (location %q) maps to %q", res.cluster.contextName, res.cluster.location, res.bucketLocation))
			}
			return "", fmt.Errorf("clusters have differing bucket locations: %s; specify an explicit GCS bucket location using --%s (e.g. a multi-region, dual-region, or single region)", strings.Join(details, ", "), bucketLocationFlag)
		}
	}

	if !first.isMultiRegion && warnWriter != nil {
		fmt.Fprintf(warnWriter, "Warning: cluster location %q does not map to a GCS multi-region; using regional bucket location %q. For cross-region redundancy, specify --%s (e.g. a multi-region or dual-region).\n", first.cluster.location, first.bucketLocation, bucketLocationFlag)
	}

	return first.bucketLocation, nil
}

// mapClusterToBucketLocation maps a GKE cluster location (zonal or regional)
// to a candidate GCS bucket location. If the region is within a GCS multi-region
// ("us", "eu", or "asia"), the multi-region code is returned with isMultiRegion=true.
// Otherwise, it falls back to the regional location with isMultiRegion=false.
//
// Per GCS documentation (https://cloud.google.com/storage/docs/locations#location-mr):
// - The EU multi-region does not store data in europe-west2 (London) or europe-west6 (Zürich).
// - The ASIA multi-region excludes asia-east2 (Hong Kong) and asia-southeast2 (Jakarta).
func mapClusterToBucketLocation(location string) (string, bool) {
	loc := strings.ToLower(strings.TrimSpace(location))
	if loc == "" {
		return "", false
	}

	region := loc
	if isZonalGKELocation(loc) {
		lastHyphen := strings.LastIndex(loc, "-")
		region = loc[:lastHyphen]
	}

	switch {
	case region == "us" || strings.HasPrefix(region, "us-"):
		return "us", true
	case region == "eu":
		return "eu", true
	case strings.HasPrefix(region, "europe-"):
		if region == "europe-west2" || region == "europe-west6" {
			return region, false
		}
		return "eu", true
	case region == "asia":
		return "asia", true
	case strings.HasPrefix(region, "asia-"):
		if region == "asia-east2" || region == "asia-southeast2" {
			return region, false
		}
		return "asia", true
	default:
		return region, false
	}
}

func isZonalGKELocation(location string) bool {
	lastHyphen := strings.LastIndex(location, "-")
	if lastHyphen == -1 {
		return false
	}

	suffix := location[lastHyphen+1:]
	return len(suffix) == 1 && suffix[0] >= 'a' && suffix[0] <= 'z'
}

func gcsBucketLocation(location string) string {
	loc, _ := mapClusterToBucketLocation(location)
	return loc
}

func dedupeProjects(clusters []gkeClusterInfo) []gkeClusterInfo {
	seen := make(map[string]bool, len(clusters))
	unique := make([]gkeClusterInfo, 0, len(clusters))
	for _, c := range clusters {
		if !seen[c.project] {
			seen[c.project] = true
			unique = append(unique, c)
		}
	}

	return unique
}

const (
	gcpWorkloadIdentityNamespace = "cnrm-system"
	gcpMCLServiceAccountName     = "kcc-mcl"
	gcpSyncerServiceAccountName  = "kcc-syncer"
)

func (p *gcpProvisioner) runIAMBindingWithRetry(ctx context.Context, retry bool, name string, args ...string) error {
	err := p.runner.run(ctx, name, args...)
	if err == nil || !retry {
		return err
	}

	if p.outWriter != nil {
		fmt.Fprintln(p.outWriter, "Waiting for newly created service account to propagate in IAM...")
	}

	timeout := p.retryTimeout
	if timeout <= 0 {
		timeout = defaultIAMRetryTimeout
	}
	interval := p.retryInterval
	if interval <= 0 {
		interval = defaultIAMRetryInterval
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	lastErr := err
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			return lastErr
		case <-ticker.C:
			lastErr = p.runner.run(ctx, name, args...)
			if lastErr == nil {
				return nil
			}
		}
	}
}

func (p *gcpProvisioner) configureProjectIAM(ctx context.Context, projectID string, allProjects []gkeClusterInfo, retry bool) error {
	mclGSA := gsaEmail(gcpMCLServiceAccountName, projectID)
	if err := p.configureIAMForLeaseBucket(ctx, mclGSA, retry); err != nil {
		return err
	}

	syncerGSA := gsaEmail(gcpSyncerServiceAccountName, projectID)

	if err := p.configureIAMForWorkloadIdentity(ctx, projectID, mclGSA, syncerGSA, retry); err != nil {
		return err
	}

	if err := p.configureIAMForServiceAccounts(ctx, allProjects, syncerGSA, retry); err != nil {
		return err
	}

	return nil
}

func gsaEmail(serviceAccountName, projectID string) string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", serviceAccountName, projectID)
}

func (p *gcpProvisioner) configureIAMForLeaseBucket(ctx context.Context, mclGSA string, retry bool) error {
	roles := []string{
		"roles/storage.objectAdmin",
		"roles/storage.legacyBucketReader",
	}
	for _, role := range roles {
		if err := p.runIAMBindingWithRetry(ctx, retry, "gcloud", "storage", "buckets", "add-iam-policy-binding", p.bucketURI,
			"--member=serviceAccount:"+mclGSA,
			"--role="+role,
		); err != nil {
			return fmt.Errorf("granting %s on %q to %q: %w", role, p.bucketURI, mclGSA, err)
		}
	}

	return nil
}

func (p *gcpProvisioner) configureIAMForWorkloadIdentity(ctx context.Context, projectID, mclGSA, syncerGSA string, retry bool) error {
	bindings := []struct {
		gsa string
		ksa string
	}{
		{gsa: mclGSA, ksa: gcpMCLServiceAccountName},
		{gsa: syncerGSA, ksa: gcpSyncerServiceAccountName},
	}

	for _, b := range bindings {
		wiMember := fmt.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]", projectID, gcpWorkloadIdentityNamespace, b.ksa)
		if err := p.runIAMBindingWithRetry(ctx, retry, "gcloud", "iam", "service-accounts", "add-iam-policy-binding", b.gsa,
			"--project="+projectID,
			"--role=roles/iam.workloadIdentityUser",
			"--member="+wiMember,
		); err != nil {
			return fmt.Errorf("binding Workload Identity on %q for %q: %w", b.gsa, wiMember, err)
		}
	}
	return nil
}

// syncerClusterAccessRole lets a syncer find and connect to a peer cluster: it looks up the
// cluster's endpoint and CA through the GKE API (container.clusters.get), and the cluster's DNS
// endpoint requires container.clusters.connect.
//
// The role grants no access to the objects in the cluster. A syncer reads Config Connector
// resources through the cnrm-syncer-peers ClusterRoleBinding that setup creates on each cluster.
const syncerClusterAccessRole = "roles/container.clusterViewer"

func (p *gcpProvisioner) configureIAMForServiceAccounts(ctx context.Context, allProjects []gkeClusterInfo, syncerGSA string, retry bool) error {
	for _, peer := range allProjects {
		if err := p.runIAMBindingWithRetry(ctx, retry, "gcloud", "projects", "add-iam-policy-binding", peer.project,
			"--member=serviceAccount:"+syncerGSA,
			"--role="+syncerClusterAccessRole,
		); err != nil {
			return fmt.Errorf("granting %s on project %q to %q: %w", syncerClusterAccessRole, peer.project, syncerGSA, err)
		}
	}

	return nil
}
