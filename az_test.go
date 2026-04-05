package main

import (
	"os/exec"
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

func TestFetchResourceGroupsMock(t *testing.T) {
	execCommand = func(name string, args ...string) *exec.Cmd {
		return exec.Command("echo", `[{"name":"rg1","location":"eastus"}]`)
	}
	defer func() { execCommand = exec.Command }() // restore

	rgs, err := FetchResourceGroups()
	if err != nil {
		t.Fatalf("FetchResourceGroups failed: %v", err)
	}
	if len(rgs) != 1 || rgs[0].Name != "rg1" {
		t.Errorf("unexpected result: %+v", rgs)
	}
}
