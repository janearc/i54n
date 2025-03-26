// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Verb interface {
	Name() string
	Execute(args []string) error
}

type SimpleVerb struct {
	name        string
	executeFunc func(args []string) error
}

func (sv SimpleVerb) Name() string {
	return sv.name
}

func (sv SimpleVerb) Execute(args []string) error {
	return sv.executeFunc(args)
}

func NewSimpleVerb(name string, executeFunc func(args []string) error) Verb {
	return SimpleVerb{name: name, executeFunc: executeFunc}
}

var verbRegistry = make(map[string]Verb)

func RegisterVerb(v Verb) {
	verbRegistry[v.Name()] = v
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "bior",
		Short: "Bioreactor management tool",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Add all registered verbs as subcommands.
	for _, verb := range verbRegistry {
		v := verb // capture range variable
		cmd := &cobra.Command{
			Use:   v.Name(),
			Short: fmt.Sprintf("Execute the %s verb", v.Name()),
			RunE: func(cmd *cobra.Command, args []string) error {
				return v.Execute(args)
			},
		}
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RegisterVerb(NewSimpleVerb("about", func(args []string) error {
		fmt.Println("bior is the ultimate bioreactor management tool, built for the Anthropocene.")
		return nil
	}))

	RegisterVerb(NewSimpleVerb("license", func(args []string) error {
		fmt.Println("Licensed under the MIT License.")
		return nil
	}))

	RegisterVerb(NewSimpleVerb("license", func(args []string) error {
		licenseBytes, err := os.ReadFile("LICENSE.md")
		if err != nil {
			return fmt.Errorf("error reading LICENSE.md: %v", err)
		}
		noticeBytes, err := os.ReadFile("NOTICE")
		if err != nil {
			return fmt.Errorf("error reading NOTICE: %v", err)
		}
		fmt.Println(string(licenseBytes))
		fmt.Println()
		fmt.Println(string(noticeBytes))
		return nil
	}))
}
