// Package greeting is the fixture's code: the agentlab workspace proofs
// change it, build it and open pull requests on it.
package greeting

import "fmt"

// Hello returns the greeting for name, or for the world when name is empty.
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
