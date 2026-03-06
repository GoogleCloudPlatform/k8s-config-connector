// Copyright 2026 Google LLC
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

package migrate

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/kubecli"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

var (
	clusterOptions kubecli.ClusterOptions
	outputFile     string

	migrateGKEAddonCmd = &cobra.Command{
		Use:   "gke-addon",
		Short: "Migrate from GKE Addon to Manual Installation",
		Long:  `Helper tool to migrate Config Connector from a GKE Addon to a manual installation in the same cluster.`,
	}

	migrateGKEAddonPrepareCmd = &cobra.Command{
		Use:   "prepare",
		Short: "Prepare for migration: backup resources, pause reconciliation, and remove finalizers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPrepare(cmd.Context())
		},
	}

	migrateGKEAddonFinishCmd = &cobra.Command{
		Use:   "finish",
		Short: "Finish migration: re-apply resources and resume reconciliation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFinish(cmd.Context())
		},
	}
)

func init() {
	migrateGKEAddonCmd.AddCommand(migrateGKEAddonPrepareCmd)
	migrateGKEAddonCmd.AddCommand(migrateGKEAddonFinishCmd)

	clusterOptions.AddFlags(migrateGKEAddonPrepareCmd)
	migrateGKEAddonPrepareCmd.Flags().StringVarP(&outputFile, "output", "o", "kcc-backup.yaml", "Path to the backup file")

	clusterOptions.AddFlags(migrateGKEAddonFinishCmd)
	migrateGKEAddonFinishCmd.Flags().StringVarP(&outputFile, "input", "i", "kcc-backup.yaml", "Path to the backup file")
}

func runPrepare(ctx context.Context) error {
	kubeClient, err := kubecli.NewClient(ctx, clusterOptions)
	if err != nil {
		return err
	}

	fmt.Println("1. Discovering all Config Connector resources...")
	resources, err := getAllKCCResources(ctx, kubeClient)
	if err != nil {
		return fmt.Errorf("error getting KCC resources: %w", err)
	}
	fmt.Printf("Found %d resources.\n", len(resources))

	fmt.Printf("2. Backing up resources to %s...\n", outputFile)
	if err := backupResources(resources, outputFile); err != nil {
		return fmt.Errorf("error backing up resources: %w", err)
	}

	fmt.Println("3. Pausing reconciliation...")
	if err := pauseReconciliation(ctx, kubeClient); err != nil {
		return fmt.Errorf("error pausing reconciliation: %w", err)
	}

	fmt.Println("4. Removing finalizers from resources to decouple them from GCP...")
	if err := removeFinalizers(ctx, kubeClient, resources); err != nil {
		return fmt.Errorf("error removing finalizers: %w", err)
	}

	fmt.Println("\nPreparation complete!")
	fmt.Println("Next steps:")
	fmt.Println("1. Disable the Config Connector addon in your GKE cluster:")
	fmt.Println("   gcloud container clusters update <cluster-name> --update-addons ConfigConnector=DISABLED")
	fmt.Println("2. Install the Config Connector operator manually:")
	fmt.Println("   https://cloud.google.com/config-connector/docs/how-to/install-manually")
	fmt.Println("3. Run the finish command to resume reconciliation:")
	fmt.Println("   config-connector migrate gke-addon finish")

	return nil
}

func runFinish(ctx context.Context) error {
	kubeClient, err := kubecli.NewClient(ctx, clusterOptions)
	if err != nil {
		return err
	}

	fmt.Printf("1. Loading backed up resources from %s...\n", outputFile)
	resources, err := loadResources(outputFile)
	if err != nil {
		return fmt.Errorf("error loading resources: %w", err)
	}

	fmt.Println("2. Re-applying resources...")
	if err := applyResources(ctx, kubeClient, resources); err != nil {
		return fmt.Errorf("error applying resources: %w", err)
	}

	fmt.Println("3. Resuming reconciliation...")
	if err := resumeReconciliation(ctx, kubeClient); err != nil {
		return fmt.Errorf("error resuming reconciliation: %w", err)
	}

	fmt.Println("\nMigration complete!")
	return nil
}

func getAllKCCResources(ctx context.Context, kubeClient *kubecli.Client) ([]unstructured.Unstructured, error) {
	crdList := &unstructured.UnstructuredList{}
	crdList.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	})
	if err := kubeClient.List(ctx, crdList); err != nil {
		return nil, fmt.Errorf("error listing CRDs: %w", err)
	}

	var kccResources []unstructured.Unstructured
	for _, crd := range crdList.Items {
		group, found, err := unstructured.NestedString(crd.Object, "spec", "group")
		if err != nil || !found {
			continue
		}
		if !strings.HasSuffix(group, ".cnrm.cloud.google.com") {
			continue
		}
		kind, found, err := unstructured.NestedString(crd.Object, "spec", "names", "kind")
		if err != nil || !found {
			continue
		}

		versions, found, err := unstructured.NestedSlice(crd.Object, "spec", "versions")
		if err != nil || !found || len(versions) == 0 {
			continue
		}

		var version string
		for _, vIntf := range versions {
			v, ok := vIntf.(map[string]interface{})
			if !ok {
				continue
			}
			name, _ := v["name"].(string)
			served, _ := v["served"].(bool)
			storage, _ := v["storage"].(bool)
			if served {
				version = name
				if storage {
					break
				}
			}
		}
		if version == "" {
			v0, _ := versions[0].(map[string]interface{})
			version, _ = v0["name"].(string)
		}

		gvk := schema.GroupVersionKind{
			Group:   group,
			Version: version,
			Kind:    kind,
		}
		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)
		if err := kubeClient.List(ctx, list); err != nil {
			fmt.Printf("Warning: failed to list %s: %v\n", gvk.Kind, err)
			continue
		}
		kccResources = append(kccResources, list.Items...)
	}
	return kccResources, nil
}

