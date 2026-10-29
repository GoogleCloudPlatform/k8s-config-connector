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

package connectors

import (
	"context"
	"testing"
)

func TestAdapterForURL(t *testing.T) {
	ctx := context.Background()
	m := &model{}

	// Non-matching URLs must return (nil, nil) so other direct models in the registry can match.
	unrelatedURLs := []string{
		"//bigquery.googleapis.com/projects/p/datasets/d",
		"projects/p/locations/l/keyRings/k",
		"//pubsub.googleapis.com/projects/p/topics/t",
		"//secretmanager.googleapis.com/projects/p/secrets/s",
		"",
	}
	for _, u := range unrelatedURLs {
		adapter, err := m.AdapterForURL(ctx, u)
		if err != nil {
			t.Errorf("AdapterForURL(%q) returned unexpected error: %v, expected (nil, nil)", u, err)
		}
		if adapter != nil {
			t.Errorf("AdapterForURL(%q) returned unexpected adapter: %v, expected nil", u, adapter)
		}
	}
}
