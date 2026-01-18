package a

import (
	"encoding/json"
	"log"
)

type MyStruct struct {
	Field string
}

func BadUnmarshalUsage() {
	var data []byte = []byte(`["a", "b"]`)

	// Case 1: make with len > 0
	var t []string = make([]string, 5)
	if err := json.Unmarshal(data, &t); err != nil { // want "potential reuse of variable created with non-zero length"
		log.Fatal(err)
	}

	// Case 2: non-empty composite literal (slice)
	var v = []string{"foo"}
	if err := json.Unmarshal(data, &v); err != nil { // want "potential reuse of non-empty variable"
		log.Fatal(err)
	}

	// Case 3: non-empty composite literal (map)
	var m = map[string]string{"foo": "bar"}
	if err := json.Unmarshal(data, &m); err != nil { // want "potential reuse of non-empty variable"
		log.Fatal(err)
	}

	// Case 4: non-empty composite literal (struct)
	var st = MyStruct{Field: "foo"}
	if err := json.Unmarshal(data, &st); err != nil { // want "potential reuse of non-empty variable"
		log.Fatal(err)
	}
}