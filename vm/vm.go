package vm

import (
	"fmt"
	"strconv"
)

type VMRuntimeError struct {
	pc  int
	err string
}

const invalidInstruction = "Invalid Instruction"
const stackOverflow = "Stack Overflow"
const stackUnderflow = "Stack Underflow"

func (v VMRuntimeError) Error() string {
	return fmt.Sprintf("Error executing bytecode at pc=%d: %s", v.pc, v.err)
}

type vm struct {
	stack         []int32
	maxStackDepth int

	bytecode []int32
	pc       int
	top      int
	safe     bool
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

func Run(bytecode []int32, maxStackDepth int, safe bool) {
	var insn int32

	stack := make([]int32, maxStackDepth)
	vm := vm{
		stack:         stack[0:len(stack)], // doesn't copy stack
		maxStackDepth: maxStackDepth,
		bytecode:      bytecode,
		pc:            0,
		top:           -1,
		safe:          safe,
	}

	for vm.pc < len(bytecode) {
		insn = bytecode[vm.pc]
		switch insn {
		case 0x28: // i32_load
			if safe && vm.pc+1 == len(bytecode) {
				panic(VMRuntimeError{vm.pc, invalidInstruction})
			}
			vm.push(vm.bytecode[vm.pc+1])
			vm.pc += 1

		case 0x6a: //i32_add
			v1 := vm.pop()
			v2 := vm.pop()
			vm.push(v1 + v2)
		case 0x6b: //i32_mult
			v1 := vm.pop()
			v2 := vm.pop()
			vm.push(v1 * v2)
		case 0x6c: //i32_sub
			v1 := vm.pop()
			v2 := vm.pop()
			vm.push(v1 - v2)
		case 0xcc: //i32_print
			fmt.Println(strconv.FormatInt(int64(vm.pop()), 10))
		default:
			panic(VMRuntimeError{vm.pc, invalidInstruction})
		}
		vm.pc += 1
	}
}
