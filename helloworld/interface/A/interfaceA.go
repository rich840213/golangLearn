package main

/*PHP Interface Class轉Golang Struct*/

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (h Human) Run() {
	fmt.Printf("%s, I am running\n", h.name)
}

func (h Human) Say() {
	fmt.Printf("%s, I am %d years\n", h.name, h.age)
}

type Singer struct {
	Human
	collection string
}

func (s Singer) Sing() {
	fmt.Println("I am sing")
	fmt.Printf("%s\n", s.collection)
}

type Man interface {
	Run()
	Say()
}

func main() {
	foo := Human{
		name: "foo",
		age:  20,
	}

	bar := Singer{
		Human: Human{
			name: "bar",
			age:  25,
		},
		collection: "test",
	}

	/*foo.Run()
	foo.Say()

	bar.Run()
	bar.Say()
	bar.Sing()*/

	test := []Man{}
	test = append(test, foo, bar)
	for _, v := range test {
		v.Run()
		v.Say()
	}
}
