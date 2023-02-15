package animal

type Animal interface {
	Name() string
}

type animal struct {
	name string
}

func (a *animal) Name() string {
	return a.name
}

func PrintName(a Animal) {
	println(a.Name())
}
