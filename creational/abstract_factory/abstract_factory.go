package abstract_factory

import "fmt"

type Shoe interface {
	Wear() string
}

type Shirt interface {
	PutOn() string
}

type BrandFactory interface {
	CreateShoe() Shoe
	CreateShirt() Shirt
}

func GetBrandFactory(brand string) (BrandFactory, error) {
	switch brand {
	case "nike":
		return new(NikeFactory), nil

	case "adidas":
		return new(AdidasFactory), nil

	default:
		return nil, fmt.Errorf("Unknown brand: %s", brand)
	}
}
