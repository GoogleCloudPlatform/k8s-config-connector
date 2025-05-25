// Copyright 2024 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PubSubTopicRef struct {
	// If provided must be in the format `projects/[project_id]/topics/[topic_id]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `PubSubTopic` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `PubSubTopic` resource.
	Namespace string `json:"namespace,omitempty"`
}

type PubSubTopic struct {
	projectID string
	topicID   string
}

func ResolvePubSubTopic(ctx context.Context, reader client.Reader, src client.Object, ref *PubSubTopicRef) (*PubSubTopic, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on PubSubTopicRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on PubSubTopicRef")
	}

	// External is provided.
	if ref.External != "" {
		// External should be in the `projects/[project_id]/topics/[topic_id]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "topics" {
			return &PubSubTopic{
				projectID: tokens[1],
				topicID:   tokens[3],
			}, nil
		}
		return nil, fmt.Errorf("format of PubSubTopicRef external=%q was not known (use projects/[project_id]/topics/[topic_id])", ref.External)
	}

	// Fetch PubSubTopic object to construct the external form.
	pubSubTopic := &unstructured.Unstructured{}
	pubSubTopic.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "pubsub.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "PubSubTopic",
	})
	nn := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = src.GetNamespace()
	}
	if err := reader.Get(ctx, nn, pubSubTopic); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced PubSubTopic %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced PubSubTopic %v: %w", nn, err)
	}
	resource, err := k8s.NewResource(pubSubTopic)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(pubSubTopic.GroupVersionKind(), nn)
	}

	projectID, err := ResolveProjectID(ctx, reader, pubSubTopic)
	if err != nil {
		return nil, err
	}
	topicID, err := GetResourceID(pubSubTopic)
	if err != nil {
		return nil, err
	}
	return &PubSubTopic{
		projectID: projectID,
		topicID:   topicID,
	}, nil
}

func (t *PubSubTopic) String() string {
	return fmt.Sprintf("projects/%s/topics/%s", t.projectID, t.topicID)
}

type PubSubSubscriptionRef struct {
	// If provided must be in the format `projects/[project_id]/subscriptions/[subscription_id]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `PubSubSubscription` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `PubSubSubscription` resource.
	Namespace string `json:"namespace,omitempty"`
}

type PubSubSubscription struct {
	projectID string
	subID     string
}

func ResolvePubSubSubscription(ctx context.Context, reader client.Reader, src client.Object, ref *PubSubSubscriptionRef) (*PubSubSubscription, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on PubSubSubscriptionRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on PubSubSubscriptionRef")
	}

	// External is provided.
	if ref.External != "" {
		// External should be in the `projects/[project_id]/subscriptions/[subscriptions_id]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "subscriptions" {
			return &PubSubSubscription{
				projectID: tokens[1],
				subID:     tokens[3],
			}, nil
		}
		return nil, fmt.Errorf("format of PubSubSubscriptionRef external=%q was not known (use projects/[project_id]/subscriptions/[subscriptions_id])", ref.External)
	}

	// Fetch PubSubSubscription object to construct the external form.
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "pubsub.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "PubSubSubscription",
	})
	nn := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = src.GetNamespace()
	}
	if err := reader.Get(ctx, nn, u); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced PubSubSubscription %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced PubSubSubscription %v: %w", nn, err)
	}
	projectID, err := ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}
	subID, err := GetResourceID(u)
	if err != nil {
		return nil, err
	}
	return &PubSubSubscription{
		projectID: projectID,
		subID:     subID,
	}, nil
}

func (t *PubSubSubscription) String() string {
	return fmt.Sprintf("projects/%s/subscriptions/%s", t.projectID, t.subID)
}
