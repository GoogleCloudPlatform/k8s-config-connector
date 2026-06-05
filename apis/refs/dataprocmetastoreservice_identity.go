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

package refs

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

var (
	_ identity.IdentityV2 = &DataprocMetastoreServiceIdentity{}
)

var DataprocMetastoreServiceIdentityFormat = gcpurls.Template[DataprocMetastoreServiceIdentity]("metastore.googleapis.com", "projects/{project}/locations/{location}/services/{service}")

// DataprocMetastoreServiceIdentity is the identity of a GCP DataprocMetastoreService resource.
// +k8s:deepcopy-gen=false
type DataprocMetastoreServiceIdentity struct {
	Project  string
	Location string
	Service  string
}

func (i *DataprocMetastoreServiceIdentity) String() string {
	return DataprocMetastoreServiceIdentityFormat.ToString(*i)
}

func (i *DataprocMetastoreServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := DataprocMetastoreServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataprocMetastoreService external=%q was not known (use %s): %w", ref, DataprocMetastoreServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataprocMetastoreService external=%q was not known (use %s)", ref, DataprocMetastoreServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataprocMetastoreServiceIdentity) Host() string {
	return DataprocMetastoreServiceIdentityFormat.Host()
}

// DataprocMetastoreService_IdentityFromSpec gets the identity of a DataprocMetastoreService from its spec.
func DataprocMetastoreService_IdentityFromSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataprocMetastoreServiceIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refsv1beta1.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &DataprocMetastoreServiceIdentity{
		Project:  projectID,
		Location: location,
		Service:  resourceID,
	}
	return identity, nil
}
