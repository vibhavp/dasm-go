package vm

import (
	"testing"

	"github.com/vibhavp/dasm-go/read"
)

func TestRun(t *testing.T) {
	c := []int32{
		read.I32_SETJMP, 0,
		read.I32_JMPNOT1, 9,
		read.I32_LOAD, 2,
		read.I32_PRINT,
		read.JMP, 14,

		read.I32_LOAD, 3,
		read.I32_PRINT,
		read.I32_LONGJMP, 0,

		read.I32_LOAD, 1,
		read.I32_PRINT,
	}
	b1 := read.Bytecode{
		Bytecode:      c,
		MaxStackDepth: 5,
	}
	Run(b1.Bytecode, b1.MaxStackDepth, false)
}
