// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package kms

import (
	"context"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
	"google.golang.org/api/cloudkms/v1"
)

var (
	cryptoKeyVersionRegexp = regexp.MustCompile(`^(//[^/]*/[^/]*/)?(projects/[^/]+/locations/[^/]+/keyRings/[^/]+/cryptoKeys/[^/]+/cryptoKeyVersions/[^/]+)$`)
)

func DataSourceGoogleKmsSecretAsymmetric() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGoogleKmsSecretAsymmetricReadContext,
		Schema: map[string]*schema.Schema{
			"crypto_key_version": {
				Type:         schema.TypeString,
				Description:  "The fully qualified KMS crypto key version name",
				ValidateFunc: verify.ValidateRegexp(cryptoKeyVersionRegexp.String()),
				Required:     true,
			},
			"ciphertext": {
				Type:         schema.TypeString,
				Description:  "The public key encrypted ciphertext in base64 encoding",
				ValidateFunc: validateBase64WithWhitespaces,
				Required:     true,
			},
			"crc32": {
				Type:         schema.TypeString,
				Description:  "The crc32 checksum of the ciphertext, hexadecimal encoding. If not specified, it will be computed",
				ValidateFunc: validateHexadecimalUint32,
				Optional:     true,
			},
			"plaintext": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceGoogleKmsSecretAsymmetricReadContext(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	err := dataSourceGoogleKmsSecretAsymmetricRead(ctx, d, meta)
	if err != nil {
		diags = diag.FromErr(err)
	}
	return diags
}

func dataSourceGoogleKmsSecretAsymmetricRead(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	// `google_kms_crypto_key_version` returns an id with the prefix
	// //cloudkms.googleapis.com/v1, which is an invalid name. To allow for the most elegant
	// configuration, we will allow it as an input.
	keyVersion := cryptoKeyVersionRegexp.FindStringSubmatch(d.Get("crypto_key_version").(string))
	cryptoKeyVersion := keyVersion[len(keyVersion)-1]

	base64CipherText := removeWhiteSpaceFromString(d.Get("ciphertext").(string))
	ciphertext, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return err
	}

	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)
	}

	ciphertextCRC32C := crc32c(ciphertext)
	if s, ok := d.Get("crc32").(string); ok && s != "" {
		u, err := strconv.ParseUint(s, 16, 32)
		if err != nil {
			return fmt.Errorf("failed to convert crc32 into uint32, %s", err)
		}
		ciphertextCRC32C = uint32(u)
	} else {
		if err := d.Set("crc32", fmt.Sprintf("%x", ciphertextCRC32C)); err != nil {
			return fmt.Errorf("failed to set crc32, %s", err)
		}
	}

	req := cloudkms.AsymmetricDecryptRequest{
		Ciphertext:       base64CipherText,
		CiphertextCrc32c: int64(ciphertextCRC32C)}

	client := config.NewKmsClientWithCtx(ctx, userAgent)
	if client == nil {
		return fmt.Errorf("failed to get a KMS client")
	}

	result, err := client.Projects.Locations.KeyRings.CryptoKeys.CryptoKeyVersions.AsymmetricDecrypt(cryptoKeyVersion, &req).Do()
	if err != nil {
		return fmt.Errorf("failed to decrypt ciphertext: %v", err)
	}
	plaintext, err := base64.StdEncoding.DecodeString(result.Plaintext)
	if err != nil {
		return fmt.Errorf("failed to base64 decode plaintext: %v", err)
	}

	plaintextCrc32c := int64(crc32c(plaintext))
	if !result.VerifiedCiphertextCrc32c || plaintextCrc32c != result.PlaintextCrc32c {
		return fmt.Errorf("asymmetricDecrypt response corrupted in-transit, got %x, expected %x",
			plaintextCrc32c, result.PlaintextCrc32c)
	}

	if err := d.Set("plaintext", string(plaintext)); err != nil {
		return fmt.Errorf("error setting plaintext: %s", err)
	}

	d.SetId(fmt.Sprintf("%s:%x:%s", cryptoKeyVersion, ciphertextCRC32C, base64CipherText))
	return nil
}

func removeWhiteSpaceFromString(s string) string {
	whitespaceRegexp := regexp.MustCompile(`(?m)[\s]+`)
	return whitespaceRegexp.ReplaceAllString(s, "")
}

func validateBase64WithWhitespaces(i interface{}, val string) ([]string, []error) {
	_, err := base64.StdEncoding.DecodeString(removeWhiteSpaceFromString(i.(string)))
	if err != nil {
		return nil, []error{fmt.Errorf("could not decode %q as a valid base64 value. Please use the terraform base64 functions such as base64encode() or filebase64() to supply a valid base64 string", val)}
	}
	return nil, nil
}

func validateHexadecimalUint32(i interface{}, val string) ([]string, []error) {
	_, err := strconv.ParseUint(i.(string), 16, 32)
	if err != nil {
		return nil, []error{fmt.Errorf("could not decode %q as a unsigned 32 bit hexadecimal integer", val)}
	}
	return nil, nil
}
