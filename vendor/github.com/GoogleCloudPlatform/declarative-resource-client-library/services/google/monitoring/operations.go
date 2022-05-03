// Copyright 2020 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package monitoring contains handwritten support code for the monitoring service.
package monitoring

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Operation can be parsed from the returned API operation and waited on.
type Operation struct {
	Name string `json:"name"`
}

// Wait waits for an Operation to complete by fetching the operation until it completes.
func (op *Operation) Wait(ctx context.Context, c *dcl.Config) error {
	if op.Name != "" {
		// Names come in the form "accessPolicies/{{name}}"
		parts := strings.Split(op.Name, "/")
		op.Name = parts[len(parts)-1]
	}
	return nil
}

// FetchName will fetch the operation and return the name of the resource created.
// Monitoring creates resources with machine generated names.
// It must be called after the resource has been created.
func (op *Operation) FetchName() (*string, error) {
	if op.Name == "" {
		return nil, fmt.Errorf("this operation (%s) has no name and probably hasn't been run before", op.Name)
	}
	return &op.Name, nil
}
