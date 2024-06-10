package beaut

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

func (s Set[T]) Remove(t T) {
	delete(s, t)
}

type runeWithLength struct {
	char      rune
	byteCount int32
}

func TestFindRuneAtIndex(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.String().Draw(t, "s")
		startIndexes := map[int]runeWithLength{}
		bytesSoFar := 0
		for _, c := range s {
			_, n := explodeRune(c)
			startIndexes[bytesSoFar] = runeWithLength{char: c, byteCount: n}
			bytesSoFar += int(n)
		}
		bytes := []byte(s)
		for i := 0; i < len(s); i++ {
			savedRune, ok := startIndexes[i]
			r, runeLength := findRuneAtIndex(bytes, ByteIndex(i))
			if runeLength != -1 {
				require.Equalf(t, savedRune.char, r, "index %d, map: %v", i, startIndexes)
				require.Equal(t, savedRune.byteCount, runeLength)
			} else {
				require.Falsef(t, ok, "map says rune '%c' starts at byte %d, but has invalid start byte '%x'",
					savedRune.char, i, bytes[i])
			}
		}
	})
}

func TestSafelySliceable(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		bytes := []byte(rapid.String().Draw(t, "s"))
		start := rapid.IntRange(0, max(0, len(bytes)-1)).Draw(t, "start")
		end := rapid.IntRange(start, len(bytes)).Draw(t, "end")
		err := checkSafelySliceableUTF8(bytes, ByteIndex(start), ByteIndex(end))
		isValidUTF8 := utf8.Valid(bytes[start:end])
		if err == nil {
			require.True(t, isValidUTF8, "checkSafelySliceable was OK but is invalid UTF-8")
		} else {
			require.False(t, isValidUTF8, "checkSafelySliceable returned error but is valid UTF-8")
		}
	})
}
