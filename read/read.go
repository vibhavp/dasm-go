package read

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type Bytecode struct {
	Bytecode      []int32
	MaxStackDepth int
}

const (
	I32_LOAD    = 0x28
	I32_ADD     = 0x6a
	I32_MULT    = 0x6b
	I32_SUB     = 0x6c
	I32_PRINT   = 0xcc
	I32_SETJMP  = 0x11
	I32_LONGJMP = 0x10
)

// How many operands (if any) does an op have?
var NumOperands = map[int32]int{
	I32_LOAD: 1,
} //everything else becomes 0

// How much does an op increase the stack depth by?
var stackAdj = map[int32]int{
	I32_LOAD: 1,
}

func FromReader(reader io.Reader) (*Bytecode, error) {
	sig := make([]byte, 4)
	if err := binary.Read(reader, binary.LittleEndian, &sig); err != nil {
		return nil, err
	}
	if bytes.Compare(sig[0:], []byte{0, 'a', 's', 'm'}) != 0 {
		return nil, errors.New("read: missing or invalid signature")
	}
	b := &Bytecode{[]int32{}, 0}

	for {
		var op byte
		err := binary.Read(reader, binary.LittleEndian, &op)
		if err == io.EOF {
			return b, nil
		} else if err != nil {
			return nil, err
		}
		// FIXME: this takes 4 bytes bytes for each op, which wastes space
		b.Bytecode = append(b.Bytecode, int32(op))

		if ar := NumOperands[int32(op)]; ar != 0 {
			i := 0
			for i < ar {
				var op int32
				err := binary.Read(reader, binary.LittleEndian, &op)
				if err == io.EOF {
					return nil, fmt.Errorf("read: Incomplete instruction %d", op)
				} else if err != nil {
					return nil, err
				}
				b.Bytecode = append(b.Bytecode, op)
				i += 1
			}
		}

		b.MaxStackDepth += stackAdj[int32(op)]
	}
}

func FromFile(path string) (*Bytecode, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return FromReader(f)
}
