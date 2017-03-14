package vm

import (
	"testing"

	"github.com/vibhavp/dasm-go/read"
	"github.com/vibhavp/dasm-go/read/opcode"
)

var bCode2 = []int32{
	opcode.I32_SETJMP, 0,
	opcode.I32_JMPNOT1, 9,
	opcode.I32_LOAD, 2,
	opcode.I32_PRINT,
	opcode.JMP, 14,

	opcode.I32_LOAD, 3,
	opcode.I32_PRINT,
	opcode.I32_LONGJMP, 0,

	opcode.I32_LOAD, 1,
	opcode.I32_PRINT,
}

var bCode1 = []int32{
	opcode.I32_SETJMP, 0,
	opcode.I32_JMPNOT1, 9,
	opcode.I32_LOAD, 2,
	opcode.NOOP,
	opcode.JMP, 14,

	opcode.I32_LOAD, 3,
	opcode.NOOP,
	opcode.I32_LONGJMP, 0,

	opcode.I32_LOAD, 1,
	opcode.NOOP,
}

func TestRun(t *testing.T) {
	b1 := read.Bytecode{
		Bytecode:      bCode2,
		MaxStackDepth: 5,
	}
	Run(b1.Bytecode, b1.MaxStackDepth, false)
}

func BenchmarkRun(b *testing.B) {
	b1 := read.Bytecode{
		Bytecode:      bCode1,
		MaxStackDepth: 5,
	}
	for i := 0; i < b.N; i++ {
		Run(b1.Bytecode, b1.MaxStackDepth, false)
	}
}
