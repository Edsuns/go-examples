package animal

import "testing"

func TestOOP(t *testing.T) {
	var (
		cat = MakeCat("a")
		dog = MakeDog("b")
	)

	PrintName(cat)
	PrintName(dog)

	cat.Eat("x")
	dog.Eat("y")

	cat.Climb("1")
	dog.Swim(2)
}
