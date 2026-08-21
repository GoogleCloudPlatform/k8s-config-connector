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

package testvariable

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"
)

func NewUniqueID() string {
	var id string
	for i := 0; i < 3; i++ {
		// generate a globally unique ID
		id = randomid.New().String()
		if isValidID(id) {
			return id
		}
	}
	panic(fmt.Sprintf("unable to generate a valid id, last attempt resulted in value: %v", id))
}

func isValidID(id string) bool {
	// ssl and google are reserved words and cannot be used in project ids, a valid lookbehind regex would be
	// "((?!ssl|google).*)", but golang does not support lookbehinds. For that reason, just throw out any id
	// that contains these reserved words
	return !strings.Contains(id, "google") && !strings.Contains(id, "ssl")
}

// RandomIDGenerator generates a randomized string of a given length and
// contains only the characters defined in the given regex
// the IDs returned by this are NOT globally unique so this should only be used in locations where there cannot be
// collisions across tests
func RandomIDGenerator(reg *regexp.Regexp, length uint) string {
	var stringBuffer strings.Builder

	// Keep generating and appending randomized characters until the length is hit
	for uint(len(stringBuffer.String())) < length {
		b := make([]byte, length*5) // Add more to account for filtering out unwanted chars
		_, err := rand.Read(b)
		if err != nil {
			log.Fatalf("Could not read random bytes: %v", err)
		}

		s := base64.URLEncoding.EncodeToString(b)
		allowedChars := reg.FindAllString(s, -1)
		stringBuffer.WriteString(strings.Join(allowedChars, ""))
	}
	return stringBuffer.String()[0:length]
}
