package vm

import (
	"fmt"
	"strconv"

	"github.com/vibhavp/dasm-go/read/opcode"
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

func (v *vm) next()  { v.pc += 1 }
func (v *vm) next2() { v.pc += 2 }

var opTable = map[int32]func(v *vm){
	opcode.I32_LOAD: func(vm *vm) {
		if vm.safe && vm.pc+1 == len(vm.bytecode) {
			panic(VMRuntimeError{vm.pc, invalidInstruction})
		}
		vm.push(vm.fetch())
		vm.next()
	},
	opcode.I32_ADD: func(vm *vm) {
		v1 := vm.pop()
		v2 := vm.pop()
		vm.push(v1 + v2)
		vm.next()
	},
	opcode.I32_MULT: func(vm *vm) {
		v1 := vm.pop()
		v2 := vm.pop()
		vm.push(v1 * v2)
		vm.next()
	},
	opcode.I32_SUB: func(vm *vm) {
		v1 := vm.pop()
		v2 := vm.pop()
		vm.push(v1 - v2)
		vm.next()
	},
	opcode.I32_PRINT: func(vm *vm) {
		fmt.Println(strconv.FormatInt(int64(vm.pop()), 10))
		vm.next()
	},
	opcode.I32_SETJMP: func(vm *vm) {
		saved := &context{
			stack: make([]int32, vm.top+1),
			pc:    vm.pc,
			top:   vm.top,
		}
		copy(saved.stack, vm.stack)
		vm.savedContexts[vm.fetch()] = saved
		vm.push(0)
		vm.next()
	},
	opcode.I32_LONGJMP: func(vm *vm) {
		ctxt := vm.savedContexts[vm.fetch()]
		if ctxt == nil {
			panic(VMRuntimeError{vm.pc, invalidContext})
		}
		vm.top = ctxt.top
		copy(vm.stack, ctxt.stack)
		vm.pc = ctxt.pc
		vm.push(1)
		vm.next2()
	},
	opcode.I32_JMP1: func(vm *vm) {
		v1 := vm.pop()
		addr := vm.fetch()
		if v1 == 1 {
			if int(addr) >= len(vm.bytecode) {
				panic(VMRuntimeError{vm.pc, invalidAddr})
			}
			vm.pc = int(addr)
		} else {
			vm.next()
		}
	},
	opcode.I32_JMPNOT1: func(vm *vm) {
		v1 := vm.pop()
		addr := vm.fetch()
		if v1 != 1 {
			if int(addr) >= len(vm.bytecode) {
				panic(VMRuntimeError{vm.pc, invalidAddr})
			}
			vm.pc = int(addr)
		} else {
			vm.next()
		}
	},
	opcode.JMP: func(vm *vm) {
		addr := vm.fetch()
		if int(addr) >= len(vm.bytecode) {
			panic(VMRuntimeError{vm.pc, invalidAddr})
		}
		vm.pc = int(addr)
	},
	opcode.NOOP: func(vm *vm) { vm.next() },
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
		opTable[insn](vm)
		// fn(vm)
	}
}
