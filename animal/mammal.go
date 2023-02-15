package animal

type Mammal interface {
	Name() string
	Eat(food string)
}

type mammal struct {
	animal
}

func (m *mammal) Eat(food string) {
	println(m.Name(), "eats food", food)
}
