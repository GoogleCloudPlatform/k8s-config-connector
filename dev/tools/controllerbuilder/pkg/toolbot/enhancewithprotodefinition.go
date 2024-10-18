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
)


type protoService struct {
	FilePath   string
	Definition []string
}

type EnhanceWithProtoDefinition struct {
	ProtoDirectory string
}

var _ Enhancer = &EnhanceWithProtoDefinition{}

func (x *EnhanceWithProtoDefinition) EnhanceDataPoint(ctx context.Context, p *DataPoint) error {
	service := p.Input["proto.service"]
	if service == "" {
		return nil
	}

	protoService, err := x.getProtoForService(ctx, service)
	if err != nil {
		return fmt.Errorf("getting proto for service %q: %w", service, err)
	}
	p.SetInput("proto.service.definition", strings.Join(protoService.Definition, "\n"))
	return nil
}

func (x *EnhanceWithProtoDefinition) getProtoForService(ctx context.Context, serviceName string) (*protoService, error) {
	var matches []*protoService
	if err := filepath.WalkDir(x.ProtoDirectory, func(p string, d os.DirEntry, err error) error {
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
