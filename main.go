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
	array()
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