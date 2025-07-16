package entity

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

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

	WaitingForKey bool
	WaitingReg    byte
}

func NewChip8() *Chip8 {
	c := &Chip8{
		PC:            consts.StartAddress,
		WaitingForKey: false,
		WaitingReg:    0,
	}
	c.LoadFontSet()
	rand.Seed(time.Now().UnixNano())
	return c
}

func (c *Chip8) LoadFontSet() {
	copy(c.Memory[consts.FontStartAddress:], consts.FontSet[:])
}

// DebugScreen renders the current screen state to the console.
// It uses a strings.Builder to build the entire screen in a buffer
// before printing once, which prevents flickering.
func (c *Chip8) DebugScreen() {
	var screen strings.Builder
	screen.Grow(consts.DisplaySize + consts.DisplayHeight)

	// Move cursor to home position. This overwrites the previous frame
	// instead of clearing the screen, which is faster.
	screen.WriteString("\033[H")

	for y := 0; y < consts.DisplayHeight; y++ {
		for x := 0; x < consts.DisplayWidth; x++ {
			if c.Screen[y*consts.DisplayWidth+x] {
				screen.WriteString("⬛")
			} else {
				screen.WriteString("⬜")
			}
		}
		screen.WriteString("\n")
	}
	fmt.Print(screen.String())
}

// Load ROM data into memory, starting at 0x200.
func (c *Chip8) LoadROM(data []byte) {
	const maxRomSize = consts.MemorySize - consts.StartAddress
	if len(data) > maxRomSize {
		panic("ROM too large")
	}
	copy(c.Memory[consts.StartAddress:], data[:])
}

// Fetch reads the 2-byte opcode from memory at the PC and advances the PC.
func (c *Chip8) Fetch() uint16 {
	// Garante que o PC não ultrapasse os limites da memória.
	if c.PC+1 >= consts.MemorySize {
		// Ação de erro, como pausar ou parar o emulador.
		// Por simplicidade, vamos apenas logar e parar.
		fmt.Printf("Erro: PC (0x%X) fora dos limites da memória.\n", c.PC)
		// Para evitar pânico, podemos entrar em um loop infinito ou definir um estado de erro.
		// Aqui, vamos apenas retornar um opcode NOP (0x0000) para evitar crash.
		return 0x0000
	}
	high := uint16(c.Memory[c.PC])
	low := uint16(c.Memory[c.PC+1])
	opcode := (high << 8) | low
	c.PC += 2
	return opcode
}

// Tick executes one CPU cycle: fetch, decode, and execute.
// It also handles the waiting-for-key state.
func (c *Chip8) Tick() {
	// If the emulator is waiting for a key press, halt execution.
	if c.WaitingForKey {
		for key, pressed := range c.Keys {
			if pressed {
				c.V[c.WaitingReg] = byte(key)
				c.WaitingForKey = false
				break
			}
		}
		return
	}

	opcode := c.Fetch()
	c.Execute(opcode)
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
				return
			}
			c.op5XY0(opcode)
		case 0x8000:
			switch opcode & 0x000F {
			case 0x0:
				c.op8XY0(opcode)
			case 0x1:
				c.op8XY1(opcode)
			case 0x2:
				c.op8XY2(opcode)
			case 0x3:
				c.op8XY3(opcode)
			case 0x4:
				c.op8XY4(opcode)
			case 0x5:
				c.op8XY5(opcode)
			case 0x6:
				c.op8XY6(opcode)
			case 0x7:
				c.op8XY7(opcode)
			case 0xE:
				c.op8XYE(opcode)
			}
		case 0x9000:
			if (opcode & consts.NMask) != 0 {
				return
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
		case 0xB000:
			c.opBNNN(opcode)
		case 0xF000:
			switch opcode & 0x00FF {
			case 0x07:
				c.opFX07(opcode)
			case 0x0A:
				c.opFX0A(opcode)
			case 0x15:
				c.opFX15(opcode)
			case 0x18:
				c.opFX18(opcode)
			case 0x1E:
				c.opFX1E(opcode)
			case 0x29:
				c.opFX29(opcode)
			case 0x33:
				c.opFX33(opcode)
			case 0x55:
				c.opFX55(opcode)
			case 0x65:
				c.opFX65(opcode)
			}
		case 0xE000:
			switch opcode & 0x00FF {
			case 0x9E:
				c.opEX9E(opcode)
			case 0xA1:
				c.opEXA1(opcode)
			}
		}
	}
}

