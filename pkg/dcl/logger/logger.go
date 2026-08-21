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

package logger

import (
	"fmt"
	"os"

	"k8s.io/klog/v2"
)

var dclDebug = os.Getenv("DCL_DEBUG")

// A wrapper of klog that implements DCL logger interface and only logs fatal and warnings, unless DCL_DEBUG env variable is set.
type dclLogger struct {
	debug bool
}

func SimpleDCLLogger() dclLogger { //nolint:revive
	return dclLogger{
		debug: dclDebug != "",
	}
}

func (l dclLogger) Fatal(args ...interface{}) {
	klog.Fatal(args...)
}

func (l dclLogger) Fatalf(format string, args ...interface{}) {
	klog.Fatalf(fmt.Sprintf("[DCL FATAL] %s", format), args...)
}

func (l dclLogger) Info(args ...interface{}) {
	if l.debug {
		klog.Info(args...)
	}
}

func (l dclLogger) Infof(format string, args ...interface{}) {
	if l.debug {
		klog.Infof(fmt.Sprintf("[DCL INFO] %s", format), args...)
	}
}

func (l dclLogger) Warningf(format string, args ...interface{}) {
	klog.Warningf(fmt.Sprintf("[DCL WARNING] %s", format), args...)
}

func (l dclLogger) Warning(args ...interface{}) {
	klog.Warning(args...)
}
