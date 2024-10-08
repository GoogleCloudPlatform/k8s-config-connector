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

package contexts

import "time"

func init() {
	resourceContextMap["firestoredatabase-minimal"] = ResourceContext{
		ResourceKind: "FirestoreDatabase",
		// Firestore Databases return success for create / update before the changes are actually visible
		// in GCP. It's probably a bug with the service, or some issue with eventual consistency.
		PostModifyDelay: 60 * time.Second,
		// After deleting a FirestoreDatabase, the database name is reserved for a while, so we cannot
		// immediately re-create the database with the same name.
		SkipDriftDetection: true,
	}

	resourceContextMap["firestoredatabase-full"] = ResourceContext{
		ResourceKind: "FirestoreDatabase",
		// Firestore Databases return success for create / update before the changes are actually visible
		// in GCP. It's probably a bug with the service, or some issue with eventual consistency.
		PostModifyDelay: 60 * time.Second,
		// After deleting a FirestoreDatabase, the database name is reserved for a while, so we cannot
		// immediately re-create the database with the same name.
		SkipDriftDetection: true,
	}

	resourceContextMap["firestoreindex"] = ResourceContext{
		ResourceKind: "FirestoreIndex",
		SkipUpdate:   true,
	}
}
