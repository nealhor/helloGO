package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Master", age: 24}
	fmt.Println(p)

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	// pointer 7+1
	pp := 7
	inc(&pp)
	fmt.Println(pp)
	//
	var x int = 5 // x:=5
	var y int
	y = 10
	fmt.Println(x, y)
	var sum int = x + y // sum := x + y
	fmt.Println(sum)
	i := 999
	fmt.Println(i)

	if i < 100 {
		fmt.Println("Less that 100")
	} else if x > 2 {
		fmt.Println("x bigger that 2")
	}
	var a [5]int // a:= []int{0, 0, 0, 0, 0,}
	a[2] = 7
	fmt.Println(a) // [ 0 0 7 0 0]

	b := []int{5, 4, 3, 2, 1}
	b = append(b, 13) // add 13 to the end

	fmt.Println(b)

	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

	for g := 0; g < 5; g++ {
		fmt.Println(g)
	}
	// while loop
	summ := 1
	for summ < 5 {
		fmt.Println(summ)
		summ++
	}
	fmt.Println(summ)

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value", value)
	}

	m := make(map[string]string)
	m["a"] = "master"
	m["b"] = "slave"

	for key, value := range m {
		fmt.Println("key:", key, "value", value)
	}
	result := sumint(2, 4)
	fmt.Println("2 + 4 =", result)

	result1, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}
	TTT()
	ZZZ()
}

func sumint(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func inc(x *int) {
	*x++
}
