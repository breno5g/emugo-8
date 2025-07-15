package entity

import (
	"fmt"
	"math/rand"

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
	case 0x00EE:
		c.op00EE()

	default:
		switch opcode & 0xF000 {
		case 0x1000:
			c.op1NNN(opcode)
		case 0x6000:
			c.op6XNN(opcode)
		case 0x7000:
			c.op7XNN(opcode)
		case 0x3000:
			c.op3XNN(opcode)
		case 0x4000:
			c.op4XNN(opcode)
		case 0x5000:
			if (opcode & consts.NMask) != 0 {
				panic("Invalid opcode")
			}
			c.op5XY0(opcode)
		case 0x9000:
			if (opcode & consts.NMask) != 0 {
				panic("Invalid opcode")
			}
			c.op9XY0(opcode)
		case 0x2000:
			c.op2NNN(opcode)
		case 0xA000:
			c.opANNN(opcode)
		case 0xD000:
			c.opDXYN(opcode)
		case 0xC000:
			c.opCXNN(opcode)
		case 0xF000:
			switch opcode & 0x00FF {
			case 0x1E:
				c.opFX1E(opcode)
			case 0x07:
				c.opFX07(opcode)
			case 0x15:
				c.opFX15(opcode)
			case 0x18:
				c.opFX18(opcode)
			case 0x29:
				c.opFX29(opcode)
			case 0x33:
				c.opFX33(opcode)
			case 0x55:
				c.opFX55(opcode)
			}
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
	address := opcode & consts.NNNMask
	c.PC = address
}

// 6XNN - LD Vx, byte - Load Vx with NN
func (c *Chip8) op6XNN(opcode uint16) {
	// get the register
	x := (opcode & consts.XMask) >> 8
	// get the value
	nn := opcode & consts.NNask
	// load the value into the register
	c.V[x] = byte(nn)
}

// 7XNN - ADD Vx, byte - Add NN to Vx
func (c *Chip8) op7XNN(opcode uint16) {
	// get the register
	x := (opcode & consts.XMask) >> 8
	// get the value
	nn := opcode & consts.NNask
	// add the value to the register
	c.V[x] += byte(nn)
}

// 3XNN - SE Vx, byte - Skip next instruction if Vx = NN
func (c *Chip8) op3XNN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	nn := opcode & consts.NNask
	if c.V[x] == byte(nn) {
		c.PC += 2
	}
}

// 4XNN - SNE Vx, byte - Skip next instruction if Vx != NN
func (c *Chip8) op4XNN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	nn := opcode & consts.NNask
	if c.V[x] != byte(nn) {
		c.PC += 2
	}
}

// 5XY0 - SE Vx, Vy - Skip next instruction if Vx = Vy
func (c *Chip8) op5XY0(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	if c.V[x] == c.V[y] {
		c.PC += 2
	}
}

// 9XY0 - SNE Vx, Vy - Skip next instruction if Vx != Vy
func (c *Chip8) op9XY0(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	if c.V[x] != c.V[y] {
		c.PC += 2
	}
}

// 2NNN - CALL addr - Call subroutine at NNN
func (c *Chip8) op2NNN(opcode uint16) {
	address := opcode & consts.NNNMask
	// push the current PC to the stack
	c.Stack[c.SP] = c.PC
	// increment the stack pointer
	c.SP++
	// set the PC to the address
	c.PC = address
}

// 00EE - RET - Return from subroutine
func (c *Chip8) op00EE() {
	// decrement the stack pointer
	c.SP--
	// set the PC to the address in the stack
	c.PC = c.Stack[c.SP]
}

// ANNN - LD I, addr - Load I with NNN
func (c *Chip8) opANNN(opcode uint16) {
	addr := opcode & consts.NNNMask
	c.I = addr
}

// DXYN - DRW Vx, Vy, N - Draw sprite at Vx, Vy with N rows
func (c *Chip8) opDXYN(opcode uint16) {
	// get the registers
	x := c.V[(opcode&consts.XMask)>>8]
	y := c.V[(opcode&consts.YMask)>>4]
	// get the number of rows
	n := opcode & consts.NMask

	c.V[0xF] = 0 // VF = 0 (without colision)

	for row := range n {
		// get the sprite byte
		spriteByte := c.Memory[c.I+row]
		for col := range uint16(8) {
			// calculate the pixel coordinates
			pixelX := (uint16(x) + col) % consts.DisplayWidth
			pixelY := (uint16(y) + row) % consts.DisplayHeight
			idx := pixelY*consts.DisplayWidth + pixelX

			// get the sprite bit
			spriteBit := (spriteByte >> (7 - col)) & 1
			if spriteBit == 1 {
				// check if the pixel is already on
				if c.Screen[idx] {
					// VF = 1 (colision happened)
					c.V[0xF] = 1
				}
				// toggle the pixel
				c.Screen[idx] = !c.Screen[idx]
			}
		}
	}
}

func (c *Chip8) opCXNN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	nn := byte(opcode & 0x00FF)

	random := byte(rand.Intn(256)) // 0x00–0xFF
	c.V[x] = random & nn
}

func (c *Chip8) opFX1E(opcode uint16) {
	x := (opcode & consts.XMask) >> 8

	c.I += uint16(c.V[x])
}

func (c *Chip8) opFX07(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.V[x] = c.DT
}

func (c *Chip8) opFX15(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.DT = c.V[x]
}

func (c *Chip8) opFX18(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.ST = c.V[x]
}

func (c *Chip8) opFX29(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	digit := c.V[x] & 0x0F

	c.I = uint16(digit) * 5
}

func (c *Chip8) opFX33(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	val := c.V[x]

	c.Memory[c.I] = val / 100
	c.Memory[c.I+1] = (val / 10) % 10
	c.Memory[c.I+2] = val % 10
}

func (c *Chip8) opFX55(opcode uint16) {
	x := (opcode & consts.XMask) >> 8

	for i := uint16(0); i <= x; i++ {
		c.Memory[c.I+i] = c.V[i]
	}
}
