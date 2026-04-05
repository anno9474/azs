package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type ResourceGroup struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func ParseGroups(data []byte) ([]ResourceGroup, error) {
	var groups []ResourceGroup
	err := json.Unmarshal(data, &groups)
	return groups, err
}

// FetchResourceGroups runs az CLI and returns parsed resource groups
var execCommand = exec.Command

func FetchResourceGroups() ([]ResourceGroup, error) {
	cmd := execCommand("az", "group", "list", "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az command failed: %w", err)
	}
	return ParseGroups(output)
}
