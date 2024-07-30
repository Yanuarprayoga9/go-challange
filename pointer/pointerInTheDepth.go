package main

import "fmt"

type person struct {
	name    string
	age     int
	city    string
	country string
}

func (psn *person) ifMeet(action string) {
    if action == "greet" {
        fmt.Printf("Hello, my name is %s. Nice to meet you!\n", psn.name)
    } else if action == "bye" {
        fmt.Printf("Goodbye from %s. See you next time!\n", psn.name)
    } else {
        fmt.Printf("%s doesn't understand the action: %s\n", psn.name, action)
    }
}

func edit(a *int) {
	*a = 10 // Change the value of the integer that the pointer points to
}

func editPerson(o *person) {
	o.age = o.age + 3
}
func main() {
	yanuar := person{"yanuar", 21, "kebumen", "indonesia "}
	choerul := person{"choerul", 21, "cilacap", "indonesia "}

	personList := []person{
		yanuar,
		choerul,
	}

	for _, p := range personList {
		fmt.Println(p)

	}
	editPerson(&yanuar)
	editPerson(&choerul)
	yanuar.ifMeet("greet")  // Output: Hello, my name is Yanuar. Nice to meet you!
    choerul.ifMeet("bye")   // Output: Goodbye from Choerul. See you next time!
    yanuar.ifMeet("dance")  	// yanuar  := 1;
	// dimas  := 1;

	// edit (&yanuar)
	// edit (&dimas)
	fmt.Println(yanuar)
	fmt.Println(choerul)
}
