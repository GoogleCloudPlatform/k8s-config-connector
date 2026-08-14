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

package composer

import (
	"go/ast"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode"

	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// buildPatches is a test helper that evaluates fieldUpdaters against desired and actual,
// collecting the generated patch protos keyed by update mask.
func buildPatches(desired *krm.ComposerEnvironment, desiredPb, actualPb *composerpb.Environment) map[string]*composerpb.Environment {
	updates := make(map[string]*composerpb.Environment)
	for _, u := range fieldUpdaters {
		if patch := u.build(desired, desiredPb, actualPb); patch != nil {
			updates[u.mask] = patch
		}
	}
	return updates
}

func TestFieldUpdatersBuild(t *testing.T) {
	t.Run("no diff when desired matches actual", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Labels: map[string]string{"env": "prod"},
				Config: &krm.EnvironmentConfig{
					EnvironmentSize: direct.LazyPtr("ENVIRONMENT_SIZE_SMALL"),
				},
			},
		}
		actual := &composerpb.Environment{
			Labels: map[string]string{"env": "prod"},
			Config: &composerpb.EnvironmentConfig{
				EnvironmentSize: composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL,
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		updates := buildPatches(desired, desiredPb, actual)
		if len(updates) != 0 {
			t.Errorf("expected 0 pending updates, got %d", len(updates))
		}
	})

	t.Run("detects scheduled snapshots config update", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					RecoveryConfig: &krm.RecoveryConfig{
						ScheduledSnapshotsConfig: &krm.ScheduledSnapshotsConfig{
							Enabled:                  direct.LazyPtr(true),
							SnapshotCreationSchedule: direct.LazyPtr("20 15 * * *"),
							SnapshotLocation:         direct.LazyPtr("gs://my-bucket/snapshots"),
							TimeZone:                 direct.LazyPtr("UTC"),
						},
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				RecoveryConfig: &composerpb.RecoveryConfig{
					ScheduledSnapshotsConfig: &composerpb.ScheduledSnapshotsConfig{
						Enabled: false,
					},
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		updates := buildPatches(desired, desiredPb, actual)
		if len(updates) != 1 {
			t.Fatalf("expected 1 pending update, got %d", len(updates))
		}
		patch, ok := updates["config.recovery_config.scheduled_snapshots_config"]
		if !ok {
			t.Fatalf("expected update for 'config.recovery_config.scheduled_snapshots_config', but was missing")
		}
		if !patch.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig().GetEnabled() {
			t.Errorf("expected enabled to be true in patch")
		}
		if patch.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig().GetSnapshotCreationSchedule() != "20 15 * * *" {
			t.Errorf("expected schedule to be '20 15 * * *', got %q", patch.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig().GetSnapshotCreationSchedule())
		}
	})

	t.Run("detects workloads config update", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					WorkloadsConfig: &krm.WorkloadsConfig{
						Triggerer: &krm.WorkloadsConfig_TriggererResource{
							Count: direct.LazyPtr(int32(2)),
						},
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				WorkloadsConfig: &composerpb.WorkloadsConfig{
					Triggerer: &composerpb.WorkloadsConfig_TriggererResource{
						Count: 1,
					},
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		updates := buildPatches(desired, desiredPb, actual)
		if len(updates) != 1 {
			t.Fatalf("expected 1 pending update, got %d", len(updates))
		}
		patch, ok := updates["config.workloads_config"]
		if !ok {
			t.Fatalf("expected update for 'config.workloads_config', but was missing")
		}
		if patch.GetConfig().GetWorkloadsConfig().GetTriggerer().GetCount() != 2 {
			t.Errorf("expected triggerer count to be 2, got %d", patch.GetConfig().GetWorkloadsConfig().GetTriggerer().GetCount())
		}
	})

	t.Run("detects labels and maintenance window updates simultaneously", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Labels: map[string]string{"env": "staging"},
				Config: &krm.EnvironmentConfig{
					MaintenanceWindow: &krm.MaintenanceWindow{
						Recurrence: direct.LazyPtr("FREQ=WEEKLY;BYDAY=SA,SU"),
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Labels: map[string]string{"env": "prod"},
			Config: &composerpb.EnvironmentConfig{
				MaintenanceWindow: &composerpb.MaintenanceWindow{
					Recurrence: "FREQ=WEEKLY;BYDAY=FR,SA,SU",
					StartTime:  timestamppb.Now(),
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		updates := buildPatches(desired, desiredPb, actual)
		expected := []string{"labels", "config.maintenance_window"}
		if len(updates) != len(expected) {
			t.Fatalf("expected %d updates, got %d", len(expected), len(updates))
		}
		for _, m := range expected {
			if _, ok := updates[m]; !ok {
				t.Errorf("expected update for mask %q, but was missing", m)
			}
		}
	})

	t.Run("ignores unset spec fields against actual state", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				// No labels or config specified
			},
		}
		actual := &composerpb.Environment{
			Labels: map[string]string{"default-label": "val"},
			Config: &composerpb.EnvironmentConfig{
				EnvironmentSize: composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_SMALL,
				WorkloadsConfig: &composerpb.WorkloadsConfig{
					Scheduler: &composerpb.WorkloadsConfig_SchedulerResource{
						Count: 1,
					},
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		updates := buildPatches(desired, desiredPb, actual)
		if len(updates) != 0 {
			t.Errorf("expected 0 pending updates for unset spec, got %d", len(updates))
		}
	})

	t.Run("detects workloadsConfig, scheduledSnapshotsConfig, and resilienceMode updates simultaneously", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					ResilienceMode: direct.LazyPtr("HIGH_RESILIENCE"),
					RecoveryConfig: &krm.RecoveryConfig{
						ScheduledSnapshotsConfig: &krm.ScheduledSnapshotsConfig{
							Enabled:                  direct.LazyPtr(true),
							SnapshotCreationSchedule: direct.LazyPtr("0 5 * * *"),
							SnapshotLocation:         direct.LazyPtr("gs://my-bucket/snapshots"),
							TimeZone:                 direct.LazyPtr("UTC"),
						},
					},
					WorkloadsConfig: &krm.WorkloadsConfig{
						Scheduler: &krm.WorkloadsConfig_SchedulerResource{
							CPU:       direct.LazyPtr("1.25"),
							MemoryGB:  direct.LazyPtr("2.5"),
							StorageGB: direct.LazyPtr("2"),
							Count:     direct.LazyPtr(int32(2)),
						},
						WebServer: &krm.WorkloadsConfig_WebServerResource{
							CPU:       direct.LazyPtr("1"),
							MemoryGB:  direct.LazyPtr("3"),
							StorageGB: direct.LazyPtr("2"),
						},
						Worker: &krm.WorkloadsConfig_WorkerResource{
							CPU:       direct.LazyPtr("1.25"),
							MemoryGB:  direct.LazyPtr("2.5"),
							StorageGB: direct.LazyPtr("2"),
							MinCount:  direct.LazyPtr(int32(2)),
							MaxCount:  direct.LazyPtr(int32(6)),
						},
						Triggerer: &krm.WorkloadsConfig_TriggererResource{
							Count:    direct.LazyPtr(int32(2)),
							CPU:      direct.LazyPtr("0.75"),
							MemoryGB: direct.LazyPtr("1.5"),
						},
						DagProcessor: &krm.WorkloadsConfig_DagProcessorResource{
							Count:     direct.LazyPtr(int32(2)),
							CPU:       direct.LazyPtr("1.5"),
							MemoryGB:  direct.LazyPtr("4"),
							StorageGB: direct.LazyPtr("2"),
						},
					},
				},
			},
		}

		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				ResilienceMode: composerpb.EnvironmentConfig_RESILIENCE_MODE_UNSPECIFIED,
				RecoveryConfig: &composerpb.RecoveryConfig{
					ScheduledSnapshotsConfig: &composerpb.ScheduledSnapshotsConfig{
						Enabled: false,
					},
				},
				WorkloadsConfig: &composerpb.WorkloadsConfig{
					Scheduler: &composerpb.WorkloadsConfig_SchedulerResource{
						Cpu: 0.5, MemoryGb: 1.875, StorageGb: 1, Count: 1,
					},
					WebServer: &composerpb.WorkloadsConfig_WebServerResource{
						Cpu: 0.5, MemoryGb: 1.875, StorageGb: 1,
					},
					Worker: &composerpb.WorkloadsConfig_WorkerResource{
						Cpu: 0.5, MemoryGb: 1.875, StorageGb: 1, MinCount: 1, MaxCount: 3,
					},
					Triggerer: &composerpb.WorkloadsConfig_TriggererResource{
						Cpu: 0.5, MemoryGb: 0.5, Count: 1,
					},
					DagProcessor: &composerpb.WorkloadsConfig_DagProcessorResource{
						Cpu: 0.5, MemoryGb: 1.875, StorageGb: 1, Count: 1,
					},
				},
			},
		}

		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		updates := buildPatches(desired, desiredPb, actual)

		if len(updates) != 3 {
			t.Fatalf("expected 3 patches, got %d", len(updates))
		}

		resiliencePatch, ok := updates["config.resilience_mode"]
		if !ok {
			t.Errorf("missing update for 'config.resilience_mode'")
		} else if resiliencePatch.GetConfig().GetResilienceMode() != composerpb.EnvironmentConfig_HIGH_RESILIENCE {
			t.Errorf("expected HIGH_RESILIENCE, got %v", resiliencePatch.GetConfig().GetResilienceMode())
		}

		snapshotPatch, ok := updates["config.recovery_config.scheduled_snapshots_config"]
		if !ok {
			t.Errorf("missing update for 'config.recovery_config.scheduled_snapshots_config'")
		} else if !snapshotPatch.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig().GetEnabled() {
			t.Errorf("expected snapshot enabled=true, got false")
		}

		workloadPatch, ok := updates["config.workloads_config"]
		if !ok {
			t.Errorf("missing update for 'config.workloads_config'")
		} else if workloadPatch.GetConfig().GetWorkloadsConfig().GetWorker().GetMaxCount() != 6 {
			t.Errorf("expected worker maxCount=6, got %v", workloadPatch.GetConfig().GetWorkloadsConfig().GetWorker().GetMaxCount())
		}
	})
}

