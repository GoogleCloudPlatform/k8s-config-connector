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

	// Filter out nested slice fields (Rrdatas and SignatureRrdatas) since they are not mapped
	// in Geo/Wrr item mappers (which use RrdatasRefs in KRM), but cannot be easily matched via
	// fieldOverrides because they reset to ".Rrdatas" and conflict with the top-level spec field ".Rrdatas".
	filter := func(in *api.ResourceRecordSet) {
		if in.RoutingPolicy != nil {
			if in.RoutingPolicy.Geo != nil {
				for _, item := range in.RoutingPolicy.Geo.Items {
					item.Rrdatas = nil
					item.SignatureRrdatas = nil
				}
			}
			if in.RoutingPolicy.Wrr != nil {
				for _, item := range in.RoutingPolicy.Wrr.Items {
					item.Rrdatas = nil
					item.SignatureRrdatas = nil
				}
			}
			if in.RoutingPolicy.PrimaryBackup != nil {
				if in.RoutingPolicy.PrimaryBackup.BackupGeoTargets != nil {
					for _, item := range in.RoutingPolicy.PrimaryBackup.BackupGeoTargets.Items {
						item.Rrdatas = nil
						item.SignatureRrdatas = nil
					}
				}
			}
		}
	}
	f.FilterSpec = filter
	f.FilterStatus = filter

	return f
}
