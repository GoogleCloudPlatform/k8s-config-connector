// Copyright 2025 Google LLC
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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	loggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &LinkIdentity{}

// LinkIdentity defines the resource reference to LoggingLink, which "External" field
// holds the GCP identifier for the KRM object.
type LinkIdentity struct {
	parent *loggingv1beta1.LogBucketIdentity
	id     string
}

func (i *LinkIdentity) String() string {
	return i.parent.String() + "/links/" + i.id
}

func (i *LinkIdentity) ID() string {
	return i.id
}

func (i *LinkIdentity) Parent() *loggingv1beta1.LogBucketIdentity {
	return i.parent
}

func (i *LinkIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/links/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of LoggingLink external=%q was not known (use projects/{{projectID}}/locations/{{location}}/buckets/{{bucketID}}/links/{{linkID}})", ref)
	}
	i.parent = &loggingv1beta1.LogBucketIdentity{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("linkID was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &LoggingLink{}

func (obj *LoggingLink) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &LinkIdentity{}

	// Resolve Parent
	if err := obj.Spec.ParentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	newIdentity.parent = &loggingv1beta1.LogBucketIdentity{}
	if err := newIdentity.parent.FromExternal(obj.Spec.ParentRef.GetExternal()); err != nil {
		return nil, fmt.Errorf("parsing parentRef.external=%q: %w", obj.Spec.ParentRef.GetExternal(), err)
	}

	// Get desired ID
	newIdentity.id = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.id == "" {
		newIdentity.id = obj.GetName()
	}
	if newIdentity.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &LinkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}

