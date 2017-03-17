package vm

import (
	"fmt"
	"strconv"

	"github.com/vibhavp/dasm-go/read/opcode"
	"math"
)

type VMRuntimeError struct {
	pc  int
	err string
}

const invalidInstruction = "Invalid Instruction"
const stackOverflow = "Stack Overflow"
const stackUnderflow = "Stack Underflow"
const invalidContext = "Invalid Context"
const invalidAddr = "Invalid Address"

func (v VMRuntimeError) Error() string {
	return fmt.Sprintf("Error executing bytecode at pc=%d: %s", v.pc, v.err)
}

type context struct {
	stack []int32
	pc    int
	top   int
}

type vm struct {
	context
	maxStackDepth int

	bytecode []int32
	safe     bool

	savedContexts map[int32]*context
}

// I should manually inline all of this, panicking in a function makes it
// impossible for go to inline the function, resulting in unecessary overhead.

func (v *vm) pop() int32 {
	if v.safe && v.top < -1 {
		panic(VMRuntimeError{v.pc, stackUnderflow})
	}

	v1 := v.stack[v.top]
	v.top -= 1
	return v1
}

func (v *vm) push(i int32) {
	v.top += 1
	if v.safe && v.top == v.maxStackDepth {
		panic(VMRuntimeError{v.pc, stackOverflow})
	}

	v.stack[v.top] = i
}

func (v *vm) fetch() int32 {
	if v.safe && v.pc+1 == len(v.bytecode) {
		panic(VMRuntimeError{v.pc, invalidInstruction})
	}
	v.pc += 1
	return v.bytecode[v.pc]
}

func Run(bytecode []int32, maxStackDepth int, safe bool) {
	var insn int32

	stack := make([]int32, maxStackDepth)
	vm := &vm{
		context: context{
			stack: stack[0:len(stack)], // doesn't copy stack
			pc:    0,
			top:   -1,
		},
		maxStackDepth: maxStackDepth,
		bytecode:      bytecode,
		safe:          safe,
		savedContexts: make(map[int32]*context),
	}

	for vm.pc < len(bytecode) {
		insn = bytecode[vm.pc]
		switch insn {
		case opcode.I32_LOAD: // i32_load
			vm.push(vm.fetch())
		case opcode.F32_LOAD: // f32_load
			vm.push(vm.fetch())
		case opcode.I32_ADD: //i32_add
			v1 := vm.pop()
			v2 := vm.pop()
			vm.push(v1 + v2)
		case opcode.F32_ADD:
			v1 := math.Float32frombits(uint32(vm.pop()))
			v2 := math.Float32frombits(uint32(vm.pop()))
			vm.push(int32(math.Float32bits(v1 + v2)))
		case opcode.I32_MULT: //i32_mult
			v1 := vm.pop()
			v2 := vm.pop()
			vm.push(v1 * v2)
		case opcode.F32_MULT:
			v1 := math.Float32frombits(uint32(vm.pop()))
			v2 := math.Float32frombits(uint32(vm.pop()))
			vm.push(int32(math.Float32bits(v1 * v2)))
		case opcode.I32_SUB: //i32_sub
			v1 := vm.pop()
			v2 := vm.pop()
			vm.push(v1 - v2)
		case opcode.F32_SUB:
			v1 := math.Float32frombits(uint32(vm.pop()))
			v2 := math.Float32frombits(uint32(vm.pop()))
			vm.push(int32(math.Float32bits(v1 - v2)))
		case opcode.I32_PRINT: //i32_print
			fmt.Println(strconv.FormatInt(int64(vm.pop()), 10))
		case opcode.F32_PRINT:
			fmt.Println(math.Float32frombits(uint32(vm.pop())))
		case opcode.I32_SETJMP:
			saved := &context{
				stack: make([]int32, vm.top+1),
				pc:    vm.pc + 1,
				top:   vm.top,
			}
			copy(saved.stack, vm.stack)
			vm.savedContexts[vm.fetch()] = saved
			vm.push(0)
		case opcode.I32_LONGJMP:
			ctxt := vm.savedContexts[vm.fetch()]
			if ctxt == nil {
				panic(VMRuntimeError{vm.pc, invalidContext})
			}
			vm.top = ctxt.top
			copy(vm.stack, ctxt.stack)
			vm.pc = ctxt.pc
			vm.push(1)
		case opcode.I32_JMP1:
			v1 := vm.pop()
			addr := vm.fetch()
			if v1 == 1 {
				if int(addr) >= len(vm.bytecode) {
					panic(VMRuntimeError{vm.pc, invalidAddr})
				}
				vm.pc = int(addr)
				continue
			}
		case opcode.I32_JMPNOT1:
			v1 := vm.pop()
			addr := vm.fetch()
			if v1 != 1 {
				if int(addr) >= len(vm.bytecode) {
					panic(VMRuntimeError{vm.pc, invalidAddr})
				}
				vm.pc = int(addr)
				continue
			}
		case opcode.JMP:
			addr := vm.fetch()
			if int(addr) >= len(vm.bytecode) {
				panic(VMRuntimeError{vm.pc, invalidAddr})
			}
			vm.pc = int(addr)
			continue
		case opcode.NOOP:
		case opcode.I32_EQ:
			// go wouldn't let me do vm.push(int32(vm.pop() == vm.pop()))
			if vm.pop() == vm.pop() {
				vm.push(1)
			} else {
				vm.push(0)
			}
		case opcode.I32_GREATER:
			if vm.pop() > vm.pop() {
				vm.push(1)
			} else {
				vm.push(0)
			}
		case opcode.I32_GEQ:
			if vm.pop() >= vm.pop() {
				vm.push(1)
			} else {
				vm.push(0)
			}
		case opcode.I32_LESS:
			if vm.pop() < vm.pop() {
				vm.push(1)
			} else {
				vm.push(0)
			}
		case opcode.I32_LEQ:
			if vm.pop() <= vm.pop() {
				vm.push(1)
			} else {
				vm.push(0)
			}
		default:
			panic(VMRuntimeError{vm.pc, invalidInstruction})
		}
		vm.pc += 1
	}
}
