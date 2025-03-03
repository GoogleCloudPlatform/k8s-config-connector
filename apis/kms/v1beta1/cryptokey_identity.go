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

package v1beta1

import (
	"fmt"
	"strings"
)

type KMSCryptoKeyIdentity struct {
	parent *KMSKeyRingIdentity
	id     string
}

func (i *KMSCryptoKeyIdentity) String() string {
	return i.parent.String() + "/cryptoKeys/" + i.id
}

func ParseKMSCryptoKeyExternal(external string) (parent *KMSCryptoKeyIdentity, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	p, err := ParseKMSKeyRingExternal(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return nil, err
	}
	if tokens[len(tokens)-2] == "cryptoKeys" {
		return &KMSCryptoKeyIdentity{parent: p, id: tokens[len(tokens)-1]}, nil
	}
	return nil, fmt.Errorf("format of KMSCryptoKey external=%q was not known (use %s/cryptoKeys/{{key})", external, p)

}
