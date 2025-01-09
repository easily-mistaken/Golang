package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)



func main() {
	// fmt.Println("hello world")
	// values()
	// variables()
	// constants()
	// loops()
	// ifelse()
	// switchcase()
	// switchAsIfelse()
	// array()
	// slices()
	// test()
	// fmt.Println(fact(7), fib(7))
	rangeExample()
}

func values() {
	fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, reflect.TypeOf(b), c, reflect.TypeOf(c))

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)


	f := "apple"
	fmt.Println(f)

	var sl []int
	var m map[string]int
	var ch chan int
	var fn func()

	fmt.Println(sl, m, ch, fn) // zero values when not initialized
}

const s string = "constant"

func constants() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func loops() {
	i := 1
	for i <= 3 {
		fmt.Println()
		i = i + 1
	}

	for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    for i := range 3 {
        fmt.Println("range", i)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}

func ifelse() {
	if i,j:=5,6; i==5 || j==10 {
		fmt.Println("i is 5")
	} else {
		fmt.Println("i is not 5")
	}
}

func switchcase() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
func switchAsIfelse(){
	num := 10
	switch {
    case num < 5:
        fmt.Println("Less than 5")
    case num == 10:
        fmt.Println("Equals 10")
    case num > 10:
        fmt.Println("Greater than 10")
    }
}
func array(){
	var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2d2: ", twoD2)
}

func slices(){

	var s[]string
	fmt.Println("uninit:", s,s==nil, len(s)==0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
func test(){
	// Slice is a descriptor of an array segment
	// consists of a pointer to the array,
	// the length of the segment,
	// and its capacity

	var a [8]int
	var s = a[0:5]
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))
}

func fact(n int) int{
	if n==0 {return 1}
	return n*fact(n-1)
}

func fib(n int) int{
	if n<2 {return n}
	return fib(n-1)+fib(n-2)
}

func rangeExample(){
	nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for i, c := range "go" {
		fmt.Println(i, string(c))
	}
}