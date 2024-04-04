package solid

import "fmt"

// In this code you can see each time developer wanted to visit Filter object and need to add new function for new filter requirement.
// here we modified the object each time, we need to do extend instead of modification
// If we did not provide extendability then this filter will end up in 100s of filter implementations

type ShirtFilter struct {
}

func (f *ShirtFilter) FilterByColor(shirts []Shirt, color Color) []*Shirt {
	filtered := make([]*Shirt, 0)
	for i, shirt := range shirts {
		if shirt.Color == color {
			filtered = append(filtered, &shirts[i])
		}
	}
	return filtered
}

func (f *ShirtFilter) FilterBySize(shirts []Shirt, size Size) []*Shirt {
	filtered := make([]*Shirt, 0)
	for i, shirt := range shirts {
		if shirt.Size == size {
			filtered = append(filtered, &shirts[i])
		}
	}
	return filtered
}

func (f *ShirtFilter) FilterByColorAndSize(shirts []Shirt, color Color, size Size) []*Shirt {
	filtered := make([]*Shirt, 0)
	for i, shirt := range shirts {
		if shirt.Color == color && shirt.Size == size {
			filtered = append(filtered, &shirts[i])
		}
	}
	return filtered
}

func PrintShirts(shirts []*Shirt) {
	for _, shirt := range shirts {
		fmt.Println(shirt.String())
	}
}
