package utils

type Action[T any] struct {
	Fn   T
	Name string
}

func (action *Action[T]) GetFn() T {
	return action.Fn
}

func (action *Action[T]) GetName() string {
	return action.Name
}

func (action *Action[T]) String() string {
	return action.Name
}
