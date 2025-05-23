package tuple

type Tuple[A any, B any] struct {
	First  A
	Second B
}

func T[A any, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{
		First:  a,
		Second: b,
	}
}
