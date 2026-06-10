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

// +tool:fuzz-gen
// proto.message: google.cloud.kms.v1.KeyRing
// api.group: kms.cnrm.cloud.google.com

package kms

import (
	"strings"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(KMSKeyRingFuzzer())
}

func KMSKeyRingFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.KeyRing{},
		KMSKeyRingSpec_FromProto, KMSKeyRingSpec_ToProto,
		KMSKeyRingStatus_FromProto, KMSKeyRingStatus_ToProto,
	)

	f.SpecField(".location")
	f.SpecField(".resource_id")

	f.StatusField(".name")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".create_time")

	f.FilterSpec = func(in *pb.KeyRing) {
		if in.Name != "" {
			in.Name = "projects/p/locations/l/keyRings/" + sanitizeKeyRingID(in.Name)
		}
	}

	f.FilterStatus = func(in *pb.KeyRing) {
		if in.Name != "" {
			in.Name = "projects/p/locations/l/keyRings/" + sanitizeKeyRingID(in.Name)
		}
	}

	return f
}

func sanitizeKeyRingID(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			sb.WriteRune(r)
		}
	}
	if sb.Len() == 0 {
		return "keyring"
	}
	return sb.String()
}