// 00E0 - CLS - Clear screen
func (c *Chip8) op00E0() {
	// Reset the array to clear screen
	c.Screen = [consts.DisplaySize]bool{}
}

// 00EE - RET - Return from subroutine
func (c *Chip8) op00EE() {
	c.SP--
	c.PC = c.Stack[c.SP]
}

// 1NNN - JP addr - Jump to address NNN
func (c *Chip8) op1NNN(opcode uint16) {
	address := opcode & consts.NNNMask
	c.PC = address
}

// 2NNN - CALL addr - Call subroutine at NNN
func (c *Chip8) op2NNN(opcode uint16) {
	address := opcode & consts.NNNMask
	c.Stack[c.SP] = c.PC
	c.SP++
	c.PC = address
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

// 6XNN - LD Vx, byte - Load Vx with NN
func (c *Chip8) op6XNN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	nn := opcode & consts.NNask
	c.V[x] = byte(nn)
}

// 7XNN - ADD Vx, byte - Add NN to Vx
func (c *Chip8) op7XNN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	nn := opcode & consts.NNask
	c.V[x] += byte(nn)
}

// 8XY0 - LD Vx, Vy - Set Vx = Vy
func (c *Chip8) op8XY0(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	c.V[x] = c.V[y]
}

// 8XY1 - OR Vx, Vy - Set Vx = Vx OR Vy
func (c *Chip8) op8XY1(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	c.V[x] |= c.V[y]
}

// 8XY2 - AND Vx, Vy - Set Vx = Vx AND Vy
func (c *Chip8) op8XY2(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	c.V[x] &= c.V[y]
}

// 8XY3 - XOR Vx, Vy - Set Vx = Vx XOR Vy
func (c *Chip8) op8XY3(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	c.V[x] ^= c.V[y]
}

// 8XY4 - ADD Vx, Vy - Set Vx = Vx + Vy, set VF = carry
func (c *Chip8) op8XY4(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	sum := uint16(c.V[x]) + uint16(c.V[y])
	if sum > 0xFF {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[x] = byte(sum)
}

// 8XY5 - SUB Vx, Vy - Set Vx = Vx - Vy, set VF = NOT borrow
func (c *Chip8) op8XY5(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	if c.V[x] >= c.V[y] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[x] -= c.V[y]
}

// 8XY6 - SHR Vx {, Vy} - Set Vx = Vx >> 1, set VF = LSB
func (c *Chip8) op8XY6(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.V[0xF] = c.V[x] & 0x1
	c.V[x] >>= 1
}

// 8XY7 - SUBN Vx, Vy - Set Vx = Vy - Vx, set VF = NOT borrow
func (c *Chip8) op8XY7(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	if c.V[y] >= c.V[x] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[x] = c.V[y] - c.V[x]
}

// 8XYE - SHL Vx {, Vy} - Set Vx = Vx << 1, set VF = MSB
func (c *Chip8) op8XYE(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.V[0xF] = (c.V[x] & 0x80) >> 7
	c.V[x] <<= 1
}

// 9XY0 - SNE Vx, Vy - Skip next instruction if Vx != Vy
func (c *Chip8) op9XY0(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	if c.V[x] != c.V[y] {
		c.PC += 2
	}
}

// ANNN - LD I, addr - Load I with NNN
func (c *Chip8) opANNN(opcode uint16) {
	addr := opcode & consts.NNNMask
	c.I = addr
}

// BNNN - JP V0, addr - jump to NNN + V0
func (c *Chip8) opBNNN(opcode uint16) {
	addr := opcode & consts.NNNMask
	c.PC = addr + uint16(c.V[0])
}

// CXNN - RND Vx, byte - Generate random number and AND with NN
func (c *Chip8) opCXNN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	nn := byte(opcode & consts.NNask)
	random := byte(rand.Intn(256))
	c.V[x] = random & nn
}

