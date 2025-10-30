// Copyright 2025 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	runv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ProjectPrefix = "//cloudresourcemanager.googleapis.com/projects"
	// OrganizationPrefix = "//cloudresourcemanager.googleapis.com/organizations"
	BucketPrefix = "//storage.googleapis.com/projects/_/buckets"
	RunJobPrefix = "//run.googleapis.com/"
)

var _ identity.Identity = &TagBindingParent{}

type TagBindingParent struct {
	parent identity.Identity
}

func (p *TagBindingParent) String() string {
	var fullResourceName string
	switch p := p.parent.(type) {
	case *refsv1beta1.Project:
		fullResourceName = ProjectPrefix + "/" + p.ProjectID
	case *storagev1beta1.StorageBucketIdentity:
		// TagBindings on buckets have a special parent format with `projects/_`.
		// The identity's ID() method returns just the bucket name.
		fullResourceName = BucketPrefix + "/" + p.ID()
	case *runv1beta1.JobIdentity:
		fullResourceName = RunJobPrefix + p.String()
	default:
		fullResourceName = ""
	}

	return fullResourceName
}

func (p *TagBindingParent) FromExternal(ref string) error {
	switch {
	case strings.HasPrefix(ref, ProjectPrefix):
		r := strings.TrimPrefix(ref, ProjectPrefix+"/")

		p.parent = &refsv1beta1.Project{}
		return p.parent.FromExternal(r)
	case strings.HasPrefix(ref, BucketPrefix):
		r := strings.TrimPrefix(ref, BucketPrefix+"/")

		p.parent = &storagev1beta1.StorageBucketIdentity{}
		return p.parent.FromExternal(r)
	case strings.HasPrefix(ref, RunJobPrefix):
		r := strings.TrimPrefix(ref, RunJobPrefix+"/")

		p.parent = &runv1beta1.JobIdentity{}
		return p.parent.FromExternal(r)
	}
	return fmt.Errorf("unknown parent format for %q", ref)
}

func (r *ParentRef) Normalize(ctx context.Context, reader client.Reader, namespace string) error {
	// Validate the full resource name.
	if r.External != "" {
		parent := &TagBindingParent{}
		return parent.FromExternal(r.External)
	}

	if r.RunJobRef != nil {
		external, err := r.RunJobRef.NormalizedExternal(ctx, reader, namespace)
		if err != nil {
			return err
		}
		runjobIdentity := &runv1beta1.JobIdentity{}
		if err := runjobIdentity.FromExternal(external); err != nil {
			return err
		}
		r.External = RunJobPrefix + "/" + external
	}

	if r.StorageBucketRef != nil {
		external, err := r.StorageBucketRef.NormalizedExternal(ctx, reader, namespace)
		if err != nil {
			return err
		}
		storagebucketIdentity := &storagev1beta1.StorageBucketIdentity{}
		if err := storagebucketIdentity.FromExternal(external); err != nil {
			return err
		}
		r.External = BucketPrefix + "/" + external
	}

	if r.Name != "" {
		projectRef := &refsv1beta1.ProjectRef{
			Name:      r.Name,
			Namespace: r.Namespace,
		}
		project, err := refsv1beta1.ResolveProject(ctx, reader, namespace, projectRef)
		if err != nil {
			return err
		}
		r.External = ProjectPrefix + "/" + project.String()
	}

	// TODO: support organizationRef.
	return nil
}
