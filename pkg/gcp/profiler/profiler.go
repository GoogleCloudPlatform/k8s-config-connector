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

package profiler

import (
	"fmt"

	"cloud.google.com/go/profiler"
	flag "github.com/spf13/pflag"
)

var profilerServiceName string

func AddFlag(flagSet *flag.FlagSet) {
	flagSet.StringVar(&profilerServiceName, "profiler-service-name", "", "when specified, the profiler agent is enabled with a service name equal to the value. See https://cloud.google.com/profiler/docs/profiling-go#gke for details.")
}

// Starts the Cloud Profiler agent if the 'profile' option is true
func StartIfEnabled() error {
	if profilerServiceName == "" {
		return nil
	}
	return start()
}

// Starts the Cloud Profiler agent, should be called as soon as possible
func start() error {
	cfg := profiler.Config{
		Service:      profilerServiceName,
		DebugLogging: true,
	}
	// Profiler initialization, best done as early as possible.
	if err := profiler.Start(cfg); err != nil {
		return fmt.Errorf("error starting profiler: %w", err)
	}
	return nil
}
