package init

import "fmt"

//main執行前才會執行init，不管順序
func init() {
	fmt.Println("init run init()!")
}
