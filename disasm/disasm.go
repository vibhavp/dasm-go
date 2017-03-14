package disasm

import (
	"fmt"

	"github.com/vibhavp/dasm-go/read"
	"github.com/vibhavp/dasm-go/read/opcode"
)

func ToDasm(b *read.Bytecode) (string, error) {
	if len(b.Bytecode) == 0 {
		return "", nil
	}
	out := ""
	var insn int32
	for pc := 0; pc < len(b.Bytecode); pc += 1 {
		insn = b.Bytecode[pc]
		opstr, ok := opcode.OpStr[insn]

		if !ok {
			return "", fmt.Errorf("Invalid Instruction: %d", insn)
		} else {
			out += fmt.Sprintf("%s", opstr)
		}

		if ar := opcode.NumOperands[insn]; ar != 0 {
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
	}

	return out[:len(out)-1], nil
}
