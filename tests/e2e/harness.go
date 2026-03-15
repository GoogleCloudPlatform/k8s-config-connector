package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

type Harness struct {
	*create.Harness
}

type HarnessInterface interface {
	GetClient() client.Client

	Context() context.Context

	Fatalf(format string, args ...interface{})
}

func NewHarness(ctx context.Context, t *testing.T) *Harness {
	h := create.NewHarness(ctx, t)
	return &Harness{
		Harness: h,
	}
}

func (h *Harness) ApplyObject(obj *unstructured.Unstructured) {
	if err := h.GetClient().Patch(h.Ctx, obj, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
		h.Fatalf("error applying resource: %v", err)
	}
}

func (h *Harness) ExportObject(obj *unstructured.Unstructured) *unstructured.Unstructured {
	return exportResourceAsUnstructured(h.Harness, obj)
}

func (h *Harness) WaitForReady(timeout time.Duration, objects ...*unstructured.Unstructured) {
	create.WaitForReady(h.Harness, timeout, objects...)
}

func (h *Harness) DeleteResources(objects []*unstructured.Unstructured) {
	create.DeleteResources(h.Harness, create.CreateDeleteTestOptions{Create: objects})
}
func (h *Harness) AsYAML(obj *unstructured.Unstructured) []byte {
	yamlBytes, err := yaml.Marshal(obj.Object)
	if err != nil {
		h.Fatalf("error marshalling object to YAML: %v", err)
	}
	return yamlBytes
}
