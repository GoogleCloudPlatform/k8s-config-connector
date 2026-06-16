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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PubSubSubscriptionIdentity{}
	_ identity.Resource   = &PubSubSubscription{}
)

var PubSubSubscriptionIdentityFormat = gcpurls.Template[PubSubSubscriptionIdentity]("pubsub.googleapis.com", "projects/{project}/subscriptions/{subscription}")

// +k8s:deepcopy-gen=false

// PubSubSubscriptionIdentity is the identity of a Google Cloud PubSubSubscription resource.
type PubSubSubscriptionIdentity struct {
	Project      string
	Subscription string
}

func (i *PubSubSubscriptionIdentity) String() string {
	return PubSubSubscriptionIdentityFormat.ToString(*i)
}

func (i *PubSubSubscriptionIdentity) FromExternal(ref string) error {
	parsed, match, err := PubSubSubscriptionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PubSubSubscription external=%q was not known (use %s): %w", ref, PubSubSubscriptionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PubSubSubscription external=%q was not known (use %s)", ref, PubSubSubscriptionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PubSubSubscriptionIdentity) Host() string {
	return PubSubSubscriptionIdentityFormat.Host()
}

func (i *PubSubSubscriptionIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromPubSubSubscriptionSpec(ctx context.Context, reader client.Reader, obj *PubSubSubscription) (*PubSubSubscriptionIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &PubSubSubscriptionIdentity{
		Project:      projectID,
		Subscription: resourceID,
	}
	return identity, nil
}

func (obj *PubSubSubscription) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromPubSubSubscriptionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
