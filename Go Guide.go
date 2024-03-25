package main

// Imported packages
import (
	"fmt"
	"math"
	"strconv"

	"go-guide/strutil"
)

// This is a function
func reverseStr(str string) {
	fmt.Println(strutil.Reverse(str)) // Reverses any string
}


// This is a function
func someMath() {
	fmt.Println(math.Floor(2.7)) // 2
	fmt.Println(math.Ceil(2.7)) // 3
	fmt.Println(math.Sqrt(16)) // 4
}


// This is a struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age int
}

// Greeting method (value reciever (ou seja: não muda nada, só recebe um valor e faz um cálculo))
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever (ou seja: quando chamado, muda algo no map criado))
func (p *Person) hasBirthday() {
	p.age++
}


// This is the main function. Go executes this.
func main() {
	
	fmt.Println()

	// VARIABLES
	var age int = 37
	age = 38
	name := "Brad"
	const profession, email = "Teacher", "brad@gmail.com"
	fmt.Println("Hello", name, "!") // Hello Brad !
	fmt.Print(name, " is a ", age, " years old ", profession, " whose email is ", email, "\n") // Brad is a 38 years old Teacher whose email is brad@gmail.com


	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	
	fmt.Println()


	// FUNCTIONS
	reverseStr(name)
	someMath()

	
	fmt.Println()


	// ARRAYS or SLICES
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr) // [Apple Orange]
	fmt.Println(fruitArr[1]) // Orange

	fruitSlice := []string{"Pinapple", "Strawberry", "Grape", "Cherry"}
	fmt.Println(len(fruitSlice)) // 4
	fmt.Println(fruitSlice[1:3]) // [Strawberry Grape]

	
	fmt.Println()


	// CONDITIONALS
	color := "red"
	
	// else if
	if color == "red" {
		fmt.Println("ELSE IF: color is red")
	} else if color == "blue" {
		fmt.Println("ELSE IF: color is blue")
	} else {
		fmt.Println("ELSE IF: color is NOT blue or red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("SWITCH: color is red")
	case "blue":
		fmt.Println("SWITCH: color is blue")
	default:
		fmt.Println("SWITCH: color is NOT blue or red")
	}


	// LOOP
	for i := 1; i <= 5; i++ { // Number1/ Number2/ Number3/ Number4/ Number5/
		fmt.Print("Number", i, "/ ")
	}


	fmt.Println()
	fmt.Println()


	// MAP (equivalent to JS object)
	// emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	emails := make(map[string]string) // map[KEY type]VALUE type
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails) // map[Bob:bob@gmail.com Mike:mike@gmail.com Sharon:sharon@gmail.com]
	fmt.Println(emails["Bob"]) // bob@gmail.com
	fmt.Println(len(emails)) // 3
	delete(emails, "Bob")
	fmt.Println(emails) // map[Mike:mike@gmail.com Sharon:sharon@gmail.com]


	fmt.Println()


	// RANGE
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Print(i, " - ID: ", id, "/ ") // 0 - ID: 33/ 1 - ID: 76/ 2 - ID: 54/ 3 - ID: 23/ 4 - ID: 11/ 5 - ID: 2/
	}
	fmt.Println()

	// Not using index
	for _, id := range ids {
		fmt.Print("ID: ", id, "/ ") // ID: 33/ ID: 76/ ID: 54/ ID: 23/ ID: 11/ ID: 2/
	}
	fmt.Println()

	// Range with map
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v) // Mike: mike@gmail.com Sharon: sharon@gmail.com
	}

	// Range with map
	for k := range emails {
		fmt.Println("Name: " + k) // Name: Sharon Name: Mike
	}


	fmt.Println()


	// POINTERS
	a := 5
	b := &a // adress of a

	fmt.Println(a, b) // 5 0xc0000121b8
	fmt.Printf("%T\n", b) // typeof b == *int

	//  Use * to read val from address
	fmt.Println(*b) // 5
	// fmt.Println(*&a) // 5 (commented out because it throws a warning)

	// Change val with pointer
	*b = 10
	fmt.Println(a) // 10


	fmt.Println()


	// STRUCT
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}
	fmt.Println(person1) // {Samantha Smith Boston f 25}
	fmt.Println(person2) // {Bob Johnson New York m 30}
	fmt.Println(person1.greet()) // Hello, my name is Samantha Smith and I am 25

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet()) // Hello, my name is Samantha Smith and I am 28

}