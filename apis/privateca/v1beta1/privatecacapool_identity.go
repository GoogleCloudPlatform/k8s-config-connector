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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &PrivateCACAPoolIdentity{}
	_ identity.Resource   = &PrivateCACAPool{}
)

var PrivateCACAPoolIdentityFormat = gcpurls.Template[PrivateCACAPoolIdentity]("privateca.googleapis.com", "projects/{project}/locations/{location}/caPools/{caPool}")

// +k8s:deepcopy-gen=false
type PrivateCACAPoolIdentity struct {
	Project  string
	Location string
	CAPool   string
}

func (i *PrivateCACAPoolIdentity) String() string {
	return PrivateCACAPoolIdentityFormat.ToString(*i)
}

func (i *PrivateCACAPoolIdentity) FromExternal(ref string) error {
	parsed, match, err := PrivateCACAPoolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of PrivateCACAPool external=%q was not known (use %s): %w", ref, PrivateCACAPoolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of PrivateCACAPool external=%q was not known (use %s)", ref, PrivateCACAPoolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *PrivateCACAPoolIdentity) Host() string {
	return PrivateCACAPoolIdentityFormat.Host()
}

func getIdentityFromPrivateCACAPoolSpec(ctx context.Context, reader client.Reader, obj client.Object) (*PrivateCACAPoolIdentity, error) {
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

	identity := &PrivateCACAPoolIdentity{
		Project:  projectID,
		Location: location,
		CAPool:   resourceID,
	}
	return identity, nil
}

func (obj *PrivateCACAPool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromPrivateCACAPoolSpec(ctx, reader, obj)
}
