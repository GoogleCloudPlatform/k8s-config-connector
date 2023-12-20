// Copyright 2023 Google LLC
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

package runtime

import (
	"context"
	"fmt"
	"log"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsub/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"cloud.google.com/go/pubsub"
)

// Create a struct type named PubSubConverter which implements the OnePlatformConverter interface from converter.go

type PubSubConverter struct {
}

// GetResource will convert obj to the Go type in pubsubtopic_types.go
func (c *PubSubConverter) GetResource(ctx context.Context, obj *unstructured.Unstructured) (u *unstructured.Unstructured, err error) {
	strongTypeObj := &v1beta1.PubSubTopic{}

	// Convert the unstructured object to a strongly typed object
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, strongTypeObj)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	projectID := strongTypeObj.ObjectMeta.Annotations["cnrm.cloud.google.com/project-id"]

	client, err := pubsub.NewClient(ctx, projectID)

	topic, _ := getGCPPubSubTopic(ctx, client, *strongTypeObj.Spec.ResourceID)

	if topic == nil {
		return nil, nil
	}

	sdkObj, _ := topic.Config(ctx)

	krmObj := sdkToKRMObj(&sdkObj, strongTypeObj)

	us, err := toUnstructured(krmObj)

	return us, nil
}

// CreateResource will convert obj to the Go type in pubsubtopic_types.go and then call createGCPPubSubTopics() method to create a pubsub topic resource using PubSub Go client SDK
func (c *PubSubConverter) CreateResource(ctx context.Context, obj *unstructured.Unstructured) (u *unstructured.Unstructured, err error) {

	strongTypeObj := &v1beta1.PubSubTopic{}

	// Convert the unstructured object to a strongly typed object
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, strongTypeObj)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	projectID := strongTypeObj.ObjectMeta.Annotations["cnrm.cloud.google.com/project-id"]

	client, err := pubsub.NewClient(ctx, projectID)

	sdkObj := krmToSdkObj(strongTypeObj)

	topic, err := client.CreateTopicWithConfig(ctx, *strongTypeObj.Spec.ResourceID, sdkObj)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
		return nil, err
	}

	config, _ := topic.Config(ctx)

	krmObj := sdkToKRMObj(&config, strongTypeObj)

	us, err := toUnstructured(krmObj)

	return us, nil
}

// UpdateResource will convert obj to the Go type in pubsubtopic_types.go
// convert the Go type to pubsub.TopicConfig and then update a pubsub topic resource using PubSub Go client SDK
func (c *PubSubConverter) UpdateResource(ctx context.Context, obj *unstructured.Unstructured) (u *unstructured.Unstructured, err error) {

	strongTypeObj := &v1beta1.PubSubTopic{}

	// Convert the unstructured object to a strongly typed object
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, strongTypeObj)
	if err != nil {
		log.Fatal(err)
	}
	projectID := strongTypeObj.ObjectMeta.Annotations["cnrm.cloud.google.com/project-id"]

	client, err := pubsub.NewClient(ctx, projectID)

	sdkObj := krmToSdkObj(strongTypeObj)

	topicID := strongTypeObj.Spec.ResourceID
	topic := client.Topic(*topicID)

	// Convert TopicConfig to TopicConfigToUpdate
	topicConfigToUpdate := convertToTopicConfigToUpdate(*sdkObj)

	updatedConfig, err := topic.Update(ctx, topicConfigToUpdate)
	if err != nil {
		log.Fatalf("Failed to update topic: %v", err)
		return nil, err
	}

	log.Printf("Updated topic: %+v\n", updatedConfig)

	krmObj := sdkToKRMObj(&updatedConfig, strongTypeObj)

	us, err := toUnstructured(krmObj)

	return us, err
}

// DeleteResource will extract the resourceID from the obj and then delete a pubsub topic resource using PubSub Go client SDK
func (c *PubSubConverter) DeleteResource(ctx context.Context, obj *unstructured.Unstructured) error {
	strongTypeObj := &v1beta1.PubSubTopic{}

	// Convert the unstructured object to a strongly typed object
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, strongTypeObj)
	if err != nil {
		log.Fatal(err)
		return err
	}
	projectID := strongTypeObj.ObjectMeta.Annotations["cnrm.cloud.google.com/project-id"]

	client, err := pubsub.NewClient(ctx, projectID)

	topicID := strongTypeObj.Spec.ResourceID

	// use client to delete topic
	topic := client.Topic(*topicID)
	err = topic.Delete(ctx)
	if err != nil {
		log.Fatalf("Failed to delete topic: %v", err)
		return err
	}

	return nil
}

