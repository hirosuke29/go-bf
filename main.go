package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	RIGHT      byte = '>'
	LEFT       byte = '<'
	INCREMENT  byte = '+'
	DECREMENT  byte = '-'
	INPUT      byte = ','
	OUTPUT     byte = '.'
	LOOP_START byte = '['
	LOOP_END   byte = ']'
)

const MEMORY_SIZE uint = 2000000000

type Stack struct {
	data []uint
}

func NewStack() *Stack {
	s := new(Stack)
	s.data = []uint{}
	return s
}

func (s *Stack) push(arg uint) error {
	s.data = append(s.data, arg)
	return nil
}

func (s *Stack) pop() uint {
	retVal := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return retVal
}

func readSource(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getChar() (*byte, error) {
	r := bufio.NewReader(os.Stdin)
	b, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func exec(bytes []byte) error {
	var memory [MEMORY_SIZE]uint8

	var ptr uint = 0

	length := len(bytes)
	var cur uint = 0
	var loopStack *Stack = NewStack()
	for int(cur) < length {
		token := bytes[cur]
		switch token {
		case RIGHT:
			if ptr >= MEMORY_SIZE-1 {
				ptr = 0
			} else {
				ptr++
			}
		case LEFT:
			if ptr <= 0 {
				ptr = MEMORY_SIZE - 1
			} else {
				ptr--
			}
		case INCREMENT:
			memory[ptr]++
		case DECREMENT:
			memory[ptr]--
		case INPUT:
			b, err := getChar()
			if err != nil {
				return err
			}
			memory[ptr] = uint8(*b)
		case OUTPUT:
			fmt.Print(string(memory[ptr]))
		case LOOP_START:
			loopStack.push(cur)
			if memory[ptr] == 0 {
				depth := 1
				for depth > 0 {
					cur++
					if bytes[cur] == LOOP_START {
						depth++
					}
					if bytes[cur] == LOOP_END {
						depth--
					}
				}
				loopStack.pop()
			}
		case LOOP_END:
			cur = loopStack.pop() - 1
		}
		cur++
	}
	fmt.Printf("\n")
	return nil
}

func main() {
	filename := os.Args[1]
	bytes, err := readSource(filename)
	if err != nil {
		fmt.Print(err)
		return
	}
	if err := exec(bytes); err != nil {
		fmt.Print(err)
		return
	}
}
