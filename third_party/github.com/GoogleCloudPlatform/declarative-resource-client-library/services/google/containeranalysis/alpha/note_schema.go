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
package alpha

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLNoteSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "ContainerAnalysis/Note",
			Description: "The ContainerAnalysis Note resource",
			StructName:  "Note",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Note",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "note",
						Required:    true,
						Description: "A full instance of a Note",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Note",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "note",
						Required:    true,
						Description: "A full instance of a Note",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Note",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "note",
						Required:    true,
						Description: "A full instance of a Note",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Note",
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
				Description: "The function used to list information about many Note",
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
				"Note": &dcl.Component{
					Title:           "Note",
					ID:              "projects/{{project}}/notes/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"attestation": &dcl.Property{
								Type:        "object",
								GoName:      "Attestation",
								GoType:      "NoteAttestation",
								Description: "A note describing an attestation role.",
								Conflicts: []string{
									"vulnerability",
									"build",
									"image",
									"package",
									"deployment",
									"discovery",
								},
								Properties: map[string]*dcl.Property{
									"hint": &dcl.Property{
										Type:        "object",
										GoName:      "Hint",
										GoType:      "NoteAttestationHint",
										Description: "Hint hints at the purpose of the attestation authority.",
										Required: []string{
											"humanReadableName",
										},
										Properties: map[string]*dcl.Property{
											"humanReadableName": &dcl.Property{
												Type:        "string",
												GoName:      "HumanReadableName",
												Description: "Required. The human readable name of this attestation authority, for example \"qa\".",
											},
										},
									},
								},
							},
							"build": &dcl.Property{
								Type:        "object",
								GoName:      "Build",
								GoType:      "NoteBuild",
								Description: "A note describing build provenance for a verifiable build.",
								Conflicts: []string{
									"vulnerability",
									"image",
									"package",
									"deployment",
									"discovery",
									"attestation",
								},
								Required: []string{
									"builderVersion",
								},
								Properties: map[string]*dcl.Property{
									"builderVersion": &dcl.Property{
										Type:        "string",
										GoName:      "BuilderVersion",
										Description: "Required. Immutable. Version of the builder which produced this build.",
									},
									"signature": &dcl.Property{
										Type:        "object",
										GoName:      "Signature",
										GoType:      "NoteBuildSignature",
										Description: "Signature of the build in occurrences pointing to this build note containing build details.",
										Required: []string{
											"signature",
										},
										Properties: map[string]*dcl.Property{
											"keyId": &dcl.Property{
												Type:        "string",
												GoName:      "KeyId",
												Description: "An ID for the key used to sign. This could be either an ID for the key stored in `public_key` (such as the ID or fingerprint for a PGP key, or the CN for a cert), or a reference to an external key (such as a reference to a key in Cloud Key Management Service).",
											},
											"keyType": &dcl.Property{
												Type:        "string",
												GoName:      "KeyType",
												GoType:      "NoteBuildSignatureKeyTypeEnum",
												Description: "The type of the key, either stored in `public_key` or referenced in `key_id`. Possible values: KEY_TYPE_UNSPECIFIED, PGP_ASCII_ARMORED, PKIX_PEM",
												Enum: []string{
													"KEY_TYPE_UNSPECIFIED",
													"PGP_ASCII_ARMORED",
													"PKIX_PEM",
												},
											},
											"publicKey": &dcl.Property{
												Type:        "string",
												GoName:      "PublicKey",
												Description: "Public key of the builder which can be used to verify that the related findings are valid and unchanged. If `key_type` is empty, this defaults to PEM encoded public keys. This field may be empty if `key_id` references an external key. For Cloud Build based signatures, this is a PEM encoded public key. To verify the Cloud Build signature, place the contents of this field into a file (public.pem). The signature field is base64-decoded into its binary representation in signature.bin, and the provenance bytes from `BuildDetails` are base64-decoded into a binary representation in signed.bin. OpenSSL can then verify the signature: `openssl sha256 -verify public.pem -signature signature.bin signed.bin`",
											},
											"signature": &dcl.Property{
												Type:        "string",
												GoName:      "Signature",
												Description: "Required. Signature of the related `BuildProvenance`. In JSON, this is base-64 encoded.",
											},
										},
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time this note was created. This field can be used as a filter in list requests.",
								Immutable:   true,
							},
							"deployment": &dcl.Property{
								Type:        "object",
								GoName:      "Deployment",
								GoType:      "NoteDeployment",
								Description: "A note describing something that can be deployed.",
								Conflicts: []string{
									"vulnerability",
									"build",
									"image",
									"package",
									"discovery",
									"attestation",
								},
								Required: []string{
									"resourceUri",
								},
								Properties: map[string]*dcl.Property{
									"resourceUri": &dcl.Property{
										Type:        "array",
										GoName:      "ResourceUri",
										Description: "Required. Resource URI for the artifact being deployed.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
								},
							},
							"discovery": &dcl.Property{
								Type:        "object",
								GoName:      "Discovery",
								GoType:      "NoteDiscovery",
								Description: "A note describing the initial analysis of a resource.",
								Conflicts: []string{
									"vulnerability",
									"build",
									"image",
									"package",
									"deployment",
									"attestation",
								},
								Required: []string{
									"analysisKind",
								},
								Properties: map[string]*dcl.Property{
									"analysisKind": &dcl.Property{
										Type:        "string",
										GoName:      "AnalysisKind",
										GoType:      "NoteDiscoveryAnalysisKindEnum",
										Description: "The kind of analysis that is handled by this discovery. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE",
										Enum: []string{
											"NOTE_KIND_UNSPECIFIED",
											"VULNERABILITY",
											"BUILD",
											"IMAGE",
											"PACKAGE",
											"DEPLOYMENT",
											"DISCOVERY",
											"ATTESTATION",
											"UPGRADE",
										},
									},
								},
							},
							"expirationTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "ExpirationTime",
								Description: "Time of expiration for this note. Empty if note does not expire.",
							},
							"image": &dcl.Property{
								Type:        "object",
								GoName:      "Image",
								GoType:      "NoteImage",
								Description: "A note describing a base image.",
								Conflicts: []string{
									"vulnerability",
									"build",
									"package",
									"deployment",
									"discovery",
									"attestation",
								},
								Required: []string{
									"resourceUrl",
									"fingerprint",
								},
								Properties: map[string]*dcl.Property{
									"fingerprint": &dcl.Property{
										Type:        "object",
										GoName:      "Fingerprint",
										GoType:      "NoteImageFingerprint",
										Description: "Required. Immutable. The fingerprint of the base image.",
										Required: []string{
											"v1Name",
											"v2Blob",
										},
										Properties: map[string]*dcl.Property{
											"v1Name": &dcl.Property{
												Type:        "string",
												GoName:      "V1Name",
												Description: "Required. The layer ID of the final layer in the Docker image's v1 representation.",
											},
											"v2Blob": &dcl.Property{
												Type:        "array",
												GoName:      "V2Blob",
												Description: "Required. The ordered list of v2 blobs that represent a given image.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"v2Name": &dcl.Property{
												Type:        "string",
												GoName:      "V2Name",
												ReadOnly:    true,
												Description: "Output only. The name of the image's v2 blobs computed via: ) Only the name of the final blob is kept.",
											},
										},
									},
									"resourceUrl": &dcl.Property{
										Type:        "string",
										GoName:      "ResourceUrl",
										Description: "Required. Immutable. The resource_url for the resource representing the basis of associated occurrence images.",
									},
								},
							},
							"longDescription": &dcl.Property{
								Type:        "string",
								GoName:      "LongDescription",
								Description: "A detailed description of this note.",
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. Immutable. The name of the package.",
								Immutable:   true,
							},
							"package": &dcl.Property{
								Type:        "object",
								GoName:      "Package",
								GoType:      "NotePackage",
								Description: "Required for non-Windows OS. The package this Upgrade is for.",
								Conflicts: []string{
									"vulnerability",
									"build",
									"image",
									"deployment",
									"discovery",
									"attestation",
								},
								Required: []string{
									"name",
								},
								Properties: map[string]*dcl.Property{
									"distribution": &dcl.Property{
										Type:        "array",
										GoName:      "Distribution",
										Description: "The various channels by which a package is distributed.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "NotePackageDistribution",
											Required: []string{
												"cpeUri",
											},
											Properties: map[string]*dcl.Property{
												"architecture": &dcl.Property{
													Type:        "string",
													GoName:      "Architecture",
													GoType:      "NotePackageDistributionArchitectureEnum",
													Description: "The CPU architecture for which packages in this distribution channel were built Possible values: ARCHITECTURE_UNSPECIFIED, X86, X64",
													Enum: []string{
														"ARCHITECTURE_UNSPECIFIED",
														"X86",
														"X64",
													},
												},
												"cpeUri": &dcl.Property{
													Type:        "string",
													GoName:      "CpeUri",
													Description: "The cpe_uri in [cpe format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.",
												},
												"description": &dcl.Property{
													Type:        "string",
													GoName:      "Description",
													Description: "The distribution channel-specific description of this package.",
												},
												"latestVersion": &dcl.Property{
													Type:        "object",
													GoName:      "LatestVersion",
													GoType:      "NotePackageDistributionLatestVersion",
													Description: "The latest available version of this package in this distribution channel.",
													Required: []string{
														"kind",
													},
													Properties: map[string]*dcl.Property{
														"epoch": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "Epoch",
															Description: "Used to correct mistakes in the version numbering scheme.",
														},
														"fullName": &dcl.Property{
															Type:        "string",
															GoName:      "FullName",
															Description: "Human readable version string. This string is of the form :- and is only set when kind is NORMAL.",
														},
														"kind": &dcl.Property{
															Type:        "string",
															GoName:      "Kind",
															GoType:      "NotePackageDistributionLatestVersionKindEnum",
															Description: "Distinguish between sentinel MIN/MAX versions and normal versions. If kind is not NORMAL, then the other fields are ignored. Possible values: VERSION_KIND_UNSPECIFIED, NORMAL, MINIMUM, MAXIMUM",
															Enum: []string{
																"VERSION_KIND_UNSPECIFIED",
																"NORMAL",
																"MINIMUM",
																"MAXIMUM",
															},
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "The main part of the version name.",
														},
														"revision": &dcl.Property{
															Type:        "string",
															GoName:      "Revision",
															Description: "The iteration of the package build from the above version.",
														},
													},
												},
												"maintainer": &dcl.Property{
													Type:        "string",
													GoName:      "Maintainer",
													Description: "A freeform string denoting the maintainer of this package.",
												},
												"url": &dcl.Property{
													Type:        "string",
													GoName:      "Url",
													Description: "The distribution channel-specific homepage for this package.",
												},
											},
										},
									},
									"name": &dcl.Property{
										Type:        "string",
										GoName:      "Name",
										Description: "The name of the package.",
									},
								},
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
							"relatedNoteNames": &dcl.Property{
								Type:        "array",
								GoName:      "RelatedNoteNames",
								Description: "Other notes related to this note.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
									ResourceReferences: []*dcl.PropertyResourceReference{
										&dcl.PropertyResourceReference{
											Resource: "Containeranalysis/Note",
											Field:    "name",
										},
									},
								},
							},
							"relatedUrl": &dcl.Property{
								Type:        "array",
								GoName:      "RelatedUrl",
								Description: "URLs associated with this note.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "NoteRelatedUrl",
									Properties: map[string]*dcl.Property{
										"label": &dcl.Property{
											Type:        "string",
											GoName:      "Label",
											Description: "Label to describe usage of the URL",
										},
										"url": &dcl.Property{
											Type:        "string",
											GoName:      "Url",
											Description: "Specific URL to associate with the note",
										},
									},
								},
							},
							"shortDescription": &dcl.Property{
								Type:        "string",
								GoName:      "ShortDescription",
								Description: "A one sentence description of this note.",
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The time this note was last updated. This field can be used as a filter in list requests.",
								Immutable:   true,
							},
							"vulnerability": &dcl.Property{
								Type:        "object",
								GoName:      "Vulnerability",
								GoType:      "NoteVulnerability",
								Description: "A note describing a package vulnerability.",
								Conflicts: []string{
									"build",
									"image",
									"package",
									"deployment",
									"discovery",
									"attestation",
								},
								Properties: map[string]*dcl.Property{
									"cvssScore": &dcl.Property{
										Type:        "number",
										Format:      "double",
										GoName:      "CvssScore",
										Description: "The CVSS score of this vulnerability. CVSS score is on a scale of 0 - 10 where 0 indicates low severity and 10 indicates high severity.",
									},
									"cvssV3": &dcl.Property{
										Type:        "object",
										GoName:      "CvssV3",
										GoType:      "NoteVulnerabilityCvssV3",
										Description: "The full description of the CVSSv3 for this vulnerability.",
										Properties: map[string]*dcl.Property{
											"attackComplexity": &dcl.Property{
												Type:        "string",
												GoName:      "AttackComplexity",
												GoType:      "NoteVulnerabilityCvssV3AttackComplexityEnum",
												Description: " Possible values: ATTACK_COMPLEXITY_UNSPECIFIED, ATTACK_COMPLEXITY_LOW, ATTACK_COMPLEXITY_HIGH",
												Enum: []string{
													"ATTACK_COMPLEXITY_UNSPECIFIED",
													"ATTACK_COMPLEXITY_LOW",
													"ATTACK_COMPLEXITY_HIGH",
												},
											},
											"attackVector": &dcl.Property{
												Type:        "string",
												GoName:      "AttackVector",
												GoType:      "NoteVulnerabilityCvssV3AttackVectorEnum",
												Description: "Base Metrics Represents the intrinsic characteristics of a vulnerability that are constant over time and across user environments. Possible values: ATTACK_VECTOR_UNSPECIFIED, ATTACK_VECTOR_NETWORK, ATTACK_VECTOR_ADJACENT, ATTACK_VECTOR_LOCAL, ATTACK_VECTOR_PHYSICAL",
												Enum: []string{
													"ATTACK_VECTOR_UNSPECIFIED",
													"ATTACK_VECTOR_NETWORK",
													"ATTACK_VECTOR_ADJACENT",
													"ATTACK_VECTOR_LOCAL",
													"ATTACK_VECTOR_PHYSICAL",
												},
											},
											"availabilityImpact": &dcl.Property{
												Type:        "string",
												GoName:      "AvailabilityImpact",
												GoType:      "NoteVulnerabilityCvssV3AvailabilityImpactEnum",
												Description: " Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW, IMPACT_NONE",
												Enum: []string{
													"IMPACT_UNSPECIFIED",
													"IMPACT_HIGH",
													"IMPACT_LOW",
													"IMPACT_NONE",
												},
											},
											"baseScore": &dcl.Property{
												Type:        "number",
												Format:      "double",
												GoName:      "BaseScore",
												Description: "The base score is a function of the base metric scores.",
											},
											"confidentialityImpact": &dcl.Property{
												Type:        "string",
												GoName:      "ConfidentialityImpact",
												GoType:      "NoteVulnerabilityCvssV3ConfidentialityImpactEnum",
												Description: " Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW, IMPACT_NONE",
												Enum: []string{
													"IMPACT_UNSPECIFIED",
													"IMPACT_HIGH",
													"IMPACT_LOW",
													"IMPACT_NONE",
												},
											},
											"exploitabilityScore": &dcl.Property{
												Type:   "number",
												Format: "double",
												GoName: "ExploitabilityScore",
											},
											"impactScore": &dcl.Property{
												Type:   "number",
												Format: "double",
												GoName: "ImpactScore",
											},
											"integrityImpact": &dcl.Property{
												Type:        "string",
												GoName:      "IntegrityImpact",
												GoType:      "NoteVulnerabilityCvssV3IntegrityImpactEnum",
												Description: " Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW, IMPACT_NONE",
												Enum: []string{
													"IMPACT_UNSPECIFIED",
													"IMPACT_HIGH",
													"IMPACT_LOW",
													"IMPACT_NONE",
												},
											},
											"privilegesRequired": &dcl.Property{
												Type:        "string",
												GoName:      "PrivilegesRequired",
												GoType:      "NoteVulnerabilityCvssV3PrivilegesRequiredEnum",
												Description: " Possible values: PRIVILEGES_REQUIRED_UNSPECIFIED, PRIVILEGES_REQUIRED_NONE, PRIVILEGES_REQUIRED_LOW, PRIVILEGES_REQUIRED_HIGH",
												Enum: []string{
													"PRIVILEGES_REQUIRED_UNSPECIFIED",
													"PRIVILEGES_REQUIRED_NONE",
													"PRIVILEGES_REQUIRED_LOW",
													"PRIVILEGES_REQUIRED_HIGH",
												},
											},
											"scope": &dcl.Property{
												Type:        "string",
												GoName:      "Scope",
												GoType:      "NoteVulnerabilityCvssV3ScopeEnum",
												Description: " Possible values: SCOPE_UNSPECIFIED, SCOPE_UNCHANGED, SCOPE_CHANGED",
												Enum: []string{
													"SCOPE_UNSPECIFIED",
													"SCOPE_UNCHANGED",
													"SCOPE_CHANGED",
												},
											},
											"userInteraction": &dcl.Property{
												Type:        "string",
												GoName:      "UserInteraction",
												GoType:      "NoteVulnerabilityCvssV3UserInteractionEnum",
												Description: " Possible values: USER_INTERACTION_UNSPECIFIED, USER_INTERACTION_NONE, USER_INTERACTION_REQUIRED",
												Enum: []string{
													"USER_INTERACTION_UNSPECIFIED",
													"USER_INTERACTION_NONE",
													"USER_INTERACTION_REQUIRED",
												},
											},
										},
									},
									"details": &dcl.Property{
										Type:        "array",
										GoName:      "Details",
										Description: "Details of all known distros and packages affected by this vulnerability.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "NoteVulnerabilityDetails",
											Required: []string{
												"affectedCpeUri",
												"affectedPackage",
											},
											Properties: map[string]*dcl.Property{
												"affectedCpeUri": &dcl.Property{
													Type:        "string",
													GoName:      "AffectedCpeUri",
													Description: "Required. The (https://cpe.mitre.org/specification/) this vulnerability affects.",
												},
												"affectedPackage": &dcl.Property{
													Type:        "string",
													GoName:      "AffectedPackage",
													Description: "Required. The package this vulnerability affects.",
												},
												"affectedVersionEnd": &dcl.Property{
													Type:        "object",
													GoName:      "AffectedVersionEnd",
													GoType:      "NoteVulnerabilityDetailsAffectedVersionEnd",
													Description: "The version number at the end of an interval in which this vulnerability exists. A vulnerability can affect a package between version numbers that are disjoint sets of intervals (example: ) each of which will be represented in its own Detail. If a specific affected version is provided by a vulnerability database, affected_version_start and affected_version_end will be the same in that Detail.",
													Required: []string{
														"kind",
													},
													Properties: map[string]*dcl.Property{
														"epoch": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "Epoch",
															Description: "Used to correct mistakes in the version numbering scheme.",
														},
														"fullName": &dcl.Property{
															Type:        "string",
															GoName:      "FullName",
															Description: "Human readable version string. This string is of the form :- and is only set when kind is NORMAL.",
														},
														"kind": &dcl.Property{
															Type:        "string",
															GoName:      "Kind",
															GoType:      "NoteVulnerabilityDetailsAffectedVersionEndKindEnum",
															Description: "Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE",
															Enum: []string{
																"NOTE_KIND_UNSPECIFIED",
																"VULNERABILITY",
																"BUILD",
																"IMAGE",
																"PACKAGE",
																"DEPLOYMENT",
																"DISCOVERY",
																"ATTESTATION",
																"UPGRADE",
															},
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Required only when version kind is NORMAL. The main part of the version name.",
														},
														"revision": &dcl.Property{
															Type:        "string",
															GoName:      "Revision",
															Description: "The iteration of the package build from the above version.",
														},
													},
												},
												"affectedVersionStart": &dcl.Property{
													Type:        "object",
													GoName:      "AffectedVersionStart",
													GoType:      "NoteVulnerabilityDetailsAffectedVersionStart",
													Description: "The version number at the start of an interval in which this vulnerability exists. A vulnerability can affect a package between version numbers that are disjoint sets of intervals (example: ) each of which will be represented in its own Detail. If a specific affected version is provided by a vulnerability database, affected_version_start and affected_version_end will be the same in that Detail.",
													Required: []string{
														"kind",
													},
													Properties: map[string]*dcl.Property{
														"epoch": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "Epoch",
															Description: "Used to correct mistakes in the version numbering scheme.",
														},
														"fullName": &dcl.Property{
															Type:          "string",
															GoName:        "FullName",
															Description:   "Human readable version string. This string is of the form :- and is only set when kind is NORMAL.",
															ServerDefault: true,
														},
														"kind": &dcl.Property{
															Type:        "string",
															GoName:      "Kind",
															GoType:      "NoteVulnerabilityDetailsAffectedVersionStartKindEnum",
															Description: "Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE",
															Enum: []string{
																"NOTE_KIND_UNSPECIFIED",
																"VULNERABILITY",
																"BUILD",
																"IMAGE",
																"PACKAGE",
																"DEPLOYMENT",
																"DISCOVERY",
																"ATTESTATION",
																"UPGRADE",
															},
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Required only when version kind is NORMAL. The main part of the version name.",
														},
														"revision": &dcl.Property{
															Type:        "string",
															GoName:      "Revision",
															Description: "The iteration of the package build from the above version.",
														},
													},
												},
												"description": &dcl.Property{
													Type:        "string",
													GoName:      "Description",
													Description: "A vendor-specific description of this vulnerability.",
												},
												"fixedCpeUri": &dcl.Property{
													Type:        "string",
													GoName:      "FixedCpeUri",
													Description: "The distro recommended (https://cpe.mitre.org/specification/) to update to that contains a fix for this vulnerability. It is possible for this to be different from the affected_cpe_uri.",
												},
												"fixedPackage": &dcl.Property{
													Type:        "string",
													GoName:      "FixedPackage",
													Description: "The distro recommended package to update to that contains a fix for this vulnerability. It is possible for this to be different from the affected_package.",
												},
												"fixedVersion": &dcl.Property{
													Type:        "object",
													GoName:      "FixedVersion",
													GoType:      "NoteVulnerabilityDetailsFixedVersion",
													Description: "The distro recommended version to update to that contains a fix for this vulnerability. Setting this to VersionKind.MAXIMUM means no such version is yet available.",
													Required: []string{
														"kind",
													},
													Properties: map[string]*dcl.Property{
														"epoch": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "Epoch",
															Description: "Used to correct mistakes in the version numbering scheme.",
														},
														"fullName": &dcl.Property{
															Type:        "string",
															GoName:      "FullName",
															Description: "Human readable version string. This string is of the form :- and is only set when kind is NORMAL.",
														},
														"kind": &dcl.Property{
															Type:        "string",
															GoName:      "Kind",
															GoType:      "NoteVulnerabilityDetailsFixedVersionKindEnum",
															Description: "Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE",
															Enum: []string{
																"NOTE_KIND_UNSPECIFIED",
																"VULNERABILITY",
																"BUILD",
																"IMAGE",
																"PACKAGE",
																"DEPLOYMENT",
																"DISCOVERY",
																"ATTESTATION",
																"UPGRADE",
															},
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Required only when version kind is NORMAL. The main part of the version name.",
														},
														"revision": &dcl.Property{
															Type:        "string",
															GoName:      "Revision",
															Description: "The iteration of the package build from the above version.",
														},
													},
												},
												"isObsolete": &dcl.Property{
													Type:        "boolean",
													GoName:      "IsObsolete",
													Description: "Whether this detail is obsolete. Occurrences are expected not to point to obsolete details.",
												},
												"packageType": &dcl.Property{
													Type:        "string",
													GoName:      "PackageType",
													Description: "The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.).",
												},
												"severityName": &dcl.Property{
													Type:        "string",
													GoName:      "SeverityName",
													Description: "The distro assigned severity of this vulnerability.",
												},
												"sourceUpdateTime": &dcl.Property{
													Type:        "string",
													Format:      "date-time",
													GoName:      "SourceUpdateTime",
													Description: "The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker.",
												},
											},
										},
									},
									"severity": &dcl.Property{
										Type:        "string",
										GoName:      "Severity",
										GoType:      "NoteVulnerabilitySeverityEnum",
										Description: "The note provider assigned severity of this vulnerability. Possible values: SEVERITY_UNSPECIFIED, MINIMAL, LOW, MEDIUM, HIGH, CRITICAL",
										Enum: []string{
											"SEVERITY_UNSPECIFIED",
											"MINIMAL",
											"LOW",
											"MEDIUM",
											"HIGH",
											"CRITICAL",
										},
									},
									"sourceUpdateTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "SourceUpdateTime",
										Description: "The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker.",
									},
									"windowsDetails": &dcl.Property{
										Type:        "array",
										GoName:      "WindowsDetails",
										Description: "Windows details get their own format because the information format and model don't match a normal detail. Specifically Windows updates are done as patches, thus Windows vulnerabilities really are a missing package, rather than a package being at an incorrect version.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "NoteVulnerabilityWindowsDetails",
											Required: []string{
												"cpeUri",
												"name",
												"fixingKbs",
											},
											Properties: map[string]*dcl.Property{
												"cpeUri": &dcl.Property{
													Type:        "string",
													GoName:      "CpeUri",
													Description: "Required. The (https://cpe.mitre.org/specification/) this vulnerability affects.",
												},
												"description": &dcl.Property{
													Type:        "string",
													GoName:      "Description",
													Description: "The description of this vulnerability.",
												},
												"fixingKbs": &dcl.Property{
													Type:        "array",
													GoName:      "FixingKbs",
													Description: "Required. The names of the KBs which have hotfixes to mitigate this vulnerability. Note that there may be multiple hotfixes (and thus multiple KBs) that mitigate a given vulnerability. Currently any listed KBs presence is considered a fix.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "NoteVulnerabilityWindowsDetailsFixingKbs",
														Properties: map[string]*dcl.Property{
															"name": &dcl.Property{
																Type:        "string",
																GoName:      "Name",
																Description: "The KB name (generally of the form KB+ (e.g., KB123456)).",
															},
															"url": &dcl.Property{
																Type:        "string",
																GoName:      "Url",
																Description: "A link to the KB in the (https://www.catalog.update.microsoft.com/).",
															},
														},
													},
												},
												"name": &dcl.Property{
													Type:        "string",
													GoName:      "Name",
													Description: "Required. The name of this vulnerability.",
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
