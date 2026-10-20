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
}

// deprecatedResourceReason reports whether a resource proto message belongs to
// a deprecated service:
//  1. Its proto package is in knownDeprecatedProtoPackages (for retired GCP
//     products whose .proto descriptors omit `option deprecated = true;`), or
//  2. The proto service(s) managing the resource have `option deprecated = true;`
//     at the service level, or have `option deprecated = true;` on ALL RPC
//     methods in the service.
func deprecatedResourceReason(api *protoapi.Proto, msg protoreflect.MessageDescriptor) (string, bool) {
	if msg == nil {
		return "", false
	}
	pkg := msg.ParentFile().Package()
	if reason, ok := knownDeprecatedProtoPackages[string(pkg)]; ok {
		return reason, true
	}
	if api == nil {
		return "", false
	}

	services := managingServicesForResource(api, msg)
	if len(services) == 0 {
		return "", false
	}

	var firstReason string
	for _, svc := range services {
		reason, deprecated := deprecatedServiceReason(svc)
		if !deprecated {
			return "", false
		}
		if firstReason == "" {
			firstReason = reason
		}
	}
	return firstReason, true
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
	for i := 0; i < n; i++ {
		m := svc.Methods().Get(i)
		opts, ok := m.Options().(*descriptorpb.MethodOptions)
		if !ok || !opts.GetDeprecated() {
			return "", false
		}
	}
	return fmt.Sprintf("all %d RPCs in service %s are marked deprecated in proto", n, svc.FullName()), true
}

// managingServicesForResource returns the ServiceDescriptors in msg's proto
// package that manage msg via CRUD RPC names or request/response types. If no
// service in the package references msg via RPCs, it falls back to the services
// declared in msg's own .proto file.
func managingServicesForResource(api *protoapi.Proto, msg protoreflect.MessageDescriptor) []protoreflect.ServiceDescriptor {
	pkg := msg.ParentFile().Package()
	var matched []protoreflect.ServiceDescriptor
	api.Files().RangeFilesByPackage(pkg, func(fd protoreflect.FileDescriptor) bool {
		for i := 0; i < fd.Services().Len(); i++ {
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
	for i := 0; i < fd.Services().Len(); i++ {
		matched = append(matched, fd.Services().Get(i))
	}
	return matched
}

// serviceManagesResource reports whether any RPC method in svc manages msg,
// identified by standard CRUD method naming (`Create<Name>`, `Get<Name>`,
// `Update<Name>`, `Delete<Name>`), taking or returning msg directly, or taking
// a request message with a field of type msg (such as `Create<Name>Request`).
func serviceManagesResource(svc protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) bool {
	name := string(msg.Name())
	fullName := msg.FullName()
	for i := 0; i < svc.Methods().Len(); i++ {
		m := svc.Methods().Get(i)
		switch string(m.Name()) {
		case "Create" + name, "Get" + name, "Update" + name, "Delete" + name:
			return true
		}
		if m.Input().FullName() == fullName || m.Output().FullName() == fullName {
			return true
		}
		fields := m.Input().Fields()
		for j := 0; j < fields.Len(); j++ {
			f := fields.Get(j)
			if f.Kind() == protoreflect.MessageKind && f.Message().FullName() == fullName {
				return true
			}
		}
	}
	return false
}
