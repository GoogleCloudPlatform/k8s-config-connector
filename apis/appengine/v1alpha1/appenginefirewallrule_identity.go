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
	"math"
	"strconv"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getIdentityFromAppEngineFirewallRuleSpec(ctx context.Context, reader client.Reader, obj *AppEngineFirewallRule) (*AppEngineFirewallRuleIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}
	// Parse and canonicalize as int32 here, but defer the [1, MaxInt32] range check to
	// Priority() during Create/Update so pre-existing "0"/negative CRs carrying a finalizer
	// can still resolve identity, observe 404 on Find, and complete deletion.
	parsed, err := strconv.ParseInt(resourceID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid resource ID %q: must be a valid 32-bit integer: %w", resourceID, err)
	}

	var projectID string
	if obj.Spec.Project != nil && *obj.Spec.Project != "" {
		projectID = *obj.Spec.Project
	} else {
		projectID, err = refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
	}

	return &AppEngineFirewallRuleIdentity{
		App:         projectID,
		IngressRule: strconv.FormatInt(parsed, 10),
	}, nil
}

// Priority parses and validates the firewall rule priority in the valid App Engine range [1, MaxInt32].
func (i *AppEngineFirewallRuleIdentity) Priority() (int32, error) {
	if i == nil {
		return 0, fmt.Errorf("missing identity for AppEngineFirewallRule")
	}
	priority, err := strconv.ParseInt(i.IngressRule, 10, 32)
	if err != nil || priority < 1 {
		return 0, fmt.Errorf("invalid priority %q: must be an integer between 1 and %d", i.IngressRule, math.MaxInt32)
	}
	return int32(priority), nil
}
