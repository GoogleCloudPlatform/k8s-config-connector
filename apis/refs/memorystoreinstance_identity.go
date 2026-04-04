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
	_ identity.IdentityV2 = &MemorystoreInstanceIdentity{}
)

var MemorystoreInstanceIdentityFormat = gcpurls.Template[MemorystoreInstanceIdentity]("memorystore.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// MemorystoreInstanceIdentity defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type MemorystoreInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *MemorystoreInstanceIdentity) String() string {
	return MemorystoreInstanceIdentityFormat.ToString(*i)
}

func (i *MemorystoreInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := MemorystoreInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MemorystoreInstance external=%q was not known (use %s): %w", ref, MemorystoreInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MemorystoreInstance external=%q was not known (use %s)", ref, MemorystoreInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MemorystoreInstanceIdentity) Host() string {
	return MemorystoreInstanceIdentityFormat.Host()
}

func (i *MemorystoreInstanceIdentity) ID() string {
	return i.Instance
}

func (i *MemorystoreInstanceIdentity) Parent() *MemorystoreInstanceParent {
	return &MemorystoreInstanceParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

type MemorystoreInstanceParent struct {
	ProjectID string
	Location  string
}

func (p *MemorystoreInstanceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

func ParseInstanceExternal(external string) (parent *MemorystoreInstanceParent, resourceID string, err error) {
	id := &MemorystoreInstanceIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.Parent(), id.ID(), nil
}

// MemorystoreInstance_IdentityFromSpec gets the identity of a MemorystoreInstance from its spec.
// We could have a registry mapping GVK to the type, once all the types implemented identity.Resource,
// then we could move this helper into the resource type.
func MemorystoreInstance_IdentityFromSpec(ctx context.Context, reader client.Reader, obj client.Object) (*MemorystoreInstanceIdentity, error) {
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

	identity := &MemorystoreInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}
	return identity, nil
}