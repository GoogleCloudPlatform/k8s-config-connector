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

package grafeas

import (
	pb "google.golang.org/genproto/googleapis/grafeas/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(grafeasNoteFuzzer())
}

func grafeasNoteFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Note{},
		GrafeasNoteSpec_FromProto, GrafeasNoteSpec_ToProto,
		GrafeasNoteObservedState_FromProto, GrafeasNoteObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".secret")
	f.Unimplemented_NotYetTriaged(".vulnerability.advisory_publish_time")

	f.SpecField(".short_description")
	f.SpecField(".long_description")
	f.SpecField(".related_url")
	f.SpecField(".expiration_time")
	f.SpecField(".related_note_names")
	f.SpecField(".vulnerability")
	f.SpecField(".build")
	f.SpecField(".image")
	f.SpecField(".package")
	f.SpecField(".deployment")
	f.SpecField(".discovery")
	f.SpecField(".attestation")
	f.SpecField(".upgrade")
	f.SpecField(".compliance")
	f.SpecField(".dsse_attestation")
	f.SpecField(".vulnerability_assessment")
	f.SpecField(".sbom_reference")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".kind")

	return f
}
