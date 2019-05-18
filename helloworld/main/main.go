package main

import (
	"errors"
	"fmt"
	_ "golangLearn/helloworld/init"
	"golangLearn/helloworld/src"
	"strings"
	"time"
)

//變數宣告-全域
//第一種
var str1 string
var num1 int

//第二種
var (
	str2 string = "str456"
	num2 int    = 456
)

//常數宣告
//第一種
const (
	Monday1    = 1
	Tuesday1   = 2
	Wednesday1 = 3
)

//第二種
//iota預設值為0，後面的會自動+1
const (
	Monday2 = iota
	Tuesday2
	Wednesday2
)

//主程式
func main() {
	//變數宣告-區域(建議使用)
	str3 := "str123456"
	num3 := 123456

	str1 = "str123"
	num1 = 123

	//變數結果
	fmt.Println(str1)
	fmt.Println(num1)
	fmt.Println(str2)
	fmt.Println(num2)
	fmt.Println(str3)
	fmt.Println(num3)

	//常數結果
	fmt.Println(Monday1)
	fmt.Println(Tuesday1)
	fmt.Println(Wednesday1)
	fmt.Println(Monday2)
	fmt.Println(Tuesday2)
	fmt.Println(Wednesday2)

	//不同package函式呼叫結果
	fmt.Println(hello.HelloWorld())

	//數值相加結果
	fmt.Println(add(5, 5))

	//數值互換結果
	a := "A"
	b := "B"
	a, b = swap(a, b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	//數值互換簡寫結果
	c := "C"
	d := "D"
	c, d = d, c
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	//回傳函式結果
	test1 := callFunc()
	fmt.Printf("類型: %T\n", test1)
	fmt.Println(test1())

	//變數宣告成函式
	foo := func(i, j float32) float32 {
		return i + j
	}
	fmt.Printf("類型: %T\n", foo)
	fmt.Println(foo(5.5, 5.0))

	//anonymous(匿名函式) - 第一種
	anyfunc := func() {
		fmt.Println("匿名函式呼叫1")
	}
	anyfunc()

	//anonymous(匿名函式) - 第二種
	func() {
		fmt.Println("匿名函式呼叫2")
	}()

	//利用關鍵字go背景執行函式(如果函式在main最後才執行，會看不到結果。因為main已經結束)
	go func(i, j int) {
		fmt.Println(i + j)
	}(3, 3)

	//產生SQL函式結果
	fmt.Println(getUserListSQL("jk", ""))
	fmt.Println(getUserListSQL("jk", "10663126@gm.nfu.edu.tw"))

	//產生SQL函式結果(struct版)
	fmt.Println("struct版-1: " + getUserListOptsSQL(SearchOpts{
		username: "jk",
		email:    "10663126@gm.nfu.edu.tw",
	}))
	fmt.Println("struct版-2: " + getUserListOptsSQL(SearchOpts{
		sexy: 2,
	}))

	//錯誤處理結果
	if _, err := checkUserNameExist("jk"); err != nil {
		fmt.Println(err)
	}
	if _, err := checkUserNameExist("andy"); err != nil {
		fmt.Println(err)
	}
	if _, err := checkUserNameExist("user"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		}
	}
	if _, err := checkUserNameExist("error"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("isErrUserNameExist is false")
		}
	}

	//pass slice as function argument
	fooarr := []string{"a", "b"}
	fmt.Println("before fooarr: ", fooarr)
	modify(fooarr)
	fmt.Println("after fooarr: ", fooarr)
	modifyAddValue(fooarr)
	fmt.Println("add after fooarr: ", fooarr)
	//fooarr = modifyAddReturnValue(fooarr)
	//fmt.Println("add after fooarr: ", fooarr)
	modifyAddAddress(&fooarr)
	fmt.Println("add after fooarr: ", fooarr)
	fooarr = modifyAddReturnValue(fooarr)
	fmt.Println("return value after fooarr: ", fooarr)
	bar := fooarr[:1]
	fmt.Println("bar: ", bar)
	s1 := append(bar, "c")
	fmt.Println("fooarr: ", fooarr)
	fmt.Println("s1: ", s1)
	// 因為沒有超出bar的最大上限，所以s2還是保持2個值
	s2 := append(bar, "d")
	fmt.Println("fooarr: ", fooarr)
	fmt.Println("s2: ", s2)
	// 因為超出bar的最大上限，所以變複製成新的s3
	s3 := append(bar, "e", "f")
	fmt.Println("fooarr: ", fooarr)
	fmt.Println("s3: ", s3)

	//switch case
	testSwitch(0)
	testSwitch(1)
	testSwitch(2)
	testSwitch(3)

	//Struct Method 的 Pointers vs. Values 差異
	val := &Car{}
	fmt.Printf("val address: %p\n", val)
	//傳值，複製一個新的(struct量多時容易浪費記憶體)
	val.setName1("123")
	fmt.Println(val.name)
	//傳址，指標指向舊的
	val.setName2("456")
	fmt.Println(val.name)

	//struct寄信例子(使用指標)
	e := &Email{}
	//循序
	for i := 0; i < 10; i++ {
		e.setFrom(fmt.Sprintf("user%02d@test.com", i)).
			setTo(fmt.Sprintf("user%02d@test.com", i+1)).
			sendMail()
	}
	//並行處理(goroutine)
	for i := 0; i < 10; i++ {
		go func(i int) {
			e.setFrom(fmt.Sprintf("goroutine%02d@test.com", i)).
				setTo(fmt.Sprintf("goroutine%02d@test.com", i+1)).
				sendMail()
		}(i)
	}

	time.Sleep(1 * time.Second)
}

