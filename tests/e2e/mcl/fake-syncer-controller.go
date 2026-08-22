package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var sourceContext, targetContext, kubeconfig string
	flag.StringVar(&sourceContext, "source-context", "", "Kubeconfig context for source cluster")
	flag.StringVar(&targetContext, "target-context", "", "Kubeconfig context for target cluster")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig")
	flag.Parse()

	if sourceContext == "" || targetContext == "" || kubeconfig == "" {
		fmt.Println("All flags --source-context, --target-context, and --kubeconfig must be set")
		os.Exit(1)
	}

	srcCfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{CurrentContext: sourceContext}).ClientConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building source config: %v\n", err)
		os.Exit(1)
	}
	tgtCfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{CurrentContext: targetContext}).ClientConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building target config: %v\n", err)
		os.Exit(1)
	}

	srcClient := dynamic.NewForConfigOrDie(srcCfg)
	tgtClient := dynamic.NewForConfigOrDie(tgtCfg)

	gvr := schema.GroupVersionResource{
		Group:    "essentialcontacts.cnrm.cloud.google.com",
		Version:  "v1beta1",
		Resource: "essentialcontactscontacts",
	}

	fmt.Printf("Fake Syncer starting. Syncing %s from %s to %s\n", gvr.Resource, sourceContext, targetContext)

	for {
		list, err := srcClient.Resource(gvr).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing resources in source: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, item := range list.Items {
			name := item.GetName()
			namespace := item.GetNamespace()

			// Simple copy: remove fields that should not be synced
			newItem := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": item.GetAPIVersion(),
					"kind":       item.GetKind(),
					"metadata": map[string]interface{}{
						"name":      name,
						"namespace": namespace,
					},
					"spec": item.Object["spec"],
				},
			}
			// Copy status if present
			if status, ok := item.Object["status"]; ok {
				newItem.Object["status"] = status
			}

			created, err := tgtClient.Resource(gvr).Namespace(namespace).Create(context.Background(), newItem, metav1.CreateOptions{})
			if err != nil {
				if errors.IsAlreadyExists(err) {
					// Update if needed
					existing, _ := tgtClient.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
					newItem.SetResourceVersion(existing.GetResourceVersion())
					created, err = tgtClient.Resource(gvr).Namespace(namespace).Update(context.Background(), newItem, metav1.UpdateOptions{})
					if err != nil {
						fmt.Printf("Error updating %s in target: %v\n", name, err)
						continue
					}
					fmt.Printf("Updated %s in target cluster\n", name)
				} else {
					fmt.Printf("Error creating %s in target: %v\n", name, err)
					continue
				}
			} else {
				fmt.Printf("Synced %s to target cluster\n", name)
			}

			// Explicitly sync status using the status subresource
			if status, ok := newItem.Object["status"]; ok {
				statusUpdate := created.DeepCopy()
				statusUpdate.Object["status"] = status
				_, err = tgtClient.Resource(gvr).Namespace(namespace).UpdateStatus(context.Background(), statusUpdate, metav1.UpdateOptions{})
				if err != nil {
					fmt.Printf("Error syncing status for %s: %v\n", name, err)
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}
