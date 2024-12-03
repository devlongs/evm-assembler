package evm

import (
	"fmt"

	"github.com/devlongs/evm-assembler/asm"
)

// GenerateBytecode converts parsed instructions into EVM bytecode.
func GenerateBytecode(instructions []asm.Instruction) ([]byte, error) {
	var bytecode []byte

	for _, inst := range instructions {
		opcode, valid := asm.ValidateInstruction(inst.Name)
		if !valid {
			return nil, fmt.Errorf("invalid instruction: %s", inst.Name)
		}

		bytecode = append(bytecode, opcode)
		for _, arg := range inst.Arguments {
			bytecode = append(bytecode, []byte(arg)...)
		}
	}

	return bytecode, nil
}
