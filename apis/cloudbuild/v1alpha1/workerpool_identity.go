// Copyright 2024 Google LLC
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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// WorkerPoolIdentity defines the resource reference to CloudBuildWorkerPool, which "External" field
// holds the GCP identifier for the KRM object.
type WorkerPoolIdentity struct {
	parent *WorkerPoolParent
	id     string
}

func (i *WorkerPoolIdentity) String() string {
	return i.parent.String() + "/workerPools/" + i.id
}

func (i *WorkerPoolIdentity) ID() string {
	return i.id
}

func (i *WorkerPoolIdentity) Location() string {
	return i.parent.Location
}

func (i *WorkerPoolIdentity) Parent() *WorkerPoolParent {
	return i.parent
}

type WorkerPoolParent struct {
	ProjectID string
	Location  string
}

func (p *WorkerPoolParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a WorkerPoolIdentity from the Config Connector WorkerPool object.
func NewWorkerPoolIdentity(ctx context.Context, reader client.Reader, obj *CloudBuildWorkerPool) (*WorkerPoolIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	return &WorkerPoolIdentity{
		parent: &WorkerPoolParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseWorkerPoolIdentityFromExternal(ctx context.Context, external string) (id *WorkerPoolIdentity, err error) {
	actualExternal := external
	// For backwards compatbility on:
	// Format: //cloudbuild.googleapis.com/projects/<project>/lcoations/<location>/workerPools/<id>
	if strings.HasPrefix(external, "//cloudbuild.googleapis.com/") {
		// log for visbility
		log.FromContext(ctx).Info("external url %s did not have service prefix %s", external, "//cloudbuild.googleapis.com/")
		actualExternal = strings.TrimPrefix(external, "//cloudbuild.googleapis.com/")
	}
	tokens := strings.Split(actualExternal, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workerPools" {
		return nil, fmt.Errorf("format of CloudBuildWorkerPool external=%q was not known (use projects/{{projectID}}/locations/{{location}}/workerPools/{{workerpoolID}})", external)
	}

	return &WorkerPoolIdentity{
		parent: &WorkerPoolParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		},
		id: tokens[5],
	}, nil
}
