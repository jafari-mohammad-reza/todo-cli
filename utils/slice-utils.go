package utils

func Remove[T any](slice []T, index int) {
	slice = append(slice[:index], slice[index+1:]...)
}
