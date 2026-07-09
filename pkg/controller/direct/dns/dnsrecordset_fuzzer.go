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

package dns

import (
	api "google.golang.org/api/dns/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(dnsRecordSetFuzzer())
}

func dnsRecordSetFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.ResourceRecordSet{},
		DNSRecordSetSpec_FromAPI, DNSRecordSetSpec_ToAPI,
		DNSRecordSetStatus_FromAPI, DNSRecordSetStatus_ToAPI,
	)

	f.SpecField(".Name")
	f.SpecField(".RoutingPolicy")
	f.SpecField(".Rrdatas")
	f.SpecField(".Ttl")
	f.SpecField(".Type")

	// Top level unimplemented fields
	f.Unimplemented_NotYetTriaged(".Kind")
	f.Unimplemented_NotYetTriaged(".SignatureRrdatas")

	// RoutingPolicy top-level (not in slice) unimplemented fields
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.HealthCheck")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Kind")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Geo.Kind")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.PrimaryBackup.Kind")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.PrimaryBackup.BackupGeoTargets.Kind")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Wrr.Kind")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.PrimaryBackup.PrimaryTargets.ExternalEndpoints")

	// Nested fields inside slices of Geo, BackupGeoTargets, and Wrr items
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Geo.Items[].HealthCheckedTargets.ExternalEndpoints")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.PrimaryBackup.BackupGeoTargets.Items[].HealthCheckedTargets.ExternalEndpoints")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Wrr.Items[].HealthCheckedTargets.ExternalEndpoints")

	// Nested fields inside slices (Rrdatas and SignatureRrdatas are unmapped in routing policies)
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Geo.Items[].Rrdatas")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Geo.Items[].SignatureRrdatas")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Wrr.Items[].Rrdatas")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.Wrr.Items[].SignatureRrdatas")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.PrimaryBackup.BackupGeoTargets.Items[].Rrdatas")
	f.Unimplemented_NotYetTriaged(".RoutingPolicy.PrimaryBackup.BackupGeoTargets.Items[].SignatureRrdatas")

	return f
}
