// Copyright 2023 Google LLC
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

package main

import (
	"context"
	"fmt"
	"os"

	github "github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/commonoperator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/githubapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/reconcilers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/universe/github/pkg/reconcilers/githubreporeconciler"
)

func main() {
	ctx := ctrl.SetupSignalHandler()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set (try `export GITHUB_TOKEN=$(gh auth token)`)")
	}
	githubTokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	githubRawClient := github.NewClient(oauth2.NewClient(ctx, githubTokenSource))
	externalClient, err := githubapi.NewClient(ctx, githubRawClient)
	if err != nil {
		return err
	}

	op := commonoperator.Operator{}
	op.RegisterSchema(v1alpha1.AddToScheme)
	reposBridge := reconcilers.NewReconcilerBridge(&githubreporeconciler.GithubRepoReconciler{}, externalClient)
	op.RegisterReconciler(reposBridge)
	return op.Run(ctx, commonoperator.Options{LeaderElectionID: "github.experimental.cnrm.cloud.google.com"})
}
