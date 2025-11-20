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

package configconnector

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/types"
)

func TestKeygen(t *testing.T) {
	ctx := context.TODO()

	serviceID := types.NamespacedName{Name: "myservice", Namespace: "mynamespace"}

	dnsName := serviceNameToDNSName(serviceID)
	certificate, err := newWebhookCertificate(dnsName)
	if err != nil {
		t.Fatalf("error from newWebhookCertificate: %v", err)
	}

	secretID := serviceNameToSecretID(serviceID)

	newSecret := buildSecretForCertificate(certificate, secretID)

	if !isSecretValid(ctx, newSecret, serviceID, minimumCertValidityDuration) {
		t.Fatalf("generated secret is not valid for minimum duration %v", minimumCertValidityDuration)
	}

	fiveYears := time.Duration(5 * 365 * 24 * time.Hour)
	if !isSecretValid(ctx, newSecret, serviceID, fiveYears) {
		t.Fatalf("generated secret is not valid for five years")
	}

	justPastExpiry := certificateDuration + time.Hour
	if isSecretValid(ctx, newSecret, serviceID, justPastExpiry) {
		t.Fatalf("generated secret is valid past expiry")
	}
}
