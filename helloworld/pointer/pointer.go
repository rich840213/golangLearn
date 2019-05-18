package main

import "fmt"

func foo_value(x int) {
	fmt.Println(&x) // function內x的記憶體位置
}

func foo_point(x *int) {
	fmt.Println(x) // function內x的記憶體位置
}

func main() {
	p := 1

	//取"位址"需用"&"符號。 "*"號是取該"位址"的值
	fmt.Println(&p)
	fmt.Println(p)

	//p1指向p的位址
	p1 := &p
	//透過p1取p的值
	fmt.Println(*p1)
	//透過p1更改p的值
	*p1 = 2
	fmt.Println(p)
	fmt.Println(*p1)
	//位址的確相同
	fmt.Println(p1)

	//宣告一個空的int
	n := new(int)
	//直接把值指向記憶體位置
	*n = 2
	fmt.Println(n)
	fmt.Println(*n)

	b := new(int)
	*b = 2
	fmt.Println(b)
	foo_value(*b)
	foo_point(b)
}
