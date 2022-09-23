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

package clientconfig

import (
	"time"
)

// kindToTimeout lists the kinds for which we configure a custom DCL timeout
// (i.e. the time taken by DCL operations like Get/List/Apply/Delete before
// timing out). Note that DCL only supports kind-specific timeouts (i.e. a kind
// uses the same timeout for all DCL operations).
var kindToTimeout = map[string]time.Duration{
	// ApigeeOrganization operations could take a long time to complete. Setting
	// it to 80 minutes to match the DCL configuration.
	"ApigeeOrganization": 80 * time.Minute,
	// DataprocCluster creations can take 35+ minutes because the long-running
	// operation kicked off by DCL on GCP can take that long to time out and
	// report errors if there are any issues (e.g. b/187470070#comment6). Let
	// us configure a long enough DCL timeout to give the long-running
	// operation a chance to finish so that its error can be surfaced back to
	// the user.
	"DataprocCluster": 60 * time.Minute,
	// The DataFusionInstance integration test fails due to timeouts.
	// Use a custom timeout of 80 minutes which is what the DCL team uses for their integration test.
	"DataFusionInstance": 80 * time.Minute,
}
