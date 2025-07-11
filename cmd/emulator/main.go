package main

import (
	"fmt"

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

	// 3XNN - SE Vx, byte - Skip next instruction if Vx = NN
	chip.V[4] = 0x42
	fmt.Println("PC", chip.PC)
	chip.Execute(0x3442) // SE V4, 0x42
	fmt.Printf("PC: 0x%04X\n", chip.PC)
}
