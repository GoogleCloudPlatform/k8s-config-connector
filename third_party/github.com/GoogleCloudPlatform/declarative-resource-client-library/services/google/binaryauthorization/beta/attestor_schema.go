// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLAttestorSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "BinaryAuthorization/Attestor",
			Description: "The BinaryAuthorization Attestor resource",
			StructName:  "Attestor",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Attestor",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "attestor",
						Required:    true,
						Description: "A full instance of a Attestor",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Attestor",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "attestor",
						Required:    true,
						Description: "A full instance of a Attestor",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Attestor",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "attestor",
						Required:    true,
						Description: "A full instance of a Attestor",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Attestor",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Attestor",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Attestor": &dcl.Component{
					Title:           "Attestor",
					ID:              "projects/{{project}}/attestors/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.",
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Time when the attestor was last updated.",
								Immutable:   true,
							},
							"userOwnedDrydockNote": &dcl.Property{
								Type:        "object",
								GoName:      "UserOwnedDrydockNote",
								GoType:      "AttestorUserOwnedDrydockNote",
								Description: "This specifies how an attestation will be read, and how it will be used during policy enforcement.",
								Required: []string{
									"noteReference",
								},
								Properties: map[string]*dcl.Property{
									"delegationServiceAccountEmail": &dcl.Property{
										Type:        "string",
										GoName:      "DelegationServiceAccountEmail",
										ReadOnly:    true,
										Description: "Output only. This field will contain the service account email address that this Attestor will use as the principal when querying Container Analysis. Attestor administrators must grant this service account the IAM role needed to read attestations from the in Container Analysis (`containeranalysis.notes.occurrences.viewer`). This email address is fixed for the lifetime of the Attestor, but callers should not make any other assumptions about the service account email; future versions may use an email based on a different naming pattern.",
										Immutable:   true,
									},
									"noteReference": &dcl.Property{
										Type:        "string",
										GoName:      "NoteReference",
										Description: "Required. The Drydock resource name of a Attestation. Authority Note, created by the user, in the format: `projects/*/notes/*`. This field may not be updated. An attestation by this attestor is stored as a Grafeas Attestation. Authority Occurrence that names a container image and that links to this Note. Grafeas is an external dependency.",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Containeranalysis/Note",
												Field:    "name",
											},
										},
									},
									"publicKeys": &dcl.Property{
										Type:        "array",
										GoName:      "PublicKeys",
										Description: "Optional. Public keys that verify attestations signed by this attestor. This field may be updated. If this field is non-empty, one of the specified public keys must verify that an attestation was signed by this attestor for the image specified in the admission request. If this field is empty, this attestor always returns that no valid attestations exist.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "AttestorUserOwnedDrydockNotePublicKeys",
											Properties: map[string]*dcl.Property{
												"asciiArmoredPgpPublicKey": &dcl.Property{
													Type:        "string",
													GoName:      "AsciiArmoredPgpPublicKey",
													Description: "ASCII-armored representation of a PGP public key, as the entire output by the command `gpg --export --armor foo@example.com` (either LF or CRLF line endings). When using this field, `id` should be left blank. The BinAuthz API handlers will calculate the ID and fill it in automatically. BinAuthz computes this ID as the OpenPGP RFC4880 V4 fingerprint, represented as upper-case hex. If `id` is provided by the caller, it will be overwritten by the API-calculated ID.",
												},
												"comment": &dcl.Property{
													Type:        "string",
													GoName:      "Comment",
													Description: "Optional. A descriptive comment. This field may be updated.",
												},
												"id": &dcl.Property{
													Type:          "string",
													GoName:        "Id",
													Description:   "The ID of this public key. Signatures verified by BinAuthz must include the ID of the public key that can be used to verify them, and that ID must match the contents of this field exactly. Additional restrictions on this field can be imposed based on which public key type is encapsulated. See the documentation on `public_key` cases below for details.",
													ServerDefault: true,
												},
												"pkixPublicKey": &dcl.Property{
													Type:        "object",
													GoName:      "PkixPublicKey",
													GoType:      "AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey",
													Description: "A raw PKIX SubjectPublicKeyInfo format public key. NOTE: `id` may be explicitly provided by the caller when using this type of public key, but it MUST be a valid RFC3986 URI. If `id` is left blank, a default one will be computed based on the digest of the DER encoding of the public key.",
													Properties: map[string]*dcl.Property{
														"publicKeyPem": &dcl.Property{
															Type:        "string",
															GoName:      "PublicKeyPem",
															Description: "A PEM-encoded public key, as described in https://tools.ietf.org/html/rfc7468#section-13",
														},
														"signatureAlgorithm": &dcl.Property{
															Type:        "string",
															GoName:      "SignatureAlgorithm",
															GoType:      "AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum",
															Description: "The signature algorithm used to verify a message against a signature using this key. These signature algorithm must match the structure and any object identifiers encoded in `public_key_pem` (i.e. this algorithm must match that of the public key). Possible values: SIGNATURE_ALGORITHM_UNSPECIFIED, RSA_PSS_2048_SHA256, RSA_PSS_3072_SHA256, RSA_PSS_4096_SHA256, RSA_PSS_4096_SHA512, RSA_SIGN_PKCS1_2048_SHA256, RSA_SIGN_PKCS1_3072_SHA256, RSA_SIGN_PKCS1_4096_SHA256, RSA_SIGN_PKCS1_4096_SHA512, ECDSA_P256_SHA256, EC_SIGN_P256_SHA256, ECDSA_P384_SHA384, EC_SIGN_P384_SHA384, ECDSA_P521_SHA512, EC_SIGN_P521_SHA512",
															Enum: []string{
																"SIGNATURE_ALGORITHM_UNSPECIFIED",
																"RSA_PSS_2048_SHA256",
																"RSA_PSS_3072_SHA256",
																"RSA_PSS_4096_SHA256",
																"RSA_PSS_4096_SHA512",
																"RSA_SIGN_PKCS1_2048_SHA256",
																"RSA_SIGN_PKCS1_3072_SHA256",
																"RSA_SIGN_PKCS1_4096_SHA256",
																"RSA_SIGN_PKCS1_4096_SHA512",
																"ECDSA_P256_SHA256",
																"EC_SIGN_P256_SHA256",
																"ECDSA_P384_SHA384",
																"EC_SIGN_P384_SHA384",
																"ECDSA_P521_SHA512",
																"EC_SIGN_P521_SHA512",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
