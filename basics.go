package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// variable declaration
	var age int
	age = 21

	num := 7
	var b float64 = 3.14

	name := "vishal"
	var n int = 4

	fmt.Println("my age is:", age, "my name is:", name, "pi is", b)

	//if statement
	if age > 23 {
		fmt.Println("you can drink")
	} else {
		fmt.Println("you cannot drink")
	}

	if num%2 == 0 {
		fmt.Println("divisble bt 2")
	} else {
		fmt.Println("not dividble by 2")
	}

	//for loop
	for i := 0; i < n; i++ {
		fmt.Println("hello")
	}

	arr := []int{10, 20, 30}
	var sum int
	for _, j := range arr {
		sum += j
	}
	fmt.Println("sum is:", sum)

	for i := 0; i < 5; i++ {
		if i == 4 {
			break
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println(" ")
	for i := 1; i <= 5; i++ {
		if i == 4 {
			continue
		}
		fmt.Print(i)
	}

	fmt.Println("")

	//switch-case
	var day string = "mon"

	switch day {
	case "mon":
		fmt.Println("happ mon")
		break
	case "tue":
		fmt.Println("happy tue")
		break
	case "wed":
		fmt.Println("happy wed")
		break
	default:
		fmt.Println("invalid day")
	}

	//functions

	mult := add(5, 6)
	fmt.Println("multiplied:", mult)

	//error
	result, er := divison(10, 0)
	fmt.Println("result is:", result)
	if er != nil {
		fmt.Println("error:", er)
	}
	//struct->user defined datatype
	type person struct {
		name string
		age  int
		city string
	}

	//creating instance for this struct
	p := person{"vishal", 22, "kochi"}
	fmt.Println(p.name, p.age)

	//lets create instance for the array
	var brr []person
	brr = append(brr, person{"anaz", 22, " kozhikode"})
	brr = append(brr, person{"nihal", 21, "kuttiyadi"})
	fmt.Println(brr[1].name)

	//different ways to update struct
	brr[1] = person{"vishal", 21, "trivandrum"}
	fmt.Println(brr[1].name, brr[1].age)

	//maping
	//string->key int->value
	ages1 := make(map[string]int) //creating memory
	ages1["vishal"] = 22
	ages1["yedhu"] = 25

	fmt.Println(ages1["vishal"]) //accessing map value
	value, ok := ages1["vishal"] //map returns two values
	if ok {
		fmt.Println("found:", ok, "value", value)
	} else {
		fmt.Println("value not found")
	}
	//itrating through value
	for key, value := range ages1 {
		fmt.Println(key, value)
	}

}

func add(a int, b int) int {
	mult := a * b
	return mult
}
func divison(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator can't zero")
	}
	div := a / b
	return div, nil
}
