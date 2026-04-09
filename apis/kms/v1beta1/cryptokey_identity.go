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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KMSCryptoKeyIdentity struct {
	parent *KMSKeyRingIdentity
	id     string
}

func (i *KMSCryptoKeyIdentity) String() string {
	return i.parent.String() + "/cryptoKeys/" + i.id
}

func (i *KMSCryptoKeyIdentity) ID() string {
	return i.id
}

func (i *KMSCryptoKeyIdentity) Parent() *KMSKeyRingIdentity {
	return i.parent
}

// New builds a KMSCryptoKeyIdentity from the Config Connector KMSCryptoKey object.
func NewKMSCryptoKeyIdentity(ctx context.Context, reader client.Reader, obj *KMSCryptoKey) (*KMSCryptoKeyIdentity, error) {
	// Get Parent
	kmsKeyRingExternal, err := obj.Spec.KeyRingRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, err
	}
	kmsKeyRing, err := ParseKMSKeyRingExternal(kmsKeyRingExternal)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity, err := ParseKMSCryptoKeyExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.Parent().String() != kmsKeyRing.String() {
			return nil, fmt.Errorf("spec.keyRingRef changed, expect %s, got %s", actualIdentity.Parent().String(), kmsKeyRing.String())
		}
		if actualIdentity.ID() != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.ID())
		}
	}
	return &KMSCryptoKeyIdentity{
		parent: kmsKeyRing,
		id:     resourceID,
	}, nil
}

func ParseKMSCryptoKeyExternal(external string) (*KMSCryptoKeyIdentity, error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	// projects/{{projectId}}/locations/{{location}}/keyRings/{{keyRingId}}/cryptoKeys/{{key}}
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" && tokens[6] == "cryptoKeys" {
		return &KMSCryptoKeyIdentity{parent: &KMSKeyRingIdentity{
			Parent: &parent.ProjectAndLocationParent{
				ProjectID: tokens[1], Location: tokens[3],
			}, ID: tokens[5],
		}, id: tokens[7]}, nil
	}
	return nil, fmt.Errorf("format of KMSCryptoKey external=%q was not known (use projects/{{projectId}}/locations/{{location}}/keyRings/{{keyRingId}}/cryptoKeys/{{keyId}})", external)
}
