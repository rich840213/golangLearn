package main

import "fmt"

type Info struct {
	id   int
	name string
}

func main() {
	s := &Info{
		id:   1,
		name: "123",
	}

	//%v：單純把內容都印出來
	fmt.Printf("Info is %v\n", s)
	//%+v：除了值以外，還有欄位名稱
	fmt.Printf("Info is %+v\n", s)
	//%#v：可以得到最多資訊，連結構體的名稱、所在的函數都會一起印出
	fmt.Printf("Info is %#v\n", s)

	s.id = 2
	s.name = "456"
	//%v：單純把內容都印出來
	fmt.Printf("Info is %v\n", s)
	//%+v：除了值以外，還有欄位名稱
	fmt.Printf("Info is %+v\n", s)
	//%#v：可以得到最多資訊，連結構體的名稱、所在的函數都會一起印出
	fmt.Printf("Info is %#v\n", s)

	fmt.Printf("Info all value is %d, %s\n", s.id, s.name)
}
