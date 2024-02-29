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

package logging_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/google/go-cmp/cmp"
)

type ExampleError struct{}

func (e *ExampleError) Error() string {
	return "ExampleError"
}

// Tests that buildLogger (used by SetupLogger) sets up log entries to be in JSON and
// with fields compatible with GCP Cloud Logging.
func TestBuildLoggerFormat(t *testing.T) {
	var buffer bytes.Buffer
	logger := logging.BuildLogger(&buffer)
	inputInfo := "testLoggingFormat"
	logger.Info(inputInfo)
	logger.Error(&ExampleError{}, "test")
	wantList := [...]map[string]interface{}{
		{
			"msg":      inputInfo,
			"severity": "info",
		},
		{
			"error":    "ExampleError",
			"msg":      "test",
			"severity": "error",
		},
	}
	for _, want := range wantList {
		severity := want["severity"].(string)
		t.Run(severity, func(t *testing.T) {
			loggedLine, err := buffer.ReadString('\n')
			if err != nil {
				t.Fatalf("ReadString('\n') failed: %s", err)
			}
			var got map[string]interface{}
			if err := json.Unmarshal([]byte(loggedLine), &got); err != nil {
				t.Fatal(err)
			}
			if _, ok := got["timestamp"]; !ok {
				t.Fatalf("the log message %v doesn't contain `timestamp`", got)
			}
			delete(got, "timestamp")
			if diff := cmp.Diff(want, got); diff != "" {
				t.Fatalf("logger %s returned unexpected diff (-want +got): \n%s", severity, diff)
			}
		})
	}
}
