package label

import "fmt"

// Label represents a label and its offset.
type Label struct {
	Name   string
	Offset int
}

// ResolveLabels resolves label references in the instructions.
func ResolveLabels(labels map[string]int, unresolved []Label) error {
	for _, label := range unresolved {
		if _, exists := labels[label.Name]; !exists {
			return fmt.Errorf("undefined label: %s", label.Name)
		}
	}
	return nil
}
