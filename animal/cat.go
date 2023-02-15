package animal

type Cat interface {
	Name() string
	Eat(food string)
	Climb(tree string)
}

type cat struct {
	mammal
}

func MakeCat(name string) Cat {
	return &cat{
		mammal: mammal{animal{name}},
	}
}

func (c *cat) Climb(tree string) {
	println("Cat", c.Name(), "climbs", tree)
}
