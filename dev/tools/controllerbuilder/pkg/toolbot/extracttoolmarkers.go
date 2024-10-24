package toolbot

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"k8s.io/klog"
)

type ExtractToolMarkers struct {
}

func (x *ExtractToolMarkers) Extract(ctx context.Context, b []byte) ([]*DataPoint, error) {
	var dataPoints []*DataPoint

	r := bytes.NewReader(b)
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("scanning code: %w", err)
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)
			if strings.HasPrefix(comment, "+tool:") {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := strings.TrimPrefix(comment, "+tool:")
				dataPoint := &DataPoint{
					Type:   toolName,
					Output: string(b),
				}

				for {
					line, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return nil, fmt.Errorf("scanning code: %w", err)
					}
					line = strings.TrimSpace(line)
					if !strings.HasPrefix(line, "//") {
						break
					}
					toolLine := strings.TrimPrefix(line, "//")
					toolLine = strings.TrimPrefix(toolLine, " ")

					tokens := strings.SplitN(toolLine, ":", 2)
					if len(tokens) == 2 {
						dataPoint.SetInput(tokens[0], strings.TrimSpace(tokens[1]))
					} else {
						return nil, fmt.Errorf("cannot parse tool line %q", toolLine)
					}
				}
				dataPoints = append(dataPoints, dataPoint)
			}

			if strings.HasPrefix(comment, "+kcc:proto=") {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := "kcc-proto"
				dataPoint := &DataPoint{
					Type: toolName,
				}

				proto := strings.TrimPrefix(comment, "+kcc:proto=")
				dataPoint.SetInput("proto.message", proto)

				var bb bytes.Buffer
				for {
					line, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return nil, fmt.Errorf("scanning code: %w", err)
					}

					bb.WriteString(line)

					s := strings.TrimSpace(line)
					if strings.HasPrefix(s, "}") {
						break
					}
				}
				dataPoint.Output = bb.String()
				dataPoints = append(dataPoints, dataPoint)
			}
		}
	}
	return dataPoints, nil
}
