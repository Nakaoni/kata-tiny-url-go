package katatinyurl

import (
	"regexp"
	"testing"
)

func TestCodeGenerator(t *testing.T) {
	t.Run("it returns a 6 char-length code", func(t *testing.T) {
		code := GenerateCode()

		wantLength := 6
		gotLength := len(code)

		assertEquals(t, gotLength, wantLength)
	})
	t.Run("it returns a different code each time", func(t *testing.T) {
		code1 := GenerateCode()
		code2 := GenerateCode()

		assertNotEquals(t, code1, code2)
	})
	t.Run("it contains only alphanumeric characters", func(t *testing.T) {
		pattern := regexp.MustCompile(`^[a-zA-Z0-9]{CODE_LENGTH}$`)

		for i := range 100 {
			code := GenerateCode()
			if !pattern.MatchString(code) {
				t.Errorf("Iteration %d: Code %q is invalid", i, code)
			}
		}
	})
}

func BenchmarkGenerateCode(b *testing.B) {
	for b.Loop() {
		GenerateCode()
	}
}

func assertEquals(t *testing.T, got any, want any) {
	t.Helper()
	if got != want {
		t.Errorf("Want %v does not equal to Got %v.", want, got)
	}
}

func assertNotEquals(t *testing.T, got any, want any) {
	t.Helper()
	if got == want {
		t.Errorf("Want %v equals to Got %v.", want, got)
	}
}
