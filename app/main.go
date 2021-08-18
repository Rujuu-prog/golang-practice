package main

import (
	"fmt"
	// "time"
)

//関数外での変数定義
//i5 := 500 これは動かない
var i5 int = 500

func outer(){
	var s4 string = "outer"
	fmt.Println(s4)
}

func main(){
	//HelloWorld
	/*
	fmt.Println("Hello World")
	fmt.Println(time.Now())
	*/

	//変数
	//明示的な定義
	//var 変数名 型 = 値
	/*
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Golang"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t,f)

	var (
		i2 int = 200
		s2 string = "Go Running"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Golang"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)
	*/

	//暗黙的な定義
	//変数名 := 値
	/*
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
	*/

	//型
	//整数型
	/*
	var i int = 100
	fmt.Println(i + 50)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", int32(i))
	*/
	//浮動小数点
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)
	ninf := -1.0 / zero
	fmt.Println(ninf)
	nan := zero / zero
	fmt.Println(nan)
}