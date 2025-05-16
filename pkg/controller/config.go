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

package controller

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Common controller dependencies.
type Deps struct {
	TfProvider        *tfschema.Provider
	TfLoader          *servicemappingloader.ServiceMappingLoader
	DclConfig         *dcl.Config
	DclConverter      *conversion.Converter
	Defaulters        []k8s.Defaulter
	JitterGen         jitter.Generator
	DependencyTracker *gcpwatch.DependencyTracker
}