// DXYN - DRW Vx, Vy, N - Draw sprite at Vx, Vy with N rows
func (c *Chip8) opDXYN(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	y := (opcode & consts.YMask) >> 4
	n := opcode & consts.NMask

	vx := uint16(c.V[x])
	vy := uint16(c.V[y])
	c.V[0xF] = 0 // VF = 0 (without collision)

	// loop from 0 to n-1
	for row := uint16(0); row < n; row++ {
		spriteByte := c.Memory[c.I+row]
		pixelY := (vy + row) % consts.DisplayHeight

		// loop from 0 to 7
		for col := uint16(0); col < 8; col++ {
			spritePixel := (spriteByte >> (7 - col)) & 1
			if spritePixel == 1 {
				pixelX := (vx + col) % consts.DisplayWidth
				idx := pixelY*consts.DisplayWidth + pixelX

				if c.Screen[idx] {
					c.V[0xF] = 1 // VF = 1 (collision happened)
				}
				c.Screen[idx] = !c.Screen[idx] // toggle the pixel (XOR)
			}
		}
	}
}

// EX9E - SKP Vx - skip next instruction if Vx is pressed
func (c *Chip8) opEX9E(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	key := c.V[x] & 0x0F
	if c.Keys[key] {
		c.PC += 2
	}
}

// EXA1 - SKNP Vx - skip next instruction if Vx is not pressed
func (c *Chip8) opEXA1(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	key := c.V[x] & 0x0F
	if !c.Keys[key] {
		c.PC += 2
	}
}

// FX07 - LD Vx, DT - Load Vx with DT
func (c *Chip8) opFX07(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.V[x] = c.DT
}

// FX0A - LD Vx, K - pause execution until key pressed
func (c *Chip8) opFX0A(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.WaitingForKey = true
	c.WaitingReg = byte(x)
}

// FX15 - LD DT, Vx - Load DT with Vx
func (c *Chip8) opFX15(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.DT = c.V[x]
}

// FX18 - LD ST, Vx - Load ST with Vx
func (c *Chip8) opFX18(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.ST = c.V[x]
}

// FX1E - ADD I, Vx - Add Vx to I
func (c *Chip8) opFX1E(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	c.I += uint16(c.V[x])
}

// FX29 - LD F, Vx - Set I to the location of the sprite for the character in Vx
func (c *Chip8) opFX29(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	digit := c.V[x] & 0x0F
	c.I = consts.FontStartAddress + (uint16(digit) * 5)
}

// FX33 - LD B, Vx - Store BCD representation of Vx in memory
func (c *Chip8) opFX33(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	val := c.V[x]
	c.Memory[c.I] = val / 100
	c.Memory[c.I+1] = (val / 10) % 10
	c.Memory[c.I+2] = val % 10
}

// FX55 - LD [I], Vx - Store registers V0 to Vx in memory
func (c *Chip8) opFX55(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	for i := uint16(0); i <= x; i++ {
		c.Memory[c.I+i] = c.V[i]
	}
}

// FX65 - LD Vx, [I] - Read registers V0 to Vx from memory
func (c *Chip8) opFX65(opcode uint16) {
	x := (opcode & consts.XMask) >> 8
	for i := uint16(0); i <= x; i++ {
		c.V[i] = c.Memory[c.I+i]
	}
}

// UpdateTimers decrements the delay and sound timers.
func (c *Chip8) UpdateTimers() {
	if c.DT > 0 {
		c.DT--
	}
	if c.ST > 0 {
		c.ST--
	}
}

// // EventLoop is the main emulator loop.
// func (c *Chip8) EventLoop() {
// 	// A typical Chip-8 CPU speed is around 500-700Hz.
// 	// Screen and timers update at 60Hz.
// 	// We use separate tickers to manage this timing.
// 	cpuTicker := time.NewTicker(time.Second / 700) // ~700 cycles per second
// 	defer cpuTicker.Stop()

// 	timerTicker := time.NewTicker(time.Second / 60) // 60Hz for timers and screen
// 	defer timerTicker.Stop()

// 	for {
// 		select {
// 		case <-cpuTicker.C:
// 			// handleInput(c) // TODO: add input handler method [@breno5g]
// 			c.Tick()
// 		case <-timerTicker.C:
// 			c.UpdateTimers()
// 			c.DebugScreen()
// 			if c.ST > 0 {
// 				fmt.Print("\a")
// 			}
// 		}
// 	}
// }