func backupResources(resources []unstructured.Unstructured, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for i, res := range resources {
		// Clean up metadata before backup. Use a copy to avoid modifying the original.
		resCopy := res.DeepCopy()
		cleanResourceForBackup(resCopy)

		y, err := yaml.Marshal(resCopy.Object)
		if err != nil {
			return err
		}
		if i > 0 {
			if _, err := f.WriteString("---\n"); err != nil {
				return err
			}
		}
		if _, err := f.Write(y); err != nil {
			return err
		}
	}
	return nil
}

func cleanResourceForBackup(res *unstructured.Unstructured) {
	unstructured.RemoveNestedField(res.Object, "metadata", "resourceVersion")
	unstructured.RemoveNestedField(res.Object, "metadata", "uid")
	unstructured.RemoveNestedField(res.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(res.Object, "metadata", "managedFields")
	unstructured.RemoveNestedField(res.Object, "metadata", "deletionTimestamp")
	unstructured.RemoveNestedField(res.Object, "metadata", "deletionGracePeriodSeconds")
	unstructured.RemoveNestedField(res.Object, "status")
}

func pauseReconciliation(ctx context.Context, kubeClient *kubecli.Client) error {
	gvkList := []schema.GroupVersionKind{
		{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnector"},
		{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnectorContext"},
	}

	for _, gvk := range gvkList {
		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)
		if err := kubeClient.List(ctx, list); err != nil {
			fmt.Printf("Warning: failed to list %s: %v\n", gvk.Kind, err)
			continue
		}

		for i := range list.Items {
			res := &list.Items[i]
			if err := unstructured.SetNestedField(res.Object, "Paused", "spec", "actuationMode"); err != nil {
				return err
			}
			if err := kubeClient.Update(ctx, res); err != nil {
				return fmt.Errorf("error updating %s %s/%s: %w", res.GetKind(), res.GetNamespace(), res.GetName(), err)
			}
			fmt.Printf("Paused %s %s/%s\n", res.GetKind(), res.GetNamespace(), res.GetName())
		}
	}
	return nil
}

func removeFinalizers(ctx context.Context, kubeClient *kubecli.Client, resources []unstructured.Unstructured) error {
	for i := range resources {
		res := &resources[i]
		if len(res.GetFinalizers()) == 0 {
			continue
		}
		res.SetFinalizers(nil)
		if err := kubeClient.Update(ctx, res); err != nil {
			// Don't fail the whole process if one finalizer removal fails, but warn the user.
			fmt.Printf("Warning: failed to remove finalizers from %s %s/%s: %v\n", res.GetKind(), res.GetNamespace(), res.GetName(), err)
		}
	}
	return nil
}

func loadResources(path string) ([]unstructured.Unstructured, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var resources []unstructured.Unstructured
	docs := strings.Split(string(data), "\n---\n")
	for _, doc := range docs {
		if strings.TrimSpace(doc) == "" {
			continue
		}
		res := unstructured.Unstructured{}
		if err := yaml.Unmarshal([]byte(doc), &res.Object); err != nil {
			return nil, err
		}
		resources = append(resources, res)
	}
	return resources, nil
}

func applyResources(ctx context.Context, kubeClient *kubecli.Client, resources []unstructured.Unstructured) error {
	for _, res := range resources {
		// Use Patch to create or update
		data, err := yaml.Marshal(res.Object)
		if err != nil {
			return err
		}
		if err := kubeClient.Patch(ctx, &res, client.RawPatch(types.ApplyPatchType, data), client.FieldOwner("kcc-migration-tool"), client.ForceOwnership); err != nil {
			fmt.Printf("Warning: failed to apply %s %s/%s: %v\n", res.GetKind(), res.GetNamespace(), res.GetName(), err)
		}
	}
	return nil
}

func resumeReconciliation(ctx context.Context, kubeClient *kubecli.Client) error {
	gvkList := []schema.GroupVersionKind{
		{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnector"},
		{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnectorContext"},
	}

	for _, gvk := range gvkList {
		list := &unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)
		if err := kubeClient.List(ctx, list); err != nil {
			fmt.Printf("Warning: failed to list %s: %v\n", gvk.Kind, err)
			continue
		}

		for i := range list.Items {
			res := &list.Items[i]
			if err := unstructured.SetNestedField(res.Object, "Reconciling", "spec", "actuationMode"); err != nil {
				return err
			}
			if err := kubeClient.Update(ctx, res); err != nil {
				return fmt.Errorf("error updating %s %s/%s: %w", res.GetKind(), res.GetNamespace(), res.GetName(), err)
			}
			fmt.Printf("Resumed %s %s/%s\n", res.GetKind(), res.GetNamespace(), res.GetName())
		}
	}
	return nil
}
