// go test ./examples -run 'TestMakeMultiplier|TestMakeAdder|TestApplyTwice' -v

package examples

import "testing"

func TestMakeMultiplier(t *testing.T) {
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	if double(5) != 10 {
		t.Errorf("expected 10, got %d", double(5))
	}

	if triple(4) != 12 {
		t.Errorf("expected 12, got %d", triple(4))
	}
}

func TestMakeAdder(t *testing.T) {
	add10 := makeAdder(10)
	add5 := makeAdder(5)

	if add10(5) != 15 {
		t.Errorf("expected 15, got %d", add10(5))
	}

	if add5(3) != 8 {
		t.Errorf("expected 8, got %d", add5(3))
	}
}

func TestApplyTwice(t *testing.T) {
	double := func(x int) int {
		return x * 2
	}

	f := applyTwice(double)

	if f(3) != 12 {
		t.Errorf("expected 12, got %d", f(3))
	}

	increment := func(x int) int {
		return x + 1
	}

	g := applyTwice(increment)

	if g(5) != 7 {
		t.Errorf("expected 7, got %d", g(5))
	}
}