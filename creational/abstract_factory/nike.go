package abstract_factory

type NikeShoe struct{}

func (s *NikeShoe) Wear() string {
	return "Wearing Nike shoes"
}

type NikeShirt struct{}

func (s *NikeShirt) PutOn() string {
	return "Putting on a Nike shirt"
}

type NikeFactory struct{}

func (f *NikeFactory) CreateShoe() Shoe {
	return new(NikeShoe)
}

func (f *NikeFactory) CreateShirt() Shirt {
	return new(NikeShirt)
}
