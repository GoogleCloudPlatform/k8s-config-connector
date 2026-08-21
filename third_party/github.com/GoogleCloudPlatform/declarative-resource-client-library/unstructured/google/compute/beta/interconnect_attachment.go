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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type InterconnectAttachment struct{}

func InterconnectAttachmentToUnstructured(r *dclService.InterconnectAttachment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "beta",
			Type:    "InterconnectAttachment",
		},
		Object: make(map[string]interface{}),
	}
	if r.AdminEnabled != nil {
		u.Object["adminEnabled"] = *r.AdminEnabled
	}
	if r.Bandwidth != nil {
		u.Object["bandwidth"] = string(*r.Bandwidth)
	}
	var rCandidateSubnets []interface{}
	for _, rCandidateSubnetsVal := range r.CandidateSubnets {
		rCandidateSubnets = append(rCandidateSubnets, rCandidateSubnetsVal)
	}
	u.Object["candidateSubnets"] = rCandidateSubnets
	if r.CloudRouterIPAddress != nil {
		u.Object["cloudRouterIPAddress"] = *r.CloudRouterIPAddress
	}
	if r.CustomerRouterIPAddress != nil {
		u.Object["customerRouterIPAddress"] = *r.CustomerRouterIPAddress
	}
	if r.DataplaneVersion != nil {
		u.Object["dataplaneVersion"] = *r.DataplaneVersion
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EdgeAvailabilityDomain != nil {
		u.Object["edgeAvailabilityDomain"] = string(*r.EdgeAvailabilityDomain)
	}
	if r.Encryption != nil {
		u.Object["encryption"] = string(*r.Encryption)
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Interconnect != nil {
		u.Object["interconnect"] = *r.Interconnect
	}
	var rIpsecInternalAddresses []interface{}
	for _, rIpsecInternalAddressesVal := range r.IpsecInternalAddresses {
		rIpsecInternalAddresses = append(rIpsecInternalAddresses, rIpsecInternalAddressesVal)
	}
	u.Object["ipsecInternalAddresses"] = rIpsecInternalAddresses
	if r.LabelFingerprint != nil {
		u.Object["labelFingerprint"] = *r.LabelFingerprint
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Mtu != nil {
		u.Object["mtu"] = *r.Mtu
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.OperationalStatus != nil {
		u.Object["operationalStatus"] = string(*r.OperationalStatus)
	}
	if r.PairingKey != nil {
		u.Object["pairingKey"] = *r.PairingKey
	}
	if r.PartnerAsn != nil {
		u.Object["partnerAsn"] = *r.PartnerAsn
	}
	if r.PartnerMetadata != nil && r.PartnerMetadata != dclService.EmptyInterconnectAttachmentPartnerMetadata {
		rPartnerMetadata := make(map[string]interface{})
		if r.PartnerMetadata.InterconnectName != nil {
			rPartnerMetadata["interconnectName"] = *r.PartnerMetadata.InterconnectName
		}
		if r.PartnerMetadata.PartnerName != nil {
			rPartnerMetadata["partnerName"] = *r.PartnerMetadata.PartnerName
		}
		if r.PartnerMetadata.PortalUrl != nil {
			rPartnerMetadata["portalUrl"] = *r.PartnerMetadata.PortalUrl
		}
		u.Object["partnerMetadata"] = rPartnerMetadata
	}
	if r.PrivateInterconnectInfo != nil && r.PrivateInterconnectInfo != dclService.EmptyInterconnectAttachmentPrivateInterconnectInfo {
		rPrivateInterconnectInfo := make(map[string]interface{})
		if r.PrivateInterconnectInfo.Tag8021q != nil {
			rPrivateInterconnectInfo["tag8021q"] = *r.PrivateInterconnectInfo.Tag8021q
		}
		u.Object["privateInterconnectInfo"] = rPrivateInterconnectInfo
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.Router != nil {
		u.Object["router"] = *r.Router
	}
	if r.SatisfiesPzs != nil {
		u.Object["satisfiesPzs"] = *r.SatisfiesPzs
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.VlanTag8021q != nil {
		u.Object["vlanTag8021q"] = *r.VlanTag8021q
	}
	return u
}

func UnstructuredToInterconnectAttachment(u *unstructured.Resource) (*dclService.InterconnectAttachment, error) {
	r := &dclService.InterconnectAttachment{}
	if _, ok := u.Object["adminEnabled"]; ok {
		if b, ok := u.Object["adminEnabled"].(bool); ok {
			r.AdminEnabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.AdminEnabled: expected bool")
		}
	}
	if _, ok := u.Object["bandwidth"]; ok {
		if s, ok := u.Object["bandwidth"].(string); ok {
			r.Bandwidth = dclService.InterconnectAttachmentBandwidthEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Bandwidth: expected string")
		}
	}
	if _, ok := u.Object["candidateSubnets"]; ok {
		if s, ok := u.Object["candidateSubnets"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.CandidateSubnets = append(r.CandidateSubnets, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.CandidateSubnets: expected []interface{}")
		}
	}
	if _, ok := u.Object["cloudRouterIPAddress"]; ok {
		if s, ok := u.Object["cloudRouterIPAddress"].(string); ok {
			r.CloudRouterIPAddress = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CloudRouterIPAddress: expected string")
		}
	}
	if _, ok := u.Object["customerRouterIPAddress"]; ok {
		if s, ok := u.Object["customerRouterIPAddress"].(string); ok {
			r.CustomerRouterIPAddress = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CustomerRouterIPAddress: expected string")
		}
	}
	if _, ok := u.Object["dataplaneVersion"]; ok {
		if i, ok := u.Object["dataplaneVersion"].(int64); ok {
			r.DataplaneVersion = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.DataplaneVersion: expected int64")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["edgeAvailabilityDomain"]; ok {
		if s, ok := u.Object["edgeAvailabilityDomain"].(string); ok {
			r.EdgeAvailabilityDomain = dclService.InterconnectAttachmentEdgeAvailabilityDomainEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.EdgeAvailabilityDomain: expected string")
		}
	}
	if _, ok := u.Object["encryption"]; ok {
		if s, ok := u.Object["encryption"].(string); ok {
			r.Encryption = dclService.InterconnectAttachmentEncryptionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Encryption: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if i, ok := u.Object["id"].(int64); ok {
			r.Id = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Id: expected int64")
		}
	}
	if _, ok := u.Object["interconnect"]; ok {
		if s, ok := u.Object["interconnect"].(string); ok {
			r.Interconnect = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Interconnect: expected string")
		}
	}
	if _, ok := u.Object["ipsecInternalAddresses"]; ok {
		if s, ok := u.Object["ipsecInternalAddresses"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.IpsecInternalAddresses = append(r.IpsecInternalAddresses, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.IpsecInternalAddresses: expected []interface{}")
		}
	}
	if _, ok := u.Object["labelFingerprint"]; ok {
		if s, ok := u.Object["labelFingerprint"].(string); ok {
			r.LabelFingerprint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LabelFingerprint: expected string")
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
	if _, ok := u.Object["mtu"]; ok {
		if i, ok := u.Object["mtu"].(int64); ok {
			r.Mtu = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Mtu: expected int64")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["operationalStatus"]; ok {
		if s, ok := u.Object["operationalStatus"].(string); ok {
			r.OperationalStatus = dclService.InterconnectAttachmentOperationalStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.OperationalStatus: expected string")
		}
	}
	if _, ok := u.Object["pairingKey"]; ok {
		if s, ok := u.Object["pairingKey"].(string); ok {
			r.PairingKey = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PairingKey: expected string")
		}
	}
	if _, ok := u.Object["partnerAsn"]; ok {
		if i, ok := u.Object["partnerAsn"].(int64); ok {
			r.PartnerAsn = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.PartnerAsn: expected int64")
		}
	}
	if _, ok := u.Object["partnerMetadata"]; ok {
		if rPartnerMetadata, ok := u.Object["partnerMetadata"].(map[string]interface{}); ok {
			r.PartnerMetadata = &dclService.InterconnectAttachmentPartnerMetadata{}
			if _, ok := rPartnerMetadata["interconnectName"]; ok {
				if s, ok := rPartnerMetadata["interconnectName"].(string); ok {
					r.PartnerMetadata.InterconnectName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PartnerMetadata.InterconnectName: expected string")
				}
			}
			if _, ok := rPartnerMetadata["partnerName"]; ok {
				if s, ok := rPartnerMetadata["partnerName"].(string); ok {
					r.PartnerMetadata.PartnerName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PartnerMetadata.PartnerName: expected string")
				}
			}
			if _, ok := rPartnerMetadata["portalUrl"]; ok {
				if s, ok := rPartnerMetadata["portalUrl"].(string); ok {
					r.PartnerMetadata.PortalUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.PartnerMetadata.PortalUrl: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PartnerMetadata: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["privateInterconnectInfo"]; ok {
		if rPrivateInterconnectInfo, ok := u.Object["privateInterconnectInfo"].(map[string]interface{}); ok {
			r.PrivateInterconnectInfo = &dclService.InterconnectAttachmentPrivateInterconnectInfo{}
			if _, ok := rPrivateInterconnectInfo["tag8021q"]; ok {
				if i, ok := rPrivateInterconnectInfo["tag8021q"].(int64); ok {
					r.PrivateInterconnectInfo.Tag8021q = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.PrivateInterconnectInfo.Tag8021q: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PrivateInterconnectInfo: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
		}
	}
	if _, ok := u.Object["router"]; ok {
		if s, ok := u.Object["router"].(string); ok {
			r.Router = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Router: expected string")
		}
	}
	if _, ok := u.Object["satisfiesPzs"]; ok {
		if b, ok := u.Object["satisfiesPzs"].(bool); ok {
			r.SatisfiesPzs = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.SatisfiesPzs: expected bool")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.InterconnectAttachmentStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.InterconnectAttachmentTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	if _, ok := u.Object["vlanTag8021q"]; ok {
		if i, ok := u.Object["vlanTag8021q"].(int64); ok {
			r.VlanTag8021q = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.VlanTag8021q: expected int64")
		}
	}
	return r, nil
}

func GetInterconnectAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInterconnectAttachment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetInterconnectAttachment(ctx, r)
	if err != nil {
		return nil, err
	}
	return InterconnectAttachmentToUnstructured(r), nil
}

func ListInterconnectAttachment(ctx context.Context, config *dcl.Config, project string, region string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListInterconnectAttachment(ctx, project, region)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, InterconnectAttachmentToUnstructured(r))
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

func ApplyInterconnectAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInterconnectAttachment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInterconnectAttachment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyInterconnectAttachment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return InterconnectAttachmentToUnstructured(r), nil
}

func InterconnectAttachmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInterconnectAttachment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInterconnectAttachment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyInterconnectAttachment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteInterconnectAttachment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInterconnectAttachment(u)
	if err != nil {
		return err
	}
	return c.DeleteInterconnectAttachment(ctx, r)
}

func InterconnectAttachmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToInterconnectAttachment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *InterconnectAttachment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"InterconnectAttachment",
		"beta",
	}
}

func (r *InterconnectAttachment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InterconnectAttachment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InterconnectAttachment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *InterconnectAttachment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InterconnectAttachment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InterconnectAttachment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InterconnectAttachment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetInterconnectAttachment(ctx, config, resource)
}

func (r *InterconnectAttachment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyInterconnectAttachment(ctx, config, resource, opts...)
}

func (r *InterconnectAttachment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return InterconnectAttachmentHasDiff(ctx, config, resource, opts...)
}

func (r *InterconnectAttachment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteInterconnectAttachment(ctx, config, resource)
}

func (r *InterconnectAttachment) ID(resource *unstructured.Resource) (string, error) {
	return InterconnectAttachmentID(resource)
}

func init() {
	unstructured.Register(&InterconnectAttachment{})
}
