//go:build integration
// +build integration

package kccmanager_test

import (
	"context"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

func TestSkipNameValidation(t *testing.T) {
	ctx := context.TODO()
	// Use the config from the existing clusterModeManager which is set up in TestMain
	cfg := clusterModeManager.GetConfig()

	// Helper to create manager
	createManager := func(skip bool) (manager.Manager, error) {
		kccConfig := kccmanager.Config{
			ManagerOptions: manager.Options{
				// disable prometheus metrics as by default, the metrics server binds to the same port in all instances
				Metrics: server.Options{
					BindAddress: "0",
				},
			},
			SkipNameValidation: skip,
		}
		return kccmanager.New(ctx, cfg, kccConfig)
	}

	t.Run("FailsOnDuplicateControllerNameWithoutSkip", func(t *testing.T) {
		// First manager creation should succeed
		_, err := createManager(false)
		if err != nil {
			t.Fatalf("failed to create first manager: %v", err)
		}

		// Second manager creation should fail because "registration-controller" is already registered
		// and SkipNameValidation is false.
		_, err = createManager(false)
		if err == nil {
			// If it succeeds, it might be because the test environment behaves differently than expected.
			// However, given the feedback, we expect an error.
			// Let's log warning if it doesn't fail, but ideally we want to assert failure.
			// Since I see TestSchemeIsUniqueAcrossManagers passing in the codebase, maybe there is something subtle.
			// But I will follow instructions.
			t.Fatal("expected error when creating second manager without SkipNameValidation, but got success")
		} else {
			if !strings.Contains(err.Error(), "already exists") {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})

	t.Run("SucceedsWithSkipNameValidation", func(t *testing.T) {
		// First manager creation with skip=true
		_, err := createManager(true)
		if err != nil {
			t.Fatalf("failed to create first manager with SkipNameValidation: %v", err)
		}

		// Second manager creation with skip=true should also succeed
		_, err = createManager(true)
		if err != nil {
			t.Fatalf("failed to create second manager with SkipNameValidation: %v", err)
		}
	})
}
