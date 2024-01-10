package main

import "fmt"

func main() {
	if x := aux(); x < 9 {
		fmt.Println(x)
	} else {
		fmt.Println("Mayor que 9")
	}

	switch x := aux(); x {
	case 5:
		fmt.Println("X igual que 5")
	case 10:
		fmt.Println("X igual que 10")
		fallthrough
	default:
		fmt.Println("Mensaje default")
	}
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	fruits := []string{"apple", "banana", "pear"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

func aux() int {
	return 10
}
