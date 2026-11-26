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

package generatetypes

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// knownDeprecatedProtoPackages lists GA proto packages for Google Cloud
// products that have been deprecated or shut down, but whose .proto definitions
// in googleapis cannot be detected automatically because they omit
// `option deprecated = true;` on both the service block and its RPC methods.
var knownDeprecatedProtoPackages = map[string]string{
	// Cannot be detected from proto descriptors: google/cloud/automl/v1/service.proto
	// defines `service AutoMl` and 22 RPCs (CreateDataset, CreateModel, etc.),
	// and none of the service, RPCs, or resource messages in google/cloud/automl/v1/
	// set `option deprecated = true;`, even though Cloud AutoML was deprecated
	// and shut down in favor of Vertex AI (google.cloud.aiplatform.v1).
	"google.cloud.automl.v1": "Cloud AutoML (google.cloud.automl.v1) is deprecated in favor of Vertex AI (google.cloud.aiplatform.v1)",

	// Cannot be detected from proto descriptors: google/cloud/iot/v1/device_manager.proto
	// defines `service DeviceManager` and 14 RPCs (CreateDeviceRegistry,
	// CreateDevice, etc.), and none of the service, RPCs, or resource messages
	// set `option deprecated = true;`, even though Cloud IoT Core was deprecated
	// and retired in August 2023.
	"google.cloud.iot.v1": "Cloud IoT Core (google.cloud.iot.v1) is deprecated and retired",

	// Cannot be detected from proto descriptors: google/cloud/pubsublite/v1/admin.proto
	// defines `service AdminService` and 17 CRUD RPCs (CreateTopic,
	// CreateSubscription, CreateReservation, etc.) without `option deprecated = true;`
	// on the service or any RPC (only a single nested field in common.proto is
	// marked deprecated), even though Pub/Sub Lite was deprecated and retired.
	"google.cloud.pubsublite.v1": "Pub/Sub Lite (google.cloud.pubsublite.v1) is deprecated and retired",

	// Cannot be detected from proto descriptors: google/cloud/notebooks/v1/service.proto
	// (`service NotebookService`) and managed_service.proto (`service ManagedNotebookService`)
	// do not set `option deprecated = true;` on either service or on all of their
	// RPCs, even though Vertex AI Workbench legacy v1 user-managed and managed
	// notebooks were deprecated and shut down in favor of google.cloud.notebooks.v2.
	"google.cloud.notebooks.v1": "AI Platform / Vertex AI Workbench v1 (google.cloud.notebooks.v1) is deprecated in favor of google.cloud.notebooks.v2",

	// Shut down January 16, 2025 without `option deprecated = true;` on services/RPCs.
	"google.cloud.contentwarehouse.v1": "Document AI Warehouse (google.cloud.contentwarehouse.v1) was shut down on January 16, 2025",

	// End of sale July 30, 2024 and shut down June 30, 2025 without `option deprecated = true;`.
	"google.cloud.rapidmigrationassessment.v1": "Rapid Migration Assessment (google.cloud.rapidmigrationassessment.v1) shut down on June 30, 2025 in favor of Migration Center",

	// Deprecated June 15, 2026 and EOL September 30, 2026 without `option deprecated = true;`.
	"google.cloud.visionai.v1": "Vertex AI Vision (google.cloud.visionai.v1) reached EOL on September 30, 2026",

	// End of sale June 17, 2024 (closed to new organizations and standalone projects).
	"google.devtools.sourcerepo.v1": "Cloud Source Repositories (google.devtools.sourcerepo.v1) reached End of Sale on June 17, 2024",
}

// knownDeprecatedProtoMessages lists specific deprecated/non-creatable resource
// messages inside otherwise active GA proto packages when the resource itself
// (or its read RPC) omits `option deprecated = true;`.
var knownDeprecatedProtoMessages = map[string]string{
	"google.cloud.aiplatform.v1.DataLabelingJob":          "Vertex AI Data Labeling Service (DataLabelingJob) shut down on October 3, 2024",
	"google.cloud.backupdr.v1.ManagementServer":           "Backup and DR legacy Management Console (ManagementServer) ended new deployments on July 31, 2026",
	"google.cloud.iap.v1.Brand":                           "IAP OAuth Admin API (Brand) shut down on March 19, 2026",
	"google.cloud.iap.v1.IdentityAwareProxyClient":        "IAP OAuth Admin API (IdentityAwareProxyClient) shut down on March 19, 2026",
	"google.cloud.securitycenter.v1.OrganizationSettings": "SCC OrganizationSettings singleton only retains read access (UpdateOrganizationSettings is deprecated following Asset Discovery retirement on June 20, 2024)",
}

