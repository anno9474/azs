package main

import "encoding/json"

type ResourceGroup struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// ParseGroups parses JSON into a slice of ResourceGroup
func ParseGroups(data []byte) ([]ResourceGroup, error) {
	var groups []ResourceGroup
	err := json.Unmarshal(data, &groups)
	return groups, err
}
