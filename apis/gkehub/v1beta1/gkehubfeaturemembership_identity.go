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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &GKEHubFeatureMembershipIdentity{}
	_ identity.Resource   = &GKEHubFeatureMembership{}

	GKEHubFeatureMembershipIdentityFormat = gcpurls.Template[GKEHubFeatureMembershipIdentity](
		"gkehub.googleapis.com",
		"projects/{project}/locations/{location}/memberships/{membership}/features/{feature}",
	)
)

// +k8s:deepcopy-gen=false
type GKEHubFeatureMembershipIdentity struct {
	Project         string
	Location        string // membership location
	Membership      string
	Feature         string
	FeatureLocation string // feature location
}

func (i *GKEHubFeatureMembershipIdentity) String() string {
	return GKEHubFeatureMembershipIdentityFormat.ToString(*i)
}

func (i *GKEHubFeatureMembershipIdentity) FromExternal(external string) error {
	out, match, err := GKEHubFeatureMembershipIdentityFormat.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubFeatureMembership external=%q was not known (use %s)", external, GKEHubFeatureMembershipIdentityFormat.CanonicalForm())
	}
	*i = *out
	return nil
}

func (i *GKEHubFeatureMembershipIdentity) Host() string {
	return GKEHubFeatureMembershipIdentityFormat.Host()
}

func getIdentityFromGKEHubFeatureMembershipSpec(ctx context.Context, reader client.Reader, obj client.Object) (*GKEHubFeatureMembershipIdentity, error) {
	typed, ok := obj.(*GKEHubFeatureMembership)
	if !ok {
		u := obj.(*unstructured.Unstructured)
		typed = &GKEHubFeatureMembership{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, typed); err != nil {
			return nil, fmt.Errorf("error converting to GKEHubFeatureMembership: %w", err)
		}
	}

	membershipID, err := ResolveGKEHubMembershipRef(ctx, reader, typed, &typed.Spec.MembershipRef)
	if err != nil {
		return nil, err
	}

	// For the feature ID, we need to extract the resource name (e.g. "configmanagement")
	featureName := ""
	featureLocation := ""
	if typed.Spec.FeatureRef.External != "" {
		tokens := strings.Split(typed.Spec.FeatureRef.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "features" {
			featureName = tokens[5]
			featureLocation = tokens[3]
		} else {
			return nil, fmt.Errorf("format of feature external=%q was not known", typed.Spec.FeatureRef.External)
		}
	} else {
		// Resolve from KCC
		key := types.NamespacedName{
			Name:      typed.Spec.FeatureRef.Name,
			Namespace: typed.Spec.FeatureRef.Namespace,
		}
		if key.Namespace == "" {
			key.Namespace = typed.GetNamespace()
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "gkehub.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "GKEHubFeature",
		})
		if err := reader.Get(ctx, key, u); err != nil {
			return nil, err
		}
		var err error
		featureName, _, err = unstructured.NestedString(u.Object, "spec", "resourceID")
		if err != nil {
			return nil, err
		}
		if featureName == "" {
			featureName = u.GetName()
		}
		featureLocation, _, _ = unstructured.NestedString(u.Object, "spec", "location")
	}

	if featureLocation == "" {
		featureLocation = typed.Spec.Location
	}
	if featureLocation == "" {
		featureLocation = "global"
	}

	return &GKEHubFeatureMembershipIdentity{
		Project:         membershipID.ProjectID,
		Location:        membershipID.Location,
		Membership:      membershipID.MembershipID,
		Feature:         featureName,
		FeatureLocation: featureLocation,
	}, nil
}

func (obj *GKEHubFeatureMembership) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromGKEHubFeatureMembershipSpec(ctx, reader, obj)
}
