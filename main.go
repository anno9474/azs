package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("az", "group", "list", "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rgs, err := ParseGroups(output)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, g := range rgs {
		fmt.Printf("[%d] %s (%s)\n", i, g.Name, g.Location)
	}
}
