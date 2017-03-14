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
