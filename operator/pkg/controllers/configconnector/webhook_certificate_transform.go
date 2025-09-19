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
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1 "k8s.io/api/core/v1"
)

const (
	// CAKeyName is the name of the CA private key
	CAKeyName = "ca-key.pem"
	// CACertName is the name of the CA certificate
	CACertName = "ca-cert.pem"
	// ServerKeyName is the name of the server private key
	ServerKeyName = "key.pem"
	// ServerCertName is the name of the serving certificate
	ServerCertName = "cert.pem"

	// InjectCAAnnotation is the annotation to add to webhook configurations to indicate
	// that the CA bundle should be injected.
	// The value is the DNS name.
	InjectCAAnnotation = "cnrm.cloud.google.com/inject-ca"

	// ForceRotationAnnotation is the annotation to add to webhook pods to force a restart
	// when the certificate changes.
	ForceRotationAnnotation = "cnrm.cloud.google.com/certificate-revision"

	// minimumCertValidityDuration is the minimum duration that a certificate must be valid for
	// before it is rotated.
	minimumCertValidityDuration = time.Hour * 24 * 30
)

// injectWebhookCertificates will create and inject CA certificates into webhook configurations
func (r *Reconciler) injectWebhookCertificates(kubeClient client.Client) declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		log := klog.FromContext(ctx)

		injectCAAnnotations := sets.NewString()
		for _, obj := range m.Items {
			annotations := obj.UnstructuredObject().GetAnnotations()
			s := annotations[InjectCAAnnotation]
			if s == "" {
				continue
			}
			injectCAAnnotations.Insert(s)
		}

		if len(injectCAAnnotations) == 0 {
			// No certificates to inject.
			return nil
		}

		for _, injectCAAnnotation := range injectCAAnnotations.List() {
			tokens := strings.Split(injectCAAnnotation, ".")
			if len(tokens) != 2 {
				log.Info("ignoring invalid CA annotation", "annotation", InjectCAAnnotation, "value", injectCAAnnotation)
				continue
			}
			serviceID := types.NamespacedName{Namespace: tokens[1], Name: tokens[0]}

			secret := &corev1.Secret{}
			secretID := serviceNameToSecretID(serviceID)

			if err := kubeClient.Get(ctx, secretID, secret); err != nil {
				if client.IgnoreNotFound(err) != nil {
					return fmt.Errorf("error getting webhook certificate secret %s: %w", secretID, err)
				}
			}

			if !isSecretValid(ctx, secret, serviceID, minimumCertValidityDuration) {
				// TODO: Generate and store a new certificate.
				newSecret, err := uploadWebhookCertificateSecret(ctx, kubeClient, serviceID, secret)
				if err != nil {
					return fmt.Errorf("error rotating webhook certificate secret %s: %w", secretID, err)
				}
				secret = newSecret
			}

			// Add the CA bundle to the webhook configurations.
			caBundle := secret.Data[CACertName]
			if len(caBundle) == 0 {
				// This should not happen because isSecretValid should have caught this.
				return fmt.Errorf("webhook certificate secret %s does not contain CA cert key %q", secretID, CACertName)
			}

			// Set the CA bundle in each webhook configuration.
			for i, obj := range m.Items {
				var hook *unstructured.Unstructured
				switch obj.Kind {
				case "ValidatingWebhookConfiguration":
					hook = obj.UnstructuredObject()

				case "MutatingWebhookConfiguration":
					hook = obj.UnstructuredObject()
				}

				if hook == nil {
					continue
				}

				webhooks, found, err := unstructured.NestedSlice(hook.Object, "webhooks")
				if err != nil {
					return fmt.Errorf("error getting webhooks from %s/%s: %w", obj.GetName(), obj.GetNamespace(), err)
				}
				if !found {
					return fmt.Errorf("could not get webhooks from %s/%s", obj.GetName(), obj.GetNamespace())
				}
				for i := range webhooks {
					webhook, ok := webhooks[i].(map[string]interface{})
					if !ok {
						return fmt.Errorf("expected webhook to be a map[string]interface{}, got %T", webhooks[i])
					}
					if err := unstructured.SetNestedField(webhook, string(caBundle), "clientConfig", "caBundle"); err != nil {
						return fmt.Errorf("error setting caBundle for webhook %s/%s: %w", obj.GetName(), obj.GetNamespace(), err)
					}
				}

				item, err := manifest.NewObject(hook)
				if err != nil {
					return fmt.Errorf("error creating new manifest object for %s/%s: %w", obj.GetName(), obj.GetNamespace(), err)
				}
				m.Items[i] = item
			}

			// Add an annotation to the webhook pods to force a restart when the certificate changes.
			caBundleHash := sha256.Sum256(caBundle)
			caBundleHashHex := hex.EncodeToString(caBundleHash[:])

			annotations := map[string]string{}
			annotations[ForceRotationAnnotation] = caBundleHashHex
			annotations[InjectCAAnnotation] = injectCAAnnotation

			for _, obj := range m.Items {
				u := obj.UnstructuredObject()
				oldAnnotations, _, _ := unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")

				// Make sure we are injecting the correct certificate
				if annotations[InjectCAAnnotation] != oldAnnotations[InjectCAAnnotation] {
					continue
				}

				obj.AddAnnotations(annotations)
			}
		}

		return nil
	}
}