// deprecatedResourceReason reports whether a resource proto message is
// deprecated:
//  1. Its proto package is in knownDeprecatedProtoPackages (for retired GCP
//     products whose .proto descriptors omit `option deprecated = true;`), or
//  2. Its proto message full name is in knownDeprecatedProtoMessages (for
//     retired/non-creatable resources inside an otherwise active package), or
//  3. The resource message itself has `option deprecated = true;`, or
//  4. The proto service(s) managing the resource have `option deprecated = true;`
//     at the service level (or on all RPC methods in the service), or
//  5. All RPC methods managing this specific resource in a shared service have
//     `option deprecated = true;`.
func deprecatedResourceReason(api *protoapi.Proto, msg protoreflect.MessageDescriptor) (string, bool) {
	if api == nil || msg == nil {
		return "", false
	}
	pkg := msg.ParentFile().Package()
	if reason, ok := knownDeprecatedProtoPackages[string(pkg)]; ok {
		return reason, true
	}
	if reason, ok := knownDeprecatedProtoMessages[string(msg.FullName())]; ok {
		return reason, true
	}
	if opts, ok := msg.Options().(*descriptorpb.MessageOptions); ok && opts.GetDeprecated() {
		return fmt.Sprintf("message %s is marked deprecated in proto", msg.FullName()), true
	}

	services := managingServicesForResource(api, msg)
	if len(services) == 0 {
		return "", false
	}

	allServicesDeprecated := true
	var firstReason string
	for _, svc := range services {
		reason, deprecated := deprecatedServiceReason(svc)
		if !deprecated {
			allServicesDeprecated = false
			break
		}
		if firstReason == "" {
			firstReason = reason
		}
	}
	if allServicesDeprecated {
		return firstReason, true
	}

	return deprecatedResourceRPCsReason(api, msg)
}

// deprecatedServiceReason reports whether a proto service has
// `option deprecated = true;` on the service descriptor itself or on every RPC
// method in the service.
func deprecatedServiceReason(svc protoreflect.ServiceDescriptor) (string, bool) {
	if opts, ok := svc.Options().(*descriptorpb.ServiceOptions); ok && opts.GetDeprecated() {
		return fmt.Sprintf("service %s is marked deprecated in proto", svc.FullName()), true
	}
	n := svc.Methods().Len()
	if n == 0 {
		return "", false
	}
	for i := range n {
		m := svc.Methods().Get(i)
		opts, ok := m.Options().(*descriptorpb.MethodOptions)
		if !ok || !opts.GetDeprecated() {
			return "", false
		}
	}
	return fmt.Sprintf("all %d RPCs in service %s are marked deprecated in proto", n, svc.FullName()), true
}

// deprecatedResourceRPCsReason checks whether all RPCs managing msg across its
// proto package are marked deprecated (either on the RPC method itself or on
// its enclosing service). This catches deprecated resources inside an otherwise
// active shared service (such as google.cloud.edgenetwork.v1.Zone).
func deprecatedResourceRPCsReason(api *protoapi.Proto, msg protoreflect.MessageDescriptor) (string, bool) {
	pkg := msg.ParentFile().Package()
	var matchedRPCs []protoreflect.MethodDescriptor
	api.Files().RangeFilesByPackage(pkg, func(fd protoreflect.FileDescriptor) bool {
		for i := range fd.Services().Len() {
			svc := fd.Services().Get(i)
			for j := range svc.Methods().Len() {
				m := svc.Methods().Get(j)
				if methodManagesResource(m, msg) {
					matchedRPCs = append(matchedRPCs, m)
				}
			}
		}
		return true
	})
	if len(matchedRPCs) == 0 {
		return "", false
	}
	for _, m := range matchedRPCs {
		svcOpts, _ := m.Parent().(protoreflect.ServiceDescriptor).Options().(*descriptorpb.ServiceOptions)
		mOpts, _ := m.Options().(*descriptorpb.MethodOptions)
		if !svcOpts.GetDeprecated() && !mOpts.GetDeprecated() {
			return "", false
		}
	}
	return fmt.Sprintf("all %d RPCs managing %s in %s are marked deprecated in proto", len(matchedRPCs), msg.Name(), pkg), true
}

// managingServicesForResource returns the ServiceDescriptors in msg's proto
// package that manage msg via CRUD RPC names or request/response types. If no
// service in the package references msg via RPCs, it falls back to the services
// declared in msg's own .proto file.
func managingServicesForResource(api *protoapi.Proto, msg protoreflect.MessageDescriptor) []protoreflect.ServiceDescriptor {
	pkg := msg.ParentFile().Package()
	var matched []protoreflect.ServiceDescriptor
	api.Files().RangeFilesByPackage(pkg, func(fd protoreflect.FileDescriptor) bool {
		for i := range fd.Services().Len() {
			svc := fd.Services().Get(i)
			if serviceManagesResource(svc, msg) {
				matched = append(matched, svc)
			}
		}
		return true
	})
	if len(matched) > 0 {
		return matched
	}
	fd := msg.ParentFile()
	for i := range fd.Services().Len() {
		matched = append(matched, fd.Services().Get(i))
	}
	return matched
}

// serviceManagesResource reports whether any RPC method in svc manages msg.
func serviceManagesResource(svc protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) bool {
	for i := range svc.Methods().Len() {
		if methodManagesResource(svc.Methods().Get(i), msg) {
			return true
		}
	}
	return false
}

// methodManagesResource reports whether RPC method m manages msg, identified by
// standard CRUD method naming (`Create<Name>`, `Get<Name>`, `Update<Name>`,
// `Delete<Name>`), taking or returning msg directly, or taking a request
// message with a field of type msg (such as `Create<Name>Request`).
func methodManagesResource(m protoreflect.MethodDescriptor, msg protoreflect.MessageDescriptor) bool {
	name := string(msg.Name())
	fullName := msg.FullName()
	switch string(m.Name()) {
	case "Create" + name, "Get" + name, "Update" + name, "Delete" + name:
		return true
	}
	if m.Input().FullName() == fullName || m.Output().FullName() == fullName {
		return true
	}
	fields := m.Input().Fields()
	for j := range fields.Len() {
		f := fields.Get(j)
		if f.Kind() == protoreflect.MessageKind && f.Message().FullName() == fullName {
			return true
		}
	}
	return false
}
