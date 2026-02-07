package a

import (
	"encoding/json"
	"log"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
)

func GoodUnmarshalUsage() {
	var data []byte = []byte(`["a", "b"]`)

	// Implicitly nil slice (zero value) - Safe
	var implicitNil []string
	if err := json.Unmarshal(data, &implicitNil); err != nil {
		log.Fatal(err)
	}

	// Explicitly initialize with an empty composite literal
	var s []string = []string{}
	if err := json.Unmarshal(data, &s); err != nil {
		log.Fatal(err)
	}

	// Explicitly initialize with nil
	var t []string = nil
	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
	}

	// Direct initialization in short declaration
	u := []string{}
	if err := json.Unmarshal(data, &u); err != nil {
		log.Fatal(err)
	}

	// Pointer to a new empty slice
	var v = new([]string)
	if err := json.Unmarshal(data, v); err != nil {
		log.Fatal(err)
	}
	
	// make with len 0
	var w = make([]string, 0)
	if err := json.Unmarshal(data, &w); err != nil {
		log.Fatal(err)
	}

	// Clean map
	var m = map[string]string{}
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatal(err)
	}

	// Clean struct
	var st = struct{ Field string }{}
	if err := json.Unmarshal(data, &st); err != nil {
		log.Fatal(err)
	}

	// util.Marshal good case
	var u2 []string
	if err := util.Marshal(data, &u2); err != nil {
		log.Fatal(err)
	}

	// Case: make with map capacity - Safe (always empty initially)
	const mapCapacity = 10
	var myMap = make(map[string]string, mapCapacity)
	if err := json.Unmarshal(data, &myMap); err != nil {
		log.Fatal(err)
	}
}
