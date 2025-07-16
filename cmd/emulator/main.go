package main

import (
	"log"
	"os"
	"time"

	"github.com/breno5g/emugo-8/internal/entity"
)

func main() {
	chip := entity.NewChip8()

	// fmt.Printf("initial PC: 0x%X\n", chip.PC)

	// fmt.Println("Fontset bytes:")
	// for i := range chip.Memory[:80] {
	// 	fmt.Printf("0x%02X ", chip.Memory[i])
	// 	if (i+1)%8 == 0 {
	// 		fmt.Println()
	// 	}
	// }

	// chip.LoadROM(consts.TestROM)

	// fmt.Println("ROM bytes:")
	// for i := range consts.TestROM {
	// 	fmt.Printf("0x%02X ", consts.TestROM[i])
	// 	if (i+1)%8 == 0 {
	// 		fmt.Println()
	// 	}
	// }

	// chip.DebugScreen()

	// first byte is 0x60 and second is 0x10
	// chip.LoadROM(consts.TestROM)
	// opcode := chip.Fetch()
	// fmt.Printf("opcode: 0x%04X\n", opcode)

	// for i := range chip.Screen {
	// 	chip.Screen[i] = true
	// }

	// chip.Execute(0x00E0)

	// chip.DebugScreen()

	// chip.LoadROM(consts.TestROM)
	// opcode := chip.Fetch()
	// fmt.Printf("opcode: 0x%04X\n", opcode)
	// chip.Execute(opcode)
	// fmt.Printf("PC: 0x%04X\n", chip.PC)

	// 6XNN - LD Vx, byte - Load Vx with NN
	// chip.LoadROM(consts.TestROM)
	// opcode := chip.Fetch()
	// fmt.Printf("opcode: 0x%04X\n", opcode)
	// chip.Execute(opcode)
	// fmt.Printf("V[0]: 0x%02X\n", chip.V[0])

	// 7XNN - ADD Vx, byte - Add NN to Vx
	// LD V5, 0x10
	// chip.Execute(0x6510)
	// // ADD V5, 0x20
	// chip.Execute(0x7520)
	// fmt.Printf("V[5]: 0x%02X\n", chip.V[5])

	// // 3XNN - SE Vx, byte - Skip next instruction if Vx = NN
	// chip.V[4] = 0x42
	// fmt.Println("PC", chip.PC)
	// chip.Execute(0x3442) // SE V4, 0x42
	// fmt.Printf("PC: 0x%04X\n", chip.PC)

	// 4XNN - SNE Vx, byte - Skip next instruction if Vx != NN
	// chip.V[4] = 0x42
	// 	fmt.Printf("PC before: 0x%04X\n", chip.PC)
	// 	chip.Execute(0x4442) // SNE V4, 0x42
	// 	fmt.Printf("PC after: 0x%04X\n", chip.PC)

	// 5XY0 - SE Vx, Vy - Skip next instruction if Vx = Vy
	// chip.V[1] = 0x10
	// chip.V[2] = 0x10
	// // chip.V[2] = 0x20
	// fmt.Printf("PC before: 0x%04X\n", chip.PC)
	// chip.Execute(0x5120) // SE V1, V2
	// fmt.Printf("PC after: 0x%04X\n", chip.PC)

	// 9XY0 - SNE Vx, Vy - Skip next instruction if Vx != Vy
	// chip.V[1] = 0x10
	// // chip.V[2] = 0x20
	// chip.V[2] = 0x10
	// fmt.Printf("PC before: 0x%04X\n", chip.PC)
	// chip.Execute(0x9120) // SNE V1, V2
	// fmt.Printf("PC after: 0x%04X\n", chip.PC)

	// 2NNN - CALL addr - Call subroutine at NNN
	// chip.Execute(0x2200) // CALL 0x200
	// fmt.Printf("PC: 0x%04X\n", chip.PC)
	// fmt.Printf("SP: 0x%04X\n", chip.SP)
	// fmt.Printf("Stack: %v\n", chip.Stack)

	// // 00EE - RET - Return from subroutine
	// chip.PC = 0x200
	// chip.SP = 1
	// chip.Stack[0] = 0x100
	// chip.Execute(0x00EE) // RET
	// fmt.Printf("PC: 0x%04X\n", chip.PC)
	// fmt.Printf("SP: 0x%04X\n", chip.SP)
	// fmt.Printf("Stack: %v\n", chip.Stack)

	// ANNN - LD I, addr - Load I with NNN
	// chip.Execute(0xA123) // LD I, 0x123
	// fmt.Printf("I: 0x%04X\n", chip.I)

	// DXYN - DRW Vx, Vy, N - Draw sprite at Vx, Vy with N rows

	// chip.Memory[0x300] = 0b11111111
	// chip.I = 0x300

	// chip.V[0] = 0 // X
	// chip.V[1] = 0 // Y

	// chip.Execute(0xD011) // DRW V0, V1, 1

	// chip.DebugScreen()

	// // CXNN - RND Vx, byte - Generate random number and AND with NN
	// chip.Execute(0xC0FF) // RND V0, 0xFF
	// fmt.Printf("V[0]: 0x%02X\n", chip.V[0])

	// fmt.Printf("PC: 0x%04X\n", byte(177&0xF0))

	// // FX1E - ADD I, Vx - Add Vx to I
	// chip.I = 0x300
	// chip.V[5] = 0x20
	// chip.Execute(0xF51E) // ADD I, V5
	// fmt.Printf("I: 0x%04X\n", chip.I)

	// // FX07 - LD Vx, DT - Load Vx with DT
	// chip.DT = 0x42
	// chip.Execute(0xF307) // LD V3, DT
	// fmt.Printf("V[3]: 0x%02X\n", chip.V[3])

	// // FX15 - LD DT, Vx - Load DT with Vx
	// chip.V[3] = 0x42
	// chip.Execute(0xF315) // LD DT, V3
	// fmt.Printf("DT: 0x%02X\n", chip.DT)

	// // FX18 - LD ST, Vx - Load ST with Vx
	// chip.V[7] = 0x55
	// chip.Execute(0xF718) // LD ST, V7
	// fmt.Printf("ST: 0x%02X\n", chip.ST)

	// // FX29 - LD F, Vx - Set I with Vx sprite
	// chip.V[1] = 0x0A // A
	// chip.Execute(0xF129)
	// fmt.Printf("I: 0x%02X\n", chip.I)

	// // FX33 - LD B, Vx - Store BCD representation of Vx in memory
	// chip.V[2] = 0x66
	// chip.Execute(0xF233)
	// fmt.Printf("Memory: %v\n", chip.Memory[chip.I:chip.I+3])

	// // FX55 - LD [I], Vx - Store registers V0 to Vx in memory
	// chip.V[0] = 0x11
	// chip.V[1] = 0x22
	// chip.V[2] = 0x33
	// chip.V[3] = 0x44
	// chip.Execute(0xF355)
	// fmt.Printf("Memory: %v\n", chip.Memory[chip.I:chip.I+4])

	// // FX65 - LD Vx, [I] - Read registers V0 to Vx from memory
	// chip.Memory[0x300] = 0x11
	// chip.Memory[0x301] = 0x22
	// chip.Memory[0x302] = 0x33
	// chip.Memory[0x303] = 0x44
	// chip.I = 0x300
	// chip.Execute(0xF365)
	// fmt.Printf("V: %v\n", chip.V[:4])

	// romData, err := os.ReadFile("assets/roms/ibm-logo.ch8")
	// romData, err := os.ReadFile("assets/roms/chip8-logo.ch8")
	// romData, err := os.ReadFile("assets/roms/corax+.ch8")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// chip.LoadROM(romData)
	// RunCycles(chip, 1000)
	// chip.DebugScreen()

	// // 8XY1 - OR Vx, Vy - Set Vx = Vx OR Vy
	// chip.V[0] = 0b10101010
	// chip.V[1] = 0b01010101
	// chip.Execute(0x8011)
	// fmt.Printf("V[0]: 0b%08b\n", chip.V[0])

	// // 8XY2 - AND Vx, Vy - Set Vx = Vx AND Vy
	// chip.V[0] = 0b10101011
	// chip.V[1] = 0b01010101
	// chip.Execute(0x8012)
	// fmt.Printf("V[0]: 0b%08b\n", chip.V[0])

	// // 8XY3 - XOR Vx, Vy - Set Vx = Vx XOR Vy
	// chip.V[0] = 0b11101011
	// chip.V[1] = 0b01011101
	// chip.Execute(0x8013)
	// fmt.Printf("V[0]: 0b%08b\n", chip.V[0])

	// // 8XY4 - ADD Vx, Vy - Sum Vx with Vy and add overflow to carry
	// chip.V[3] = 0xF0
	// chip.V[4] = 0x30
	// chip.Execute(0x8344)

	// fmt.Printf("V[3]: 0x%02X, V[15]: 0x%02X\n", chip.V[3], chip.V[0xF])

	// // 8XY5 SUB Vx, Vy - Sub Vx with Vy and add underflow to borrow
	// chip.V[3] = 10
	// chip.V[4] = 20
	// chip.Execute(0x8345)
	// fmt.Printf("V[3]: 0x%02X, V[F]: %d\n", chip.V[3], chip.V[0xF])

	// // 8XY6 SHR Vx - Shift right Vx and set LSB to VF
	// chip.V[2] = 0b00000011 // 3
	// chip.Execute(0x8206)
	// fmt.Printf("V2: %d, VF: %d\n", chip.V[2], chip.V[0xF])

	// chip.V[3] = 0b00000100 // 4
	// chip.Execute(0x8306)
	// fmt.Printf("V3: %d, VF: %d", chip.V[3], chip.V[0xF])

	// // 8XYE SHL Vx - Shift left Vx and set MSB to VF
	// chip.V[2] = 0b10000001 // 129
	// chip.Execute(0x820E)
	// fmt.Printf("V2: %d, VF: %d\n", chip.V[2], chip.V[0xF])

	// chip.V[3] = 0b01000001 // 65
	// chip.Execute(0x830E)
	// fmt.Printf("V3: %d, VF: %d", chip.V[3], chip.V[0xF])

	chip.LoadFontSet() // carrega fontes na memória
	// romData, err := os.ReadFile("assets/roms/ibm-logo.ch8")
	// romData, err := os.ReadFile("assets/roms/corax+.ch8")
	// romData, err := os.ReadFile("assets/roms/flags.ch8")
	// romData, err := os.ReadFile("assets/roms/gradsim.ch8")
	romData, err := os.ReadFile("assets/roms/beep.ch8")
	if err != nil {
		log.Fatal(err)
	}
	chip.LoadROM(romData)

	const fps = 60
	ticker := time.NewTicker(time.Second / fps)
	defer ticker.Stop()

	chip.EventLoop()
}
