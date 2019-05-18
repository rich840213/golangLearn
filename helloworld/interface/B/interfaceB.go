package main

/*使用Interface為函數參數(fmt的Stringer Interface)*/

import "fmt"

type IPv4 []byte

func (a IPv4) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}

func main() {
	ipv4 := map[string]IPv4{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}

	for i, v := range ipv4 {
		fmt.Printf("%v : %v\n", i, v)
	}
}
