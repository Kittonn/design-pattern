package abstract_factory

type AdidasShoe struct{}

func (s *AdidasShoe) Wear() string {
	return "Wearing Adidas shoes"
}

type AdidasShirt struct{}

func (s *AdidasShirt) PutOn() string {
	return "Putting on an Adidas shirt"
}

type AdidasFactory struct{}

func (f *AdidasFactory) CreateShoe() Shoe {
	return new(AdidasShoe)
}

func (f *AdidasFactory) CreateShirt() Shirt {
	return new(AdidasShirt)
}
