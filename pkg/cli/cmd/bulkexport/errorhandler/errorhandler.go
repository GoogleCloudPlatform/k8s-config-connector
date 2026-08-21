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

package errorhandler

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/log"
)

type Handler interface {
	Handle(err error) error
}

type haltHandler struct{}
type continueHandler struct{}
type ignoreHandler struct{}

func NewHalt() Handler {
	return &haltHandler{}
}

func (h *haltHandler) Handle(err error) error {
	return err
}

func NewContinue() Handler {
	return &continueHandler{}
}

func (c *continueHandler) Handle(err error) error {
	log.Err(err)
	return nil
}

func NewIgnore() Handler {
	return &ignoreHandler{}
}

func (i *ignoreHandler) Handle(_ error) error {
	return nil
}

func NewErrorHandler(params *parameters.Parameters) (Handler, error) {
	switch params.OnError {
	case parameters.HaltOnErrorOption:
		return NewHalt(), nil
	case parameters.ContinueOnErrorOption:
		return NewContinue(), nil
	case parameters.IgnoreOnErrorOption:
		return NewIgnore(), nil
	default:
		return nil, fmt.Errorf("unknown '%v' option '%v', must be one of: %v",
			parameters.OnErrorParam, params.OnError, strings.Join(parameters.AllErrorOptions, ", "))
	}
}
