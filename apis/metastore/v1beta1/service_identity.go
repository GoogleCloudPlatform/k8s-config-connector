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
	_ identity.IdentityV2 = &MetastoreServiceIdentity{}
)

var MetastoreServiceIdentityFormat = gcpurls.Template[MetastoreServiceIdentity]("metastore.googleapis.com", "projects/{project}/locations/{location}/services/{service}")

// +k8s:deepcopy-gen=false
type MetastoreServiceIdentity struct {
	Project  string
	Location string
	Service  string
}

func (i *MetastoreServiceIdentity) String() string {
	return MetastoreServiceIdentityFormat.ToString(*i)
}

func (i *MetastoreServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := MetastoreServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MetastoreService external=%q was not known (use %s): %w", ref, MetastoreServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MetastoreService external=%q was not known (use %s)", ref, MetastoreServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MetastoreServiceIdentity) Host() string {
	return MetastoreServiceIdentityFormat.Host()
}

func getIdentityFromMetastoreServiceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*MetastoreServiceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &MetastoreServiceIdentity{
		Project:  projectID,
		Location: location,
		Service:  resourceID,
	}
	return identity, nil
}
