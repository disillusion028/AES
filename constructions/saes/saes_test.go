package saes

import (
	"testing"
)

var key = [16]byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33, 33, 33, 33, 33}

func TestSubByte(t *testing.T) {
	constr := Construction{key}

	if constr.subByte(0x00) != 0x63 {
	  t.Fatalf("Affine component of SubByte is wrong!")
	}

	if constr.subByte(0x53) != 0xED {
	  t.Fatalf("Linear component of SubByte is wrong! #1")
	}

  if constr.subByte(0x02) != 0x77 {
	  t.Fatalf("Linear component of SubByte is wrong! #2")
	}

	// Test subWord
	wordA := uint32((0x10 << 24) | (0x53 << 16) | (0x86 << 8) | 0xED)
	wordB := uint32((0xCA << 24) | (0xED << 16) | (0x44 << 8) | 0x55)

	if constr.subWord(wordA) != wordB {
	  t.Fatalf("constr.subWord gave incorrect output!")
	}
}

// func TestKeyStretching(t *testing.T) {
//   real := [11][16]byte{
//     [16]byte{ 72, 101, 108, 108, 111,  32,  87, 111, 114, 108, 100,  33,  33,  33,  33,  33},
//     [16]byte{180, 152, 145, 145, 219, 184, 198, 254, 169, 212, 162, 223, 136, 245, 131, 254},
//     [16]byte{ 80, 116,  42,  85, 139, 204, 236, 171,  34,  24,  78, 116, 170, 237, 205, 138},
//     [16]byte{  1, 201,  84, 249, 138,   5, 184,  82, 168,  29, 246,  38,   2, 240,  59, 172},
//     [16]byte{133,  43, 197, 142,  15,  46, 125, 220, 167,  51, 139, 250, 165, 195, 176,  86},
//     [16]byte{187, 204, 116, 136, 180, 226,   9,  84,  19, 209, 130, 174, 182,  18,  50, 248},
//     [16]byte{ 82, 239,  53, 198, 230,  13,  60, 146, 245, 220, 190,  60,  67, 206, 140, 196},
//     [16]byte{153, 139,  41, 220, 127, 134,  21,  78, 138,  90, 171, 114, 201, 148,  39, 182},
//     [16]byte{ 59,  71, 103,   1,  68, 193, 114,  79, 206, 155, 217,  61,   7,  15, 254, 139},
//     [16]byte{ 86, 252,  90, 196,  18,  61,  40, 139, 220, 166, 241, 182, 219, 169,  15,  61},
//     [16]byte{179, 138, 125, 125, 161, 183,  85, 246, 125,  17, 164,  64, 166, 184, 171, 125},
//   }
//
//   constr := Construction{key}
//
//   cand := constr.stretchedKey()
//
//   fmt.Println(cand)
//   for i := 0; i < 11; i++ {
//     for j := 0; j < 16; j++ {
//       if real[i][j] != cand[i][j] {
//         t.Fatalf("Byte (%v, %v) is wrong! %v != %v", i, j, real[i][j], cand[i][j])
//       }
//     }
//   }
// }