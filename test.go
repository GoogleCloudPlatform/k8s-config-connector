package main

import (
    "fmt"
    pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
)

func main() {
    fmt.Printf("%T\n", pb.FirewallEndpoint{})
}
