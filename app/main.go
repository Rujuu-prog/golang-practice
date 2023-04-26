package main

import (
	"fmt"
	"os"
	"time"
)

// Hello world
/*
複数行のコメント
*/

// i5 := 500
var i5 int = 500

// 定数
const Pi = 3.14
const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const Big = 9223372036854775807 + 1

// iotaは連続した値を生成する
const (
	c0 = iota
	c1
	c2
)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func section1() {
	//明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// エラー
	// i4 := 500
	// fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// fmt.Println(s4)

	var s5 string = "Not Use"
	fmt.Println(s5)
}

func section2Int(){
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i +50)

	// fmt.Println(i+ i2)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}

func section2Float(){
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)

	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	fmt.Println(fl64 + float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}

func section2Bool(){
	var t, f bool = true, false
	fmt.Println(t, f)
}

func section2String(){
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Printf(`test
	test
	 test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))
}

func section2Byte(){
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}

func section2Array(){
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)

	// エラー
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	fmt.Println(len(arr1))
}

func section2Interface(){
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1,2,3}
	fmt.Println(x)

	x = 2
	// fmt.Println(x + 3)
}

func section2Convert(){
	/*
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl64 = %T\n", fl64)

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2)

		fl := 10.5
		i3 := int(fl)
		fmt.Printf("i3 = %T\n", i3)
		fmt.Println(i3)
	*/

	/*
	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)
	*/

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}

func section5Const(){
	fmt.Println(Pi)

	// エラー
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}

func Plus(x, y int)int{
	return x + y
}

func Div(x, y int)(int, int) {
	q := x/y
	r := x%y
	return q, r
}

func Double(price int)(result int){
	result = price * 2
	return}

func Noreturn(){
	fmt.Println("No Return")
	return
}

func ReturnFunc()func(){
	return func(){
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()){
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string{
		s := store
		store = next
		return s
	}
}

func Integers() func() int{
	var i int = 0
	return func() int{
		i++
		return i
	}
}

func SwitchSample(){
	/*
	n := 3
	switch n {
	case 1,2:
		fmt.Println("1,2")
	case 3,4:
		fmt.Println("3,4")
	default:
		fmt.Println("default")
	}
	*/

	/*
	switch n:=3; n {
	case 1,2:
		fmt.Println("1,2")
	case 3,4:
		fmt.Println("3,4")
	default:
		fmt.Println("default")
	}
	*/

	switch n:=5; {
	case n>0 && n<=2:
		fmt.Println("1~2")
	case n>2 && n<=5:
		fmt.Println("3~5")
	default:
		fmt.Println("default")
	}
}


func anything(a interface{}) {
	// fmt.Println(a)
	switch v:=a.(type){
	case string:
		fmt.Println(v+"string")
	case int:
		fmt.Println(v+1000)
	default:
		fmt.Println("default")
	}
}

func SwitchType(){
	anything("aaa")
	anything(1)

	var x interface{} = "e"
	// i, isInt := x.(int)
	// fmt.Println(i+2, isInt)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	if x==nil{
		fmt.Println("nil")
	}else if i, isInt := x.(int); isInt{
		fmt.Println(i, "int")
	}else if f, isFloat64 := x.(float64); isFloat64{
		fmt.Println(f, "float64")
	}else if s, isString := x.(string); isString{
		fmt.Println(s, "string")
	}else {
		fmt.Println("other")
	}

	//上のifよりこっちのswitchを使う方を推奨されてるぽい
	switch x.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}

	switch v:=x.(type){
	case bool:
		fmt.Println(v, "bool")
	case string:
		fmt.Println(v, "string")
	case int:
		fmt.Println(v, "int")
	default:
		fmt.Println("default")
	}
}

func LabelFor(){
	/*
	Loop:
	for {
		for {
			for {
				fmt.Println("Start")
				break Loop
			}
			fmt.Println("処理しない")
		}
		fmt.Println("処理しない2")
	}
	fmt.Println("end")
	*/

	Loop:
	for i:=0;i<3;i++{
		for j:=1;j<3;j++{
			if j>1{
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}

// defer
func TestDefer(){
	defer fmt.Println("END")
	
	defer func(){
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
	}()
	fmt.Println("START")
}

// deferの活用事例
func TestTextCreate(){
	file, err := os.Create("test.txt")
	if err == nil{
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}

func sub(){
	for{
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

// init 基本は一つの関数にまとめる
func init(){
	fmt.Println("init")
}
// func init(){
// 	fmt.Println("init2")
// }

func SliceTest(){
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100,200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1,2,3,4,5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1:len(sl5)-1])
}

func Sum2(s ...int) int{
	var n int = 0
	for _, v := range s{
		n += v
	}
	return n
}

func MapSample(){
	var m = map[string]int{"A":100,"B":200}
	fmt.Println(m)
	m2 := map[string]int{"A":100,"B":200}
	fmt.Println(m2)
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)
	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)
	
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])
	
	_, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(s, ok)
	
	m4[2] = "US"
	fmt.Println(m4)
	m4[3] = "China"
	fmt.Println(m4)
	
	delete(m4, 3)
	fmt.Println(m4)
	
	fmt.Println(len(m4))
}

func MapFor(){
	m := map[string]int{
		"Apple": 100,
		"banana": 200,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// channel
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言、操作
func ChannelSample(){
	//指定しない場合は双方向のチャンネルになる
	var ch1 chan int

	//受信専用
	// var ch2 <-chan int

	//送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))
	
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))
	
	ch3 <- 1
	fmt.Println(len(ch3))
	ch3 <-2
	ch3 <-3
	fmt.Println("len",len(ch3))
	
	i := <-ch3
	fmt.Println(i)
	fmt.Println("len",len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len",len(ch3))
	
	fmt.Println(<-ch3)
	fmt.Println("len",len(ch3))
	
	
	fmt.Println("#-------")
	ch3<-1
	fmt.Println("len",len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len",len(ch3))
	ch3<-2
	ch3<-3
	ch3<-4
	ch3<-5
	ch3<-6

	fmt.Println(len(ch3))

}

func main() {
	// section1()
	// section2Int()
	// section2Float()
	// section2Bool()
	// section2String()
	// section2Byte()
	// section2Array()
	// section2Interface()
	// section2Convert()
	// section5Const()
	// i := Plus(1, 2)
	// fmt.Println(i)
	// i2, _ := Div(9, 4)
	// fmt.Println(i2)
	// i4 := Double(1000)
	// fmt.Println(i4)
	// Noreturn()

	//　無名関数
	// f := func(x,y int) int {
	// 	return x+y
	// }
	// i := f(1, 2)
	// fmt.Println(i)

	// i2 := func(x,y int)int{
	// 	return x*y
	// }(2,3)
	// fmt.Println(i2)

	// f := ReturnFunc()
	// f()
	// CallFunction(func() {
	// 	fmt.Println("function.")
	// })

	//クロージャー
	// f :=Later()
	// fmt.Println(f("Hello"))
	// fmt.Println(f("My"))
	// fmt.Println(f("name"))

	// 擬似ジェネレーター
	// f := Integers()
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

	// otherints := Integers()
	// fmt.Println(otherints())
	// fmt.Println(otherints())
	// fmt.Println(otherints())
	// fmt.Println(otherints())
	// fmt.Println(otherints())

	// 簡易文if
	// x := 0
	// if x:=2;x==2 {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// 無限ループ
	// for {
	//  fmt.Println("loop")
	// }

	// whileのような
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// 通常のやつ
	// for i:=0;i<10;i++{
	// 	if i==3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// 配列とインデックスを合わせて回す
	// arr := [3]int{1,2,3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }
	
	// SwitchSample()
	// SwitchType()
	// LabelFor()

	// TestDefer()
	// TestTextCreate()

	// go goroutin
	// go sub()
	// go sub()
	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	//init
	// fmt.Println("Main")

	// SliceTest()

	// fmt.Println(Sum2(1,2,3))
	// var sl []int = []int{1,2,3,4}
	// fmt.Println(Sum2(sl...))
	// MapSample()
	// MapFor()

	ChannelSample()
}