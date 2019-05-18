package main

/*接收輸入參數*/
/*go run osargs.go user 123456*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//變數宣告4種方式
	/*s := "" (常用)
	var s string (常用)
	var s = ""
	var s string = ""*/

	var s1, sep1 string
	s2, sep2 := "", ""
	//s3, sep3 := "", ""

	//for共有二種
	//第一種for
	for i := 1; i < len(os.Args); i++ {
		s1 += sep1 + os.Args[i]
		sep1 = " "
	}
	fmt.Println(s1)

	//第二種for
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)

	//strings.join效率比較好
	fmt.Println(strings.Join(os.Args[1:], " "))
}
