// Copyright 2024 Google LLC
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

package controllers

import (
	"testing"
)

func TestImageTransform(t *testing.T) {
	grid := []struct {
		ImagePrefix string
		Image       string
		Want        string
	}{
		{
			ImagePrefix: "foo/bar",
			Image:       "gcr.io/gke-release/cnrm/deletiondefender:1.124.0-rc.1",
			Want:        "foo/bar/deletiondefender:1.124.0-rc.1",
		},
		{
			ImagePrefix: "foo/bar",
			Image:       "gcr.io/gke-release/cnrm/deletiondefender:latest",
			Want:        "foo/bar/deletiondefender:latest",
		},
		{
			ImagePrefix: "foo/bar",
			Image:       "gcr.io/gke-release/cnrm/deletiondefender",
			Want:        "foo/bar/deletiondefender",
		},
		{
			ImagePrefix: "foo/bar/",
			Image:       "gcr.io/gke-release/cnrm/deletiondefender",
			Want:        "foo/bar/deletiondefender",
		},
		{
			ImagePrefix: "",
			Image:       "gcr.io/gke-release/cnrm/deletiondefender",
			Want:        "deletiondefender",
		},
	}

	for _, g := range grid {
		x := NewImageTransform(g.ImagePrefix)
		got := x.remapImage(g.Image)
		if got != g.Want {
			t.Errorf("unexpected remapping with ImagePrefix=%q from %q; got %q, want %q", g.ImagePrefix, g.Image, got, g.Want)
		}
	}
}
