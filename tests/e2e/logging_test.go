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

package e2e

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/log"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func TestStructuredReportingLogging(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	// Capture logs
	var logBuffer bytes.Buffer

	// Create a zap logger that writes to the buffer
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(&logBuffer),
		zap.DebugLevel,
	)
	zapLogger := zap.New(core)
	logger := zapr.NewLogger(zapLogger)

	// Inject logger into context
	ctx := context.Background()
	ctx = log.IntoContext(ctx, logger)
	ctx = structuredreporting.ContextWithListener(ctx, &structuredreporting.DebugLogListener{})

	// Run the test in a subtest to ensure cleanup
	t.Run("sqlinstance-logging", func(t *testing.T) {
		// Setup Harness
		h := create.NewHarness(ctx, t)

		uniqueID := testvariable.NewUniqueID()
		project := h.Project

		yamlCreate := `
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: sqlinstance-${uniqueId}
spec:
  databaseVersion: POSTGRES_15
  region: us-central1
  settings:
    tier: db-custom-1-3840
`
		yamlUpdate := `
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: sqlinstance-${uniqueId}
spec:
  databaseVersion: POSTGRES_15
  region: us-central1
  settings:
    tier: db-custom-2-7680
`

		// Replace variables
		yamlCreate = strings.ReplaceAll(yamlCreate, "${uniqueId}", uniqueID)
		yamlUpdate = strings.ReplaceAll(yamlUpdate, "${uniqueId}", uniqueID)

		// Run Create/Update
		createObj := bytesToUnstructured(t, []byte(yamlCreate), uniqueID, project)
		updateObj := bytesToUnstructured(t, []byte(yamlUpdate), uniqueID, project)

		create.SetupNamespacesAndApplyDefaults(h, []*unstructured.Unstructured{createObj}, project)
		create.SetupNamespacesAndApplyDefaults(h, []*unstructured.Unstructured{updateObj}, project)

		opt := create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{createObj},
			Updates:          []*unstructured.Unstructured{updateObj},
			CleanupResources: true,
			PrimaryResource:  createObj,
		}

		create.RunCreateDeleteTest(h, opt)
	})

	// Check logs
	logOutput := logBuffer.String()
	if !strings.Contains(logOutput, "structuredreporting OnDiff") {
		t.Errorf("Expected 'structuredreporting OnDiff' in logs, but not found.")
	}

	if !strings.Contains(logOutput, ".settings.tier") {
		t.Errorf("Expected '.settings.tier' in logs, but not found.")
	}

	if t.Failed() {
		t.Logf("Logs:\n%s", logOutput)
	}
}
