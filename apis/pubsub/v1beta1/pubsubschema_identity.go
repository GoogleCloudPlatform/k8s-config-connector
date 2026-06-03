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
	_ identity.IdentityV2 = &PubSubSchemaIdentity{}
	_ identity.Resource   = &PubSubSchema{}
)

var PubSubSchemaIdentityFormat = gcpurls.Template[PubSubSchemaIdentity]("pubsub.googleapis.com", "projects/{project}/schemas/{schema}")

// +k8s:deepcopy-gen=false
type PubSubSchemaIdentity struct {
	Project string
	Schema  string
}

func (i *PubSubSchemaIdentity) String() string {
	return PubSubSchemaIdentityFormat.ToString(*i)
}

func (i *PubSubSchemaIdentity) FromExternal(ref string) error {
	parsed, match, err := PubSubSchemaIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PubSubSchema external=%q was not known (use %s): %w", ref, PubSubSchemaIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PubSubSchema external=%q was not known (use %s)", ref, PubSubSchemaIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PubSubSchemaIdentity) Host() string {
	return PubSubSchemaIdentityFormat.Host()
}

func getIdentityFromPubSubSchemaSpec(ctx context.Context, reader client.Reader, obj client.Object) (*PubSubSchemaIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &PubSubSchemaIdentity{
		Project: projectID,
		Schema:  resourceID,
	}
	return identity, nil
}

func (obj *PubSubSchema) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromPubSubSchemaSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
