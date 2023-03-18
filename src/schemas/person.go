package schemas

type Person struct {
	Name string
	Age  int
	Address
}

func (p Person) Greeting() string {
	return "Hello" + " " + p.Name
}
