package main

import (
	"reflect"
	"testing"
)

func TestParseGroups(t *testing.T) {
	sampleJSON := []byte(`
        [
            {"name": "rg1", "location": "westeurope"},
            {"name": "rg2", "location": "northeurope"}
        ]
    `)

	got, err := ParseGroups(sampleJSON)
	if err != nil {
		t.Fatalf("ParseGroups returned error: %v", err)
	}

	want := []ResourceGroup{
		{Name: "rg1", Location: "westeurope"},
		{Name: "rg2", Location: "northeurope"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseGroups = %+v, want %+v", got, want)
	}
}
