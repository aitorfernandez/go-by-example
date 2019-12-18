package main

import "fmt"

func main() {
	_, err := returnID()

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Println(greet("Hello", "Go!"))
	fmt.Printf("greet func type %T\n", greet)

	// Self execution
	func(i int) {
		fmt.Printf("I'm self-execution! %v\n", i)
	}(1)

	// Dynamic params
	avg := average(43, 5, 87, 12, 3)
	// data := []float64{43, 5, 87, 12, 3}
	// avg := average(data...)
	// avg := average(data)
	fmt.Printf("dynamic params ... %v\n", avg)

	printWithCallback([]int{1, 2, 3, 4, 5}, func(i int) {
		fmt.Printf("in callback %v\n", i)
	})

	// Panic and Recover
	defer func() {
		str := recover()
		fmt.Println(str) // panic message
	}()
	panic("PANIC!")
}

// func average(values []float64) float64 { ... }
func average(values ...float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}

// return naming
func returnID() (id int, err error) {
	// declare the same variable name twice, no new variables on left side of
	id = 10 // id := 10
	if id == 10 {
		// scoped the err variable
		// err := fmt.Errorf("Invalid id") // err is shadowed during return we are not using err in the definition
		err = fmt.Errorf("Invalid id")
		return
	}
	return
}

// Multiple return, params
func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

// return func
func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func printWithCallback(values []int, callback func(int)) {
	for _, i := range values {
		callback(i)
	}
}
