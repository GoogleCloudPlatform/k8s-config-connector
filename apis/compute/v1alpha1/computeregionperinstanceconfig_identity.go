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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeRegionPerInstanceConfigIdentity{}
	_ identity.Resource   = &ComputeRegionPerInstanceConfig{}
)

var ComputeRegionPerInstanceConfigIdentityFormat = gcpurls.Template[ComputeRegionPerInstanceConfigIdentity](
	"compute.googleapis.com",
	"projects/{project}/regions/{region}/instanceGroupManagers/{regionInstanceGroupManager}/{name}",
)

// ComputeRegionPerInstanceConfigIdentity is the identity of a GCP ComputeRegionPerInstanceConfig resource.
// +k8s:deepcopy-gen=false
type ComputeRegionPerInstanceConfigIdentity struct {
	Project                    string
	Region                     string
	RegionInstanceGroupManager string
	Name                       string
}

func (i *ComputeRegionPerInstanceConfigIdentity) String() string {
	return ComputeRegionPerInstanceConfigIdentityFormat.ToString(*i)
}

func (i *ComputeRegionPerInstanceConfigIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeRegionPerInstanceConfigIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeRegionPerInstanceConfig external=%q was not known (use %s): %w", ref, ComputeRegionPerInstanceConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeRegionPerInstanceConfig external=%q was not known (use %s)", ref, ComputeRegionPerInstanceConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeRegionPerInstanceConfigIdentity) Host() string {
	return ComputeRegionPerInstanceConfigIdentityFormat.Host()
}

func (i *ComputeRegionPerInstanceConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s/instanceGroupManagers/%s", i.Project, i.Region, i.RegionInstanceGroupManager)
}

func ParseComputeRegionPerInstanceConfigExternal(external string) (*ComputeRegionPerInstanceConfigIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeRegionPerInstanceConfig external value")
	}
	id := &ComputeRegionPerInstanceConfigIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func parseRegionInstanceGroupManager(ctx context.Context, reader client.Reader, obj *ComputeRegionPerInstanceConfig) (string, error) {
	ref := obj.Spec.RegionInstanceGroupManagerRef
	if ref.External != "" {
		external := apirefs.TrimComputeURIPrefix(ref.External)
		if strings.Contains(external, "instanceGroupManagers/") {
			parts := strings.Split(external, "instanceGroupManagers/")
			if len(parts) == 2 && parts[1] != "" {
				subparts := strings.Split(parts[1], "/")
				if subparts[0] != "" {
					return subparts[0], nil
				}
			}
			return "", fmt.Errorf("invalid ComputeRegionInstanceGroupManager external reference format: %q", ref.External)
		}
		if strings.Contains(external, "/") {
			return "", fmt.Errorf("invalid ComputeRegionInstanceGroupManager external reference format: %q", ref.External)
		}
		return external, nil
	}
	if ref.Name != "" {
		ns := ref.Namespace
		if ns == "" {
			ns = obj.Namespace
		}
		if reader != nil {
			igm := &computev1beta1.ComputeInstanceGroupManager{}
			if err := reader.Get(ctx, types.NamespacedName{Name: ref.Name, Namespace: ns}, igm); err != nil {
				return "", fmt.Errorf("cannot get referenced ComputeInstanceGroupManager %s/%s: %w", ns, ref.Name, err)
			}
			id, err := igm.GetIdentity(ctx, reader)
			if err != nil {
				return "", fmt.Errorf("cannot get identity of referenced ComputeInstanceGroupManager %s/%s: %w", ns, ref.Name, err)
			}
			if igmId, ok := id.(*computev1beta1.ComputeInstanceGroupManagerIdentity); ok && igmId.InstanceGroupManager != "" {
				return igmId.InstanceGroupManager, nil
			}
			resourceID, err := refs.GetResourceID(igm)
			if err != nil {
				return "", fmt.Errorf("cannot get resource ID of referenced ComputeInstanceGroupManager %s/%s: %w", ns, ref.Name, err)
			}
			if resourceID != "" {
				return resourceID, nil
			}
			return ref.Name, nil
		}
		return ref.Name, nil
	}
	return "", fmt.Errorf("cannot resolve regionInstanceGroupManager: neither external nor name is set")
}

func getIdentityFromComputeRegionPerInstanceConfigSpec(ctx context.Context, reader client.Reader, obj *ComputeRegionPerInstanceConfig) (*ComputeRegionPerInstanceConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	region := obj.Spec.Region
	if region == "" {
		return nil, fmt.Errorf("cannot resolve region: spec.region is empty")
	}

	igm, err := parseRegionInstanceGroupManager(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve regionInstanceGroupManager: %w", err)
	}

	identity := &ComputeRegionPerInstanceConfigIdentity{
		Project:                    projectID,
		Region:                     region,
		RegionInstanceGroupManager: igm,
		Name:                       resourceID,
	}
	return identity, nil
}

func (obj *ComputeRegionPerInstanceConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeRegionPerInstanceConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
