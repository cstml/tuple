package tuple

import (
	"fmt"
	"testing"
)

func TestEquality(m *testing.T) {
	t1 := T(1, 2)
	t2 := T(1, 2)
	if t1 != t2 {
		m.Error("should be equal")
	}

	t3 := T("abc", 2)
	if t3 != t3 {
		m.Error("should be equal")
	}
}

func TestSendChannel(t *testing.T) {
	c := make(chan (Tuple[error, int]))

	go func() {
		r := <-c
		if r.First.Error() != "random" {
			t.Error("unexpected error")
		}
		if r.Second != 0 {
			t.Error("unexpected number")
		}
	}()

	c <- T(fmt.Errorf("random"), 0)
}

func TestNestedTuple(t *testing.T) {
	z := T(1, T(2, T(3, "4")))
	q := z
	q.First = 2
	if q == z {
		t.Error("Referenced not copied")
	}
}
