package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Mary")
	if s != "Welcome my dear Mary" {
		t.Error("got", s, "expected", "Welcome my dear Mary")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Mary"))
	// Output:
	// Welcome my dear Mary
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Mary")
	}
}
