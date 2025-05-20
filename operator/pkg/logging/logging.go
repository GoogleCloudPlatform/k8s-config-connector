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

// Package logging adds common logging hooks for cnrm applications
package logging

import (
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var logger = log.Log

// BuildLogger constructs a logr.Logger object that matches the standard
// configuration across cnrm applications, writing to the io.Writer passed.
func BuildLogger(output io.Writer) logr.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "severity",
		NameKey:        "logger",
		TimeKey:        "timestamp",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	return zapr.NewLogger(zap.NewRaw(zap.WriteTo(output), zap.Encoder(encoder)))
}

// Fatal is a utility function to replace log.Fatal, which doesn't exist
// for logr loggers.
func Fatal(err error, msg string) {
	logger.Error(err, msg)
	os.Exit(1)
}