func TestValidateUpdatableFields(t *testing.T) {
	t.Run("returns nil for valid updates on mutable fields", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Labels: map[string]string{"env": "prod"},
				Config: &krm.EnvironmentConfig{
					EnvironmentSize: direct.LazyPtr("ENVIRONMENT_SIZE_SMALL"),
				},
			},
		}
		actual := &composerpb.Environment{
			Labels: map[string]string{"env": "dev"},
			Config: &composerpb.EnvironmentConfig{
				EnvironmentSize: composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_MEDIUM,
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		err := validateUpdatableFields(desiredPb, actual)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

	t.Run("fails on nodeConfig.machineType modification", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					NodeConfig: &krm.NodeConfig{
						MachineType: direct.LazyPtr("n1-standard-4"),
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				NodeConfig: &composerpb.NodeConfig{
					MachineType: "n1-standard-1",
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		err := validateUpdatableFields(desiredPb, actual)
		if err == nil {
			t.Fatalf("expected error for unsupported machineType modification, got nil")
		}
		expectedMsg := `updating field(s) [config.node_config.machine_type] is not supported`
		if err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
		}
	})

	t.Run("fails on storageConfig.bucketRef modification", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				StorageConfig: &krm.StorageConfig{
					BucketRef: &storagev1beta1.StorageBucketRef{
						External: "projects/p1/buckets/new-bucket",
					},
				},
			},
		}
		actual := &composerpb.Environment{
			StorageConfig: &composerpb.StorageConfig{
				Bucket: "projects/p1/buckets/old-bucket",
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		err := validateUpdatableFields(desiredPb, actual)
		if err == nil {
			t.Fatalf("expected error for unsupported bucket modification, got nil")
		}
		expectedMsg := `updating field(s) [storage_config.bucket] is not supported`
		if err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
		}
	})

	t.Run("fails on encryptionConfig.kmsKeyRef modification", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					EncryptionConfig: &krm.EncryptionConfig{
						KMSKeyRef: &refs.KMSCryptoKeyRef{
							External: "projects/p1/locations/l1/keyRings/r1/cryptoKeys/k2",
						},
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				EncryptionConfig: &composerpb.EncryptionConfig{
					KmsKeyName: "projects/p1/locations/l1/keyRings/r1/cryptoKeys/k1",
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		err := validateUpdatableFields(desiredPb, actual)
		if err == nil {
			t.Fatalf("expected error for unsupported KMS key modification, got nil")
		}
		expectedMsg := `updating field(s) [config.encryption_config.kms_key_name] is not supported`
		if err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
		}
	})

	t.Run("fails on databaseConfig.zone modification", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				Config: &krm.EnvironmentConfig{
					DatabaseConfig: &krm.DatabaseConfig{
						Zone: direct.LazyPtr("us-central1-b"),
					},
				},
			},
		}
		actual := &composerpb.Environment{
			Config: &composerpb.EnvironmentConfig{
				DatabaseConfig: &composerpb.DatabaseConfig{
					Zone: "us-central1-a",
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		err := validateUpdatableFields(desiredPb, actual)
		if err == nil {
			t.Fatalf("expected error for unsupported databaseConfig.zone modification, got nil")
		}
		expectedMsg := `updating field(s) [config.database_config.zone] is not supported`
		if err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
		}
	})

	t.Run("fails on multiple unsupported field modifications", func(t *testing.T) {
		desired := &krm.ComposerEnvironment{
			Spec: krm.ComposerEnvironmentSpec{
				StorageConfig: &krm.StorageConfig{
					BucketRef: &storagev1beta1.StorageBucketRef{
						External: "projects/p1/buckets/new-bucket",
					},
				},
				Config: &krm.EnvironmentConfig{
					DatabaseConfig: &krm.DatabaseConfig{
						Zone: direct.LazyPtr("us-central1-b"),
					},
					NodeConfig: &krm.NodeConfig{
						MachineType: direct.LazyPtr("n1-standard-4"),
					},
				},
			},
		}
		actual := &composerpb.Environment{
			StorageConfig: &composerpb.StorageConfig{
				Bucket: "projects/p1/buckets/old-bucket",
			},
			Config: &composerpb.EnvironmentConfig{
				DatabaseConfig: &composerpb.DatabaseConfig{
					Zone: "us-central1-a",
				},
				NodeConfig: &composerpb.NodeConfig{
					MachineType: "n1-standard-1",
				},
			},
		}
		mapCtx := &direct.MapContext{}
		desiredPb := ComposerEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("unexpected error converting desired spec: %v", mapCtx.Err())
		}

		err := validateUpdatableFields(desiredPb, actual)
		if err == nil {
			t.Fatalf("expected error for multiple unsupported field modifications, got nil")
		}
		expectedMsg := `updating field(s) [config.database_config.zone config.node_config.machine_type storage_config.bucket] is not supported`
		if err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %q", expectedMsg, err.Error())
		}
	})
}

