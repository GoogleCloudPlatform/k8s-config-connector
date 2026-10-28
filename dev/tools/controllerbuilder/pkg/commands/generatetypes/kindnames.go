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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/kindnames"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/sourcelinks"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// kindResource is a new resource generate-types is about to scaffold.
type kindResource struct {
	Kind  string
	Proto string
	// API is the API name from the proto's google.api.default_host, e.g.
	// sqladmin.
	API string
	// ResourceType is the proto's google.api.resource type, e.g.
	// sqladmin.googleapis.com/Instance.
	ResourceType string
}

// newKindResource fills in the API and resource type of a resource's proto.
// Either is "" when the proto does not say.
func newKindResource(api *protoapi.Proto, kind string, msg protoreflect.MessageDescriptor) kindResource {
	r := kindResource{Kind: kind, Proto: string(msg.FullName())}
	if opts := msg.Options(); opts != nil {
		if rd, ok := proto.GetExtension(opts, annotations.E_Resource).(*annotations.ResourceDescriptor); ok && rd != nil {
			r.ResourceType = rd.GetType()
		}
	}
	r.API = kindnames.APIForResource(sourcelinks.DefaultHost(msg.ParentFile(), api.Files()), r.ResourceType)
	return r
}

// kindNameMismatches returns a line for each resource that cannot be
// scaffolded with its Kind: the resource's API has no service prefix in either
// file, the API's service prefix is still waiting for approval, or the Kind
// differs from the one the table gives (the API's service prefix followed by
// the singular resource type). A resource whose proto has no
// google.api.resource type, or whose API cannot be determined, is not checked.
func kindNameMismatches(table *kindnames.Table, resources []kindResource) []string {
	var out []string
	for _, r := range resources {
		if r.API == "" || kindnames.ResourceType(r.ResourceType) == "" {
			continue
		}
		e, ok := table.Lookup(r.API)
		if !ok {
			out = append(out, fmt.Sprintf("  - resource %s (%s): API %s has no service prefix in %s or %s; propose one in %s", r.Kind, r.Proto, r.API, kindnames.FileName, kindnames.ExceptionsFileName, kindnames.ExceptionsFileName))
			continue
		}
		want, _ := table.Kind(r.API, r.ResourceType)
		if e.Pending {
			out = append(out, fmt.Sprintf("  - resource %s (%s): the service prefix of API %s (%s) is pending approval in %s", r.Kind, r.Proto, r.API, e.ServicePrefix, kindnames.ExceptionsFileName))
			continue
		}
		if want != r.Kind {
			out = append(out, fmt.Sprintf("  - resource %s (%s): API %s has service prefix %s, so the Kind should be %s", r.Kind, r.Proto, r.API, e.ServicePrefix, want))
		}
	}
	return out
}
