// Copyright 2026 Google LLC
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

package shadow

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
)

func TestShadowEvaluation_ExistingControllers(t *testing.T) {
	root, err := options.RepoRoot()
	if err != nil {
		t.Fatalf("failed to find repo root: %v", err)
	}

	sampleRelativePaths := []string{
		"pkg/controller/direct/securesourcemanager/securesourcemanagerinstance_controller.go",
		"pkg/controller/direct/privilegedaccessmanager/privilegedaccessmanagerentitlement_controller.go",
		"pkg/controller/direct/saasservicemgmt/saasservicemgmtunit_controller.go",
		"pkg/controller/direct/cloudquota/apiquotapreference_controller.go",
		"pkg/controller/direct/dataform/dataformrepository_controller.go",
		"pkg/controller/direct/tags/tagstagkey_controller.go",
		"pkg/controller/direct/assuredworkloads/assuredworkloadsworkload_controller.go",
		"pkg/controller/direct/billing/billingaccount_controller.go",
	}

	var sampleFiles []string
	for _, rel := range sampleRelativePaths {
		sampleFiles = append(sampleFiles, filepath.Join(root, rel))
	}

	summary := BatchEvaluate(sampleFiles)
	if summary.TotalEvaluated == 0 || len(summary.Controllers) == 0 {
		t.Fatalf("no controllers evaluated")
	}

	t.Logf("=== Shadow Mode Batch Agreement Report ===")
	t.Logf("Total Controllers Evaluated: %d", len(summary.Controllers))
	t.Logf("Average Structural Agreement: %.1f%%", summary.AverageStructuralScore*100)
	t.Logf("RegisterModel Agreement:     %.1f%%", summary.RegisterModelRate*100)
	t.Logf("Find() Agreement:            %.1f%%", summary.FindAgreementRate*100)
	t.Logf("Create() Agreement:          %.1f%%", summary.CreateAgreementRate*100)
	t.Logf("Update() Agreement:          %.1f%%", summary.UpdateAgreementRate*100)
	t.Logf("Export() Agreement:          %.1f%%", summary.ExportAgreementRate*100)
	t.Logf("Delete() Agreement:          %.1f%%", summary.DeleteAgreementRate*100)

	t.Logf("\n--- Detailed Controller Breakdown ---")
	for _, ctrl := range summary.Controllers {
		baseName := filepath.Base(ctrl.FilePath)
		t.Logf("[%s] Structural Agreement: %.1f%% | Existing: %d lines, Shadow Generated: %d lines",
			baseName, ctrl.StructuralScore*100, ctrl.ExistingLines, ctrl.GeneratedLines)
		if len(ctrl.CustomFunctions) > 0 {
			t.Logf("   -> Custom / Hand-crafted extensions: %v", ctrl.CustomFunctions)
		}
	}

	if summary.AverageStructuralScore < 0.85 {
		t.Errorf("Average structural agreement %.2f below threshold 0.85", summary.AverageStructuralScore)
	}
}

func TestShadowEvaluation_BroadDirectControllers(t *testing.T) {
	root, err := options.RepoRoot()
	if err != nil {
		t.Fatalf("failed to find repo root: %v", err)
	}

	directDir := filepath.Join(root, "pkg/controller/direct")
	var candidateFiles []string

	err = filepath.Walk(directDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "_controller.go") {
			// Skip test or mock files
			if !strings.HasSuffix(info.Name(), "_test.go") {
				candidateFiles = append(candidateFiles, path)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking direct dir: %v", err)
	}

	summary := BatchEvaluate(candidateFiles)
	t.Logf("=== Repository-Wide Shadow Agreement (All %d Direct Controllers) ===", len(summary.Controllers))
	t.Logf("Average Structural Agreement: %.1f%%", summary.AverageStructuralScore*100)
	t.Logf("RegisterModel Rate:          %.1f%%", summary.RegisterModelRate*100)
	t.Logf("Find() Rate:                 %.1f%%", summary.FindAgreementRate*100)
	t.Logf("Create() Rate:               %.1f%%", summary.CreateAgreementRate*100)
	t.Logf("Update() Rate:               %.1f%%", summary.UpdateAgreementRate*100)
	t.Logf("Export() Rate:               %.1f%%", summary.ExportAgreementRate*100)
	t.Logf("Delete() Rate:               %.1f%%", summary.DeleteAgreementRate*100)
}
