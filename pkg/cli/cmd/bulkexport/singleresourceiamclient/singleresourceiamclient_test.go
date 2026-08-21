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

//go:build integration
// +build integration

package singleresourceiamclient_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/singleresourceiamclient"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	testiam "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/iam"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var mgr manager.Manager

func TestGetPolicy(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a DCL schema loader: %v", err)
	}
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		if fixture.Type != resourcefixture.ExternalRef {
			return false
		}
		return testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, tstContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		singleResourceClient := singleresourceiamclient.New(sysContext.TFProvider, sysContext.SMLoader)
		policy, err := singleResourceClient.GetPolicy(context.TODO(), tstContext.CreateUnstruct)
		if err != nil {
			t.Fatalf("unexpected error returned by GetPolicy: %v", err)
		}
		if policy.Name != tstContext.CreateUnstruct.GetName() {
			t.Fatalf("name mismatch: got '%v', want '%v'", policy.Name, tstContext.CreateUnstruct.GetName())
		}
		if policy.Spec.ResourceReference.External == "" {
			t.Fatalf("external reference is empty, expected a value")
		}
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testFunc)
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
