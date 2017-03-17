package opcode

const (
	I32_LOAD    = 0x28
	F32_LOAD    = 0x29
	I32_ADD     = 0x6a
	F32_ADD     = 0x30
	I32_MULT    = 0x6b
	F32_MULT    = 0x31
	I32_SUB     = 0x6c
	F32_SUB     = 0x32
	I32_PRINT   = 0xcc
	F32_PRINT   = 0x33
	I32_SETJMP  = 0x11
	I32_LONGJMP = 0x10
	I32_JMP1    = 0x12
	I32_JMPNOT1 = 0x13
	JMP         = 0x14
	I32_EQ      = 0x15
	I32_GREATER = 0x16
	I32_GEQ     = 0x17
	I32_LESS    = 0x18
	I32_LEQ     = 0x19
	F32_EQ      = 0x20
	F32_GREATER = 0x1a
	F32_GEQ     = 0x1b
	F32_LESS    = 0x1c
	F32_LEQ     = 0x1d
	NOOP        = 0x00
)

// How many operands (if any) does an op have?
var NumOperands = map[int32]int{
	I32_LOAD:    1,
	F32_LOAD:    1,
	I32_SETJMP:  1,
	I32_LONGJMP: 1,
} //everything else becomes 0

// How much does an op increase the stack depth by?
var StackAdj = map[int32]int{
	I32_LOAD:    1,
	F32_LOAD:    1,
	I32_SETJMP:  1,
	I32_LONGJMP: 1,
}

var OpStr = map[int32]string{
	I32_LOAD:    "i32_load",
	F32_LOAD:    "f32_load",
	I32_ADD:     "i32_add",
	F32_ADD:     "f32_add",
	I32_MULT:    "i32_mult",
	F32_MULT:    "f32_mult",
	I32_SUB:     "i32_sub",
	F32_SUB:     "f32_sub",
	I32_PRINT:   "i32_print",
	F32_PRINT:   "f32_print",
	I32_SETJMP:  "i32_setjmp",
	I32_LONGJMP: "i32_longjmp",
	I32_JMP1:    "i32_jmp1",
	I32_JMPNOT1: "i32_jmpno1",
	I32_EQ:      "i32_eq",
	I32_GREATER: "i32_greater",
	I32_GEQ:     "i32_geq",
	I32_LESS:    "i32_less",
	I32_LEQ:     "i32_leq",
	F32_GREATER: "f32_greater",
	F32_GEQ:     "f32_geq",
	F32_LESS:    "f32_less",
	F32_LEQ:     "f32_leq",
	JMP:         "jmp",
	NOOP:        "noop",
}
