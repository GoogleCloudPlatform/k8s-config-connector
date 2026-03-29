package main
import (
	"fmt"
	"strings"
)
func main() {
	method := "/google.bigtable.admin.v2.BigtableTableAdmin/GetSchemaBundle"
	fmt.Printf("method: %q\n", method)
	fmt.Printf("Contains /Get: %v\n", strings.Contains(method, "/Get"))
}
