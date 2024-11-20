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

package parameters

import (
	"fmt"
	"os"

	bulkparams "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/valutil"
)

const (
	InputParam = "input-file"
)

type Parameters struct {
	Input       string
	OAuth2Token string
	Verbose     bool
}

func Validate(p *Parameters, stdin *os.File) error {
	inputFile := &p.Input
	piped, err := bulkparams.IsInputPiped(stdin)
	if err != nil {
		return err
	}
	if piped {
		if !valutil.IsDefaultValue(inputFile) {
			return fmt.Errorf("cannot supply input on stdin with the '%v' parameter", InputParam)
		}
	}
	if valutil.IsDefaultValue(inputFile) {
		return fmt.Errorf("'%v' parameter cannot be empty", InputParam)
	}
	return nil
}
