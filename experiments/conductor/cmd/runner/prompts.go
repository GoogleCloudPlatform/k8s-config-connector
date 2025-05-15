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

package runner

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed *.txt
var fs embed.FS

func getTemplate(key string) (*template.Template, error) {
	b, err := fs.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("reading embedded file %q: %w", key, err)
	}

	t, err := template.New(key).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("parsing template %q: %w", key, err)
	}
	return t, nil
}

type MockgcpGenerateGcloudTestPrompt struct {
	GcloudCommand string
	Group         string
	Resource      string
}

func (o *MockgcpGenerateGcloudTestPrompt) Generate() ([]byte, error) {
	key := "mockgcp-generate-gcloud-test.prompt.txt"
	t, err := getTemplate(key)
	if err != nil {
		return nil, err
	}

	var bb bytes.Buffer
	if err := t.Execute(&bb, o); err != nil {
		return nil, fmt.Errorf("running %q template: %w", key, err)
	}

	return bb.Bytes(), nil
}