//相加函式
func add(i, j int) int {
	return i + j
}

//互換函式
func swap(i, j string) (string, string) {
	return j, i
}

//回傳函式
func callFunc() func() int {
	return func() int {
		return 100
	}
}

//產生SQL函式
func getUserListSQL(username, email string) string {
	sql := "select * from user"
	where := []string{}

	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}
	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

//產生SQL函式(struct版)
type SearchOpts struct {
	username string
	email    string
	sexy     int
}

func getUserListOptsSQL(opts SearchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}
	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}
	if opts.sexy != 0 {
		where = append(where, fmt.Sprintf("sexy = '%d'", opts.sexy))
	}

	return sql + " where " + strings.Join(where, " or ")
}

//錯誤處理
type ErrUserNameExist struct {
	user string
}

func (e ErrUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.user)
}

func isErrUserNameExist(err error) bool {
	_, ok := err.(ErrUserNameExist)
	return ok
}

func checkUserNameExist(username string) (bool, error) {
	if username == "jk" {
		return true, fmt.Errorf("username %s already exist", username)
	}

	if username == "error" {
		return true, errors.New("username error already exist")
	}

	// 常用方式errors.New
	if username == "andy" {
		return true, errors.New("username andy already exist")
	}

	// 常用方式，透過struct
	if username == "user" {
		return true, ErrUserNameExist{user: username}
	}

	return false, nil
}

//pass slice as function argument
func modify(fooarr []string) {
	fooarr[1] = "c"
	fmt.Println("modify fooarr: ", fooarr)
}

//傳值
func modifyAddValue(fooarr []string) {
	fooarr = append(fooarr, "c")
	fmt.Println("modify fooarr: ", fooarr)
}

//傳址
func modifyAddAddress(fooarr *[]string) {
	*fooarr = append(*fooarr, "c")
	fmt.Println("modify fooarr: ", fooarr)
}

//傳值(透過回傳值方式)
func modifyAddReturnValue(fooarr []string) []string {
	fooarr = append(fooarr, "d")
	fmt.Println("return modify fooarr: ", fooarr)
	return fooarr
}

//switch case
func testSwitch(i int) {
	switch i {
	case 0:
		fallthrough //強制執行後面的程式
	case 1:
		fmt.Println("select value is ", i)
	case 2, 3: //2 or 3都會進來
		fmt.Println("select value is ", i)
	}
}

//Struct Method 的 Pointers vs. Values 差異
type Car struct {
	name string
}

func (val Car) setName1(s string) {
	//複製(struct量多時容易浪費記憶體)
	fmt.Printf("setName1 address: %p\n", &val)
	val.name = s
}

func (val *Car) setName2(s string) {
	//指標
	fmt.Printf("setName2 address: %p\n", val)
	val.name = s
}

//struct寄信例子(使用指標)
type Email struct {
	from string
	to   string
}

//設定from
func (e Email) setFrom(str string) Email {
	e.from = str
	return e
}

//設定to
func (e Email) setTo(str string) Email {
	e.to = str
	return e
}

//傳送郵件
func (e Email) sendMail() {
	fmt.Printf("From: %s, To: %s\n", e.from, e.to)
}
