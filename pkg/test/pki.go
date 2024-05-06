package test

import (
	"bytes"
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type CommonHarness struct {
	T *testing.T
}

func (h *CommonHarness) Init(t *testing.T) {
	h.T = t
}

func (h *CommonHarness) MustWriteFile(p string, data []byte, perm os.FileMode) {
	MustWriteFile(h.T, p, data, perm)
}

type PKIHarness struct {
	CommonHarness

	CertsDir string

	caPEM []byte
}

func (h *PKIHarness) CreateServerCertificates() {
	caCert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "kcctest-ca",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caPrivateKey, err := rsa.GenerateKey(cryptorand.Reader, 1024)
	if err != nil {
		h.T.Fatalf("generating rsa private key: %v", err)
	}
	caCertBytes, err := x509.CreateCertificate(cryptorand.Reader, caCert, caCert, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		h.T.Fatalf("generating ca certificate: %v", err)
	}
	h.caPEM = h.PEMEncodeCertificate(caCertBytes)

	serverCert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "kcctest-server",
		},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1)},
		DNSNames:    []string{"localhost"},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
	}
	serverPrivateKey, err := rsa.GenerateKey(cryptorand.Reader, 1024)
	if err != nil {
		h.T.Fatalf("generating rsa private key: %v", err)
	}
	serverCertBytes, err := x509.CreateCertificate(cryptorand.Reader, serverCert, caCert, &serverPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		h.T.Fatalf("generating server certificate: %v", err)
	}

	certsDir := h.T.TempDir()
	h.CertsDir = certsDir

	keyData := h.PEMEncodePrivateKey(serverPrivateKey)
	certData := h.PEMEncodeCertificate(serverCertBytes)

	// h.MustWriteFile(filepath.Join(certsDir, "ca.crt"),h.caPem, 0640)
	h.MustWriteFile(filepath.Join(certsDir, "server.crt"), certData, 0640)
	h.MustWriteFile(filepath.Join(certsDir, "server.key"), keyData, 0640)
}

func (h *PKIHarness) CABundle() []byte {

	return h.caPEM
}

func (h *PKIHarness) pemEncodeRaw(raw []byte, pemType string) []byte {
	var certPEM bytes.Buffer
	pem.Encode(&certPEM, &pem.Block{
		Type:  pemType,
		Bytes: raw,
	})
	return certPEM.Bytes()
}

func (h *PKIHarness) PEMEncodePrivateKey(key crypto.PrivateKey) []byte {
	rsaPrivateKey := key.(*rsa.PrivateKey)
	certBytes := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	pemType := "RSA PRIVATE KEY"
	return h.pemEncodeRaw(certBytes, pemType)
}

func (h *PKIHarness) PEMEncodeCertificate(certBytes []byte) []byte {
	pemType := "CERTIFICATE"
	return h.pemEncodeRaw(certBytes, pemType)
}
