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
	"strings"
)

// gcpProvisioner provisions the GCS lease bucket, GCP service accounts, and IAM bindings required for multi-cluster HA.
type gcpProvisioner struct {
	runner    commandRunner
	bucketURI string
	dryRun    bool
}

func newGCPProvisioner(runner commandRunner, bucket string, dryRun bool) *gcpProvisioner {
	return &gcpProvisioner{
		runner:    runner,
		bucketURI: "gs://" + bucket,
		dryRun:    dryRun,
	}
}

func (p *gcpProvisioner) provision(ctx context.Context, clusters []gkeClusterInfo) error {
	if err := p.ensureLeaseBucket(ctx, clusters[0]); err != nil {
		return err
	}

	projects := dedupeProjects(clusters)
	for _, project := range projects {
		if err := p.configureProjectIAM(ctx, project.project, projects); err != nil {
			return err
		}
	}

	return nil
}

func (p *gcpProvisioner) ensureLeaseBucket(ctx context.Context, primary gkeClusterInfo) error {
	if !p.dryRun && p.runner.run(ctx, "gcloud", "storage", "buckets", "describe", p.bucketURI, "--project="+primary.project) == nil {
		return nil
	}

	location := gcsBucketLocation(primary.location)
	if err := p.runner.run(ctx, "gcloud", "storage", "buckets", "create", p.bucketURI, "--project="+primary.project, "--location="+location); err != nil {
		return fmt.Errorf("creating GCS lease bucket %q in location %q (project %q): %w", p.bucketURI, location, primary.project, err)
	}

	return nil
}

// gcsBucketLocation returns a valid GCS bucket region derived from a GKE cluster location.
// If the GKE cluster is zonal (e.g. "us-central1-a"), the zone suffix is stripped ("us-central1").
func gcsBucketLocation(location string) string {
	if isZonalGKELocation(location) {
		lastHyphen := strings.LastIndex(location, "-")
		return location[:lastHyphen]
	}

	return location
}

func isZonalGKELocation(location string) bool {
	if strings.Count(location, "-") != 2 {
		return false
	}

	lastHyphen := strings.LastIndex(location, "-")
	suffix := location[lastHyphen+1:]
	return len(suffix) == 1 && suffix[0] >= 'a' && suffix[0] <= 'z'
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

func (p *gcpProvisioner) configureProjectIAM(ctx context.Context, projectID string, allProjects []gkeClusterInfo) error {
	mclGSA := gsaEmail(gcpMCLServiceAccountName, projectID)
	syncerGSA := gsaEmail(gcpSyncerServiceAccountName, projectID)

	if err := p.ensureServiceAccounts(ctx, projectID, mclGSA, syncerGSA); err != nil {
		return err
	}

	if err := p.configureIAMForLeaseBucket(ctx, mclGSA); err != nil {
		return err
	}

	if err := p.configureIAMForWorkloadIdentity(ctx, projectID, mclGSA, syncerGSA); err != nil {
		return err
	}

	if err := p.configureIAMForServiceAccounts(ctx, allProjects, syncerGSA); err != nil {
		return err
	}

	return nil
}

func gsaEmail(serviceAccountName, projectID string) string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", serviceAccountName, projectID)
}

func (p *gcpProvisioner) ensureServiceAccounts(ctx context.Context, projectID, mclGSA, syncerGSA string) error {
	serviceAccounts := []struct {
		name  string
		email string
	}{
		{name: gcpMCLServiceAccountName, email: mclGSA},
		{name: gcpSyncerServiceAccountName, email: syncerGSA},
	}

	for _, sa := range serviceAccounts {
		if !p.dryRun && p.runner.run(ctx, "gcloud", "iam", "service-accounts", "describe", sa.email, "--project="+projectID) == nil {
			continue
		}

		if err := p.runner.run(ctx, "gcloud", "iam", "service-accounts", "create", sa.name, "--project="+projectID); err != nil {
			return fmt.Errorf("creating GCP service account %q in project %q: %w", sa.email, projectID, err)
		}
	}

	return nil
}

func (p *gcpProvisioner) configureIAMForLeaseBucket(ctx context.Context, mclGSA string) error {
	roles := []string{
		"roles/storage.objectAdmin",
		"roles/storage.legacyBucketReader",
	}
	for _, role := range roles {
		if err := p.runner.run(ctx, "gcloud", "storage", "buckets", "add-iam-policy-binding", p.bucketURI,
			"--member=serviceAccount:"+mclGSA,
			"--role="+role,
		); err != nil {
			return fmt.Errorf("granting %s on %q to %q: %w", role, p.bucketURI, mclGSA, err)
		}
	}

	return nil
}

func (p *gcpProvisioner) configureIAMForWorkloadIdentity(ctx context.Context, projectID, mclGSA, syncerGSA string) error {
	bindings := []struct {
		gsa string
		ksa string
	}{
		{gsa: mclGSA, ksa: gcpMCLServiceAccountName},
		{gsa: syncerGSA, ksa: gcpSyncerServiceAccountName},
	}

	for _, b := range bindings {
		wiMember := fmt.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]", projectID, gcpWorkloadIdentityNamespace, b.ksa)
		if err := p.runner.run(ctx, "gcloud", "iam", "service-accounts", "add-iam-policy-binding", b.gsa,
			"--project="+projectID,
			"--role=roles/iam.workloadIdentityUser",
			"--member="+wiMember,
		); err != nil {
			return fmt.Errorf("binding Workload Identity on %q for %q: %w", b.gsa, wiMember, err)
		}
	}

	return nil
}

func (p *gcpProvisioner) configureIAMForServiceAccounts(ctx context.Context, allProjects []gkeClusterInfo, syncerGSA string) error {
	for _, peer := range allProjects {
		if err := p.runner.run(ctx, "gcloud", "projects", "add-iam-policy-binding", peer.project,
			"--member=serviceAccount:"+syncerGSA,
			"--role=roles/container.viewer",
		); err != nil {
			return fmt.Errorf("granting container.viewer on project %q to %q: %w", peer.project, syncerGSA, err)
		}
	}

	return nil
}
