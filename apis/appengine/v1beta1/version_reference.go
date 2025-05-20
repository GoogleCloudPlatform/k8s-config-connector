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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &AppEngineVersionRef{}
var AppEngineVersionGVK = GroupVersion.WithKind("AppEngineVersion")

// AppEngineVersionRef defines the resource reference to AppEngineVersion, which "External" field
// holds the GCP identifier for the KRM object.
type AppEngineVersionRef struct {
	// The value of an externally managed AppEngineVersion resource.
	// Should be in the format "projects/{{projectID}}/iap_web/appengine-{{appID}}/service/{{service_id}}/version/{{versionID}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	The name of a AppEngineVersion resource.

	Name string `json:"name,omitempty"`

	// The namespace of a AppEngineVersion resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

// NormalizedExternal provision the "External" value for other resource that depends on AppEngineVersion.
// If the "External" is given in the other resource's spec.AppEngineVersionRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual AppEngineVersion object from the cluster.
func (r *AppEngineVersionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" {
		return "", fmt.Errorf("invalid AppEngineVersionRef: must specify external")
	}
	return r.External, nil
}
