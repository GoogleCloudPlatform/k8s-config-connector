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

package resourcefixture

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testconstants "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/constants"
)

type ShouldRunFunc func(fixture ResourceFixture) bool
type TestCaseFunc func(ctx context.Context, t *testing.T, fixture ResourceFixture)

func RunTests(ctx context.Context, t *testing.T, shouldRun ShouldRunFunc, testCaseFunc TestCaseFunc) {
	testCases := Load(t)
	for _, tc := range testCases {
		if !shouldRun(tc) {
			continue
		}
		runTestCase(ctx, t, tc, testCaseFunc)
	}
}

func RunSpecificTests(ctx context.Context, t *testing.T, fixtures []ResourceFixture, testCaseFunc TestCaseFunc) {
	for _, f := range fixtures {
		runTestCase(ctx, t, f, testCaseFunc)
	}
}

func runTestCase(ctx context.Context, t *testing.T, fixture ResourceFixture, testCaseFunc TestCaseFunc) {
	testName := FormatTestName(fixture)
	if test.StringMatchesRegexList(t, testconstants.TestNameRegexesToSkip, testName) {
		return
	}
	t.Run(FormatTestName(fixture), func(t *testing.T) {
		t.Parallel()
		ctx = test.WithContext(ctx, t)
		testCaseFunc(ctx, t, fixture)
		// note, this function, runTestCase(...) almost always returns before testCaseFunc(...) returns
	})
}

func FormatTestName(tc ResourceFixture) string {
	return fmt.Sprintf("%v-%v", string(tc.Type), tc.Name)
}
