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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MetastoreServiceRef struct {
	// +required
	/* The self-link of an existing Dataproc Metastore service , when not managed by Config Connector. */
	External string `json:"external,omitempty"`
}

type MetastoreService struct {
	ProjectID string
	Location  string
	ServiceID string
}

func (s *MetastoreService) String() string {
	return "projects/" + s.ProjectID + "/locations/" + s.Location + "/services/" + s.ServiceID
}

func ResolveMetastoreServiceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *MetastoreServiceRef) (*MetastoreService, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		// External must be in form `projects/<projectID>/locations/<location>/services/<ServiceID>`.
		// see https://cloud.google.com/dataproc-metastore/docs/reference/rest/v1beta/projects.locations.services/get
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "services" {
			return &MetastoreService{
				ProjectID: tokens[1],
				Location:  tokens[3],
				ServiceID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of MetastoreService external=%q was not known (use projects/<projectID>/locations/<location>/services/<ServiceID>)", ref.External)
	}

	return nil, fmt.Errorf("must specify external on MetastoreServiceRef")
}
