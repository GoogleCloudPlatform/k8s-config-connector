// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/tier2/alpha/tier2_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tier2/alpha"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceSkuTierEnum converts a InstanceSkuTierEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceSkuTierEnum(e alphapb.Tier2AlphaInstanceSkuTierEnum) *alpha.InstanceSkuTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceSkuTierEnum_name[int32(e)]; ok {
		e := alpha.InstanceSkuTierEnum(n[len("Tier2AlphaInstanceSkuTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSkuSizeEnum converts a InstanceSkuSizeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceSkuSizeEnum(e alphapb.Tier2AlphaInstanceSkuSizeEnum) *alpha.InstanceSkuSizeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceSkuSizeEnum_name[int32(e)]; ok {
		e := alpha.InstanceSkuSizeEnum(n[len("Tier2AlphaInstanceSkuSizeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceStateEnum(e alphapb.Tier2AlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("Tier2AlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum converts a InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(e alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum) *alpha.InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(n[len("Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessCreateRecipeStepsActionEnum converts a InstancePreprocessCreateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum) *alpha.InstancePreprocessCreateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessCreateRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceCreateRecipeStepsActionEnum converts a InstanceCreateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum) *alpha.InstanceCreateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceCreateRecipeStepsActionEnum(n[len("Tier2AlphaInstanceCreateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDeleteRecipeStepsActionEnum converts a InstanceDeleteRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum) *alpha.InstanceDeleteRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceDeleteRecipeStepsActionEnum(n[len("Tier2AlphaInstanceDeleteRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUpdateRecipeStepsActionEnum converts a InstanceUpdateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum) *alpha.InstanceUpdateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceUpdateRecipeStepsActionEnum(n[len("Tier2AlphaInstanceUpdateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessResetRecipeStepsActionEnum converts a InstancePreprocessResetRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum) *alpha.InstancePreprocessResetRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessResetRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceResetRecipeStepsActionEnum converts a InstanceResetRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum) *alpha.InstanceResetRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceResetRecipeStepsActionEnum(n[len("Tier2AlphaInstanceResetRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceResetRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceResetRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceResetRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessRepairRecipeStepsActionEnum converts a InstancePreprocessRepairRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum) *alpha.InstancePreprocessRepairRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessRepairRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceRepairRecipeStepsActionEnum converts a InstanceRepairRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum) *alpha.InstanceRepairRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceRepairRecipeStepsActionEnum(n[len("Tier2AlphaInstanceRepairRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessDeleteRecipeStepsActionEnum converts a InstancePreprocessDeleteRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum) *alpha.InstancePreprocessDeleteRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessDeleteRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUpdateRecipeStepsActionEnum converts a InstancePreprocessUpdateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum) *alpha.InstancePreprocessUpdateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUpdateRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessFreezeRecipeStepsActionEnum converts a InstancePreprocessFreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum) *alpha.InstancePreprocessFreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessFreezeRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFreezeRecipeStepsActionEnum converts a InstanceFreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum) *alpha.InstanceFreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceFreezeRecipeStepsActionEnum(n[len("Tier2AlphaInstanceFreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsActionEnum converts a InstancePreprocessUnfreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum) *alpha.InstancePreprocessUnfreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUnfreezeRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUnfreezeRecipeStepsActionEnum converts a InstanceUnfreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum) *alpha.InstanceUnfreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceUnfreezeRecipeStepsActionEnum(n[len("Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsActionEnum converts a InstancePreprocessReportInstanceHealthRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessReportInstanceHealthRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReportInstanceHealthRecipeStepsActionEnum converts a InstanceReportInstanceHealthRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum) *alpha.InstanceReportInstanceHealthRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceReportInstanceHealthRecipeStepsActionEnum(n[len("Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessGetRecipeStepsActionEnum converts a InstancePreprocessGetRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum) *alpha.InstancePreprocessGetRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessGetRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsActionEnum converts a InstanceNotifyKeyAvailableRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum) *alpha.InstanceNotifyKeyAvailableRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceNotifyKeyAvailableRecipeStepsActionEnum(n[len("Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsActionEnum converts a InstanceNotifyKeyUnavailableRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum) *alpha.InstanceNotifyKeyUnavailableRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceNotifyKeyUnavailableRecipeStepsActionEnum(n[len("Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReadonlyRecipeStepsActionEnum converts a InstanceReadonlyRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum) *alpha.InstanceReadonlyRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceReadonlyRecipeStepsActionEnum(n[len("Tier2AlphaInstanceReadonlyRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReconcileRecipeStepsActionEnum converts a InstanceReconcileRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum) *alpha.InstanceReconcileRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceReconcileRecipeStepsActionEnum(n[len("Tier2AlphaInstanceReconcileRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessPassthroughRecipeStepsActionEnum converts a InstancePreprocessPassthroughRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum) *alpha.InstancePreprocessPassthroughRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessPassthroughRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessReconcileRecipeStepsActionEnum converts a InstancePreprocessReconcileRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum) *alpha.InstancePreprocessReconcileRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessReconcileRecipeStepsActionEnum(n[len("Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(n[len("Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(e alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum) *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(n[len("Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSku converts a InstanceSku resource from its proto representation.
func ProtoToTier2AlphaInstanceSku(p *alphapb.Tier2AlphaInstanceSku) *alpha.InstanceSku {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceSku{
		Tier: ProtoToTier2AlphaInstanceSkuTierEnum(p.GetTier()),
		Size: ProtoToTier2AlphaInstanceSkuSizeEnum(p.GetSize()),
	}
	return obj
}

// ProtoToInstanceReferences converts a InstanceReferences resource from its proto representation.
func ProtoToTier2AlphaInstanceReferences(p *alphapb.Tier2AlphaInstanceReferences) *alpha.InstanceReferences {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReferences{
		Name:           dcl.StringOrNil(p.Name),
		Type:           dcl.StringOrNil(p.Type),
		SourceResource: dcl.StringOrNil(p.SourceResource),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceReferencesDetails(r))
	}
	return obj
}

// ProtoToInstanceReferencesDetails converts a InstanceReferencesDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceReferencesDetails(p *alphapb.Tier2AlphaInstanceReferencesDetails) *alpha.InstanceReferencesDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReferencesDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceEncryptionKeys converts a InstanceEncryptionKeys resource from its proto representation.
func ProtoToTier2AlphaInstanceEncryptionKeys(p *alphapb.Tier2AlphaInstanceEncryptionKeys) *alpha.InstanceEncryptionKeys {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceEncryptionKeys{
		KeyOrVersion: dcl.StringOrNil(p.KeyOrVersion),
		Grant:        dcl.StringOrNil(p.Grant),
		Delegate:     dcl.StringOrNil(p.Delegate),
		KeyState:     ProtoToTier2AlphaInstanceEncryptionKeysKeyState(p.GetKeyState()),
	}
	return obj
}

// ProtoToInstanceEncryptionKeysKeyState converts a InstanceEncryptionKeysKeyState resource from its proto representation.
func ProtoToTier2AlphaInstanceEncryptionKeysKeyState(p *alphapb.Tier2AlphaInstanceEncryptionKeysKeyState) *alpha.InstanceEncryptionKeysKeyState {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceEncryptionKeysKeyState{
		KeyStateVersion: dcl.Int64OrNil(p.KeyStateVersion),
		Availability:    ProtoToTier2AlphaInstanceEncryptionKeysKeyStateAvailability(p.GetAvailability()),
	}
	return obj
}

// ProtoToInstanceEncryptionKeysKeyStateAvailability converts a InstanceEncryptionKeysKeyStateAvailability resource from its proto representation.
func ProtoToTier2AlphaInstanceEncryptionKeysKeyStateAvailability(p *alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailability) *alpha.InstanceEncryptionKeysKeyStateAvailability {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceEncryptionKeysKeyStateAvailability{
		PermissionDenied: dcl.Bool(p.PermissionDenied),
		UnknownFailure:   dcl.Bool(p.UnknownFailure),
		KeyVersionState:  ProtoToTier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(p.GetKeyVersionState()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipe converts a InstancePreprocessCreateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipe(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipe) *alpha.InstancePreprocessCreateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessCreateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeSteps converts a InstancePreprocessCreateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeSteps) *alpha.InstancePreprocessCreateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsStatus converts a InstancePreprocessCreateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus) *alpha.InstancePreprocessCreateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsStatusDetails converts a InstancePreprocessCreateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails) *alpha.InstancePreprocessCreateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsQuotaRequestDeltas converts a InstancePreprocessCreateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessCreateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPreprocessUpdate converts a InstancePreprocessCreateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessCreateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsRequestedTenantProject converts a InstancePreprocessCreateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfo converts a InstancePreprocessCreateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceGoogleprotobufstruct converts a InstanceGoogleprotobufstruct resource from its proto representation.
func ProtoToTier2AlphaInstanceGoogleprotobufstruct(p *alphapb.Tier2AlphaInstanceGoogleprotobufstruct) *alpha.InstanceGoogleprotobufstruct {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGoogleprotobufstruct{}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoResource converts a InstancePreprocessCreateRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceCreateRecipe converts a InstanceCreateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipe(p *alphapb.Tier2AlphaInstanceCreateRecipe) *alpha.InstanceCreateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceCreateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceCreateRecipeSteps converts a InstanceCreateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeSteps(p *alphapb.Tier2AlphaInstanceCreateRecipeSteps) *alpha.InstanceCreateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceCreateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceCreateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceCreateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsStatus converts a InstanceCreateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatus) *alpha.InstanceCreateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceCreateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsStatusDetails converts a InstanceCreateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatusDetails) *alpha.InstanceCreateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsQuotaRequestDeltas converts a InstanceCreateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas) *alpha.InstanceCreateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPreprocessUpdate converts a InstanceCreateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate) *alpha.InstanceCreateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsRequestedTenantProject converts a InstanceCreateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject) *alpha.InstanceCreateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfo converts a InstanceCreateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo) *alpha.InstanceCreateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoPolicyName converts a InstanceCreateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoIamPermissions converts a InstanceCreateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceCreateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoResource converts a InstanceCreateRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoResource) *alpha.InstanceCreateRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdate converts a InstanceCreateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceDeleteRecipe converts a InstanceDeleteRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipe(p *alphapb.Tier2AlphaInstanceDeleteRecipe) *alpha.InstanceDeleteRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceDeleteRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceDeleteRecipeSteps converts a InstanceDeleteRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeSteps(p *alphapb.Tier2AlphaInstanceDeleteRecipeSteps) *alpha.InstanceDeleteRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceDeleteRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceDeleteRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsStatus converts a InstanceDeleteRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatus) *alpha.InstanceDeleteRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceDeleteRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsStatusDetails converts a InstanceDeleteRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails) *alpha.InstanceDeleteRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsQuotaRequestDeltas converts a InstanceDeleteRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas) *alpha.InstanceDeleteRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPreprocessUpdate converts a InstanceDeleteRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate) *alpha.InstanceDeleteRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsRequestedTenantProject converts a InstanceDeleteRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject) *alpha.InstanceDeleteRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfo converts a InstanceDeleteRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo) *alpha.InstanceDeleteRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoPolicyName converts a InstanceDeleteRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoIamPermissions converts a InstanceDeleteRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceDeleteRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoResource converts a InstanceDeleteRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResource) *alpha.InstanceDeleteRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdate converts a InstanceDeleteRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceUpdateRecipe converts a InstanceUpdateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipe(p *alphapb.Tier2AlphaInstanceUpdateRecipe) *alpha.InstanceUpdateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceUpdateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceUpdateRecipeSteps converts a InstanceUpdateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeSteps(p *alphapb.Tier2AlphaInstanceUpdateRecipeSteps) *alpha.InstanceUpdateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceUpdateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceUpdateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsStatus converts a InstanceUpdateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatus) *alpha.InstanceUpdateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceUpdateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsStatusDetails converts a InstanceUpdateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails) *alpha.InstanceUpdateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsQuotaRequestDeltas converts a InstanceUpdateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas) *alpha.InstanceUpdateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPreprocessUpdate converts a InstanceUpdateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate) *alpha.InstanceUpdateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsRequestedTenantProject converts a InstanceUpdateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject) *alpha.InstanceUpdateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfo converts a InstanceUpdateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo) *alpha.InstanceUpdateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoPolicyName converts a InstanceUpdateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoIamPermissions converts a InstanceUpdateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceUpdateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoResource converts a InstanceUpdateRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResource) *alpha.InstanceUpdateRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdate converts a InstanceUpdateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipe converts a InstancePreprocessResetRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipe(p *alphapb.Tier2AlphaInstancePreprocessResetRecipe) *alpha.InstancePreprocessResetRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessResetRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeSteps converts a InstancePreprocessResetRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeSteps) *alpha.InstancePreprocessResetRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsStatus converts a InstancePreprocessResetRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatus) *alpha.InstancePreprocessResetRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsStatusDetails converts a InstancePreprocessResetRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails) *alpha.InstancePreprocessResetRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsQuotaRequestDeltas converts a InstancePreprocessResetRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessResetRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPreprocessUpdate converts a InstancePreprocessResetRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessResetRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsRequestedTenantProject converts a InstancePreprocessResetRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfo converts a InstancePreprocessResetRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoResource converts a InstancePreprocessResetRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceResetRecipe converts a InstanceResetRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipe(p *alphapb.Tier2AlphaInstanceResetRecipe) *alpha.InstanceResetRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceResetRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceResetRecipeSteps converts a InstanceResetRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeSteps(p *alphapb.Tier2AlphaInstanceResetRecipeSteps) *alpha.InstanceResetRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceResetRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceResetRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceResetRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsStatus converts a InstanceResetRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceResetRecipeStepsStatus) *alpha.InstanceResetRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceResetRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsStatusDetails converts a InstanceResetRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceResetRecipeStepsStatusDetails) *alpha.InstanceResetRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsQuotaRequestDeltas converts a InstanceResetRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas) *alpha.InstanceResetRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPreprocessUpdate converts a InstanceResetRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate) *alpha.InstanceResetRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsRequestedTenantProject converts a InstanceResetRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject) *alpha.InstanceResetRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfo converts a InstanceResetRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo) *alpha.InstanceResetRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoPolicyName converts a InstanceResetRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceResetRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoIamPermissions converts a InstanceResetRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceResetRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoResource converts a InstanceResetRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoResource) *alpha.InstanceResetRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdate converts a InstanceResetRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipe converts a InstancePreprocessRepairRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipe(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipe) *alpha.InstancePreprocessRepairRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessRepairRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeSteps converts a InstancePreprocessRepairRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeSteps) *alpha.InstancePreprocessRepairRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsStatus converts a InstancePreprocessRepairRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus) *alpha.InstancePreprocessRepairRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsStatusDetails converts a InstancePreprocessRepairRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails) *alpha.InstancePreprocessRepairRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsQuotaRequestDeltas converts a InstancePreprocessRepairRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessRepairRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPreprocessUpdate converts a InstancePreprocessRepairRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessRepairRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsRequestedTenantProject converts a InstancePreprocessRepairRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfo converts a InstancePreprocessRepairRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoResource converts a InstancePreprocessRepairRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceRepairRecipe converts a InstanceRepairRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipe(p *alphapb.Tier2AlphaInstanceRepairRecipe) *alpha.InstanceRepairRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceRepairRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceRepairRecipeSteps converts a InstanceRepairRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeSteps(p *alphapb.Tier2AlphaInstanceRepairRecipeSteps) *alpha.InstanceRepairRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceRepairRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceRepairRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceRepairRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsStatus converts a InstanceRepairRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatus) *alpha.InstanceRepairRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceRepairRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsStatusDetails converts a InstanceRepairRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatusDetails) *alpha.InstanceRepairRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsQuotaRequestDeltas converts a InstanceRepairRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas) *alpha.InstanceRepairRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPreprocessUpdate converts a InstanceRepairRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate) *alpha.InstanceRepairRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsRequestedTenantProject converts a InstanceRepairRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject) *alpha.InstanceRepairRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfo converts a InstanceRepairRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo) *alpha.InstanceRepairRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoPolicyName converts a InstanceRepairRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoIamPermissions converts a InstanceRepairRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceRepairRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoResource converts a InstanceRepairRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoResource) *alpha.InstanceRepairRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdate converts a InstanceRepairRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipe converts a InstancePreprocessDeleteRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipe(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipe) *alpha.InstancePreprocessDeleteRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeSteps converts a InstancePreprocessDeleteRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeSteps) *alpha.InstancePreprocessDeleteRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsStatus converts a InstancePreprocessDeleteRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus) *alpha.InstancePreprocessDeleteRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsStatusDetails converts a InstancePreprocessDeleteRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails) *alpha.InstancePreprocessDeleteRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas converts a InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPreprocessUpdate converts a InstancePreprocessDeleteRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessDeleteRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsRequestedTenantProject converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfo converts a InstancePreprocessDeleteRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoResource converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipe converts a InstancePreprocessUpdateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipe(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipe) *alpha.InstancePreprocessUpdateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeSteps converts a InstancePreprocessUpdateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeSteps) *alpha.InstancePreprocessUpdateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsStatus converts a InstancePreprocessUpdateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus) *alpha.InstancePreprocessUpdateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsStatusDetails converts a InstancePreprocessUpdateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails) *alpha.InstancePreprocessUpdateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas converts a InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPreprocessUpdate converts a InstancePreprocessUpdateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessUpdateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsRequestedTenantProject converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfo converts a InstancePreprocessUpdateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoResource converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipe converts a InstancePreprocessFreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipe(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipe) *alpha.InstancePreprocessFreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeSteps converts a InstancePreprocessFreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeSteps) *alpha.InstancePreprocessFreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsStatus converts a InstancePreprocessFreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus) *alpha.InstancePreprocessFreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsStatusDetails converts a InstancePreprocessFreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails) *alpha.InstancePreprocessFreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas converts a InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPreprocessUpdate converts a InstancePreprocessFreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessFreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsRequestedTenantProject converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfo converts a InstancePreprocessFreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoResource converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceFreezeRecipe converts a InstanceFreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipe(p *alphapb.Tier2AlphaInstanceFreezeRecipe) *alpha.InstanceFreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceFreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceFreezeRecipeSteps converts a InstanceFreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeSteps(p *alphapb.Tier2AlphaInstanceFreezeRecipeSteps) *alpha.InstanceFreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceFreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceFreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsStatus converts a InstanceFreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatus) *alpha.InstanceFreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceFreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsStatusDetails converts a InstanceFreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails) *alpha.InstanceFreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsQuotaRequestDeltas converts a InstanceFreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas) *alpha.InstanceFreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPreprocessUpdate converts a InstanceFreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate) *alpha.InstanceFreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsRequestedTenantProject converts a InstanceFreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject) *alpha.InstanceFreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfo converts a InstanceFreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo) *alpha.InstanceFreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoPolicyName converts a InstanceFreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoIamPermissions converts a InstanceFreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceFreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoResource converts a InstanceFreezeRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResource) *alpha.InstanceFreezeRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdate converts a InstanceFreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipe converts a InstancePreprocessUnfreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipe(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipe) *alpha.InstancePreprocessUnfreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeSteps converts a InstancePreprocessUnfreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps) *alpha.InstancePreprocessUnfreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsStatus converts a InstancePreprocessUnfreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus) *alpha.InstancePreprocessUnfreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsStatusDetails converts a InstancePreprocessUnfreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails) *alpha.InstancePreprocessUnfreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas converts a InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate converts a InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfo converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipe converts a InstanceUnfreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipe(p *alphapb.Tier2AlphaInstanceUnfreezeRecipe) *alpha.InstanceUnfreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceUnfreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeSteps converts a InstanceUnfreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeSteps(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeSteps) *alpha.InstanceUnfreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsStatus converts a InstanceUnfreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatus) *alpha.InstanceUnfreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsStatusDetails converts a InstanceUnfreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails) *alpha.InstanceUnfreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsQuotaRequestDeltas converts a InstanceUnfreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas) *alpha.InstanceUnfreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPreprocessUpdate converts a InstanceUnfreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate) *alpha.InstanceUnfreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsRequestedTenantProject converts a InstanceUnfreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject) *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfo converts a InstanceUnfreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName converts a InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions converts a InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoResource converts a InstanceUnfreezeRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResource) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdate converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipe converts a InstancePreprocessReportInstanceHealthRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipe(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipe) *alpha.InstancePreprocessReportInstanceHealthRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeSteps converts a InstancePreprocessReportInstanceHealthRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeSteps) *alpha.InstancePreprocessReportInstanceHealthRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsStatus converts a InstancePreprocessReportInstanceHealthRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatus) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails converts a InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas converts a InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate converts a InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject converts a InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipe converts a InstanceReportInstanceHealthRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipe(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipe) *alpha.InstanceReportInstanceHealthRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceReportInstanceHealthRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeSteps converts a InstanceReportInstanceHealthRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeSteps(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeSteps) *alpha.InstanceReportInstanceHealthRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsStatus converts a InstanceReportInstanceHealthRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatus) *alpha.InstanceReportInstanceHealthRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsStatusDetails converts a InstanceReportInstanceHealthRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetails) *alpha.InstanceReportInstanceHealthRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas converts a InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas) *alpha.InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsPreprocessUpdate converts a InstanceReportInstanceHealthRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdate) *alpha.InstanceReportInstanceHealthRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsRequestedTenantProject converts a InstanceReportInstanceHealthRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProject) *alpha.InstanceReportInstanceHealthRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsPermissionsInfo converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfo) *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource) *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate converts a InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate) *alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipe converts a InstancePreprocessGetRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipe(p *alphapb.Tier2AlphaInstancePreprocessGetRecipe) *alpha.InstancePreprocessGetRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessGetRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeSteps converts a InstancePreprocessGetRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeSteps) *alpha.InstancePreprocessGetRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessGetRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessGetRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsStatus converts a InstancePreprocessGetRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsStatus) *alpha.InstancePreprocessGetRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessGetRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsStatusDetails converts a InstancePreprocessGetRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsStatusDetails) *alpha.InstancePreprocessGetRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsQuotaRequestDeltas converts a InstancePreprocessGetRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessGetRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsPreprocessUpdate converts a InstancePreprocessGetRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessGetRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsRequestedTenantProject converts a InstancePreprocessGetRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessGetRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsPermissionsInfo converts a InstancePreprocessGetRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfo) *alpha.InstancePreprocessGetRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsPermissionsInfoResource converts a InstancePreprocessGetRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessGetRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipe converts a InstanceNotifyKeyAvailableRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipe(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipe) *alpha.InstanceNotifyKeyAvailableRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeSteps converts a InstanceNotifyKeyAvailableRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeSteps(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeSteps) *alpha.InstanceNotifyKeyAvailableRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsStatus converts a InstanceNotifyKeyAvailableRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatus) *alpha.InstanceNotifyKeyAvailableRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsStatusDetails converts a InstanceNotifyKeyAvailableRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetails) *alpha.InstanceNotifyKeyAvailableRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas converts a InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas) *alpha.InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate converts a InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate) *alpha.InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject converts a InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject) *alpha.InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo) *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource) *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate converts a InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate) *alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipe converts a InstanceNotifyKeyUnavailableRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipe(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipe) *alpha.InstanceNotifyKeyUnavailableRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeSteps converts a InstanceNotifyKeyUnavailableRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeSteps(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeSteps) *alpha.InstanceNotifyKeyUnavailableRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsStatus converts a InstanceNotifyKeyUnavailableRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatus) *alpha.InstanceNotifyKeyUnavailableRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsStatusDetails converts a InstanceNotifyKeyUnavailableRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetails) *alpha.InstanceNotifyKeyUnavailableRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas converts a InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas) *alpha.InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate converts a InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate) *alpha.InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject converts a InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject) *alpha.InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo) *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource) *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate converts a InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate) *alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipe converts a InstanceReadonlyRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipe(p *alphapb.Tier2AlphaInstanceReadonlyRecipe) *alpha.InstanceReadonlyRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceReadonlyRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeSteps converts a InstanceReadonlyRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeSteps(p *alphapb.Tier2AlphaInstanceReadonlyRecipeSteps) *alpha.InstanceReadonlyRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceReadonlyRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsStatus converts a InstanceReadonlyRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatus) *alpha.InstanceReadonlyRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsStatusDetails converts a InstanceReadonlyRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails) *alpha.InstanceReadonlyRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsQuotaRequestDeltas converts a InstanceReadonlyRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas) *alpha.InstanceReadonlyRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPreprocessUpdate converts a InstanceReadonlyRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate) *alpha.InstanceReadonlyRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsRequestedTenantProject converts a InstanceReadonlyRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject) *alpha.InstanceReadonlyRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfo converts a InstanceReadonlyRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo) *alpha.InstanceReadonlyRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoPolicyName converts a InstanceReadonlyRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions converts a InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoResource converts a InstanceReadonlyRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResource) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdate converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceReconcileRecipe converts a InstanceReconcileRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipe(p *alphapb.Tier2AlphaInstanceReconcileRecipe) *alpha.InstanceReconcileRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceReconcileRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceReconcileRecipeSteps converts a InstanceReconcileRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeSteps(p *alphapb.Tier2AlphaInstanceReconcileRecipeSteps) *alpha.InstanceReconcileRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceReconcileRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceReconcileRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceReconcileRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceReconcileRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsStatus converts a InstanceReconcileRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsStatus) *alpha.InstanceReconcileRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceReconcileRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsStatusDetails converts a InstanceReconcileRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsStatusDetails) *alpha.InstanceReconcileRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsQuotaRequestDeltas converts a InstanceReconcileRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltas) *alpha.InstanceReconcileRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsPreprocessUpdate converts a InstanceReconcileRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPreprocessUpdate) *alpha.InstanceReconcileRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsRequestedTenantProject converts a InstanceReconcileRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProject) *alpha.InstanceReconcileRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsPermissionsInfo converts a InstanceReconcileRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfo) *alpha.InstanceReconcileRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsPermissionsInfoPolicyName converts a InstanceReconcileRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceReconcileRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsPermissionsInfoIamPermissions converts a InstanceReconcileRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceReconcileRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsPermissionsInfoResource converts a InstanceReconcileRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResource) *alpha.InstanceReconcileRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsKeyNotificationsUpdate converts a InstanceReconcileRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdate) *alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipe converts a InstancePreprocessPassthroughRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipe(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipe) *alpha.InstancePreprocessPassthroughRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessPassthroughRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeSteps converts a InstancePreprocessPassthroughRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeSteps) *alpha.InstancePreprocessPassthroughRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsStatus converts a InstancePreprocessPassthroughRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatus) *alpha.InstancePreprocessPassthroughRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsStatusDetails converts a InstancePreprocessPassthroughRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetails) *alpha.InstancePreprocessPassthroughRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas converts a InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsPreprocessUpdate converts a InstancePreprocessPassthroughRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessPassthroughRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsRequestedTenantProject converts a InstancePreprocessPassthroughRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessPassthroughRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsPermissionsInfo converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfo) *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipe converts a InstancePreprocessReconcileRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipe(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipe) *alpha.InstancePreprocessReconcileRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessReconcileRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeSteps converts a InstancePreprocessReconcileRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeSteps) *alpha.InstancePreprocessReconcileRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
		PublicErrorMessage:             dcl.StringOrNil(p.PublicErrorMessage),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsStatus converts a InstancePreprocessReconcileRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatus) *alpha.InstancePreprocessReconcileRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsStatusDetails converts a InstancePreprocessReconcileRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetails) *alpha.InstancePreprocessReconcileRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas converts a InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.StringOrNil(p.MetricName),
		Amount:            dcl.Int64OrNil(p.Amount),
		QuotaLocationName: dcl.StringOrNil(p.QuotaLocationName),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsPreprocessUpdate converts a InstancePreprocessReconcileRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessReconcileRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsRequestedTenantProject converts a InstancePreprocessReconcileRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessReconcileRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsPermissionsInfo converts a InstancePreprocessReconcileRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfo) *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfo{
		PolicyName:     ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath:   dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:       ProtoToTier2AlphaInstanceGoogleprotobufstruct(p.GetApiAttrs()),
		PolicyNameMode: ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(p.GetPolicyNameMode()),
		Resource:       ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResource(p.GetResource()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsPermissionsInfoResource converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoResource resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResource(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResource) *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoResource {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoResource{
		Name:    dcl.StringOrNil(p.Name),
		Type:    dcl.StringOrNil(p.Type),
		Service: dcl.StringOrNil(p.Service),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	for _, r := range p.GetKeyNotificationConfigs() {
		obj.KeyNotificationConfigs = append(obj.KeyNotificationConfigs, *ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(r))
	}
	return obj
}

// ProtoToInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs converts a InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs(p *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
		Grant:            dcl.StringOrNil(p.Grant),
		DelegatorGaiaId:  dcl.Int64OrNil(p.DelegatorGaiaId),
	}
	return obj
}

// ProtoToInstanceHistory converts a InstanceHistory resource from its proto representation.
func ProtoToTier2AlphaInstanceHistory(p *alphapb.Tier2AlphaInstanceHistory) *alpha.InstanceHistory {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceHistory{
		Timestamp:           dcl.StringOrNil(p.GetTimestamp()),
		OperationHandle:     dcl.StringOrNil(p.OperationHandle),
		Description:         dcl.StringOrNil(p.Description),
		StepIndex:           dcl.Int64OrNil(p.StepIndex),
		TenantProjectNumber: dcl.Int64OrNil(p.TenantProjectNumber),
		TenantProjectId:     dcl.StringOrNil(p.TenantProjectId),
		P4ServiceAccount:    dcl.StringOrNil(p.P4ServiceAccount),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.Tier2AlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:                                 dcl.StringOrNil(p.Name),
		DisplayName:                          dcl.StringOrNil(p.DisplayName),
		Zone:                                 dcl.StringOrNil(p.Zone),
		AlternativeZone:                      dcl.StringOrNil(p.AlternativeZone),
		Sku:                                  ProtoToTier2AlphaInstanceSku(p.GetSku()),
		AuthorizedNetworkId:                  dcl.StringOrNil(p.AuthorizedNetworkId),
		ReservedIPRange:                      dcl.StringOrNil(p.ReservedIpRange),
		Host:                                 dcl.StringOrNil(p.Host),
		Port:                                 dcl.Int64OrNil(p.Port),
		CurrentZone:                          dcl.StringOrNil(p.CurrentZone),
		CreateTime:                           dcl.StringOrNil(p.GetCreateTime()),
		State:                                ProtoToTier2AlphaInstanceStateEnum(p.GetState()),
		StatusMessage:                        dcl.StringOrNil(p.StatusMessage),
		UpdateTime:                           dcl.StringOrNil(p.GetUpdateTime()),
		MutateUserId:                         dcl.Int64OrNil(p.MutateUserId),
		ReadUserId:                           dcl.Int64OrNil(p.ReadUserId),
		PreprocessCreateRecipe:               ProtoToTier2AlphaInstancePreprocessCreateRecipe(p.GetPreprocessCreateRecipe()),
		CreateRecipe:                         ProtoToTier2AlphaInstanceCreateRecipe(p.GetCreateRecipe()),
		DeleteRecipe:                         ProtoToTier2AlphaInstanceDeleteRecipe(p.GetDeleteRecipe()),
		UpdateRecipe:                         ProtoToTier2AlphaInstanceUpdateRecipe(p.GetUpdateRecipe()),
		PreprocessResetRecipe:                ProtoToTier2AlphaInstancePreprocessResetRecipe(p.GetPreprocessResetRecipe()),
		ResetRecipe:                          ProtoToTier2AlphaInstanceResetRecipe(p.GetResetRecipe()),
		PreprocessRepairRecipe:               ProtoToTier2AlphaInstancePreprocessRepairRecipe(p.GetPreprocessRepairRecipe()),
		RepairRecipe:                         ProtoToTier2AlphaInstanceRepairRecipe(p.GetRepairRecipe()),
		PreprocessDeleteRecipe:               ProtoToTier2AlphaInstancePreprocessDeleteRecipe(p.GetPreprocessDeleteRecipe()),
		PreprocessUpdateRecipe:               ProtoToTier2AlphaInstancePreprocessUpdateRecipe(p.GetPreprocessUpdateRecipe()),
		PreprocessFreezeRecipe:               ProtoToTier2AlphaInstancePreprocessFreezeRecipe(p.GetPreprocessFreezeRecipe()),
		FreezeRecipe:                         ProtoToTier2AlphaInstanceFreezeRecipe(p.GetFreezeRecipe()),
		PreprocessUnfreezeRecipe:             ProtoToTier2AlphaInstancePreprocessUnfreezeRecipe(p.GetPreprocessUnfreezeRecipe()),
		UnfreezeRecipe:                       ProtoToTier2AlphaInstanceUnfreezeRecipe(p.GetUnfreezeRecipe()),
		PreprocessReportInstanceHealthRecipe: ProtoToTier2AlphaInstancePreprocessReportInstanceHealthRecipe(p.GetPreprocessReportInstanceHealthRecipe()),
		ReportInstanceHealthRecipe:           ProtoToTier2AlphaInstanceReportInstanceHealthRecipe(p.GetReportInstanceHealthRecipe()),
		PreprocessGetRecipe:                  ProtoToTier2AlphaInstancePreprocessGetRecipe(p.GetPreprocessGetRecipe()),
		NotifyKeyAvailableRecipe:             ProtoToTier2AlphaInstanceNotifyKeyAvailableRecipe(p.GetNotifyKeyAvailableRecipe()),
		NotifyKeyUnavailableRecipe:           ProtoToTier2AlphaInstanceNotifyKeyUnavailableRecipe(p.GetNotifyKeyUnavailableRecipe()),
		ReadonlyRecipe:                       ProtoToTier2AlphaInstanceReadonlyRecipe(p.GetReadonlyRecipe()),
		ReconcileRecipe:                      ProtoToTier2AlphaInstanceReconcileRecipe(p.GetReconcileRecipe()),
		PreprocessPassthroughRecipe:          ProtoToTier2AlphaInstancePreprocessPassthroughRecipe(p.GetPreprocessPassthroughRecipe()),
		PreprocessReconcileRecipe:            ProtoToTier2AlphaInstancePreprocessReconcileRecipe(p.GetPreprocessReconcileRecipe()),
		EnableCallHistory:                    dcl.Bool(p.EnableCallHistory),
		PublicResourceViewOverride:           dcl.StringOrNil(p.PublicResourceViewOverride),
		ExtraInfo:                            dcl.StringOrNil(p.ExtraInfo),
		Uid:                                  dcl.StringOrNil(p.Uid),
		Etag:                                 dcl.StringOrNil(p.Etag),
		Project:                              dcl.StringOrNil(p.Project),
		Location:                             dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetReferences() {
		obj.References = append(obj.References, *ProtoToTier2AlphaInstanceReferences(r))
	}
	for _, r := range p.GetEncryptionKeys() {
		obj.EncryptionKeys = append(obj.EncryptionKeys, *ProtoToTier2AlphaInstanceEncryptionKeys(r))
	}
	for _, r := range p.GetHistory() {
		obj.History = append(obj.History, *ProtoToTier2AlphaInstanceHistory(r))
	}
	return obj
}

// InstanceSkuTierEnumToProto converts a InstanceSkuTierEnum enum to its proto representation.
func Tier2AlphaInstanceSkuTierEnumToProto(e *alpha.InstanceSkuTierEnum) alphapb.Tier2AlphaInstanceSkuTierEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceSkuTierEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceSkuTierEnum_value["InstanceSkuTierEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceSkuTierEnum(v)
	}
	return alphapb.Tier2AlphaInstanceSkuTierEnum(0)
}

// InstanceSkuSizeEnumToProto converts a InstanceSkuSizeEnum enum to its proto representation.
func Tier2AlphaInstanceSkuSizeEnumToProto(e *alpha.InstanceSkuSizeEnum) alphapb.Tier2AlphaInstanceSkuSizeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceSkuSizeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceSkuSizeEnum_value["InstanceSkuSizeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceSkuSizeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceSkuSizeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func Tier2AlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.Tier2AlphaInstanceStateEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceStateEnum(v)
	}
	return alphapb.Tier2AlphaInstanceStateEnum(0)
}

// InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnumToProto converts a InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum enum to its proto representation.
func Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnumToProto(e *alpha.InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum) alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum_value["InstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(v)
	}
	return alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnum(0)
}

// InstancePreprocessCreateRecipeStepsActionEnumToProto converts a InstancePreprocessCreateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessCreateRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum_value["InstancePreprocessCreateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(0)
}

// InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceCreateRecipeStepsActionEnumToProto converts a InstanceCreateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsActionEnumToProto(e *alpha.InstanceCreateRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum_value["InstanceCreateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum(0)
}

// InstanceCreateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceDeleteRecipeStepsActionEnumToProto converts a InstanceDeleteRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsActionEnumToProto(e *alpha.InstanceDeleteRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum_value["InstanceDeleteRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum(0)
}

// InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceUpdateRecipeStepsActionEnumToProto converts a InstanceUpdateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsActionEnumToProto(e *alpha.InstanceUpdateRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum_value["InstanceUpdateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum(0)
}

// InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessResetRecipeStepsActionEnumToProto converts a InstancePreprocessResetRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessResetRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum_value["InstancePreprocessResetRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum(0)
}

// InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceResetRecipeStepsActionEnumToProto converts a InstanceResetRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsActionEnumToProto(e *alpha.InstanceResetRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum_value["InstanceResetRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum(0)
}

// InstanceResetRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceResetRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceResetRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceResetRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessRepairRecipeStepsActionEnumToProto converts a InstancePreprocessRepairRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessRepairRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum_value["InstancePreprocessRepairRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(0)
}

// InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceRepairRecipeStepsActionEnumToProto converts a InstanceRepairRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsActionEnumToProto(e *alpha.InstanceRepairRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum_value["InstanceRepairRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum(0)
}

// InstanceRepairRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessDeleteRecipeStepsActionEnumToProto converts a InstancePreprocessDeleteRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessDeleteRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum_value["InstancePreprocessDeleteRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(0)
}

// InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessUpdateRecipeStepsActionEnumToProto converts a InstancePreprocessUpdateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessUpdateRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum_value["InstancePreprocessUpdateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(0)
}

// InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessFreezeRecipeStepsActionEnumToProto converts a InstancePreprocessFreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessFreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum_value["InstancePreprocessFreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(0)
}

// InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceFreezeRecipeStepsActionEnumToProto converts a InstanceFreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsActionEnumToProto(e *alpha.InstanceFreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum_value["InstanceFreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum(0)
}

// InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessUnfreezeRecipeStepsActionEnumToProto converts a InstancePreprocessUnfreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessUnfreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum_value["InstancePreprocessUnfreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(0)
}

// InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceUnfreezeRecipeStepsActionEnumToProto converts a InstanceUnfreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsActionEnumToProto(e *alpha.InstanceUnfreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum_value["InstanceUnfreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum(0)
}

// InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessReportInstanceHealthRecipeStepsActionEnumToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessReportInstanceHealthRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum_value["InstancePreprocessReportInstanceHealthRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnum(0)
}

// InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceReportInstanceHealthRecipeStepsActionEnumToProto converts a InstanceReportInstanceHealthRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnumToProto(e *alpha.InstanceReportInstanceHealthRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum_value["InstanceReportInstanceHealthRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnum(0)
}

// InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessGetRecipeStepsActionEnumToProto converts a InstancePreprocessGetRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessGetRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum_value["InstancePreprocessGetRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsActionEnum(0)
}

// InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceNotifyKeyAvailableRecipeStepsActionEnumToProto converts a InstanceNotifyKeyAvailableRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnumToProto(e *alpha.InstanceNotifyKeyAvailableRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum_value["InstanceNotifyKeyAvailableRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnum(0)
}

// InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceNotifyKeyUnavailableRecipeStepsActionEnumToProto converts a InstanceNotifyKeyUnavailableRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnumToProto(e *alpha.InstanceNotifyKeyUnavailableRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum_value["InstanceNotifyKeyUnavailableRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnum(0)
}

// InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceReadonlyRecipeStepsActionEnumToProto converts a InstanceReadonlyRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsActionEnumToProto(e *alpha.InstanceReadonlyRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum_value["InstanceReadonlyRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum(0)
}

// InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceReconcileRecipeStepsActionEnumToProto converts a InstanceReconcileRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsActionEnumToProto(e *alpha.InstanceReconcileRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum_value["InstanceReconcileRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReconcileRecipeStepsActionEnum(0)
}

// InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessPassthroughRecipeStepsActionEnumToProto converts a InstancePreprocessPassthroughRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessPassthroughRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum_value["InstancePreprocessPassthroughRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnum(0)
}

// InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstancePreprocessReconcileRecipeStepsActionEnumToProto converts a InstancePreprocessReconcileRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessReconcileRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum_value["InstancePreprocessReconcileRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnum(0)
}

// InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnumToProto converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(e *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum) alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum_value["InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnum(0)
}

// InstanceSkuToProto converts a InstanceSku resource to its proto representation.
func Tier2AlphaInstanceSkuToProto(o *alpha.InstanceSku) *alphapb.Tier2AlphaInstanceSku {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceSku{
		Tier: Tier2AlphaInstanceSkuTierEnumToProto(o.Tier),
		Size: Tier2AlphaInstanceSkuSizeEnumToProto(o.Size),
	}
	return p
}

// InstanceReferencesToProto converts a InstanceReferences resource to its proto representation.
func Tier2AlphaInstanceReferencesToProto(o *alpha.InstanceReferences) *alphapb.Tier2AlphaInstanceReferences {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReferences{
		Name:           dcl.ValueOrEmptyString(o.Name),
		Type:           dcl.ValueOrEmptyString(o.Type),
		SourceResource: dcl.ValueOrEmptyString(o.SourceResource),
		CreateTime:     dcl.ValueOrEmptyString(o.CreateTime),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceReferencesDetailsToProto(&r))
	}
	return p
}

// InstanceReferencesDetailsToProto converts a InstanceReferencesDetails resource to its proto representation.
func Tier2AlphaInstanceReferencesDetailsToProto(o *alpha.InstanceReferencesDetails) *alphapb.Tier2AlphaInstanceReferencesDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReferencesDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceEncryptionKeysToProto converts a InstanceEncryptionKeys resource to its proto representation.
func Tier2AlphaInstanceEncryptionKeysToProto(o *alpha.InstanceEncryptionKeys) *alphapb.Tier2AlphaInstanceEncryptionKeys {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceEncryptionKeys{
		KeyOrVersion: dcl.ValueOrEmptyString(o.KeyOrVersion),
		Grant:        dcl.ValueOrEmptyString(o.Grant),
		Delegate:     dcl.ValueOrEmptyString(o.Delegate),
		KeyState:     Tier2AlphaInstanceEncryptionKeysKeyStateToProto(o.KeyState),
	}
	return p
}

// InstanceEncryptionKeysKeyStateToProto converts a InstanceEncryptionKeysKeyState resource to its proto representation.
func Tier2AlphaInstanceEncryptionKeysKeyStateToProto(o *alpha.InstanceEncryptionKeysKeyState) *alphapb.Tier2AlphaInstanceEncryptionKeysKeyState {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceEncryptionKeysKeyState{
		KeyStateVersion: dcl.ValueOrEmptyInt64(o.KeyStateVersion),
		Availability:    Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityToProto(o.Availability),
	}
	return p
}

// InstanceEncryptionKeysKeyStateAvailabilityToProto converts a InstanceEncryptionKeysKeyStateAvailability resource to its proto representation.
func Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityToProto(o *alpha.InstanceEncryptionKeysKeyStateAvailability) *alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailability {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceEncryptionKeysKeyStateAvailability{
		PermissionDenied: dcl.ValueOrEmptyBool(o.PermissionDenied),
		UnknownFailure:   dcl.ValueOrEmptyBool(o.UnknownFailure),
		KeyVersionState:  Tier2AlphaInstanceEncryptionKeysKeyStateAvailabilityKeyVersionStateEnumToProto(o.KeyVersionState),
	}
	return p
}

