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

package bulkexport

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/errorhandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/filteredinputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/inputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/outputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Execute(ctx context.Context, params *parameters.Parameters) error {
	errorHandler, err := errorhandler.NewErrorHandler(params)
	if err != nil {
		return err
	}
	tfProvider, err := tf.NewProvider(ctx, params.OAuth2Token)
	if err != nil {
		return err
	}

	// Initialize direct controllers/exporters
	controllerConfig, err := params.NewControllerConfig(ctx)
	if err != nil {
		return err
	}
	if err := registry.Init(ctx, controllerConfig); err != nil {
		return err
	}

	assetStream, err := newFilteredAssetStream(ctx, params, tfProvider)
	if err != nil {
		return err
	}
	defer assetStream.Close()
	yamlStream, err := outputstream.NewResourceByteStream(tfProvider, params, assetStream)
	if err != nil {
		return err
	}
	recoverableStream := stream.NewRecoverableByteStream(yamlStream)
	outputSink, err := outputsink.New(tfProvider, params.Output, outputsink.ResourceFormat(params.ResourceFormat))
	if err != nil {
		return err
	}
	defer outputSink.Close()
	for bytes, unstructured, err := recoverableStream.Next(ctx); !errors.Is(err, io.EOF); bytes, unstructured, err = recoverableStream.Next(ctx) {
		if err != nil {
			if err := errorHandler.Handle(fmt.Errorf("error getting next YAML: %w", err)); err != nil {
				return err
			}
			continue
		}
		if err := outputSink.Receive(ctx, bytes, unstructured); err != nil {
			if err := errorHandler.Handle(err); err != nil {
				return err
			}
		}
	}
	return nil
}

func newFilteredAssetStream(ctx context.Context, params *parameters.Parameters, tfProvider *schema.Provider) (stream.AssetStream, error) {
	config, err := params.NewControllerConfig(ctx)
	if err != nil {
		return nil, err
	}
	assetStream, err := inputstream.NewAssetStream(params, os.Stdin)
	if err != nil {
		return nil, err
	}
	return filteredinputstream.NewFilteredAssetStream(ctx, assetStream, tfProvider, config)
}
