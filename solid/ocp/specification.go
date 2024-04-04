package solid

// Here each and every time we found new specification then we will create
// new implementation of Specification interface instead of modifying existing object
// we will never modify the BetterFilter
// Here this design blocking modification but allowing extendability, because developer has to provide
// new implementation of a specification

type Specification interface {
	IsSatisfied(p *Shirt) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Shirt) bool {
	return p.Color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Shirt) bool {
	return p.Size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Shirt) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Shirt, spec Specification) []*Shirt {
	result := make([]*Shirt, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
