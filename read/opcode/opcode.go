package opcode

const (
	I32_LOAD    = 0x28
	I32_ADD     = 0x6a
	I32_MULT    = 0x6b
	I32_SUB     = 0x6c
	I32_PRINT   = 0xcc
	I32_SETJMP  = 0x11
	I32_LONGJMP = 0x10
	I32_JMP1    = 0x12
	I32_JMPNOT1 = 0x13
	JMP         = 0x14
	NOOP        = 0x00
)

// How many operands (if any) does an op have?
var NumOperands = map[int32]int{
	I32_LOAD:    1,
	I32_SETJMP:  1,
	I32_LONGJMP: 1,
} //everything else becomes 0

// How much does an op increase the stack depth by?
var StackAdj = map[int32]int{
	I32_LOAD:    1,
	I32_SETJMP:  1,
	I32_LONGJMP: 1,
}

var OpStr = map[int32]string{
	I32_LOAD:    "i32_load",
	I32_ADD:     "i32_add",
	I32_MULT:    "i32_mult",
	I32_SUB:     "i32_sub",
	I32_PRINT:   "i32_print",
	I32_SETJMP:  "i32_setjmp",
	I32_LONGJMP: "i32_longjmp",
	I32_JMP1:    "i32_jmp1",
	I32_JMPNOT1: "i32_jmpno1",
	JMP:         "jmp",
	NOOP:        "noop",
}
