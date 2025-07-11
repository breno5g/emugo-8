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
