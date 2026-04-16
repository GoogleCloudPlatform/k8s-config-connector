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

package gkehub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	gkehubapi "google.golang.org/api/gkehub/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func waitForLRO(ctx context.Context, client *gkehubapi.ProjectsLocationsOperationsService, op *gkehubapi.Operation) error {
	if op.Done {
		if op.Error != nil {
			return fmt.Errorf("operation %q completed with error: %v", op.Name, op.Error)
		}
		return nil
	}

	retryPeriod := 5 * time.Second
	timeoutDuration := 20 * time.Minute
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := client.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q failed: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		if time.Now().After(timeoutAt) {
			return fmt.Errorf("operation timed out waiting for LRO after %s", timeoutDuration.String())
		}
		time.Sleep(retryPeriod)
		if retryPeriod < 30*time.Second {
			retryPeriod = retryPeriod * 2
		}
	}
}

func Convert_v1_pb_to_api(in proto.Message, out interface{}) error {
	b, err := protojson.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func Convert_v1_api_to_pb(in interface{}, out proto.Message) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(b, out)
}

func StringMapEqual(a, b map[string]string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || v != bv {
			return false
		}
	}
	return true
}
