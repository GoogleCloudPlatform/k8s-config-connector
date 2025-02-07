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

package v1alpha1


// +kcc:proto=google.cloud.binaryauthorization.v1.Attestor
type Attestor struct {
	// Required. The resource name, in the format:
	//  `projects/*/attestors/*`. This field may not be updated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Attestor.name
	Name *string `json:"name,omitempty"`

	// Optional. A descriptive comment.  This field may be updated.
	//  The field may be displayed in chooser dialogs.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Attestor.description
	Description *string `json:"description,omitempty"`

	// This specifies how an attestation will be read, and how it will be used
	//  during policy enforcement.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Attestor.user_owned_grafeas_note
	UserOwnedGrafeasNote *UserOwnedGrafeasNote `json:"userOwnedGrafeasNote,omitempty"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.AttestorPublicKey
type AttestorPublicKey struct {
	// Optional. A descriptive comment. This field may be updated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AttestorPublicKey.comment
	Comment *string `json:"comment,omitempty"`

	// The ID of this public key.
	//  Signatures verified by BinAuthz must include the ID of the public key that
	//  can be used to verify them, and that ID must match the contents of this
	//  field exactly.
	//  Additional restrictions on this field can be imposed based on which public
	//  key type is encapsulated. See the documentation on `public_key` cases below
	//  for details.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AttestorPublicKey.id
	ID *string `json:"id,omitempty"`

	// ASCII-armored representation of a PGP public key, as the entire output by
	//  the command `gpg --export --armor foo@example.com` (either LF or CRLF
	//  line endings).
	//  When using this field, `id` should be left blank.  The BinAuthz API
	//  handlers will calculate the ID and fill it in automatically.  BinAuthz
	//  computes this ID as the OpenPGP RFC4880 V4 fingerprint, represented as
	//  upper-case hex.  If `id` is provided by the caller, it will be
	//  overwritten by the API-calculated ID.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AttestorPublicKey.ascii_armored_pgp_public_key
	AsciiArmoredPgpPublicKey *string `json:"asciiArmoredPgpPublicKey,omitempty"`

	// A raw PKIX SubjectPublicKeyInfo format public key.
	//
	//  NOTE: `id` may be explicitly provided by the caller when using this
	//  type of public key, but it MUST be a valid RFC3986 URI. If `id` is left
	//  blank, a default one will be computed based on the digest of the DER
	//  encoding of the public key.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.AttestorPublicKey.pkix_public_key
	PkixPublicKey *PkixPublicKey `json:"pkixPublicKey,omitempty"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.PkixPublicKey
type PkixPublicKey struct {
	// A PEM-encoded public key, as described in
	//  https://tools.ietf.org/html/rfc7468#section-13
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.PkixPublicKey.public_key_pem
	PublicKeyPem *string `json:"publicKeyPem,omitempty"`

	// The signature algorithm used to verify a message against a signature using
	//  this key.
	//  These signature algorithm must match the structure and any object
	//  identifiers encoded in `public_key_pem` (i.e. this algorithm must match
	//  that of the public key).
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.PkixPublicKey.signature_algorithm
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.UserOwnedGrafeasNote
type UserOwnedGrafeasNote struct {
	// Required. The Grafeas resource name of a Attestation.Authority Note,
	//  created by the user, in the format: `projects/*/notes/*`. This field may
	//  not be updated.
	//
	//  An attestation by this attestor is stored as a Grafeas
	//  Attestation.Authority Occurrence that names a container image and that
	//  links to this Note. Grafeas is an external dependency.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.UserOwnedGrafeasNote.note_reference
	NoteReference *string `json:"noteReference,omitempty"`

	// Optional. Public keys that verify attestations signed by this
	//  attestor.  This field may be updated.
	//
	//  If this field is non-empty, one of the specified public keys must
	//  verify that an attestation was signed by this attestor for the
	//  image specified in the admission request.
	//
	//  If this field is empty, this attestor always returns that no
	//  valid attestations exist.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.UserOwnedGrafeasNote.public_keys
	PublicKeys []AttestorPublicKey `json:"publicKeys,omitempty"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.Attestor
type AttestorObservedState struct {
	// This specifies how an attestation will be read, and how it will be used
	//  during policy enforcement.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Attestor.user_owned_grafeas_note
	UserOwnedGrafeasNote *UserOwnedGrafeasNoteObservedState `json:"userOwnedGrafeasNote,omitempty"`

	// Output only. Time when the attestor was last updated.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.Attestor.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.binaryauthorization.v1.UserOwnedGrafeasNote
type UserOwnedGrafeasNoteObservedState struct {
	// Output only. This field will contain the service account email address
	//  that this Attestor will use as the principal when querying Container
	//  Analysis. Attestor administrators must grant this service account the
	//  IAM role needed to read attestations from the [note_reference][Note] in
	//  Container Analysis (`containeranalysis.notes.occurrences.viewer`).
	//
	//  This email address is fixed for the lifetime of the Attestor, but callers
	//  should not make any other assumptions about the service account email;
	//  future versions may use an email based on a different naming pattern.
	// +kcc:proto:field=google.cloud.binaryauthorization.v1.UserOwnedGrafeasNote.delegation_service_account_email
	DelegationServiceAccountEmail *string `json:"delegationServiceAccountEmail,omitempty"`
}
