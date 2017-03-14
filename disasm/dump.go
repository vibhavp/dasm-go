package disasm

import (
	"fmt"

	"github.com/vibhavp/dasm-go/read"
)

func Dump(b *read.Bytecode) (string, error) {
	if len(b.Bytecode) == 0 {
		return "", nil
	}
	out := "\n"
	var insn int32
	for pc := 0; pc < len(b.Bytecode); pc += 1 {
		insn = b.Bytecode[pc]
		if insn == 0x28 {
			pc += 1
			out += fmt.Sprintf("0x%x 0x%x\n", insn, b.Bytecode[pc])
		} else {
			out += fmt.Sprintf("0x%x\n", insn)
		}
	}

	return out[:len(out)-1], nil

}
