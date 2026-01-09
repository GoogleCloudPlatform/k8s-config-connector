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

package iam

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"google.golang.org/api/iam/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
)

func init() {
	registry.RegisterModel(krm.IAMServiceAccountKeyGVK, NewIAMServiceAccountKeyModel)
}

func NewIAMServiceAccountKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIAMServiceAccountKey{config: *config}, nil
}

var _ directbase.Model = &modelIAMServiceAccountKey{}

type modelIAMServiceAccountKey struct {
	config config.ControllerConfig
}

func (m *modelIAMServiceAccountKey) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.IAMServiceAccountKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	gcpClient, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCP client: %w", err)
	}

	return &IAMServiceAccountKeyAdapter{
		gcpClient:  gcpClient,
		kubeClient: reader.(client.Client), // Reconciler passes client.Client as reader
		desired:    obj,
	}, nil
}

func (m *modelIAMServiceAccountKey) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type IAMServiceAccountKeyAdapter struct {
	gcpClient  *iam.Service
	kubeClient client.Client
	desired    *krm.IAMServiceAccountKey
	actual     *iam.ServiceAccountKey
}

var _ directbase.Adapter = &IAMServiceAccountKeyAdapter{}

func (a *IAMServiceAccountKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("Finding IAMServiceAccountKey", "name", a.desired.Name)

	if a.desired.Status.Name == nil || *a.desired.Status.Name == "" {
		return false, nil
	}

	keyName := *a.desired.Status.Name
	gcpKey, err := a.gcpClient.Projects.ServiceAccounts.Keys.Get(keyName).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IAMServiceAccountKey %q: %w", keyName, err)
	}

	a.actual = gcpKey
	return true, nil
}

func (a *IAMServiceAccountKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("Creating IAMServiceAccountKey", "name", a.desired.Name)

	// 1. Resolve parent Service Account
	parentSA, err := a.resolveServiceAccountRef(ctx) // todo acpana: resolve with the current pattern?
	if err != nil {
		return err
	}

	// 2. Client-side Key Generation
	privateKey, publicKeyPEM, err := generateRSAKeyPair()
	if err != nil {
		return fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	// 3. Persist Private Key to Secret FIRST (Write-Ahead)
	// We use the resource name as the secret name by default
	secretName := a.desired.Name
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: a.desired.Namespace,
			Labels: map[string]string{
				label.CnrmManagedKey: "true",
			},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(a.desired, krm.IAMServiceAccountKeyGVK),
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"privateKey": privateKey, // Temporary field for recovery
		},
	}

	if err := a.kubeClient.Create(ctx, secret); err != nil {
		return fmt.Errorf("failed to create Secret %s for private key: %w", secretName, err)
	}
	log.Info("created Secret for private key", "secret", secretName)

	// 4. Upload Public Key to GCP
	uploadReq := &iam.UploadServiceAccountKeyRequest{
		PublicKeyData: base64.StdEncoding.EncodeToString(publicKeyPEM),
	}
	gcpKey, err := a.gcpClient.Projects.ServiceAccounts.Keys.Upload(parentSA, uploadReq).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to upload public key to GCP for %s: %w", parentSA, err)
	}
	log.Info("successfully uploaded public key to GCP", "key", gcpKey.Name)

	// 5. Update Secret with standard JSON format
	// The standard format includes project_id, private_key_id, etc.
	projectID := strings.Split(parentSA, "/")[1]
	saEmail := strings.Split(parentSA, "/")[3]
	keyID := strings.Split(gcpKey.Name, "/")[5]

	credentialJSON, err := constructCredentialJSON(projectID, saEmail, keyID, string(privateKey))
	if err != nil {
		return fmt.Errorf("failed to construct credential JSON: %w", err)
	}

	secret.Data["key.json"] = []byte(credentialJSON)
	delete(secret.Data, "privateKey") // Remove recovery field after success
	if err := a.kubeClient.Update(ctx, secret); err != nil {
		// Log error but don't fail reconciliation; we have the key in GCP now.
		// Next reconcile can retry secret update.
		log.Error(err, "failed to update Secret with final key.json", "secret", secretName)
	}

	// 6. Update Status
	status := &krm.IAMServiceAccountKeyStatus{
		Name:        &gcpKey.Name,
		ValidAfter:  &gcpKey.ValidAfterTime,
		ValidBefore: &gcpKey.ValidBeforeTime,
		// Note: We deliberately DON'T put the private key in status for security and to avoid the lost-key race condition
	}
	// We might want to set status.PublicKey if needed, but it's optional

	readyCondition := k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)
	return createOp.UpdateStatus(ctx, status, &readyCondition)
}

