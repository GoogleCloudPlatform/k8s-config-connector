// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	"strconv"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GlobalNetworkEndpointIdentity defines the resource reference to ComputeGlobalNetworkEndpoint.
// The identity uses a composite key of (project, neg, port) + optional fqdn/ip.
type GlobalNetworkEndpointIdentity struct {
	parent *GlobalNetworkEndpointParent
	// endpoint identifier: port (required), plus fqdn or ipAddress
	port      int32
	fqdn      string
	ipAddress string
}

// GlobalNetworkEndpointParent holds the parent fields for a GlobalNetworkEndpoint.
type GlobalNetworkEndpointParent struct {
	ProjectID                  string
	GlobalNetworkEndpointGroup string
}

func (p *GlobalNetworkEndpointParent) String() string {
	return "projects/" + p.ProjectID + "/global/networkEndpointGroups/" + p.GlobalNetworkEndpointGroup
}

// String returns the canonical external reference string.
// Format: projects/{project}/global/networkEndpointGroups/{neg}/endpoints/{port}/fqdn/{fqdn}
// or: projects/{project}/global/networkEndpointGroups/{neg}/endpoints/{port}/ipAddress/{ip}
func (i *GlobalNetworkEndpointIdentity) String() string {
	base := i.parent.String() + "/endpoints/" + strconv.Itoa(int(i.port))
	if i.fqdn != "" {
		return base + "/fqdn/" + i.fqdn
	}
	return base + "/ipAddress/" + i.ipAddress
}

func (i *GlobalNetworkEndpointIdentity) Parent() *GlobalNetworkEndpointParent {
	return i.parent
}

func (i *GlobalNetworkEndpointIdentity) Port() int32 {
	return i.port
}

func (i *GlobalNetworkEndpointIdentity) Fqdn() string {
	return i.fqdn
}

func (i *GlobalNetworkEndpointIdentity) IPAddress() string {
	return i.ipAddress
}

// NewGlobalNetworkEndpointIdentity builds an identity from the Config Connector object.
func NewGlobalNetworkEndpointIdentity(ctx context.Context, reader client.Reader, obj *ComputeGlobalNetworkEndpoint) (*GlobalNetworkEndpointIdentity, error) {
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	neg := obj.Spec.GlobalNetworkEndpointGroup
	if neg == "" {
		return nil, fmt.Errorf("spec.globalNetworkEndpointGroup is required")
	}

	port := obj.Spec.Port
	if port == 0 {
		return nil, fmt.Errorf("spec.port is required")
	}

	fqdn := ""
	if obj.Spec.Fqdn != nil {
		fqdn = *obj.Spec.Fqdn
	}
	ipAddress := ""
	if obj.Spec.IPAddress != nil {
		ipAddress = *obj.Spec.IPAddress
	}

	if fqdn == "" && ipAddress == "" {
		return nil, fmt.Errorf("exactly one of spec.fqdn or spec.ipAddress must be specified")
	}
	if fqdn != "" && ipAddress != "" {
		return nil, fmt.Errorf("exactly one of spec.fqdn or spec.ipAddress must be specified, not both")
	}

	return &GlobalNetworkEndpointIdentity{
		parent: &GlobalNetworkEndpointParent{
			ProjectID:                  projectID,
			GlobalNetworkEndpointGroup: neg,
		},
		port:      port,
		fqdn:      fqdn,
		ipAddress: ipAddress,
	}, nil
}

// ParseGlobalNetworkEndpointExternal parses an external reference string.
// Expected formats:
//
//	projects/{project}/global/networkEndpointGroups/{neg}/endpoints/{port}/fqdn/{fqdn}
//	projects/{project}/global/networkEndpointGroups/{neg}/endpoints/{port}/ipAddress/{ip}
//
// Token indices (0-based): 0=projects 1={project} 2=global 3=networkEndpointGroups 4={neg} 5=endpoints 6={port} 7=fqdn|ipAddress 8={value}
func ParseGlobalNetworkEndpointExternal(external string) (*GlobalNetworkEndpointIdentity, error) {
	tokens := strings.Split(external, "/")
	// 9 tokens: projects / {project} / global / networkEndpointGroups / {neg} / endpoints / {port} / fqdn|ipAddress / {value}
	if len(tokens) != 9 {
		return nil, fmt.Errorf("format of ComputeGlobalNetworkEndpoint external=%q was not known (expected 9 tokens, got %d)", external, len(tokens))
	}
	if tokens[0] != "projects" || tokens[2] != "global" || tokens[3] != "networkEndpointGroups" || tokens[5] != "endpoints" {
		return nil, fmt.Errorf("format of ComputeGlobalNetworkEndpoint external=%q was not known", external)
	}

	portInt, err := strconv.Atoi(tokens[6])
	if err != nil {
		return nil, fmt.Errorf("invalid port in external ref %q: %w", external, err)
	}

	id := &GlobalNetworkEndpointIdentity{
		parent: &GlobalNetworkEndpointParent{
			ProjectID:                  tokens[1],
			GlobalNetworkEndpointGroup: tokens[4],
		},
		port: int32(portInt),
	}

	switch tokens[7] {
	case "fqdn":
		id.fqdn = tokens[8]
	case "ipAddress":
		id.ipAddress = tokens[8]
	default:
		return nil, fmt.Errorf("format of ComputeGlobalNetworkEndpoint external=%q was not known (unknown key %q)", external, tokens[7])
	}

	return id, nil
}
