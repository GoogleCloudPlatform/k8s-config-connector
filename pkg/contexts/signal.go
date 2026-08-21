// Copyright 2025 Google LLC
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

package contexts

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler registers for SIGTERM and SIGINT. A context is returned
// which is canceled on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
//
// This code is a variant of the code from controller-runtime/pkg/signals, modified to use
// context.WithCancelCause so that the reason for cancellation can be propagated.
// This allows callers to distinguish between normal cancellation and cancellation
// due to termination signals.
func SetupSignalHandler() context.Context {
	close(onlyOneSignalHandler) // panics when called twice

	ctx, cancel := context.WithCancelCause(context.Background())

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		sig := <-c
		cancel(fmt.Errorf("received termination signal %s: %w", sig, context.Canceled))
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}

// IsContextCanceledErr checks whether the given error is caused by context cancellation.
func IsContextCanceledErr(err error) bool {
	return err != nil && errors.Is(err, context.Canceled)
}
