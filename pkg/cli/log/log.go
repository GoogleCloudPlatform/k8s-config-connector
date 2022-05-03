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

package log

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	defaultLogger = New(false)
	ePrintf       = color.New(color.FgRed).FprintfFunc()
)

// a logger suitable for a CLI: no printing of timestamps, filenames, etc.
type Logger struct {
	enableVerbose bool
}

func New(enableVerbose bool) *Logger {
	return &Logger{
		enableVerbose: enableVerbose,
	}
}

func SetDefault(logger *Logger) {
	defaultLogger = logger
}

func (l *Logger) Info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, fmt.Sprintf("%s\n", format), args...)
}

func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

func (l *Logger) Verbose(format string, args ...interface{}) {
	if l.enableVerbose {
		l.Info(format, args...)
	}
}

func Verbose(format string, args ...interface{}) {
	defaultLogger.Verbose(format, args...)
}

func (l *Logger) Err(err error) {
	l.Error("%v", err)
}

func Err(err error) {
	defaultLogger.Err(err)
}

func (l *Logger) Error(format string, args ...interface{}) {
	ePrintf(os.Stderr, fmt.Sprintf("%s\n", format), args...)
}

func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}
