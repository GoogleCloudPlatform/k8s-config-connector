/*
Copyright 2022 The Kubernetes Authors.

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

package mockkubeapiserver

import (
	"context"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

// resourceRequestBase holds the common field for single-resource requests
type resourceRequestBase struct {
	baseRequest

	Group     string
	Version   string
	Resource  string
	Namespace string
	Name      string

	SubResource string
}

// getResource is a request to get a single resource
type getResource struct {
	resourceRequestBase
}

// Run serves the http request
func (req *getResource) Run(ctx context.Context, s *MockKubeAPIServer) error {
	gr := schema.GroupResource{Group: req.Group, Resource: req.Resource}
	resource := s.storage.FindResource(gr)
	if resource == nil {
		return req.writeErrorResponse(http.StatusNotFound)
	}

	id := types.NamespacedName{Namespace: req.Namespace, Name: req.Name}

	object, found, err := s.storage.GetObject(ctx, resource, id)
	if err != nil {
		return err
	}
	if !found {
		return req.writeErrorResponse(http.StatusNotFound)
	}

	return req.writeResponse(object)
}