// GetDiff will convert both local and remote to strong type and then compare them to see if there are any diffs
// An important assumption is if certain field from local object is nil, it means the KRM user has no opinion on the field, so we don't
// compare it with the remote object.
func (c *PubSubConverter) GetDiff(ctx context.Context, local, remote *unstructured.Unstructured) (bool, error) {
	// convert local to strong type
	strongTypeLocalObj := &v1beta1.PubSubTopic{}

	// Convert the unstructured object to a strongly typed object
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(local.Object, strongTypeLocalObj)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// convert remote to strong type
	strongTypeRemoteObj := &v1beta1.PubSubTopic{}

	// Convert the unstructured object to a strongly typed object
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(remote.Object, strongTypeRemoteObj)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// Compare local and remote objects on the fields we are interested in
	if strongTypeLocalObj.Spec.KmsKeyRef != nil && strongTypeRemoteObj.Spec.KmsKeyRef != nil {
		if strongTypeLocalObj.Spec.KmsKeyRef.External != strongTypeRemoteObj.Spec.KmsKeyRef.External {
			return true, nil
		}
	} else if strongTypeLocalObj.Spec.KmsKeyRef != nil && strongTypeRemoteObj.Spec.KmsKeyRef == nil {
		return true, nil
	} else if strongTypeLocalObj.Spec.KmsKeyRef == nil && strongTypeRemoteObj.Spec.KmsKeyRef != nil {
		return true, nil
	}

	if strongTypeLocalObj.Spec.MessageRetentionDuration != strongTypeRemoteObj.Spec.MessageRetentionDuration {
		return true, nil
	}

	// compare Spec.SchemaSettings for local and remote objects
	if strongTypeLocalObj.Spec.SchemaSettings != nil && strongTypeRemoteObj.Spec.SchemaSettings != nil {
		if strongTypeLocalObj.Spec.SchemaSettings.Encoding != strongTypeRemoteObj.Spec.SchemaSettings.Encoding {
			return true, nil
		}

		if strongTypeLocalObj.Spec.SchemaSettings.SchemaRef.External != strongTypeRemoteObj.Spec.SchemaSettings.SchemaRef.External {
			return true, nil
		}
	} else if strongTypeLocalObj.Spec.SchemaSettings == nil && strongTypeRemoteObj.Spec.SchemaSettings != nil {
		return true, nil
	} else if strongTypeLocalObj.Spec.SchemaSettings != nil && strongTypeRemoteObj.Spec.SchemaSettings == nil {
		return true, nil
	}

	// compare the string slices under MessageStoragePolicy
	if strongTypeLocalObj.Spec.MessageStoragePolicy != nil {
		if strongTypeRemoteObj.Spec.MessageStoragePolicy != nil {
			if len(strongTypeLocalObj.Spec.MessageStoragePolicy.AllowedPersistenceRegions) != len(strongTypeRemoteObj.Spec.MessageStoragePolicy.AllowedPersistenceRegions) {
				return true, nil
			}
			for i, region := range strongTypeLocalObj.Spec.MessageStoragePolicy.AllowedPersistenceRegions {
				if region != strongTypeRemoteObj.Spec.MessageStoragePolicy.AllowedPersistenceRegions[i] {
					return true, nil
				}
			}
		} else if strongTypeRemoteObj.Spec.MessageStoragePolicy == nil {
			return true, nil
		}
	}

	return false, nil
}

// Unfortunately we need a different struct TopicConfigToUpdate to update Topic instead of TopicConfig
func convertToTopicConfigToUpdate(config pubsub.TopicConfig) pubsub.TopicConfigToUpdate {
	updateConfig := pubsub.TopicConfigToUpdate{}

	if config.Labels != nil {
		updateConfig.Labels = make(map[string]string)
		for key, value := range config.Labels {
			updateConfig.Labels[key] = value
		}
	}

	if config.RetentionDuration != 0 {
		updateConfig.RetentionDuration = &config.RetentionDuration
	}

	// Add other fields as necessary
	updateConfig.MessageStoragePolicy = &config.MessageStoragePolicy
	updateConfig.SchemaSettings = config.SchemaSettings

	return updateConfig
}

func toUnstructured(obj runtime.Object) (*unstructured.Unstructured, error) {
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: unstructuredObj}, nil
}

