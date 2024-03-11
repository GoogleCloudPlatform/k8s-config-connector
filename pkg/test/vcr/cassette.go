// Copyright 2024 Google LLC
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

package vcr

import (
	"os"
	"reflect"
	"sync"

	"k8s.io/klog/v2"

	"gopkg.in/yaml.v2"
)

const DIR string = "pkg/test/vcr/fixtures/cassette/"

type Cassette struct {
	Name              string        `yaml:"name"`
	Interactions      []Interaction `yaml:"interactions"`
	Mu                sync.RWMutex  `yaml:"-"`
	NextInteractionID int           `yaml:"-"`
}

type Interaction struct {
	ID       int          `yaml:"id"`
	Request  *VCRRequest  `yaml:"request"`
	Response *VCRResponse `yaml:"response"`
}

func (c *Cassette) fileName() string {
	return DIR + c.Name + ".yaml"
}

func (c *Cassette) Read() {
	//TODO(yuhou): read requests from cassette
}

func (c *Cassette) Write() error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	if _, err := os.Stat(DIR); os.IsNotExist(err) {
		if err = os.MkdirAll(DIR, 0755); err != nil {
			return err
		}
	}
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(c.fileName())
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte("---\n"))
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cassette) MatchInteraction(request *VCRRequest) *Interaction {
	if len(c.Interactions) == 0 {
		klog.Fatal("[VCR] No interactions!")
	}

	i := c.Interactions[0]
	expected := i.Request

	if expected.Method != request.Method {
		errorInteractionMismatch(request, "Method", expected.Method, request.Method)
	}

	if expected.URL != request.URL {
		errorInteractionMismatch(request, "URL", expected.URL, request.URL)
	}

	if !reflect.DeepEqual(expected.Body, request.Body) {
		errorInteractionMismatch(request, "Body", string(expected.Body[:]), string(request.Body[:]))
	}

	c.Interactions = c.Interactions[1:]
	return &i
}

func errorInteractionMismatch(request *VCRRequest, field string, expected string, actual string) {
	klog.Fatalf(
		"[VCR] Error with interaction: %s %s. Field %s does not match: expected: %s, got: %s.",
		request.Method,
		request.URL,
		field,
		expected,
		actual,
	)
}
