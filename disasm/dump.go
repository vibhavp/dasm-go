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
	pc := 0
	var insn int32
	for pc < len(b.Bytecode) {
		insn = b.Bytecode[pc]
		if insn == 0x28 {
			pc += 1
			out += fmt.Sprintf("0x%x 0x%x\n", insn, b.Bytecode[pc])
		} else {
			out += fmt.Sprintf("0x%x\n", insn)
		}
		pc += 1
	}

	return out[:len(out)-1], nil

}
