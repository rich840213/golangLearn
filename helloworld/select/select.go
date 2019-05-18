package main

import "fmt"

/*select只可用在channel，如果有多條channel要在一個執行緒去管理輸出*/
/*要注意select到底會進入哪一條case，外在是無法控制的，這是由golang內部演算法實作，所以如果case的情況多，
有機率不會進到某case裡面。所以要斟酌使用。*/

var channel01 chan string
var channel02 chan string

func callCh01() {
	channel01 <- "this is channel01"
}

func callCh02() {
	channel02 <- "this is channel02"
}

func main() {
	channel01 = make(chan string)
	channel02 = make(chan string)

	go callCh01()
	go callCh02()

	i := 0
	for {
		select {
		case msg := <-channel01:
			fmt.Println(msg)
		case msg := <-channel02:
			fmt.Println(msg)
		}

		//這邊因為只有兩個 sender 各送一條訊息，所以這個loop 只要執行兩次就可以 break，因為不會再有其他訊息進來
		i++
		if i >= 2 {
			break
		}

	}
	fmt.Println("end")
}
