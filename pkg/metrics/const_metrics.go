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

package metrics

import (
	"context"
	"os"

	"github.com/prometheus/procfs"
	"go.opencensus.io/stats"
)

func RecordProcessStartTime() error {
	p, err := procfs.NewProc(os.Getpid())
	if err != nil {
		return err
	}
	stat, err := p.NewStat()
	if err != nil {
		return err
	}
	startTime, err := stat.StartTime()
	if err != nil {
		return err
	}
	stats.Record(context.Background(), MProcessStartTime.M(startTime))
	return nil
}