// InstancePreprocessCreateRecipeToProto converts a InstancePreprocessCreateRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeToProto(o *alpha.InstancePreprocessCreateRecipe) *alphapb.Tier2AlphaInstancePreprocessCreateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessCreateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessCreateRecipeStepsToProto converts a InstancePreprocessCreateRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsToProto(o *alpha.InstancePreprocessCreateRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessCreateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsStatusToProto converts a InstancePreprocessCreateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsStatusToProto(o *alpha.InstancePreprocessCreateRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsStatusDetailsToProto converts a InstancePreprocessCreateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessCreateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessCreateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessCreateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessCreateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessCreateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessCreateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceGoogleprotobufstructToProto converts a InstanceGoogleprotobufstruct resource to its proto representation.
func Tier2AlphaInstanceGoogleprotobufstructToProto(o *alpha.InstanceGoogleprotobufstruct) *alphapb.Tier2AlphaInstanceGoogleprotobufstruct {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceGoogleprotobufstruct{}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceCreateRecipeToProto converts a InstanceCreateRecipe resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeToProto(o *alpha.InstanceCreateRecipe) *alphapb.Tier2AlphaInstanceCreateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceCreateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceCreateRecipeStepsToProto converts a InstanceCreateRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsToProto(o *alpha.InstanceCreateRecipeSteps) *alphapb.Tier2AlphaInstanceCreateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceCreateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceCreateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsStatusToProto converts a InstanceCreateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsStatusToProto(o *alpha.InstanceCreateRecipeStepsStatus) *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceCreateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsStatusDetailsToProto converts a InstanceCreateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsStatusDetailsToProto(o *alpha.InstanceCreateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceCreateRecipeStepsQuotaRequestDeltasToProto converts a InstanceCreateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceCreateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceCreateRecipeStepsPreprocessUpdateToProto converts a InstanceCreateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceCreateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceCreateRecipeStepsRequestedTenantProjectToProto converts a InstanceCreateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceCreateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoToProto converts a InstanceCreateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceCreateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceCreateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoResourceToProto converts a InstanceCreateRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceDeleteRecipeToProto converts a InstanceDeleteRecipe resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeToProto(o *alpha.InstanceDeleteRecipe) *alphapb.Tier2AlphaInstanceDeleteRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceDeleteRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceDeleteRecipeStepsToProto converts a InstanceDeleteRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsToProto(o *alpha.InstanceDeleteRecipeSteps) *alphapb.Tier2AlphaInstanceDeleteRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceDeleteRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceDeleteRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsStatusToProto converts a InstanceDeleteRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsStatusToProto(o *alpha.InstanceDeleteRecipeStepsStatus) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceDeleteRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsStatusDetailsToProto converts a InstanceDeleteRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsStatusDetailsToProto(o *alpha.InstanceDeleteRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceDeleteRecipeStepsQuotaRequestDeltasToProto converts a InstanceDeleteRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceDeleteRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceDeleteRecipeStepsPreprocessUpdateToProto converts a InstanceDeleteRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceDeleteRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceDeleteRecipeStepsRequestedTenantProjectToProto converts a InstanceDeleteRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceDeleteRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoToProto converts a InstanceDeleteRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceDeleteRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceDeleteRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoResourceToProto converts a InstanceDeleteRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceUpdateRecipeToProto converts a InstanceUpdateRecipe resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeToProto(o *alpha.InstanceUpdateRecipe) *alphapb.Tier2AlphaInstanceUpdateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceUpdateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceUpdateRecipeStepsToProto converts a InstanceUpdateRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsToProto(o *alpha.InstanceUpdateRecipeSteps) *alphapb.Tier2AlphaInstanceUpdateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceUpdateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceUpdateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsStatusToProto converts a InstanceUpdateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsStatusToProto(o *alpha.InstanceUpdateRecipeStepsStatus) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceUpdateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsStatusDetailsToProto converts a InstanceUpdateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsStatusDetailsToProto(o *alpha.InstanceUpdateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceUpdateRecipeStepsQuotaRequestDeltasToProto converts a InstanceUpdateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceUpdateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceUpdateRecipeStepsPreprocessUpdateToProto converts a InstanceUpdateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceUpdateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceUpdateRecipeStepsRequestedTenantProjectToProto converts a InstanceUpdateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceUpdateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoToProto converts a InstanceUpdateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceUpdateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceUpdateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoResourceToProto converts a InstanceUpdateRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessResetRecipeToProto converts a InstancePreprocessResetRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeToProto(o *alpha.InstancePreprocessResetRecipe) *alphapb.Tier2AlphaInstancePreprocessResetRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessResetRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessResetRecipeStepsToProto converts a InstancePreprocessResetRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsToProto(o *alpha.InstancePreprocessResetRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessResetRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessResetRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessResetRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsStatusToProto converts a InstancePreprocessResetRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsStatusToProto(o *alpha.InstancePreprocessResetRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsStatusDetailsToProto converts a InstancePreprocessResetRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessResetRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessResetRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessResetRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessResetRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessResetRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessResetRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessResetRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessResetRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceResetRecipeToProto converts a InstanceResetRecipe resource to its proto representation.
func Tier2AlphaInstanceResetRecipeToProto(o *alpha.InstanceResetRecipe) *alphapb.Tier2AlphaInstanceResetRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceResetRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceResetRecipeStepsToProto converts a InstanceResetRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsToProto(o *alpha.InstanceResetRecipeSteps) *alphapb.Tier2AlphaInstanceResetRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceResetRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceResetRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceResetRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceResetRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsStatusToProto converts a InstanceResetRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsStatusToProto(o *alpha.InstanceResetRecipeStepsStatus) *alphapb.Tier2AlphaInstanceResetRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceResetRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsStatusDetailsToProto converts a InstanceResetRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsStatusDetailsToProto(o *alpha.InstanceResetRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceResetRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceResetRecipeStepsQuotaRequestDeltasToProto converts a InstanceResetRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceResetRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceResetRecipeStepsPreprocessUpdateToProto converts a InstanceResetRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceResetRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceResetRecipeStepsRequestedTenantProjectToProto converts a InstanceResetRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceResetRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoToProto converts a InstanceResetRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceResetRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceResetRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceResetRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoResourceToProto converts a InstanceResetRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessRepairRecipeToProto converts a InstancePreprocessRepairRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeToProto(o *alpha.InstancePreprocessRepairRecipe) *alphapb.Tier2AlphaInstancePreprocessRepairRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessRepairRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessRepairRecipeStepsToProto converts a InstancePreprocessRepairRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsToProto(o *alpha.InstancePreprocessRepairRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessRepairRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsStatusToProto converts a InstancePreprocessRepairRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsStatusToProto(o *alpha.InstancePreprocessRepairRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsStatusDetailsToProto converts a InstancePreprocessRepairRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessRepairRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessRepairRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessRepairRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessRepairRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessRepairRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessRepairRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceRepairRecipeToProto converts a InstanceRepairRecipe resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeToProto(o *alpha.InstanceRepairRecipe) *alphapb.Tier2AlphaInstanceRepairRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceRepairRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceRepairRecipeStepsToProto converts a InstanceRepairRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsToProto(o *alpha.InstanceRepairRecipeSteps) *alphapb.Tier2AlphaInstanceRepairRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceRepairRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceRepairRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsStatusToProto converts a InstanceRepairRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsStatusToProto(o *alpha.InstanceRepairRecipeStepsStatus) *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceRepairRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsStatusDetailsToProto converts a InstanceRepairRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsStatusDetailsToProto(o *alpha.InstanceRepairRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceRepairRecipeStepsQuotaRequestDeltasToProto converts a InstanceRepairRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceRepairRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceRepairRecipeStepsPreprocessUpdateToProto converts a InstanceRepairRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceRepairRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceRepairRecipeStepsRequestedTenantProjectToProto converts a InstanceRepairRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceRepairRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoToProto converts a InstanceRepairRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceRepairRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceRepairRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoResourceToProto converts a InstanceRepairRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessDeleteRecipeToProto converts a InstancePreprocessDeleteRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeToProto(o *alpha.InstancePreprocessDeleteRecipe) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessDeleteRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsToProto converts a InstancePreprocessDeleteRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsToProto(o *alpha.InstancePreprocessDeleteRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsStatusToProto converts a InstancePreprocessDeleteRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusToProto(o *alpha.InstancePreprocessDeleteRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsStatusDetailsToProto converts a InstancePreprocessDeleteRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessDeleteRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessUpdateRecipeToProto converts a InstancePreprocessUpdateRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeToProto(o *alpha.InstancePreprocessUpdateRecipe) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessUpdateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsToProto converts a InstancePreprocessUpdateRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsToProto(o *alpha.InstancePreprocessUpdateRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsStatusToProto converts a InstancePreprocessUpdateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusToProto(o *alpha.InstancePreprocessUpdateRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsStatusDetailsToProto converts a InstancePreprocessUpdateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessUpdateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessFreezeRecipeToProto converts a InstancePreprocessFreezeRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeToProto(o *alpha.InstancePreprocessFreezeRecipe) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessFreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsToProto converts a InstancePreprocessFreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsToProto(o *alpha.InstancePreprocessFreezeRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsStatusToProto converts a InstancePreprocessFreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusToProto(o *alpha.InstancePreprocessFreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsStatusDetailsToProto converts a InstancePreprocessFreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessFreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceFreezeRecipeToProto converts a InstanceFreezeRecipe resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeToProto(o *alpha.InstanceFreezeRecipe) *alphapb.Tier2AlphaInstanceFreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceFreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceFreezeRecipeStepsToProto converts a InstanceFreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsToProto(o *alpha.InstanceFreezeRecipeSteps) *alphapb.Tier2AlphaInstanceFreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceFreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceFreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsStatusToProto converts a InstanceFreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsStatusToProto(o *alpha.InstanceFreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceFreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsStatusDetailsToProto converts a InstanceFreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsStatusDetailsToProto(o *alpha.InstanceFreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceFreezeRecipeStepsQuotaRequestDeltasToProto converts a InstanceFreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceFreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceFreezeRecipeStepsPreprocessUpdateToProto converts a InstanceFreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceFreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceFreezeRecipeStepsRequestedTenantProjectToProto converts a InstanceFreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceFreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoToProto converts a InstanceFreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceFreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceFreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoResourceToProto converts a InstanceFreezeRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeToProto converts a InstancePreprocessUnfreezeRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeToProto(o *alpha.InstancePreprocessUnfreezeRecipe) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsToProto converts a InstancePreprocessUnfreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsToProto(o *alpha.InstancePreprocessUnfreezeRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsStatusToProto converts a InstancePreprocessUnfreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsStatusDetailsToProto converts a InstancePreprocessUnfreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceUnfreezeRecipeToProto converts a InstanceUnfreezeRecipe resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeToProto(o *alpha.InstanceUnfreezeRecipe) *alphapb.Tier2AlphaInstanceUnfreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceUnfreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceUnfreezeRecipeStepsToProto converts a InstanceUnfreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsToProto(o *alpha.InstanceUnfreezeRecipeSteps) *alphapb.Tier2AlphaInstanceUnfreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceUnfreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceUnfreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsStatusToProto converts a InstanceUnfreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsStatusToProto(o *alpha.InstanceUnfreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsStatusDetailsToProto converts a InstanceUnfreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetailsToProto(o *alpha.InstanceUnfreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceUnfreezeRecipeStepsQuotaRequestDeltasToProto converts a InstanceUnfreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceUnfreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPreprocessUpdateToProto converts a InstanceUnfreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceUnfreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceUnfreezeRecipeStepsRequestedTenantProjectToProto converts a InstanceUnfreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoResourceToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeToProto converts a InstancePreprocessReportInstanceHealthRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipe) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsToProto converts a InstancePreprocessReportInstanceHealthRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsStatusToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceReportInstanceHealthRecipeToProto converts a InstanceReportInstanceHealthRecipe resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeToProto(o *alpha.InstanceReportInstanceHealthRecipe) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceReportInstanceHealthRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsToProto converts a InstanceReportInstanceHealthRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsToProto(o *alpha.InstanceReportInstanceHealthRecipeSteps) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceReportInstanceHealthRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsStatusToProto converts a InstanceReportInstanceHealthRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsStatus) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsStatusDetailsToProto converts a InstanceReportInstanceHealthRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetailsToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasToProto converts a InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsPreprocessUpdateToProto converts a InstanceReportInstanceHealthRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectToProto converts a InstanceReportInstanceHealthRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsPermissionsInfoToProto converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsPermissionsInfoResourceToProto converts a InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateToProto converts a InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessGetRecipeToProto converts a InstancePreprocessGetRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeToProto(o *alpha.InstancePreprocessGetRecipe) *alphapb.Tier2AlphaInstancePreprocessGetRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessGetRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessGetRecipeStepsToProto converts a InstancePreprocessGetRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsToProto(o *alpha.InstancePreprocessGetRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessGetRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessGetRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessGetRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessGetRecipeStepsStatusToProto converts a InstancePreprocessGetRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsStatusToProto(o *alpha.InstancePreprocessGetRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessGetRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessGetRecipeStepsStatusDetailsToProto converts a InstancePreprocessGetRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessGetRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessGetRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessGetRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessGetRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessGetRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessGetRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessGetRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessGetRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessGetRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessGetRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessGetRecipeStepsPermissionsInfoToProto converts a InstancePreprocessGetRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessGetRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessGetRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessGetRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessGetRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessGetRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessGetRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeToProto converts a InstanceNotifyKeyAvailableRecipe resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeToProto(o *alpha.InstanceNotifyKeyAvailableRecipe) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsToProto converts a InstanceNotifyKeyAvailableRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsToProto(o *alpha.InstanceNotifyKeyAvailableRecipeSteps) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsStatusToProto converts a InstanceNotifyKeyAvailableRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsStatus) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsStatusDetailsToProto converts a InstanceNotifyKeyAvailableRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetailsToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasToProto converts a InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdateToProto converts a InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectToProto converts a InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoToProto converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResourceToProto converts a InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateToProto converts a InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeToProto converts a InstanceNotifyKeyUnavailableRecipe resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeToProto(o *alpha.InstanceNotifyKeyUnavailableRecipe) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsToProto converts a InstanceNotifyKeyUnavailableRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeSteps) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsStatusToProto converts a InstanceNotifyKeyUnavailableRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsStatus) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsStatusDetailsToProto converts a InstanceNotifyKeyUnavailableRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetailsToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasToProto converts a InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdateToProto converts a InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectToProto converts a InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoToProto converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResourceToProto converts a InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateToProto converts a InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceReadonlyRecipeToProto converts a InstanceReadonlyRecipe resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeToProto(o *alpha.InstanceReadonlyRecipe) *alphapb.Tier2AlphaInstanceReadonlyRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceReadonlyRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceReadonlyRecipeStepsToProto converts a InstanceReadonlyRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsToProto(o *alpha.InstanceReadonlyRecipeSteps) *alphapb.Tier2AlphaInstanceReadonlyRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceReadonlyRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceReadonlyRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsStatusToProto converts a InstanceReadonlyRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsStatusToProto(o *alpha.InstanceReadonlyRecipeStepsStatus) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceReadonlyRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsStatusDetailsToProto converts a InstanceReadonlyRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsStatusDetailsToProto(o *alpha.InstanceReadonlyRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceReadonlyRecipeStepsQuotaRequestDeltasToProto converts a InstanceReadonlyRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceReadonlyRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceReadonlyRecipeStepsPreprocessUpdateToProto converts a InstanceReadonlyRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceReadonlyRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceReadonlyRecipeStepsRequestedTenantProjectToProto converts a InstanceReadonlyRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceReadonlyRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoToProto converts a InstanceReadonlyRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoResourceToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceReconcileRecipeToProto converts a InstanceReconcileRecipe resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeToProto(o *alpha.InstanceReconcileRecipe) *alphapb.Tier2AlphaInstanceReconcileRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceReconcileRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceReconcileRecipeStepsToProto converts a InstanceReconcileRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsToProto(o *alpha.InstanceReconcileRecipeSteps) *alphapb.Tier2AlphaInstanceReconcileRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceReconcileRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceReconcileRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceReconcileRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceReconcileRecipeStepsStatusToProto converts a InstanceReconcileRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsStatusToProto(o *alpha.InstanceReconcileRecipeStepsStatus) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceReconcileRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceReconcileRecipeStepsStatusDetailsToProto converts a InstanceReconcileRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsStatusDetailsToProto(o *alpha.InstanceReconcileRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceReconcileRecipeStepsQuotaRequestDeltasToProto converts a InstanceReconcileRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceReconcileRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstanceReconcileRecipeStepsPreprocessUpdateToProto converts a InstanceReconcileRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceReconcileRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceReconcileRecipeStepsRequestedTenantProjectToProto converts a InstanceReconcileRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceReconcileRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceReconcileRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceReconcileRecipeStepsPermissionsInfoToProto converts a InstanceReconcileRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoToProto(o *alpha.InstanceReconcileRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceReconcileRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceReconcileRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceReconcileRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceReconcileRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceReconcileRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceReconcileRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceReconcileRecipeStepsPermissionsInfoResourceToProto converts a InstanceReconcileRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstanceReconcileRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstanceReconcileRecipeStepsKeyNotificationsUpdateToProto converts a InstanceReconcileRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessPassthroughRecipeToProto converts a InstancePreprocessPassthroughRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeToProto(o *alpha.InstancePreprocessPassthroughRecipe) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessPassthroughRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsToProto converts a InstancePreprocessPassthroughRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsToProto(o *alpha.InstancePreprocessPassthroughRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessPassthroughRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsStatusToProto converts a InstancePreprocessPassthroughRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsStatusDetailsToProto converts a InstancePreprocessPassthroughRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessPassthroughRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessPassthroughRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsPermissionsInfoToProto converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstancePreprocessReconcileRecipeToProto converts a InstancePreprocessReconcileRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeToProto(o *alpha.InstancePreprocessReconcileRecipe) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessReconcileRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsToProto converts a InstancePreprocessReconcileRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsToProto(o *alpha.InstancePreprocessReconcileRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessReconcileRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
		PublicErrorMessage:             dcl.ValueOrEmptyString(o.PublicErrorMessage),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsStatusToProto converts a InstancePreprocessReconcileRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusToProto(o *alpha.InstancePreprocessReconcileRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsStatusDetailsToProto converts a InstancePreprocessReconcileRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessReconcileRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas{
		MetricName:        dcl.ValueOrEmptyString(o.MetricName),
		Amount:            dcl.ValueOrEmptyInt64(o.Amount),
		QuotaLocationName: dcl.ValueOrEmptyString(o.QuotaLocationName),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessReconcileRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessReconcileRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessReconcileRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessReconcileRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsPermissionsInfoToProto converts a InstancePreprocessReconcileRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfo{
		PolicyName:     Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath:   dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:       Tier2AlphaInstanceGoogleprotobufstructToProto(o.ApiAttrs),
		PolicyNameMode: Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameModeEnumToProto(o.PolicyNameMode),
		Resource:       Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResourceToProto(o.Resource),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsPermissionsInfoResourceToProto converts a InstancePreprocessReconcileRecipeStepsPermissionsInfoResource resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResourceToProto(o *alpha.InstancePreprocessReconcileRecipeStepsPermissionsInfoResource) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResource {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsPermissionsInfoResource{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Type:    dcl.ValueOrEmptyString(o.Type),
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	for _, r := range o.KeyNotificationConfigs {
		p.KeyNotificationConfigs = append(p.KeyNotificationConfigs, Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(&r))
	}
	return p
}

// InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto converts a InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigsToProto(o *alpha.InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) *alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
		Grant:            dcl.ValueOrEmptyString(o.Grant),
		DelegatorGaiaId:  dcl.ValueOrEmptyInt64(o.DelegatorGaiaId),
	}
	return p
}

// InstanceHistoryToProto converts a InstanceHistory resource to its proto representation.
func Tier2AlphaInstanceHistoryToProto(o *alpha.InstanceHistory) *alphapb.Tier2AlphaInstanceHistory {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceHistory{
		Timestamp:           dcl.ValueOrEmptyString(o.Timestamp),
		OperationHandle:     dcl.ValueOrEmptyString(o.OperationHandle),
		Description:         dcl.ValueOrEmptyString(o.Description),
		StepIndex:           dcl.ValueOrEmptyInt64(o.StepIndex),
		TenantProjectNumber: dcl.ValueOrEmptyInt64(o.TenantProjectNumber),
		TenantProjectId:     dcl.ValueOrEmptyString(o.TenantProjectId),
		P4ServiceAccount:    dcl.ValueOrEmptyString(o.P4ServiceAccount),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.Tier2AlphaInstance {
	p := &alphapb.Tier2AlphaInstance{
		Name:                                 dcl.ValueOrEmptyString(resource.Name),
		DisplayName:                          dcl.ValueOrEmptyString(resource.DisplayName),
		Zone:                                 dcl.ValueOrEmptyString(resource.Zone),
		AlternativeZone:                      dcl.ValueOrEmptyString(resource.AlternativeZone),
		Sku:                                  Tier2AlphaInstanceSkuToProto(resource.Sku),
		AuthorizedNetworkId:                  dcl.ValueOrEmptyString(resource.AuthorizedNetworkId),
		ReservedIpRange:                      dcl.ValueOrEmptyString(resource.ReservedIPRange),
		Host:                                 dcl.ValueOrEmptyString(resource.Host),
		Port:                                 dcl.ValueOrEmptyInt64(resource.Port),
		CurrentZone:                          dcl.ValueOrEmptyString(resource.CurrentZone),
		CreateTime:                           dcl.ValueOrEmptyString(resource.CreateTime),
		State:                                Tier2AlphaInstanceStateEnumToProto(resource.State),
		StatusMessage:                        dcl.ValueOrEmptyString(resource.StatusMessage),
		UpdateTime:                           dcl.ValueOrEmptyString(resource.UpdateTime),
		MutateUserId:                         dcl.ValueOrEmptyInt64(resource.MutateUserId),
		ReadUserId:                           dcl.ValueOrEmptyInt64(resource.ReadUserId),
		PreprocessCreateRecipe:               Tier2AlphaInstancePreprocessCreateRecipeToProto(resource.PreprocessCreateRecipe),
		CreateRecipe:                         Tier2AlphaInstanceCreateRecipeToProto(resource.CreateRecipe),
		DeleteRecipe:                         Tier2AlphaInstanceDeleteRecipeToProto(resource.DeleteRecipe),
		UpdateRecipe:                         Tier2AlphaInstanceUpdateRecipeToProto(resource.UpdateRecipe),
		PreprocessResetRecipe:                Tier2AlphaInstancePreprocessResetRecipeToProto(resource.PreprocessResetRecipe),
		ResetRecipe:                          Tier2AlphaInstanceResetRecipeToProto(resource.ResetRecipe),
		PreprocessRepairRecipe:               Tier2AlphaInstancePreprocessRepairRecipeToProto(resource.PreprocessRepairRecipe),
		RepairRecipe:                         Tier2AlphaInstanceRepairRecipeToProto(resource.RepairRecipe),
		PreprocessDeleteRecipe:               Tier2AlphaInstancePreprocessDeleteRecipeToProto(resource.PreprocessDeleteRecipe),
		PreprocessUpdateRecipe:               Tier2AlphaInstancePreprocessUpdateRecipeToProto(resource.PreprocessUpdateRecipe),
		PreprocessFreezeRecipe:               Tier2AlphaInstancePreprocessFreezeRecipeToProto(resource.PreprocessFreezeRecipe),
		FreezeRecipe:                         Tier2AlphaInstanceFreezeRecipeToProto(resource.FreezeRecipe),
		PreprocessUnfreezeRecipe:             Tier2AlphaInstancePreprocessUnfreezeRecipeToProto(resource.PreprocessUnfreezeRecipe),
		UnfreezeRecipe:                       Tier2AlphaInstanceUnfreezeRecipeToProto(resource.UnfreezeRecipe),
		PreprocessReportInstanceHealthRecipe: Tier2AlphaInstancePreprocessReportInstanceHealthRecipeToProto(resource.PreprocessReportInstanceHealthRecipe),
		ReportInstanceHealthRecipe:           Tier2AlphaInstanceReportInstanceHealthRecipeToProto(resource.ReportInstanceHealthRecipe),
		PreprocessGetRecipe:                  Tier2AlphaInstancePreprocessGetRecipeToProto(resource.PreprocessGetRecipe),
		NotifyKeyAvailableRecipe:             Tier2AlphaInstanceNotifyKeyAvailableRecipeToProto(resource.NotifyKeyAvailableRecipe),
		NotifyKeyUnavailableRecipe:           Tier2AlphaInstanceNotifyKeyUnavailableRecipeToProto(resource.NotifyKeyUnavailableRecipe),
		ReadonlyRecipe:                       Tier2AlphaInstanceReadonlyRecipeToProto(resource.ReadonlyRecipe),
		ReconcileRecipe:                      Tier2AlphaInstanceReconcileRecipeToProto(resource.ReconcileRecipe),
		PreprocessPassthroughRecipe:          Tier2AlphaInstancePreprocessPassthroughRecipeToProto(resource.PreprocessPassthroughRecipe),
		PreprocessReconcileRecipe:            Tier2AlphaInstancePreprocessReconcileRecipeToProto(resource.PreprocessReconcileRecipe),
		EnableCallHistory:                    dcl.ValueOrEmptyBool(resource.EnableCallHistory),
		PublicResourceViewOverride:           dcl.ValueOrEmptyString(resource.PublicResourceViewOverride),
		ExtraInfo:                            dcl.ValueOrEmptyString(resource.ExtraInfo),
		Uid:                                  dcl.ValueOrEmptyString(resource.Uid),
		Etag:                                 dcl.ValueOrEmptyString(resource.Etag),
		Project:                              dcl.ValueOrEmptyString(resource.Project),
		Location:                             dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.References {
		p.References = append(p.References, Tier2AlphaInstanceReferencesToProto(&r))
	}
	for _, r := range resource.EncryptionKeys {
		p.EncryptionKeys = append(p.EncryptionKeys, Tier2AlphaInstanceEncryptionKeysToProto(&r))
	}
	for _, r := range resource.History {
		p.History = append(p.History, Tier2AlphaInstanceHistoryToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyTier2AlphaInstanceRequest) (*alphapb.Tier2AlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyTier2AlphaInstance(ctx context.Context, request *alphapb.ApplyTier2AlphaInstanceRequest) (*alphapb.Tier2AlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteTier2AlphaInstance(ctx context.Context, request *alphapb.DeleteTier2AlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListTier2AlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListTier2AlphaInstance(ctx context.Context, request *alphapb.ListTier2AlphaInstanceRequest) (*alphapb.ListTier2AlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, ProtoToInstance(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.Tier2AlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListTier2AlphaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
