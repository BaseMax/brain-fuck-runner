package bf

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	TapeSize = 30000
)

type Interpreter struct {
	Code    string
	Tape    []byte
	Pointer int
	loops   []int
	Output  strings.Builder
	Reader  *bufio.Reader
}

func New(code string) *Interpreter {
	return &Interpreter{
		Code:   strings.TrimSpace(code),
		Tape:   make([]byte, TapeSize),
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (b *Interpreter) Execute() (string, error) {
	code := b.Code
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			b.Tape[b.Pointer]++
		case '-':
			b.Tape[b.Pointer]--
		case '>':
			b.Pointer++
			if b.Pointer >= TapeSize {
				return "", errors.New("pointer out of bounds (>)")
			}
		case '<':
			b.Pointer--
			if b.Pointer < 0 {
				return "", errors.New("pointer out of bounds (<)")
			}
		case '.':
			b.Output.WriteByte(b.Tape[b.Pointer])
		case ',':
			char, err := b.Reader.ReadByte()
			if err != nil {
				return "", err
			}
			b.Tape[b.Pointer] = char
		case '[':
			if b.Tape[b.Pointer] == 0 {
				// Skip to matching ]
				skip := 1
				for skip > 0 {
					i++
					if i >= len(code) {
						return "", errors.New("unmatched '['")
					}
					if code[i] == '[' {
						skip++
					} else if code[i] == ']' {
						skip--
					}
				}
			} else {
				b.loops = append(b.loops, i)
			}
		case ']':
			if b.Tape[b.Pointer] != 0 {
				if len(b.loops) == 0 {
					return "", errors.New("unmatched ']'")
				}
				i = b.loops[len(b.loops)-1]
			} else {
				b.loops = b.loops[:len(b.loops)-1]
			}
		default:
			// Ignore unknown characters
		}
	}
	return b.Output.String(), nil
}
