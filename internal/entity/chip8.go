package entity

import "github.com/breno5g/emugo-8/internal/consts"

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
