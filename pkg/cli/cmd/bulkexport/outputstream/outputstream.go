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

package outputstream

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/singleresourceiamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/commonparams"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewResourceByteStream(tfProvider *schema.Provider, params *parameters.Parameters, assetStream stream.AssetStream) (stream.ByteStream, error) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, fmt.Errorf("error creating service mapping loader: %w", err)
	}
	unstructuredStream, err := NewUnstructuredStream(params, assetStream, tfProvider, smLoader)
	if err != nil {
		return nil, err
	}
	return stream.NewByteStream(outputsink.ResourceFormat(params.ResourceFormat), unstructuredStream, smLoader, tfProvider)
}

func NewUnstructuredStream(params *parameters.Parameters, assetStream stream.AssetStream, provider *schema.Provider, smLoader *servicemappingloader.ServiceMappingLoader) (stream.UnstructuredStream, error) {
	ctx := context.TODO()
	httpClient, err := serviceclient.NewHTTPClient(ctx, params.OAuth2Token)
	if err != nil {
		return nil, fmt.Errorf("error creating http client: %w", err)
	}
	config, err := params.NewControllerConfig(ctx)
	if err != nil {
		return nil, err
	}
	serviceClient := serviceclient.NewServiceClient(httpClient)
	gcpClient := gcpclient.New(provider, smLoader)
	unstructuredResourceStream, err := stream.NewUnstructuredResourceStreamFromAssetStream(assetStream, gcpClient, provider, &serviceClient, config)
	if err != nil {
		return nil, fmt.Errorf("error creating unstructured resource stream: %w", err)
	}
	fixupStream := stream.NewUnstructuredResourceFixupStream(unstructuredResourceStream)
	if params.IAMFormat == commonparams.NoneIAMFormatOption {
		return fixupStream, nil
	}
	iamClient := singleresourceiamclient.New(provider, smLoader)
	iamFormat, err := commonparams.IAMFormatParamToStreamIAMFormat(params.IAMFormat)
	if err != nil {
		return nil, err
	}
	unstructuredResourceAndPolicyStream := stream.NewUnstructuredResourceAndIAMPolicyStream(fixupStream, iamClient, iamFormat, params.FilterDeletedIAMMembers)
	return unstructuredResourceAndPolicyStream, nil
}
