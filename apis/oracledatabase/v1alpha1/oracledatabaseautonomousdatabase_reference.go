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

// --- OracleDatabaseAutonomousDatabaseRef ---

var OracleDatabaseAutonomousDatabaseRefGVK = schema.GroupVersionKind{
	Group:   "oracledatabase.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "OracleDatabaseAutonomousDatabase",
}

var _ refs.Ref = &OracleDatabaseAutonomousDatabaseRef{}

// OracleDatabaseAutonomousDatabaseRef is a reference to an OracleDatabaseAutonomousDatabase.
type OracleDatabaseAutonomousDatabaseRef struct {
	// A reference to an externally managed OracleDatabaseAutonomousDatabase resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/autonomousDatabases/{{autonomousDatabase}}".
	External string `json:"external,omitempty"`

	// The name of an OracleDatabaseAutonomousDatabase resource.
	Name string `json:"name,omitempty"`

	// The namespace of an OracleDatabaseAutonomousDatabase resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&OracleDatabaseAutonomousDatabaseRef{}, &OracleDatabaseAutonomousDatabase{})
}

func (r *OracleDatabaseAutonomousDatabaseRef) GetGVK() schema.GroupVersionKind {
	return OracleDatabaseAutonomousDatabaseRefGVK
}

func (r *OracleDatabaseAutonomousDatabaseRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *OracleDatabaseAutonomousDatabaseRef) GetExternal() string {
	return r.External
}

func (r *OracleDatabaseAutonomousDatabaseRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *OracleDatabaseAutonomousDatabaseRef) ValidateExternal(ref string) error {
	id := &OracleDatabaseAutonomousDatabaseIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *OracleDatabaseAutonomousDatabaseRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &OracleDatabaseAutonomousDatabaseIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *OracleDatabaseAutonomousDatabaseRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// --- OracleDatabaseAutonomousDatabaseBackupRef ---

var OracleDatabaseAutonomousDatabaseBackupRefGVK = schema.GroupVersionKind{
	Group:   "oracledatabase.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "OracleDatabaseAutonomousDatabaseBackup",
}

var _ refs.Ref = &OracleDatabaseAutonomousDatabaseBackupRef{}

// OracleDatabaseAutonomousDatabaseBackupRef is a reference to an OracleDatabaseAutonomousDatabaseBackup.
type OracleDatabaseAutonomousDatabaseBackupRef struct {
	// A reference to an externally managed OracleDatabaseAutonomousDatabaseBackup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/autonomousDatabaseBackups/{{autonomousDatabaseBackup}}".
	External string `json:"external,omitempty"`

	// The name of an OracleDatabaseAutonomousDatabaseBackup resource.
	Name string `json:"name,omitempty"`

	// The namespace of an OracleDatabaseAutonomousDatabaseBackup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&OracleDatabaseAutonomousDatabaseBackupRef{}, nil)
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) GetGVK() schema.GroupVersionKind {
	return OracleDatabaseAutonomousDatabaseBackupRefGVK
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) GetExternal() string {
	return r.External
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

var OracleDatabaseAutonomousDatabaseBackupIdentityFormat = gcpurls.Template[OracleDatabaseAutonomousDatabaseBackupIdentity]("oracledatabase.googleapis.com", "projects/{project}/locations/{location}/autonomousDatabaseBackups/{autonomousDatabaseBackup}")

// OracleDatabaseAutonomousDatabaseBackupIdentity is the identity of a GCP OracleDatabaseAutonomousDatabaseBackup resource.
// +k8s:deepcopy-gen=false
type OracleDatabaseAutonomousDatabaseBackupIdentity struct {
	Project                  string
	Location                 string
	AutonomousDatabaseBackup string
}

func (i *OracleDatabaseAutonomousDatabaseBackupIdentity) String() string {
	return OracleDatabaseAutonomousDatabaseBackupIdentityFormat.ToString(*i)
}

func (i *OracleDatabaseAutonomousDatabaseBackupIdentity) FromExternal(ref string) error {
	parsed, match, err := OracleDatabaseAutonomousDatabaseBackupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of OracleDatabaseAutonomousDatabaseBackup external=%q was not known (use %s): %w", ref, OracleDatabaseAutonomousDatabaseBackupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of OracleDatabaseAutonomousDatabaseBackup external=%q was not known (use %s)", ref, OracleDatabaseAutonomousDatabaseBackupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) ValidateExternal(ref string) error {
	id := &OracleDatabaseAutonomousDatabaseBackupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &OracleDatabaseAutonomousDatabaseBackupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *OracleDatabaseAutonomousDatabaseBackupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", OracleDatabaseAutonomousDatabaseBackupRefGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
