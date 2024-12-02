package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var opcodes = map[string]byte{
	"stop":   0x00,
	"add":    0x01,
	"push1":  0x60,
	"jump":   0x56,
	"jumpi":  0x57,
	"mstore": 0x52,
	"return": 0xf3,
	"revert": 0xfd,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: assembler <input file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	labels := make(map[string]int)
	var bytecode []byte
	var unresolved []struct {
		label string
		pos   int
	}

	scanner := bufio.NewScanner(file)
	currentOffset := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, ";") {
			continue
		}

		// Handle labels
		if strings.HasSuffix(line, ":") {
			label := strings.TrimSuffix(line, ":")
			labels[label] = currentOffset
			continue
		}

		// Parse instruction and arguments
		parts := strings.Fields(line)
		instruction := parts[0]

		if opcode, found := opcodes[instruction]; found {
			bytecode = append(bytecode, opcode)
			currentOffset++

			if strings.HasPrefix(instruction, "push") && len(parts) > 1 {
				arg := parts[1]
				if strings.HasPrefix(arg, "@") {
					unresolved = append(unresolved, struct {
						label string
						pos   int
					}{label: arg[1:], pos: currentOffset})
					bytecode = append(bytecode, 0)
				} else {
					val, err := strconv.ParseInt(arg, 0, 8)
					if err != nil {
						fmt.Printf("Invalid argument for %s: %s\n", instruction, arg)
						return
					}
					bytecode = append(bytecode, byte(val))
				}
				currentOffset++
			}
		} else {
			fmt.Printf("Unknown instruction: %s\n", instruction)
			return
		}
	}

	for _, ref := range unresolved {
		if pos, found := labels[ref.label]; found {
			bytecode[ref.pos] = byte(pos)
		} else {
			fmt.Printf("Undefined label: %s\n", ref.label)
			return
		}
	}

	fmt.Printf("Bytecode: %x\n", bytecode)
}
