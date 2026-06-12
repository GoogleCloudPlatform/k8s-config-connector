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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ManagedKafkaConsumerGroupIdentity{}
	_ identity.Resource   = &ManagedKafkaConsumerGroup{}
)

var ManagedKafkaConsumerGroupIdentityFormat = gcpurls.Template[ManagedKafkaConsumerGroupIdentity]("managedkafka.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}/consumerGroups/{consumerGroup}")

// +k8s:deepcopy-gen=false
type ManagedKafkaConsumerGroupIdentity struct {
	Project       string
	Location      string
	Cluster       string
	ConsumerGroup string
}

func (i *ManagedKafkaConsumerGroupIdentity) String() string {
	return ManagedKafkaConsumerGroupIdentityFormat.ToString(*i)
}

func (i *ManagedKafkaConsumerGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := ManagedKafkaConsumerGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ManagedKafkaConsumerGroup external=%q was not known (use %s): %w", ref, ManagedKafkaConsumerGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ManagedKafkaConsumerGroup external=%q was not known (use %s)", ref, ManagedKafkaConsumerGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ManagedKafkaConsumerGroupIdentity) Host() string {
	return ManagedKafkaConsumerGroupIdentityFormat.Host()
}

func GetManagedKafkaConsumerGroupSpecIdentity(ctx context.Context, reader client.Reader, obj client.Object) (*ManagedKafkaConsumerGroupIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	consumerGroupObj, ok := obj.(*ManagedKafkaConsumerGroup)
	if !ok {
		consumerGroupObj = &ManagedKafkaConsumerGroup{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, consumerGroupObj); err != nil {
			return nil, fmt.Errorf("error converting to ManagedKafkaConsumerGroup: %w", err)
		}
	}

	if err := consumerGroupObj.Spec.ClusterRef.Normalize(ctx, reader, consumerGroupObj.GetNamespace()); err != nil {
		return nil, err
	}
	clusterExternalRef := consumerGroupObj.Spec.ClusterRef.External
	clusterIdentity := &ClusterIdentity{}
	if err := clusterIdentity.FromExternal(clusterExternalRef); err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	identity := &ManagedKafkaConsumerGroupIdentity{
		Project:       projectID,
		Location:      location,
		Cluster:       clusterIdentity.Cluster,
		ConsumerGroup: resourceID,
	}
	return identity, nil
}

func (obj *ManagedKafkaConsumerGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := GetManagedKafkaConsumerGroupSpecIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ManagedKafkaConsumerGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ManagedKafkaConsumerGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *ManagedKafkaConsumerGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
