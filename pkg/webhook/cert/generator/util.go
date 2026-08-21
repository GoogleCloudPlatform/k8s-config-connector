/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generator

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"k8s.io/klog/v2"
)

// ValidCACert think cert and key are valid if they meet the following requirements:
// - key and cert are valid pair
// - caCert is the root ca of cert
// - cert is for dnsName
// - cert won't expire before time
func ValidCACert(key, cert, caCert []byte, dnsName string, time time.Time) bool {
	ctx := context.TODO()
	log := klog.FromContext(ctx)
	if len(key) == 0 || len(cert) == 0 || len(caCert) == 0 {
		log.Error(fmt.Errorf("empty key, cert or caCert"), "key, cert or caCert is empty")
		return false
	}
	// Verify key and cert are valid pair
	_, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Error(err, "failed to parse key pair")
		return false
	}

	// Verify cert is valid for at least 1 year.
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caCert) {
		log.Error(fmt.Errorf("failed to parse caCert"), "failed to append caCert to pool")
		return false
	}
	block, _ := pem.Decode([]byte(cert))
	if block == nil {
		log.Error(fmt.Errorf("failed to decode cert"), "failed to decode cert")
		return false
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Error(err, "failed to parse cert")
		return false
	}
	ops := x509.VerifyOptions{
		DNSName:     dnsName,
		Roots:       pool,
		CurrentTime: time,
	}
	_, err = c.Verify(ops)
	if err != nil {
		log.Error(err, "failed to verify cert")
	}
	return err == nil
}
