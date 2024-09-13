package core

type BaseService[T int | float32 | uint] struct {
}

func (service *BaseService[T]) Sum(a T, b T) T {
	return a + b
}
