package entity_test

import (
	"testing"

	"github.com/breno5g/emugo-8/internal/consts"
	"github.com/breno5g/emugo-8/internal/entity"
)

func TestNewChip8Initialization(t *testing.T) {
	chip := entity.NewChip8()

	if chip.PC != consts.StartAddress {
		t.Errorf("Initial PC should be 0x%X, but it is 0x%X", consts.StartAddress, chip.PC)
	}

	for i, v := range consts.FontSet {
		if chip.Memory[i] != v {
			t.Errorf("Fontset not loaded correctly at index %d", i)
		}
	}
}

func TestLoadROM(t *testing.T) {
	chip := entity.NewChip8()

	chip.LoadROM(consts.TestROM)

	start := consts.StartAddress
	for i, b := range consts.TestROM {
		if chip.Memory[start+i] != b {
			t.Errorf("ROM not loaded correctly at address 0x%X: expected 0x%X, got 0x%X", start+i, b, chip.Memory[start+i])
		}
	}
}

func TestFetch(t *testing.T) {
	chip := entity.NewChip8()

	// Simulating a fake instruction 0xA2F0 in memory
	chip.Memory[0x200] = 0xA2
	chip.Memory[0x201] = 0xF0

	opcode := chip.Fetch()

	if opcode != 0xA2F0 {
		t.Errorf("Wrong opcode fetched: expected 0xA2F0, got 0x%04X", opcode)
	}

	if chip.PC != 0x202 {
		t.Errorf("PC was not incremented correctly: expected 0x202, got 0x%04X", chip.PC)
	}
}

func TestOpcode00E0_CLS(t *testing.T) {
	chip := entity.NewChip8()

	// Turn on all pixels
	for i := range chip.Screen {
		chip.Screen[i] = true
	}

	// Execute the CLS opcode
	chip.Execute(0x00E0)

	// Check if the screen was cleared
	for i, pixel := range chip.Screen {
		if pixel {
			t.Errorf("Pixel %d was not cleared", i)
		}
	}
}

func TestOpcode1NNN_Jump(t *testing.T) {
	chip := entity.NewChip8()

	chip.Execute(0x1234)

	if chip.PC != 0x0234 {
		t.Errorf("JP falhou: esperado PC = 0x0234, obtido 0x%04X", chip.PC)
	}
}
