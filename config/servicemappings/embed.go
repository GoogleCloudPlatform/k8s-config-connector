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

package servicemappings

import (
	"embed"
	"fmt"
)

//go:embed *.yaml
var servicemappings embed.FS

func ServiceMapping(key string) ([]byte, error) {
	b, err := servicemappings.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("error reading embedded file %q: %w", key, err)
	}
	return b, nil
}

func AllKeys() ([]string, error) {
	p := "."
	entries, err := servicemappings.ReadDir(p)
	if err != nil {
		return nil, fmt.Errorf("error reading embedded directory %q: %w", p, err)
	}
	var keys []string
	for _, entry := range entries {
		keys = append(keys, entry.Name())
	}
	return keys, nil
}