func TestFieldUpdatersConsistency(t *testing.T) {
	seenMasks := make(map[string]bool)
	for _, u := range fieldUpdaters {
		if u.mask == "" {
			t.Errorf("fieldUpdater with empty mask found")
		}
		if seenMasks[u.mask] {
			t.Errorf("duplicate mask %q in fieldUpdaters", u.mask)
		}
		seenMasks[u.mask] = true

		if !isValidUpdatePrefix(u.mask) {
			t.Errorf("isValidUpdatePrefix(%q) returned false for registered fieldUpdater", u.mask)
		}
	}
}

// TestProtoUpdatableFieldCoverageViaGoDoc verifies that our fieldUpdaters registry covers
// all updatable field paths documented by Google in the Cloud Composer Go client library.
//
// Instead of reading raw .proto files from disk, this test uses Go's standard library
// (go/build, go/parser, and go/doc) to parse the Go package documentation of
// cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb.
func TestProtoUpdatableFieldCoverageViaGoDoc(t *testing.T) {
	// Step 1: Locate the compiled Go package on the local system using standard Go build tooling.
	pkg, err := build.Import("cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb", "", 0)
	if err != nil {
		t.Fatalf("failed to import servicepb package: %v", err)
	}

	// Step 2: Parse all Go source files in the package directory into an Abstract Syntax Tree (AST)
	// while preserving code comments (parser.ParseComments).
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, pkg.Dir, nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse servicepb directory: %v", err)
	}

	astPkg, ok := pkgs[pkg.Name]
	if !ok {
		t.Fatalf("package %q not found in parsed packages", pkg.Name)
	}

	// Step 3: Extract structured Go package documentation from the AST.
	docPkg := doc.New(astPkg, pkg.ImportPath, doc.AllDecls)

	// Step 4: Find the 'UpdateEnvironmentRequest' struct and retrieve docstrings from the 'UpdateMask' field.
	// Google's protobuf compiler places the list of supported update masks in the docstring of UpdateMask.
	var reqDoc string
	for _, typ := range docPkg.Types {
		if typ.Name == "UpdateEnvironmentRequest" {
			reqDoc = typ.Doc
			if ts, ok := typ.Decl.Specs[0].(*ast.TypeSpec); ok {
				if st, ok := ts.Type.(*ast.StructType); ok {
					for _, f := range st.Fields.List {
						for _, name := range f.Names {
							if name.Name == "UpdateMask" && f.Doc != nil {
								reqDoc += "\n" + f.Doc.Text()
							}
						}
					}
				}
			}
			break
		}
	}
	if reqDoc == "" {
		t.Fatalf("doc comments for UpdateEnvironmentRequest not found in package %s", pkg.ImportPath)
	}

	// Step 5: Extract all documented update paths using regular expressions.
	// Google formats bullet points like: `* `config.softwareConfig.imageVersion`` or `* `labels``
	re := regexp.MustCompile(`\* ` + "`" + `([a-zA-Z0-9\._]+)` + "`")
	matches := re.FindAllStringSubmatch(reqDoc, -1)
	if len(matches) == 0 {
		// Fallback for plain-text bullet points without backticks
		re = regexp.MustCompile(`\*\s+([a-zA-Z0-9\._]+)`)
		matches = re.FindAllStringSubmatch(reqDoc, -1)
	}
	if len(matches) == 0 {
		t.Fatalf("no updatable fields extracted from UpdateEnvironmentRequest doc: %s", reqDoc)
	}

	// Step 6: Convert each extracted camelCase path to snake_case and collect unique mutable paths from Go doc.
	mutablePathsInGoDocSet := make(map[string]bool)
	for _, m := range matches {
		rawPath := strings.TrimSuffix(m[1], ".")
		snakePath := camelToSnake(rawPath)
		mutablePathsInGoDocSet[snakePath] = true
	}

	var mutablePathsInGoDoc []string
	for p := range mutablePathsInGoDocSet {
		mutablePathsInGoDoc = append(mutablePathsInGoDoc, p)
	}
	slices.Sort(mutablePathsInGoDoc)

	// Step 8: Subset validation.
	// Note: The docstring in cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb reflects
	// the legacy Composer 1 baseline and was not updated by Google to list newer Composer 2/3 features
	// (e.g. workloads_config, environment_size, maintenance_window, data_retention_config, etc.) that
	// are supported by the live GCP v1 backend.
	// Therefore, the documented paths in Go doc must at least be a strict subset of our registered fieldUpdaters.
	for _, docPath := range mutablePathsInGoDoc {
		if !isValidUpdatePrefix(docPath) {
			t.Errorf("Go package documents updatable path %q, but it is missing from fieldUpdaters / isValidUpdatePrefix", docPath)
		}
	}
}

