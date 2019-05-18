package main

//並行執行。只要main thread(主執行緒)中斷或結束，所有的子執行緒(goroutine)都會中斷。

import (
	"fmt"
	"sync"
	"time"
)

var money = 1500
var wg sync.WaitGroup

//mutex鎖頭
var mu sync.Mutex

func withdraw() {
	//上鎖
	mu.Lock()
	balance := money
	time.Sleep(1000 * time.Millisecond)
	balance -= 1000
	money = balance
	//解鎖
	mu.Unlock()

	//顯示結果
	fmt.Println("After withdrawing $1000, balace: ", money)
	wg.Done()
}

func main() {
	fmt.Println("We have $1500")
	wg.Add(2)
	go withdraw()
	go withdraw()
	wg.Wait()
}
