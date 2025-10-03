// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.firestore.admin.v1.Index
// api.group: firestore.cnrm.cloud.google.com

package firestore

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/firestore/admin/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(firestoreIndexFuzzer())
}

func firestoreIndexFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Index{},
		FirestoreIndexSpec_FromProto, FirestoreIndexSpec_ToProto,
		FirestoreIndexStatus_FromProto, FirestoreIndexStatus_ToProto,
	)

	f.SpecField(".query_scope")
	f.SpecField(".fields")

	f.StatusField(".name")

	f.IdentityField(".database")
	f.IdentityField(".collection")

	f.Unimplemented_NotYetTriaged(".api_scope")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".fields[].order")
	f.Unimplemented_NotYetTriaged(".fields[].vector_config")

	// We do map array_config, but we can't handle the unspecified value correctly
	f.Unimplemented_NotYetTriaged(".fields[].array_config")

	return f
}
