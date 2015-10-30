package chow

import (
	"github.com/OpenWhiteBox/AES/primitives/encoding"
	"github.com/OpenWhiteBox/AES/primitives/matrix"
	"github.com/OpenWhiteBox/AES/primitives/table"

	"github.com/OpenWhiteBox/AES/constructions/common"
	"github.com/OpenWhiteBox/AES/constructions/saes"
)

func generateKeys(rs *common.RandomSource, opts common.KeyGenerationOpts, out *Construction, inputMask, outputMask *matrix.Matrix, shift func(int) int, skinny func(int) table.Byte, wide func(int, int) table.Word) {
	// Generate input and output encodings.
	common.GenerateMasks(rs, opts, inputMask, outputMask)

	// Generate the Input Mask table and the 10th T-Box/Output Mask table
	for pos := 0; pos < 16; pos++ {
		out.InputMask[pos] = encoding.BlockTable{
			encoding.IdentityByte{},
			BlockMaskEncoding(rs, pos, common.Inside, shift),
			common.BlockMatrix{*inputMask, [16]byte{}, pos},
		}

		inEnc := common.MixingBijection(rs, 8, 8, pos)
		inInv, _ := inEnc.Invert()

		out.TBoxOutputMask[pos] = encoding.BlockTable{
			encoding.ComposedBytes{
				encoding.ByteLinear{inEnc, inInv},
				ByteRoundEncoding(rs, 8, pos, common.Outside),
			},
			BlockMaskEncoding(rs, pos, common.Outside, shift),
			table.ComposedToBlock{
				skinny(pos),
				common.BlockMatrix{*outputMask, [16]byte{}, pos},
			},
		}
	}

	// Generate the XOR Tables for the Input and Output Masks.
	out.InputXORTable = blockXORTables(rs, common.Inside, shift)
	out.OutputXORTable = blockXORTables(rs, common.Outside, shift)

	// Generate round material.
	for round := 0; round < 9; round++ {
		for pos := 0; pos < 16; pos++ {
			// Generate a word-sized mixing bijection and stick it on the end of the T-Box/Tyi Table.
			mb := common.MixingBijection(rs, 32, round, pos/4)

			inEnc := common.MixingBijection(rs, 8, round-1, pos)
			inInv, _ := inEnc.Invert()

			// Build the T-Box and Tyi Table for this round and position in the state matrix.
			out.TBoxTyiTable[round][pos] = encoding.WordTable{
				encoding.ComposedBytes{
					encoding.ByteLinear{inEnc, inInv},
					encoding.ConcatenatedByte{
						RoundEncoding(rs, round-1, 2*pos+0, common.Outside),
						RoundEncoding(rs, round-1, 2*pos+1, common.Outside),
					},
				},
				encoding.ComposedWords{
					encoding.ConcatenatedWord{
						encoding.ByteLinear{common.MixingBijection(rs, 8, round, shift(pos/4*4+0)), nil},
						encoding.ByteLinear{common.MixingBijection(rs, 8, round, shift(pos/4*4+1)), nil},
						encoding.ByteLinear{common.MixingBijection(rs, 8, round, shift(pos/4*4+2)), nil},
						encoding.ByteLinear{common.MixingBijection(rs, 8, round, shift(pos/4*4+3)), nil},
					},
					encoding.WordLinear{mb, nil},
					WordStepEncoding(rs, round, pos, common.Inside),
				},
				wide(round, pos),
			}

			// Encode the inverse of the mixing bijection from above in the MB^(-1) table for this round and position.
			mbInv, _ := mb.Invert()

			out.MBInverseTable[round][pos] = encoding.WordTable{
				encoding.ConcatenatedByte{
					RoundEncoding(rs, round, 2*pos+0, common.Inside),
					RoundEncoding(rs, round, 2*pos+1, common.Inside),
				},
				WordStepEncoding(rs, round, pos, common.Outside),
				MBInverseTable{mbInv, uint(pos) % 4},
			}
		}
	}

	// Generate the High and Low XOR Tables for reach round.
	out.HighXORTable = xorTables(rs, common.Inside, shift)
	out.LowXORTable = xorTables(rs, common.Outside, shift)
}

// GenerateEncryptionKeys creates a white-boxed version of the AES key `key` for encryption, with any non-determinism
// generated by `seed`.  The `opts` specifies what type of input and output masks we put on the construction and should
// be either IndependentMasks, SameMasks, or MatchingMasks.
func GenerateEncryptionKeys(key, seed []byte, opts common.KeyGenerationOpts) (out Construction, inputMask, outputMask matrix.Matrix) {
	rs := common.NewRandomSource("Chow Encryption", seed)

	constr := saes.Construction{key}
	roundKeys := constr.StretchedKey()

	// Apply ShiftRows to round keys 0 to 9.
	for k := 0; k < 10; k++ {
		constr.ShiftRows(roundKeys[k])
	}

	skinny := func(pos int) table.Byte {
		return common.TBox{constr, roundKeys[9][pos], roundKeys[10][pos]}
	}

	wide := func(round, pos int) table.Word {
		return table.ComposedToWord{
			common.TBox{constr, roundKeys[round][pos], 0x00},
			common.TyiTable(pos % 4),
		}
	}

	generateKeys(&rs, opts, &out, &inputMask, &outputMask, common.ShiftRows, skinny, wide)

	return
}

// GenerateDecryptionKeys creates a white-boxed version of the AES key `key` for decryption, with any non-determinism
// generated by `seed`.  The `opts` argument works the same as above.
func GenerateDecryptionKeys(key, seed []byte, opts common.KeyGenerationOpts) (out Construction, inputMask, outputMask matrix.Matrix) {
	rs := common.NewRandomSource("Chow Decryption", seed)

	constr := saes.Construction{key}
	roundKeys := constr.StretchedKey()

	// Last key needs to be unshifted for decryption to work right.
	constr.UnShiftRows(roundKeys[10])

	skinny := func(pos int) table.Byte {
		return common.InvTBox{constr, 0x00, roundKeys[0][pos]}
	}

	wide := func(round, pos int) table.Word {
		if round == 0 {
			return table.ComposedToWord{
				common.InvTBox{constr, roundKeys[10][pos], roundKeys[9][pos]},
				common.InvTyiTable(pos % 4),
			}
		} else {
			return table.ComposedToWord{
				common.InvTBox{constr, 0x00, roundKeys[9-round][pos]},
				common.InvTyiTable(pos % 4),
			}
		}
	}

	generateKeys(&rs, opts, &out, &inputMask, &outputMask, common.UnShiftRows, skinny, wide)

	return
}
