package table

import (
	"testing"
)

type AddTable byte

func (at AddTable) Get(i byte) byte { return i + byte(at) }

type TimesTable byte

func (tt TimesTable) Get(i byte) byte { return i * byte(tt) }

type ShiftTable uint

func (st ShiftTable) Get(i byte) uint32 { return uint32(i) << uint(st) }

func TestCompose(t *testing.T) {
	x := ComposedBytes{TimesTable(5), AddTable(3)}
	y := ComposedBytes{AddTable(3), TimesTable(5)}

	a := ComposedToWord{x, ShiftTable(24)}

	if x.Get(7) != ((7 * 5) + 3) {
		t.Fatalf("X's Composition is wrong.")
	}

	if y.Get(7) != ((7 + 3) * 5) {
		t.Fatalf("Y's Composition is wrong.")
	}

	if a.Get(7) != ((7*5)+3)<<24 {
		t.Fatalf("A's Composition is wrong.")
	}

	// Incorrect table compositions should refuse to compile.
}

func TestPersist(t *testing.T) {
	x := AddTable(3)
	serializedX := SerializeByte(x)
	parsedX := ParsedByte(serializedX)

	y := ShiftTable(13)
	serializedY := SerializeWord(y)
	parsedY := ParsedWord(serializedY)

	for i := 0; i < 256; i++ {
		if x.Get(byte(i)) != parsedX.Get(byte(i)) {
			t.Fatalf("X and ParsedX disagree at point %v: %v != %v", i, x.Get(byte(i)), parsedX.Get(byte(i)))
		}

		if y.Get(byte(i)) != parsedY.Get(byte(i)) {
			t.Fatalf("Y and ParsedY disagree at point %v: %v != %v", i, y.Get(byte(i)), parsedY.Get(byte(i)))
		}
	}
}