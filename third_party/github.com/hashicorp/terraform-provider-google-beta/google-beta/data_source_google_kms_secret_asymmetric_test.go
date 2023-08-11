// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"hash/crc32"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccKmsSecretAsymmetricBasic(t *testing.T) {
	// Nested tests confuse VCR
	acctest.SkipIfVcr(t)
	t.Parallel()

	projectOrg := envvar.GetTestOrgFromEnv(t)
	projectBillingAccount := envvar.GetTestBillingAccountFromEnv(t)

	projectID := "tf-test-" + acctest.RandString(t, 10)
	keyRingName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	cryptoKeyName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	plaintext := fmt.Sprintf("secret-%s", acctest.RandString(t, 10))

	// The first test creates resources needed to encrypt plaintext and produce ciphertext
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: kmsCryptoKeyAsymmetricDecryptBasic(projectID, projectOrg, projectBillingAccount, keyRingName, cryptoKeyName),
				Check: func(s *terraform.State) error {
					ciphertext, cryptoKeyVersionID, crc, err := testAccEncryptSecretDataAsymmetricWithPublicKey(t, s, "data.google_kms_crypto_key_version.crypto_key", plaintext)
					if err != nil {
						return err
					}

					// The second test asserts that the data source has the correct plaintext, given the created ciphertext
					acctest.VcrTest(t, resource.TestCase{
						PreCheck:                 func() { acctest.AccTestPreCheck(t) },
						ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
						Steps: []resource.TestStep{
							{
								Config: googleKmsSecretAsymmetricDatasource(cryptoKeyVersionID, ciphertext),
								Check:  resource.TestCheckResourceAttr("data.google_kms_secret_asymmetric.acceptance", "plaintext", plaintext),
							},
							{
								Config: googleKmsSecretAsymmetricDatasourceWithCrc(cryptoKeyVersionID, ciphertext, crc),
								Check:  resource.TestCheckResourceAttr("data.google_kms_secret_asymmetric.acceptance_with_crc", "plaintext", plaintext),
							},
						},
					})

					return nil
				},
			},
		},
	})
}

func testAccEncryptSecretDataAsymmetricWithPublicKey(t *testing.T, s *terraform.State, cryptoKeyResourceName, plaintext string) (string, string, uint32, error) {
	rs, ok := s.RootModule().Resources[cryptoKeyResourceName]
	if !ok {
		return "", "", 0, fmt.Errorf("resource not found: %s", cryptoKeyResourceName)
	}

	cryptoKeyVersionID := rs.Primary.Attributes["id"]

	block, _ := pem.Decode([]byte(rs.Primary.Attributes["public_key.0.pem"]))
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", "", 0, fmt.Errorf("failed to parse public key: %v", err)
	}
	rsaKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return "", "", 0, fmt.Errorf("public key is not rsa")
	}

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaKey, []byte(plaintext), nil)
	if err != nil {
		return "", "", 0, fmt.Errorf("rsa.EncryptOAEP: %v", err)
	}

	crc := crc32.Checksum(ciphertext, crc32.MakeTable(crc32.Castagnoli))

	result := base64.StdEncoding.EncodeToString(ciphertext)
	log.Printf("[INFO] Successfully encrypted plaintext and got ciphertext: %s", result)

	return result, cryptoKeyVersionID, crc, nil
}

func googleKmsSecretAsymmetricDatasource(cryptoKeyTerraformID, ciphertext string) string {
	return fmt.Sprintf(`
data "google_kms_secret_asymmetric" "acceptance" {
  crypto_key_version = "%s"
  ciphertext         = "%s"
}
`, cryptoKeyTerraformID, ciphertext)
}

func googleKmsSecretAsymmetricDatasourceWithCrc(cryptoKeyTerraformID, ciphertext string, crc uint32) string {
	return fmt.Sprintf(`
data "google_kms_secret_asymmetric" "acceptance_with_crc" {
  crypto_key_version = "%s"
  ciphertext         = "%s"
  crc32              = "%x"
}
`, cryptoKeyTerraformID, ciphertext, crc)
}

func kmsCryptoKeyAsymmetricDecryptBasic(projectID, projectOrg, projectBillingAccount, keyRingName, cryptoKeyName string) string {
	return fmt.Sprintf(`
resource "google_project" "acceptance" {
  name            = "%s"
  project_id      = "%s"
  org_id          = "%s"
  billing_account = "%s"
}

resource "google_project_service" "acceptance" {
  project = google_project.acceptance.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_kms_key_ring" "key_ring" {
  project  = google_project_service.acceptance.project
  name     = "%s"
  location = "us-central1"
  depends_on = [google_project_service.acceptance]
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "%s"
  key_ring = google_kms_key_ring.key_ring.id
  purpose  = "ASYMMETRIC_DECRYPT"
  version_template {
    algorithm = "RSA_DECRYPT_OAEP_4096_SHA256"
  }
}

data "google_kms_crypto_key_version" "crypto_key" {
  crypto_key = google_kms_crypto_key.crypto_key.id
}
`, projectID, projectID, projectOrg, projectBillingAccount, keyRingName, cryptoKeyName)
}
