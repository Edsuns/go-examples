package animal

type Dog interface {
	Name() string
	Eat(food string)
	Swim(distance int)
}

type dog struct {
	mammal
}

func MakeDog(name string) Dog {
	return &dog{
		mammal: mammal{animal{name}},
	}
}

func (d *dog) Swim(distance int) {
	println("Dog", d.Name(), "swims", distance)
}
