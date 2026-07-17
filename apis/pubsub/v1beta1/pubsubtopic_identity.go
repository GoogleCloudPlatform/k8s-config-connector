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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PubSubTopicIdentity{}
	_ identity.Resource   = &PubSubTopic{}
)

var PubSubTopicIdentityFormat = gcpurls.Template[PubSubTopicIdentity]("pubsub.googleapis.com", "projects/{project}/topics/{topic}")

// +k8s:deepcopy-gen=false

// PubSubTopicIdentity is the identity of a GCP PubSubTopic resource.
type PubSubTopicIdentity struct {
	Project string
	Topic   string
}

func (i *PubSubTopicIdentity) String() string {
	return PubSubTopicIdentityFormat.ToString(*i)
}

func (i *PubSubTopicIdentity) FromExternal(ref string) error {
	parsed, match, err := PubSubTopicIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PubSubTopic external=%q was not known (use %s): %w", ref, PubSubTopicIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PubSubTopic external=%q was not known (use %s)", ref, PubSubTopicIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PubSubTopicIdentity) Host() string {
	return PubSubTopicIdentityFormat.Host()
}

func getIdentityFromPubSubTopicSpec(ctx context.Context, reader client.Reader, obj *PubSubTopic) (*PubSubTopicIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &PubSubTopicIdentity{
		Project: projectID,
		Topic:   resourceID,
	}
	return identity, nil
}

func (obj *PubSubTopic) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromPubSubTopicSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
