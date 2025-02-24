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
package binaryauthorization

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Attestor struct{}

func AttestorToUnstructured(r *dclService.Attestor) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "binaryauthorization",
			Version: "beta",
			Type:    "Attestor",
		},
		Object: make(map[string]interface{}),
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.UserOwnedDrydockNote != nil && r.UserOwnedDrydockNote != dclService.EmptyAttestorUserOwnedDrydockNote {
		rUserOwnedDrydockNote := make(map[string]interface{})
		if r.UserOwnedDrydockNote.DelegationServiceAccountEmail != nil {
			rUserOwnedDrydockNote["delegationServiceAccountEmail"] = *r.UserOwnedDrydockNote.DelegationServiceAccountEmail
		}
		if r.UserOwnedDrydockNote.NoteReference != nil {
			rUserOwnedDrydockNote["noteReference"] = *r.UserOwnedDrydockNote.NoteReference
		}
		var rUserOwnedDrydockNotePublicKeys []interface{}
		for _, rUserOwnedDrydockNotePublicKeysVal := range r.UserOwnedDrydockNote.PublicKeys {
			rUserOwnedDrydockNotePublicKeysObject := make(map[string]interface{})
			if rUserOwnedDrydockNotePublicKeysVal.AsciiArmoredPgpPublicKey != nil {
				rUserOwnedDrydockNotePublicKeysObject["asciiArmoredPgpPublicKey"] = *rUserOwnedDrydockNotePublicKeysVal.AsciiArmoredPgpPublicKey
			}
			if rUserOwnedDrydockNotePublicKeysVal.Comment != nil {
				rUserOwnedDrydockNotePublicKeysObject["comment"] = *rUserOwnedDrydockNotePublicKeysVal.Comment
			}
			if rUserOwnedDrydockNotePublicKeysVal.Id != nil {
				rUserOwnedDrydockNotePublicKeysObject["id"] = *rUserOwnedDrydockNotePublicKeysVal.Id
			}
			if rUserOwnedDrydockNotePublicKeysVal.PkixPublicKey != nil && rUserOwnedDrydockNotePublicKeysVal.PkixPublicKey != dclService.EmptyAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
				rUserOwnedDrydockNotePublicKeysValPkixPublicKey := make(map[string]interface{})
				if rUserOwnedDrydockNotePublicKeysVal.PkixPublicKey.PublicKeyPem != nil {
					rUserOwnedDrydockNotePublicKeysValPkixPublicKey["publicKeyPem"] = *rUserOwnedDrydockNotePublicKeysVal.PkixPublicKey.PublicKeyPem
				}
				if rUserOwnedDrydockNotePublicKeysVal.PkixPublicKey.SignatureAlgorithm != nil {
					rUserOwnedDrydockNotePublicKeysValPkixPublicKey["signatureAlgorithm"] = string(*rUserOwnedDrydockNotePublicKeysVal.PkixPublicKey.SignatureAlgorithm)
				}
				rUserOwnedDrydockNotePublicKeysObject["pkixPublicKey"] = rUserOwnedDrydockNotePublicKeysValPkixPublicKey
			}
			rUserOwnedDrydockNotePublicKeys = append(rUserOwnedDrydockNotePublicKeys, rUserOwnedDrydockNotePublicKeysObject)
		}
		rUserOwnedDrydockNote["publicKeys"] = rUserOwnedDrydockNotePublicKeys
		u.Object["userOwnedDrydockNote"] = rUserOwnedDrydockNote
	}
	return u
}

