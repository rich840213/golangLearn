package main

import (
	"fmt"
	"strconv"
	"time"
)

/*channel，中文翻譯為"通道"。實作上通常搭配Goroutine，作為彼此之間的通訊機制，可以接收(receiver)與發送(sender)資料。*/

/*channel如果沒有宣告buffer，sender端送進去的訊息，receiver端沒收之前，sender端送不進去第二則資料，
此時，sender端會block，需注意控制這個問題，以避免造成deadlock。

解決deadlock的方法有兩種:
1. 透過for loop不斷接收資料
2. 透過channel buffer
*/

//發送
//ch <- 5

//接收，將箭頭指向變數
//rec := <-ch

//通道被關閉後，資料就只能讀取不能寫入，read ok，not write
//close(ch)

//宣告全域channel
var message chan string

func Bot() {
	msg := <-message
	fmt.Println("msg is ", msg)
}

func main() {
	//宣告區域channel，並建立
	ch1 := make(chan int)
	ch2 := make(chan string)

	//使用goroutine，有參數
	go func(s int) {
		ch1 <- s
		for i := 1; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}(0)

	//使用goroutine，沒參數
	go func() {
		for i := 10; i < 15; i++ {
			//int to string
			ch2 <- strconv.FormatInt(int64(i), 10)
		}
		close(ch2)
	}()

	//取channel
	for v := range ch1 {
		fmt.Println("int is", v)
	}
	for v := range ch2 {
		fmt.Println("string is", v)
	}

	//建立5個channel buffer，意思是可以quene 5個string
	message = make(chan string, 5)
	go Bot()
	message <- "first message"
	fmt.Println("first message send finish")
	message <- "second message"
	fmt.Println("second message send finish")

	time.Sleep(1000 * time.Millisecond)
}
