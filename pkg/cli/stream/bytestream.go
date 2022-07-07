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

package stream

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ByteStream interface {
	Next(ctx context.Context) ([]byte, *unstructured.Unstructured, error)
}

func NewByteStream(resourceFormat outputsink.ResourceFormat, uStream UnstructuredStream, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider) (ByteStream, error) {
	switch resourceFormat {
	case outputsink.KRMResourceFormat:
		return NewYAMLStream(uStream), nil
	case outputsink.HCLResourceFormat:
		return NewHCLStream(uStream, smLoader, tfProvider), nil
	default:
		return nil, fmt.Errorf("unhandled resource format '%v'", resourceFormat)
	}
}
