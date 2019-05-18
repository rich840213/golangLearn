package main

import "fmt"

/*透過interface實現多型*/

type Animal interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Speak() string {
	return "汪汪"
}

func (c Cat) Speak() string {
	return "喵喵"
}

func main() {
	animal := []Animal{Dog{}, Cat{}}
	for _, a := range animal {
		fmt.Println(a.Speak())
	}
}
