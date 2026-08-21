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

package asset

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/storage/v1"
)

type Asset struct {
	Name      string
	AssetType string `json:"asset_type,omitempty"`
	Ancestors []string
}

type Stream struct {
	decoder   *json.Decoder
	srcReader io.Reader
}

func (a *Stream) Next() (*Asset, error) {
	var asset Asset
	err := a.decoder.Decode(&asset)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			err = fmt.Errorf("error decoding asset from reader: %w", err)
		}

		return nil, err
	}
	return &asset, nil
}

func (a *Stream) Close() error {
	// do not close stdin if it is the source
	if a.srcReader == os.Stdin {
		return nil
	}
	closer, ok := a.srcReader.(io.Closer)
	if !ok {
		return nil
	}
	return closer.Close()
}

func NewStream(r io.Reader) *Stream {
	reader := r
	if _, ok := r.(*bufio.Reader); !ok {
		reader = bufio.NewReader(r)
	}
	decoder := json.NewDecoder(reader)
	assetStream := Stream{
		decoder:   decoder,
		srcReader: r,
	}
	return &assetStream
}

func NewStreamFromFile(filePath string) (*Stream, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%v': %w", filePath, err)
	}
	return NewStream(file), nil
}

func NewStreamFromStorageObject(ctx context.Context, httpClient *http.Client, bucketName, objectName string) (*Stream, error) {
	storageClient, err := storage.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("error creating storage client: %w", err)
	}
	response, err := storageClient.Objects.Get(bucketName, objectName).Download()
	if err != nil {
		return nil, fmt.Errorf("error initiating download of %v/%v: %w", bucketName, objectName, err)
	}
	return NewStream(response.Body), nil
}
