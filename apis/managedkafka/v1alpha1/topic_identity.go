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
	_ identity.IdentityV2 = &TopicIdentity{}
	_ identity.Resource   = &ManagedKafkaTopic{}
)

var TopicIdentityFormat = gcpurls.Template[TopicIdentity]("managedkafka.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}/topics/{topic}")

// +k8s:deepcopy-gen=false
type TopicIdentity struct {
	Project  string
	Location string
	Cluster  string
	Topic    string
}

func (i *TopicIdentity) String() string {
	return TopicIdentityFormat.ToString(*i)
}

func (i *TopicIdentity) FromExternal(ref string) error {
	parsed, match, err := TopicIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ManagedKafkaTopic external=%q was not known (use %s): %w", ref, TopicIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ManagedKafkaTopic external=%q was not known (use %s)", ref, TopicIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TopicIdentity) Host() string {
	return TopicIdentityFormat.Host()
}

func getIdentityFromManagedKafkaTopicSpec(ctx context.Context, reader client.Reader, obj client.Object) (*TopicIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	topicObj, ok := obj.(*ManagedKafkaTopic)
	if !ok {
		topicObj = &ManagedKafkaTopic{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.(*unstructured.Unstructured).Object, topicObj); err != nil {
			return nil, fmt.Errorf("error converting to ManagedKafkaTopic: %w", err)
		}
	}

	if err := topicObj.Spec.ClusterRef.Normalize(ctx, reader, topicObj.GetNamespace()); err != nil {
		return nil, err
	}
	clusterExternalRef := topicObj.Spec.ClusterRef.External
	clusterIdentity := &ClusterIdentity{}
	if err := clusterIdentity.FromExternal(clusterExternalRef); err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	identity := &TopicIdentity{
		Project:  projectID,
		Location: location,
		Cluster:  clusterIdentity.Cluster,
		Topic:    resourceID,
	}
	return identity, nil
}

func (obj *ManagedKafkaTopic) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromManagedKafkaTopicSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &TopicIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ManagedKafkaTopic identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
