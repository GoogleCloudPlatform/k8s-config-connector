/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudbuild

import (
	"fmt"
	"strings"

	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/externalresource"
)

const (
	serviceBaseURL = "https://cloudbuild.googleapis.com/v1/"
)

func NewResourceRef(gcpObj *cloudbuildpb.WorkerPool) (*ResourceRef, error) {
	baseResourceRef := externalresource.New(serviceBaseURL, gcpObj)
	extResRef := direct.ValueOf(baseResourceRef.Get())
	segments := strings.Split(extResRef, "/projects/")
	if len(segments) != 2 {
		return nil, fmt.Errorf("externalReference should be <baseUrl>/projects/<project>/locations/<location>/workerPools/<workerPool>, got %s",
			extResRef)
	}
	segments = strings.Split(segments[1], "/")
	if len(segments) == 5 && segments[1] == "locations" && segments[3] == "workerPools" {
		return &ResourceRef{
			project:     segments[0],
			location:    segments[2],
			resourceID:  segments[4],
			externalRef: baseResourceRef.Get(),
		}, nil
	}
	return nil, fmt.Errorf("externalReference should be in the form of <baseUrl>/projects/<project>/locations/<location>/workerPools/<workerPool>, got %s",
		extResRef)
}

type ResourceRef struct {
	resourceID  string
	location    string
	project     string
	externalRef *string
}

func (e *ResourceRef) GetExternalReference() *string {
	return e.externalRef
}

func (e *ResourceRef) GetResourceID() string {
	return e.resourceID
}

func (e *ResourceRef) GetLocation() string {
	return e.location
}

func (e *ResourceRef) GetProject() string {
	return e.project
}