func (a *IAMServiceAccountKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// IAM Service Account Keys are immutable.
	// We might support updating annotations/labels if they were relevant, but typically we just return.
	return nil
}

func (a *IAMServiceAccountKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return true, nil
	}

	log.Info("deleting IAMServiceAccountKey", "name", a.actual.Name)
	_, err := a.gcpClient.Projects.ServiceAccounts.Keys.Delete(a.actual.Name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting IAMServiceAccountKey %q: %w", a.actual.Name, err)
	}

	log.Info("successfully deleted IAMServiceAccountKey", "name", a.actual.Name)
	return true, nil
}

func (a *IAMServiceAccountKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Helpers

func (a *IAMServiceAccountKeyAdapter) resolveServiceAccountRef(ctx context.Context) (string, error) {
	ref := a.desired.Spec.ServiceAccountRef
	if ref.External != "" {
		if !strings.HasPrefix(ref.External, "projects/") {
			// Assume it's just the email, we need to construct the full resource name
			// But we don't have the project easily...
			// In KCC, we usually resolve the project from the resource or namespace.
			// For now, let's assume External is the full resource name or we fail.
			return "", fmt.Errorf("external serviceAccountRef must be a full resource name (projects/.../serviceAccounts/...) or use name/namespace")
		}
		return ref.External, nil
	}

	// Resolve local reference
	sa := &unstructured.Unstructured{}
	sa.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "iam.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "IAMServiceAccount",
	})

	ns := ref.Namespace
	if ns == "" {
		ns = a.desired.Namespace
	}
	key := client.ObjectKey{Name: ref.Name, Namespace: ns}
	if err := a.kubeClient.Get(ctx, key, sa); err != nil {
		return "", fmt.Errorf("failed to resolve serviceAccountRef %v: %w", key, err)
	}

	email, found, err := unstructured.NestedString(sa.Object, "status", "email")
	if err != nil || !found || email == "" {
		return "", fmt.Errorf("referenced IAMServiceAccount %v does not have an email in status", key)
	}

	projectID, found, err := unstructured.NestedString(sa.Object, "spec", "projectID")
	// If projectID not in spec, we'd normally resolve it from annotations or namespace.
	// Simplified for now.
	if !found || projectID == "" {
		// Try annotation
		annotations := sa.GetAnnotations()
		if val, ok := annotations["cnrm.cloud.google.com/project-id"]; ok {
			projectID = val
		} else {
			projectID = ns // Default fallback
		}
	}

	return fmt.Sprintf("projects/%s/serviceAccounts/%s", projectID, email), nil
}

func generateRSAKeyPair() (privatePEM []byte, publicPEM []byte, err error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	privatePEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})

	pubBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})

	return privatePEM, publicPEM, nil
}

func constructCredentialJSON(projectID, saEmail, keyID, privateKey string) (string, error) {
	// Standard Google Service Account Key JSON format
	creds := map[string]string{
		"type":                        "service_account",
		"project_id":                  projectID,
		"private_key_id":              keyID,
		"private_key":                 privateKey,
		"client_email":                saEmail,
		"client_id":                   "", // Optional?
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        fmt.Sprintf("https://www.googleapis.com/robot/v1/metadata/x509/%s", saEmail),
	}
	b, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
