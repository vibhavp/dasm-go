package read

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
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
		switch op {
		case 0x28:
			var i int32
			err := binary.Read(reader, binary.LittleEndian, &i)
			if err == io.EOF {
				return nil, errors.New("read: incomplete instruction i32_load")
			} else if err != nil {
				return nil, err
			}
			b.Bytecode = append(b.Bytecode, i)
			b.MaxStackDepth += 1
		}
	}
}

func FromFile(path string) (*Bytecode, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return FromReader(f)
}
