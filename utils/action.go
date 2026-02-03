package utils

type Action[T any] struct {
	Fn   T
	Name string
}

func (action *Action[T]) String() string {
	return action.Name
}
