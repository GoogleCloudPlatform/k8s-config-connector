package a

import (
	"encoding/json"
	"log"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
)

type MyStruct struct {
	Field string
}

type Embedded struct {
	Inner string
}

type Wrapper struct {
	Embedded `json:",inline"`
}

func BadUnmarshalUsage() {
	const sliceLen = 5
	var data []byte = []byte(`["a", "b"]`)
	var s []string 
	
	if err := json.Unmarshal(data, &s); err != nil { 
		log.Fatal(err)
	}

	// Case 1: make with len > 0
	var t []string = make([]string, 5)
	if err := json.Unmarshal(data, &t); err != nil { // want "potential reuse of variable created with non-zero length"
		log.Fatal(err)
	}

	// Case 2: non-empty composite literal (slice)
	var v = []string{"foo"}
	if err := json.Unmarshal(data, &v); err != nil { // want "potential reuse of non-empty slice; existing elements will be lost"
		log.Fatal(err)
	}

	// Case 3: non-empty composite literal (map)
	var m = map[string]string{"foo": "bar"}
	if err := json.Unmarshal(data, &m); err != nil { // want "potential reuse of non-empty map; existing elements will be merged"
		log.Fatal(err)
	}

	// Case 4: non-empty composite literal (struct)
	var st = MyStruct{Field: "foo"}
	if err := json.Unmarshal(data, &st); err != nil { // want "potential reuse of non-empty variable"
		log.Fatal(err)
	}

	// Case 5: util.Marshal with bad reuse
	var v2 = []string{"foo"}
	if err := util.Marshal(data, &v2); err != nil { // want "potential reuse of non-empty slice; existing elements will be lost"
		log.Fatal(err)
	}

	// Case 6: non-empty inlined field
	var w = Wrapper{Embedded: Embedded{Inner: "foo"}}
	if err := json.Unmarshal(data, &w); err != nil { // want "potential reuse of non-empty variable"
		log.Fatal(err)
	}

	// Case 7: Pointer to non-empty struct literal (AssignStmt)
	sPtr := &MyStruct{Field: "foo"}
	if err := json.Unmarshal(data, sPtr); err != nil { // want "potential reuse of non-empty variable"
		log.Fatal(err)
	}

	// Case 8: make with named constant for length
	var t2 []string = make([]string, sliceLen)
	if err := json.Unmarshal(data, &t2); err != nil { // want "potential reuse of variable created with non-zero length"
		log.Fatal(err)
	}
}

