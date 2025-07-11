package entity

import (
	"fmt"

	"github.com/breno5g/emugo-8/internal/consts"
)

type Chip8 struct {
	Memory [consts.MemorySize]byte
	PC     uint16
	Stack  [consts.StackSize]uint16
	SP     uint16
	V      [16]byte
	I      uint16
	DT     byte
	ST     byte
	Screen [consts.DisplaySize]bool
	Keys   [16]bool
}

func NewChip8() *Chip8 {
	c := &Chip8{
		PC: consts.StartAddress,
	}
	c.LoadFontSet()
	return c
}

func (c *Chip8) LoadFontSet() {
	// for i := 0; i < 80; i++ {
	// 	c.Memory[consts.FontStartAddress+i] = consts.FontSet[i]
	// }
	copy(c.Memory[consts.FontStartAddress:], consts.FontSet[:])
}

func (c *Chip8) DebugScreen() {
	for y := range consts.DisplayHeight {
		for x := range consts.DisplayWidth {
			if c.Screen[y*consts.DisplayWidth+x] {
				fmt.Print("⬛")
			} else {
				fmt.Print("⬜")
			}
		}
		fmt.Println()
	}
}

func (c *Chip8) LoadROM(data []byte) {
	const maxRomSize = consts.MemorySize - consts.StartAddress
	if len(data) > maxRomSize {
		panic("ROM too large")
	}

	copy(c.Memory[consts.StartAddress:], data[:])
}

func (c *Chip8) Fetch() uint16 {
	// get the first byte of the opcode
	high := uint16(c.Memory[c.PC])
	// get the second byte of the opcode
	low := uint16(c.Memory[c.PC+1])
	// << is left shift, | is bitwise OR
	// left shift move the bits x positions to the left
	// bitwise OR combines the bits of two numbers
	// 0x0100 | 0x0001 = 0x0101
	opcode := (high << 8) | low
	// increment the program counter by 2 to point to the next opcode
	c.PC += 2
	return opcode
}

func (c *Chip8) Execute(opcode uint16) {
	switch opcode {
	case 0x00E0:
		c.op00E0()
	default:
		switch opcode & 0xF000 {
		case 0x1000:
			c.op1NNN(opcode)
		}
	}
}

// 00E0 - CLS - Clear screen
func (c *Chip8) op00E0() {
	for i := range c.Screen {
		c.Screen[i] = false
	}
}

// 1NNN - JP addr - Jump to address NNN
func (c *Chip8) op1NNN(opcode uint16) {
	address := opcode & 0x0FFF
	fmt.Printf("jumping to 0x%04X\n", address)
	c.PC = address
}
