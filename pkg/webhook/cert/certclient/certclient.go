// Copyright 2022 Google LLC
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

package certclient

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/provisioner"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/writer"

	admissionregistration "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	maxJitterFactor = 0.1
)

// default interval for checking cert is 90 days (~3 months)
var defaultCertRefreshInterval = 3 * 30 * 24 * time.Hour

// CertClient is responsible for installing webhook manifests and provisioning and rotating their certs.
type CertClient struct {
	webhookManifests []client.Object
	svc              *corev1.Service
	kubeClient       client.Client
	provisioner      *provisioner.Provisioner
}

type Options struct {
	WebhookManifests []client.Object
	Service          *corev1.Service
	KubeClient       client.Client
	CertDir          string
	CertWriter       writer.CertWriter
}

func New(opts Options) (*CertClient, error) {
	certWriter, err := certWriterFromOptsOrNew(opts)
	if err != nil {
		return nil, err
	}
	return &CertClient{
		webhookManifests: opts.WebhookManifests,
		svc:              opts.Service,
		provisioner: &provisioner.Provisioner{
			CertWriter: certWriter,
		},
		kubeClient: opts.KubeClient,
	}, nil
}

func certWriterFromOptsOrNew(opts Options) (writer.CertWriter, error) {
	if opts.CertWriter != nil {
		return opts.CertWriter, nil
	}
	certWriter, err := writer.NewFSCertWriter(
		writer.FSCertWriterOptions{
			Path: opts.CertDir,
		})
	if err != nil {
		return nil, fmt.Errorf("error creating FS cert writer: %w", err)
	}
	return certWriter, nil
}

func (c *CertClient) RefreshCertsAndInstall(ctx context.Context) error {
	_, err := c.provisioner.Provision(provisioner.Options{
		ClientConfig: &admissionregistration.WebhookClientConfig{
			CABundle: []byte{},
			Service: &admissionregistration.ServiceReference{
				Name:      c.svc.GetName(),
				Namespace: c.svc.GetNamespace(),
			},
		},
		Objects: c.webhookManifests,
	})
	if err != nil {
		return fmt.Errorf("error provisioning certs: %w", err)
	}
	objects := append([]client.Object{c.svc}, c.webhookManifests...)
	return batchCreateOrReplace(ctx, c.kubeClient, objects...)
}

func (c *CertClient) Start(ctx context.Context) error {
	timer := time.Tick(wait.Jitter(defaultCertRefreshInterval, maxJitterFactor))
	for {
		select {
		case <-timer:
			if err := c.RefreshCertsAndInstall(ctx); err != nil {
				return fmt.Errorf("error refreshing certs: %w", err)
			}

			// We force-exit to reload the certs.
			// We are missing logic to apply the new certs here;
			// it's also possible that another pod rotated this cert.
			// Because we want to move this to the operator anyway,
			// we simply exit here (relying on kubelet to restart us)
			// rather than trying to add update logic.
			// b/267353534
			klog.Warningf("forcing process exit after ~%v to reload webhook certificates", defaultCertRefreshInterval)
			os.Exit(1)

		case <-ctx.Done():
			return nil
		}
	}
}

type mutateFn func(current, desired *client.Object) error

var serviceFn = func(current, desired *client.Object) error {
	typedC := (*current).(*corev1.Service)
	typedD := (*desired).(*corev1.Service)
	typedC.Spec.Selector = typedD.Spec.Selector
	typedC.Spec.Ports = typedD.Spec.Ports
	return nil
}

var mutatingWebhookConfigFn = func(current, desired *client.Object) error {
	typedC := (*current).(*admissionregistration.MutatingWebhookConfiguration)
	typedD := (*desired).(*admissionregistration.MutatingWebhookConfiguration)
	typedC.Webhooks = typedD.Webhooks
	return nil
}

var validatingWebhookConfigFn = func(current, desired *client.Object) error {
	typedC := (*current).(*admissionregistration.ValidatingWebhookConfiguration)
	typedD := (*desired).(*admissionregistration.ValidatingWebhookConfiguration)
	typedC.Webhooks = typedD.Webhooks
	return nil
}

var genericFn = func(current, desired *client.Object) error {
	*current = *desired
	return nil
}

// createOrReplaceHelper creates the object if it doesn't exist;
// otherwise, it will replace it.
// When replacing, fn  should know how to preserve existing fields in the object GET from the APIServer.
func createOrReplaceHelper(ctx context.Context, c client.Client, obj client.Object, fn mutateFn) error {
	if obj == nil {
		return nil
	}

	objectKey := client.ObjectKeyFromObject(obj)

	// this cast is not safe per-se but is added to get around the transition from runtime.Object to client.Object
	existing, ok := obj.DeepCopyObject().(client.Object)
	if !ok {
		return fmt.Errorf("unable to cast deep copy %T to client.Object", obj)
	}

	// NOTE: create *before* get does not always work for services.
	// For services, if we are out of clusterIPs, then the create fails (with an error that is not NotFound)

	if err := c.Get(ctx, objectKey, existing); err != nil {
		if apierrors.IsNotFound(err) {
			return c.Create(ctx, obj)
		}
		return err
	}

	if err := fn(&existing, &obj); err != nil {
		return err
	}

	return c.Update(context.Background(), existing)
}

// createOrReplace creates the object if it doesn't exist;
// otherwise, it will replace it.
// When replacing, it knows how to preserve existing fields in the object GET from the APIServer.
// It currently only support MutatingWebhookConfiguration, ValidatingWebhookConfiguration and Service.
// For other kinds, it uses genericFn to replace the whole object.
func createOrReplace(ctx context.Context, c client.Client, obj client.Object) error {
	if obj == nil {
		return nil
	}
	switch obj.(type) {
	case *admissionregistration.MutatingWebhookConfiguration:
		return createOrReplaceHelper(ctx, c, obj, mutatingWebhookConfigFn)
	case *admissionregistration.ValidatingWebhookConfiguration:
		return createOrReplaceHelper(ctx, c, obj, validatingWebhookConfigFn)
	case *corev1.Service:
		return createOrReplaceHelper(ctx, c, obj, serviceFn)
	default:
		return createOrReplaceHelper(ctx, c, obj, genericFn)
	}
}

func batchCreateOrReplace(ctx context.Context, c client.Client, objs ...client.Object) error {
	for i := range objs {
		obj := objs[i]
		// TODO: We really should be using apply
		err := createOrReplace(ctx, c, obj)
		if err != nil {
			return fmt.Errorf("failed to apply object %T %v: %w", obj, obj.GetName(), err)
		}
	}
	return nil
}
