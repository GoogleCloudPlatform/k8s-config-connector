package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	var ( // TODO: convert to input parameters
		apiDirectory          = "/usr/local/google/home/jingyih/Project/kcc/scifi/src/github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
		newFieldFile          = "/usr/local/google/home/jingyih/Project/kcc/scifi/src/github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/newtypes.txt"
		parentMessageFullName = "google.monitoring.dashboard.v1.Dashboard"
	)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	// Define tools for listing, reading, and writing files
	fileTools := &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			{
				Name:        "listFiles",
				Description: "Lists all Go files in a directory.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"directory": {
							Type:        genai.TypeString,
							Description: "The directory to list files from.",
						},
					},
					Required: []string{"directory"},
				},
			},
			{
				Name:        "readFile",
				Description: "Reads the content of a Go file.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"filePath": {
							Type:        genai.TypeString,
							Description: "The path of the file to read.",
						},
					},
					Required: []string{"filePath"},
				},
			},
			// Note: Gemini does not seem to get the line number consistently. Every time
			// it says a different random number.
			// Use Go struct name instead.
			{
				Name:        "replaceGoStruct",
				Description: "replace the content of a Go struct",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"filePath": {
							Type:        genai.TypeString,
							Description: "The path of the file to be modified.",
						},
						"goStructName": {
							Type:        genai.TypeString,
							Description: "The name of Go struct whose content to be replaced.",
						},
						"originalContent": {
							Type:        genai.TypeString,
							Description: "The original content inside the Go struct",
						},
						"content": {
							Type:        genai.TypeString,
							Description: "The content of the Go struct to be replaced with.",
						},
					},
					Required: []string{"filePath", "goStructName", "content"},
				},
			},
			/*
				// Note: Gemini cannot return a super large response in function call.
				{
					Name:        "writeFile",
					Description: "Writes modified content back to a Go file.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"filePath": {
								Type:        genai.TypeString,
								Description: "The path of the file to write.",
							},
							"content": {
								Type:        genai.TypeString,
								Description: "The content to write to the file.",
							},
						},
						Required: []string{"filePath", "content"},
					},
				},
			*/
		},
	}

	model.Tools = []*genai.Tool{fileTools}

	/*
		prompt := fmt.Sprintf(
			`
			I have Go files under the directory %s. I also have Go field written in file %s.
			Could you first read content from files in the directory, find the Go struct that is commented
			with "+kcc:proto=%s" (no following suffix), and insert the Go field in that struct?
			You can call updateFile as many times to insert one line of content at a time.
			After you are done, just respond with stop.
			Please try to make minimum changes if possible. Only insert the Go field to the target file.
			Do not return text response until you are done.
			`, apiDirectory, newFieldFile, parentMessageFullName)

		fmt.Println(prompt)
	*/

	// Start new chat session
	session := model.StartChat()
	session.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(fmt.Sprintf(`
						I have some Go structs written in Go files under a directory.
						I also have a Go field written in a file.
						Could you:
						1) read content from all files in the directory
						2) find the Go struct that has comment "+kcc:proto=%s" with no following suffix
						3) insert the Go field into the found Go struct, write it back
						You only need to tell me what does the updated Go struct look like. No need to send the entire file.
						Do not return text response until you are finished all steps. Then respond with text "stop".
					`, parentMessageFullName)),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(fmt.Sprintf("The directory is %s.", apiDirectory)),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(fmt.Sprintf(`
				The Go file is written in file %s.
				`, newFieldFile)),
			},
			Role: "user",
		},
	}

	resp, err := session.SendMessage(ctx, genai.Text("you can start now."))
	if err != nil {
		log.Fatalf("Error receiving message: %v\n", err)
	}

	for {
		// Process the response
		fmt.Println("============")
		time.Sleep(1 * time.Second) // do not throttle the API

		part := resp.Candidates[0].Content.Parts[0]
		switch part.(type) {
		case genai.FunctionCall:
			fmt.Printf("Received function call: %#v\n", part.(genai.FunctionCall))
			funcall, ok := part.(genai.FunctionCall)
			if !ok {
				log.Fatalf("Expected type FunctionCall, got %T", part)
			}

			switch funcall.Name {
			case "readFile":
				filePath := funcall.Args["filePath"].(string)
				content, err := readFile(filePath)
				if err != nil {
					log.Fatalf("error reading file: %v\n", err)
				}
				apiResult := map[string]any{
					"filePath": filePath,
					"content":  content,
				}

				// Send the file content back to Gemini
				resp, err = session.SendMessage(ctx, genai.FunctionResponse{
					Name:     "readFile",
					Response: apiResult,
				})
				if err != nil {
					log.Fatalf("error sending message: %v\n", err)
				}

			case "writeFile":
				filePath := funcall.Args["filePath"].(string)
				modifiedContent := funcall.Args["content"].(string)

				// Write the modified content back to the file
				err := writeFile(filePath, modifiedContent)
				if err != nil {
					log.Fatalf("eeror writing file: %v\n", err)
				}

				resp, err = session.SendMessage(ctx, genai.Text("file written successfully."))
				if err != nil {
					log.Fatalf("Error receiving message: %v\n", err)
				}
				os.Exit(0)

			case "replaceGoStruct":
				filePath := funcall.Args["filePath"].(string)
				goStructName := funcall.Args["goStructName"].(string)
				content := funcall.Args["content"].(string)

				// Update the file with the new content
				err := updateFile(filePath, goStructName, content)
				if err != nil {
					log.Fatalf("error updating file: %v\n", err)
				}
				resp, err = session.SendMessage(ctx, genai.Text("file updated successfully."))
				if err != nil {
					log.Fatalf("Error receiving message: %v\n", err)
				}

			case "listFiles":
				directory := funcall.Args["directory"].(string)
				files, err := listFiles(directory)
				if err != nil {
					log.Fatalf("error listing files: %v\n", err)
				}
				anySlice := make([]any, len(files))
				for _, f := range files {
					fmt.Printf("found file %#v\n", f)
					anySlice = append(anySlice, f)
				}
				apiResult := map[string]any{
					"directory": directory,
					"files":     anySlice,
				}

				// Send the list of files back to Gemini
				resp, err = session.SendMessage(ctx, genai.FunctionResponse{
					Name:     "listFiles",
					Response: apiResult,
				})
				if err != nil {
					log.Fatalf("error sending message: %v\n", err)
				}

			default:
				log.Fatalf("unknown function call: %s", funcall.Name)
			}
		case genai.Text:
			fmt.Println(fmt.Sprintf("received text: %s\n", part.(genai.Text)))
			if strings.Contains(string(part.(genai.Text)), "stop") {
				os.Exit(0)
			}
			os.Exit(1)
		}

	}
}

func listFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func updateFile(filePath, goStructName, content string) error {
	content = renderEscapedString(content)

	fmt.Printf("updating file %s, goStruct %s with the following content:\n", filePath, goStructName)
	fmt.Println(content)

	// TODO: update file

	return nil

	/*file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
	}

	if lineNumber < 1 || lineNumber > len(lines)+1 {
		return fmt.Errorf("line number %d out of bounds", lineNumber)
	}

	newLines := strings.Split(newText, "\n")
	for i := range newLines {
		newLines[i] += "\n" // Add the newline character back to each line
	}

	lines = append(lines[:lineNumber-1], append(newLines, lines[lineNumber-1:]...)...)

	return os.WriteFile(filePath, []byte(strings.Join(lines, "")), 0644)*/
}

func renderEscapedString(escapedStr string) string {
	replacer := strings.NewReplacer(
		"\\t", "\t",
		"\\n", "\n",
		"\\\"", "\"",
		"\\\\", "\\",
	)
	return replacer.Replace(escapedStr)
}

func writeFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, content)
	return nil
}
