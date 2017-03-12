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
		if insn == 0x28 {
			pc += 1
			out += fmt.Sprintf("%s %d\n", opstr, b.Bytecode[pc])
		} else {
			if !ok {
				return "", fmt.Errorf("Invalid Instruction: %d", insn)
			}
			out += fmt.Sprintf("%s\n", opstr)
		}
		pc += 1
	}

	return out[:len(out)-1], nil
}
