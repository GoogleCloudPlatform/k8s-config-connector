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

package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

const (
	name    = "kompanion"
	unknown = "unknown"
)

func GetVersion() string {
	version := "" // todo acpana in the future factor this out for major releases
	vcsrevision := unknown
	vcstimestamp := unknown
	vcsdirty := ""

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, v := range info.Settings {
			switch v.Key {
			case "vcs.revision":
				vcsrevision = v.Value
			case "vcs.modified":
				if v.Value == "true" {
					vcsdirty = "-dirty"
				}
			case "vcs.time":
				vcstimestamp = v.Value
			}
		}
	}

	version = fmt.Sprintf("devel (%s)", vcsrevision)
	return fmt.Sprintf("%s/%s (%s/%s) %s%s/%s", name, version, runtime.GOOS, runtime.GOARCH, vcsrevision, vcsdirty, vcstimestamp)
}
