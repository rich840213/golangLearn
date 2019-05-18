package main

/*儲存變數類型*/

import "fmt"

type Human struct {
	name string
	age  int
}

func main() {
	temp := []interface{}{}

	a := 1
	b := "foo"
	c := Human{
		name: "bar",
		age:  25,
	}
	d := 100.10

	temp = append(temp, a, b, c, d)

	//if方式
	for i, value := range temp {
		//判斷型態是否可以轉
		if v, ok := value.(int); ok {
			fmt.Printf("%d, value: %v\n", i, v)
		} else if v, ok := value.(string); ok {
			fmt.Printf("%d, string value: %v\n", i, v)
		} else if v, ok := value.(Human); ok {
			fmt.Printf("%d, Human value: %v\n", i, v.name)
		} else {
			fmt.Printf("%d, Unknownn\n", i)
		}
	}

	//switch方式
	for i, value := range temp {
		switch v := value.(type) {
		case int:
			fmt.Printf("%d, value: %v\n", i, v)
		case string:
			fmt.Printf("%d, string value: %v\n", i, v)
		case Human:
			fmt.Printf("%d, Human value: %v\n", i, v.name)
		default:
			fmt.Printf("%d, Unknownn\n", i)
		}
	}
}
