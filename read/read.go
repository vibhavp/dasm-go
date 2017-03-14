package read

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/vibhavp/dasm-go/read/opcode"
)

type Bytecode struct {
	Bytecode      []int32
	MaxStackDepth int
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

		if ar := opcode.NumOperands[int32(op)]; ar != 0 {
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

		b.MaxStackDepth += opcode.StackAdj[int32(op)]
	}
}

func FromFile(path string) (*Bytecode, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return FromReader(f)
}
