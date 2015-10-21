package chow

import (
	"fmt"

	"github.com/OpenWhiteBox/AES/primitives/encoding"
	"github.com/OpenWhiteBox/AES/primitives/matrix"
	"github.com/OpenWhiteBox/AES/primitives/table"

	"github.com/OpenWhiteBox/AES/constructions/saes"
)

type TableAsEncoding struct {
	Forwards, Backwards table.InvertibleTable
}

func (tae TableAsEncoding) Encode(i byte) byte { return tae.Forwards.Get(i) }
func (tae TableAsEncoding) Decode(i byte) byte { return tae.Backwards.Get(i) }

var powx = [16]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1b, 0x36, 0x6c, 0xd8, 0xab, 0x4d, 0x9a, 0x2f}

// BackOneRound takes round key i and returns round key i-1.
func BackOneRound(roundKey []byte, round int) (out []byte) {
	out = make([]byte, 16)
	constr := saes.Construction{}

	// Recover everything except the first word by XORing consecutive blocks.
	for pos := 4; pos < 16; pos++ {
		out[pos] = roundKey[pos] ^ roundKey[pos-4]
	}

	// Recover the first word by XORing the first block of the roundKey with f(last block of roundKey), where f is a
	// subroutine of AES' key scheduling algorithm.
	for pos := 0; pos < 4; pos++ {
		out[pos] = roundKey[pos] ^ constr.SubByte(out[12+(pos+1)%4])
	}
	out[0] ^= powx[round-1]

	return
}

// FunctionFromBasis produces the element of S according a basis and a specified combination of its elements.  It is an
// isomorphism from (0..2^8-1, xor) -> (S, compose).
func FunctionFromBasis(i int, basis []table.Byte) table.Byte {
	// Generate the function specified by i.
	vect := table.ComposedBytes{
		table.IdentityByte{},
	}

	for j := uint(0); j < uint(len(basis)); j++ {
		if (i>>j)&1 == 1 {
			vect = append(vect, basis[j])
		}
	}

	return vect
}

// DecomposeAffineEncoding is an efficient way to factor an unknown affine encoding into its component linear and
// affine parts.
func DecomposeAffineEncoding(e encoding.Byte) (matrix.Matrix, byte) {
	m := matrix.Matrix{
		matrix.Row{0}, matrix.Row{0}, matrix.Row{0}, matrix.Row{0},
		matrix.Row{0}, matrix.Row{0}, matrix.Row{0}, matrix.Row{0},
	}
	c := e.Encode(0)

	for i := uint(0); i < 8; i++ {
		x := e.Encode(1<<i) ^ c

		for j := uint(0); j < 8; j++ {
			if (x>>j)&1 == 1 {
				m[j][0] += 1 << i
			}
		}
	}

	return m, c
}

// isAffine returns true if the given encoding is affine and false if not.
func isAffine(aff encoding.Byte) bool {
	m, c := DecomposeAffineEncoding(aff)
	test := encoding.ByteAffine{encoding.ByteLinear(m), c}

	for i := 0; i < 256; i++ {
		a, b := aff.Encode(byte(i)), test.Encode(byte(i))
		if a != b {
			return false
		}
	}

	return true
}

// FindCharacteristic finds the characteristic number of a matrix.  This number is invariant to matrix similarity.
func FindCharacteristic(A matrix.Matrix) (a byte) {
	AkEnc := encoding.ComposedBytes{}

	for k := uint(0); k < 8; k++ {
		AkEnc = append(AkEnc, encoding.ByteLinear(A))
		Ak, _ := DecomposeAffineEncoding(AkEnc)
		a ^= Ak.Trace() << k
	}

	return
}

// FindDuplicate returns the first duplicate matrix it finds in a set of matrices.
func FindDuplicate(ms []matrix.Matrix) matrix.Matrix {
	// Forall pairs of matrices without repetition
	for i, m := range ms {
		for _, n := range ms[i+1:] {
			if fmt.Sprintf("%x", m) == fmt.Sprintf("%x", n) {
				return m
			}
		}
	}

	return nil
}

// Index in, index out.  Example: shiftRows(5) = 1 because ShiftRows(block) returns [16]byte{block[0], block[5], ...
func shiftRows(i int) int {
	return []int{0, 13, 10, 7, 4, 1, 14, 11, 8, 5, 2, 15, 12, 9, 6, 3}[i]
}

func unshiftRows(i int) int {
	return []int{0, 5, 10, 15, 4, 9, 14, 3, 8, 13, 2, 7, 12, 1, 6, 11}[i]
}