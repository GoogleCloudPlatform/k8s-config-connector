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

package krmtotf

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func NewErrorFromDiagnostics(diagnostics diag.Diagnostics) error {
	if !diagnostics.HasError() {
		return nil
	}
	msgs := make([]string, 0)
	for _, d := range diagnostics {
		if d.Severity == diag.Error {
			// Only append detail if it's rendered.
			if d.Detail == "" {
				msgs = append(msgs, fmt.Sprintf("summary: %v", d.Summary))
			} else {
				msgs = append(msgs, fmt.Sprintf("summary: %v, detail: %v", d.Summary, d.Detail))
			}
		}
	}
	return errors.New(strings.Join(msgs, "\n"))
}
