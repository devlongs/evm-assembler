package asm

// Opcodes maps instruction names to their respective bytecodes.
var Opcodes = map[string]byte{
	"stop":   0x00,
	"add":    0x01,
	"push1":  0x60,
	"jump":   0x56,
	"jumpi":  0x57,
	"mstore": 0x52,
	"return": 0xf3,
	"revert": 0xfd,
}

func ValidateInstruction(name string) (byte, bool) {
	opcode, exists := Opcodes[name]
	return opcode, exists
}
