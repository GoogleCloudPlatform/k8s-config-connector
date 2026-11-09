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

// gcpProvisioner provisions the GCS lease bucket and IAM bindings required for multi-cluster HA.
type gcpProvisioner struct {
	runner    commandRunner
	bucketURI string
}

func newGCPProvisioner(runner commandRunner, bucket string) *gcpProvisioner {
	return &gcpProvisioner{
		runner:    runner,
		bucketURI: "gs://" + bucket,
	}
}

func (p *gcpProvisioner) provision(ctx context.Context, clusters []gkeClusterInfo) error {
	if err := p.createLeaseBucket(ctx, clusters[0]); err != nil {
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

func (p *gcpProvisioner) createLeaseBucket(ctx context.Context, primary gkeClusterInfo) error {
	location := gcsBucketLocation(primary.location)
	if err := p.runner.run(ctx, "gcloud", "storage", "buckets", "create", p.bucketURI, "--project="+primary.project, "--location="+location); err != nil {
		return fmt.Errorf("creating GCS lease bucket %q in location %q (project %q): %w", p.bucketURI, location, primary.project, err)
	}

	return nil
}

// gcsBucketLocation returns a GCS multi-region location derived from a GKE cluster location.
// Regional and zonal GKE locations are mapped to their corresponding GCS multi-region
// ("us", "eu", or "asia"), defaulting to "us" if no match is found. This ensures that
// the multi-cluster lease bucket remains accessible even during a regional outage.
func gcsBucketLocation(location string) string {
	loc := strings.ToLower(strings.TrimSpace(location))
	switch {
	case strings.HasPrefix(loc, "us-") || loc == "us":
		return "us"
	case strings.HasPrefix(loc, "europe-") || loc == "eu":
		return "eu"
	case strings.HasPrefix(loc, "asia-") || loc == "asia":
		return "asia"
	default:
		return "us"
	}
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
	if err := p.configureIAMForLeaseBucket(ctx, mclGSA); err != nil {
		return err
	}

	syncerGSA := gsaEmail(gcpSyncerServiceAccountName, projectID)

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
