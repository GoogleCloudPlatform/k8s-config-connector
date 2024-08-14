// Copyright 2024 Google LLC
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

package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "[::]:50051", "the address to connect to")
	cmd  = flag.String("command", "validate", "validate|evaluate")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error closing connection: %v", err)
		}
	}()
	c := pb.NewExpanderClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if *cmd == "validate" {
		r, err := c.Validate(ctx,
			&pb.ValidateRequest{
				Config:  []byte{},
				Context: []byte{},
				Facade:  []byte{},
				Value:   []byte{},
			})
		if err != nil {
			log.Fatalf("could not validate: %v", err)
		}
		log.Printf("status: %s", r.GetStatus())
	} else {
		r, err := c.Evaluate(ctx,
			&pb.EvaluateRequest{
				Config:  []byte{},
				Context: []byte{},
				Facade:  []byte{},
				Value:   []byte{},
			})
		if err != nil {
			log.Fatalf("could not evaluate: %v", err)
		}
		log.Printf("status: %s", r.GetStatus())
	}
}
