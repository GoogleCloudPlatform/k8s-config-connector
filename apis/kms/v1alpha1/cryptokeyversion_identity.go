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

package v1alpha1

import (
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"strings"
)

type KMSCryptoKeyVersionIdentity struct {
	parent *v1beta1.KMSCryptoKeyIdentity
	id     string
}

func (i *KMSCryptoKeyVersionIdentity) String() string {
	return i.parent.String() + "/cryptoKeyVersions/" + i.id
}

func ParseKMSCryptoKeyVersionExternal(external string) (parent *KMSCryptoKeyVersionIdentity, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	p, err := v1beta1.ParseKMSCryptoKeyExternal(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return nil, err
	}
	if tokens[len(tokens)-2] == "cryptoKeyVersions" {
		return &KMSCryptoKeyVersionIdentity{parent: p, id: tokens[len(tokens)-1]}, nil
	}
	return nil, fmt.Errorf("format of KMSCryptoKeyVersion external=%q was not known (use %s/cryptoKeyVersions/{{key})", external, p)

}
