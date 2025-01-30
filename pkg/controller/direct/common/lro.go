// Copyright 2025 Google LLC
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

package common

import (
	"context"
	"fmt"
	"time"
)

func WaitForDoneOrTimeout(ctx context.Context, pollInterval time.Duration, doneFunc func() (bool, error)) error {
	for {
		// Check for timeout
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout exceeded")
		default:
		}

		// Check for condition
		result, err := doneFunc()
		if err != nil {
			return err
		}
		if result {
			return nil
		}

		// Wait for pollInterval
		time.Sleep(pollInterval)
	}
}
