package common

func PointerOf[T any](v T) *T {
	return &v
}
