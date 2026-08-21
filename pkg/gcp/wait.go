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

package gcp

import (
	"fmt"
	"time"

	"google.golang.org/api/bigtableadmin/v2"
	"google.golang.org/api/cloudasset/v1"
	resourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/compute/v1"
	container "google.golang.org/api/container/v1beta1"
	"google.golang.org/api/redis/v1"
	"google.golang.org/api/servicenetworking/v1"
	"google.golang.org/api/spanner/v1"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
	"k8s.io/apimachinery/pkg/util/wait"
)

type AssetInventoryWaitCallback func(operation *cloudasset.Operation) error
type BigtableWaitCallback func(operation *bigtableadmin.Operation) error
type RedisWaitCallback func(operation *redis.Operation) error
type SpannerWaitCallback func(operation *spanner.Operation) error
type SQLWaitCallback func(operation *sqladmin.Operation) error
type ResourceManagerCallback func(operation *resourcemanager.Operation) error
type ComputeWaitCallback func(operation *compute.Operation) error
type ContainerWaitCallback func(operation *container.Operation) error
type ServiceNetworkingWaitCallback func(operation *servicenetworking.Operation) error

func WaitForAssetInventoryOperationDefaultTimeout(assetClient *cloudasset.Service,
	operation *cloudasset.Operation, projectNum string, callback AssetInventoryWaitCallback) (*cloudasset.Operation, error) {
	return WaitForAssetInventoryOperation(assetClient, operation, projectNum, 10*time.Second, 30*time.Minute, callback)
}

// The project number should be the project number for the credentials which are making this request
func WaitForAssetInventoryOperation(assetClient *cloudasset.Service, operation *cloudasset.Operation, projectNum string,
	interval, timeout time.Duration, callback AssetInventoryWaitCallback) (*cloudasset.Operation, error) {
	err := wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		request := assetClient.Operations.Get(operation.Name)
		// CAI needs this value to use end user credentials instead of a service account
		//   https://cloud.google.com/asset-inventory/docs/faq
		// when the export is run on a project, the project number is in the operation name, however, when the operation
		// is run on an organization or folder it is not in the name.
		request.Header().Add("X-Goog-User-Project", projectNum)
		newOp, err := request.Do()
		if err != nil {
			return false, fmt.Errorf("error getting operation %v: %w", operation.Name,
				err)
		}
		operation = newOp
		if callback != nil {
			if err = callback(operation); err != nil {
				return false, fmt.Errorf("error returned by wait callback: %w", err)
			}
		}
		return operation.Done, nil
	})
	return operation, err
}

func WaitForResourceManagerOperationDefaultTimeout(rmClient *resourcemanager.Service, operation *resourcemanager.Operation,
	callback ResourceManagerCallback) (*resourcemanager.Operation, error) {
	return WaitForResourceManagerOperation(rmClient, operation, 10*time.Second, 10*time.Minute, callback)
}

func WaitForResourceManagerOperation(rmClient *resourcemanager.Service, operation *resourcemanager.Operation,
	interval, timeout time.Duration, callback ResourceManagerCallback) (*resourcemanager.Operation, error) {
	err := wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		newOp, err := rmClient.Operations.Get(operation.Name).Do()
		if err != nil {
			return false, fmt.Errorf("error getting operation %v: %w", operation.Name,
				err)
		}
		operation = newOp
		if callback != nil {
			if err = callback(operation); err != nil {
				return false, fmt.Errorf("error returned by wait callback: %w", err)
			}
		}
		return operation.Done, nil
	})
	return operation, err
}

func WaitForComputeOperationDefaultTimeout(computeClient *compute.Service, operation *compute.Operation,
	projectID string, callback ComputeWaitCallback) (*compute.Operation, error) {
	return WaitForComputeOperation(computeClient, operation, projectID, 30*time.Second, 30*time.Minute, callback)
}

func WaitForComputeOperation(computeClient *compute.Service, operation *compute.Operation,
	projectID string, interval, timeout time.Duration, callback ComputeWaitCallback) (*compute.Operation, error) {
	err := wait.PollImmediate(interval, timeout, func() (done bool, err error) {
		var newOp *compute.Operation
		if operation.Zone != "" {
			zone := FullResourceNameToShortName(operation.Zone)
			newOp, err = computeClient.ZoneOperations.Get(projectID, zone, operation.Name).Do()
		} else if operation.Region != "" {
			region := FullResourceNameToShortName(operation.Region)
			newOp, err = computeClient.RegionOperations.Get(projectID, region, operation.Name).Do()
		} else {
			newOp, err = computeClient.GlobalOperations.Get(projectID, operation.Name).Do()
		}
		if err != nil {
			return false, fmt.Errorf("error getting operation %v/%v: %w", projectID, operation.Name,
				err)
		}
		operation = newOp
		if callback != nil {
			if err = callback(operation); err != nil {
				return false, fmt.Errorf("error returned by wait callback: %w", err)
			}
		}
		return operation.Status == "DONE", nil
	})
	return operation, err
}
