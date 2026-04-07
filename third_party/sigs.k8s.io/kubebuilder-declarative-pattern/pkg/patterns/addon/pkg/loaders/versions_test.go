/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loaders

import (
	"context"
	"testing"
)

func TestCompareVersions(t *testing.T) {
	ctx := context.TODO()

	grid := []struct {
		Bigger  Version
		Smaller Version
	}{
		// Basic ordering applies
		{
			Bigger:  Version{Package: "", Version: "2.0.0"},
			Smaller: Version{Package: "", Version: "1.0.0"},
		},
		// Parsing is tolerant
		{
			Bigger:  Version{Package: "", Version: "1"},
			Smaller: Version{Package: "", Version: "0.1"},
		},
		// Named packages have higher priority
		{
			Bigger:  Version{Package: "foo", Version: "v1"},
			Smaller: Version{Package: "", Version: "v1"},
		},
		// Valid versions are higher priority than invalid versions
		{
			Bigger:  Version{Package: "", Version: "v1"},
			Smaller: Version{Package: "", Version: "corrupt2"},
		},
	}

	for _, g := range grid {
		{
			l := &g.Bigger
			r := &g.Smaller
			v := l.Compare(ctx, r)

			if v <= 0 {
				t.Errorf("comparison failed: %+v vs %+v was %d", l, r, v)
			}
		}

		{
			l := &g.Smaller
			r := &g.Bigger
			v := l.Compare(ctx, r)

			if v >= 0 {
				t.Errorf("comparison failed: %+v vs %+v was %d", l, r, v)
			}
		}

		{
			l := &g.Bigger
			r := &g.Bigger
			v := l.Compare(ctx, r)

			if v != 0 {
				t.Errorf("comparison failed: %+v vs %+v was %d", l, r, v)
			}
		}

		{
			l := &g.Smaller
			r := &g.Smaller
			v := l.Compare(ctx, r)

			if v != 0 {
				t.Errorf("comparison failed: %+v vs %+v was %d", l, r, v)
			}
		}
	}
}
