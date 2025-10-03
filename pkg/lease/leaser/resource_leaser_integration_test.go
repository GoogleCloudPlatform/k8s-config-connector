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

package leaser_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testkrmtotf "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func TestObtainAndReleaseResourceLease(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		// only need to test contention for a single resource since the logic will apply to all resources
		return fixture.GVK.Kind == "BigQueryDataset" && fixture.Type == resourcefixture.Basic
	}
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext) {
		u := testContext.CreateUnstruct
		// create the resource with a no management conflict policy to prevent the controller from obtaining a lease on the resource
		k8s.SetAnnotation(managementconflict.FullyQualifiedAnnotation, managementconflict.ManagementConflictPreventionPolicyNone, u)
		if err := systemContext.Manager.GetClient().Create(ctx, u); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		resourceCleanup := systemContext.Reconciler.BuildCleanupFunc(ctx, testContext.CreateUnstruct, testreconciler.CleanupPolicyAlways)
		defer resourceCleanup()
		systemContext.Reconciler.Reconcile(ctx, u, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, u), nil)
		sm, err := systemContext.SMLoader.GetServiceMapping(u.GroupVersionKind().Group)
		if err != nil {
			t.Fatalf("error getting service mapping: %v", err)
		}
		krmResource1 := testkrmtotf.NewKRMResource(t, u, sm, systemContext.TFProvider)
		krmResource2 := testkrmtotf.NewKRMResource(t, u, sm, systemContext.TFProvider)
		krmResource2.SetNamespace(testvariable.NewUniqueID())
		resourceLeaser := leaser.NewResourceLeaser(systemContext.TFProvider, systemContext.SMLoader, systemContext.Manager.GetClient())
		liveState1 := testkrmtotf.FetchLiveState(t, krmResource1, systemContext.TFProvider, systemContext.Manager.GetClient(), systemContext.SMLoader)
		// obtain a lease for the first namespace
		liveLabels1 := krmtotf.GetLabelsFromState(krmResource1, liveState1)
		if err := resourceLeaser.SoftObtain(ctx, &krmResource1.Resource, liveLabels1); err != nil {
			t.Fatalf("error obtaining lease on first namespace: %v", err)
		}
		apply(t, context.Background(), systemContext.TFProvider, systemContext.Manager.GetClient(), krmResource1, liveState1, systemContext.SMLoader)
		// try to obtain a lease for the second namespace, verify this fails
		liveState2 := testkrmtotf.FetchLiveState(t, krmResource2, systemContext.TFProvider, systemContext.Manager.GetClient(), systemContext.SMLoader)
		liveLabels2 := krmtotf.GetLabelsFromState(krmResource2, liveState2)
		if err := resourceLeaser.SoftObtain(ctx, &krmResource2.Resource, liveLabels2); err == nil {
			t.Fatal("expected error when obtaining lease for second namespace, instead got 'nil'")
		}
		// release the lease for the first namespace
		if err := resourceLeaser.Release(ctx, u); err != nil {
			t.Fatalf("error releasing lease on first namespace: %v", err)
		}
		// now that the resource has been released, obtain a lease for the second namespace
		liveState2 = testkrmtotf.FetchLiveState(t, krmResource2, systemContext.TFProvider, systemContext.Manager.GetClient(), systemContext.SMLoader)
		liveLabels2 = krmtotf.GetLabelsFromState(krmResource2, liveState2)
		if err := resourceLeaser.SoftObtain(ctx, &krmResource2.Resource, liveLabels2); err != nil {
			t.Fatalf("error obtaining lease for second namespace after it was released by the first: %v", err)
		}
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testFunc)
}

func apply(t *testing.T, ctx context.Context, provider *schema.Provider, kubeClient client.Client, resource *krmtotf.Resource, liveState *terraform.InstanceState, smLoader *servicemappingloader.ServiceMappingLoader) {
	config, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, kubeClient, smLoader)
	if err != nil {
		t.Fatalf("error converting resource to tf config: %v", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, config, provider.Meta())
	if err != nil {
		t.Fatalf("error diffing changes: %v", err)
	}
	if diff.Empty() {
		t.Fatalf("unexpected empty diff")
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		t.Fatalf("error applying changes: %v", err)
	}
}
