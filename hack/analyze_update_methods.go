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

type RPCInfo struct {
	MethodName  string
	RequestMsg  string
	Description string
	FilePath    string
}

type MessageInfo struct {
	Fields   map[string]string
	FilePath string
}

func toSnakeCase(s string) string {
	var res string
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				res += "_"
			}
			res += strings.ToLower(string(r))
		} else {
			res += string(r)
		}
	}
	return res
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	gcpFile := filepath.Join(rootDir, "gcp_mutability.json")
	content, err := ioutil.ReadFile(gcpFile)
	if err != nil {
		fmt.Printf("Error reading gcp_mutability.json: %v\n", err)
		os.Exit(1)
	}

	var mutabilityData map[string]map[string]interface{}
	if err := json.Unmarshal(content, &mutabilityData); err != nil {
		fmt.Printf("Error unmarshaling gcp_mutability.json: %v\n", err)
		os.Exit(1)
	}

	protoFiles := []string{}
	// Search in googleapis
	filepath.Walk(filepath.Join(rootDir, ".build/third_party/googleapis/google"), func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})
	// Search in mockgcp/apis (for generated protos like DNS)
	filepath.Walk(filepath.Join(rootDir, "mockgcp/apis/google"), func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	// Map of message name to its info (multiple files might have same message name, we handle it during lookup)
	messageFields := make(map[string][]MessageInfo)
	rpcs := []RPCInfo{}

	messageStartRegex := regexp.MustCompile(`^message\s+([A-Za-z0-9_]+)\s*\{`)
	fieldRegex := regexp.MustCompile(`^\s*(?:repeated\s+|optional\s+|required\s+)?([a-zA-Z0-9_.<>, ]+)\s+([a-zA-Z0-9_]+)\s*=\s*\d+`)
	rpcRegex := regexp.MustCompile(`^rpc\s+([A-Za-z0-9_]+)\s*\(\s*([A-Za-z0-9_.]+)\s*\)`)

	for _, file := range protoFiles {
		contentBytes, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		lines := strings.Split(string(contentBytes), "\n")
		
		var currentMessage string
		var depth int
		var commentBlock []string
		
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			
			if strings.HasPrefix(trimmed, "//") {
				commentBlock = append(commentBlock, strings.TrimSpace(strings.TrimPrefix(trimmed, "//")))
			} else if strings.HasPrefix(trimmed, "rpc ") {
				matches := rpcRegex.FindStringSubmatch(trimmed)
				if len(matches) > 2 {
					rpcs = append(rpcs, RPCInfo{
						MethodName:  matches[1],
						RequestMsg:  matches[2],
						Description: strings.Join(commentBlock, " "),
						FilePath:    file,
					})
				}
				commentBlock = nil
			} else if trimmed == "" {
				// Keep collecting comments
			} else {
				commentBlock = nil
			}

			if strings.Contains(line, "{") {
				depth++
			}
			if strings.Contains(line, "}") {
				depth--
				if depth == 0 {
					currentMessage = ""
				}
			}

			if currentMessage == "" {
				matches := messageStartRegex.FindStringSubmatch(trimmed)
				if len(matches) > 1 {
					currentMessage = matches[1]
					messageFields[currentMessage] = append(messageFields[currentMessage], MessageInfo{
						Fields:   make(map[string]string),
						FilePath: file,
					})
					depth = 1 
				}
			} else if depth == 1 {
				matches := fieldRegex.FindStringSubmatch(trimmed)
				if len(matches) > 2 {
					fieldType := matches[1]
					fieldName := matches[2]
					if strings.HasPrefix(fieldType, "map<") {
						fieldType = "map"
					}
					// Add to the last message info added for this message name
					idx := len(messageFields[currentMessage]) - 1
					messageFields[currentMessage][idx].Fields[fieldName] = fieldType
				}
			}
		}
	}

	// Recursive function to find a field in the message graph, now filtered by file path relevance
	var findField func(msgName string, targetField string, visited map[string]bool, relevantPathKeyword string) (bool, string)
	findField = func(msgName string, targetField string, visited map[string]bool, relevantPathKeyword string) (bool, string) {
		key := msgName + ":" + relevantPathKeyword
		if visited[key] {
			return false, ""
		}
		visited[key] = true
		
		allInfos, ok := messageFields[msgName]
		if !ok {
			return false, ""
		}
		
		// Filter infos by relevance
		var infos []MessageInfo
		for _, info := range allInfos {
			if strings.Contains(strings.ToLower(info.FilePath), relevantPathKeyword) {
				infos = append(infos, info)
			}
		}
		// Fallback to all if none match (risky but might be needed for common shared messages)
		if len(infos) == 0 {
			infos = allInfos
		}

		for _, info := range infos {
			fields := info.Fields
			// 1. Check top-level first
			for f := range fields {
				if f == targetField || strings.ReplaceAll(f, "_", "") == strings.ReplaceAll(targetField, "_", "") {
					return true, f
				}
			}
			
			for f := range fields {
				if strings.HasPrefix(targetField, f+"_") || strings.HasSuffix(targetField, "_"+f) || (f == "labels" && strings.Contains(targetField, "label")) {
					return true, f
				}
			}
			
			// 2. Traverse nested messages
			// Sort fields to ensure deterministic traversal
			var sortedFields []string
			for f := range fields {
				sortedFields = append(sortedFields, f)
			}
			sort.Strings(sortedFields)

			for _, f := range sortedFields {
				fType := fields[f]
				if fType == "string" || fType == "int32" || fType == "int64" || fType == "uint32" || fType == "uint64" || fType == "bool" || fType == "float" || fType == "double" || fType == "bytes" || fType == "map" || strings.HasPrefix(fType, "google.protobuf.") {
					continue
				}
				
				parts := strings.Split(fType, ".")
				cleanType := parts[len(parts)-1]
				
				found, foundName := findField(cleanType, targetField, visited, relevantPathKeyword)
				if found {
					return true, f + "." + foundName
				}
			}
		}
		
		return false, ""
	}

	// Special cases for KCC kinds that don't map cleanly to a Proto Message of the same name
	specialCases := map[string]string{
		// Faceted / Sub-field resources: KCC resource represents a specific field or facet of a larger GCP resource.
		"ComputeProjectMetadata":               "Project", // Metadata is a field on the Project resource
		"BigtableGCPolicy":                     "Table",   // GC policy is a field in the Table resource

		// Legacy / Unified resources: KCC resource maps to a modern, unified GCP API message.
		"ComputeHTTPHealthCheck":               "HealthCheck", // Mapped to the unified HealthCheck API
		"ComputeHTTPSHealthCheck":              "HealthCheck", // Mapped to the unified HealthCheck API

		// Subcomponent resources: Typically managed via specialized parent/child methods or distinct but related messages.
		"StorageBucketAccessControl":           "BucketAccessControl",
		"StorageDefaultObjectAccessControl":    "DefaultObjectAccessControl",

		// Naming differences: Direct mappings with different names (often [ServiceName][ResourceName] in KCC).
		"SecretManagerSecretVersion":           "SecretVersion",
		"LoggingLogSink":                       "LogSink",
		"LoggingLogBucket":                     "LogBucket",
		"LoggingLogView":                       "LogView",
		"LoggingLogExclusion":                  "LogExclusion",
		"DNSManagedZone":                       "ManagedZone",
		"DNSRecordSet":                         "ResourceRecordSet",
		"IdentityPlatformOAuthIDPConfig":       "OAuthIdpConfig",
		"IdentityPlatformTenantOAuthIDPConfig": "OAuthIdpConfig",
	}

	totalWithoutUpdate := 0

	for kind, data := range mutabilityData {
		// Apply special case overrides
		if protoOverride, ok := specialCases[kind]; ok {
			data["proto_message"] = protoOverride
		}
		
		protoMessage := data["proto_message"].(string)
		
		expectedType := ""
		if et, ok := data["expected_type"].(string); ok {
			expectedType = et
		}

		relevantPathKeyword := ""
		if expectedType != "" {
			// Extract service prefix. e.g., "compute.googleapis.com/Instance" -> "compute" or "dataproc/Cluster" -> "dataproc"
			firstPart := strings.Split(expectedType, "/")[0]
			relevantPathKeyword = strings.Split(firstPart, ".")[0]
		}
		if relevantPathKeyword == "" {
			// Fallback: try to extract from the KCC Group or Kind
			if g, ok := data["group"].(string); ok && g != "" {
				relevantPathKeyword = strings.Split(g, ".")[0]
			} else {
				relevantPathKeyword = strings.ToLower(strings.Split(kind, "Compute")[0]) 
			}
		}

		// Let's assume we can get it from the Kind for common ones
		if strings.HasPrefix(kind, "Compute") {
			relevantPathKeyword = "compute"
		} else if strings.HasPrefix(kind, "Logging") {
			relevantPathKeyword = "logging"
		} else if strings.HasPrefix(kind, "Bigtable") {
			relevantPathKeyword = "bigtable"
		} else if strings.HasPrefix(kind, "Spanner") {
			relevantPathKeyword = "spanner"
		} else if strings.HasPrefix(kind, "DNS") {
			relevantPathKeyword = "dns"
		} else if strings.HasPrefix(kind, "IAM") {
			relevantPathKeyword = "iam"
		} else if strings.HasPrefix(kind, "Storage") {
			relevantPathKeyword = "storage"
		} else if strings.HasPrefix(kind, "BigQuery") {
			relevantPathKeyword = "bigquery"
		} else if strings.HasPrefix(kind, "Redis") {
			relevantPathKeyword = "redis"
		} else if strings.HasPrefix(kind, "KMS") {
			relevantPathKeyword = "kms"
		} else if strings.HasPrefix(kind, "SQL") {
			relevantPathKeyword = "sql"
		} else if strings.HasPrefix(kind, "CloudFunctions") {
			relevantPathKeyword = "functions"
		}

		// Some normalization for keywords
		if relevantPathKeyword == "accesscontextmanager" {
			relevantPathKeyword = "access_context_manager"
		} else if relevantPathKeyword == "cloudresourcemanager" {
			relevantPathKeyword = "resourcemanager"
		}

		updateReqMsgLower := strings.ToLower("Update" + protoMessage + "Request")
		patchReqMsgLower := strings.ToLower("Patch" + protoMessage + "Request")
		
		var fallbackUpdateLower, fallbackPatchLower string
		prefixes := []string{"Log", "Managed", "Compute", "Bucket", "FlexTemplate", "ResourceRecord", "OAuthIdp"}
		for _, p := range prefixes {
			if strings.HasPrefix(protoMessage, p) && len(protoMessage) > len(p) {
				short := protoMessage[len(p):]
				fallbackUpdateLower = strings.ToLower("Update" + short + "Request")
				fallbackPatchLower = strings.ToLower("Patch" + short + "Request")
				break
			}
		}

		hasUpdate := false
		requestMsg := ""
		
		for _, rpc := range rpcs {
			// Ensure RPC is from a relevant file
			if relevantPathKeyword != "" && !strings.Contains(strings.ToLower(rpc.FilePath), relevantPathKeyword) {
				continue
			}

			reqLower := strings.ToLower(rpc.RequestMsg)
			if reqLower == updateReqMsgLower || reqLower == patchReqMsgLower || 
			   (fallbackUpdateLower != "" && (reqLower == fallbackUpdateLower || reqLower == "update"+fallbackUpdateLower)) || 
			   (fallbackPatchLower != "" && (reqLower == fallbackPatchLower || reqLower == "patch"+fallbackPatchLower)) {
				hasUpdate = true
				requestMsg = rpc.RequestMsg
				break
			}
		}
		
		data["has_update_method"] = hasUpdate
		if hasUpdate {
			data["update_request_message"] = requestMsg
		} else {
			data["update_request_message"] = "" // Clear it!
			totalWithoutUpdate++
		}

		// Find setters (case-insensitive for protoMessage)
		setters := make(map[string]map[string]interface{})
		protoMsgLower := strings.ToLower(protoMessage)
		for _, rpc := range rpcs {
			// Ensure RPC is from a relevant file
			if relevantPathKeyword != "" && !strings.Contains(strings.ToLower(rpc.FilePath), relevantPathKeyword) {
				continue
			}

			msgLower := strings.ToLower(rpc.RequestMsg)
			method := rpc.MethodName

			if method == "SetIamPolicy" {
				continue
			}

			var inferredField string
			if strings.HasPrefix(msgLower, "set") && strings.HasSuffix(msgLower, protoMsgLower+"request") {
				originalFieldStart := 3
				originalFieldEnd := len(rpc.RequestMsg) - len(protoMessage) - 7
				if originalFieldEnd > originalFieldStart {
					inferredField = rpc.RequestMsg[originalFieldStart:originalFieldEnd]
				}
			} else if strings.HasPrefix(msgLower, protoMsgLower+"set") && strings.HasSuffix(msgLower, "request") {
				originalFieldStart := len(protoMessage) + 3
				originalFieldEnd := len(rpc.RequestMsg) - 7
				if originalFieldEnd > originalFieldStart {
					inferredField = rpc.RequestMsg[originalFieldStart:originalFieldEnd]
				}
			} else if strings.HasPrefix(msgLower, "set") && strings.HasSuffix(msgLower, "projectrequest") && protoMessage == "Project" {
				originalFieldStart := 3
				originalFieldEnd := len(rpc.RequestMsg) - 14
				if originalFieldEnd > originalFieldStart {
					inferredField = rpc.RequestMsg[originalFieldStart:originalFieldEnd]
				}
			}
			
			if inferredField != "" {
				// Strip common scope suffixes from the inferred field name
				// e.g. SecurityPolicyRegion -> SecurityPolicy
				cleanInferredField := inferredField
				scopes := []string{"Region", "Global", "Zonal", "Project"}
				for _, s := range scopes {
					if strings.HasSuffix(inferredField, s) && len(inferredField) > len(s) {
						cleanInferredField = inferredField[:len(inferredField)-len(s)]
						break
					}
				}

				snakeField := toSnakeCase(cleanInferredField)
				fieldExists, matchedField := findField(protoMessage, snakeField, make(map[string]bool), relevantPathKeyword)
				
				// If not found with clean name, try original name
				if !fieldExists && cleanInferredField != inferredField {
					snakeField = toSnakeCase(inferredField)
					fieldExists, matchedField = findField(protoMessage, snakeField, make(map[string]bool), relevantPathKeyword)
				}
				
				setters[method] = map[string]interface{}{
					"description":               rpc.Description,
					"inferred_field":            snakeField,
					"matched_field_in_resource": matchedField,
					"is_field_in_resource":      fieldExists,
				}
			}
		}
		data["setter_methods"] = setters
	}
	
	fmt.Printf("Found %d resources WITHOUT an Update or Patch method\n", totalWithoutUpdate)

	outFile := filepath.Join(rootDir, "gcp_mutability.json")
	outData, err := json.MarshalIndent(mutabilityData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		os.Exit(1)
	}

	ioutil.WriteFile(outFile, outData, 0644)
}