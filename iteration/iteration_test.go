package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"
	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 5)
	fmt.Println(result)
	// Output: bbbbb
}
