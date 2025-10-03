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
package cloudkms

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type CryptoKey struct{}

func CryptoKeyToUnstructured(r *dclService.CryptoKey) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "cloudkms",
			Version: "beta",
			Type:    "CryptoKey",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DestroyScheduledDuration != nil {
		u.Object["destroyScheduledDuration"] = *r.DestroyScheduledDuration
	}
	if r.ImportOnly != nil {
		u.Object["importOnly"] = *r.ImportOnly
	}
	if r.KeyRing != nil {
		u.Object["keyRing"] = *r.KeyRing
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.NextRotationTime != nil {
		u.Object["nextRotationTime"] = *r.NextRotationTime
	}
	if r.Primary != nil && r.Primary != dclService.EmptyCryptoKeyPrimary {
		rPrimary := make(map[string]interface{})
		if r.Primary.Algorithm != nil {
			rPrimary["algorithm"] = string(*r.Primary.Algorithm)
		}
		if r.Primary.Attestation != nil && r.Primary.Attestation != dclService.EmptyCryptoKeyPrimaryAttestation {
			rPrimaryAttestation := make(map[string]interface{})
			if r.Primary.Attestation.CertChains != nil && r.Primary.Attestation.CertChains != dclService.EmptyCryptoKeyPrimaryAttestationCertChains {
				rPrimaryAttestationCertChains := make(map[string]interface{})
				var rPrimaryAttestationCertChainsCaviumCerts []interface{}
				for _, rPrimaryAttestationCertChainsCaviumCertsVal := range r.Primary.Attestation.CertChains.CaviumCerts {
					rPrimaryAttestationCertChainsCaviumCerts = append(rPrimaryAttestationCertChainsCaviumCerts, rPrimaryAttestationCertChainsCaviumCertsVal)
				}
				rPrimaryAttestationCertChains["caviumCerts"] = rPrimaryAttestationCertChainsCaviumCerts
				var rPrimaryAttestationCertChainsGoogleCardCerts []interface{}
				for _, rPrimaryAttestationCertChainsGoogleCardCertsVal := range r.Primary.Attestation.CertChains.GoogleCardCerts {
					rPrimaryAttestationCertChainsGoogleCardCerts = append(rPrimaryAttestationCertChainsGoogleCardCerts, rPrimaryAttestationCertChainsGoogleCardCertsVal)
				}
				rPrimaryAttestationCertChains["googleCardCerts"] = rPrimaryAttestationCertChainsGoogleCardCerts
				var rPrimaryAttestationCertChainsGooglePartitionCerts []interface{}
				for _, rPrimaryAttestationCertChainsGooglePartitionCertsVal := range r.Primary.Attestation.CertChains.GooglePartitionCerts {
					rPrimaryAttestationCertChainsGooglePartitionCerts = append(rPrimaryAttestationCertChainsGooglePartitionCerts, rPrimaryAttestationCertChainsGooglePartitionCertsVal)
				}
				rPrimaryAttestationCertChains["googlePartitionCerts"] = rPrimaryAttestationCertChainsGooglePartitionCerts
				rPrimaryAttestation["certChains"] = rPrimaryAttestationCertChains
			}
			if r.Primary.Attestation.Content != nil {
				rPrimaryAttestation["content"] = *r.Primary.Attestation.Content
			}
			if r.Primary.Attestation.Format != nil {
				rPrimaryAttestation["format"] = string(*r.Primary.Attestation.Format)
			}
			rPrimary["attestation"] = rPrimaryAttestation
		}
		if r.Primary.CreateTime != nil {
			rPrimary["createTime"] = *r.Primary.CreateTime
		}
		if r.Primary.DestroyEventTime != nil {
			rPrimary["destroyEventTime"] = *r.Primary.DestroyEventTime
		}
		if r.Primary.DestroyTime != nil {
			rPrimary["destroyTime"] = *r.Primary.DestroyTime
		}
		if r.Primary.ExternalProtectionLevelOptions != nil && r.Primary.ExternalProtectionLevelOptions != dclService.EmptyCryptoKeyPrimaryExternalProtectionLevelOptions {
			rPrimaryExternalProtectionLevelOptions := make(map[string]interface{})
			if r.Primary.ExternalProtectionLevelOptions.ExternalKeyUri != nil {
				rPrimaryExternalProtectionLevelOptions["externalKeyUri"] = *r.Primary.ExternalProtectionLevelOptions.ExternalKeyUri
			}
			rPrimary["externalProtectionLevelOptions"] = rPrimaryExternalProtectionLevelOptions
		}
		if r.Primary.GenerateTime != nil {
			rPrimary["generateTime"] = *r.Primary.GenerateTime
		}
		if r.Primary.ImportFailureReason != nil {
			rPrimary["importFailureReason"] = *r.Primary.ImportFailureReason
		}
		if r.Primary.ImportJob != nil {
			rPrimary["importJob"] = *r.Primary.ImportJob
		}
		if r.Primary.ImportTime != nil {
			rPrimary["importTime"] = *r.Primary.ImportTime
		}
		if r.Primary.Name != nil {
			rPrimary["name"] = *r.Primary.Name
		}
		if r.Primary.ProtectionLevel != nil {
			rPrimary["protectionLevel"] = string(*r.Primary.ProtectionLevel)
		}
		if r.Primary.ReimportEligible != nil {
			rPrimary["reimportEligible"] = *r.Primary.ReimportEligible
		}
		if r.Primary.State != nil {
			rPrimary["state"] = string(*r.Primary.State)
		}
		u.Object["primary"] = rPrimary
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Purpose != nil {
		u.Object["purpose"] = string(*r.Purpose)
	}
	if r.RotationPeriod != nil {
		u.Object["rotationPeriod"] = *r.RotationPeriod
	}
	if r.VersionTemplate != nil && r.VersionTemplate != dclService.EmptyCryptoKeyVersionTemplate {
		rVersionTemplate := make(map[string]interface{})
		if r.VersionTemplate.Algorithm != nil {
			rVersionTemplate["algorithm"] = string(*r.VersionTemplate.Algorithm)
		}
		if r.VersionTemplate.ProtectionLevel != nil {
			rVersionTemplate["protectionLevel"] = string(*r.VersionTemplate.ProtectionLevel)
		}
		u.Object["versionTemplate"] = rVersionTemplate
	}
	return u
}

func UnstructuredToCryptoKey(u *unstructured.Resource) (*dclService.CryptoKey, error) {
	r := &dclService.CryptoKey{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["destroyScheduledDuration"]; ok {
		if s, ok := u.Object["destroyScheduledDuration"].(string); ok {
			r.DestroyScheduledDuration = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DestroyScheduledDuration: expected string")
		}
	}
	if _, ok := u.Object["importOnly"]; ok {
		if b, ok := u.Object["importOnly"].(bool); ok {
			r.ImportOnly = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.ImportOnly: expected bool")
		}
	}
	if _, ok := u.Object["keyRing"]; ok {
		if s, ok := u.Object["keyRing"].(string); ok {
			r.KeyRing = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.KeyRing: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["nextRotationTime"]; ok {
		if s, ok := u.Object["nextRotationTime"].(string); ok {
			r.NextRotationTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextRotationTime: expected string")
		}
	}
	if _, ok := u.Object["primary"]; ok {
		if rPrimary, ok := u.Object["primary"].(map[string]interface{}); ok {
			r.Primary = &dclService.CryptoKeyPrimary{}
			if _, ok := rPrimary["algorithm"]; ok {
				if s, ok := rPrimary["algorithm"].(string); ok {
					r.Primary.Algorithm = dclService.CryptoKeyPrimaryAlgorithmEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Primary.Algorithm: expected string")
				}
			}
			if _, ok := rPrimary["attestation"]; ok {
				if rPrimaryAttestation, ok := rPrimary["attestation"].(map[string]interface{}); ok {
					r.Primary.Attestation = &dclService.CryptoKeyPrimaryAttestation{}
					if _, ok := rPrimaryAttestation["certChains"]; ok {
						if rPrimaryAttestationCertChains, ok := rPrimaryAttestation["certChains"].(map[string]interface{}); ok {
							r.Primary.Attestation.CertChains = &dclService.CryptoKeyPrimaryAttestationCertChains{}
							if _, ok := rPrimaryAttestationCertChains["caviumCerts"]; ok {
								if s, ok := rPrimaryAttestationCertChains["caviumCerts"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Primary.Attestation.CertChains.CaviumCerts = append(r.Primary.Attestation.CertChains.CaviumCerts, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Primary.Attestation.CertChains.CaviumCerts: expected []interface{}")
								}
							}
							if _, ok := rPrimaryAttestationCertChains["googleCardCerts"]; ok {
								if s, ok := rPrimaryAttestationCertChains["googleCardCerts"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Primary.Attestation.CertChains.GoogleCardCerts = append(r.Primary.Attestation.CertChains.GoogleCardCerts, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Primary.Attestation.CertChains.GoogleCardCerts: expected []interface{}")
								}
							}
							if _, ok := rPrimaryAttestationCertChains["googlePartitionCerts"]; ok {
								if s, ok := rPrimaryAttestationCertChains["googlePartitionCerts"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Primary.Attestation.CertChains.GooglePartitionCerts = append(r.Primary.Attestation.CertChains.GooglePartitionCerts, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Primary.Attestation.CertChains.GooglePartitionCerts: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Primary.Attestation.CertChains: expected map[string]interface{}")
						}
					}
					if _, ok := rPrimaryAttestation["content"]; ok {
						if s, ok := rPrimaryAttestation["content"].(string); ok {
							r.Primary.Attestation.Content = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Primary.Attestation.Content: expected string")
						}
					}
					if _, ok := rPrimaryAttestation["format"]; ok {
						if s, ok := rPrimaryAttestation["format"].(string); ok {
							r.Primary.Attestation.Format = dclService.CryptoKeyPrimaryAttestationFormatEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Primary.Attestation.Format: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Primary.Attestation: expected map[string]interface{}")
				}
			}
			if _, ok := rPrimary["createTime"]; ok {
				if s, ok := rPrimary["createTime"].(string); ok {
					r.Primary.CreateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.CreateTime: expected string")
				}
			}
			if _, ok := rPrimary["destroyEventTime"]; ok {
				if s, ok := rPrimary["destroyEventTime"].(string); ok {
					r.Primary.DestroyEventTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.DestroyEventTime: expected string")
				}
			}
			if _, ok := rPrimary["destroyTime"]; ok {
				if s, ok := rPrimary["destroyTime"].(string); ok {
					r.Primary.DestroyTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.DestroyTime: expected string")
				}
			}
			if _, ok := rPrimary["externalProtectionLevelOptions"]; ok {
				if rPrimaryExternalProtectionLevelOptions, ok := rPrimary["externalProtectionLevelOptions"].(map[string]interface{}); ok {
					r.Primary.ExternalProtectionLevelOptions = &dclService.CryptoKeyPrimaryExternalProtectionLevelOptions{}
					if _, ok := rPrimaryExternalProtectionLevelOptions["externalKeyUri"]; ok {
						if s, ok := rPrimaryExternalProtectionLevelOptions["externalKeyUri"].(string); ok {
							r.Primary.ExternalProtectionLevelOptions.ExternalKeyUri = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Primary.ExternalProtectionLevelOptions.ExternalKeyUri: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Primary.ExternalProtectionLevelOptions: expected map[string]interface{}")
				}
			}
			if _, ok := rPrimary["generateTime"]; ok {
				if s, ok := rPrimary["generateTime"].(string); ok {
					r.Primary.GenerateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.GenerateTime: expected string")
				}
			}
			if _, ok := rPrimary["importFailureReason"]; ok {
				if s, ok := rPrimary["importFailureReason"].(string); ok {
					r.Primary.ImportFailureReason = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.ImportFailureReason: expected string")
				}
			}
			if _, ok := rPrimary["importJob"]; ok {
				if s, ok := rPrimary["importJob"].(string); ok {
					r.Primary.ImportJob = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.ImportJob: expected string")
				}
			}
			if _, ok := rPrimary["importTime"]; ok {
				if s, ok := rPrimary["importTime"].(string); ok {
					r.Primary.ImportTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.ImportTime: expected string")
				}
			}
			if _, ok := rPrimary["name"]; ok {
				if s, ok := rPrimary["name"].(string); ok {
					r.Primary.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Primary.Name: expected string")
				}
			}
			if _, ok := rPrimary["protectionLevel"]; ok {
				if s, ok := rPrimary["protectionLevel"].(string); ok {
					r.Primary.ProtectionLevel = dclService.CryptoKeyPrimaryProtectionLevelEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Primary.ProtectionLevel: expected string")
				}
			}
			if _, ok := rPrimary["reimportEligible"]; ok {
				if b, ok := rPrimary["reimportEligible"].(bool); ok {
					r.Primary.ReimportEligible = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Primary.ReimportEligible: expected bool")
				}
			}
			if _, ok := rPrimary["state"]; ok {
				if s, ok := rPrimary["state"].(string); ok {
					r.Primary.State = dclService.CryptoKeyPrimaryStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Primary.State: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Primary: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["purpose"]; ok {
		if s, ok := u.Object["purpose"].(string); ok {
			r.Purpose = dclService.CryptoKeyPurposeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Purpose: expected string")
		}
	}
	if _, ok := u.Object["rotationPeriod"]; ok {
		if s, ok := u.Object["rotationPeriod"].(string); ok {
			r.RotationPeriod = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RotationPeriod: expected string")
		}
	}
	if _, ok := u.Object["versionTemplate"]; ok {
		if rVersionTemplate, ok := u.Object["versionTemplate"].(map[string]interface{}); ok {
			r.VersionTemplate = &dclService.CryptoKeyVersionTemplate{}
			if _, ok := rVersionTemplate["algorithm"]; ok {
				if s, ok := rVersionTemplate["algorithm"].(string); ok {
					r.VersionTemplate.Algorithm = dclService.CryptoKeyVersionTemplateAlgorithmEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.VersionTemplate.Algorithm: expected string")
				}
			}
			if _, ok := rVersionTemplate["protectionLevel"]; ok {
				if s, ok := rVersionTemplate["protectionLevel"].(string); ok {
					r.VersionTemplate.ProtectionLevel = dclService.CryptoKeyVersionTemplateProtectionLevelEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.VersionTemplate.ProtectionLevel: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.VersionTemplate: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCryptoKey(ctx, r)
	if err != nil {
		return nil, err
	}
	return CryptoKeyToUnstructured(r), nil
}

func ListCryptoKey(ctx context.Context, config *dcl.Config, project string, location string, keyRing string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCryptoKey(ctx, project, location, keyRing)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, CryptoKeyToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCryptoKey(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCryptoKey(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return CryptoKeyToUnstructured(r), nil
}

func CryptoKeyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCryptoKey(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCryptoKey(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func CryptoKeyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *CryptoKey) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"cloudkms",
		"CryptoKey",
		"beta",
	}
}

func SetPolicyCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberCryptoKey(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToCryptoKey(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *CryptoKey) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberCryptoKey(ctx, config, resource, member)
}

func (r *CryptoKey) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberCryptoKey(ctx, config, resource, role, member)
}

func (r *CryptoKey) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberCryptoKey(ctx, config, resource, member)
}

func (r *CryptoKey) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyCryptoKey(ctx, config, resource, policy)
}

func (r *CryptoKey) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagCryptoKey(ctx, config, resource, policy)
}

func (r *CryptoKey) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyCryptoKey(ctx, config, resource)
}

func (r *CryptoKey) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCryptoKey(ctx, config, resource)
}

func (r *CryptoKey) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCryptoKey(ctx, config, resource, opts...)
}

func (r *CryptoKey) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return CryptoKeyHasDiff(ctx, config, resource, opts...)
}

func (r *CryptoKey) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCryptoKey(ctx, config, resource)
}

func (r *CryptoKey) ID(resource *unstructured.Resource) (string, error) {
	return CryptoKeyID(resource)
}

func init() {
	unstructured.Register(&CryptoKey{})
}
