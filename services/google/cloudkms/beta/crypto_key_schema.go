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

func DCLCryptoKeySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Cloudkms/CryptoKey",
			Description: "The Cloudkms CryptoKey resource",
			StructName:  "CryptoKey",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a CryptoKey",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "cryptoKey",
						Required:    true,
						Description: "A full instance of a CryptoKey",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a CryptoKey",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "cryptoKey",
						Required:    true,
						Description: "A full instance of a CryptoKey",
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many CryptoKey",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "keyRing",
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
				"CryptoKey": &dcl.Component{
					Title:           "CryptoKey",
					ID:              "projects/{{project}}/locations/{{location}}/keyRings/{{key_ring}}/cryptoKeys/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"purpose",
							"project",
							"location",
							"keyRing",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time at which this CryptoKey was created.",
								Immutable:   true,
							},
							"destroyScheduledDuration": &dcl.Property{
								Type:          "string",
								GoName:        "DestroyScheduledDuration",
								Description:   "Immutable. The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED. If not specified at creation time, the default duration is 24 hours.",
								Immutable:     true,
								ServerDefault: true,
							},
							"importOnly": &dcl.Property{
								Type:        "boolean",
								GoName:      "ImportOnly",
								Description: "Immutable. Whether this key may contain imported versions only.",
								Immutable:   true,
							},
							"keyRing": &dcl.Property{
								Type:        "string",
								GoName:      "KeyRing",
								Description: "The key ring for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudkms/KeyRing",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Labels with user-defined metadata. For more information, see [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The resource name for this CryptoKey in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*`.",
								Immutable:   true,
							},
							"nextRotationTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "NextRotationTime",
								Description: "At next_rotation_time, the Key Management Service will automatically: 1. Create a new version of this CryptoKey. 2. Mark the new version as primary. Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.",
							},
							"primary": &dcl.Property{
								Type:        "object",
								GoName:      "Primary",
								GoType:      "CryptoKeyPrimary",
								ReadOnly:    true,
								Description: "Output only. A copy of the \"primary\" CryptoKeyVersion that will be used by Encrypt when this CryptoKey is given in EncryptRequest.name. The CryptoKey's primary version can be updated via UpdateCryptoKeyPrimaryVersion. Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys, this field will be omitted.",
								Properties: map[string]*dcl.Property{
									"algorithm": &dcl.Property{
										Type:        "string",
										GoName:      "Algorithm",
										GoType:      "CryptoKeyPrimaryAlgorithmEnum",
										ReadOnly:    true,
										Description: "Output only. The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports. Possible values: CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED, GOOGLE_SYMMETRIC_ENCRYPTION, RSA_SIGN_PSS_2048_SHA256, RSA_SIGN_PSS_3072_SHA256, RSA_SIGN_PSS_4096_SHA256, RSA_SIGN_PSS_4096_SHA512, RSA_SIGN_PKCS1_2048_SHA256, RSA_SIGN_PKCS1_3072_SHA256, RSA_SIGN_PKCS1_4096_SHA256, RSA_SIGN_PKCS1_4096_SHA512, RSA_DECRYPT_OAEP_2048_SHA256, RSA_DECRYPT_OAEP_3072_SHA256, RSA_DECRYPT_OAEP_4096_SHA256, RSA_DECRYPT_OAEP_4096_SHA512, EC_SIGN_P256_SHA256, EC_SIGN_P384_SHA384, EC_SIGN_SECP256K1_SHA256, HMAC_SHA256, EXTERNAL_SYMMETRIC_ENCRYPTION",
										Immutable:   true,
										Enum: []string{
											"CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED",
											"GOOGLE_SYMMETRIC_ENCRYPTION",
											"RSA_SIGN_PSS_2048_SHA256",
											"RSA_SIGN_PSS_3072_SHA256",
											"RSA_SIGN_PSS_4096_SHA256",
											"RSA_SIGN_PSS_4096_SHA512",
											"RSA_SIGN_PKCS1_2048_SHA256",
											"RSA_SIGN_PKCS1_3072_SHA256",
											"RSA_SIGN_PKCS1_4096_SHA256",
											"RSA_SIGN_PKCS1_4096_SHA512",
											"RSA_DECRYPT_OAEP_2048_SHA256",
											"RSA_DECRYPT_OAEP_3072_SHA256",
											"RSA_DECRYPT_OAEP_4096_SHA256",
											"RSA_DECRYPT_OAEP_4096_SHA512",
											"EC_SIGN_P256_SHA256",
											"EC_SIGN_P384_SHA384",
											"EC_SIGN_SECP256K1_SHA256",
											"HMAC_SHA256",
											"EXTERNAL_SYMMETRIC_ENCRYPTION",
										},
									},
									"attestation": &dcl.Property{
										Type:        "object",
										GoName:      "Attestation",
										GoType:      "CryptoKeyPrimaryAttestation",
										ReadOnly:    true,
										Description: "Output only. Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only provided for key versions with protection_level HSM.",
										Properties: map[string]*dcl.Property{
											"certChains": &dcl.Property{
												Type:        "object",
												GoName:      "CertChains",
												GoType:      "CryptoKeyPrimaryAttestationCertChains",
												ReadOnly:    true,
												Description: "Output only. The certificate chains needed to validate the attestation",
												Properties: map[string]*dcl.Property{
													"caviumCerts": &dcl.Property{
														Type:        "array",
														GoName:      "CaviumCerts",
														Description: "Cavium certificate chain corresponding to the attestation.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"googleCardCerts": &dcl.Property{
														Type:        "array",
														GoName:      "GoogleCardCerts",
														Description: "Google card certificate chain corresponding to the attestation.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"googlePartitionCerts": &dcl.Property{
														Type:        "array",
														GoName:      "GooglePartitionCerts",
														Description: "Google partition certificate chain corresponding to the attestation.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
												},
											},
											"content": &dcl.Property{
												Type:        "string",
												GoName:      "Content",
												ReadOnly:    true,
												Description: "Output only. The attestation data provided by the HSM when the key operation was performed.",
												Immutable:   true,
											},
											"format": &dcl.Property{
												Type:        "string",
												GoName:      "Format",
												GoType:      "CryptoKeyPrimaryAttestationFormatEnum",
												ReadOnly:    true,
												Description: "Output only. The format of the attestation data. Possible values: ATTESTATION_FORMAT_UNSPECIFIED, CAVIUM_V1_COMPRESSED, CAVIUM_V2_COMPRESSED",
												Immutable:   true,
												Enum: []string{
													"ATTESTATION_FORMAT_UNSPECIFIED",
													"CAVIUM_V1_COMPRESSED",
													"CAVIUM_V2_COMPRESSED",
												},
											},
										},
									},
									"createTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "CreateTime",
										ReadOnly:    true,
										Description: "Output only. The time at which this CryptoKeyVersion was created.",
										Immutable:   true,
									},
									"destroyEventTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "DestroyEventTime",
										ReadOnly:    true,
										Description: "Output only. The time this CryptoKeyVersion's key material was destroyed. Only present if state is DESTROYED.",
										Immutable:   true,
									},
									"destroyTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "DestroyTime",
										ReadOnly:    true,
										Description: "Output only. The time this CryptoKeyVersion's key material is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.",
										Immutable:   true,
									},
									"externalProtectionLevelOptions": &dcl.Property{
										Type:        "object",
										GoName:      "ExternalProtectionLevelOptions",
										GoType:      "CryptoKeyPrimaryExternalProtectionLevelOptions",
										Description: "ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level.",
										Properties: map[string]*dcl.Property{
											"externalKeyUri": &dcl.Property{
												Type:        "string",
												GoName:      "ExternalKeyUri",
												Description: "The URI for an external resource that this CryptoKeyVersion represents.",
											},
										},
									},
									"generateTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "GenerateTime",
										ReadOnly:    true,
										Description: "Output only. The time this CryptoKeyVersion's key material was generated.",
										Immutable:   true,
									},
									"importFailureReason": &dcl.Property{
										Type:        "string",
										GoName:      "ImportFailureReason",
										ReadOnly:    true,
										Description: "Output only. The root cause of the most recent import failure. Only present if state is IMPORT_FAILED.",
										Immutable:   true,
									},
									"importJob": &dcl.Property{
										Type:        "string",
										GoName:      "ImportJob",
										ReadOnly:    true,
										Description: "Output only. The name of the ImportJob used in the most recent import of this CryptoKeyVersion. Only present if the underlying key material was imported.",
										Immutable:   true,
									},
									"importTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "ImportTime",
										ReadOnly:    true,
										Description: "Output only. The time at which this CryptoKeyVersion's key material was most recently imported.",
										Immutable:   true,
									},
									"name": &dcl.Property{
										Type:                     "string",
										GoName:                   "Name",
										Description:              "Output only. The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.",
										Immutable:                true,
										ServerGeneratedParameter: true,
									},
									"protectionLevel": &dcl.Property{
										Type:        "string",
										GoName:      "ProtectionLevel",
										GoType:      "CryptoKeyPrimaryProtectionLevelEnum",
										ReadOnly:    true,
										Description: "Output only. The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. Possible values: PROTECTION_LEVEL_UNSPECIFIED, SOFTWARE, HSM, EXTERNAL, EXTERNAL_VPC",
										Immutable:   true,
										Enum: []string{
											"PROTECTION_LEVEL_UNSPECIFIED",
											"SOFTWARE",
											"HSM",
											"EXTERNAL",
											"EXTERNAL_VPC",
										},
									},
									"reimportEligible": &dcl.Property{
										Type:        "boolean",
										GoName:      "ReimportEligible",
										ReadOnly:    true,
										Description: "Output only. Whether or not this key version is eligible for reimport, by being specified as a target in ImportCryptoKeyVersionRequest.crypto_key_version.",
										Immutable:   true,
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "CryptoKeyPrimaryStateEnum",
										Description: "The current state of the CryptoKeyVersion. Possible values: CRYPTO_KEY_VERSION_STATE_UNSPECIFIED, PENDING_GENERATION, ENABLED, DISABLED, DESTROYED, DESTROY_SCHEDULED, PENDING_IMPORT, IMPORT_FAILED",
										Enum: []string{
											"CRYPTO_KEY_VERSION_STATE_UNSPECIFIED",
											"PENDING_GENERATION",
											"ENABLED",
											"DISABLED",
											"DESTROYED",
											"DESTROY_SCHEDULED",
											"PENDING_IMPORT",
											"IMPORT_FAILED",
										},
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
							"purpose": &dcl.Property{
								Type:        "string",
								GoName:      "Purpose",
								GoType:      "CryptoKeyPurposeEnum",
								Description: "Immutable. The immutable purpose of this CryptoKey. Possible values: CRYPTO_KEY_PURPOSE_UNSPECIFIED, ENCRYPT_DECRYPT, ASYMMETRIC_SIGN, ASYMMETRIC_DECRYPT, MAC",
								Immutable:   true,
								Enum: []string{
									"CRYPTO_KEY_PURPOSE_UNSPECIFIED",
									"ENCRYPT_DECRYPT",
									"ASYMMETRIC_SIGN",
									"ASYMMETRIC_DECRYPT",
									"MAC",
								},
							},
							"rotationPeriod": &dcl.Property{
								Type:        "string",
								GoName:      "RotationPeriod",
								Description: "next_rotation_time will be advanced by this period when the service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours. If rotation_period is set, next_rotation_time must also be set. Keys with purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this field must be omitted.",
							},
							"versionTemplate": &dcl.Property{
								Type:          "object",
								GoName:        "VersionTemplate",
								GoType:        "CryptoKeyVersionTemplate",
								Description:   "A template describing settings for new CryptoKeyVersion instances. The properties of new CryptoKeyVersion instances created by either CreateCryptoKeyVersion or auto-rotation are controlled by this template.",
								ServerDefault: true,
								Required: []string{
									"algorithm",
								},
								Properties: map[string]*dcl.Property{
									"algorithm": &dcl.Property{
										Type:        "string",
										GoName:      "Algorithm",
										GoType:      "CryptoKeyVersionTemplateAlgorithmEnum",
										Description: "Required. Algorithm to use when creating a CryptoKeyVersion based on this template. For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both this field is omitted and CryptoKey.purpose is ENCRYPT_DECRYPT. Possible values: CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED, GOOGLE_SYMMETRIC_ENCRYPTION, RSA_SIGN_PSS_2048_SHA256, RSA_SIGN_PSS_3072_SHA256, RSA_SIGN_PSS_4096_SHA256, RSA_SIGN_PSS_4096_SHA512, RSA_SIGN_PKCS1_2048_SHA256, RSA_SIGN_PKCS1_3072_SHA256, RSA_SIGN_PKCS1_4096_SHA256, RSA_SIGN_PKCS1_4096_SHA512, RSA_DECRYPT_OAEP_2048_SHA256, RSA_DECRYPT_OAEP_3072_SHA256, RSA_DECRYPT_OAEP_4096_SHA256, RSA_DECRYPT_OAEP_4096_SHA512, EC_SIGN_P256_SHA256, EC_SIGN_P384_SHA384, EC_SIGN_SECP256K1_SHA256, HMAC_SHA256, EXTERNAL_SYMMETRIC_ENCRYPTION",
										Enum: []string{
											"CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED",
											"GOOGLE_SYMMETRIC_ENCRYPTION",
											"RSA_SIGN_PSS_2048_SHA256",
											"RSA_SIGN_PSS_3072_SHA256",
											"RSA_SIGN_PSS_4096_SHA256",
											"RSA_SIGN_PSS_4096_SHA512",
											"RSA_SIGN_PKCS1_2048_SHA256",
											"RSA_SIGN_PKCS1_3072_SHA256",
											"RSA_SIGN_PKCS1_4096_SHA256",
											"RSA_SIGN_PKCS1_4096_SHA512",
											"RSA_DECRYPT_OAEP_2048_SHA256",
											"RSA_DECRYPT_OAEP_3072_SHA256",
											"RSA_DECRYPT_OAEP_4096_SHA256",
											"RSA_DECRYPT_OAEP_4096_SHA512",
											"EC_SIGN_P256_SHA256",
											"EC_SIGN_P384_SHA384",
											"EC_SIGN_SECP256K1_SHA256",
											"HMAC_SHA256",
											"EXTERNAL_SYMMETRIC_ENCRYPTION",
										},
									},
									"protectionLevel": &dcl.Property{
										Type:          "string",
										GoName:        "ProtectionLevel",
										GoType:        "CryptoKeyVersionTemplateProtectionLevelEnum",
										Description:   "ProtectionLevel to use when creating a CryptoKeyVersion based on this template. Immutable. Defaults to SOFTWARE. Possible values: PROTECTION_LEVEL_UNSPECIFIED, SOFTWARE, HSM, EXTERNAL, EXTERNAL_VPC",
										Immutable:     true,
										ServerDefault: true,
										Enum: []string{
											"PROTECTION_LEVEL_UNSPECIFIED",
											"SOFTWARE",
											"HSM",
											"EXTERNAL",
											"EXTERNAL_VPC",
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
