package configconnector

import (
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"k8s.io/client-go/util/cert"
	"k8s.io/client-go/util/keyutil"
)

const (
	rsaKeySize = 2048

	certificateDuration = time.Hour * 24 * 365 * 10

	maxSkew = time.Hour * -24
)

func newRSAPrivateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(cryptorand.Reader, rsaKeySize)
}

// newCACertificate creates a self signed CA certificate
func newCACertificate() (*rsa.PrivateKey, *x509.Certificate, error) {
	caKey, err := newRSAPrivateKey()
	if err != nil {
		return nil, nil, err
	}
	subject := pkix.Name{CommonName: "webhook-cert-ca"}
	now := time.Now()
	tmpl := x509.Certificate{
		SerialNumber:          new(big.Int).SetInt64(1),
		Subject:               subject,
		NotBefore:             now.Add(maxSkew).UTC(),
		NotAfter:              now.Add(certificateDuration).UTC(),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDERBytes, err := x509.CreateCertificate(cryptorand.Reader, &tmpl, &tmpl, caKey.Public(), caKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create CA certificate: %w", err)
	}
	caCert, err := x509.ParseCertificate(certDERBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse generated CA certificate: %w", err)
	}
	return caKey, caCert, nil
}

// newWebhookCertificate creates a signed CA certificate / key and server key / cert pair.
func newWebhookCertificate(dnsName string) (*CertificateData, error) {
	caKey, caCert, err := newCACertificate()
	if err != nil {
		return nil, err
	}

	serial, err := cryptorand.Int(cryptorand.Reader, new(big.Int).SetInt64(math.MaxInt64))
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %w", err)
	}
	if dnsName == "" {
		return nil, errors.New("must specify a CommonName")
	}

	serverKey, err := newRSAPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to create the private key: %w", err)
	}

	now := time.Now()

	serverCertTemplate := x509.Certificate{
		DNSNames:     []string{dnsName},
		SerialNumber: serial,
		NotBefore:    now.Add(maxSkew).UTC(),
		NotAfter:     now.Add(certificateDuration).UTC(),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	serverCertDER, err := x509.CreateCertificate(cryptorand.Reader, &serverCertTemplate, caCert, serverKey.Public(), caKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %w", err)
	}
	serverCert, err := x509.ParseCertificate(serverCertDER)
	if err != nil {
		return nil, fmt.Errorf("failed to parse generated certificate: %w", err)
	}

	return &CertificateData{
		CACert:     encodeCertPEM(caCert),
		CAKey:      encodePrivateKeyPEM(caKey),
		ServerCert: encodeCertPEM(serverCert),
		ServerKey:  encodePrivateKeyPEM(serverKey),
	}, nil
}

func encodePrivateKeyPEM(key *rsa.PrivateKey) []byte {
	block := pem.Block{
		Type:  keyutil.RSAPrivateKeyBlockType,
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	return pem.EncodeToMemory(&block)
}

func encodeCertPEM(c *x509.Certificate) []byte {
	block := pem.Block{
		Type:  cert.CertificateBlockType,
		Bytes: c.Raw,
	}
	return pem.EncodeToMemory(&block)
}
