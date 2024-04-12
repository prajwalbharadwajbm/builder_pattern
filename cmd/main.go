package main

import (
	"buildPattern/pkg/builder"
	"buildPattern/pkg/mestri"
	"fmt"
)

func main() {
	normalBuilder := builder.Get("normal")
	tentBuilder := builder.Get("tent")
	fmt.Println("NORMAL HOUSE")
	normalMestri := mestri.New(normalBuilder)
	normalHouse := normalMestri.BuildHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	fmt.Println("\nTENT HOUSE")
	tentMestri := mestri.New(tentBuilder)
	tentHouse := tentMestri.BuildHouse()
	fmt.Printf("Normal House Door Type: %s\n", tentHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", tentHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", tentHouse.Floor)
}
