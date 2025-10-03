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

package firestore

import pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"

func ApplyFirestoreDatabaseDefaults(in *pb.Database) {
	// Set default values to make sure firestore database is "declarative-friendly".
	in.Type = pb.Database_FIRESTORE_NATIVE
	in.AppEngineIntegrationMode = pb.Database_DISABLED
	in.DeleteProtectionState = pb.Database_DELETE_PROTECTION_DISABLED
}
