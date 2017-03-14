package disasm

import (
	"fmt"

	"github.com/vibhavp/dasm-go/read"
)

var opMap = map[int32]string{
	0x28: "i32_load",
	0x6a: "i32_add",
	0x6b: "i32_sub",
	0x6c: "i32_mul",
	0xcc: "i32_print",
}

func ToDasm(b *read.Bytecode) (string, error) {
	if len(b.Bytecode) == 0 {
		return "", nil
	}
	out := ""
	pc := 0
	var insn int32
	for pc < len(b.Bytecode) {
		insn = b.Bytecode[pc]
		opstr, ok := opMap[insn]

		if !ok {
			return "", fmt.Errorf("Invalid Instruction: %d", insn)
		} else {
			out += fmt.Sprintf("%s", opstr)
		}

		if ar := read.NumOperands[insn]; ar != 0 {
			if pc+ar >= len(b.Bytecode) {
				return "", fmt.Errorf("disasm: incomplete instruction %s", opstr)
			}

			i := 0
			operands := ""
			for i < ar {
				pc += 1
				i += 1
				operands += fmt.Sprintf(" %d", b.Bytecode[pc])
			}
			out += operands
		}
		out += "\n"
		pc += 1
	}

	return out[:len(out)-1], nil
}
