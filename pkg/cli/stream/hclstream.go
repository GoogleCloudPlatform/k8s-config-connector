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
	"errors"
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/krmtohcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type HCLStream struct {
	unstructuredStream UnstructuredStream
	smLoader           *servicemappingloader.ServiceMappingLoader
	tfProvider         *schema.Provider
}

func NewHCLStream(unstructuredStream UnstructuredStream, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider) *HCLStream {
	hclStream := HCLStream{
		tfProvider:         tfProvider,
		smLoader:           smLoader,
		unstructuredStream: unstructuredStream,
	}
	return &hclStream
}

func (h *HCLStream) Next(ctx context.Context) ([]byte, *unstructured.Unstructured, error) {
	unstructured, err := h.unstructuredStream.Next(ctx)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			err = fmt.Errorf("error getting next asset: %w", err)
		}
		return nil, unstructured, err
	}
	hcl, err := krmtohcl.UnstructuredToHCL(ctx, unstructured, h.smLoader, h.tfProvider)
	if err != nil {
		return nil, unstructured, fmt.Errorf("error converting krm to hcl: %w", err)
	}
	return []byte(hcl), unstructured, nil
}
