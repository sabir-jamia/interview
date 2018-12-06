package main

import (
	"interview/pkg/bag"
	"fmt"
)

func main() {
	bag := bag.New()
	bag.Add(1)
	bag.Add(2)
	bag.Add(3)
	fmt.Println(bag.Iterate())
}
