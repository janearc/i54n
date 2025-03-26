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

var verbRegistry = make(map[string]Verb)

// RegisterVerb is called by each verb library to register its function.
func RegisterVerb(v Verb) {
	verbRegistry[v.Name()] = v
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "bior",
		Short: "Bioreactor management tool",
	}

	// Automatically add a Cobra command for each registered verb.
	for _, verb := range verbRegistry {
		// Capture the current verb to avoid closure issues.
		v := verb
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
