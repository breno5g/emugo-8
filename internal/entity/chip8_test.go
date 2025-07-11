package entity_test

import (
	"testing"

	"github.com/breno5g/emugo-8/internal/consts"
	"github.com/breno5g/emugo-8/internal/entity"
)

func TestNewChip8Initialization(t *testing.T) {
	chip := entity.NewChip8()

	if chip.PC != consts.StartAddress {
		t.Errorf("PC inicial incorreto: esperado %X, obtido %X", consts.StartAddress, chip.PC)
	}

	for i, v := range consts.FontSet {
		if chip.Memory[i] != v {
			t.Errorf("Fonte não carregada corretamente na memória no índice %d", i)
		}
	}
}
