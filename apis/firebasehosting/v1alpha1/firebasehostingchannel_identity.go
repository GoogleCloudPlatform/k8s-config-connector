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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &FirebaseHostingChannelIdentity{}
var _ identity.Resource = &FirebaseHostingChannel{}

var FirebaseHostingChannelIdentityFormat = gcpurls.Template[FirebaseHostingChannelIdentity](
	"firebasehosting.googleapis.com",
	"sites/{site}/channels/{channel}",
)

// +k8s:deepcopy-gen=false
// FirebaseHostingChannelIdentity is the identity of a GCP FirebaseHostingChannel resource.
type FirebaseHostingChannelIdentity struct {
	Site    string
	Channel string
}

func (i *FirebaseHostingChannelIdentity) String() string {
	return FirebaseHostingChannelIdentityFormat.ToString(*i)
}

func (i *FirebaseHostingChannelIdentity) FromExternal(ref string) error {
	parsed, match, err := FirebaseHostingChannelIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of FirebaseHostingChannel external=%q was not known (use %s): %w", ref, FirebaseHostingChannelIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of FirebaseHostingChannel external=%q was not known (use %s)", ref, FirebaseHostingChannelIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirebaseHostingChannelIdentity) Host() string {
	return FirebaseHostingChannelIdentityFormat.Host()
}

func (i *FirebaseHostingChannelIdentity) ParentString() string {
	return "sites/" + i.Site
}

func getIdentityFromFirebaseHostingChannelSpec(ctx context.Context, reader client.Reader, obj *FirebaseHostingChannel) (*FirebaseHostingChannelIdentity, error) {
	siteID := obj.Spec.SiteID
	if siteID == "" {
		return nil, fmt.Errorf("spec.siteId is required")
	}

	channelID := common.ValueOf(obj.Spec.ResourceID)
	if channelID == "" {
		channelID = obj.GetName()
	}
	if channelID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &FirebaseHostingChannelIdentity{
		Site:    siteID,
		Channel: channelID,
	}, nil
}

func (obj *FirebaseHostingChannel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		id := &FirebaseHostingChannelIdentity{}
		if err := id.FromExternal(externalRef); err != nil {
			return nil, err
		}
		return id, nil
	}

	return getIdentityFromFirebaseHostingChannelSpec(ctx, reader, obj)
}
