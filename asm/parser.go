package asm

import (
	"strings"
)

// Instruction represents a parsed EVM instruction.
type Instruction struct {
	Name      string
	Arguments []string
	Offset    int
}

// ParseAssembly parses the raw assembly code into structured instructions.
func ParseAssembly(input string) ([]Instruction, error) {
	lines := strings.Split(input, "\n")
	var instructions []Instruction
	offset := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, ";") {
			continue
		}

		parts := strings.Fields(line)
		instruction := Instruction{
			Name:      parts[0],
			Arguments: parts[1:],
			Offset:    offset,
		}
		instructions = append(instructions, instruction)

		// Estimate the offset increment (1 byte for the opcode + argument bytes)
		offset += 1 + len(instruction.Arguments)
	}

	return instructions, nil
}
