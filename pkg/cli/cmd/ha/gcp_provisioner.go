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
		if err := p.configureProjectIAM(ctx, project, projects); err != nil {
			return err
		}
	}

	return nil
}

func (p *gcpProvisioner) createLeaseBucket(ctx context.Context, primary gkeClusterInfo) error {
	location := primary.gcsBucketLocation()
	if err := p.runner.run(ctx, "gcloud", "storage", "buckets", "create", p.bucketURI, "--project="+primary.project, "--location="+location); err != nil {
		return fmt.Errorf("creating GCS lease bucket %q in location %q (project %q): %w", p.bucketURI, location, primary.project, err)
	}

	return nil
}

func (p *gcpProvisioner) configureProjectIAM(ctx context.Context, project gkeClusterInfo, allProjects []gkeClusterInfo) error {
	mclGSA := project.mclGSAEmail()
	if err := p.runner.run(ctx, "gcloud", "storage", "buckets", "add-iam-policy-binding", p.bucketURI,
		"--member=serviceAccount:"+mclGSA,
		"--role=roles/storage.objectAdmin",
	); err != nil {
		return fmt.Errorf("granting storage.objectAdmin on %q to %q: %w", p.bucketURI, mclGSA, err)
	}

	bindings := []struct {
		gsa string
		ksa string
	}{
		{gsa: mclGSA, ksa: mclServiceAccountName},
		{gsa: project.syncerGSAEmail(), ksa: syncerServiceAccountName},
	}
	for _, b := range bindings {
		wiMember := fmt.Sprintf("serviceAccount:%s.svc.id.goog[%s/%s]", project.project, cnrmSystemNamespace, b.ksa)
		if err := p.runner.run(ctx, "gcloud", "iam", "service-accounts", "add-iam-policy-binding", b.gsa,
			"--project="+project.project,
			"--role=roles/iam.workloadIdentityUser",
			"--member="+wiMember,
		); err != nil {
			return fmt.Errorf("binding Workload Identity on %q for %q: %w", b.gsa, wiMember, err)
		}
	}

	syncerGSA := project.syncerGSAEmail()
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
