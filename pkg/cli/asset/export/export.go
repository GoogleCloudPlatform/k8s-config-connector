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
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"

	"google.golang.org/api/cloudasset/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/storage/v1"
)

// ForParentToStorageObject performs an asset inventory export to the GCS 'bucketName' bucket at 'objectName'
// see additional methods in this package for creating a temporary bucket and generating an object name
//
// parent is the fully qualified org, folder, or project, i.e.
// projects/projectID
// folders/folderNumber
// organizations/orgNumber
func ForParentToStorageObject(ctx context.Context, httpClient *http.Client, parent, bucketName, objectName string) error {
	gcsDestination := fmt.Sprintf("gs://%v/%v", bucketName, objectName)
	exportAssetsRequest := cloudasset.ExportAssetsRequest{
		OutputConfig: &cloudasset.OutputConfig{
			GcsDestination: &cloudasset.GcsDestination{
				Uri: gcsDestination,
			},
		},
	}
	assetClient, err := newAssetInventoryClient(ctx, httpClient)
	if err != nil {
		return fmt.Errorf("error creating asset client: %w", err)
	}
	projectNum, err := getDefaultProjectNumber(ctx, httpClient)
	if err != nil {
		return fmt.Errorf("error getting project number: %w", err)
	}
	projectNumString := fmt.Sprint(projectNum)
	request := assetClient.V1.ExportAssets(parent, &exportAssetsRequest)
	request.Header().Add("X-Goog-User-Project", projectNumString)
	op, err := request.Do()
	if err != nil {
		return fmt.Errorf("error response from exportassets request: %w", err)
	}
	if _, err := gcp.WaitForAssetInventoryOperationDefaultTimeout(assetClient, op, projectNumString, nil); err != nil {
		return fmt.Errorf("error waiting for operation: %w", err)
	}
	return nil
}

func NewTemporaryBucketAndObjectName(ctx context.Context, httpClient *http.Client) (string, string, error) {
	storageClient, err := newStorageClient(ctx, httpClient)
	if err != nil {
		return "", "", fmt.Errorf("error creating storage client: %w", err)
	}
	projectID, err := gcp.GetDefaultProjectID()
	if err != nil {
		return "", "", fmt.Errorf("error getting project id: %w", err)
	}
	bucketName := fmt.Sprintf("export-%v", randomid.New().String())
	bucket := &storage.Bucket{
		Name: bucketName,
	}
	_, err = storageClient.Buckets.Insert(projectID, bucket).Do()
	if err != nil {
		return "", "", fmt.Errorf("error creating bucket '%v': %w", bucketName, err)
	}
	objectName := NewObjectName()
	return bucketName, objectName, nil
}

func NewObjectName() string {
	return randomid.New().String()
}

func DeleteExport(ctx context.Context, httpClient *http.Client, bucketName, objectName string) error {
	storageClient, err := newStorageClient(ctx, httpClient)
	if err != nil {
		return fmt.Errorf("error creating storage client: %w", err)
	}
	if err := storageClient.Objects.Delete(bucketName, objectName).Do(); err != nil {
		return fmt.Errorf("error deleting storage object gs://%v/%v: %w", bucketName, objectName, err)
	}
	return err
}

func DeleteTemporaryBucket(ctx context.Context, httpClient *http.Client, bucketName string) error {
	storageClient, err := newStorageClient(ctx, httpClient)
	if err != nil {
		return fmt.Errorf("error creating storage client: %w", err)
	}
	if err := storageClient.Buckets.Delete(bucketName).Do(); err != nil {
		return fmt.Errorf("error deleting temporary bucket '%v': %w", bucketName, err)
	}
	return nil
}

func newAssetInventoryClient(ctx context.Context, httpClient *http.Client) (*cloudasset.Service, error) {
	client, err := cloudasset.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("error creating asset inventory client: %w", err)
	}
	return client, nil
}

func newStorageClient(ctx context.Context, httpClient *http.Client) (*storage.Service, error) {
	client, err := storage.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("error creating storage client: %w", err)
	}
	return client, nil
}

func getDefaultProjectNumber(ctx context.Context, httpClient *http.Client) (int64, error) {
	client, err := serviceclient.NewResourceManagerClient(ctx, httpClient)
	if err != nil {
		return 0, err
	}
	// get the default project id by shelling out to gcloud, this is questionable but *OK* because it is only reading
	// local configuration files
	projectID, err := gcp.GetDefaultProjectID()
	if err != nil {
		return 0, fmt.Errorf("error getting default gcloud sdk project id: %w", err)
	}
	project, err := client.Projects.Get(projectID).Do()
	if err != nil {
		return 0, fmt.Errorf("error getting default project '%v': %w", projectID, err)
	}
	return project.ProjectNumber, nil
}