func UnstructuredToAttestor(u *unstructured.Resource) (*dclService.Attestor, error) {
	r := &dclService.Attestor{}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["userOwnedDrydockNote"]; ok {
		if rUserOwnedDrydockNote, ok := u.Object["userOwnedDrydockNote"].(map[string]interface{}); ok {
			r.UserOwnedDrydockNote = &dclService.AttestorUserOwnedDrydockNote{}
			if _, ok := rUserOwnedDrydockNote["delegationServiceAccountEmail"]; ok {
				if s, ok := rUserOwnedDrydockNote["delegationServiceAccountEmail"].(string); ok {
					r.UserOwnedDrydockNote.DelegationServiceAccountEmail = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.UserOwnedDrydockNote.DelegationServiceAccountEmail: expected string")
				}
			}
			if _, ok := rUserOwnedDrydockNote["noteReference"]; ok {
				if s, ok := rUserOwnedDrydockNote["noteReference"].(string); ok {
					r.UserOwnedDrydockNote.NoteReference = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.UserOwnedDrydockNote.NoteReference: expected string")
				}
			}
			if _, ok := rUserOwnedDrydockNote["publicKeys"]; ok {
				if s, ok := rUserOwnedDrydockNote["publicKeys"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rUserOwnedDrydockNotePublicKeys dclService.AttestorUserOwnedDrydockNotePublicKeys
							if _, ok := objval["asciiArmoredPgpPublicKey"]; ok {
								if s, ok := objval["asciiArmoredPgpPublicKey"].(string); ok {
									rUserOwnedDrydockNotePublicKeys.AsciiArmoredPgpPublicKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rUserOwnedDrydockNotePublicKeys.AsciiArmoredPgpPublicKey: expected string")
								}
							}
							if _, ok := objval["comment"]; ok {
								if s, ok := objval["comment"].(string); ok {
									rUserOwnedDrydockNotePublicKeys.Comment = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rUserOwnedDrydockNotePublicKeys.Comment: expected string")
								}
							}
							if _, ok := objval["id"]; ok {
								if s, ok := objval["id"].(string); ok {
									rUserOwnedDrydockNotePublicKeys.Id = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rUserOwnedDrydockNotePublicKeys.Id: expected string")
								}
							}
							if _, ok := objval["pkixPublicKey"]; ok {
								if rUserOwnedDrydockNotePublicKeysPkixPublicKey, ok := objval["pkixPublicKey"].(map[string]interface{}); ok {
									rUserOwnedDrydockNotePublicKeys.PkixPublicKey = &dclService.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{}
									if _, ok := rUserOwnedDrydockNotePublicKeysPkixPublicKey["publicKeyPem"]; ok {
										if s, ok := rUserOwnedDrydockNotePublicKeysPkixPublicKey["publicKeyPem"].(string); ok {
											rUserOwnedDrydockNotePublicKeys.PkixPublicKey.PublicKeyPem = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rUserOwnedDrydockNotePublicKeys.PkixPublicKey.PublicKeyPem: expected string")
										}
									}
									if _, ok := rUserOwnedDrydockNotePublicKeysPkixPublicKey["signatureAlgorithm"]; ok {
										if s, ok := rUserOwnedDrydockNotePublicKeysPkixPublicKey["signatureAlgorithm"].(string); ok {
											rUserOwnedDrydockNotePublicKeys.PkixPublicKey.SignatureAlgorithm = dclService.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumRef(s)
										} else {
											return nil, fmt.Errorf("rUserOwnedDrydockNotePublicKeys.PkixPublicKey.SignatureAlgorithm: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rUserOwnedDrydockNotePublicKeys.PkixPublicKey: expected map[string]interface{}")
								}
							}
							r.UserOwnedDrydockNote.PublicKeys = append(r.UserOwnedDrydockNote.PublicKeys, rUserOwnedDrydockNotePublicKeys)
						}
					}
				} else {
					return nil, fmt.Errorf("r.UserOwnedDrydockNote.PublicKeys: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.UserOwnedDrydockNote: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetAttestor(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAttestor(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAttestor(ctx, r)
	if err != nil {
		return nil, err
	}
	return AttestorToUnstructured(r), nil
}

func ListAttestor(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAttestor(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AttestorToUnstructured(r))
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

func ApplyAttestor(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAttestor(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAttestor(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAttestor(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AttestorToUnstructured(r), nil
}

func AttestorHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAttestor(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAttestor(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAttestor(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAttestor(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAttestor(u)
	if err != nil {
		return err
	}
	return c.DeleteAttestor(ctx, r)
}

func AttestorID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAttestor(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Attestor) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"binaryauthorization",
		"Attestor",
		"beta",
	}
}

func (r *Attestor) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Attestor) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Attestor) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Attestor) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Attestor) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Attestor) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Attestor) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAttestor(ctx, config, resource)
}

func (r *Attestor) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAttestor(ctx, config, resource, opts...)
}

func (r *Attestor) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AttestorHasDiff(ctx, config, resource, opts...)
}

func (r *Attestor) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAttestor(ctx, config, resource)
}

func (r *Attestor) ID(resource *unstructured.Resource) (string, error) {
	return AttestorID(resource)
}

func init() {
	unstructured.Register(&Attestor{})
}
