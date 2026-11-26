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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &AgentRegistryPublisherIdentity{}
)

var (
	AgentRegistryPublisherIdentityFormat = gcpurls.Template[AgentRegistryPublisherIdentity]("agentregistry.googleapis.com", "projects/{project}/locations/{location}/publishers/{publisher}")
)

// AgentRegistryPublisherIdentity is the identity of a GCP AgentRegistryPublisher resource.
// +k8s:deepcopy-gen=false
type AgentRegistryPublisherIdentity struct {
	Project   string
	Location  string
	Publisher string
}

func (i *AgentRegistryPublisherIdentity) String() string {
	return AgentRegistryPublisherIdentityFormat.ToString(*i)
}

func (i *AgentRegistryPublisherIdentity) Host() string {
	return AgentRegistryPublisherIdentityFormat.Host()
}

func (i *AgentRegistryPublisherIdentity) FromExternal(ref string) error {
	parsed, match, err := AgentRegistryPublisherIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AgentRegistryPublisher external=%q was not known (use %s): %w", ref, AgentRegistryPublisherIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AgentRegistryPublisher external=%q was not known (use %s)", ref, AgentRegistryPublisherIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}
