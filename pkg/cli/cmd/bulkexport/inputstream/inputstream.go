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

package inputstream

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/log"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/storage"
)

var (
	requestTimeout = 10 * time.Second
	exportTimeout  = 2 * time.Minute
)

func NewAssetStream(params *parameters.Parameters, stdin *os.File) (*asset.Stream, error) {
	piped, err := parameters.IsInputPiped(stdin)
	if err != nil {
		return nil, fmt.Errorf("error checking if stdin has piped input: %w", err)
	}
	if piped {
		return asset.NewStream(stdin), nil
	}
	if params.Input != "" {
		return asset.NewStreamFromFile(params.Input)
	}
	return newStorageObjectStream(params)
}

func newStorageObjectStream(params *parameters.Parameters) (*asset.Stream, error) {
	var err error
	var bucketName, objectName string
	httpClient, err := serviceclient.NewHTTPClient(context.TODO(), params.OAuth2Token)
	if err != nil {
		return nil, fmt.Errorf("error creating http client: %w", err)
	}
	if params.StorageKey == "" {
		log.Verbose("Creating temporary bucket for export...")
		ctx, cancel := newRequestContext()
		defer cancel()
		bucketName, objectName, err = export.NewTemporaryBucketAndObjectName(ctx, httpClient)
		if err != nil {
			return nil, fmt.Errorf("error creating temporary bucket and prefix: %w", err)
		}
		defer deleteTemporaryBucket(httpClient, bucketName)
	} else {
		bucketName, objectName, err = storage.GetBucketAndPrefix(params.StorageKey)
		if err != nil {
			return nil, fmt.Errorf("error extracting bucket and prefix from parameter '%v': %w", parameters.StorageKeyParam, err)
		}
		if objectName == "" {
			objectName = export.NewObjectName()
		} else {
			return newStreamFromStorageObject(httpClient, bucketName, objectName)
		}
	}
	parent, err := getParentName(params)
	if err != nil {
		return nil, err
	}
	exportCtx, cancel := context.WithTimeout(context.Background(), exportTimeout)
	defer cancel()
	log.Verbose("Creating asset inventory export at %v", storage.GetFullURI(bucketName, objectName))
	if err := export.ForParentToStorageObject(exportCtx, httpClient, parent, bucketName, objectName); err != nil {
		return nil, fmt.Errorf("error exporting asset inventory: %w", err)
	}
	defer deleteExport(httpClient, bucketName, objectName)
	return newStreamFromStorageObject(httpClient, bucketName, objectName)
}

func newStreamFromStorageObject(httpClient *http.Client, bucketName, objectName string) (*asset.Stream, error) {
	ctx, cancel := newRequestContext()
	defer cancel()
	return asset.NewStreamFromStorageObject(ctx, httpClient, bucketName, objectName)
}

func deleteTemporaryBucket(httpClient *http.Client, bucketName string) {
	log.Verbose("Deleting temporary bucket '%v'...", bucketName)
	ctx, cancel := newRequestContext()
	defer cancel()
	if err := export.DeleteTemporaryBucket(ctx, httpClient, bucketName); err != nil {
		log.Verbose("Error deleting temporary bucket '%v': %v", bucketName, err)
	}
}

func deleteExport(httpClient *http.Client, bucketName, objectName string) {
	objectURI := storage.GetFullURI(bucketName, objectName)
	log.Verbose("Deleting export at %v...", objectURI)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	if err := export.DeleteExport(ctx, httpClient, bucketName, objectName); err != nil {
		log.Verbose("Error deleting export at %v: %v", objectURI, err)
	}
}

func getParentName(params *parameters.Parameters) (string, error) {
	if params.ProjectID != "" {
		return fmt.Sprintf("projects/%v", params.ProjectID), nil
	}
	if params.FolderID != 0 {
		return fmt.Sprintf("folders/%v", params.FolderID), nil
	}
	if params.OrganizationID != 0 {
		return fmt.Sprintf("organizations/%v", params.OrganizationID), nil
	}
	// technically the parameters validation methods guard against this case but we return a helpful error in case there
	// is a bug in the validation
	return "", fmt.Errorf("one of the '%v', '%v', or '%v' parameters must be defined when exporting",
		parameters.ProjectIDParam, parameters.FolderIDParam, parameters.OrganizationIDParam)
}

func newRequestContext() (context.Context, func()) {
	return context.WithTimeout(context.Background(), requestTimeout)
}
