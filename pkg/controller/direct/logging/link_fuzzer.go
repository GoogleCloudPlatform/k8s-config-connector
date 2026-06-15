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
// proto.message: google.logging.v2.Link
// api.group: logging.cnrm.cloud.google.com

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(loggingLinkFuzzer())
}

func loggingLinkFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Link{},
		LoggingLinkSpec_FromProto, LoggingLinkSpec_ToProto,
		LoggingLinkObservedState_FromProto, LoggingLinkObservedState_ToProto,
	)

	// Explicitly compare the KRM Spec fields with the GCP proto fields:
	// - loggingLogBucketRef (KRM Spec only): Resolved at identity / adapter creation time, maps to bucket name hierarchy.
	// - resourceID (KRM Spec only): Mapped to leaf part of name.
	// - description (KRM Spec): Mapped to GCP .description.
	// - bigqueryDatasetRef (KRM Spec only, commented out/not yet supported): GCP .bigquery_dataset is output-only and mapped in ObservedState.
	f.SpecField(".description")

	f.StatusField(".create_time")
	f.StatusField(".lifecycle_state")
	f.StatusField(".bigquery_dataset")

	f.Unimplemented_Identity(".name")

	return f
}
