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

package protoapi

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/scanner"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func ApplyOverlay(fds *descriptorpb.FileDescriptorSet, overlayPath string) error {
	f, err := os.Open(overlayPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return applyOverlayFromReader(fds, f)
}

func applyOverlayFromReader(fds *descriptorpb.FileDescriptorSet, r io.Reader) error {
	var s scanner.Scanner
	s.Init(r)
	s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanStrings | scanner.SkipComments

	l := &lexer{scanner: s}
	l.next()

	var currentPackage string
	fieldBehaviors := make(map[string][]annotations.FieldBehavior)

	if err := parseFile(l, &currentPackage, fieldBehaviors); err != nil {
		return fmt.Errorf("parsing overlay: %w", err)
	}

	if len(fieldBehaviors) == 0 {
		return nil
	}

	for _, file := range fds.File {
		applyToMessages(file.GetPackage(), file.MessageType, fieldBehaviors)
	}

	return nil
}

type lexer struct {
	scanner scanner.Scanner
	tok     rune
}

func (l *lexer) next() {
	l.tok = l.scanner.Scan()
}

func (l *lexer) text() string {
	return l.scanner.TokenText()
}

func parseFile(l *lexer, currentPackage *string, behaviors map[string][]annotations.FieldBehavior) error {
	for l.tok != scanner.EOF {
		if l.tok == scanner.Ident && l.text() == "package" {
			l.next()
			pkg, err := parseIdentPath(l)
			if err != nil {
				return err
			}
			*currentPackage = pkg
			if l.tok != ';' {
				return fmt.Errorf("expected ';' at %s", l.scanner.Position)
			}
			l.next()
		} else if l.tok == scanner.Ident && l.text() == "message" {
			l.next()
			if l.tok != scanner.Ident {
				return fmt.Errorf("expected message name at %s", l.scanner.Position)
			}
			msgName := l.text()
			l.next()

			if err := parseMessage(l, []string{msgName}, behaviors, currentPackage); err != nil {
				return err
			}
		} else {
			// skip other top-level elements (e.g. syntax, import, option)
			if l.tok == scanner.Ident && (l.text() == "syntax" || l.text() == "import" || l.text() == "option") {
				for l.tok != scanner.EOF && l.tok != ';' {
					l.next()
				}
				if l.tok == ';' {
					l.next()
				}
			} else {
				l.next()
			}
		}
	}
	return nil
}

func parseIdentPath(l *lexer) (string, error) {
	var sb strings.Builder
	for {
		if l.tok != scanner.Ident {
			return "", fmt.Errorf("expected identifier at %s", l.scanner.Position)
		}
		sb.WriteString(l.text())
		l.next()
		if l.tok == '.' {
			sb.WriteRune('.')
			l.next()
		} else {
			break
		}
	}
	return sb.String(), nil
}

