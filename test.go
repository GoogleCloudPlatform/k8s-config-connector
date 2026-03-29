package main
import (
	"fmt"
	"strings"
)
func main() {
	method := "/google.bigtable.admin.v2.BigtableTableAdmin/GetSchemaBundle"
	fmt.Println(strings.Contains(method, "/Get"))
}
