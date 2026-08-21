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

package plan

import (
	"context"
	"io"
	"net/http"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
)

// PlanOptions holds options for a plan operation.
type PlanOptions struct {
	Out     io.Writer
	Objects []*unstructured.Unstructured

	RESTConfig *rest.Config
	HTTPClient *http.Client
}

// RunPlan executes a plan operation.
func RunPlan(ctx context.Context, opt *PlanOptions) error {
	target, err := buildTarget(ctx, opt.RESTConfig, opt.HTTPClient)
	if err != nil {
		return err
	}

	p := &Planner{}

	plan, err := p.BuildPlan(ctx, opt.Objects, target)
	if err != nil {
		return err
	}

	printPlan(ctx, plan, opt.Out)

	return nil
}

func buildTarget(ctx context.Context, restConfig *rest.Config, httpClient *http.Client) (*ClusterTarget, error) {
	return NewClusterTarget(restConfig, httpClient)
}
