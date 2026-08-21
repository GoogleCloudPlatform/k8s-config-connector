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

package export

import (
	"context"
	"errors"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/outputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/stream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func Execute(ctx context.Context, params *parameters.Parameters) error {
	tfProvider, err := tf.NewProvider(ctx, params.GCPAccessToken)
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

	byteStream, err := outputstream.NewResourceByteStream(tfProvider, params)
	if err != nil {
		return err
	}
	recoverableStream := stream.NewRecoverableByteStream(byteStream)

	outputSink, err := outputsink.New(tfProvider, params.Output, outputsink.ResourceFormat(params.ResourceFormat))
	if err != nil {
		return err
	}
	defer outputSink.Close()
	for bytes, unstructured, err := recoverableStream.Next(ctx); !errors.Is(err, io.EOF); bytes, unstructured, err = recoverableStream.Next(ctx) {
		if err != nil {
			return err
		}
		if err := outputSink.Receive(ctx, bytes, unstructured); err != nil {
			return err
		}
	}
	return nil
}
