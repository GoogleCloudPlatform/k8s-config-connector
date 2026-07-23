package utils

import (
	"context"
	"fmt"
	"net"
	"strings"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	dnsDomain = "cluster.local"
	dnsIP     = "10.96.0.10"
)

// getKubernetesService fetches the Kubernetes Service
func getKubernetesService(ctx context.Context, c client.Client) (*corev1.Service, error) {
	kubernetesService := &corev1.Service{}
	id := client.ObjectKey{Namespace: metav1.NamespaceDefault, Name: "kubernetes"}

	// Get the Kubernetes Service
	err := c.Get(ctx, id, kubernetesService)

	return kubernetesService, err
}

// FindDNSClusterIP tries to find the Cluster IP to be used by the DNS service
// It is usually the 10th address to the Kubernetes Service Cluster IP
// If the Kubernetes Service Cluster IP is not found, we default it to be "10.96.0.10"
func FindDNSClusterIP(ctx context.Context, c client.Client) (string, error) {
	kubernetesService, err := getKubernetesService(ctx, c)
	if err != nil && !apierrors.IsNotFound(err) {
		return "", err
	}

	if apierrors.IsNotFound(err) {
		// If it cannot determine the Cluster IP, we default it to "10.96.0.10"
		return dnsIP, nil
	}

	ip := net.ParseIP(kubernetesService.Spec.ClusterIP)
	if ip == nil {
		return "", fmt.Errorf("cannot parse kubernetes ClusterIP %q", kubernetesService.Spec.ClusterIP)
	}

	// The kubernetes Service ClusterIP is the 1st IP in the Service Subnet.
	// Increment the right-most byte by 9 to get to the 10th address, canonically used for DNS.
	// This works for both IPV4, IPV6, and 16-byte IPV4 addresses.
	ip[len(ip)-1] += 9

	result := ip.String()
	klog.Infof("determined ClusterIP for cluster should be %q", result)
	return result, nil
}

// GetDNSDomain returns Kubernetes DNS cluster domain
// If it cannot determine the domain, we default it to "cluster.local"
// TODO (rajansandeep): find a better way to implement this?
func GetDNSDomain() string {
	svc := "kubernetes.default.svc"

	cname, err := net.LookupCNAME(svc)
	if err != nil {
		// If it cannot determine the domain, we default it to "cluster.local"
		klog.Infof("could not determine the domain, the DNS Domain for the cluster will default to %q", dnsDomain)
		return dnsDomain
	}

	domain := strings.TrimPrefix(cname, svc)
	domain = strings.TrimPrefix(domain, ".")
	domain = strings.TrimSuffix(domain, ".")

	klog.Infof("determined DNS Domain for DNS should be %q", domain)

	return domain
}