func parseMessage(l *lexer, stack []string, behaviors map[string][]annotations.FieldBehavior, currentPackage *string) error {
	if l.tok != '{' {
		return fmt.Errorf("expected '{' at %s", l.scanner.Position)
	}
	l.next()

	for l.tok != scanner.EOF && l.tok != '}' {
		if l.tok == scanner.Ident && l.text() == "message" {
			l.next()
			if l.tok != scanner.Ident {
				return fmt.Errorf("expected message name at %s", l.scanner.Position)
			}
			msgName := l.text()
			l.next()

			newStack := append([]string{}, stack...)
			newStack = append(newStack, msgName)
			if err := parseMessage(l, newStack, behaviors, currentPackage); err != nil {
				return err
			}
		} else if l.tok == scanner.Ident && l.text() == "option" {
			for l.tok != scanner.EOF && l.tok != ';' {
				l.next()
			}
			if l.tok == ';' {
				l.next()
			}
		} else if l.tok == scanner.Ident && (l.text() == "oneof" || l.text() == "enum") {
			l.next()
			if l.tok == scanner.Ident {
				l.next()
			}
			if l.tok == '{' {
				skipBlock(l)
			}
		} else if l.tok == scanner.Ident {
			// A field declaration
			var fieldName string
			for l.tok != scanner.EOF && l.tok != '=' && l.tok != ';' && l.tok != '{' {
				if l.tok == scanner.Ident {
					fieldName = l.text()
				}
				l.next()
			}

			if l.tok == '{' {
				// E.g. nested extension or unexpected block
				skipBlock(l)
			} else if l.tok == '=' {
				l.next() // consume '='

				for l.tok != scanner.EOF && l.tok != ';' && l.tok != '[' {
					l.next()
				}

				if l.tok == '[' {
					l.next() // consume '['
					for l.tok != scanner.EOF && l.tok != ']' {
						if l.tok == '(' {
							l.next()
							optName, _ := parseIdentPath(l)
							if l.tok == ')' {
								l.next()
							}
							if l.tok == '=' {
								l.next()
								if l.tok == scanner.Ident {
									optVal := l.text()
									l.next()
									if optName == "google.api.field_behavior" {
										behavior, err := parseFieldBehavior(optVal)
										if err != nil {
											return err
										}

										fullName := *currentPackage
										for _, s := range stack {
											if fullName != "" {
												fullName += "."
											}
											fullName += s
										}
										if fullName != "" {
											fullName += "."
										}
										fullName += fieldName

										behaviors[fullName] = append(behaviors[fullName], behavior)
									}
								}
							}
						} else {
							l.next()
						}
					}
					if l.tok == ']' {
						l.next()
					}
				}

				if l.tok == ';' {
					l.next()
				} else {
					return fmt.Errorf("expected ';' after field at %s", l.scanner.Position)
				}
			} else if l.tok == ';' {
				l.next()
			}
		} else {
			l.next() // ignore unexpected tokens inside message
		}
	}

	if l.tok == '}' {
		l.next()
	} else {
		return fmt.Errorf("expected '}' at %s", l.scanner.Position)
	}

	return nil
}

func skipBlock(l *lexer) {
	if l.tok != '{' {
		return
	}
	l.next()
	depth := 1
	for l.tok != scanner.EOF && depth > 0 {
		if l.tok == '{' {
			depth++
		} else if l.tok == '}' {
			depth--
		}
		if depth > 0 {
			l.next()
		}
	}
}

func parseFieldBehavior(val string) (annotations.FieldBehavior, error) {
	switch val {
	case "OUTPUT_ONLY":
		return annotations.FieldBehavior_OUTPUT_ONLY, nil
	case "REQUIRED":
		return annotations.FieldBehavior_REQUIRED, nil
	case "OPTIONAL":
		return annotations.FieldBehavior_OPTIONAL, nil
	case "IMMUTABLE":
		return annotations.FieldBehavior_IMMUTABLE, nil
	case "INPUT_ONLY":
		return annotations.FieldBehavior_INPUT_ONLY, nil
	default:
		return 0, fmt.Errorf("unknown field behavior: %q", val)
	}
}

func applyToMessages(pkg string, msgs []*descriptorpb.DescriptorProto, fieldBehaviors map[string][]annotations.FieldBehavior) {
	for _, msg := range msgs {
		msgName := msg.GetName()
		fullMsgName := msgName
		if pkg != "" {
			fullMsgName = pkg + "." + msgName
		}

		for _, field := range msg.Field {
			fullName := fullMsgName + "." + field.GetName()
			if behaviors, ok := fieldBehaviors[fullName]; ok {
				if field.Options == nil {
					field.Options = &descriptorpb.FieldOptions{}
				}
				proto.SetExtension(field.Options, annotations.E_FieldBehavior, behaviors)
			}
		}

		if len(msg.NestedType) > 0 {
			applyToMessages(fullMsgName, msg.NestedType, fieldBehaviors)
		}
	}
}
