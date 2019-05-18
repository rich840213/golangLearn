package main

import "fmt"

/*透過struct & func實現OOP的class & function*/

type User struct {
	id   int
	name string
}

func (u *User) SetId(i int) {
	u.id = i
}

func (u *User) SetName(n string) {
	u.name = n
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func main() {
	user := &User{
		id:   123,
		name: "abc",
	}

	fmt.Printf("User all value is %d, %s\n", user.GetId(), user.GetName())

	user.SetId(456)
	user.SetName("xyz")
	fmt.Printf("User all value is %d, %s", user.GetId(), user.GetName())
}
