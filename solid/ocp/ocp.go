package solid

import "fmt"

// creating a Predicate type function, this can be any named function in future that should return bool
type Predicate[T any] func(T) bool

// This filter can be added to any objects
// We will not modify the Filter for new filter requirements
// This will accept a Predicate to test, So for each requirement developer as to provide new Predicate
// Here this design blocking modification but allowing extendability
func Filter[T any](values []T, predicate Predicate[T]) []T {
	var filtered []T
	for _, value := range values {
		if predicate(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func PrintShirtsOCP(shirts []Shirt) {
	for _, shirt := range shirts {
		fmt.Println(shirt.String())
	}
}