// sdkToKRMObj() takes object of type pubsub.TopicConfig and convert it to type v1beta1.PubSubTopic and return it
func sdkToKRMObj(sdkObj *pubsub.TopicConfig, originalObject *v1beta1.PubSubTopic) *v1beta1.PubSubTopic {
	// Create a new v1beta1.PubSubTopic object
	krmObj := &v1beta1.PubSubTopic{}
	krmObj.TypeMeta = originalObject.TypeMeta
	krmObj.ObjectMeta = originalObject.ObjectMeta

	// Iterate through the fields of pubsub.TopicConfig and map the corresponding
	// fields from pubsub.TopicConfig object to a new v1beta1.PubSubTopic object
	if sdkObj.KMSKeyName != "" {
		krmObj.Spec.KmsKeyRef.External = sdkObj.KMSKeyName
	}
	krmObj.Spec.MessageStoragePolicy = &v1beta1.TopicMessageStoragePolicy{
		AllowedPersistenceRegions: sdkObj.MessageStoragePolicy.AllowedPersistenceRegions,
	}
	if sdkObj.SchemaSettings != nil {
		encoding := schemaEncodingToString(sdkObj.SchemaSettings.Encoding)
		krmObj.Spec.SchemaSettings.Encoding = &encoding
		krmObj.Spec.SchemaSettings.SchemaRef.External = sdkObj.SchemaSettings.Schema
	}
	if sdkObj.RetentionDuration != nil {
		str := fmt.Sprintf("%v", sdkObj.RetentionDuration) // Convert it to string
		krmObj.Spec.MessageRetentionDuration = &str
	}

	return krmObj
}

// krmToSdkObj() takes object of type v1beta1.PubSubTopic and convert it to type pubsub.Topic, during
// the conversion, it iterates through the fields of v1beta1.PubSubTopic and maps the corresponding
// fields from v1beta1.PubSubTopic object to a new pubsub.Topic object
func krmToSdkObj(krmObj *v1beta1.PubSubTopic) *pubsub.TopicConfig {
	// Create a new pubsub.Topic object
	sdkObj := &pubsub.TopicConfig{}

	// Iterate through the fields of v1beta1.PubSubTopic and map the corresponding
	// fields from v1beta1.PubSubTopic object to a new pubsub.TopicConfig object
	if krmObj.Spec.KmsKeyRef != nil {
		sdkObj.KMSKeyName = krmObj.Spec.KmsKeyRef.External
	}
	if krmObj.Spec.MessageStoragePolicy != nil {
		sdkObj.MessageStoragePolicy = pubsub.MessageStoragePolicy{
			AllowedPersistenceRegions: krmObj.Spec.MessageStoragePolicy.AllowedPersistenceRegions,
		}
	}
	if krmObj.Spec.SchemaSettings != nil {
		sdkObj.SchemaSettings.Encoding, _ = stringToSchemaEncoding(*krmObj.Spec.SchemaSettings.Encoding)
		sdkObj.SchemaSettings.Schema = krmObj.Spec.SchemaSettings.SchemaRef.External
	}
	if krmObj.Spec.MessageRetentionDuration != nil {
		duration, _ := time.ParseDuration(*krmObj.Spec.MessageRetentionDuration)
		sdkObj.RetentionDuration = duration
	}

	return sdkObj
}

// schemaEncodingToString converts a pubsub.SchemaEncoding value to a string.
func schemaEncodingToString(encoding pubsub.SchemaEncoding) string {
	switch encoding {
	case pubsub.EncodingUnspecified:
		return "ENCODING_UNSPECIFIED"
	case pubsub.EncodingJSON:
		return "JSON"
	case pubsub.EncodingBinary:
		return "BINARY"
	default:
		return "UNKNOWN"
	}
}

// stringToSchemaEncoding converts a string to the corresponding pubsub.SchemaEncoding value.
func stringToSchemaEncoding(encoding string) (pubsub.SchemaEncoding, error) {
	switch encoding {
	case "ENCODING_UNSPECIFIED":
		return pubsub.EncodingUnspecified, nil
	case "JSON":
		return pubsub.EncodingJSON, nil
	case "BINARY":
		return pubsub.EncodingBinary, nil
	default:
		return 0, fmt.Errorf("unknown encoding: %s", encoding)
	}
}

// getGCPPubSubTopics() method will get a pubsub topic resource using PubSub Go client SDK. It will return a pubsub topic resource or nil if the topic does not exist.
func getGCPPubSubTopic(ctx context.Context, client *pubsub.Client, topicID string) (*pubsub.Topic, error) {
	// Get a topic
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if topic exists: %w", err)
	}
	if !exists {
		return nil, nil
	}
	return topic, nil
}