// camelToSnake converts proto camelCase field names to snake_case.
// For example: "config.softwareConfig.imageVersion" -> "config.software_config.image_version"
func camelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && s[i-1] != '.' && !unicode.IsUpper(rune(s[i-1])) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func TestMaximumComposerEnvironment(t *testing.T) {
	createSpec := &krm.ComposerEnvironment{
		Spec: krm.ComposerEnvironmentSpec{
			Labels: map[string]string{
				"env":        "staging",
				"department": "platform",
			},
			StorageConfig: &krm.StorageConfig{
				BucketRef: &storagev1beta1.StorageBucketRef{
					External: "gs://bucket-12345",
				},
			},
			Config: &krm.EnvironmentConfig{
				EnvironmentSize: direct.LazyPtr("ENVIRONMENT_SIZE_MEDIUM"),
				ResilienceMode:  direct.LazyPtr("HIGH_RESILIENCE"),
				EncryptionConfig: &krm.EncryptionConfig{
					KMSKeyRef: &refs.KMSCryptoKeyRef{
						External: "projects/p1/locations/us-central1/keyRings/r1/cryptoKeys/k1",
					},
				},
				MaintenanceWindow: &krm.MaintenanceWindow{
					StartTime:  direct.LazyPtr("2026-01-01T04:00:00Z"),
					EndTime:    direct.LazyPtr("2026-01-01T08:00:00Z"),
					Recurrence: direct.LazyPtr("FREQ=WEEKLY;BYDAY=TU,WE,TH"),
				},
				SoftwareConfig: &krm.SoftwareConfig{
					ImageVersion: direct.LazyPtr("composer-3-airflow-3.1.7-build.5"),
					AirflowConfigOverrides: map[string]string{
						"core-dags_are_paused_at_creation": "true",
						"webserver-navbar_color":           "#112233",
					},
					PypiPackages: map[string]string{
						"scipy": "==1.11.4",
					},
					EnvVariables: map[string]string{
						"TEST_ENV": "value1",
					},
					SchedulerCount:       direct.LazyPtr(int32(2)),
					WebServerPluginsMode: direct.LazyPtr("PLUGINS_DISABLED"),
				},
				NodeConfig: &krm.NodeConfig{
					NetworkRef: &computerefs.ComputeNetworkRef{
						External: "projects/p1/global/networks/network-12345",
					},
					SubnetworkRef: &computev1beta1.ComputeSubnetworkRef{
						External: "projects/p1/regions/us-central1/subnetworks/subnet-12345",
					},
					ServiceAccountRef: &refs.IAMServiceAccountRef{
						External: "sa-12345@p1.iam.gserviceaccount.com",
					},
					MachineType: direct.LazyPtr("n1-standard-2"),
					DiskSizeGB:  direct.LazyPtr(int32(120)),
					ComposerNetworkAttachmentRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "projects/p1/regions/us-central1/networkAttachments/attachment-12345",
					},
					ComposerInternalIPv4CIDRBlock: direct.LazyPtr("100.64.0.0/20"),
					Tags:                         []string{"composer-node", "data-platform"},
					EnableIPMasqAgent:            direct.LazyPtr(false),
					IPAllocationPolicy: &krm.IPAllocationPolicy{
						UseIPAliases:               direct.LazyPtr(true),
						ClusterSecondaryRangeName:  direct.LazyPtr("pods"),
						ServicesSecondaryRangeName: direct.LazyPtr("services"),
					},
				},
				PrivateEnvironmentConfig: &krm.PrivateEnvironmentConfig{
					EnablePrivateEnvironment: direct.LazyPtr(true),
					EnablePrivateBuildsOnly:  direct.LazyPtr(true),
					CloudComposerConnectionSubnetworkRef: &computev1beta1.ComputeSubnetworkRef{
						External: "projects/p1/regions/us-central1/subnetworks/subnetwork-private-12345",
					},
					EnablePrivatelyUsedPublicIPs: direct.LazyPtr(false),
					CloudComposerNetworkIPv4CIDRBlock: direct.LazyPtr("172.31.248.0/24"),
					CloudSQLIPv4CIDRBlock:            direct.LazyPtr("10.0.0.0/24"),
					WebServerIPv4CIDRBlock:           direct.LazyPtr("10.0.1.0/24"),
					PrivateClusterConfig: &krm.PrivateClusterConfig{
						EnablePrivateEndpoint: direct.LazyPtr(true),
						MasterIPV4CIDRBlock:   direct.LazyPtr("172.16.0.0/28"),
					},
					NetworkingConfig: &krm.NetworkingConfig{
						ConnectionType: direct.LazyPtr("PRIVATE_SERVICE_CONNECT"),
					},
				},
				WebServerNetworkAccessControl: &krm.WebServerNetworkAccessControl{
					AllowedIPRanges: []krm.WebServerNetworkAccessControl_AllowedIPRange{
						{
							Value:       direct.LazyPtr("192.168.1.0/24"),
							Description: direct.LazyPtr("Corporate VPN"),
						},
						{
							Value:       direct.LazyPtr("10.0.0.0/8"),
							Description: direct.LazyPtr("Internal VPC"),
						},
					},
				},
				DatabaseConfig: &krm.DatabaseConfig{
					MachineType: direct.LazyPtr("db-custom-4-15360"),
					Zone:        direct.LazyPtr("us-central1-a"),
				},
				WorkloadsConfig: &krm.WorkloadsConfig{
					Scheduler: &krm.WorkloadsConfig_SchedulerResource{
						CPU:       direct.LazyPtr("1.0"),
						MemoryGB:  direct.LazyPtr("4"),
						StorageGB: direct.LazyPtr("2"),
						Count:     direct.LazyPtr(int32(2)),
					},
					WebServer: &krm.WorkloadsConfig_WebServerResource{
						CPU:       direct.LazyPtr("1.0"),
						MemoryGB:  direct.LazyPtr("4"),
						StorageGB: direct.LazyPtr("2"),
					},
					Worker: &krm.WorkloadsConfig_WorkerResource{
						CPU:       direct.LazyPtr("1.0"),
						MemoryGB:  direct.LazyPtr("4"),
						StorageGB: direct.LazyPtr("2"),
						MinCount:  direct.LazyPtr(int32(2)),
						MaxCount:  direct.LazyPtr(int32(6)),
					},
					Triggerer: &krm.WorkloadsConfig_TriggererResource{
						CPU:      direct.LazyPtr("1.0"),
						MemoryGB: direct.LazyPtr("2"),
						Count:    direct.LazyPtr(int32(2)),
					},
					DagProcessor: &krm.WorkloadsConfig_DagProcessorResource{
						CPU:       direct.LazyPtr("1.0"),
						MemoryGB:  direct.LazyPtr("4"),
						StorageGB: direct.LazyPtr("2"),
						Count:     direct.LazyPtr(int32(2)),
					},
				},
				MasterAuthorizedNetworksConfig: &krm.MasterAuthorizedNetworksConfig{
					Enabled: direct.LazyPtr(true),
					CIDRBlocks: []krm.MasterAuthorizedNetworksConfig_CIDRBlock{
						{
							CIDRBlock:   direct.LazyPtr("192.168.1.0/24"),
							DisplayName: direct.LazyPtr("Admin Subnet"),
						},
					},
				},
				RecoveryConfig: &krm.RecoveryConfig{
					ScheduledSnapshotsConfig: &krm.ScheduledSnapshotsConfig{
						Enabled:                  direct.LazyPtr(true),
						SnapshotCreationSchedule: direct.LazyPtr("0 4 * * *"),
						SnapshotLocation:         direct.LazyPtr("gs://bucket-12345/snapshots"),
						TimeZone:                 direct.LazyPtr("America/New_York"),
					},
				},
				DataRetentionConfig: &krm.DataRetentionConfig{
					AirflowMetadataRetentionConfig: &krm.AirflowMetadataRetentionPolicyConfig{
						RetentionMode: direct.LazyPtr("RETENTION_MODE_ENABLED"),
						RetentionDays: direct.LazyPtr(int32(45)),
					},
				},
			},
		},
	}

	mapCtx := &direct.MapContext{}
	createPb := ComposerEnvironmentSpec_ToProto(mapCtx, &createSpec.Spec)
	if mapCtx.Err() != nil {
		t.Fatalf("failed to convert maximum createSpec to proto: %v", mapCtx.Err())
	}

	// Verify non-default proto fields are populated
	if createPb.GetStorageConfig().GetBucket() != "bucket-12345" {
		t.Errorf("expected bucket-12345, got %v", createPb.GetStorageConfig().GetBucket())
	}
	if createPb.GetConfig().GetEnvironmentSize() != composerpb.EnvironmentConfig_ENVIRONMENT_SIZE_MEDIUM {
		t.Errorf("expected size medium, got %v", createPb.GetConfig().GetEnvironmentSize())
	}
	if createPb.GetConfig().GetWorkloadsConfig().GetWorker().GetMaxCount() != 6 {
		t.Errorf("expected worker max count 6, got %v", createPb.GetConfig().GetWorkloadsConfig().GetWorker().GetMaxCount())
	}
	if createPb.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig().GetSnapshotCreationSchedule() != "0 4 * * *" {
		t.Errorf("expected schedule '0 4 * * *', got %v", createPb.GetConfig().GetRecoveryConfig().GetScheduledSnapshotsConfig().GetSnapshotCreationSchedule())
	}
	if createPb.GetConfig().GetNodeConfig().GetDiskSizeGb() != 120 {
		t.Errorf("expected disk size 120, got %v", createPb.GetConfig().GetNodeConfig().GetDiskSizeGb())
	}
	if createPb.GetConfig().GetNodeConfig().GetComposerInternalIpv4CidrBlock() != "100.64.0.0/20" {
		t.Errorf("expected CIDR '100.64.0.0/20', got %v", createPb.GetConfig().GetNodeConfig().GetComposerInternalIpv4CidrBlock())
	}
	if createPb.GetConfig().GetDataRetentionConfig().GetAirflowMetadataRetentionConfig().GetRetentionDays() != 45 {
		t.Errorf("expected retention days 45, got %v", createPb.GetConfig().GetDataRetentionConfig().GetAirflowMetadataRetentionConfig().GetRetentionDays())
	}

	// Verify roundtrip Spec_FromProto
	roundtripSpec := ComposerEnvironmentSpec_FromProto(mapCtx, createPb)
	if mapCtx.Err() != nil {
		t.Fatalf("failed to convert proto back to Spec: %v", mapCtx.Err())
	}
	if *roundtripSpec.Config.EnvironmentSize != "ENVIRONMENT_SIZE_MEDIUM" {
		t.Errorf("expected roundtrip environmentSize ENVIRONMENT_SIZE_MEDIUM, got %v", *roundtripSpec.Config.EnvironmentSize)
	}

	// Construct updated spec with distinct non-default values
	updateSpec := &krm.ComposerEnvironment{
		Spec: krm.ComposerEnvironmentSpec{
			Labels: map[string]string{
				"env":        "production",
				"department": "platform",
				"updated":    "true",
			},
			StorageConfig: createSpec.Spec.StorageConfig,
			Config: &krm.EnvironmentConfig{
				EnvironmentSize:  direct.LazyPtr("ENVIRONMENT_SIZE_LARGE"),
				ResilienceMode:   createSpec.Spec.Config.ResilienceMode,
				EncryptionConfig: createSpec.Spec.Config.EncryptionConfig,
				MaintenanceWindow: &krm.MaintenanceWindow{
					StartTime:  direct.LazyPtr("2026-01-01T08:00:00Z"),
					EndTime:    direct.LazyPtr("2026-01-01T12:00:00Z"),
					Recurrence: direct.LazyPtr("FREQ=WEEKLY;BYDAY=MO,TU,WE"),
				},
				SoftwareConfig: &krm.SoftwareConfig{
					ImageVersion: createSpec.Spec.Config.SoftwareConfig.ImageVersion,
					AirflowConfigOverrides: map[string]string{
						"core-dags_are_paused_at_creation": "true",
						"webserver-navbar_color":           "#445566",
					},
					PypiPackages: map[string]string{
						"scipy":  "==1.11.4",
						"numpy":  "==1.26.4",
						"pandas": "==2.2.0",
					},
					EnvVariables: map[string]string{
						"TEST_ENV":  "value_updated",
						"EXTRA_VAR": "true",
					},
					SchedulerCount:       direct.LazyPtr(int32(2)),
					WebServerPluginsMode: direct.LazyPtr("PLUGINS_DISABLED"),
				},
				NodeConfig:               createSpec.Spec.Config.NodeConfig,
				PrivateEnvironmentConfig: createSpec.Spec.Config.PrivateEnvironmentConfig,
				WebServerNetworkAccessControl: &krm.WebServerNetworkAccessControl{
					AllowedIPRanges: []krm.WebServerNetworkAccessControl_AllowedIPRange{
						{
							Value:       direct.LazyPtr("192.168.1.0/24"),
							Description: direct.LazyPtr("Corporate VPN"),
						},
						{
							Value:       direct.LazyPtr("10.0.0.0/8"),
							Description: direct.LazyPtr("Internal VPC"),
						},
						{
							Value:       direct.LazyPtr("172.16.0.0/12"),
							Description: direct.LazyPtr("Branch Office"),
						},
					},
				},
				DatabaseConfig: &krm.DatabaseConfig{
					MachineType: direct.LazyPtr("db-custom-8-30720"),
					Zone:        direct.LazyPtr("us-central1-a"),
				},
				WorkloadsConfig: &krm.WorkloadsConfig{
					Scheduler: &krm.WorkloadsConfig_SchedulerResource{
						CPU:       direct.LazyPtr("1.5"),
						MemoryGB:  direct.LazyPtr("6"),
						StorageGB: direct.LazyPtr("3"),
						Count:     direct.LazyPtr(int32(3)),
					},
					WebServer: &krm.WorkloadsConfig_WebServerResource{
						CPU:       direct.LazyPtr("1.5"),
						MemoryGB:  direct.LazyPtr("6"),
						StorageGB: direct.LazyPtr("3"),
					},
					Worker: &krm.WorkloadsConfig_WorkerResource{
						CPU:       direct.LazyPtr("1.5"),
						MemoryGB:  direct.LazyPtr("6"),
						StorageGB: direct.LazyPtr("3"),
						MinCount:  direct.LazyPtr(int32(3)),
						MaxCount:  direct.LazyPtr(int32(10)),
					},
					Triggerer: &krm.WorkloadsConfig_TriggererResource{
						CPU:      direct.LazyPtr("1.5"),
						MemoryGB: direct.LazyPtr("3"),
						Count:    direct.LazyPtr(int32(3)),
					},
					DagProcessor: &krm.WorkloadsConfig_DagProcessorResource{
						CPU:       direct.LazyPtr("1.5"),
						MemoryGB:  direct.LazyPtr("6"),
						StorageGB: direct.LazyPtr("3"),
						Count:     direct.LazyPtr(int32(3)),
					},
				},
				MasterAuthorizedNetworksConfig: &krm.MasterAuthorizedNetworksConfig{
					Enabled: direct.LazyPtr(true),
					CIDRBlocks: []krm.MasterAuthorizedNetworksConfig_CIDRBlock{
						{
							CIDRBlock:   direct.LazyPtr("192.168.1.0/24"),
							DisplayName: direct.LazyPtr("Admin Subnet"),
						},
						{
							CIDRBlock:   direct.LazyPtr("10.0.0.0/16"),
							DisplayName: direct.LazyPtr("Dev Subnet"),
						},
					},
				},
				RecoveryConfig: &krm.RecoveryConfig{
					ScheduledSnapshotsConfig: &krm.ScheduledSnapshotsConfig{
						Enabled:                  direct.LazyPtr(true),
						SnapshotCreationSchedule: direct.LazyPtr("0 8 * * *"),
						SnapshotLocation:         direct.LazyPtr("gs://bucket-12345/snapshots"),
						TimeZone:                 direct.LazyPtr("America/Chicago"),
					},
				},
				DataRetentionConfig: &krm.DataRetentionConfig{
					AirflowMetadataRetentionConfig: &krm.AirflowMetadataRetentionPolicyConfig{
						RetentionMode: direct.LazyPtr("RETENTION_MODE_ENABLED"),
						RetentionDays: direct.LazyPtr(int32(90)),
					},
				},
			},
		},
	}

	updatePb := ComposerEnvironmentSpec_ToProto(mapCtx, &updateSpec.Spec)
	if mapCtx.Err() != nil {
		t.Fatalf("failed to convert updateSpec to proto: %v", mapCtx.Err())
	}

	// Validate updatability of maximum spec transition
	if err := validateUpdatableFields(updatePb, createPb); err != nil {
		t.Fatalf("unexpected validation error during maximum update transition: %v", err)
	}

	// Verify buildPatches generates expected individual patches
	patches := buildPatches(updateSpec, updatePb, createPb)
	expectedMasks := []string{
		"labels",
		"config.environment_size",
		"config.maintenance_window",
		"config.software_config.airflow_config_overrides",
		"config.software_config.pypi_packages",
		"config.software_config.env_variables",
		"config.web_server_network_access_control",
		"config.database_config.machine_type",
		"config.workloads_config",
		"config.master_authorized_networks_config",
		"config.recovery_config.scheduled_snapshots_config",
		"config.data_retention_config",
	}

	for _, mask := range expectedMasks {
		if _, ok := patches[mask]; !ok {
			t.Errorf("expected update mask %q to be built, but was missing", mask)
		}
	}
}

