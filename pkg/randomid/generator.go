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

package randomid

import (
	cryptorand "crypto/rand"
	"encoding/base32"

	"k8s.io/klog/v2"
)

// ID holds a random identifier
type ID struct {
	b []byte
}

// New builds a random identifier.
// We always generate a more secure random value; we could in future expose less secure options if this becomes a bottleneck.
func New() ID {
	// The forwarding rule name for PSC Google APIs must be an 1-20 characters string with lowercase letters and numbers and must start with a letter.
	// 9 bytes is equivalent to 72 bits of data. Each character in Base32 represents 5 bits.
	// Calculation: 72 bits / 5 bits per character = 14.4 characters.
	// Round up to the nearest whole number, which is 15. the ID string will have 15 characters.
	b := make([]byte, 9)
	if _, err := cryptorand.Read(b); err != nil {
		klog.Fatalf("failed to read from crypto/rand: %v", err)
	}
	return ID{b}
}

// encoding is like base32.StdEncoding but in lower case and without padding
var encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)

// String returns an ascii (base32) encoding of the random identifier.
func (i ID) String() string {
	return encoding.EncodeToString(i.b)
}
