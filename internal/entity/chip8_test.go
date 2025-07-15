package entity_test

import (
	"math/rand"
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

func TestOpcode1NNN_JP(t *testing.T) {
	chip := entity.NewChip8()

	chip.Execute(0x1234)

	if chip.PC != 0x0234 {
		t.Errorf("Jump failed: expected PC = 0x0234, got 0x%04X", chip.PC)
	}
}

func TestOpcode6XNN_LD(t *testing.T) {
	chip := entity.NewChip8()

	// LD V5, 0xAB
	chip.Execute(0x65AB)

	if chip.V[5] != 0xAB {
		t.Errorf("Expected V5 = 0xAB, got 0x%02X", chip.V[5])
	}
}

func TestOpcode7XNN_ADD(t *testing.T) {
	chip := entity.NewChip8()

	// LD V5, 0x10
	chip.Execute(0x6510)
	// ADD V5, 0x20
	chip.Execute(0x7520)

	if chip.V[5] != 0x30 {
		t.Errorf("Expected V5 = 0x30, got 0x%02X", chip.V[5])
	}
}

func TestOpcode3XNN_SE(t *testing.T) {
	chip := entity.NewChip8()
	chip.PC = 0x200
	chip.V[5] = 0xAA

	// SE V5, 0xAA → true condition (AA == AA) → PC += 2
	chip.Execute(0x35AA)

	if chip.PC != 0x202 {
		t.Errorf("Expected PC = 0x202 after true comparison, got 0x%04X", chip.PC)
	}

	chip.PC = 0x200
	chip.V[5] = 0x10

	// SE V5, 0xAA → false condition (10 ≠ AA) → doesn't change PC
	chip.Execute(0x35AA)

	if chip.PC != 0x200 {
		t.Errorf("Expected PC = 0x200 after false comparison, got 0x%04X", chip.PC)
	}
}

func TestOpcode4XNN_SNE(t *testing.T) {
	chip := entity.NewChip8()
	chip.PC = 0x200
	chip.V[3] = 0x55

	// SNE V3, 0x42 → true condition (55 ≠ 42) → PC += 2
	chip.Execute(0x4342)
	if chip.PC != 0x202 {
		t.Errorf("Expected PC = 0x202 after true comparison, got 0x%04X", chip.PC)
	}

	chip.PC = 0x200
	chip.V[3] = 0x42

	// SNE V3, 0x42 → false condition (42 == 42) → doesn't change PC
	chip.Execute(0x4342)
	if chip.PC != 0x200 {
		t.Errorf("Expected PC = 0x200 after false comparison, got 0x%04X", chip.PC)
	}
}

func TestOpcode5XY0_SE(t *testing.T) {
	chip := entity.NewChip8()
	chip.PC = 0x200
	chip.V[1] = 0x10
	chip.V[2] = 0x10

	// SE V1, V2 → true condition (10 == 10) → PC += 2
	chip.Execute(0x5120)
	if chip.PC != 0x202 {
		t.Errorf("Expected PC = 0x202 after true comparison, got 0x%04X", chip.PC)
	}

	chip.PC = 0x200
	chip.V[1] = 0x10
	chip.V[2] = 0x20

	// SE V1, V2 → false condition (10 ≠ 20) → doesn't change PC
	chip.Execute(0x5120)
	if chip.PC != 0x200 {
		t.Errorf("Expected PC = 0x200 after false comparison, got 0x%04X", chip.PC)
	}
}

func TestOpcode9XY0_SNE(t *testing.T) {
	chip := entity.NewChip8()
	chip.PC = 0x200
	chip.V[1] = 0x10
	chip.V[2] = 0x20

	// SNE V1, V2 → true condition (10 ≠ 20) → PC += 2
	chip.Execute(0x9120)
	if chip.PC != 0x202 {
		t.Errorf("Expected PC = 0x202 after true comparison, got 0x%04X", chip.PC)
	}

	chip.PC = 0x200
	chip.V[1] = 0x10
	chip.V[2] = 0x10

	// SNE V1, V2 → false condition (10 == 10) → doesn't change PC
	chip.Execute(0x9120)
	if chip.PC != 0x200 {
		t.Errorf("Expected PC = 0x200 after false comparison, got 0x%04X", chip.PC)
	}
}

func TestOpcode2NNN_CALL(t *testing.T) {
	chip := entity.NewChip8()
	chip.PC = 0x200

	chip.Execute(0x2300)

	if chip.PC != 0x300 {
		t.Errorf("Expected PC = 0x300 after call, got 0x%04X", chip.PC)
	}

	if chip.SP != 1 {
		t.Errorf("Expected SP = 1 after call, got %d", chip.SP)
	}

	if chip.Stack[0] != 0x200 {
		t.Errorf("Expected Stack[0] = 0x200 after call, got 0x%04X", chip.Stack[0])
	}
}

func TestOpcode00EE_RET(t *testing.T) {
	chip := entity.NewChip8()
	chip.PC = 0x200
	chip.SP = 1
	chip.Stack[0] = 0x100

	chip.Execute(0x00EE)

	if chip.PC != 0x100 {
		t.Errorf("Expected PC = 0x100 after return, got 0x%04X", chip.PC)
	}

	if chip.SP != 0 {
		t.Errorf("Expected SP = 0 after return, got %d", chip.SP)
	}
}

func TestOpcodeANNN_LD(t *testing.T) {
	chip := entity.NewChip8()

	chip.Execute(0xA123)

	if chip.I != 0x123 {
		t.Errorf("Expected I = 0x123 after load, got 0x%04X", chip.I)
	}
}

func TestOpcodeDXYN_DrawSinglePixel(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[0] = 0 // x = 0
	chip.V[1] = 0 // y = 0
	chip.I = 0x300
	chip.Memory[0x300] = 0b10000000 // 1 pixel to the left (0x80)
	chip.Execute(0xD011)            // DRW V0, V1, 1

	if !chip.Screen[0] {
		t.Error("Pixel (0,0) need to be on")
	}
	if chip.V[0xF] != 0 {
		t.Error("VF should be 0 (without colision)")
	}

	// Execute again: pixel will turn off (colision)
	chip.Execute(0xD011)
	if chip.Screen[0] {
		t.Error("Pixel (0,0) should have been turned off")
	}
	if chip.V[0xF] != 1 {
		t.Error("VF should be 1 (colision happened)")
	}
}

func TestOpcodeCXNN_RND(t *testing.T) {
	rand.Seed(42) // seed to generate deterministic value

	chip := entity.NewChip8()
	chip.Execute(0xC3F0) // V3 = rand() & 0xF0

	v := chip.V[3]

	// first value 0xB0 (176)
	expected := byte(176 & 0xF0)

	if v != expected {
		t.Errorf("Expected 0x%04X after call, got 0x%04X", expected, v)
	}
}

func TestOpcodeFX1E_AddI(t *testing.T) {
	chip := entity.NewChip8()
	chip.I = 0x300
	chip.V[5] = 0x20

	chip.Execute(0xF51E)

	if chip.I != 0x320 {
		t.Errorf("Expected I = 0x320 after add, got 0x%04X", chip.I)
	}
}

func TestOpcodeFX07_LD_Vx_DT(t *testing.T) {
	chip := entity.NewChip8()
	chip.DT = 0x42
	chip.V[3] = 0x00

	chip.Execute(0xF307) // V3 = DT

	if chip.V[3] != 0x42 {
		t.Errorf("Expected V3 = 0x42 after load, got 0x%02X", chip.V[3])
	}
}

func TestOpcodeFX15_LD_DT_Vx(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[4] = 0x7F
	chip.DT = 0x00

	chip.Execute(0xF415) // DT = V4

	if chip.DT != 0x7F {
		t.Errorf("Expected DT = 0x7F, got 0x%02X", chip.DT)
	}
}

func TestOpcodeFX18_LD_ST_Vx(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[7] = 0x55
	chip.ST = 0x00

	chip.Execute(0xF718) // ST = V7

	if chip.ST != 0x55 {
		t.Errorf("Expected ST = 0x55, got 0x%02X", chip.ST)
	}
}

func TestOpcodeFX29_LD_F_Sprite(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[1] = 0x0A // A

	chip.Execute(0xF129) // I = sprite(V1) (A)

	expected := uint16(0x0A * 5) // 5 bytes per sprite

	if chip.I != expected {
		t.Errorf("Expected I = 0x%04X, got 0x%04X", expected, chip.I)
	}
}

func TestOpcodeFX33_BCD(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[4] = 154
	chip.I = 0x300

	chip.Execute(0xF433) // LD B, V4 (154) → BCD (154 = 15*100 + 5*10 + 4)

	if chip.Memory[0x300] != 1 {
		t.Errorf("Expected Memory[I] = 1, got %d", chip.Memory[0x300])
	}
	if chip.Memory[0x301] != 5 {
		t.Errorf("Expected Memory[I+1] = 5, got %d", chip.Memory[0x301])
	}
	if chip.Memory[0x302] != 4 {
		t.Errorf("Expected Memory[I+2] = 4, got %d", chip.Memory[0x302])
	}
}

func TestOpcodeFX55_LD_Memory_Vx(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[0] = 0xAA
	chip.V[1] = 0xBB
	chip.V[2] = 0xCC
	chip.I = 0x300

	chip.Execute(0xF255) // V0–V2 → Memory[0x300–0x302]

	if chip.Memory[0x300] != 0xAA || chip.Memory[0x301] != 0xBB || chip.Memory[0x302] != 0xCC {
		t.Errorf("Expected [0xAA, 0xBB, 0xCC], got [%02X, %02X, %02X]",
			chip.Memory[0x300], chip.Memory[0x301], chip.Memory[0x302])
	}
}

func TestOpcodeFX65_LD_Vx_Memory(t *testing.T) {
	chip := entity.NewChip8()
	chip.Memory[0x300] = 0x11
	chip.Memory[0x301] = 0x22
	chip.Memory[0x302] = 0x33
	chip.I = 0x300

	chip.Execute(0xF265) // LD V0–V2, [I]

	if chip.V[0] != 0x11 || chip.V[1] != 0x22 || chip.V[2] != 0x33 {
		t.Errorf("Expected [0x11, 0x22, 0x33], got [%02X, %02X, %02X]",
			chip.V[0], chip.V[1], chip.V[2])
	}
}

func TestOpcode8XY1_OR(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[1] = 0x0F
	chip.V[2] = 0xF0

	chip.Execute(0x8121) // V1 = V1 | V2

	if chip.V[1] != 0xFF {
		t.Errorf("Expected V1 = 0xFF, got 0x%02X", chip.V[1])
	}
}

func TestOpcode8XY2_AND(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[3] = 0x5A
	chip.V[4] = 0xF0

	chip.Execute(0x8342) // V3 = V3 & V4

	if chip.V[3] != 0x50 {
		t.Errorf("Expected V3 = 0x50, got 0x%02X", chip.V[3])
	}
}

func TestOpcode8XY3_XOR(t *testing.T) {
	chip := entity.NewChip8()
	chip.V[5] = 0xAA
	chip.V[6] = 0x0F

	chip.Execute(0x8563) // V5 = V5 ^ V6

	if chip.V[5] != 0xA5 {
		t.Errorf("Expected V5 = 0xA5, got 0x%02X", chip.V[5])
	}
}