func isSecretValid(ctx context.Context, secret *corev1.Secret, serviceID types.NamespacedName, duration time.Duration) bool {
	log := klog.FromContext(ctx)

	if secret == nil {
		log.Info("certificate not valid: secret is nil")
		return false
	}
	if secret.Data == nil {
		log.Info("certificate not valid: secret data is nil")
		return false
	}
	data := &CertificateData{
		CAKey:      secret.Data[CAKeyName],
		CACert:     secret.Data[CACertName],
		ServerKey:  secret.Data[ServerKeyName],
		ServerCert: secret.Data[ServerCertName],
	}
	if len(data.CAKey) == 0 || len(data.CACert) == 0 || len(data.ServerKey) == 0 || len(data.ServerCert) == 0 {
		log.Info("certificate not valid: secret data is incomplete")
	}

	dnsName := serviceNameToDNSName(serviceID)
	return validCert(ctx, data, dnsName, duration)
}

type CertificateData struct {
	CACert     []byte
	CAKey      []byte
	ServerCert []byte
	ServerKey  []byte
}

// validCert verifies if the certificate is valid, including
// additional verifications to ensure compatibility with Kubernetes
// and it's default HTTP client.
func validCert(ctx context.Context, certs *CertificateData, dnsName string, duration time.Duration) bool {
	log := klog.FromContext(ctx)
	if certs == nil {
		log.Info("certificate not valid: certs are nil")
		return false
	}

	// Verify key and cert are valid pair
	_, err := tls.X509KeyPair(certs.ServerCert, certs.ServerKey)
	if err != nil {
		log.Error(err, "certificate not valid: failed to parse key pair")
		return false
	}

	// Verify cert is good for desired DNS name and signed by CA and will be valid for desired period of time.
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(certs.CACert) {
		log.Info("certificate not valid: failed to parse CA cert")
		return false
	}
	block, _ := pem.Decode([]byte(certs.ServerCert))
	if block == nil {
		log.Info("certificate not valid: failed to parse server cert")
		return false
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Error(err, "certificate not valid: failed to parse server cert")
		return false
	}

	ops := x509.VerifyOptions{
		DNSName:     dnsName,
		Roots:       pool,
		CurrentTime: time.Now().Add(duration),
	}
	if _, err := cert.Verify(ops); err != nil {
		log.Error(err, "certificate not valid: failed to verify certificate")
		return false
	}

	if err := ValidateCertificateWorksWithK8sAPIClient(cert); err != nil {
		log.Error(err, "certificate not valid: not compatible with Kubernetes API client")
		return false
	}
	return true
}

// ValidateCertificateWorksWithK8sAPIClient returns false if the certificate
// is not compatible with Kubernetes HTTP clients.
func ValidateCertificateWorksWithK8sAPIClient(cert *x509.Certificate) error {
	// check to see if the cert has a DNSName. Certificates that
	// do not will not be considered valid for Kubernetes distributions built
	// with go 1.15 or higher.
	if len(cert.DNSNames) == 0 {
		return fmt.Errorf("certificate does not have a DNS name")
	}
	return nil
}

func uploadWebhookCertificateSecret(ctx context.Context, kubeClient client.Client, serviceID types.NamespacedName, existingSecret *corev1.Secret) (*corev1.Secret, error) {
	dnsName := serviceNameToDNSName(serviceID)
	certificate, err := newWebhookCertificate(dnsName)
	if err != nil {
		return nil, fmt.Errorf("error generating self-signed certificate: %w", err)
	}

	secretID := serviceNameToSecretID(serviceID)

	newSecret := buildSecretForCertificate(certificate, secretID)

	if existingSecret == nil {
		if err := kubeClient.Create(ctx, newSecret); err != nil {
			return nil, fmt.Errorf("error creating webhook certificate secret %s/%s: %w", newSecret.Namespace, newSecret.Name, err)
		}

		return newSecret, nil
	} else {
		existingSecret.Data = newSecret.Data
		existingSecret.Type = newSecret.Type
		if err := kubeClient.Update(ctx, existingSecret); err != nil {
			return nil, fmt.Errorf("error updating webhook certificate secret %s/%s: %w", newSecret.Namespace, newSecret.Name, err)
		}

		return existingSecret, nil
	}
}

func buildSecretForCertificate(certificate *CertificateData, secretID types.NamespacedName) *corev1.Secret {
	newSecret := corev1.Secret{}
	newSecret.Name = secretID.Name
	newSecret.Namespace = secretID.Namespace

	newSecret.Type = corev1.SecretTypeTLS

	newSecret.Data = map[string][]byte{
		CAKeyName:      certificate.CAKey,
		CACertName:     certificate.CACert,
		ServerKeyName:  certificate.ServerKey,
		ServerCertName: certificate.ServerCert,
	}

	return &newSecret
}

func serviceNameToDNSName(name types.NamespacedName) string {
	return name.Name + "." + name.Namespace
}

func serviceNameToSecretID(name types.NamespacedName) types.NamespacedName {
	return types.NamespacedName{
		Name:      "webhook-cert-" + name.Name,
		Namespace: name.Namespace,
	}
}
