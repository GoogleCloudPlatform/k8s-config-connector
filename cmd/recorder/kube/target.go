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

package kube

import (
	"fmt"
	"net/http"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// Target is a wrapper around a kubernetes client.
type Target struct {
	dynamicWatchers *dynamicWatchers
}

// NewTarget constructs a Target.
func NewTarget(restConfig *rest.Config, httpClient *http.Client) (*Target, error) {
	dynamicClient, err := dynamic.NewForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error building dynamic client: %w", err)
	}
	target := &Target{}
	target.dynamicWatchers = newDynamicWatchers(dynamicClient)
	return target, nil
}
