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
	"errors"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
)

type AssetStream interface {
	Next() (*asset.Asset, error)
	io.Closer
}

type FilteredAssetStream struct {
	assetStream   AssetStream
	shouldInclude func(*asset.Asset) bool
}

// Constructs a new asset stream with a filter function which determines if a given asset is included or not
// This is useful for filtering out asset types and resource types that we do not yet support
func NewFilteredAssetStream(stream AssetStream, shouldInclude func(*asset.Asset) bool) AssetStream {
	filteredStream := FilteredAssetStream{
		assetStream:   stream,
		shouldInclude: shouldInclude,
	}
	return &filteredStream
}

func (f *FilteredAssetStream) Close() error {
	return f.assetStream.Close()
}

func (f *FilteredAssetStream) Next() (*asset.Asset, error) {
	if f.shouldInclude == nil {
		return f.assetStream.Next()
	}
	for asset, err := f.assetStream.Next(); !errors.Is(err, io.EOF); asset, err = f.assetStream.Next() {
		if err != nil {
			return nil, err
		}
		if f.shouldInclude(asset) {
			return asset, nil
		}
	}
	return nil, io.EOF
}
