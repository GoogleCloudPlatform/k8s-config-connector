package toolbot

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"
)

type protoService struct {
	FilePath   string
	Definition []string
}

type protoMessage struct {
	FilePath   string
	Definition []string
}

type EnhanceWithProtoDefinition struct {
	protoDirectory string
	messages       map[string]*protoMessage
}

func NewEnhanceWithProtoDefinition(protoDirectory string) (*EnhanceWithProtoDefinition, error) {
	x := &EnhanceWithProtoDefinition{
		protoDirectory: protoDirectory,
		messages:       make(map[string]*protoMessage),
	}
	if err := x.findProtoMessages(); err != nil {
		return nil, err
	}
	return x, nil
}

var _ Enhancer = &EnhanceWithProtoDefinition{}

func (x *EnhanceWithProtoDefinition) EnhanceDataPoint(ctx context.Context, p *DataPoint) error {
	service := p.Input["proto.service"]
	if service != "" {
		protoService, err := x.getProtoForService(ctx, service)
		if err != nil {
			return fmt.Errorf("getting proto for service %q: %w", service, err)
		}
		p.SetInput("proto.service.definition", strings.Join(protoService.Definition, "\n"))
	}

	message := p.Input["proto.message"]
	if message != "" {
		protoMessage := x.messages[message]
		if protoMessage != nil {
			p.SetInput("proto.message.definition", strings.Join(protoMessage.Definition, "\n"))
		} else {
			klog.Infof("unable to find proto message %q", message)
		}
	}

	return nil
}

func (x *EnhanceWithProtoDefinition) getProtoForService(ctx context.Context, serviceName string) (*protoService, error) {
	var matches []*protoService
	if err := filepath.WalkDir(x.protoDirectory, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(p) {
		case ".proto":
			// OK
		default:
			return nil
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		r := bytes.NewReader(b)
		br := bufio.NewReader(r)

		packageName := ""

		for {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("scanning file %q: %w", p, err)
			}
			line = strings.TrimSuffix(line, "\n")

			tokens := strings.Fields(line)

			if len(tokens) >= 2 && tokens[0] == "package" {
				packageName = strings.TrimSuffix(tokens[1], ";")
			}

			if len(tokens) >= 2 && tokens[0] == "service" {
				found := packageName + "." + tokens[1]

				if found != serviceName {
					continue
				}

				match := &protoService{FilePath: p}
				indent := 0
				for {
					match.Definition = append(match.Definition, line)
					for _, r := range line {
						if r == '{' {
							indent++
						}
						if r == '}' {
							indent--
						}
					}
					if indent == 0 {
						break
					}
					line, err = br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return fmt.Errorf("scanning file %q: %w", p, err)
					}
					line = strings.TrimSuffix(line, "\n")
				}
				matches = append(matches, match)
			}
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("walking directory tree: %w", err)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("service %q not found", serviceName)
	}
	if len(matches) > 1 {
		return nil, fmt.Errorf("found multiple services with name %q", serviceName)
	}
	return matches[0], nil
}

func (x *EnhanceWithProtoDefinition) findProtoMessages() error {

	if err := filepath.WalkDir(x.protoDirectory, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(p) {
		case ".proto":
			// OK
		default:
			return nil
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		r := bytes.NewReader(b)
		br := bufio.NewReader(r)

		packageName := ""

		for {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("scanning file %q: %w", p, err)
			}
			line = strings.TrimSuffix(line, "\n")

			tokens := strings.Fields(line)

			if len(tokens) >= 2 && tokens[0] == "package" {
				packageName = strings.TrimSuffix(tokens[1], ";")
			}

			if len(tokens) >= 2 && tokens[0] == "message" {
				messageName := packageName + "." + tokens[1]

				message := &protoMessage{FilePath: p}
				indent := 0
				for {
					message.Definition = append(message.Definition, line)
					for _, r := range line {
						if r == '{' {
							indent++
						}
						if r == '}' {
							indent--
						}
					}
					if indent == 0 {
						break
					}
					line, err = br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return fmt.Errorf("scanning file %q: %w", p, err)
					}
					line = strings.TrimSuffix(line, "\n")
				}
				x.messages[messageName] = message
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory tree: %w", err)
	}

	return nil
}
