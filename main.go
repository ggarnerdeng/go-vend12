package main

import "fmt"

type machine struct {
	model    string
	address  string
	capacity int
}

func (m machine) String() string {
	return "This machine is located at " + m.address
}
func main() {
	var cewmachine = machine{
		model:    "123",
		address:  "100 Mitchell St",
		capacity: 100,
	}

	fmt.Println(cewmachine)
	fmt.Println("Capacity:", cewmachine.capacity)
}
