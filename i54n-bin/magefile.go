//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Build compiles the bior binary.
func Build() error {
	fmt.Println("Building bior binary...")
	cmd := exec.Command("go", "build", "-o", "bior")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Clean removes the bior binary.
func Clean() error {
	fmt.Println("Cleaning up bior binary...")
	err := os.Remove("bior")
	if err != nil {
		fmt.Println("Binary not found, nothing to clean.")
	}
	return nil
}
