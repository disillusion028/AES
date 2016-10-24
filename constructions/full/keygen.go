package full

import (
	"io"

	"github.com/OpenWhiteBox/primitives/encoding"
	"github.com/OpenWhiteBox/primitives/matrix"
	"github.com/OpenWhiteBox/primitives/random"

	"github.com/OpenWhiteBox/AES/constructions/common"
	"github.com/OpenWhiteBox/AES/constructions/saes"
)

// generateAffineMasks creates the random external masks for the construction.
func generateAffineMasks(rs *random.Source) (inputMask, outputMask *blockAffine) {
	var inputLinear, outputLinear matrix.Matrix
	common.GenerateMasks(rs, common.IndependentMasks{common.RandomMask, common.RandomMask}, &inputLinear, &outputLinear)

	reader := rs.Stream(make([]byte, 16))

	inputConstant, outputConstant := matrix.NewRow(128), matrix.NewRow(128)
	reader.Read(inputConstant[:])
	reader.Read(outputConstant[:])

	inputMask = &blockAffine{linear: inputLinear, constant: inputConstant}
	outputMask = &blockAffine{linear: outputLinear, constant: outputConstant}

	return
}

// generateSelfEquivalence returns a random self-equivalence of the S-box layer, defined by stateSize and compressSize.
func generateSelfEquivalence(r io.Reader, stateSize, compressSize int) (a, bInv *blockAffine) {
	inSize, outSize := 8*(stateSize+compressSize), 8*stateSize
	in := &blockAffine{
		linear:   matrix.GenerateEmpty(inSize, inSize),
		constant: matrix.NewRow(inSize),
	}
	out := &blockAffine{
		linear:   matrix.GenerateEmpty(outSize, outSize),
		constant: matrix.NewRow(outSize),
	}

	// The S-box portion of the self-equivalence. Set as the identity for now.
	// TODO(brendan): Implement me.
	for i := 0; i < 8*compressSize; i++ {
		in.linear[2*i+0].SetBit(2*i+0, true)
		in.linear[2*i+1].SetBit(2*i+1, true)
		out.linear[i].SetBit(i, true)
	}

	// The open portion of the self-equivalence. Fill it with a random, invertible matrix.
	ignoreAll := func(_ int) bool { return true }
	dense, denseInv := matrix.GenerateRandomPartial(r, 8*(stateSize-compressSize), matrix.IgnoreNoBytes, ignoreAll)
	for i := 0; i < 8*(stateSize-compressSize); i++ {
		copy(in.linear[8*2*compressSize+i][2*compressSize:], dense[i])
		copy(out.linear[8*compressSize+i][compressSize:], denseInv[i])
	}

	return in, out
}

// GenerateKeys creates a white-boxed version of the AES key `key`, with any non-determinism generated by `seed`.
func GenerateKeys(key, seed []byte) (out Construction, inputMask, outputMask encoding.BlockAffine) {
	rs := random.NewSource("Ful Construction", seed)

	// Generate two completely random affine transformations, to be put on input and output of SPN.
	input, output := generateAffineMasks(&rs)

	// Steal key schedule logic from the standard AES construction.
	contr := saes.Construction{key}
	roundKeys := contr.StretchedKey()

	// Generate an SPN which has the input and output masks, but is otherwise un-obfuscated.
	out[0] = decomposition[0].compose(&blockAffine{
		linear:   matrix.GenerateIdentity(128),
		constant: matrix.Row(roundKeys[0]),
	}).compose(input)
	copy(out[1:5], decomposition[1:5])

	for i := 1; i < 10; i++ {
		out[4*i+0] = decomposition[0].compose(&blockAffine{
			linear:   round,
			constant: matrix.Row(roundKeys[i]).Add(subBytesConst),
		}).compose(out[4*i+0])
		copy(out[4*i+1:4*i+5], decomposition[1:5])
	}

	out[40] = output.compose(&blockAffine{
		linear:   lastRound,
		constant: matrix.Row(roundKeys[10]).Add(subBytesConst),
	}).compose(out[40])

	// Sample a self-equivalences of the S-box layer and mix them into adjacent affine layers.
	label := make([]byte, 16)
	copy(label, []byte("Self-Eq"))
	r := rs.Stream(label)

	for i := 0; i < 40; i++ {
		a, bInv := generateSelfEquivalence(r, stateSize[i%4], compressSize[i%4])
		out[i] = a.compose(out[i])
		out[i+1] = out[i+1].compose(bInv)
	}

	return out, input.BlockAffine(), output.BlockAffine()
}
