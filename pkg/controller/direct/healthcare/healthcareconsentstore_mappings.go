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

package healthcare

import (
	pb "google.golang.org/api/healthcare/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/healthcare/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func HealthcareConsentStoreSpec_FromAPI(mapCtx *direct.MapContext, in *pb.ConsentStore) *krm.HealthcareConsentStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.HealthcareConsentStoreSpec{}
	out.DefaultConsentTTL = direct.LazyPtr(in.DefaultConsentTtl)
	out.EnableConsentCreateOnUpdate = direct.LazyPtr(in.EnableConsentCreateOnUpdate)
	return out
}

func HealthcareConsentStoreSpec_ToAPI(mapCtx *direct.MapContext, in *krm.HealthcareConsentStoreSpec) *pb.ConsentStore {
	if in == nil {
		return nil
	}
	out := &pb.ConsentStore{}
	out.DefaultConsentTtl = direct.ValueOf(in.DefaultConsentTTL)
	out.EnableConsentCreateOnUpdate = direct.ValueOf(in.EnableConsentCreateOnUpdate)
	return out
}

func HealthcareConsentStoreStatus_FromAPI(mapCtx *direct.MapContext, in *pb.ConsentStore) *krm.HealthcareConsentStoreStatus {
	if in == nil {
		return nil
	}
	out := &krm.HealthcareConsentStoreStatus{}
	return out
}

func HealthcareConsentStoreStatus_ToAPI(mapCtx *direct.MapContext, in *krm.HealthcareConsentStoreStatus) *pb.ConsentStore {
	if in == nil {
		return nil
	}
	out := &pb.ConsentStore{}
	return out
}
