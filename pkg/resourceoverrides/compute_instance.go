// Copyright 2022 Google LLC
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

package resourceoverrides

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var (
	networkInterfacePath  = []string{"networkInterface"}
	networkIpFieldPath    = []string{"networkIp"}
	networkIpRefFieldPath = []string{"networkIpRef"}
	supportedKinds        = []string{"ComputeAddress"}
)

func GetComputeInstanceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ComputeInstance",
	}
	ro.Overrides = append(ro.Overrides, addNetworkIpRefField())
	return ro
}

func addNetworkIpRefField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		if err := PreserveMutuallyExclusiveNonReferenceField(crd, networkInterfacePath, networkIpRefFieldPath[0], networkIpFieldPath[0]); err != nil {
			return err
		}
		if err := EnsureReferenceFieldIsMultiKind(crd, networkInterfacePath, networkIpRefFieldPath[0], supportedKinds); err != nil {
			return err
		}

		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorReferenceFieldOverNonReferenceFieldUnderSlice(r, networkInterfacePath, networkIpFieldPath, networkIpRefFieldPath); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", strings.Join(networkIpFieldPath, "."), strings.Join(networkIpRefFieldPath, "."), err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource) error {
		if err := PreserveUserSpecifiedLegacyFieldUnderSlice(original, reconciled, networkInterfacePath, networkIpFieldPath); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(networkIpFieldPath, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecifiedUnderSlice(original, reconciled, networkInterfacePath, networkIpFieldPath, networkIpRefFieldPath); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(networkIpRefFieldPath, "."), err)
		}
		return nil
	}
	return o
}
