package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type ResourceInfo struct {
	ControllerType     string `json:"controller_type"`
	ResourceName       string `json:"resource_name"`
	ServiceHostName    string `json:"service_hostname,omitempty"`
	APIEndpoint        string `json:"api_endpoint,omitempty"`
	OfficialRESTAPIURL string `json:"official_rest_api_url,omitempty"`
	Kind               string `json:"kind"`
	Group              string `json:"group"`
}

// Map resource kind to its probable Protobuf message name
func getProtoMessageName(kind, group string) string {
	servicePrefix := strings.Split(group, ".")[0]
	// Special cases or direct matches can be added here
	
	lowerKind := strings.ToLower(kind)
	if strings.HasPrefix(lowerKind, servicePrefix) {
		// Strip the service prefix (case-insensitive)
		return kind[len(servicePrefix):]
	}
	return kind
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	resourcesFile := filepath.Join(rootDir, "resources.json")
	content, err := ioutil.ReadFile(resourcesFile)
	if err != nil {
		fmt.Printf("Error reading resources.json: %v\n", err)
		os.Exit(1)
	}

	var resources map[string]ResourceInfo
	if err := json.Unmarshal(content, &resources); err != nil {
		fmt.Printf("Error unmarshaling resources.json: %v\n", err)
		os.Exit(1)
	}

	// 1. Collect all .proto files in googleapis
	protoFiles := []string{}
	err = filepath.Walk(filepath.Join(rootDir, ".build/third_party/googleapis/google"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // ignore errors
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking googleapis: %v\n", err)
		os.Exit(1)
	}

	// This maps a unique resource type (e.g. alloydb.googleapis.com/Cluster) to a slice of its immutable fields
	messageToImmutable := make(map[string][]string)
	
	// This regex matches "message Name {" and captures the name
	messageStartRegex := regexp.MustCompile(`(?m)^message\s+([A-Za-z0-9_]+)\s*\{`)
	
	// This regex matches the google.api.resource type
	typeRegex := regexp.MustCompile(`type:\s*"([^"]+)"`)

	// This regex matches a field definition that contains IMMUTABLE in its options block
	fieldRegex := regexp.MustCompile(`(?s)([a-zA-Z0-9_]+)\s*=\s*\d+\s*\[([^\]]*IMMUTABLE[^\]]*)\];`)

	totalParsedFields := 0

	for _, file := range protoFiles {
		contentBytes, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		content := string(contentBytes)
		
		// Find all messages and their start indices
		messageMatches := messageStartRegex.FindAllStringSubmatchIndex(content, -1)
		
		for i, match := range messageMatches {
			messageName := content[match[2]:match[3]]
			startIndex := match[1]
			
			// Determine end index (start of next message or end of file)
			endIndex := len(content)
			if i+1 < len(messageMatches) {
				endIndex = messageMatches[i+1][0]
			}
			
			messageBody := content[startIndex:endIndex]
			
			// Find google.api.resource type
			resourceType := messageName // fallback
			typeMatches := typeRegex.FindStringSubmatch(messageBody)
			if len(typeMatches) > 1 {
				resourceType = typeMatches[1]
			} else {
				// We can try to guess the service from the file path, but for now we skip 
				// or keep it as the message name as a fallback.
			}
			
			// Find fields with IMMUTABLE in this message body
			fieldMatches := fieldRegex.FindAllStringSubmatch(messageBody, -1)
			for _, fm := range fieldMatches {
				fieldName := fm[1]
				messageToImmutable[resourceType] = append(messageToImmutable[resourceType], fieldName)
				totalParsedFields++
			}
		}
	}
	fmt.Printf("Parsed %d immutable fields from proto files\n", totalParsedFields)

	// 2. Map KCC resources to actual mutability
	actualMutability := make(map[string]map[string]interface{})
	
	for kind, info := range resources {
		protoMessageName := getProtoMessageName(kind, info.Group)
		
		// Construct the expected resource type
		expectedType := protoMessageName
		if info.ServiceHostName != "" {
			expectedType = info.ServiceHostName + "/" + protoMessageName
		}
		
		fields := []string{}
		if fieldsFound, ok := messageToImmutable[expectedType]; ok {
			fields = fieldsFound
		} else if fieldsFound, ok := messageToImmutable[protoMessageName]; ok {
			// Fallback to just the message name if full type not found
			fields = fieldsFound
		}
		
		// Deduplicate fields (since multiple proto versions might exist and we read all of them)
		fieldMap := make(map[string]bool)
		uniqueFields := []string{}
		for _, f := range fields {
			if !fieldMap[f] {
				fieldMap[f] = true
				uniqueFields = append(uniqueFields, f)
			}
		}
		
		sort.Strings(uniqueFields)
		
		actualMutability[kind] = map[string]interface{}{
			"proto_message": protoMessageName,
			"expected_type": expectedType,
			"immutable_fields": uniqueFields,
		}
	}

	outFile := filepath.Join(rootDir, "gcp_mutability.json")
	outData, err := json.MarshalIndent(actualMutability, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outFile, outData, 0644); err != nil {
		fmt.Printf("Error writing gcp_mutability.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully extracted actual mutability for %d resources to gcp_mutability.json\n", len(actualMutability))
}