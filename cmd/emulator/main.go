package main

import (
	"fmt"

	"github.com/breno5g/emugo-8/internal/consts"
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

	chip.LoadROM(consts.TestROM)

	fmt.Println("ROM bytes:")
	for i := range consts.TestROM {
		fmt.Printf("0x%02X ", consts.TestROM[i])
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}

	// chip.DebugScreen()
}
