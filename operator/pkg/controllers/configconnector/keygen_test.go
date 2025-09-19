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
