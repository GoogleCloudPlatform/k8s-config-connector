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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// --- OracleDatabaseODBNetworkRef ---

var OracleDatabaseODBNetworkRefGVK = schema.GroupVersionKind{
	Group:   "oracledatabase.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "OracleDatabaseODBNetwork",
}

var _ refs.Ref = &OracleDatabaseODBNetworkRef{}

// OracleDatabaseODBNetworkRef is a reference to an OracleDatabaseODBNetwork.
type OracleDatabaseODBNetworkRef struct {
	// A reference to an externally managed OracleDatabaseODBNetwork resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/odbNetworks/{{odbNetwork}}".
	External string `json:"external,omitempty"`

	// The name of an OracleDatabaseODBNetwork resource.
	Name string `json:"name,omitempty"`

	// The namespace of an OracleDatabaseODBNetwork resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&OracleDatabaseODBNetworkRef{}, nil)
}

func (r *OracleDatabaseODBNetworkRef) GetGVK() schema.GroupVersionKind {
	return OracleDatabaseODBNetworkRefGVK
}

func (r *OracleDatabaseODBNetworkRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *OracleDatabaseODBNetworkRef) GetExternal() string {
	return r.External
}

func (r *OracleDatabaseODBNetworkRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

var OracleDatabaseODBNetworkIdentityFormat = gcpurls.Template[OracleDatabaseODBNetworkIdentity]("oracledatabase.googleapis.com", "projects/{project}/locations/{location}/odbNetworks/{odbNetwork}")

// +k8s:deepcopy-gen=false
type OracleDatabaseODBNetworkIdentity struct {
	Project    string
	Location   string
	OdbNetwork string
}

func (i *OracleDatabaseODBNetworkIdentity) String() string {
	return OracleDatabaseODBNetworkIdentityFormat.ToString(*i)
}

func (i *OracleDatabaseODBNetworkIdentity) FromExternal(ref string) error {
	parsed, match, err := OracleDatabaseODBNetworkIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of OracleDatabaseODBNetwork external=%q was not known (use %s): %w", ref, OracleDatabaseODBNetworkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of OracleDatabaseODBNetwork external=%q was not known (use %s)", ref, OracleDatabaseODBNetworkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *OracleDatabaseODBNetworkRef) ValidateExternal(ref string) error {
	id := &OracleDatabaseODBNetworkIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *OracleDatabaseODBNetworkRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &OracleDatabaseODBNetworkIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *OracleDatabaseODBNetworkRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// --- OracleDatabaseODBSubnetRef ---

var OracleDatabaseODBSubnetRefGVK = schema.GroupVersionKind{
	Group:   "oracledatabase.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "OracleDatabaseODBSubnet",
}

var _ refs.Ref = &OracleDatabaseODBSubnetRef{}

// OracleDatabaseODBSubnetRef is a reference to an OracleDatabaseODBSubnet.
type OracleDatabaseODBSubnetRef struct {
	// A reference to an externally managed OracleDatabaseODBSubnet resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/odbNetworks/{{odbNetwork}}/odbSubnets/{{odbSubnet}}".
	External string `json:"external,omitempty"`

	// The name of an OracleDatabaseODBSubnet resource.
	Name string `json:"name,omitempty"`

	// The namespace of an OracleDatabaseODBSubnet resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&OracleDatabaseODBSubnetRef{}, nil)
}

func (r *OracleDatabaseODBSubnetRef) GetGVK() schema.GroupVersionKind {
	return OracleDatabaseODBSubnetRefGVK
}

func (r *OracleDatabaseODBSubnetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *OracleDatabaseODBSubnetRef) GetExternal() string {
	return r.External
}

func (r *OracleDatabaseODBSubnetRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

var OracleDatabaseODBSubnetIdentityFormat = gcpurls.Template[OracleDatabaseODBSubnetIdentity]("oracledatabase.googleapis.com", "projects/{project}/locations/{location}/odbNetworks/{odbNetwork}/odbSubnets/{odbSubnet}")

// +k8s:deepcopy-gen=false
type OracleDatabaseODBSubnetIdentity struct {
	Project    string
	Location   string
	OdbNetwork string
	OdbSubnet  string
}

func (i *OracleDatabaseODBSubnetIdentity) String() string {
	return OracleDatabaseODBSubnetIdentityFormat.ToString(*i)
}

func (i *OracleDatabaseODBSubnetIdentity) FromExternal(ref string) error {
	parsed, match, err := OracleDatabaseODBSubnetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of OracleDatabaseODBSubnet external=%q was not known (use %s): %w", ref, OracleDatabaseODBSubnetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of OracleDatabaseODBSubnet external=%q was not known (use %s)", ref, OracleDatabaseODBSubnetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *OracleDatabaseODBSubnetRef) ValidateExternal(ref string) error {
	id := &OracleDatabaseODBSubnetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *OracleDatabaseODBSubnetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &OracleDatabaseODBSubnetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *OracleDatabaseODBSubnetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// --- OracleDatabaseExascaleDBStorageVaultRef ---

var OracleDatabaseExascaleDBStorageVaultRefGVK = schema.GroupVersionKind{
	Group:   "oracledatabase.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "OracleDatabaseExascaleDBStorageVault",
}

var _ refs.Ref = &OracleDatabaseExascaleDBStorageVaultRef{}

// OracleDatabaseExascaleDBStorageVaultRef is a reference to an OracleDatabaseExascaleDBStorageVault.
type OracleDatabaseExascaleDBStorageVaultRef struct {
	// A reference to an externally managed OracleDatabaseExascaleDBStorageVault resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/exascaleDbStorageVaults/{{exascaleDbStorageVault}}".
	External string `json:"external,omitempty"`

	// The name of an OracleDatabaseExascaleDBStorageVault resource.
	Name string `json:"name,omitempty"`

	// The namespace of an OracleDatabaseExascaleDBStorageVault resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&OracleDatabaseExascaleDBStorageVaultRef{}, nil)
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) GetGVK() schema.GroupVersionKind {
	return OracleDatabaseExascaleDBStorageVaultRefGVK
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) GetExternal() string {
	return r.External
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

var OracleDatabaseExascaleDBStorageVaultIdentityFormat = gcpurls.Template[OracleDatabaseExascaleDBStorageVaultIdentity]("oracledatabase.googleapis.com", "projects/{project}/locations/{location}/exascaleDbStorageVaults/{exascaleDbStorageVault}")

// +k8s:deepcopy-gen=false
type OracleDatabaseExascaleDBStorageVaultIdentity struct {
	Project                string
	Location               string
	ExascaleDbStorageVault string
}

func (i *OracleDatabaseExascaleDBStorageVaultIdentity) String() string {
	return OracleDatabaseExascaleDBStorageVaultIdentityFormat.ToString(*i)
}

func (i *OracleDatabaseExascaleDBStorageVaultIdentity) FromExternal(ref string) error {
	parsed, match, err := OracleDatabaseExascaleDBStorageVaultIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of OracleDatabaseExascaleDBStorageVault external=%q was not known (use %s): %w", ref, OracleDatabaseExascaleDBStorageVaultIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of OracleDatabaseExascaleDBStorageVault external=%q was not known (use %s)", ref, OracleDatabaseExascaleDBStorageVaultIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) ValidateExternal(ref string) error {
	id := &OracleDatabaseExascaleDBStorageVaultIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &OracleDatabaseExascaleDBStorageVaultIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *OracleDatabaseExascaleDBStorageVaultRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
