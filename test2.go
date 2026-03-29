package main
import "fmt"
type Request struct {
	Method string
	URL    string
}
type Response struct {
	Status string
	Code int
}
type LogEntry struct {
	Timestamp string
	Request   Request
	Response  Response
}
func main() {
	e := &LogEntry{
		Timestamp: "0001-01-01 00:00:00 +0000 UTC",
		Request: Request{
			Method: "GRPC",
			URL: "/google.bigtable.admin.v2.BigtableTableAdmin/GetSchemaBundle",
		},
		Response: Response{
			Status: "OK",
			Code: 0,
		},
	}
	fmt.Printf("%v\n", e)
}
