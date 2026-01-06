package functions

func PointerTo[T any](v T) *T {
	return &v
}
